/*
 * Copyright (c) 2025, NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dgxh200

import (
	"fmt"
	"sync"

	"github.com/google/uuid"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock"
)

type Server struct {
	mock.Interface
	mock.ExtendedInterface
	Devices           [8]nvml.Device
	DriverVersion     string
	NvmlVersion       string
	CudaDriverVersion int
}

type Device struct {
	mock.Device
	sync.RWMutex
	UUID                  string
	Name                  string
	Brand                 nvml.BrandType
	Architecture          nvml.DeviceArchitecture
	PciBusID              string
	Minor                 int
	Index                 int
	CudaComputeCapability CudaComputeCapability
	MigMode               int
	GpuInstances          map[*GpuInstance]struct{}
	GpuInstanceCounter    uint32
	MemoryInfo            nvml.Memory
	// H200+ specific: Enhanced fabric capabilities with more memory
	FabricManagerEnabled bool
	FabricPartitionID    int
}

type GpuInstance struct {
	mock.GpuInstance
	sync.RWMutex
	Info                   nvml.GpuInstanceInfo
	ComputeInstances       map[*ComputeInstance]struct{}
	ComputeInstanceCounter uint32
	// H200+ specific: Can be part of fabric partitions with enhanced memory
	FabricPartitionID int
}

type ComputeInstance struct {
	mock.ComputeInstance
	Info nvml.ComputeInstanceInfo
}

type CudaComputeCapability struct {
	Major int
	Minor int
}

var _ nvml.Interface = (*Server)(nil)
var _ nvml.Device = (*Device)(nil)
var _ nvml.GpuInstance = (*GpuInstance)(nil)
var _ nvml.ComputeInstance = (*ComputeInstance)(nil)

func New() *Server {
	server := &Server{
		Devices: [8]nvml.Device{
			NewDevice(0),
			NewDevice(1),
			NewDevice(2),
			NewDevice(3),
			NewDevice(4),
			NewDevice(5),
			NewDevice(6),
			NewDevice(7),
		},
		DriverVersion:     "550.54.15",
		NvmlVersion:       "12.550.54.15",
		CudaDriverVersion: 12040,
	}
	server.setMockFuncs()
	return server
}

func NewDevice(index int) *Device {
	device := &Device{
		UUID:         "GPU-" + uuid.New().String(),
		Name:         "Mock NVIDIA H200-SXM5-141GB",
		Brand:        nvml.BRAND_NVIDIA,
		Architecture: nvml.DEVICE_ARCH_HOPPER, // Hopper architecture (same as H100)
		PciBusID:     fmt.Sprintf("0000:%02d:00.0", index),
		Minor:        index,
		Index:        index,
		CudaComputeCapability: CudaComputeCapability{
			Major: 9, // H200 CUDA compute capability (same as H100)
			Minor: 0,
		},
		MigMode:              nvml.DEVICE_MIG_ENABLE,
		GpuInstances:         make(map[*GpuInstance]struct{}),
		GpuInstanceCounter:   0,
		MemoryInfo:           nvml.Memory{Total: 151397597184}, // 141GB for H200 (vs 80GB for H100)
		FabricManagerEnabled: true,                             // H200+ supports FabricManager
		FabricPartitionID:    -1,                               // Not assigned to fabric partition initially
	}
	return device
}

func (d *Device) GetUUID() (string, nvml.Return) {
	d.RLock()
	defer d.RUnlock()
	return d.UUID, nvml.SUCCESS
}

func (d *Device) GetName() (string, nvml.Return) {
	d.RLock()
	defer d.RUnlock()
	return d.Name, nvml.SUCCESS
}

func (d *Device) GetBrand() (nvml.BrandType, nvml.Return) {
	d.RLock()
	defer d.RUnlock()
	return d.Brand, nvml.SUCCESS
}

func (d *Device) GetArchitecture() (nvml.DeviceArchitecture, nvml.Return) {
	d.RLock()
	defer d.RUnlock()
	return d.Architecture, nvml.SUCCESS
}

func (d *Device) GetPciInfo() (nvml.PciInfo, nvml.Return) {
	d.RLock()
	defer d.RUnlock()
	var busId [32]uint8
	copy(busId[:], d.PciBusID)
	return nvml.PciInfo{
		BusId: busId,
	}, nvml.SUCCESS
}

func (d *Device) GetMinorNumber() (int, nvml.Return) {
	d.RLock()
	defer d.RUnlock()
	return d.Minor, nvml.SUCCESS
}

func (d *Device) GetIndex() (int, nvml.Return) {
	d.RLock()
	defer d.RUnlock()
	return d.Index, nvml.SUCCESS
}

func (d *Device) GetCudaComputeCapability() (int, int, nvml.Return) {
	d.RLock()
	defer d.RUnlock()
	return d.CudaComputeCapability.Major, d.CudaComputeCapability.Minor, nvml.SUCCESS
}

func (d *Device) GetMigMode() (int, int, nvml.Return) {
	d.RLock()
	defer d.RUnlock()
	return d.MigMode, d.MigMode, nvml.SUCCESS
}

func (d *Device) GetMemoryInfo() (nvml.Memory, nvml.Return) {
	d.RLock()
	defer d.RUnlock()
	return d.MemoryInfo, nvml.SUCCESS
}

func (d *Device) GetGpuInstanceProfileInfo(profile int) (nvml.GpuInstanceProfileInfo, nvml.Return) {
	d.RLock()
	defer d.RUnlock()

	// Use H200+ specific MIG profiles with enhanced memory
	if profileInfo, exists := MIGProfiles.GpuInstanceProfiles[profile]; exists {
		return profileInfo, nvml.SUCCESS
	}
	return nvml.GpuInstanceProfileInfo{}, nvml.ERROR_NOT_SUPPORTED
}

func (d *Device) GetGpuInstanceProfileInfoFunc(profile int) (nvml.GpuInstanceProfileInfo, nvml.Return) {
	return d.GetGpuInstanceProfileInfo(profile)
}

func (d *Device) CreateGpuInstance(profileInfo *nvml.GpuInstanceProfileInfo) (nvml.GpuInstance, nvml.Return) {
	d.Lock()
	defer d.Unlock()

	// Create a new GPU instance
	gi := &GpuInstance{
		Info: nvml.GpuInstanceInfo{
			Device:    d,
			Id:        d.GpuInstanceCounter,
			ProfileId: profileInfo.Id,
			Placement: nvml.GpuInstancePlacement{},
		},
		ComputeInstances:       make(map[*ComputeInstance]struct{}),
		ComputeInstanceCounter: 0,
		FabricPartitionID:      d.FabricPartitionID, // H200+ specific
	}

	d.GpuInstances[gi] = struct{}{}
	d.GpuInstanceCounter++

	return gi, nvml.SUCCESS
}

func (d *Device) CreateGpuInstanceFunc(profileInfo *nvml.GpuInstanceProfileInfo) (nvml.GpuInstance, nvml.Return) {
	return d.CreateGpuInstance(profileInfo)
}

func (d *Device) CreateGpuInstanceWithPlacement(profileInfo *nvml.GpuInstanceProfileInfo, placement *nvml.GpuInstancePlacement) (nvml.GpuInstance, nvml.Return) {
	d.Lock()
	defer d.Unlock()

	// Create a new GPU instance with placement
	gi := &GpuInstance{
		Info: nvml.GpuInstanceInfo{
			Device:    d,
			Id:        d.GpuInstanceCounter,
			ProfileId: profileInfo.Id,
			Placement: *placement,
		},
		ComputeInstances:       make(map[*ComputeInstance]struct{}),
		ComputeInstanceCounter: 0,
		FabricPartitionID:      d.FabricPartitionID, // H200+ specific
	}

	d.GpuInstances[gi] = struct{}{}
	d.GpuInstanceCounter++

	return gi, nvml.SUCCESS
}

func (d *Device) CreateGpuInstanceWithPlacementFunc(profileInfo *nvml.GpuInstanceProfileInfo, placement *nvml.GpuInstancePlacement) (nvml.GpuInstance, nvml.Return) {
	return d.CreateGpuInstanceWithPlacement(profileInfo, placement)
}

func (d *Device) GetGpuInstances(profileInfo *nvml.GpuInstanceProfileInfo) ([]nvml.GpuInstance, nvml.Return) {
	d.RLock()
	defer d.RUnlock()

	var instances []nvml.GpuInstance
	for gi := range d.GpuInstances {
		instances = append(instances, gi)
	}
	return instances, nvml.SUCCESS
}

func (d *Device) GetGpuInstancesFunc(profileInfo *nvml.GpuInstanceProfileInfo) ([]nvml.GpuInstance, nvml.Return) {
	return d.GetGpuInstances(profileInfo)
}

func (d *Device) DeleteGpuInstance(gi nvml.GpuInstance) nvml.Return {
	d.Lock()
	defer d.Unlock()

	// Find and remove the GPU instance
	for gpuInstance := range d.GpuInstances {
		if gpuInstance == gi {
			delete(d.GpuInstances, gpuInstance)
			return nvml.SUCCESS
		}
	}
	return nvml.ERROR_NOT_FOUND
}

func (d *Device) DeleteGpuInstanceFunc(gi nvml.GpuInstance) nvml.Return {
	return d.DeleteGpuInstance(gi)
}

// H200+ specific: FabricManager integration methods
func (d *Device) IsFabricManagerEnabled() bool {
	d.RLock()
	defer d.RUnlock()
	return d.FabricManagerEnabled
}

func (d *Device) GetFabricPartitionID() int {
	d.RLock()
	defer d.RUnlock()
	return d.FabricPartitionID
}

func (d *Device) SetFabricPartitionID(partitionID int) {
	d.Lock()
	defer d.Unlock()
	d.FabricPartitionID = partitionID
}

// GPU Instance methods
func (gi *GpuInstance) GetInfo() (nvml.GpuInstanceInfo, nvml.Return) {
	gi.RLock()
	defer gi.RUnlock()
	return gi.Info, nvml.SUCCESS
}

func (gi *GpuInstance) GetInfoFunc() (nvml.GpuInstanceInfo, nvml.Return) {
	return gi.GetInfo()
}

func (gi *GpuInstance) GetComputeInstanceProfileInfo(profile int, slice int) (nvml.ComputeInstanceProfileInfo, nvml.Return) {
	gi.RLock()
	defer gi.RUnlock()

	if profiles, exists := MIGProfiles.ComputeInstanceProfiles[int(gi.Info.ProfileId)]; exists {
		if profileInfo, exists := profiles[profile]; exists {
			return profileInfo, nvml.SUCCESS
		}
	}
	return nvml.ComputeInstanceProfileInfo{}, nvml.ERROR_NOT_SUPPORTED
}

func (gi *GpuInstance) GetComputeInstanceProfileInfoFunc(profile int, slice int) (nvml.ComputeInstanceProfileInfo, nvml.Return) {
	return gi.GetComputeInstanceProfileInfo(profile, slice)
}

func (gi *GpuInstance) CreateComputeInstance(profileInfo *nvml.ComputeInstanceProfileInfo) (nvml.ComputeInstance, nvml.Return) {
	gi.Lock()
	defer gi.Unlock()

	ci := &ComputeInstance{
		Info: nvml.ComputeInstanceInfo{
			Device:      gi.Info.Device,
			GpuInstance: gi,
			Id:          gi.ComputeInstanceCounter,
			ProfileId:   profileInfo.Id,
			Placement:   nvml.ComputeInstancePlacement{},
		},
	}

	gi.ComputeInstances[ci] = struct{}{}
	gi.ComputeInstanceCounter++

	return ci, nvml.SUCCESS
}

func (gi *GpuInstance) CreateComputeInstanceFunc(profileInfo *nvml.ComputeInstanceProfileInfo) (nvml.ComputeInstance, nvml.Return) {
	return gi.CreateComputeInstance(profileInfo)
}

func (gi *GpuInstance) GetComputeInstances(profileInfo *nvml.ComputeInstanceProfileInfo) ([]nvml.ComputeInstance, nvml.Return) {
	gi.RLock()
	defer gi.RUnlock()

	var instances []nvml.ComputeInstance
	for ci := range gi.ComputeInstances {
		instances = append(instances, ci)
	}
	return instances, nvml.SUCCESS
}

func (gi *GpuInstance) GetComputeInstancesFunc(profileInfo *nvml.ComputeInstanceProfileInfo) ([]nvml.ComputeInstance, nvml.Return) {
	return gi.GetComputeInstances(profileInfo)
}

func (gi *GpuInstance) DeleteComputeInstance(ci nvml.ComputeInstance) nvml.Return {
	gi.Lock()
	defer gi.Unlock()

	for computeInstance := range gi.ComputeInstances {
		if computeInstance == ci {
			delete(gi.ComputeInstances, computeInstance)
			return nvml.SUCCESS
		}
	}
	return nvml.ERROR_NOT_FOUND
}

func (gi *GpuInstance) DeleteComputeInstanceFunc(ci nvml.ComputeInstance) nvml.Return {
	return gi.DeleteComputeInstance(ci)
}

// H200+ specific: FabricManager integration for GPU instances
func (gi *GpuInstance) GetFabricPartitionID() int {
	gi.RLock()
	defer gi.RUnlock()
	return gi.FabricPartitionID
}

func (gi *GpuInstance) SetFabricPartitionID(partitionID int) {
	gi.Lock()
	defer gi.Unlock()
	gi.FabricPartitionID = partitionID
}

// Compute Instance methods
func (ci *ComputeInstance) GetInfo() (nvml.ComputeInstanceInfo, nvml.Return) {
	return ci.Info, nvml.SUCCESS
}

func (ci *ComputeInstance) GetInfoFunc() (nvml.ComputeInstanceInfo, nvml.Return) {
	return ci.GetInfo()
}

// Server methods
func (s *Server) DeviceGetCount() (int, nvml.Return) {
	return len(s.Devices), nvml.SUCCESS
}

func (s *Server) DeviceGetCountFunc() (int, nvml.Return) {
	return s.DeviceGetCount()
}

func (s *Server) DeviceGetHandleByIndex(index int) (nvml.Device, nvml.Return) {
	if index < 0 || index >= len(s.Devices) {
		return nil, nvml.ERROR_INVALID_ARGUMENT
	}
	return s.Devices[index], nvml.SUCCESS
}

func (s *Server) DeviceGetHandleByIndexFunc(index int) (nvml.Device, nvml.Return) {
	return s.DeviceGetHandleByIndex(index)
}

func (s *Server) DeviceGetHandleByUUID(uuid string) (nvml.Device, nvml.Return) {
	for _, device := range s.Devices {
		if d, ok := device.(*Device); ok {
			if d.UUID == uuid {
				return device, nvml.SUCCESS
			}
		}
	}
	return nil, nvml.ERROR_NOT_FOUND
}

func (s *Server) DeviceGetHandleByUUIDFunc(uuid string) (nvml.Device, nvml.Return) {
	return s.DeviceGetHandleByUUID(uuid)
}

func (s *Server) DeviceGetTopologyNearestGpus(device nvml.Device, gpuTopologyLevel nvml.GpuTopologyLevel) ([]nvml.Device, nvml.Return) {
	// H200+ enhanced fabric topology
	devices := []nvml.Device{}
	for i := 1; i < len(s.Devices); i++ {
		devices = append(devices, s.Devices[i])
	}
	return devices, nvml.SUCCESS
}

func (s *Server) DeviceGetTopologyNearestGpusFunc(device nvml.Device, gpuTopologyLevel nvml.GpuTopologyLevel) ([]nvml.Device, nvml.Return) {
	return s.DeviceGetTopologyNearestGpus(device, gpuTopologyLevel)
}

func (s *Server) setMockFuncs() {
	// Set up minimal mock function implementations
	s.InitFunc = func() nvml.Return {
		return nvml.SUCCESS
	}

	s.ShutdownFunc = func() nvml.Return {
		return nvml.SUCCESS
	}
}
