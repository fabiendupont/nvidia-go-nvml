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

// DGXH200Spec defines the complete server specification for DGX H200
var DGXH200Spec = shared.GPUSpec{
	Name:                gpu.H200Spec.Name,
	Architecture:        gpu.H200Spec.Architecture,
	Brand:               gpu.H200Spec.Brand,
	PciDeviceId:         gpu.H200Spec.PciDeviceId,
	PciBusIDBase:        gpu.H200Spec.PciBusIDBase,
	DriverVersion:       gpu.H200Spec.DriverVersion,
	NvmlVersion:         gpu.H200Spec.NvmlVersion,
	CudaDriverVersion:   gpu.H200Spec.CudaDriverVersion,
	CudaCapabilityMajor: gpu.H200Spec.CudaCapabilityMajor,
	CudaCapabilityMinor: gpu.H200Spec.CudaCapabilityMinor,
	TotalMemoryMB:       gpu.H200Spec.TotalMemoryMB,
	DeviceCount:         shared.DGX_DEVICE_COUNT, // DGX H200 has 8 GPUs
	MIGProfiles:         gpu.H200Spec.MIGProfiles,
}

// NewDGXH200 creates a new DGX H200 mock server
func NewDGXH200() *shared.MockGPU {
	return shared.NewMockGPU(DGXH200Spec)
}

// GetDGXH200DeviceCount returns the number of devices in a DGX H200 server
func GetDGXH200DeviceCount() int {
	return shared.DGX_DEVICE_COUNT
}

// GetDGXH200Device returns device information for a DGX H200 server
func GetDGXH200Device(i int) nvml.Device {
	if i < 0 || i >= shared.DGX_DEVICE_COUNT {
		panic("Invalid device index")
	}
	return shared.NewMockDevice(gpu.H200Spec, i)
}
