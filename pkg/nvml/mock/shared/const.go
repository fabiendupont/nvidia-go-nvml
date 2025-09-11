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

// GPU PCI Device IDs (device ID only, vendor ID 0x10DE is separate)
const (
	A100_PCI_DEVICE_ID = 0x20B0 // NVIDIA A100-SXM4-40GB
	H100_PCI_DEVICE_ID = 0x2330 // NVIDIA H100-SXM5-80GB
	H200_PCI_DEVICE_ID = 0x2339 // NVIDIA H200-SXM5-141GB
)

// GPU Memory Sizes in bytes
const (
	A100_TOTAL_MEMORY_BYTES = 42949672960  // 40GB
	H100_TOTAL_MEMORY_BYTES = 85899345920  // 80GB
	H200_TOTAL_MEMORY_BYTES = 151397597184 // 141GB
)

// GPU Memory Sizes in MB (for convenience)
const (
	A100_TOTAL_MEMORY_MB = A100_TOTAL_MEMORY_BYTES / (1024 * 1024) // 40960 MB
	H100_TOTAL_MEMORY_MB = H100_TOTAL_MEMORY_BYTES / (1024 * 1024) // 81920 MB
	H200_TOTAL_MEMORY_MB = H200_TOTAL_MEMORY_BYTES / (1024 * 1024) // 144384 MB
)

// CUDA Compute Capabilities
const (
	A100_CUDA_MAJOR = 8 // Ampere architecture
	A100_CUDA_MINOR = 0
	H100_CUDA_MAJOR = 9 // Hopper architecture
	H100_CUDA_MINOR = 0
	H200_CUDA_MAJOR = 9 // Hopper architecture
	H200_CUDA_MINOR = 0
)

// Driver and NVML Versions
const (
	DEFAULT_DRIVER_VERSION      = "550.54.15"
	DEFAULT_NVML_VERSION        = "12.550.54.15"
	DEFAULT_CUDA_DRIVER_VERSION = 12040
)

// Device Count per DGX Server
const (
	DGX_DEVICE_COUNT = 8 // Standard DGX configuration
)

// Memory conversion constants
const (
	BYTES_PER_MB = 1024 * 1024 // Bytes in a megabyte
)

// A100 MIG Memory Sizes in MB
const (
	A100_MIG_1G_5GB_MEMORY_MB  = 4864  // 1g.5gb profile
	A100_MIG_2G_10GB_MEMORY_MB = 9856  // 2g.10gb profile
	A100_MIG_3G_20GB_MEMORY_MB = 19968 // 3g.20gb profile
	A100_MIG_4G_20GB_MEMORY_MB = 19968 // 4g.20gb profile
	A100_MIG_7G_40GB_MEMORY_MB = 40192 // 7g.40gb profile (full GPU)
	// Revision variants use same memory as base profiles
	A100_MIG_1G_REV1_MEMORY_MB = 4864  // 1g revision 1
	A100_MIG_2G_REV1_MEMORY_MB = 9856  // 2g revision 1
	// Graphics variants - optimized memory allocation
	A100_MIG_1G_GFX_MEMORY_MB  = 4864  // 1g graphics variant
	A100_MIG_2G_GFX_MEMORY_MB  = 9856  // 2g graphics variant
	A100_MIG_4G_GFX_MEMORY_MB  = 19968 // 4g graphics variant
	// Memory engine variants
	A100_MIG_1G_NO_ME_MEMORY_MB  = 4864  // 1g no memory engines
	A100_MIG_2G_NO_ME_MEMORY_MB  = 9856  // 2g no memory engines
	A100_MIG_1G_ALL_ME_MEMORY_MB = 9856  // 1g all memory engines
	A100_MIG_2G_ALL_ME_MEMORY_MB = 15872 // 2g all memory engines
)

// MIG Profile Instance Counts
const (
	MIG_1_SLICE_INSTANCE_COUNT = 7 // Number of 1-slice instances possible
	MIG_2_SLICE_INSTANCE_COUNT = 3 // Number of 2-slice instances possible
	MIG_3_SLICE_INSTANCE_COUNT = 2 // Number of 3-slice instances possible
	MIG_4_SLICE_INSTANCE_COUNT = 1 // Number of 4-slice instances possible
	MIG_7_SLICE_INSTANCE_COUNT = 1 // Number of 7-slice instances possible (full GPU)
)

