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

package dgxb200

import (
	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

// MIGProfiles holds the profile information for GIs and CIs in this B200+ mock server.
// B200+ has massive memory capabilities (192GB) enabling unprecedented MIG flexibility
// and advanced fabric integration for next-generation AI workloads.
var MIGProfiles = struct {
	GpuInstanceProfiles     map[int]nvml.GpuInstanceProfileInfo
	ComputeInstanceProfiles map[int]map[int]nvml.ComputeInstanceProfileInfo
}{
	GpuInstanceProfiles: map[int]nvml.GpuInstanceProfileInfo{
		// B200+ Massive Memory MIG Profiles (192GB memory - 36% more than H200's 141GB)
		nvml.GPU_INSTANCE_PROFILE_1_SLICE: {
			Id:                  nvml.GPU_INSTANCE_PROFILE_1_SLICE,
			IsP2pSupported:      1, // B200+ supports advanced P2P in MIG
			SliceCount:          1,
			InstanceCount:       7,
			MultiprocessorCount: 14,
			CopyEngineCount:     1,
			DecoderCount:        0,
			EncoderCount:        0,
			JpegCount:           0,
			OfaCount:            0,
			MemorySizeMB:        19200, // 1g.19gb profile (massive compared to H200's 1g.14gb)
		},
		nvml.GPU_INSTANCE_PROFILE_1_SLICE_REV1: {
			Id:                  nvml.GPU_INSTANCE_PROFILE_1_SLICE_REV1,
			IsP2pSupported:      1,
			SliceCount:          1,
			InstanceCount:       1,
			MultiprocessorCount: 14,
			CopyEngineCount:     1,
			DecoderCount:        1,
			EncoderCount:        1,
			JpegCount:           1,
			OfaCount:            0,
			MemorySizeMB:        19200, // 1g.19gb profile
		},
		nvml.GPU_INSTANCE_PROFILE_2_SLICE: {
			Id:                  nvml.GPU_INSTANCE_PROFILE_2_SLICE,
			IsP2pSupported:      1,
			SliceCount:          2,
			InstanceCount:       3,
			MultiprocessorCount: 28,
			CopyEngineCount:     2,
			DecoderCount:        0,
			EncoderCount:        0,
			JpegCount:           0,
			OfaCount:            0,
			MemorySizeMB:        38400, // 2g.38gb profile (massive compared to H200's 2g.28gb)
		},
		nvml.GPU_INSTANCE_PROFILE_2_SLICE_REV1: {
			Id:                  nvml.GPU_INSTANCE_PROFILE_2_SLICE_REV1,
			IsP2pSupported:      1,
			SliceCount:          2,
			InstanceCount:       1,
			MultiprocessorCount: 28,
			CopyEngineCount:     2,
			DecoderCount:        2,
			EncoderCount:        2,
			JpegCount:           2,
			OfaCount:            0,
			MemorySizeMB:        38400, // 2g.38gb profile
		},
		nvml.GPU_INSTANCE_PROFILE_3_SLICE: {
			Id:                  nvml.GPU_INSTANCE_PROFILE_3_SLICE,
			IsP2pSupported:      1,
			SliceCount:          3,
			InstanceCount:       2,
			MultiprocessorCount: 42,
			CopyEngineCount:     3,
			DecoderCount:        0,
			EncoderCount:        0,
			JpegCount:           0,
			OfaCount:            0,
			MemorySizeMB:        57600, // 3g.57gb profile (massive compared to H200's 3g.42gb)
		},
		nvml.GPU_INSTANCE_PROFILE_4_SLICE: {
			Id:                  nvml.GPU_INSTANCE_PROFILE_4_SLICE,
			IsP2pSupported:      1,
			SliceCount:          4,
			InstanceCount:       1,
			MultiprocessorCount: 56,
			CopyEngineCount:     4,
			DecoderCount:        0,
			EncoderCount:        0,
			JpegCount:           0,
			OfaCount:            0,
			MemorySizeMB:        76800, // 4g.76gb profile (massive compared to H200's 4g.56gb)
		},
		nvml.GPU_INSTANCE_PROFILE_7_SLICE: {
			Id:                  nvml.GPU_INSTANCE_PROFILE_7_SLICE,
			IsP2pSupported:      1,
			SliceCount:          7,
			InstanceCount:       1,
			MultiprocessorCount: 98,
			CopyEngineCount:     7,
			DecoderCount:        0,
			EncoderCount:        0,
			JpegCount:           0,
			OfaCount:            0,
			MemorySizeMB:        192000, // 7g.192gb profile (full GPU - massive compared to H200's 7g.141gb)
		},
	},
	ComputeInstanceProfiles: map[int]map[int]nvml.ComputeInstanceProfileInfo{
		// 1g.19gb profile compute instances
		nvml.GPU_INSTANCE_PROFILE_1_SLICE: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
				SliceCount:            1,
				InstanceCount:         1,
				MultiprocessorCount:   14,
				SharedCopyEngineCount: 1,
				SharedDecoderCount:    0,
				SharedEncoderCount:    0,
				SharedJpegCount:       0,
				SharedOfaCount:        0,
			},
		},
		nvml.GPU_INSTANCE_PROFILE_1_SLICE_REV1: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
				SliceCount:            1,
				InstanceCount:         1,
				MultiprocessorCount:   14,
				SharedCopyEngineCount: 1,
				SharedDecoderCount:    1,
				SharedEncoderCount:    1,
				SharedJpegCount:       1,
				SharedOfaCount:        0,
			},
		},
		// 2g.38gb profile compute instances
		nvml.GPU_INSTANCE_PROFILE_2_SLICE: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
				SliceCount:            1,
				InstanceCount:         2,
				MultiprocessorCount:   14,
				SharedCopyEngineCount: 1,
				SharedDecoderCount:    0,
				SharedEncoderCount:    0,
				SharedJpegCount:       0,
				SharedOfaCount:        0,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE,
				SliceCount:            2,
				InstanceCount:         1,
				MultiprocessorCount:   28,
				SharedCopyEngineCount: 2,
				SharedDecoderCount:    0,
				SharedEncoderCount:    0,
				SharedJpegCount:       0,
				SharedOfaCount:        0,
			},
		},
		nvml.GPU_INSTANCE_PROFILE_2_SLICE_REV1: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
				SliceCount:            1,
				InstanceCount:         2,
				MultiprocessorCount:   14,
				SharedCopyEngineCount: 1,
				SharedDecoderCount:    1,
				SharedEncoderCount:    1,
				SharedJpegCount:       1,
				SharedOfaCount:        0,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE,
				SliceCount:            2,
				InstanceCount:         1,
				MultiprocessorCount:   28,
				SharedCopyEngineCount: 2,
				SharedDecoderCount:    2,
				SharedEncoderCount:    2,
				SharedJpegCount:       2,
				SharedOfaCount:        0,
			},
		},
		// 3g.57gb profile compute instances
		nvml.GPU_INSTANCE_PROFILE_3_SLICE: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
				SliceCount:            1,
				InstanceCount:         3,
				MultiprocessorCount:   14,
				SharedCopyEngineCount: 1,
				SharedDecoderCount:    0,
				SharedEncoderCount:    0,
				SharedJpegCount:       0,
				SharedOfaCount:        0,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE,
				SliceCount:            2,
				InstanceCount:         1,
				MultiprocessorCount:   28,
				SharedCopyEngineCount: 2,
				SharedDecoderCount:    0,
				SharedEncoderCount:    0,
				SharedJpegCount:       0,
				SharedOfaCount:        0,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_3_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_3_SLICE,
				SliceCount:            3,
				InstanceCount:         1,
				MultiprocessorCount:   42,
				SharedCopyEngineCount: 3,
				SharedDecoderCount:    0,
				SharedEncoderCount:    0,
				SharedJpegCount:       0,
				SharedOfaCount:        0,
			},
		},
		// 4g.76gb profile compute instances
		nvml.GPU_INSTANCE_PROFILE_4_SLICE: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
				SliceCount:            1,
				InstanceCount:         4,
				MultiprocessorCount:   14,
				SharedCopyEngineCount: 1,
				SharedDecoderCount:    0,
				SharedEncoderCount:    0,
				SharedJpegCount:       0,
				SharedOfaCount:        0,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE,
				SliceCount:            2,
				InstanceCount:         2,
				MultiprocessorCount:   28,
				SharedCopyEngineCount: 2,
				SharedDecoderCount:    0,
				SharedEncoderCount:    0,
				SharedJpegCount:       0,
				SharedOfaCount:        0,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_4_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_4_SLICE,
				SliceCount:            4,
				InstanceCount:         1,
				MultiprocessorCount:   56,
				SharedCopyEngineCount: 4,
				SharedDecoderCount:    0,
				SharedEncoderCount:    0,
				SharedJpegCount:       0,
				SharedOfaCount:        0,
			},
		},
		// 7g.192gb profile compute instances (full GPU)
		nvml.GPU_INSTANCE_PROFILE_7_SLICE: {
			nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_1_SLICE,
				SliceCount:            1,
				InstanceCount:         7,
				MultiprocessorCount:   14,
				SharedCopyEngineCount: 1,
				SharedDecoderCount:    0,
				SharedEncoderCount:    0,
				SharedJpegCount:       0,
				SharedOfaCount:        0,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_2_SLICE,
				SliceCount:            2,
				InstanceCount:         3,
				MultiprocessorCount:   28,
				SharedCopyEngineCount: 2,
				SharedDecoderCount:    0,
				SharedEncoderCount:    0,
				SharedJpegCount:       0,
				SharedOfaCount:        0,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_4_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_4_SLICE,
				SliceCount:            4,
				InstanceCount:         1,
				MultiprocessorCount:   56,
				SharedCopyEngineCount: 4,
				SharedDecoderCount:    0,
				SharedEncoderCount:    0,
				SharedJpegCount:       0,
				SharedOfaCount:        0,
			},
			nvml.COMPUTE_INSTANCE_PROFILE_7_SLICE: {
				Id:                    nvml.COMPUTE_INSTANCE_PROFILE_7_SLICE,
				SliceCount:            7,
				InstanceCount:         1,
				MultiprocessorCount:   98,
				SharedCopyEngineCount: 7,
				SharedDecoderCount:    0,
				SharedEncoderCount:    0,
				SharedJpegCount:       0,
				SharedOfaCount:        0,
			},
		},
	},
}
