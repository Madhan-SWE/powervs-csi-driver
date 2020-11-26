/*
Copyright 2020 The Kubernetes Authors.

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

package options

//DONE

import (
	"flag"

	"github.com/ppc64le-cloud/powervs-csi-driver/pkg/driver"
)

// ServerOptions contains options and configuration settings for the driver server.
type ServerOptions struct {
	// Endpoint is the endpoint that the driver server should listen on.
	Endpoint string
	// PowerVS Cloud Instance ID
	PVMCloudInstanceID string
	// Debug
	Debug bool
}

func (s *ServerOptions) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&s.Endpoint, "endpoint", driver.DefaultCSIEndpoint, "Endpoint for the CSI driver server")
	fs.StringVar(&s.PVMCloudInstanceID, "pvm-cloud-instance-id","","ID of the PVM Cloud Instance ID")
	fs.BoolVar(&s.Debug, "debug",false,"Debug option PowerVS client(Prints API requests and replies)")
}
