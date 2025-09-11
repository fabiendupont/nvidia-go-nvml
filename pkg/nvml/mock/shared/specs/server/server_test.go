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

package server

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/shared"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/shared/specs/gpu"
)

// createTestServer creates a test server with the specified device count
func createTestServer(deviceCount int) *shared.MockGPU {
	spec := gpu.A100Spec
	spec.DeviceCount = deviceCount
	return shared.NewMockGPU(spec)
}

func TestNewDGXA100(t *testing.T) {
	server := NewDGXA100()
	require.NotNil(t, server)

	// Test that it implements the expected interfaces
	require.Implements(t, (*nvml.Interface)(nil), server)
	require.Implements(t, (*nvml.ExtendedInterface)(nil), server)

	// Test basic server properties
	count, ret := server.DeviceGetCount()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, shared.DGX_DEVICE_COUNT, count) // DGX A100 has 8 GPUs

	// Test device properties
	device, ret := server.DeviceGetHandleByIndex(0)
	require.Equal(t, nvml.SUCCESS, ret)
	require.NotNil(t, device)

	name, ret := device.GetName()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, "NVIDIA A100-SXM4-40GB", name)
}

// TestMockGPUServerCreation verifies server creation and basic properties
func TestMockGPUServerCreation(t *testing.T) {
	server := createTestServer(shared.DGX_DEVICE_COUNT)
	require.NotNil(t, server)

	// Test interface compliance
	require.Implements(t, (*nvml.Interface)(nil), server)
	require.Implements(t, (*nvml.ExtendedInterface)(nil), server)

	// Test device count
	count, ret := server.DeviceGetCount()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, shared.DGX_DEVICE_COUNT, count)

	// Test system information
	driver, ret := server.SystemGetDriverVersion()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, shared.DEFAULT_DRIVER_VERSION, driver)

	nvmlVer, ret := server.SystemGetNVMLVersion()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, shared.DEFAULT_NVML_VERSION, nvmlVer)

	cudaVer, ret := server.SystemGetCudaDriverVersion()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, shared.DEFAULT_CUDA_DRIVER_VERSION, cudaVer)
}

// TestMockGPUDeviceHandling verifies device access and indexing
func TestMockGPUDeviceHandling(t *testing.T) {
	server := createTestServer(shared.DGX_DEVICE_COUNT)

	// Test valid device indices
	for i := 0; i < shared.DGX_DEVICE_COUNT; i++ {
		device, ret := server.DeviceGetHandleByIndex(i)
		require.Equal(t, nvml.SUCCESS, ret)
		require.NotNil(t, device)

		// Test device index
		index, ret := device.GetIndex()
		require.Equal(t, nvml.SUCCESS, ret)
		require.Equal(t, i, index)

		// Test minor number
		minor, ret := device.GetMinorNumber()
		require.Equal(t, nvml.SUCCESS, ret)
		require.Equal(t, i, minor)
	}

	// Test invalid device indices
	_, ret := server.DeviceGetHandleByIndex(-1)
	require.Equal(t, nvml.ERROR_INVALID_ARGUMENT, ret)

	_, ret = server.DeviceGetHandleByIndex(shared.DGX_DEVICE_COUNT)
	require.Equal(t, nvml.ERROR_INVALID_ARGUMENT, ret)
}

// TestMockGPUDeviceProperties verifies all device properties
func TestMockGPUDeviceProperties(t *testing.T) {
	server := createTestServer(shared.DGX_DEVICE_COUNT)
	device, ret := server.DeviceGetHandleByIndex(0)
	require.Equal(t, nvml.SUCCESS, ret)
	require.NotNil(t, device)

	// Test device name
	name, ret := device.GetName()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, gpu.A100Spec.Name, name)

	// Test architecture
	arch, ret := device.GetArchitecture()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, nvml.DeviceArchitecture(gpu.A100Spec.Architecture), arch)

	// Test brand
	brand, ret := device.GetBrand()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, gpu.A100Spec.Brand, brand)

	// Test CUDA compute capability
	major, minor, ret := device.GetCudaComputeCapability()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, gpu.A100Spec.CudaCapabilityMajor, major)
	require.Equal(t, gpu.A100Spec.CudaCapabilityMinor, minor)

	// Test memory info
	memory, ret := device.GetMemoryInfo()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, gpu.A100Spec.TotalMemoryMB*shared.BYTES_PER_MB, memory.Total)

	// Test PCI device ID
	pciInfo, ret := device.GetPciInfo()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, gpu.A100Spec.PciDeviceId, pciInfo.PciDeviceId)

	// Test UUID is set
	uuid, ret := device.GetUUID()
	require.Equal(t, nvml.SUCCESS, ret)
	require.NotEmpty(t, uuid)
	require.Contains(t, uuid, "GPU-")
}

