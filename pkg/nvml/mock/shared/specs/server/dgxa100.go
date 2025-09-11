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
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/shared"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/shared/specs/gpu"
)

// DGXA100Spec defines the complete server specification for DGX A100
var DGXA100Spec = shared.GPUSpec{
	Name:                gpu.A100Spec.Name,
	Architecture:        gpu.A100Spec.Architecture,
	Brand:               gpu.A100Spec.Brand,
	PciDeviceId:         gpu.A100Spec.PciDeviceId,
	DriverVersion:       gpu.A100Spec.DriverVersion,
	NvmlVersion:         gpu.A100Spec.NvmlVersion,
	CudaDriverVersion:   gpu.A100Spec.CudaDriverVersion,
	CudaCapabilityMajor: gpu.A100Spec.CudaCapabilityMajor,
	CudaCapabilityMinor: gpu.A100Spec.CudaCapabilityMinor,
	TotalMemoryMB:       gpu.A100Spec.TotalMemoryMB,
	DeviceCount:         shared.DGX_DEVICE_COUNT, // DGX A100 has 8 GPUs
	MIGProfiles:         gpu.A100Spec.MIGProfiles,
}

// NewDGXA100 creates a new DGX A100 mock server
func NewDGXA100() *shared.MockGPU {
	return shared.NewMockGPU(DGXA100Spec)
}
