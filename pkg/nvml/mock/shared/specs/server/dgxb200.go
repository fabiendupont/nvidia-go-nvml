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

// GetDGXB200DeviceCount returns the number of devices in a DGX B200 system
func GetDGXB200DeviceCount() int {
	return shared.DGX_DEVICE_COUNT
}

// GetDGXB200Device returns a mock device configured for DGX B200 at the specified index
func GetDGXB200Device(index int) nvml.Device {
	if index < 0 || index >= shared.DGX_DEVICE_COUNT {
		panic("Invalid device index for DGX B200")
	}
	return shared.NewMockDevice(gpu.B200Spec, index)
}

// NewDGXB200 creates a new DGX B200 mock GPU server
func NewDGXB200() *shared.MockGPU {
	spec := gpu.B200Spec
	spec.DeviceCount = shared.DGX_DEVICE_COUNT
	return shared.NewMockGPU(spec)
}
