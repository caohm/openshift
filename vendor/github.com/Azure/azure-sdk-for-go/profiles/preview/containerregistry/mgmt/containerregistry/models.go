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

package containerregistry

import original "github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2018-09-01/containerregistry"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type Architecture = original.Architecture

const (
	Amd64 Architecture = original.Amd64
	Arm   Architecture = original.Arm
	X86   Architecture = original.X86
)

type BaseImageDependencyType = original.BaseImageDependencyType

const (
	BuildTime BaseImageDependencyType = original.BuildTime
	RunTime   BaseImageDependencyType = original.RunTime
)

type BaseImageTriggerType = original.BaseImageTriggerType

const (
	All     BaseImageTriggerType = original.All
	Runtime BaseImageTriggerType = original.Runtime
)

type ImportMode = original.ImportMode

const (
	Force   ImportMode = original.Force
	NoForce ImportMode = original.NoForce
)

type OS = original.OS

const (
	Linux   OS = original.Linux
	Windows OS = original.Windows
)

type PasswordName = original.PasswordName

const (
	Password  PasswordName = original.Password
	Password2 PasswordName = original.Password2
)

type PolicyStatus = original.PolicyStatus

const (
	Disabled PolicyStatus = original.Disabled
	Enabled  PolicyStatus = original.Enabled
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Updating  ProvisioningState = original.Updating
)

type RegistryUsageUnit = original.RegistryUsageUnit

const (
	Bytes RegistryUsageUnit = original.Bytes
	Count RegistryUsageUnit = original.Count
)

type RunStatus = original.RunStatus

const (
	RunStatusCanceled  RunStatus = original.RunStatusCanceled
	RunStatusError     RunStatus = original.RunStatusError
	RunStatusFailed    RunStatus = original.RunStatusFailed
	RunStatusQueued    RunStatus = original.RunStatusQueued
	RunStatusRunning   RunStatus = original.RunStatusRunning
	RunStatusStarted   RunStatus = original.RunStatusStarted
	RunStatusSucceeded RunStatus = original.RunStatusSucceeded
	RunStatusTimeout   RunStatus = original.RunStatusTimeout
)

type RunType = original.RunType

const (
	AutoBuild  RunType = original.AutoBuild
	AutoRun    RunType = original.AutoRun
	QuickBuild RunType = original.QuickBuild
	QuickRun   RunType = original.QuickRun
)

type SkuName = original.SkuName

const (
	Basic    SkuName = original.Basic
	Classic  SkuName = original.Classic
	Premium  SkuName = original.Premium
	Standard SkuName = original.Standard
)

type SkuTier = original.SkuTier

const (
	SkuTierBasic    SkuTier = original.SkuTierBasic
	SkuTierClassic  SkuTier = original.SkuTierClassic
	SkuTierPremium  SkuTier = original.SkuTierPremium
	SkuTierStandard SkuTier = original.SkuTierStandard
)

type SourceControlType = original.SourceControlType

const (
	Github                  SourceControlType = original.Github
	VisualStudioTeamService SourceControlType = original.VisualStudioTeamService
)

type SourceTriggerEvent = original.SourceTriggerEvent

const (
	Commit      SourceTriggerEvent = original.Commit
	Pullrequest SourceTriggerEvent = original.Pullrequest
)

type TaskStatus = original.TaskStatus

const (
	TaskStatusDisabled TaskStatus = original.TaskStatusDisabled
	TaskStatusEnabled  TaskStatus = original.TaskStatusEnabled
)

type TokenType = original.TokenType

const (
	OAuth TokenType = original.OAuth
	PAT   TokenType = original.PAT
)

type TriggerStatus = original.TriggerStatus

const (
	TriggerStatusDisabled TriggerStatus = original.TriggerStatusDisabled
	TriggerStatusEnabled  TriggerStatus = original.TriggerStatusEnabled
)

type TrustPolicyType = original.TrustPolicyType

const (
	Notary TrustPolicyType = original.Notary
)

type Type = original.Type

const (
	TypeDockerBuildRequest    Type = original.TypeDockerBuildRequest
	TypeEncodedTaskRunRequest Type = original.TypeEncodedTaskRunRequest
	TypeFileTaskRunRequest    Type = original.TypeFileTaskRunRequest
	TypeRunRequest            Type = original.TypeRunRequest
	TypeTaskRunRequest        Type = original.TypeTaskRunRequest
)

