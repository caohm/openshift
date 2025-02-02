// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package analysisservices

import original "github.com/Azure/azure-sdk-for-go/services/analysisservices/mgmt/2017-08-01/analysisservices"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type ConnectionMode = original.ConnectionMode

const (
	All      ConnectionMode = original.All
	ReadOnly ConnectionMode = original.ReadOnly
)

type ProvisioningState = original.ProvisioningState

const (
	Deleting     ProvisioningState = original.Deleting
	Failed       ProvisioningState = original.Failed
	Paused       ProvisioningState = original.Paused
	Pausing      ProvisioningState = original.Pausing
	Preparing    ProvisioningState = original.Preparing
	Provisioning ProvisioningState = original.Provisioning
	Resuming     ProvisioningState = original.Resuming
	Scaling      ProvisioningState = original.Scaling
	Succeeded    ProvisioningState = original.Succeeded
	Suspended    ProvisioningState = original.Suspended
	Suspending   ProvisioningState = original.Suspending
	Updating     ProvisioningState = original.Updating
)

type SkuTier = original.SkuTier

const (
	Basic       SkuTier = original.Basic
	Development SkuTier = original.Development
	Standard    SkuTier = original.Standard
)

type State = original.State

const (
	StateDeleting     State = original.StateDeleting
	StateFailed       State = original.StateFailed
	StatePaused       State = original.StatePaused
	StatePausing      State = original.StatePausing
	StatePreparing    State = original.StatePreparing
	StateProvisioning State = original.StateProvisioning
	StateResuming     State = original.StateResuming
	StateScaling      State = original.StateScaling
	StateSucceeded    State = original.StateSucceeded
	StateSuspended    State = original.StateSuspended
	StateSuspending   State = original.StateSuspending
	StateUpdating     State = original.StateUpdating
)

type Status = original.Status

const (
	Live Status = original.Live
)

type CheckServerNameAvailabilityParameters = original.CheckServerNameAvailabilityParameters
type CheckServerNameAvailabilityResult = original.CheckServerNameAvailabilityResult
type ErrorResponse = original.ErrorResponse
type GatewayDetails = original.GatewayDetails
type GatewayError = original.GatewayError
type GatewayListStatusError = original.GatewayListStatusError
type GatewayListStatusLive = original.GatewayListStatusLive
type IPv4FirewallRule = original.IPv4FirewallRule
type IPv4FirewallSettings = original.IPv4FirewallSettings
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationStatus = original.OperationStatus
type Resource = original.Resource
type ResourceSku = original.ResourceSku
type Server = original.Server
type ServerAdministrators = original.ServerAdministrators
type ServerMutableProperties = original.ServerMutableProperties
type ServerProperties = original.ServerProperties
type Servers = original.Servers
type ServersCreateFuture = original.ServersCreateFuture
type ServersDeleteFuture = original.ServersDeleteFuture
type ServersResumeFuture = original.ServersResumeFuture
type ServersSuspendFuture = original.ServersSuspendFuture
type ServersUpdateFuture = original.ServersUpdateFuture
type ServerUpdateParameters = original.ServerUpdateParameters
type SkuDetailsForExistingResource = original.SkuDetailsForExistingResource
type SkuEnumerationForExistingResourceResult = original.SkuEnumerationForExistingResourceResult
type SkuEnumerationForNewResourceResult = original.SkuEnumerationForNewResourceResult
type OperationsClient = original.OperationsClient
type ServersClient = original.ServersClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleConnectionModeValues() []ConnectionMode {
	return original.PossibleConnectionModeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServersClient(subscriptionID string) ServersClient {
	return original.NewServersClient(subscriptionID)
}
func NewServersClientWithBaseURI(baseURI string, subscriptionID string) ServersClient {
	return original.NewServersClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
