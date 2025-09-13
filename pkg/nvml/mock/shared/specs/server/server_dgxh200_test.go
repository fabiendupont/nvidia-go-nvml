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

// TestDGXH200DeviceCount verifies H200 server device count
func TestDGXH200DeviceCount(t *testing.T) {
	count := GetDGXH200DeviceCount()
	require.Equal(t, shared.DGX_DEVICE_COUNT, count)
}

// TestDGXH200Device verifies H200 device properties
func TestDGXH200Device(t *testing.T) {
	device := GetDGXH200Device(0)
	require.NotNil(t, device)

	mockDevice := device.(*shared.MockDevice)

	// Test H200-specific properties
	require.Equal(t, "NVIDIA H200 141GB HBM3e", mockDevice.Spec.Name)
	require.Equal(t, uint32(shared.H200_PCI_DEVICE_ID), mockDevice.Spec.PciDeviceId)
	require.Equal(t, uint64(shared.H200_TOTAL_MEMORY_BYTES), mockDevice.MemoryInfo.Total)
	require.Equal(t, shared.H200_CUDA_MAJOR, mockDevice.Spec.CudaCapabilityMajor)
	require.Equal(t, shared.H200_CUDA_MINOR, mockDevice.Spec.CudaCapabilityMinor)
	require.Equal(t, nvml.DeviceArchitecture(nvml.DEVICE_ARCH_HOPPER), mockDevice.Spec.Architecture)

	// Test MIG profiles are set
	require.NotNil(t, mockDevice.Spec.MIGProfiles.GpuInstanceProfiles)
	require.NotNil(t, mockDevice.Spec.MIGProfiles.ComputeInstanceProfiles)
	require.NotNil(t, mockDevice.Spec.MIGProfiles.GpuInstancePlacements)

	// Verify H200 has expected device index
	require.Equal(t, 0, mockDevice.Index)
}

// TestDGXH200DeviceIndexing verifies H200 device indexing
func TestDGXH200DeviceIndexing(t *testing.T) {
	for i := 0; i < shared.DGX_DEVICE_COUNT; i++ {
		device := GetDGXH200Device(i)
		require.NotNil(t, device)

		mockDevice := device.(*shared.MockDevice)
		require.Equal(t, i, mockDevice.Minor)
		require.Equal(t, i, mockDevice.Index)
		require.NotEmpty(t, mockDevice.UUID)
		require.Equal(t, fmt.Sprintf("0000:4b:%02x:00.0", i), mockDevice.PciBusID)
	}
}

// TestDGXH200DeviceInvalidIndex verifies H200 device invalid index handling
func TestDGXH200DeviceInvalidIndex(t *testing.T) {
	require.Panics(t, func() {
		GetDGXH200Device(-1)
	})
	require.Panics(t, func() {
		GetDGXH200Device(shared.DGX_DEVICE_COUNT)
	})
}

// TestDGXH200MigProfiles verifies H200 MIG profile specifications
func TestDGXH200MigProfiles(t *testing.T) {
	profiles := make([]nvml.GpuInstanceProfileInfo, 0, len(gpu.H200MIGProfiles.GpuInstanceProfiles))
	for _, profile := range gpu.H200MIGProfiles.GpuInstanceProfiles {
		profiles = append(profiles, profile)
	}
	require.NotNil(t, profiles)
	require.Len(t, profiles, 17) // Updated to match comprehensive implementation

	// Test expected H200 profiles exist (using valid NVML constants)
	expectedProfilesMap := map[uint32]string{
		nvml.GPU_INSTANCE_PROFILE_1_SLICE: "1g.18gb",
		nvml.GPU_INSTANCE_PROFILE_2_SLICE: "2g.35gb",
		nvml.GPU_INSTANCE_PROFILE_3_SLICE: "3g.71gb",
		nvml.GPU_INSTANCE_PROFILE_4_SLICE: "4g.71gb",
		nvml.GPU_INSTANCE_PROFILE_7_SLICE: "7g.141gb",
		nvml.GPU_INSTANCE_PROFILE_8_SLICE: "8g.141gb",
	}

	profileMap := make(map[uint32]nvml.GpuInstanceProfileInfo)
	for _, profile := range profiles {
		profileMap[profile.Id] = profile
	}

	for profileId, expectedName := range expectedProfilesMap {
		profile, exists := profileMap[profileId]
		require.True(t, exists, "Profile %d (%s) should exist", profileId, expectedName)
		require.Equal(t, profileId, profile.Id)
		require.Greater(t, profile.MemorySizeMB, uint64(0))
		require.Equal(t, uint32(shared.H200_MIG_P2P_SUPPORT), profile.IsP2pSupported)
	}

	// Test specific H200 memory sizes for key profiles
	require.Equal(t, uint64(shared.H200_MIG_1G_18GB_MEMORY_MB), profileMap[nvml.GPU_INSTANCE_PROFILE_1_SLICE].MemorySizeMB)
	require.Equal(t, uint64(shared.H200_MIG_2G_35GB_MEMORY_MB), profileMap[nvml.GPU_INSTANCE_PROFILE_2_SLICE].MemorySizeMB)
	require.Equal(t, uint64(shared.H200_MIG_7G_141GB_MEMORY_MB), profileMap[nvml.GPU_INSTANCE_PROFILE_7_SLICE].MemorySizeMB)
}

