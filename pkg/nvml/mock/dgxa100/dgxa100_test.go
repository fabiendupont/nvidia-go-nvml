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

package dgxa100

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/shared"
)

// TestDGXA100Creation verifies DGX A100 server creation
func TestDGXA100Creation(t *testing.T) {
	server := New()
	require.NotNil(t, server)

	// Test that it implements the expected interfaces
	require.Implements(t, (*nvml.Interface)(nil), server)
	require.Implements(t, (*nvml.ExtendedInterface)(nil), server)

	// Test DGX A100 specific configuration
	count, ret := server.DeviceGetCount()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, shared.DGX_DEVICE_COUNT, count) // DGX A100 has 8 GPUs

	// Test device properties are A100
	device, ret := server.DeviceGetHandleByIndex(0)
	require.Equal(t, nvml.SUCCESS, ret)
	require.NotNil(t, device)

	name, ret := device.GetName()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, "NVIDIA A100-SXM4-40GB", name)
}

// TestLegacyMIGVariables verifies backward compatibility with legacy MIG variables
func TestLegacyMIGVariables(t *testing.T) {
	// Test that MIGProfiles variable is accessible for backward compatibility
	require.NotNil(t, MIGProfiles)
	require.NotNil(t, MIGProfiles.GpuInstanceProfiles)
	require.NotNil(t, MIGProfiles.ComputeInstanceProfiles)

	// Test that MIGPlacements variable is accessible for backward compatibility
	require.NotNil(t, MIGPlacements)
	require.NotNil(t, MIGPlacements.GpuInstancePossiblePlacements)
	require.NotNil(t, MIGPlacements.ComputeInstancePossiblePlacements)

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
		profile, exists := MIGProfiles.GpuInstanceProfiles[profileId]
		require.True(t, exists, "Profile %d should exist", profileId)
		require.Equal(t, uint32(profileId), profile.Id)
		require.Greater(t, profile.MemorySizeMB, uint64(0))
	}
}

// TestLegacyConstructors verifies backward compatibility with legacy constructor functions
func TestLegacyConstructors(t *testing.T) {
	// Test legacy device constructor
	device := NewDevice(0)
	require.NotNil(t, device)

	name, ret := device.GetName()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, "NVIDIA A100-SXM4-40GB", name)

	// Test legacy GPU instance constructor
	giInfo := nvml.GpuInstanceInfo{
		Device:    device,
		Id:        0,
		ProfileId: nvml.GPU_INSTANCE_PROFILE_1_SLICE,
		Placement: nvml.GpuInstancePlacement{Start: 0, Size: 1},
	}
	gi := NewGpuInstance(giInfo)
	require.NotNil(t, gi)

	// Test legacy compute instance constructor
	ciInfo := nvml.ComputeInstanceInfo{
		Device:      device,
		GpuInstance: gi,
		Id:          0,
		ProfileId:   nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
		Placement:   nvml.ComputeInstancePlacement{Start: 0, Size: 1},
	}
	ci := NewComputeInstance(ciInfo)
	require.NotNil(t, ci)
}

// TestA100SpecificCharacteristics tests A100-specific values
func TestA100SpecificCharacteristics(t *testing.T) {
	server := New()
	device, ret := server.DeviceGetHandleByIndex(0)
	require.Equal(t, nvml.SUCCESS, ret)

	// Test A100 doesn't support P2P in MIG (IsP2pSupported should be 0)
	profileInfo, ret := device.GetGpuInstanceProfileInfo(nvml.GPU_INSTANCE_PROFILE_1_SLICE)
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, uint32(shared.A100_MIG_P2P_SUPPORT), profileInfo.IsP2pSupported)

	// Test A100 memory values are correct
	profile1 := MIGProfiles.GpuInstanceProfiles[nvml.GPU_INSTANCE_PROFILE_1_SLICE]
	require.Equal(t, uint64(shared.A100_MIG_1G_5GB_MEMORY_MB), profile1.MemorySizeMB) // 1g.5gb

	profile7 := MIGProfiles.GpuInstanceProfiles[nvml.GPU_INSTANCE_PROFILE_7_SLICE]
	require.Equal(t, uint64(shared.A100_MIG_7G_40GB_MEMORY_MB), profile7.MemorySizeMB) // 7g.40gb

	// Test A100 architecture
	arch, ret := device.GetArchitecture()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, nvml.DeviceArchitecture(nvml.DEVICE_ARCH_AMPERE), arch)

	// Test A100 CUDA compute capability
	major, minor, ret := device.GetCudaComputeCapability()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, 8, major) // Ampere
	require.Equal(t, 0, minor)

	// Test A100 PCI device ID
	pciInfo, ret := device.GetPciInfo()
	require.Equal(t, nvml.SUCCESS, ret)
	require.Equal(t, uint32(shared.A100_PCI_DEVICE_ID), pciInfo.PciDeviceId) // A100-SXM4-40GB
}