// TestMockGPUDeviceAccessByUUID verifies UUID-based device access
func TestMockGPUDeviceAccessByUUID(t *testing.T) {
	server := createTestServer(shared.DGX_DEVICE_COUNT)

	// Get device by index and its UUID
	originalDevice, ret := server.DeviceGetHandleByIndex(0)
	require.Equal(t, nvml.SUCCESS, ret)

	uuid, ret := originalDevice.GetUUID()
	require.Equal(t, nvml.SUCCESS, ret)
	require.NotEmpty(t, uuid)

	// Get device by UUID
	deviceByUUID, ret := server.DeviceGetHandleByUUID(uuid)
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, originalDevice, deviceByUUID)

	// Test invalid UUID
	_, ret = server.DeviceGetHandleByUUID("invalid-uuid")
	require.Equal(t, nvml.ERROR_INVALID_ARGUMENT, ret)
}

// TestMockGPUDeviceAccessByPciBusId verifies PCI bus ID-based device access
func TestMockGPUDeviceAccessByPciBusId(t *testing.T) {
	server := createTestServer(shared.DGX_DEVICE_COUNT)

	// Test each device's PCI bus ID
	for i := 0; i < shared.DGX_DEVICE_COUNT; i++ {
		originalDevice, ret := server.DeviceGetHandleByIndex(i)
		require.Equal(t, nvml.SUCCESS, ret)

		expectedPciBusID := fmt.Sprintf("0000:%02x:00.0", i)

		// Get device by PCI bus ID
		deviceByPci, ret := server.DeviceGetHandleByPciBusId(expectedPciBusID)
		require.Equal(t, nvml.SUCCESS, ret)
		require.Equal(t, originalDevice, deviceByPci)
	}

	// Test invalid PCI bus ID
	_, ret := server.DeviceGetHandleByPciBusId("invalid-pci-id")
	require.Equal(t, nvml.ERROR_INVALID_ARGUMENT, ret)
}

// TestMockGPUMIGModeOperations verifies MIG mode handling
func TestMockGPUMIGModeOperations(t *testing.T) {
	server := createTestServer(shared.DGX_DEVICE_COUNT)
	device, ret := server.DeviceGetHandleByIndex(0)
	require.Equal(t, nvml.SUCCESS, ret)

	// Initially MIG should be disabled
	current, pending, ret := device.GetMigMode()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, 0, current)
	require.Equal(t, 0, pending)

	// Enable MIG mode
	currentRet, pendingRet := device.SetMigMode(1)
	require.Equal(t, nvml.SUCCESS, currentRet)
	require.Equal(t, nvml.SUCCESS, pendingRet)

	// Verify MIG is enabled
	current, pending, ret = device.GetMigMode()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, 1, current)
	require.Equal(t, 1, pending)

	// Disable MIG mode
	currentRet, pendingRet = device.SetMigMode(0)
	require.Equal(t, nvml.SUCCESS, currentRet)
	require.Equal(t, nvml.SUCCESS, pendingRet)

	// Verify MIG is disabled
	current, pending, ret = device.GetMigMode()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, 0, current)
	require.Equal(t, 0, pending)
}