// TestDGXH200ComputeInstanceProfiles verifies H200 compute instance profiles
func TestDGXH200ComputeInstanceProfiles(t *testing.T) {
	profiles := make([]nvml.ComputeInstanceProfileInfo, 0)
	for _, profileMap := range gpu.H200MIGProfiles.ComputeInstanceProfiles {
		for _, profile := range profileMap {
			profiles = append(profiles, profile)
		}
	}
	require.NotNil(t, profiles)
	require.Len(t, profiles, 21) // Updated to match comprehensive implementation

	// Test that all profiles have valid properties
	for _, profile := range profiles {
		require.Greater(t, profile.SliceCount, uint32(0), "Profile %d should have slice count > 0", profile.Id)
		require.Greater(t, profile.MultiprocessorCount, uint32(0), "Profile %d should have SM count > 0", profile.Id)
		require.GreaterOrEqual(t, profile.SharedCopyEngineCount, uint32(0), "Profile %d should have CE count >= 0", profile.Id)
	}
}

// TestDGXH200MigPlacements verifies H200 MIG placement configurations
func TestDGXH200MigPlacements(t *testing.T) {
	placements := make(map[uint32][]nvml.GpuInstancePlacement)
	for profileId, profilePlacements := range gpu.H200MIGProfiles.GpuInstancePlacements {
		placements[uint32(profileId)] = profilePlacements
	}
	require.NotNil(t, placements)

	// Test placement configurations for key profiles (using valid NVML constants)
	expectedPlacements := map[uint32]int{
		nvml.GPU_INSTANCE_PROFILE_1_SLICE: 7, // 1g.18gb - 7 possible placements
		nvml.GPU_INSTANCE_PROFILE_2_SLICE: 3, // 2g.35gb - 3 possible placements
		nvml.GPU_INSTANCE_PROFILE_3_SLICE: 2, // 3g.71gb - 2 possible placements
		nvml.GPU_INSTANCE_PROFILE_4_SLICE: 1, // 4g.71gb - 1 possible placement
		nvml.GPU_INSTANCE_PROFILE_7_SLICE: 1, // 7g.141gb - 1 possible placement
		nvml.GPU_INSTANCE_PROFILE_8_SLICE: 1, // 8g.141gb - 1 possible placement
	}

	for profileId, expectedCount := range expectedPlacements {
		profilePlacements, exists := placements[profileId]
		require.True(t, exists, "Placements for profile %d should exist", profileId)
		require.Len(t, profilePlacements, expectedCount, "Profile %d should have %d placements", profileId, expectedCount)

		// Verify placement sizes are correct
		for _, placement := range profilePlacements {
			switch profileId {
			case nvml.GPU_INSTANCE_PROFILE_1_SLICE: // 1-slice profiles
				require.Equal(t, uint32(shared.MIG_1_SLICE_PLACEMENT_SIZE), placement.Size)
			case nvml.GPU_INSTANCE_PROFILE_2_SLICE: // 2-slice profile
				require.Equal(t, uint32(shared.MIG_2_SLICE_PLACEMENT_SIZE), placement.Size)
			case nvml.GPU_INSTANCE_PROFILE_3_SLICE: // 3-slice profile
				require.Equal(t, uint32(shared.MIG_3_SLICE_PLACEMENT_SIZE), placement.Size)
			case nvml.GPU_INSTANCE_PROFILE_4_SLICE: // 4-slice profile
				require.Equal(t, uint32(shared.MIG_4_SLICE_PLACEMENT_SIZE), placement.Size)
			case nvml.GPU_INSTANCE_PROFILE_7_SLICE: // 7-slice profile
				require.Equal(t, uint32(shared.MIG_7_SLICE_PLACEMENT_SIZE), placement.Size)
			case nvml.GPU_INSTANCE_PROFILE_8_SLICE: // 8-slice profile
				require.Equal(t, uint32(8), placement.Size)
			}
		}
	}
}

