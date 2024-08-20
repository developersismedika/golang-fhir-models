// Copyright 2019 - 2022 The Samply Community
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

package fhir

// THIS FILE IS GENERATED BY https://github.com/developersismedika/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// OperationOutcome is documented here http://hl7.org/fhir/StructureDefinition/OperationOutcome
type OperationOutcome struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	Contained         []BaseResource          `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Issue             []OperationOutcomeIssue `bson:"issue" json:"issue"`
}
type OperationOutcomeIssue struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Severity          IssueSeverity    `bson:"severity" json:"severity"`
	Code              IssueType        `bson:"code" json:"code"`
	Details           *CodeableConcept `bson:"details,omitempty" json:"details,omitempty"`
	Diagnostics       *string          `bson:"diagnostics,omitempty" json:"diagnostics,omitempty"`
	Location          []string         `bson:"location,omitempty" json:"location,omitempty"`
	Expression        []string         `bson:"expression,omitempty" json:"expression,omitempty"`
}
type OtherOperationOutcome OperationOutcome

// MarshalJSON marshals the given OperationOutcome as JSON into a byte slice
func (r OperationOutcome) MarshalJSON() ([]byte, error) {
	return jsonMarshal(struct {
		ResourceType string `json:"resourceType"`
		OtherOperationOutcome
	}{
		OtherOperationOutcome: OtherOperationOutcome(r),
		ResourceType:          "OperationOutcome",
	})
}

// UnmarshalOperationOutcome unmarshals a OperationOutcome.
func UnmarshalOperationOutcome(b []byte) (OperationOutcome, error) {
	var operationOutcome OperationOutcome
	if err := jsonUnmarshal(b, &operationOutcome); err != nil {
		return operationOutcome, err
	}
	return operationOutcome, nil
}
