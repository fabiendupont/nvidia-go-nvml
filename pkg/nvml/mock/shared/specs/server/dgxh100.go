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
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/shared"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/shared/specs/gpu"
)

// DGXH100Spec defines the complete server specification for DGX H100
var DGXH100Spec = shared.GPUSpec{
	Name:                gpu.H100Spec.Name,
	Architecture:        gpu.H100Spec.Architecture,
	Brand:               gpu.H100Spec.Brand,
	PciDeviceId:         gpu.H100Spec.PciDeviceId,
	PciBusIDBase:        gpu.H100Spec.PciBusIDBase,
	DriverVersion:       gpu.H100Spec.DriverVersion,
	NvmlVersion:         gpu.H100Spec.NvmlVersion,
	CudaDriverVersion:   gpu.H100Spec.CudaDriverVersion,
	CudaCapabilityMajor: gpu.H100Spec.CudaCapabilityMajor,
	CudaCapabilityMinor: gpu.H100Spec.CudaCapabilityMinor,
	TotalMemoryMB:       gpu.H100Spec.TotalMemoryMB,
	DeviceCount:         shared.DGX_DEVICE_COUNT, // DGX H100 has 8 GPUs
	MIGProfiles:         gpu.H100Spec.MIGProfiles,
}

// NewDGXH100 creates a new DGX H100 mock server
func NewDGXH100() *shared.MockGPU {
	return shared.NewMockGPU(DGXH100Spec)
}

// GetDGXH100DeviceCount returns the number of devices in a DGX H100 server
func GetDGXH100DeviceCount() int {
	return shared.DGX_DEVICE_COUNT
}

// GetDGXH100Device returns device information for a DGX H100 server
func GetDGXH100Device(i int) nvml.Device {
	if i < 0 || i >= shared.DGX_DEVICE_COUNT {
		panic("Invalid device index")
	}
	return shared.NewMockDevice(gpu.H100Spec, i)
}