// TestDGXH200DeviceMemoryVerification verifies H200 memory specifications
func TestDGXH200DeviceMemoryVerification(t *testing.T) {
	device := GetDGXH200Device(0)
	mockDevice := device.(*shared.MockDevice)

	// Verify total memory matches expected H200 spec
	require.Equal(t, uint64(shared.H200_TOTAL_MEMORY_BYTES), mockDevice.MemoryInfo.Total)

	// Verify memory constants are consistent
	require.Equal(t, uint64(shared.H200_TOTAL_MEMORY_BYTES/shared.BYTES_PER_MB), uint64(shared.H200_TOTAL_MEMORY_MB))
}

// TestDGXH200DeviceUniqueness verifies all H200 devices are unique
func TestDGXH200DeviceUniqueness(t *testing.T) {
	devices := make([]nvml.Device, shared.DGX_DEVICE_COUNT)
	uuids := make(map[string]bool)
	busIds := make(map[string]bool)

	// Get all devices and verify uniqueness
	for i := 0; i < shared.DGX_DEVICE_COUNT; i++ {
		device := GetDGXH200Device(i)
		require.NotNil(t, device)
		devices[i] = device

		mockDevice := device.(*shared.MockDevice)

		// Verify UUID is unique
		require.False(t, uuids[mockDevice.UUID], "UUID %s should be unique", mockDevice.UUID)
		uuids[mockDevice.UUID] = true

		// Verify PCI Bus ID is unique
		require.False(t, busIds[mockDevice.PciBusID], "PCI Bus ID %s should be unique", mockDevice.PciBusID)
		busIds[mockDevice.PciBusID] = true

		// Verify device has correct index properties
		require.Equal(t, i, mockDevice.Minor)
		require.Equal(t, i, mockDevice.Index)
	}

	// Verify all devices are distinct objects
	for i := 0; i < shared.DGX_DEVICE_COUNT; i++ {
		for j := i + 1; j < shared.DGX_DEVICE_COUNT; j++ {
			require.NotEqual(t, devices[i], devices[j], "Devices %d and %d should be different objects", i, j)
		}
	}
}

// TestDGXH200VsH100Differences verifies H200 differs from H100 appropriately
func TestDGXH200VsH100Differences(t *testing.T) {
	h100Device := GetDGXH100Device(0).(*shared.MockDevice)
	h200Device := GetDGXH200Device(0).(*shared.MockDevice)

	// Same architecture but different names
	require.Equal(t, h100Device.Spec.Architecture, h200Device.Spec.Architecture)
	require.NotEqual(t, h100Device.Spec.Name, h200Device.Spec.Name)

	// Different PCI device IDs
	require.NotEqual(t, h100Device.Spec.PciDeviceId, h200Device.Spec.PciDeviceId)

	// Different memory sizes (H200 has more memory)
	require.Greater(t, h200Device.MemoryInfo.Total, h100Device.MemoryInfo.Total)

	// Different UUIDs and Bus IDs
	require.NotEqual(t, h100Device.UUID, h200Device.UUID)
	require.NotEqual(t, h100Device.PciBusID, h200Device.PciBusID)

	// Same CUDA compute capability (both Hopper generation)
	require.Equal(t, h100Device.Spec.CudaCapabilityMajor, h200Device.Spec.CudaCapabilityMajor)
	require.Equal(t, h100Device.Spec.CudaCapabilityMinor, h200Device.Spec.CudaCapabilityMinor)
}
