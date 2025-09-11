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

package shared

import (
	"fmt"
	"sync"

	"github.com/google/uuid"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock"
)

// GPUSpec defines the configuration for a specific GPU type
type GPUSpec struct {
	Name                string
	Architecture        nvml.DeviceArchitecture
	Brand               nvml.BrandType
	PciDeviceId         uint32
	PciBusIDBase        string // Base format for PCI Bus ID (e.g., "0000:17:%02x:00.0")
	DriverVersion       string
	NvmlVersion         string
	CudaDriverVersion   int
	CudaCapabilityMajor int
	CudaCapabilityMinor int
	TotalMemoryMB       uint64
	DeviceCount         int
	MIGProfiles         MIGProfileConfig
}

// MIGProfileConfig holds MIG profile specifications for a GPU type
type MIGProfileConfig struct {
	GpuInstanceProfiles       map[int]nvml.GpuInstanceProfileInfo
	ComputeInstanceProfiles   map[int]map[int]nvml.ComputeInstanceProfileInfo
	GpuInstancePlacements     map[int][]nvml.GpuInstancePlacement
	ComputeInstancePlacements map[int]map[int][]nvml.ComputeInstancePlacement
}

// MockGPU provides a unified base implementation for all GPU mock types
type MockGPU struct {
	mock.Interface
	mock.ExtendedInterface
	Spec    GPUSpec
	Devices []nvml.Device
}

// MockDevice represents a single GPU device with configurable specifications
type MockDevice struct {
	mock.Device
	sync.RWMutex
	Spec               GPUSpec
	UUID               string
	PciBusID           string
	Minor              int
	Index              int
	MigMode            int
	GpuInstances       map[*MockGpuInstance]struct{}
	GpuInstanceCounter uint32
	MemoryInfo         nvml.Memory
}

// MockGpuInstance represents a GPU instance with shared implementation
type MockGpuInstance struct {
	mock.GpuInstance
	sync.RWMutex
	Info                   nvml.GpuInstanceInfo
	ComputeInstances       map[*MockComputeInstance]struct{}
	ComputeInstanceCounter uint32
	MIGProfiles            MIGProfileConfig
}

// MockComputeInstance represents a compute instance with shared implementation
type MockComputeInstance struct {
	mock.ComputeInstance
	Info nvml.ComputeInstanceInfo
}

// Ensure interface compliance
var _ nvml.Interface = (*MockGPU)(nil)
var _ nvml.Device = (*MockDevice)(nil)
var _ nvml.GpuInstance = (*MockGpuInstance)(nil)
var _ nvml.ComputeInstance = (*MockComputeInstance)(nil)

// NewMockGPU creates a new GPU mock with the specified configuration
func NewMockGPU(spec GPUSpec) *MockGPU {
	gpu := &MockGPU{
		Spec:    spec,
		Devices: make([]nvml.Device, spec.DeviceCount),
	}

	// Initialize devices
	for i := 0; i < spec.DeviceCount; i++ {
		gpu.Devices[i] = NewMockDevice(spec, i)
	}

	gpu.setServerMockFuncs()
	return gpu
}

// NewMockDevice creates a new device with the specified configuration
func NewMockDevice(spec GPUSpec, index int) *MockDevice {
	// Use custom PCI Bus ID format if provided, otherwise fall back to default
	pciBusIDFormat := spec.PciBusIDBase
	if pciBusIDFormat == "" {
		pciBusIDFormat = "0000:%02x:00.0"
	}

	device := &MockDevice{
		Spec:               spec,
		UUID:               "GPU-" + uuid.New().String(),
		PciBusID:           fmt.Sprintf(pciBusIDFormat, index),
		Minor:              index,
		Index:              index,
		GpuInstances:       make(map[*MockGpuInstance]struct{}),
		GpuInstanceCounter: 0,
		MemoryInfo: nvml.Memory{
			Total: spec.TotalMemoryMB * 1024 * 1024,
			Free:  0,
			Used:  0,
		},
	}
	device.setDeviceMockFuncs()
	return device
}

// NewMockGpuInstance creates a new GPU instance
func NewMockGpuInstance(info nvml.GpuInstanceInfo, migProfiles MIGProfileConfig) *MockGpuInstance {
	gi := &MockGpuInstance{
		Info:                   info,
		ComputeInstances:       make(map[*MockComputeInstance]struct{}),
		ComputeInstanceCounter: 0,
		MIGProfiles:            migProfiles,
	}
	gi.setGpuInstanceMockFuncs()
	return gi
}

// NewMockComputeInstance creates a new compute instance
func NewMockComputeInstance(info nvml.ComputeInstanceInfo) *MockComputeInstance {
	ci := &MockComputeInstance{
		Info: info,
	}
	ci.setComputeInstanceMockFuncs()
	return ci
}
