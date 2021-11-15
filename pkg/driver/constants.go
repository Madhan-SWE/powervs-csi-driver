/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package driver

//DONE

// constants of keys in PublishContext
const (
	WWNKey = "wwn"
)

// constants of keys in PublishContext
const (
	// devicePathKey represents key for device path in PublishContext
	// devicePath is the device path where the volume is attached to
	DevicePathKey = "devicePath"
)

// constants of keys in VolumeContext
const (
	// VolumeAttributePartition represents key for partition config in VolumeContext
	// this represents the partition number on a device used to mount
	VolumeAttributePartition = "partition"
)

// constants of keys in volume parameters
const (
	// VolumeTypeKey represents key for volume type
	VolumeTypeKey = "type"

	// IopsPerGBKey represents key for IOPS per GB
	IopsPerGBKey = "iopspergb"

	// Iops represents key for IOPS for volume
	IopsKey = "iops"

	// ThroughputKey represents key for throughput
	ThroughputKey = "throughput"

	// EncryptedKey represents key for whether filesystem is encrypted
	EncryptedKey = "encrypted"
)

// constants for default command line flag values
const (
	DefaultCSIEndpoint = "unix://tmp/csi.sock"
)
