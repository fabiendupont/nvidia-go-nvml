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
	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

// setServerMockFuncs sets up the server-level mock functions
func (s *MockGPU) setServerMockFuncs() {
	s.ExtensionsFunc = func() nvml.ExtendedInterface {
		return s
	}

	s.LookupSymbolFunc = func(symbol string) error {
		return nil
	}

	s.InitFunc = func() nvml.Return {
		return nvml.SUCCESS
	}

	s.ShutdownFunc = func() nvml.Return {
		return nvml.SUCCESS
	}

	s.SystemGetDriverVersionFunc = func() (string, nvml.Return) {
		return s.Spec.DriverVersion, nvml.SUCCESS
	}

	s.SystemGetNVMLVersionFunc = func() (string, nvml.Return) {
		return s.Spec.NvmlVersion, nvml.SUCCESS
	}

	s.SystemGetCudaDriverVersionFunc = func() (int, nvml.Return) {
		return s.Spec.CudaDriverVersion, nvml.SUCCESS
	}

	s.DeviceGetCountFunc = func() (int, nvml.Return) {
		return len(s.Devices), nvml.SUCCESS
	}

	s.DeviceGetHandleByIndexFunc = func(index int) (nvml.Device, nvml.Return) {
		if index < 0 || index >= len(s.Devices) {
			return nil, nvml.ERROR_INVALID_ARGUMENT
		}
		return s.Devices[index], nvml.SUCCESS
	}

	s.DeviceGetHandleByUUIDFunc = func(uuid string) (nvml.Device, nvml.Return) {
		for _, d := range s.Devices {
			if uuid == d.(*MockDevice).UUID {
				return d, nvml.SUCCESS
			}
		}
		return nil, nvml.ERROR_INVALID_ARGUMENT
	}

	s.DeviceGetHandleByPciBusIdFunc = func(busID string) (nvml.Device, nvml.Return) {
		for _, d := range s.Devices {
			if busID == d.(*MockDevice).PciBusID {
				return d, nvml.SUCCESS
			}
		}
		return nil, nvml.ERROR_INVALID_ARGUMENT
	}

}

// setDeviceMockFuncs sets up device-level mock functions
func (d *MockDevice) setDeviceMockFuncs() {
	d.GetMinorNumberFunc = func() (int, nvml.Return) {
		return d.Minor, nvml.SUCCESS
	}

	d.GetIndexFunc = func() (int, nvml.Return) {
		return d.Index, nvml.SUCCESS
	}

	d.GetCudaComputeCapabilityFunc = func() (int, int, nvml.Return) {
		return d.Spec.CudaCapabilityMajor, d.Spec.CudaCapabilityMinor, nvml.SUCCESS
	}

	d.GetUUIDFunc = func() (string, nvml.Return) {
		return d.UUID, nvml.SUCCESS
	}

	d.GetNameFunc = func() (string, nvml.Return) {
		return d.Spec.Name, nvml.SUCCESS
	}

	d.GetBrandFunc = func() (nvml.BrandType, nvml.Return) {
		return d.Spec.Brand, nvml.SUCCESS
	}

	d.GetArchitectureFunc = func() (nvml.DeviceArchitecture, nvml.Return) {
		return d.Spec.Architecture, nvml.SUCCESS
	}

	d.GetMemoryInfoFunc = func() (nvml.Memory, nvml.Return) {
		return d.MemoryInfo, nvml.SUCCESS
	}

	d.GetPciInfoFunc = func() (nvml.PciInfo, nvml.Return) {
		return nvml.PciInfo{
			PciDeviceId: d.Spec.PciDeviceId,
		}, nvml.SUCCESS
	}

	d.SetMigModeFunc = func(mode int) (nvml.Return, nvml.Return) {
		d.MigMode = mode
		return nvml.SUCCESS, nvml.SUCCESS
	}

	d.GetMigModeFunc = func() (int, int, nvml.Return) {
		return d.MigMode, d.MigMode, nvml.SUCCESS
	}

	d.GetGpuInstanceProfileInfoFunc = func(giProfileId int) (nvml.GpuInstanceProfileInfo, nvml.Return) {
		if giProfileId < 0 || giProfileId >= nvml.GPU_INSTANCE_PROFILE_COUNT {
			return nvml.GpuInstanceProfileInfo{}, nvml.ERROR_INVALID_ARGUMENT
		}

		if profile, exists := d.Spec.MIGProfiles.GpuInstanceProfiles[giProfileId]; exists {
			return profile, nvml.SUCCESS
		}

		return nvml.GpuInstanceProfileInfo{}, nvml.ERROR_NOT_SUPPORTED
	}

	d.GetGpuInstancePossiblePlacementsFunc = func(info *nvml.GpuInstanceProfileInfo) ([]nvml.GpuInstancePlacement, nvml.Return) {
		return d.Spec.MIGProfiles.GpuInstancePlacements[int(info.Id)], nvml.SUCCESS
	}

	d.CreateGpuInstanceFunc = func(info *nvml.GpuInstanceProfileInfo) (nvml.GpuInstance, nvml.Return) {
		d.Lock()
		defer d.Unlock()
		giInfo := nvml.GpuInstanceInfo{
			Device:    d,
			Id:        d.GpuInstanceCounter,
			ProfileId: info.Id,
		}
		d.GpuInstanceCounter++
		gi := NewMockGpuInstance(giInfo, d.Spec.MIGProfiles)
		d.GpuInstances[gi] = struct{}{}
		return gi, nvml.SUCCESS
	}

	d.CreateGpuInstanceWithPlacementFunc = func(info *nvml.GpuInstanceProfileInfo, placement *nvml.GpuInstancePlacement) (nvml.GpuInstance, nvml.Return) {
		d.Lock()
		defer d.Unlock()
		giInfo := nvml.GpuInstanceInfo{
			Device:    d,
			Id:        d.GpuInstanceCounter,
			ProfileId: info.Id,
			Placement: *placement,
		}
		d.GpuInstanceCounter++
		gi := NewMockGpuInstance(giInfo, d.Spec.MIGProfiles)
		d.GpuInstances[gi] = struct{}{}
		return gi, nvml.SUCCESS
	}

	d.GetGpuInstancesFunc = func(info *nvml.GpuInstanceProfileInfo) ([]nvml.GpuInstance, nvml.Return) {
		d.RLock()
		defer d.RUnlock()
		var gis []nvml.GpuInstance
		for gi := range d.GpuInstances {
			if gi.Info.ProfileId == info.Id {
				gis = append(gis, gi)
			}
		}
		return gis, nvml.SUCCESS
	}
}