// TestMockGPUMIGProfilesExist verifies MIG profile configuration exists
func TestMockGPUMIGProfilesExist(t *testing.T) {
	migProfiles := gpu.A100Spec.MIGProfiles
	require.NotNil(t, migProfiles.GpuInstanceProfiles)
	require.NotNil(t, migProfiles.ComputeInstanceProfiles)
	require.NotNil(t, migProfiles.GpuInstancePlacements)
	require.NotNil(t, migProfiles.ComputeInstancePlacements)

	// Test expected profile types exist
	expectedProfiles := []int{
		nvml.GPU_INSTANCE_PROFILE_1_SLICE,
		nvml.GPU_INSTANCE_PROFILE_1_SLICE_REV1,
		nvml.GPU_INSTANCE_PROFILE_1_SLICE_REV2,
		nvml.GPU_INSTANCE_PROFILE_2_SLICE,
		nvml.GPU_INSTANCE_PROFILE_3_SLICE,
		nvml.GPU_INSTANCE_PROFILE_4_SLICE,
		nvml.GPU_INSTANCE_PROFILE_7_SLICE,
	}

	for _, profileId := range expectedProfiles {
		profile, exists := migProfiles.GpuInstanceProfiles[profileId]
		require.True(t, exists, "Profile %d should exist", profileId)
		require.Equal(t, uint32(profileId), profile.Id)
		require.Greater(t, profile.MemorySizeMB, uint64(0))
	}
}

// TestMockGPUGpuInstanceProfileInfo verifies GPU instance profile access
func TestMockGPUGpuInstanceProfileInfo(t *testing.T) {
	server := createTestServer(shared.DGX_DEVICE_COUNT)
	device, ret := server.DeviceGetHandleByIndex(0)
	require.Equal(t, nvml.SUCCESS, ret)

	// Test valid profile access
	profileInfo, ret := device.GetGpuInstanceProfileInfo(nvml.GPU_INSTANCE_PROFILE_1_SLICE)
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, uint32(nvml.GPU_INSTANCE_PROFILE_1_SLICE), profileInfo.Id)
	require.Equal(t, uint32(1), profileInfo.SliceCount)
	require.Equal(t, uint64(shared.A100_MIG_1G_5GB_MEMORY_MB), profileInfo.MemorySizeMB)

	// Test 7-slice profile
	profileInfo7, ret := device.GetGpuInstanceProfileInfo(nvml.GPU_INSTANCE_PROFILE_7_SLICE)
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, uint32(nvml.GPU_INSTANCE_PROFILE_7_SLICE), profileInfo7.Id)
	require.Equal(t, uint32(7), profileInfo7.SliceCount)
	require.Equal(t, uint64(shared.A100_MIG_7G_40GB_MEMORY_MB), profileInfo7.MemorySizeMB)

	// Test invalid profile
	_, ret = device.GetGpuInstanceProfileInfo(-1)
	require.Equal(t, nvml.ERROR_INVALID_ARGUMENT, ret)

	// Test out-of-range profile
	_, ret = device.GetGpuInstanceProfileInfo(nvml.GPU_INSTANCE_PROFILE_COUNT) // Out of range
	require.Equal(t, nvml.ERROR_INVALID_ARGUMENT, ret)
}

// TestMockGPUGpuInstancePlacements verifies GPU instance placement access
func TestMockGPUGpuInstancePlacements(t *testing.T) {
	server := createTestServer(shared.DGX_DEVICE_COUNT)
	device, ret := server.DeviceGetHandleByIndex(0)
	require.Equal(t, nvml.SUCCESS, ret)

	// Test 1-slice placements
	profileInfo, ret := device.GetGpuInstanceProfileInfo(nvml.GPU_INSTANCE_PROFILE_1_SLICE)
	require.Equal(t, nvml.SUCCESS, ret)

	placements, ret := device.GetGpuInstancePossiblePlacements(&profileInfo)
	require.Equal(t, nvml.SUCCESS, ret)
	require.Len(t, placements, 7) // Should have 7 possible placements for 1-slice

	// Test 7-slice placements
	profileInfo7, ret := device.GetGpuInstanceProfileInfo(nvml.GPU_INSTANCE_PROFILE_7_SLICE)
	require.Equal(t, nvml.SUCCESS, ret)

	placements7, ret := device.GetGpuInstancePossiblePlacements(&profileInfo7)
	require.Equal(t, nvml.SUCCESS, ret)
	require.Len(t, placements7, 1) // Should have 1 placement for 7-slice (full GPU)
	require.Equal(t, uint32(0), placements7[0].Start)
	require.Equal(t, uint32(8), placements7[0].Size)
}

