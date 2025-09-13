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

// TestDGXB200DeviceCount verifies B200 server device count
func TestDGXB200DeviceCount(t *testing.T) {
	count := GetDGXB200DeviceCount()
	require.Equal(t, shared.DGX_DEVICE_COUNT, count)
}

// TestDGXB200Device verifies B200 device properties
func TestDGXB200Device(t *testing.T) {
	device := GetDGXB200Device(0)
	require.NotNil(t, device)

	mockDevice := device.(*shared.MockDevice)

	// Test B200-specific properties
	require.Equal(t, "NVIDIA B200-SXM6-192GB", mockDevice.Spec.Name)
	require.Equal(t, uint32(shared.B200_PCI_DEVICE_ID), mockDevice.Spec.PciDeviceId)
	require.Equal(t, uint64(shared.B200_TOTAL_MEMORY_BYTES), mockDevice.MemoryInfo.Total)
	require.Equal(t, shared.B200_CUDA_MAJOR, mockDevice.Spec.CudaCapabilityMajor)
	require.Equal(t, shared.B200_CUDA_MINOR, mockDevice.Spec.CudaCapabilityMinor)
	require.Equal(t, nvml.DeviceArchitecture(nvml.DEVICE_ARCH_BLACKWELL), mockDevice.Spec.Architecture)

	// Test MIG profiles are set
	require.NotNil(t, mockDevice.Spec.MIGProfiles.GpuInstanceProfiles)
	require.NotNil(t, mockDevice.Spec.MIGProfiles.ComputeInstanceProfiles)
	require.NotNil(t, mockDevice.Spec.MIGProfiles.GpuInstancePlacements)

	// Verify B200 has expected device index
	require.Equal(t, 0, mockDevice.Index)
}

// TestDGXB200DeviceIndexing verifies B200 device indexing
func TestDGXB200DeviceIndexing(t *testing.T) {
	for i := 0; i < shared.DGX_DEVICE_COUNT; i++ {
		device := GetDGXB200Device(i)
		require.NotNil(t, device)

		mockDevice := device.(*shared.MockDevice)
		require.Equal(t, i, mockDevice.Minor)
		require.Equal(t, i, mockDevice.Index)
		require.NotEmpty(t, mockDevice.UUID)
		require.Equal(t, fmt.Sprintf("0000:90:%02x:00.0", i), mockDevice.PciBusID)
	}
}

// TestDGXB200DeviceInvalidIndex verifies B200 device invalid index handling
func TestDGXB200DeviceInvalidIndex(t *testing.T) {
	require.Panics(t, func() {
		GetDGXB200Device(-1)
	})
	require.Panics(t, func() {
		GetDGXB200Device(shared.DGX_DEVICE_COUNT)
	})
}

// TestDGXB200MigProfiles verifies B200 MIG profile specifications
func TestDGXB200MigProfiles(t *testing.T) {
	profiles := make([]nvml.GpuInstanceProfileInfo, 0, len(gpu.B200MIGProfiles.GpuInstanceProfiles))
	for _, profile := range gpu.B200MIGProfiles.GpuInstanceProfiles {
		profiles = append(profiles, profile)
	}
	require.NotNil(t, profiles)
	require.Len(t, profiles, 7) // B200 supports only 7 documented MIG profiles

	// Test expected B200 profiles exist (actual documented profiles)
	expectedProfilesMap := map[uint32]string{
		nvml.GPU_INSTANCE_PROFILE_1_SLICE:        "1g.23gb",
		nvml.GPU_INSTANCE_PROFILE_1_SLICE_ALL_ME: "1g.23gb+me",
		nvml.GPU_INSTANCE_PROFILE_1_SLICE_REV2:   "1g.45gb",
		nvml.GPU_INSTANCE_PROFILE_2_SLICE:        "2g.45gb",
		nvml.GPU_INSTANCE_PROFILE_3_SLICE:        "3g.90gb",
		nvml.GPU_INSTANCE_PROFILE_4_SLICE:        "4g.90gb",
		nvml.GPU_INSTANCE_PROFILE_7_SLICE:        "7g.180gb",
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
		require.Equal(t, uint32(shared.B200_MIG_P2P_SUPPORT), profile.IsP2pSupported)
	}

	// Test specific B200 memory sizes for key profiles (actual documented values)
	require.Equal(t, uint64(shared.B200_MIG_1G_23GB_MEMORY_MB), profileMap[nvml.GPU_INSTANCE_PROFILE_1_SLICE].MemorySizeMB)
	require.Equal(t, uint64(shared.B200_MIG_2G_45GB_MEMORY_MB), profileMap[nvml.GPU_INSTANCE_PROFILE_2_SLICE].MemorySizeMB)
	require.Equal(t, uint64(shared.B200_MIG_7G_180GB_MEMORY_MB), profileMap[nvml.GPU_INSTANCE_PROFILE_7_SLICE].MemorySizeMB)
}

// TestDGXB200ComputeInstanceProfiles verifies B200 compute instance profiles
func TestDGXB200ComputeInstanceProfiles(t *testing.T) {
	profiles := make([]nvml.ComputeInstanceProfileInfo, 0)
	for _, profileMap := range gpu.B200MIGProfiles.ComputeInstanceProfiles {
		for _, profile := range profileMap {
			profiles = append(profiles, profile)
		}
	}
	require.NotNil(t, profiles)
	require.Len(t, profiles, 15) // B200 supports only 7 documented MIG profiles with 15 compute instances

	// Test that all profiles have valid properties
	for _, profile := range profiles {
		require.Greater(t, profile.SliceCount, uint32(0), "Profile %d should have slice count > 0", profile.Id)
		require.Greater(t, profile.MultiprocessorCount, uint32(0), "Profile %d should have SM count > 0", profile.Id)
		require.GreaterOrEqual(t, profile.SharedCopyEngineCount, uint32(0), "Profile %d should have CE count >= 0", profile.Id)
	}
}