// MIG Profile Multiprocessor Counts
const (
	MIG_1_SLICE_SM_COUNT = 14 // Streaming multiprocessors per 1-slice
	MIG_2_SLICE_SM_COUNT = 28 // Streaming multiprocessors per 2-slice
	MIG_3_SLICE_SM_COUNT = 42 // Streaming multiprocessors per 3-slice
	MIG_4_SLICE_SM_COUNT = 56 // Streaming multiprocessors per 4-slice
	MIG_7_SLICE_SM_COUNT = 98 // Streaming multiprocessors per 7-slice (full GPU)
)

// MIG Profile Copy Engine Counts
const (
	MIG_1_SLICE_CE_COUNT = 1 // Copy engines per 1-slice
	MIG_2_SLICE_CE_COUNT = 2 // Copy engines per 2-slice
	MIG_3_SLICE_CE_COUNT = 3 // Copy engines per 3-slice
	MIG_4_SLICE_CE_COUNT = 4 // Copy engines per 4-slice
	MIG_7_SLICE_CE_COUNT = 7 // Copy engines per 7-slice
)

// H100 MIG Memory Sizes in MB
const (
	H100_MIG_1G_10GB_MEMORY_MB = 9856  // 1g.10gb profile
	H100_MIG_1G_20GB_MEMORY_MB = 19968 // 1g.20gb profile
	H100_MIG_2G_20GB_MEMORY_MB = 19968 // 2g.20gb profile
	H100_MIG_3G_40GB_MEMORY_MB = 40448 // 3g.40gb profile
	H100_MIG_4G_40GB_MEMORY_MB = 40448 // 4g.40gb profile
	H100_MIG_6G_60GB_MEMORY_MB = 61440 // 6g.60gb profile
	H100_MIG_7G_80GB_MEMORY_MB = 81408 // 7g.80gb profile (full GPU)
	H100_MIG_8G_80GB_MEMORY_MB = 81408 // 8g.80gb profile (alternative full GPU)
	// Revision variants use same memory as base profiles
	H100_MIG_1G_REV1_MEMORY_MB = 9856  // 1g revision 1
	H100_MIG_1G_REV2_MEMORY_MB = 19968 // 1g revision 2 (larger memory)
	H100_MIG_2G_REV1_MEMORY_MB = 19968 // 2g revision 1
	// Graphics variants - optimized memory allocation
	H100_MIG_1G_GFX_MEMORY_MB  = 9856  // 1g graphics variant
	H100_MIG_2G_GFX_MEMORY_MB  = 19968 // 2g graphics variant
	H100_MIG_4G_GFX_MEMORY_MB  = 40448 // 4g graphics variant
	// Memory engine variants
	H100_MIG_1G_NO_ME_MEMORY_MB  = 9856  // 1g no memory engines
	H100_MIG_2G_NO_ME_MEMORY_MB  = 19968 // 2g no memory engines
	H100_MIG_1G_ALL_ME_MEMORY_MB = 19968 // 1g all memory engines
	H100_MIG_2G_ALL_ME_MEMORY_MB = 35840 // 2g all memory engines
)

// H100 MIG Streaming Multiprocessor Counts
const (
	H100_MIG_1_SLICE_SM_COUNT = 16 // H100 has more SMs per slice than A100
	H100_MIG_2_SLICE_SM_COUNT = 32
	H100_MIG_3_SLICE_SM_COUNT = 48
	H100_MIG_4_SLICE_SM_COUNT = 64
	H100_MIG_6_SLICE_SM_COUNT = 96
	H100_MIG_7_SLICE_SM_COUNT = 112 // H100 has 114 total SMs, 112 available for MIG
	H100_MIG_8_SLICE_SM_COUNT = 112 // Same as 7-slice, alternative full GPU config
)

// H100 MIG Copy Engine Counts
const (
	H100_MIG_1_SLICE_CE_COUNT = 1
	H100_MIG_2_SLICE_CE_COUNT = 2
	H100_MIG_3_SLICE_CE_COUNT = 3
	H100_MIG_4_SLICE_CE_COUNT = 4
	H100_MIG_6_SLICE_CE_COUNT = 6
	H100_MIG_7_SLICE_CE_COUNT = 7
	H100_MIG_8_SLICE_CE_COUNT = 7 // Same as 7-slice
)

