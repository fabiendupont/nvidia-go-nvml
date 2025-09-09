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
	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

// MIGProfiles holds the profile information for GIs and CIs in this H200+ mock server.
// H200+ has enhanced MIG capabilities compared to H100, with 141GB memory enabling
// larger MIG instances and improved fabric integration capabilities.
var MIGProfiles = struct {
	GpuInstanceProfiles     map[int]nvml.GpuInstanceProfileInfo
	ComputeInstanceProfiles map[int]map[int]nvml.ComputeInstanceProfileInfo
}{
	GpuInstanceProfiles: map[int]nvml.GpuInstanceProfileInfo{
		// H200+ Enhanced MIG Profiles (141GB memory - 76% more than H100's 80GB)
		nvml.GPU_INSTANCE_PROFILE_1_SLICE: {
			Id:                  nvml.GPU_INSTANCE_PROFILE_1_SLICE,
			IsP2pSupported:      1, // H200+ supports P2P in MIG
			SliceCount:          1,
			InstanceCount:       7,
			MultiprocessorCount: 14,
			CopyEngineCount:     1,
			DecoderCount:        0,
			EncoderCount:        0,
			JpegCount:           0,
			OfaCount:            0,
			MemorySizeMB:        14100, // 1g.14gb profile (enhanced from H100's 1g.10gb)
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
			MemorySizeMB:        14100, // 1g.14gb profile
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
			MemorySizeMB:        28200, // 2g.28gb profile (enhanced from H100's 2g.20gb)
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
			MemorySizeMB:        28200, // 2g.28gb profile
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
			MemorySizeMB:        42300, // 3g.42gb profile (enhanced from H100's 3g.30gb)
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
			MemorySizeMB:        56400, // 4g.56gb profile (enhanced from H100's 4g.40gb)
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
			MemorySizeMB:        141000, // 7g.141gb profile (full GPU - enhanced from H100's 7g.80gb)
		},
	},
	ComputeInstanceProfiles: map[int]map[int]nvml.ComputeInstanceProfileInfo{
		// 1g.14gb profile compute instances
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
		// 2g.28gb profile compute instances
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
		// 3g.42gb profile compute instances
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
		// 4g.56gb profile compute instances
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
		// 7g.141gb profile compute instances (full GPU)
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
