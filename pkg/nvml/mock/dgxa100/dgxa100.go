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
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/shared"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/shared/specs/gpu"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock/shared/specs/server"
)

// Server is a type alias for the shared MockGPU implementation
type Server = shared.MockGPU

// New creates a new DGX A100 mock server
func New() *Server {
	return server.NewDGXA100()
}

// Legacy type aliases for backward compatibility
type Device = shared.MockDevice
type GpuInstance = shared.MockGpuInstance
type ComputeInstance = shared.MockComputeInstance

// Legacy constructor functions for backward compatibility
func NewDevice(index int) *Device {
	return shared.NewMockDevice(gpu.A100Spec, index)
}

func NewGpuInstance(info nvml.GpuInstanceInfo) *GpuInstance {
	return shared.NewMockGpuInstance(info, gpu.A100Spec.MIGProfiles)
}

func NewComputeInstance(info nvml.ComputeInstanceInfo) *ComputeInstance {
	return shared.NewMockComputeInstance(info)
}

// Legacy MIG profile variables for backward compatibility
// These reference the shared architecture profiles to ensure identical behavior
var (
	MIGProfiles = struct {
		GpuInstanceProfiles     map[int]nvml.GpuInstanceProfileInfo
		ComputeInstanceProfiles map[int]map[int]nvml.ComputeInstanceProfileInfo
	}{
		GpuInstanceProfiles:     gpu.A100Spec.MIGProfiles.GpuInstanceProfiles,
		ComputeInstanceProfiles: gpu.A100Spec.MIGProfiles.ComputeInstanceProfiles,
	}

	MIGPlacements = struct {
		GpuInstancePossiblePlacements     map[int][]nvml.GpuInstancePlacement
		ComputeInstancePossiblePlacements map[int]map[int][]nvml.ComputeInstancePlacement
	}{
		GpuInstancePossiblePlacements:     gpu.A100Spec.MIGProfiles.GpuInstancePlacements,
		ComputeInstancePossiblePlacements: gpu.A100Spec.MIGProfiles.ComputeInstancePlacements,
	}
)
