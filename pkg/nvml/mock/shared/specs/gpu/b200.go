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

package gpu

import (
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/shared"
)

// B200Spec defines the GPU specification for NVIDIA B200-SXM6-192GB
var B200Spec = shared.GPUSpec{
	Name:                "NVIDIA B200-SXM6-192GB",
	Architecture:        nvml.DEVICE_ARCH_BLACKWELL,
	Brand:               nvml.BRAND_TESLA,
	PciDeviceId:         shared.B200_PCI_DEVICE_ID,
	PciBusIDBase:        "0000:90:%02x:00.0", // B200 uses bus 0x90:0x00-0x07
	DriverVersion:       shared.DEFAULT_DRIVER_VERSION,
	NvmlVersion:         shared.DEFAULT_NVML_VERSION,
	CudaDriverVersion:   shared.DEFAULT_CUDA_DRIVER_VERSION,
	CudaCapabilityMajor: shared.B200_CUDA_MAJOR,
	CudaCapabilityMinor: shared.B200_CUDA_MINOR,
	TotalMemoryMB:       shared.B200_TOTAL_MEMORY_MB,
	MIGProfiles:         B200MIGProfiles,
}

// B200MIGProfiles defines the actual documented MIG profile configuration for B200
// Based on official NVIDIA documentation: 7 documented profiles only
var B200MIGProfiles = shared.MIGProfileConfig{
	GpuInstanceProfiles: map[int]nvml.GpuInstanceProfileInfo{
		// MIG 1g.23gb - 1/8 memory, 1/7 SMs, 7 instances
		nvml.GPU_INSTANCE_PROFILE_1_SLICE: {
			Id:                  nvml.GPU_INSTANCE_PROFILE_1_SLICE,
			IsP2pSupported:      shared.B200_MIG_P2P_SUPPORT,
			SliceCount:          1,
			InstanceCount:       shared.B200_MIG_1G_23GB_INSTANCE_COUNT,
			MultiprocessorCount: shared.B200_MIG_1_SLICE_SM_COUNT,
			CopyEngineCount:     shared.B200_MIG_1G_23GB_CE_COUNT,
			DecoderCount:        1,
			EncoderCount:        0,
			JpegCount:           1,
			OfaCount:            0,
			MemorySizeMB:        shared.B200_MIG_1G_23GB_MEMORY_MB,
		},
		// MIG 1g.23gb+me - 1/8 memory, 1/7 SMs, 1 instance (with media extensions)
		nvml.GPU_INSTANCE_PROFILE_1_SLICE_ALL_ME: {
			Id:                  nvml.GPU_INSTANCE_PROFILE_1_SLICE_ALL_ME,
			IsP2pSupported:      shared.B200_MIG_P2P_SUPPORT,
			SliceCount:          1,
			InstanceCount:       shared.B200_MIG_1G_23GB_ME_INSTANCE_COUNT,
			MultiprocessorCount: shared.B200_MIG_1_SLICE_SM_COUNT,
			CopyEngineCount:     shared.B200_MIG_1G_23GB_ME_CE_COUNT,
			DecoderCount:        1,
			EncoderCount:        0,
			JpegCount:           1,
			OfaCount:            1, // OFA included in +me variant
			MemorySizeMB:        shared.B200_MIG_1G_23GB_ME_MEMORY_MB,
		},
		// MIG 1g.45gb - 2/8 memory, 1/7 SMs, 4 instances
		nvml.GPU_INSTANCE_PROFILE_1_SLICE_REV2: {
			Id:                  nvml.GPU_INSTANCE_PROFILE_1_SLICE_REV2,
			IsP2pSupported:      shared.B200_MIG_P2P_SUPPORT,
			SliceCount:          1,
			InstanceCount:       shared.B200_MIG_1G_45GB_INSTANCE_COUNT,
			MultiprocessorCount: shared.B200_MIG_1_SLICE_SM_COUNT,
			CopyEngineCount:     shared.B200_MIG_1G_45GB_CE_COUNT,
			DecoderCount:        1,
			EncoderCount:        0,
			JpegCount:           1,
			OfaCount:            0,
			MemorySizeMB:        shared.B200_MIG_1G_45GB_MEMORY_MB,
		},
		// MIG 2g.45gb - 2/8 memory, 2/7 SMs, 3 instances
		nvml.GPU_INSTANCE_PROFILE_2_SLICE: {
			Id:                  nvml.GPU_INSTANCE_PROFILE_2_SLICE,
			IsP2pSupported:      shared.B200_MIG_P2P_SUPPORT,
			SliceCount:          2,
			InstanceCount:       shared.B200_MIG_2G_45GB_INSTANCE_COUNT,
			MultiprocessorCount: shared.B200_MIG_2_SLICE_SM_COUNT,
			CopyEngineCount:     shared.B200_MIG_2G_45GB_CE_COUNT,
			DecoderCount:        1,
			EncoderCount:        0,
			JpegCount:           1,
			OfaCount:            0,
			MemorySizeMB:        shared.B200_MIG_2G_45GB_MEMORY_MB,
		},
		// MIG 3g.90gb - 4/8 memory, 3/7 SMs, 2 instances
		nvml.GPU_INSTANCE_PROFILE_3_SLICE: {
			Id:                  nvml.GPU_INSTANCE_PROFILE_3_SLICE,
			IsP2pSupported:      shared.B200_MIG_P2P_SUPPORT,
			SliceCount:          3,
			InstanceCount:       shared.B200_MIG_3G_90GB_INSTANCE_COUNT,
			MultiprocessorCount: shared.B200_MIG_3_SLICE_SM_COUNT,
			CopyEngineCount:     shared.B200_MIG_3G_90GB_CE_COUNT,
			DecoderCount:        1,
			EncoderCount:        0,
			JpegCount:           1,
			OfaCount:            0,
			MemorySizeMB:        shared.B200_MIG_3G_90GB_MEMORY_MB,
		},
		// MIG 4g.90gb - 4/8 memory, 4/7 SMs, 1 instance
		nvml.GPU_INSTANCE_PROFILE_4_SLICE: {
			Id:                  nvml.GPU_INSTANCE_PROFILE_4_SLICE,
			IsP2pSupported:      shared.B200_MIG_P2P_SUPPORT,
			SliceCount:          4,
			InstanceCount:       shared.B200_MIG_4G_90GB_INSTANCE_COUNT,
			MultiprocessorCount: shared.B200_MIG_4_SLICE_SM_COUNT,
			CopyEngineCount:     shared.B200_MIG_4G_90GB_CE_COUNT,
			DecoderCount:        1,
			EncoderCount:        0,
			JpegCount:           1,
			OfaCount:            0,
			MemorySizeMB:        shared.B200_MIG_4G_90GB_MEMORY_MB,
		},
		// MIG 7g.180gb - Full memory, Full SMs, 1 instance
		nvml.GPU_INSTANCE_PROFILE_7_SLICE: {
			Id:                  nvml.GPU_INSTANCE_PROFILE_7_SLICE,
			IsP2pSupported:      shared.B200_MIG_P2P_SUPPORT,
			SliceCount:          7,
			InstanceCount:       shared.B200_MIG_7G_180GB_INSTANCE_COUNT,
			MultiprocessorCount: shared.B200_MIG_7_SLICE_SM_COUNT,
			CopyEngineCount:     shared.B200_MIG_7G_180GB_CE_COUNT,
			DecoderCount:        1,
			EncoderCount:        0,
			JpegCount:           1,
			OfaCount:            1, // Full GPU includes OFA
			MemorySizeMB:        shared.B200_MIG_7G_180GB_MEMORY_MB,
		},
	},
	ComputeInstanceProfiles: map[int]map[int]nvml.ComputeInstanceProfileInfo{
		// Each GPU instance profile has basic compute instance configurations
		nvml.GPU_INSTANCE_PROFILE_1_SLICE: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
				SliceCount:            1,
				InstanceCount:         1,
				MultiprocessorCount:   shared.B200_MIG_1_SLICE_SM_COUNT,
				SharedCopyEngineCount: shared.B200_MIG_1G_23GB_CE_COUNT,
				SharedDecoderCount:    1,
				SharedEncoderCount:    0,
				SharedJpegCount:       1,
				SharedOfaCount:        0,
			},
		},
		nvml.GPU_INSTANCE_PROFILE_1_SLICE_ALL_ME: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
				SliceCount:            1,
				InstanceCount:         1,
				MultiprocessorCount:   shared.B200_MIG_1_SLICE_SM_COUNT,
				SharedCopyEngineCount: shared.B200_MIG_1G_23GB_ME_CE_COUNT,
				SharedDecoderCount:    1,
				SharedEncoderCount:    0,
				SharedJpegCount:       1,
				SharedOfaCount:        1,
			},
		},
		nvml.GPU_INSTANCE_PROFILE_1_SLICE_REV2: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
				SliceCount:            1,
				InstanceCount:         1,
				MultiprocessorCount:   shared.B200_MIG_1_SLICE_SM_COUNT,
				SharedCopyEngineCount: shared.B200_MIG_1G_45GB_CE_COUNT,
				SharedDecoderCount:    1,
				SharedEncoderCount:    0,
				SharedJpegCount:       1,
				SharedOfaCount:        0,
			},
		},
		nvml.GPU_INSTANCE_PROFILE_2_SLICE: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
				SliceCount:            1,
				InstanceCount:         2,
				MultiprocessorCount:   shared.B200_MIG_1_SLICE_SM_COUNT,
				SharedCopyEngineCount: shared.B200_MIG_2G_45GB_CE_COUNT,
				SharedDecoderCount:    1,
				SharedEncoderCount:    0,
				SharedJpegCount:       1,
				SharedOfaCount:        0,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE,
				SliceCount:            2,
				InstanceCount:         1,
				MultiprocessorCount:   shared.B200_MIG_2_SLICE_SM_COUNT,
				SharedCopyEngineCount: shared.B200_MIG_2G_45GB_CE_COUNT,
				SharedDecoderCount:    1,
				SharedEncoderCount:    0,
				SharedJpegCount:       1,
				SharedOfaCount:        0,
			},
		},
		nvml.GPU_INSTANCE_PROFILE_3_SLICE: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
				SliceCount:            1,
				InstanceCount:         3,
				MultiprocessorCount:   shared.B200_MIG_1_SLICE_SM_COUNT,
				SharedCopyEngineCount: shared.B200_MIG_3G_90GB_CE_COUNT,
				SharedDecoderCount:    1,
				SharedEncoderCount:    0,
				SharedJpegCount:       1,
				SharedOfaCount:        0,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_3_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_3_SLICE,
				SliceCount:            3,
				InstanceCount:         1,
				MultiprocessorCount:   shared.B200_MIG_3_SLICE_SM_COUNT,
				SharedCopyEngineCount: shared.B200_MIG_3G_90GB_CE_COUNT,
				SharedDecoderCount:    1,
				SharedEncoderCount:    0,
				SharedJpegCount:       1,
				SharedOfaCount:        0,
			},
		},
		nvml.GPU_INSTANCE_PROFILE_4_SLICE: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
				SliceCount:            1,
				InstanceCount:         4,
				MultiprocessorCount:   shared.B200_MIG_1_SLICE_SM_COUNT,
				SharedCopyEngineCount: shared.B200_MIG_4G_90GB_CE_COUNT,
				SharedDecoderCount:    1,
				SharedEncoderCount:    0,
				SharedJpegCount:       1,
				SharedOfaCount:        0,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE,
				SliceCount:            2,
				InstanceCount:         2,
				MultiprocessorCount:   shared.B200_MIG_2_SLICE_SM_COUNT,
				SharedCopyEngineCount: shared.B200_MIG_4G_90GB_CE_COUNT,
				SharedDecoderCount:    1,
				SharedEncoderCount:    0,
				SharedJpegCount:       1,
				SharedOfaCount:        0,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_4_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_4_SLICE,
				SliceCount:            4,
				InstanceCount:         1,
				MultiprocessorCount:   shared.B200_MIG_4_SLICE_SM_COUNT,
				SharedCopyEngineCount: shared.B200_MIG_4G_90GB_CE_COUNT,
				SharedDecoderCount:    1,
				SharedEncoderCount:    0,
				SharedJpegCount:       1,
				SharedOfaCount:        0,
			},
		},
		nvml.GPU_INSTANCE_PROFILE_7_SLICE: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
				SliceCount:            1,
				InstanceCount:         7,
				MultiprocessorCount:   shared.B200_MIG_1_SLICE_SM_COUNT,
				SharedCopyEngineCount: shared.B200_MIG_7G_180GB_CE_COUNT,
				SharedDecoderCount:    1,
				SharedEncoderCount:    0,
				SharedJpegCount:       1,
				SharedOfaCount:        1,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE,
				SliceCount:            2,
				InstanceCount:         3,
				MultiprocessorCount:   shared.B200_MIG_2_SLICE_SM_COUNT,
				SharedCopyEngineCount: shared.B200_MIG_7G_180GB_CE_COUNT,
				SharedDecoderCount:    1,
				SharedEncoderCount:    0,
				SharedJpegCount:       1,
				SharedOfaCount:        1,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_3_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_3_SLICE,
				SliceCount:            3,
				InstanceCount:         2,
				MultiprocessorCount:   shared.B200_MIG_3_SLICE_SM_COUNT,
				SharedCopyEngineCount: shared.B200_MIG_7G_180GB_CE_COUNT,
				SharedDecoderCount:    1,
				SharedEncoderCount:    0,
				SharedJpegCount:       1,
				SharedOfaCount:        1,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_4_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_4_SLICE,
				SliceCount:            4,
				InstanceCount:         1,
				MultiprocessorCount:   shared.B200_MIG_4_SLICE_SM_COUNT,
				SharedCopyEngineCount: shared.B200_MIG_7G_180GB_CE_COUNT,
				SharedDecoderCount:    1,
				SharedEncoderCount:    0,
				SharedJpegCount:       1,
				SharedOfaCount:        1,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_7_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_7_SLICE,
				SliceCount:            7,
				InstanceCount:         1,
				MultiprocessorCount:   shared.B200_MIG_7_SLICE_SM_COUNT,
				SharedCopyEngineCount: shared.B200_MIG_7G_180GB_CE_COUNT,
				SharedDecoderCount:    1,
				SharedEncoderCount:    0,
				SharedJpegCount:       1,
				SharedOfaCount:        1,
			},
		},
	},
	GpuInstancePlacements: map[int][]nvml.GpuInstancePlacement{
		// Based on documented instance counts from B200 MIG documentation
		nvml.GPU_INSTANCE_PROFILE_1_SLICE: {
			{Start: 0, Size: 1}, {Start: 1, Size: 1}, {Start: 2, Size: 1}, {Start: 3, Size: 1},
			{Start: 4, Size: 1}, {Start: 5, Size: 1}, {Start: 6, Size: 1}, // 7 instances
		},
		nvml.GPU_INSTANCE_PROFILE_1_SLICE_ALL_ME: {
			{Start: 0, Size: 1}, // 1 instance only
		},
		nvml.GPU_INSTANCE_PROFILE_1_SLICE_REV2: {
			{Start: 0, Size: 2}, {Start: 2, Size: 2}, {Start: 4, Size: 2}, {Start: 6, Size: 2}, // 4 instances
		},
		nvml.GPU_INSTANCE_PROFILE_2_SLICE: {
			{Start: 0, Size: 2}, {Start: 2, Size: 2}, {Start: 4, Size: 2}, // 3 instances
		},
		nvml.GPU_INSTANCE_PROFILE_3_SLICE: {
			{Start: 0, Size: 4}, {Start: 4, Size: 4}, // 2 instances
		},
		nvml.GPU_INSTANCE_PROFILE_4_SLICE: {
			{Start: 0, Size: 4}, // 1 instance
		},
		nvml.GPU_INSTANCE_PROFILE_7_SLICE: {
			{Start: 0, Size: 8}, // 1 instance (full GPU)
		},
	},
	ComputeInstancePlacements: map[int]map[int][]nvml.ComputeInstancePlacement{
		nvml.GPU_INSTANCE_PROFILE_1_SLICE: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {{Start: 0, Size: 1}},
		},
		nvml.GPU_INSTANCE_PROFILE_1_SLICE_ALL_ME: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {{Start: 0, Size: 1}},
		},
		nvml.GPU_INSTANCE_PROFILE_1_SLICE_REV2: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {{Start: 0, Size: 1}},
		},
		nvml.GPU_INSTANCE_PROFILE_2_SLICE: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {{Start: 0, Size: 1}, {Start: 1, Size: 1}},
			nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE: {{Start: 0, Size: 2}},
		},
		nvml.GPU_INSTANCE_PROFILE_3_SLICE: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {{Start: 0, Size: 1}, {Start: 1, Size: 1}, {Start: 2, Size: 1}},
			nvml.COMPUTE_INSTANCE_PROFILE_3_SLICE: {{Start: 0, Size: 3}},
		},
		nvml.GPU_INSTANCE_PROFILE_4_SLICE: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {{Start: 0, Size: 1}, {Start: 1, Size: 1}, {Start: 2, Size: 1}, {Start: 3, Size: 1}},
			nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE: {{Start: 0, Size: 2}, {Start: 2, Size: 2}},
			nvml.COMPUTE_INSTANCE_PROFILE_4_SLICE: {{Start: 0, Size: 4}},
		},
		nvml.GPU_INSTANCE_PROFILE_7_SLICE: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {
				{Start: 0, Size: 1}, {Start: 1, Size: 1}, {Start: 2, Size: 1}, {Start: 3, Size: 1},
				{Start: 4, Size: 1}, {Start: 5, Size: 1}, {Start: 6, Size: 1},
			},
			nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE: {{Start: 0, Size: 2}, {Start: 2, Size: 2}, {Start: 4, Size: 2}},
			nvml.COMPUTE_INSTANCE_PROFILE_3_SLICE: {{Start: 0, Size: 3}, {Start: 4, Size: 3}},
			nvml.COMPUTE_INSTANCE_PROFILE_4_SLICE: {{Start: 0, Size: 4}},
			nvml.COMPUTE_INSTANCE_PROFILE_7_SLICE: {{Start: 0, Size: 8}},
		},
	},
}