// setGpuInstanceMockFuncs sets up GPU instance mock functions
func (gi *MockGpuInstance) setGpuInstanceMockFuncs() {
	gi.GetInfoFunc = func() (nvml.GpuInstanceInfo, nvml.Return) {
		return gi.Info, nvml.SUCCESS
	}

	gi.GetComputeInstanceProfileInfoFunc = func(ciProfileId int, ciEngProfileId int) (nvml.ComputeInstanceProfileInfo, nvml.Return) {
		if ciProfileId < 0 || ciProfileId >= nvml.COMPUTE_INSTANCE_PROFILE_COUNT {
			return nvml.ComputeInstanceProfileInfo{}, nvml.ERROR_INVALID_ARGUMENT
		}

		if ciEngProfileId != nvml.COMPUTE_INSTANCE_ENGINE_PROFILE_SHARED {
			return nvml.ComputeInstanceProfileInfo{}, nvml.ERROR_NOT_SUPPORTED
		}

		giProfileId := int(gi.Info.ProfileId)

		if giProfiles, exists := gi.MIGProfiles.ComputeInstanceProfiles[giProfileId]; exists {
			if profile, exists := giProfiles[ciProfileId]; exists {
				return profile, nvml.SUCCESS
			}
		}

		return nvml.ComputeInstanceProfileInfo{}, nvml.ERROR_NOT_SUPPORTED
	}

	gi.GetComputeInstancePossiblePlacementsFunc = func(info *nvml.ComputeInstanceProfileInfo) ([]nvml.ComputeInstancePlacement, nvml.Return) {
		return gi.MIGProfiles.ComputeInstancePlacements[int(gi.Info.Id)][int(info.Id)], nvml.SUCCESS
	}

	gi.CreateComputeInstanceFunc = func(info *nvml.ComputeInstanceProfileInfo) (nvml.ComputeInstance, nvml.Return) {
		gi.Lock()
		defer gi.Unlock()
		ciInfo := nvml.ComputeInstanceInfo{
			Device:      gi.Info.Device,
			GpuInstance: gi,
			Id:          gi.ComputeInstanceCounter,
			ProfileId:   info.Id,
		}
		gi.ComputeInstanceCounter++
		ci := NewMockComputeInstance(ciInfo)
		gi.ComputeInstances[ci] = struct{}{}
		return ci, nvml.SUCCESS
	}

	gi.GetComputeInstancesFunc = func(info *nvml.ComputeInstanceProfileInfo) ([]nvml.ComputeInstance, nvml.Return) {
		gi.RLock()
		defer gi.RUnlock()
		var cis []nvml.ComputeInstance
		for ci := range gi.ComputeInstances {
			if ci.Info.ProfileId == info.Id {
				cis = append(cis, ci)
			}
		}
		return cis, nvml.SUCCESS
	}

	gi.DestroyFunc = func() nvml.Return {
		d := gi.Info.Device.(*MockDevice)
		d.Lock()
		defer d.Unlock()
		delete(d.GpuInstances, gi)
		return nvml.SUCCESS
	}
}

// setComputeInstanceMockFuncs sets up compute instance mock functions
func (ci *MockComputeInstance) setComputeInstanceMockFuncs() {
	ci.GetInfoFunc = func() (nvml.ComputeInstanceInfo, nvml.Return) {
		return ci.Info, nvml.SUCCESS
	}

	ci.DestroyFunc = func() nvml.Return {
		gi := ci.Info.GpuInstance.(*MockGpuInstance)
		gi.Lock()
		defer gi.Unlock()
		delete(gi.ComputeInstances, ci)
		return nvml.SUCCESS
	}
}