type TypeBasicTaskStepProperties = original.TypeBasicTaskStepProperties

const (
	TypeDocker             TypeBasicTaskStepProperties = original.TypeDocker
	TypeEncodedTask        TypeBasicTaskStepProperties = original.TypeEncodedTask
	TypeFileTask           TypeBasicTaskStepProperties = original.TypeFileTask
	TypeTaskStepProperties TypeBasicTaskStepProperties = original.TypeTaskStepProperties
)

type TypeBasicTaskStepUpdateParameters = original.TypeBasicTaskStepUpdateParameters

const (
	TypeBasicTaskStepUpdateParametersTypeDocker                   TypeBasicTaskStepUpdateParameters = original.TypeBasicTaskStepUpdateParametersTypeDocker
	TypeBasicTaskStepUpdateParametersTypeEncodedTask              TypeBasicTaskStepUpdateParameters = original.TypeBasicTaskStepUpdateParametersTypeEncodedTask
	TypeBasicTaskStepUpdateParametersTypeFileTask                 TypeBasicTaskStepUpdateParameters = original.TypeBasicTaskStepUpdateParametersTypeFileTask
	TypeBasicTaskStepUpdateParametersTypeTaskStepUpdateParameters TypeBasicTaskStepUpdateParameters = original.TypeBasicTaskStepUpdateParametersTypeTaskStepUpdateParameters
)

type Variant = original.Variant

const (
	V6 Variant = original.V6
	V7 Variant = original.V7
	V8 Variant = original.V8
)

type WebhookAction = original.WebhookAction

const (
	Delete     WebhookAction = original.Delete
	Push       WebhookAction = original.Push
	Quarantine WebhookAction = original.Quarantine
)

type WebhookStatus = original.WebhookStatus

const (
	WebhookStatusDisabled WebhookStatus = original.WebhookStatusDisabled
	WebhookStatusEnabled  WebhookStatus = original.WebhookStatusEnabled
)