// TestMockGPUGpuInstanceLifecycle verifies complete GPU instance lifecycle
func TestMockGPUGpuInstanceLifecycle(t *testing.T) {
	server := createTestServer(shared.DGX_DEVICE_COUNT)
	device, ret := server.DeviceGetHandleByIndex(0)
	require.Equal(t, nvml.SUCCESS, ret)

	// Get 1-slice profile
	profileInfo, ret := device.GetGpuInstanceProfileInfo(nvml.GPU_INSTANCE_PROFILE_1_SLICE)
	require.Equal(t, nvml.SUCCESS, ret)

	// Create GPU instance
	gi, ret := device.CreateGpuInstance(&profileInfo)
	require.Equal(t, nvml.SUCCESS, ret)
	require.NotNil(t, gi)

	// Test GPU instance info
	giInfo, ret := gi.GetInfo()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, device, giInfo.Device)
	require.Equal(t, profileInfo.Id, giInfo.ProfileId)
	require.Equal(t, uint32(0), giInfo.Id) // First instance should have ID 0

	// Get GPU instances for this profile
	instances, ret := device.GetGpuInstances(&profileInfo)
	require.Equal(t, nvml.SUCCESS, ret)
	require.Len(t, instances, 1)
	require.Equal(t, gi, instances[0])

	// Destroy GPU instance
	ret = gi.Destroy()
	require.Equal(t, nvml.SUCCESS, ret)

	// Verify instance is removed
	instances, ret = device.GetGpuInstances(&profileInfo)
	require.Equal(t, nvml.SUCCESS, ret)
	require.Len(t, instances, 0)
}

// TestMockGPUGpuInstanceWithPlacement verifies GPU instance creation with placement
func TestMockGPUGpuInstanceWithPlacement(t *testing.T) {
	server := createTestServer(shared.DGX_DEVICE_COUNT)
	device, ret := server.DeviceGetHandleByIndex(0)
	require.Equal(t, nvml.SUCCESS, ret)

	// Get profile and placement
	profileInfo, ret := device.GetGpuInstanceProfileInfo(nvml.GPU_INSTANCE_PROFILE_1_SLICE)
	require.Equal(t, nvml.SUCCESS, ret)

	placements, ret := device.GetGpuInstancePossiblePlacements(&profileInfo)
	require.Equal(t, nvml.SUCCESS, ret)
	require.NotEmpty(t, placements)

	// Create GPU instance with specific placement
	gi, ret := device.CreateGpuInstanceWithPlacement(&profileInfo, &placements[0])
	require.Equal(t, nvml.SUCCESS, ret)
	require.NotNil(t, gi)

	// Verify placement in instance info
	giInfo, ret := gi.GetInfo()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, placements[0], giInfo.Placement)

	// Clean up
	ret = gi.Destroy()
	require.Equal(t, nvml.SUCCESS, ret)
}