// A100 MIG P2P Support
const (
	A100_MIG_P2P_SUPPORT = 0 // A100 doesn't support P2P in MIG
)

// H100 MIG P2P Support
const (
	H100_MIG_P2P_SUPPORT = 0 // H100 doesn't support P2P in MIG
)

// H200 MIG Memory Sizes in MB
const (
	H200_MIG_1G_18GB_MEMORY_MB  = 17920  // 1g.18gb profile
	H200_MIG_1G_35GB_MEMORY_MB  = 35840  // 1g.35gb profile
	H200_MIG_2G_35GB_MEMORY_MB  = 35840  // 2g.35gb profile
	H200_MIG_3G_71GB_MEMORY_MB  = 72192  // 3g.71gb profile
	H200_MIG_4G_71GB_MEMORY_MB  = 72192  // 4g.71gb profile
	H200_MIG_6G_106GB_MEMORY_MB = 108544 // 6g.106gb profile
	H200_MIG_7G_141GB_MEMORY_MB = 144384 // 7g.141gb profile (full GPU)
	H200_MIG_8G_141GB_MEMORY_MB = 144384 // 8g.141gb profile (alternative full GPU)
	// Revision variants use same memory as base profiles
	H200_MIG_1G_REV1_MEMORY_MB = 17920  // 1g revision 1
	H200_MIG_1G_REV2_MEMORY_MB = 35840  // 1g revision 2 (larger memory)
	H200_MIG_2G_REV1_MEMORY_MB = 35840  // 2g revision 1
	// Graphics variants - optimized memory allocation
	H200_MIG_1G_GFX_MEMORY_MB  = 17920  // 1g graphics variant
	H200_MIG_2G_GFX_MEMORY_MB  = 35840  // 2g graphics variant
	H200_MIG_4G_GFX_MEMORY_MB  = 72192  // 4g graphics variant
	// Memory engine variants
	H200_MIG_1G_NO_ME_MEMORY_MB  = 17920  // 1g no memory engines
	H200_MIG_2G_NO_ME_MEMORY_MB  = 35840  // 2g no memory engines
	H200_MIG_1G_ALL_ME_MEMORY_MB = 35840  // 1g all memory engines
	H200_MIG_2G_ALL_ME_MEMORY_MB = 53760  // 2g all memory engines
)

// H200 MIG Streaming Multiprocessor Counts (same as H100)
const (
	H200_MIG_1_SLICE_SM_COUNT = 16 // H200 uses same SM counts as H100
	H200_MIG_2_SLICE_SM_COUNT = 32
	H200_MIG_3_SLICE_SM_COUNT = 48
	H200_MIG_4_SLICE_SM_COUNT = 64
	H200_MIG_6_SLICE_SM_COUNT = 96
	H200_MIG_7_SLICE_SM_COUNT = 112
	H200_MIG_8_SLICE_SM_COUNT = 112 // Same as 7-slice, alternative full GPU config
)

// H200 MIG Copy Engine Counts (same as H100)
const (
	H200_MIG_1_SLICE_CE_COUNT = 1
	H200_MIG_2_SLICE_CE_COUNT = 2
	H200_MIG_3_SLICE_CE_COUNT = 3
	H200_MIG_4_SLICE_CE_COUNT = 4
	H200_MIG_6_SLICE_CE_COUNT = 6
	H200_MIG_7_SLICE_CE_COUNT = 7
	H200_MIG_8_SLICE_CE_COUNT = 7 // Same as 7-slice
)

// H200 MIG P2P Support
const (
	H200_MIG_P2P_SUPPORT = 0 // H200 doesn't support P2P in MIG
)

// MIG Profile Placement Sizes (universal across GPU generations)
const (
	MIG_1_SLICE_PLACEMENT_SIZE = 1
	MIG_2_SLICE_PLACEMENT_SIZE = 2
	MIG_3_SLICE_PLACEMENT_SIZE = 4 // 3-slice profiles use placement size 4
	MIG_4_SLICE_PLACEMENT_SIZE = 4
	MIG_7_SLICE_PLACEMENT_SIZE = 8 // 7-slice profiles use placement size 8
)