type Actor = original.Actor
type AgentProperties = original.AgentProperties
type Argument = original.Argument
type AuthInfo = original.AuthInfo
type AuthInfoUpdateParameters = original.AuthInfoUpdateParameters
type BaseImageDependency = original.BaseImageDependency
type BaseImageTrigger = original.BaseImageTrigger
type BaseImageTriggerUpdateParameters = original.BaseImageTriggerUpdateParameters
type CallbackConfig = original.CallbackConfig
type DockerBuildRequest = original.DockerBuildRequest
type DockerBuildStep = original.DockerBuildStep
type DockerBuildStepUpdateParameters = original.DockerBuildStepUpdateParameters
type EncodedTaskRunRequest = original.EncodedTaskRunRequest
type EncodedTaskStep = original.EncodedTaskStep
type EncodedTaskStepUpdateParameters = original.EncodedTaskStepUpdateParameters
type Event = original.Event
type EventContent = original.EventContent
type EventInfo = original.EventInfo
type EventListResult = original.EventListResult
type EventListResultIterator = original.EventListResultIterator
type EventListResultPage = original.EventListResultPage
type EventRequestMessage = original.EventRequestMessage
type EventResponseMessage = original.EventResponseMessage
type FileTaskRunRequest = original.FileTaskRunRequest
type FileTaskStep = original.FileTaskStep
type FileTaskStepUpdateParameters = original.FileTaskStepUpdateParameters
type ImageDescriptor = original.ImageDescriptor
type ImageUpdateTrigger = original.ImageUpdateTrigger
type ImportImageParameters = original.ImportImageParameters
type ImportSource = original.ImportSource
type ImportSourceCredentials = original.ImportSourceCredentials
type OperationDefinition = original.OperationDefinition
type OperationDisplayDefinition = original.OperationDisplayDefinition
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationMetricSpecificationDefinition = original.OperationMetricSpecificationDefinition
type OperationPropertiesDefinition = original.OperationPropertiesDefinition
type OperationServiceSpecificationDefinition = original.OperationServiceSpecificationDefinition
type PlatformProperties = original.PlatformProperties
type PlatformUpdateParameters = original.PlatformUpdateParameters
type ProxyResource = original.ProxyResource
type QuarantinePolicy = original.QuarantinePolicy
type RegenerateCredentialParameters = original.RegenerateCredentialParameters
type RegistriesCreateFuture = original.RegistriesCreateFuture
type RegistriesDeleteFuture = original.RegistriesDeleteFuture
type RegistriesImportImageFuture = original.RegistriesImportImageFuture
type RegistriesScheduleRunFuture = original.RegistriesScheduleRunFuture
type RegistriesUpdateFuture = original.RegistriesUpdateFuture
type RegistriesUpdatePoliciesFuture = original.RegistriesUpdatePoliciesFuture
type Registry = original.Registry
type RegistryIdentity = original.RegistryIdentity
type RegistryListCredentialsResult = original.RegistryListCredentialsResult
type RegistryListResult = original.RegistryListResult
type RegistryListResultIterator = original.RegistryListResultIterator
type RegistryListResultPage = original.RegistryListResultPage
type RegistryNameCheckRequest = original.RegistryNameCheckRequest
type RegistryNameStatus = original.RegistryNameStatus
type RegistryPassword = original.RegistryPassword
type RegistryPolicies = original.RegistryPolicies
type RegistryProperties = original.RegistryProperties
type RegistryPropertiesUpdateParameters = original.RegistryPropertiesUpdateParameters
type RegistryUpdateParameters = original.RegistryUpdateParameters
type RegistryUsage = original.RegistryUsage
type RegistryUsageListResult = original.RegistryUsageListResult
type Replication = original.Replication
type ReplicationListResult = original.ReplicationListResult
type ReplicationListResultIterator = original.ReplicationListResultIterator
type ReplicationListResultPage = original.ReplicationListResultPage
type ReplicationProperties = original.ReplicationProperties
type ReplicationsCreateFuture = original.ReplicationsCreateFuture
type ReplicationsDeleteFuture = original.ReplicationsDeleteFuture
type ReplicationsUpdateFuture = original.ReplicationsUpdateFuture
type ReplicationUpdateParameters = original.ReplicationUpdateParameters
type Request = original.Request
type Resource = original.Resource
type Run = original.Run
type RunFilter = original.RunFilter
type RunGetLogResult = original.RunGetLogResult
type RunListResult = original.RunListResult
type RunListResultIterator = original.RunListResultIterator
type RunListResultPage = original.RunListResultPage
type RunProperties = original.RunProperties
type BasicRunRequest = original.BasicRunRequest
type RunRequest = original.RunRequest
type RunsCancelFuture = original.RunsCancelFuture
type RunsUpdateFuture = original.RunsUpdateFuture
type RunUpdateParameters = original.RunUpdateParameters
type SetValue = original.SetValue
type Sku = original.Sku
type Source = original.Source
type SourceProperties = original.SourceProperties
type SourceTrigger = original.SourceTrigger
type SourceTriggerDescriptor = original.SourceTriggerDescriptor
type SourceTriggerUpdateParameters = original.SourceTriggerUpdateParameters
type SourceUpdateParameters = original.SourceUpdateParameters
type SourceUploadDefinition = original.SourceUploadDefinition
type Status = original.Status
type StorageAccountProperties = original.StorageAccountProperties
type Target = original.Target
type Task = original.Task
type TaskListResult = original.TaskListResult
type TaskListResultIterator = original.TaskListResultIterator
type TaskListResultPage = original.TaskListResultPage
type TaskProperties = original.TaskProperties
type TaskPropertiesUpdateParameters = original.TaskPropertiesUpdateParameters
type TaskRunRequest = original.TaskRunRequest
type TasksCreateFuture = original.TasksCreateFuture
type TasksDeleteFuture = original.TasksDeleteFuture
type BasicTaskStepProperties = original.BasicTaskStepProperties
type TaskStepProperties = original.TaskStepProperties
type BasicTaskStepUpdateParameters = original.BasicTaskStepUpdateParameters
type TaskStepUpdateParameters = original.TaskStepUpdateParameters
type TasksUpdateFuture = original.TasksUpdateFuture
type TaskUpdateParameters = original.TaskUpdateParameters
type TriggerProperties = original.TriggerProperties
type TriggerUpdateParameters = original.TriggerUpdateParameters
type TrustPolicy = original.TrustPolicy
type Webhook = original.Webhook
type WebhookCreateParameters = original.WebhookCreateParameters
type WebhookListResult = original.WebhookListResult
type WebhookListResultIterator = original.WebhookListResultIterator
type WebhookListResultPage = original.WebhookListResultPage
type WebhookProperties = original.WebhookProperties
type WebhookPropertiesCreateParameters = original.WebhookPropertiesCreateParameters
type WebhookPropertiesUpdateParameters = original.WebhookPropertiesUpdateParameters
type WebhooksCreateFuture = original.WebhooksCreateFuture
type WebhooksDeleteFuture = original.WebhooksDeleteFuture
type WebhooksUpdateFuture = original.WebhooksUpdateFuture
type WebhookUpdateParameters = original.WebhookUpdateParameters
type OperationsClient = original.OperationsClient
type RegistriesClient = original.RegistriesClient
type ReplicationsClient = original.ReplicationsClient
type RunsClient = original.RunsClient
type TasksClient = original.TasksClient
type WebhooksClient = original.WebhooksClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleArchitectureValues() []Architecture {
	return original.PossibleArchitectureValues()
}
func PossibleBaseImageDependencyTypeValues() []BaseImageDependencyType {
	return original.PossibleBaseImageDependencyTypeValues()
}
func PossibleBaseImageTriggerTypeValues() []BaseImageTriggerType {
	return original.PossibleBaseImageTriggerTypeValues()
}
func PossibleImportModeValues() []ImportMode {
	return original.PossibleImportModeValues()
}
func PossibleOSValues() []OS {
	return original.PossibleOSValues()
}
func PossiblePasswordNameValues() []PasswordName {
	return original.PossiblePasswordNameValues()
}
func PossiblePolicyStatusValues() []PolicyStatus {
	return original.PossiblePolicyStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleRegistryUsageUnitValues() []RegistryUsageUnit {
	return original.PossibleRegistryUsageUnitValues()
}
func PossibleRunStatusValues() []RunStatus {
	return original.PossibleRunStatusValues()
}
func PossibleRunTypeValues() []RunType {
	return original.PossibleRunTypeValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleSourceControlTypeValues() []SourceControlType {
	return original.PossibleSourceControlTypeValues()
}
func PossibleSourceTriggerEventValues() []SourceTriggerEvent {
	return original.PossibleSourceTriggerEventValues()
}
func PossibleTaskStatusValues() []TaskStatus {
	return original.PossibleTaskStatusValues()
}
func PossibleTokenTypeValues() []TokenType {
	return original.PossibleTokenTypeValues()
}
func PossibleTriggerStatusValues() []TriggerStatus {
	return original.PossibleTriggerStatusValues()
}
func PossibleTrustPolicyTypeValues() []TrustPolicyType {
	return original.PossibleTrustPolicyTypeValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleTypeBasicTaskStepPropertiesValues() []TypeBasicTaskStepProperties {
	return original.PossibleTypeBasicTaskStepPropertiesValues()
}
func PossibleTypeBasicTaskStepUpdateParametersValues() []TypeBasicTaskStepUpdateParameters {
	return original.PossibleTypeBasicTaskStepUpdateParametersValues()
}
func PossibleVariantValues() []Variant {
	return original.PossibleVariantValues()
}
func PossibleWebhookActionValues() []WebhookAction {
	return original.PossibleWebhookActionValues()
}
func PossibleWebhookStatusValues() []WebhookStatus {
	return original.PossibleWebhookStatusValues()
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRegistriesClient(subscriptionID string) RegistriesClient {
	return original.NewRegistriesClient(subscriptionID)
}
func NewRegistriesClientWithBaseURI(baseURI string, subscriptionID string) RegistriesClient {
	return original.NewRegistriesClientWithBaseURI(baseURI, subscriptionID)
}
func NewReplicationsClient(subscriptionID string) ReplicationsClient {
	return original.NewReplicationsClient(subscriptionID)
}
func NewReplicationsClientWithBaseURI(baseURI string, subscriptionID string) ReplicationsClient {
	return original.NewReplicationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRunsClient(subscriptionID string) RunsClient {
	return original.NewRunsClient(subscriptionID)
}
func NewRunsClientWithBaseURI(baseURI string, subscriptionID string) RunsClient {
	return original.NewRunsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTasksClient(subscriptionID string) TasksClient {
	return original.NewTasksClient(subscriptionID)
}
func NewTasksClientWithBaseURI(baseURI string, subscriptionID string) TasksClient {
	return original.NewTasksClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func NewWebhooksClient(subscriptionID string) WebhooksClient {
	return original.NewWebhooksClient(subscriptionID)
}
func NewWebhooksClientWithBaseURI(baseURI string, subscriptionID string) WebhooksClient {
	return original.NewWebhooksClientWithBaseURI(baseURI, subscriptionID)
}