// TestMockGPUComputeInstanceLifecycle verifies complete compute instance lifecycle
func TestMockGPUComputeInstanceLifecycle(t *testing.T) {
	server := createTestServer(shared.DGX_DEVICE_COUNT)
	device, ret := server.DeviceGetHandleByIndex(0)
	require.Equal(t, nvml.SUCCESS, ret)

	// Create GPU instance first
	giProfileInfo, ret := device.GetGpuInstanceProfileInfo(nvml.GPU_INSTANCE_PROFILE_1_SLICE)
	require.Equal(t, nvml.SUCCESS, ret)

	gi, ret := device.CreateGpuInstance(&giProfileInfo)
	require.Equal(t, nvml.SUCCESS, ret)
	require.NotNil(t, gi)

	// Get compute instance profile
	ciProfileInfo, ret := gi.GetComputeInstanceProfileInfo(
		nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
		nvml.COMPUTE_INSTANCE_ENGINE_PROFILE_SHARED,
	)
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, uint32(nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE), ciProfileInfo.Id)

	// Test invalid engine profile
	_, ret = gi.GetComputeInstanceProfileInfo(
		nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
		999, // Invalid engine profile
	)
	require.Equal(t, nvml.ERROR_NOT_SUPPORTED, ret)

	// Get compute instance placements
	_, ret = gi.GetComputeInstancePossiblePlacements(&ciProfileInfo)
	require.Equal(t, nvml.SUCCESS, ret)

	// Create compute instance
	ci, ret := gi.CreateComputeInstance(&ciProfileInfo)
	require.Equal(t, nvml.SUCCESS, ret)
	require.NotNil(t, ci)

	// Test compute instance info
	ciInfo, ret := ci.GetInfo()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, device, ciInfo.Device)
	require.Equal(t, gi, ciInfo.GpuInstance)
	require.Equal(t, ciProfileInfo.Id, ciInfo.ProfileId)
	require.Equal(t, uint32(0), ciInfo.Id) // First instance should have ID 0

	// Get compute instances for this profile
	instances, ret := gi.GetComputeInstances(&ciProfileInfo)
	require.Equal(t, nvml.SUCCESS, ret)
	require.Len(t, instances, 1)
	require.Equal(t, ci, instances[0])

	// Destroy compute instance
	ret = ci.Destroy()
	require.Equal(t, nvml.SUCCESS, ret)

	// Verify compute instance is removed
	instances, ret = gi.GetComputeInstances(&ciProfileInfo)
	require.Equal(t, nvml.SUCCESS, ret)
	require.Len(t, instances, 0)

	// Destroy GPU instance
	ret = gi.Destroy()
	require.Equal(t, nvml.SUCCESS, ret)
}

// TestMockGPUInitShutdownLifecycle verifies init/shutdown behavior
func TestMockGPUInitShutdownLifecycle(t *testing.T) {
	server := createTestServer(shared.DGX_DEVICE_COUNT)

	// Test init
	ret := server.Init()
	require.Equal(t, nvml.SUCCESS, ret)

	// Test lookup symbol
	err := server.LookupSymbol("nvmlInit")
	require.NoError(t, err)

	// Test extensions
	ext := server.Extensions()
	require.NotNil(t, ext)
	require.Equal(t, server, ext)

	// Test shutdown
	ret = server.Shutdown()
	require.Equal(t, nvml.SUCCESS, ret)
}

// TestMockGPUMultipleDevices verifies all devices are unique and correctly indexed
func TestMockGPUMultipleDevices(t *testing.T) {
	server := createTestServer(shared.DGX_DEVICE_COUNT)

	devices := make([]nvml.Device, shared.DGX_DEVICE_COUNT)
	uuids := make(map[string]bool)

	// Get all devices and verify uniqueness
	for i := 0; i < shared.DGX_DEVICE_COUNT; i++ {
		device, ret := server.DeviceGetHandleByIndex(i)
		require.Equal(t, nvml.SUCCESS, ret)
		require.NotNil(t, device)

		devices[i] = device

		// Verify UUID is unique
		uuid, ret := device.GetUUID()
		require.Equal(t, nvml.SUCCESS, ret)
		require.NotEmpty(t, uuid)
		require.False(t, uuids[uuid], "UUID %s should be unique", uuid)
		uuids[uuid] = true

		// Verify device properties are consistent
		index, ret := device.GetIndex()
		require.Equal(t, nvml.SUCCESS, ret)
		require.Equal(t, i, index)

		minor, ret := device.GetMinorNumber()
		require.Equal(t, nvml.SUCCESS, ret)
		require.Equal(t, i, minor)

		name, ret := device.GetName()
		require.Equal(t, nvml.SUCCESS, ret)
		require.Equal(t, gpu.A100Spec.Name, name)
	}

	// Verify all devices are distinct objects
	for i := 0; i < shared.DGX_DEVICE_COUNT; i++ {
		for j := i + 1; j < 8; j++ {
			require.NotEqual(t, devices[i], devices[j], "Devices %d and %d should be different objects", i, j)
		}
	}
}
