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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/developersismedika/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// DomainResource is documented here http://hl7.org/fhir/StructureDefinition/DomainResource
type DomainResource struct {
	Id                *string        `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string        `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative     `bson:"text,omitempty" json:"text,omitempty"`
	Contained         []BaseResource `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}
type OtherDomainResource DomainResource

// MarshalJSON marshals the given DomainResource as JSON into a byte slice
func (r DomainResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ResourceType string `json:"resourceType"`
		OtherDomainResource
	}{
		OtherDomainResource: OtherDomainResource(r),
		ResourceType:        "DomainResource",
	})
}

// UnmarshalDomainResource unmarshals a DomainResource.
func UnmarshalDomainResource(b []byte) (DomainResource, error) {
	var domainResource DomainResource
	if err := json.Unmarshal(b, &domainResource); err != nil {
		return domainResource, err
	}
	return domainResource, nil
}