// TestDGXB200MigPlacements verifies B200 MIG placement configurations
func TestDGXB200MigPlacements(t *testing.T) {
	placements := make(map[uint32][]nvml.GpuInstancePlacement)
	for profileId, profilePlacements := range gpu.B200MIGProfiles.GpuInstancePlacements {
		placements[uint32(profileId)] = profilePlacements
	}
	require.NotNil(t, placements)

	// Test placement configurations for actual B200 profiles
	expectedPlacements := map[uint32]int{
		nvml.GPU_INSTANCE_PROFILE_1_SLICE:        7, // 1g.23gb - 7 possible placements
		nvml.GPU_INSTANCE_PROFILE_1_SLICE_ALL_ME: 1, // 1g.23gb+me - 1 possible placement
		nvml.GPU_INSTANCE_PROFILE_1_SLICE_REV2:   4, // 1g.45gb - 4 possible placements
		nvml.GPU_INSTANCE_PROFILE_2_SLICE:        3, // 2g.45gb - 3 possible placements
		nvml.GPU_INSTANCE_PROFILE_3_SLICE:        2, // 3g.90gb - 2 possible placements
		nvml.GPU_INSTANCE_PROFILE_4_SLICE:        1, // 4g.90gb - 1 possible placement
		nvml.GPU_INSTANCE_PROFILE_7_SLICE:        1, // 7g.180gb - 1 possible placement
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

// TestDGXB200DeviceMemoryVerification verifies B200 memory specifications
func TestDGXB200DeviceMemoryVerification(t *testing.T) {
	device := GetDGXB200Device(0)
	mockDevice := device.(*shared.MockDevice)

	// Verify total memory matches expected B200 spec
	require.Equal(t, uint64(shared.B200_TOTAL_MEMORY_BYTES), mockDevice.MemoryInfo.Total)

	// Verify memory constants are consistent
	require.Equal(t, uint64(shared.B200_TOTAL_MEMORY_BYTES/shared.BYTES_PER_MB), uint64(shared.B200_TOTAL_MEMORY_MB))
}

// TestDGXB200DeviceUniqueness verifies all B200 devices are unique
func TestDGXB200DeviceUniqueness(t *testing.T) {
	devices := make([]nvml.Device, shared.DGX_DEVICE_COUNT)
	uuids := make(map[string]bool)
	busIds := make(map[string]bool)

	// Get all devices and verify uniqueness
	for i := 0; i < shared.DGX_DEVICE_COUNT; i++ {
		device := GetDGXB200Device(i)
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

// TestDGXB200P2PSupport verifies B200 P2P support in MIG (enhanced feature)
func TestDGXB200P2PSupport(t *testing.T) {
	device := GetDGXB200Device(0)
	require.NotNil(t, device)

	// Test that B200 profiles support P2P in MIG (unlike A100/H100/H200)
	profileInfo, ret := device.GetGpuInstanceProfileInfo(nvml.GPU_INSTANCE_PROFILE_1_SLICE)
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, uint32(1), profileInfo.IsP2pSupported) // B200 supports P2P in MIG

	// Test with another profile
	profileInfo7, ret := device.GetGpuInstanceProfileInfo(nvml.GPU_INSTANCE_PROFILE_7_SLICE)
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, uint32(1), profileInfo7.IsP2pSupported) // B200 supports P2P in MIG
}

// TestDGXB200EnhancedSMCounts verifies B200 enhanced SM counts
func TestDGXB200EnhancedSMCounts(t *testing.T) {
	device := GetDGXB200Device(0)
	require.NotNil(t, device)

	// Test that B200 has enhanced SM counts compared to H100/H200
	profileInfo1, ret := device.GetGpuInstanceProfileInfo(nvml.GPU_INSTANCE_PROFILE_1_SLICE)
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, uint32(shared.B200_MIG_1_SLICE_SM_COUNT), profileInfo1.MultiprocessorCount) // 26 SMs vs 16 in H100/H200

	profileInfo7, ret := device.GetGpuInstanceProfileInfo(nvml.GPU_INSTANCE_PROFILE_7_SLICE)
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, uint32(shared.B200_MIG_7_SLICE_SM_COUNT), profileInfo7.MultiprocessorCount) // 182 SMs vs 112 in H100/H200
}

// TestDGXB200NewMockServer verifies B200 server creation
func TestDGXB200NewMockServer(t *testing.T) {
	server := NewDGXB200()
	require.NotNil(t, server)

	// Test that it implements the expected interfaces
	require.Implements(t, (*nvml.Interface)(nil), server)
	require.Implements(t, (*nvml.ExtendedInterface)(nil), server)

	// Test B200 specific configuration
	count, ret := server.DeviceGetCount()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, shared.DGX_DEVICE_COUNT, count) // DGX B200 has 8 GPUs

	// Test device properties are B200
	device, ret := server.DeviceGetHandleByIndex(0)
	require.Equal(t, nvml.SUCCESS, ret)
	require.NotNil(t, device)

	name, ret := device.GetName()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, "NVIDIA B200-SXM6-192GB", name)

	// Test Blackwell architecture
	arch, ret := device.GetArchitecture()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, nvml.DeviceArchitecture(nvml.DEVICE_ARCH_BLACKWELL), arch)

	// Test CUDA 10.0 capability
	major, minor, ret := device.GetCudaComputeCapability()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, shared.B200_CUDA_MAJOR, major) // Blackwell = CUDA 10.0
	require.Equal(t, shared.B200_CUDA_MINOR, minor)
}
