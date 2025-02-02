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

package policy

import original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-12-01/policy"

type AssignmentsClient = original.AssignmentsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type DefinitionsClient = original.DefinitionsClient
type Mode = original.Mode

const (
	All          Mode = original.All
	Indexed      Mode = original.Indexed
	NotSpecified Mode = original.NotSpecified
)

type Type = original.Type

const (
	TypeBuiltIn      Type = original.TypeBuiltIn
	TypeCustom       Type = original.TypeCustom
	TypeNotSpecified Type = original.TypeNotSpecified
)

type Assignment = original.Assignment
type AssignmentListResult = original.AssignmentListResult
type AssignmentListResultIterator = original.AssignmentListResultIterator
type AssignmentListResultPage = original.AssignmentListResultPage
type AssignmentProperties = original.AssignmentProperties
type Definition = original.Definition
type DefinitionListResult = original.DefinitionListResult
type DefinitionListResultIterator = original.DefinitionListResultIterator
type DefinitionListResultPage = original.DefinitionListResultPage
type DefinitionProperties = original.DefinitionProperties

func NewAssignmentsClient(subscriptionID string) AssignmentsClient {
	return original.NewAssignmentsClient(subscriptionID)
}
func NewAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) AssignmentsClient {
	return original.NewAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewDefinitionsClient(subscriptionID string) DefinitionsClient {
	return original.NewDefinitionsClient(subscriptionID)
}
func NewDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) DefinitionsClient {
	return original.NewDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleModeValues() []Mode {
	return original.PossibleModeValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
