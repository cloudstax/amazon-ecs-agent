// Copyright 2014-2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package config

import (
	"encoding/json"

	"github.com/aws/amazon-ecs-agent/agent/engine/dockerclient"
)

type Config struct {
	// DEPRECATED
	// ClusterArn is the Name or full ARN of a Cluster to register into. It has
	// been deprecated (and will eventually be removed) in favor of Cluster
	ClusterArn string `deprecated:"Please use Cluster instead"`
	// Cluster can either be the Name or full ARN of a Cluster. This is the
	// cluster the agent should register this ContainerInstance into. If this
	// value is not set, it will default to "default"
	Cluster string `trim:"true"`
	// APIEndpoint is the endpoint, such as "ecs.us-east-1.amazonaws.com", to
	// make calls against. If this value is not set, it will default to the
	// endpoint for your current AWSRegion
	APIEndpoint string `trim:"true"`
	// DockerEndpoint is the address the agent will attempt to connect to the
	// Docker daemon at. This should have the same value as "DOCKER_HOST"
	// normally would to interact with the daemon. It defaults to
	// unix:///var/run/docker.sock
	DockerEndpoint string
	// AWSRegion is the region to run in (such as "us-east-1"). This value will
	// be inferred from the EC2 metadata service, but if it cannot be found this
	// will be fatal.
	AWSRegion string `missing:"fatal" trim:"true"`

	// ReservedPorts is an array of ports which should be registerd as
	// unavailable. If not set, they default to [22,2375,2376,51678].
	ReservedPorts []uint16
	// ReservedPortsUDP is an array of UDP ports which should be registered as
	// unavailable. If not set, it defaults to [].
	ReservedPortsUDP []uint16

	// DataDir is the directory data is saved to in order to preserve state
	// across agent restarts. It is only used if "Checkpoint" is true as well.
	DataDir string
	// Checkpoint configures whether data should be periodically to a checkpoint
	// file, in DataDir, such that on instance or agent restarts it will resume
	// as the same ContainerInstance. It defaults to false.
	Checkpoint bool

	// EngineAuthType configures what type of data is in EngineAuthData.
	// Supported types, right now, can be found in the dockerauth package: https://godoc.org/github.com/aws/amazon-ecs-agent/agent/engine/dockerauth
	EngineAuthType string `trim:"true"`
	// EngineAuthData contains authentication data. Please see the documentation
	// for EngineAuthType for more information.
	EngineAuthData json.RawMessage

	// UpdatesEnabled specifies whether updates should be applied to this agent.
	// Default true
	UpdatesEnabled bool
	// UpdateDownloadDir specifies where new agent versions should be placed
	// within the container in order for the external updating process to
	// correctly handle them.
	UpdateDownloadDir string

	// DisableMetrics configures whether task utilization metrics should be
	// sent to the ECS telemetry endpoint
	DisableMetrics bool

	// DockerGraphPath specifies the path for docker graph directory.
	DockerGraphPath string

	// ReservedMemory specifies the amount of memory (in MB) to reserve for things
	// other than containers managed by ECS
	ReservedMemory uint16

	// AvailableLoggingDrivers specifies the logging drivers available for use
	// with Docker.  If not set, it defaults to ["json-file"].
	AvailableLoggingDrivers []dockerclient.LoggingDriver

	// SELinxuCapable specifies whether the Agent is capable of using SELinux
	// security options
	SELinuxCapable bool

	// AppArmorCapable specifies whether the Agent is capable of using AppArmor
	// security options
	AppArmorCapable bool
}
