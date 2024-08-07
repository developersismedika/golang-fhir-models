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

// Group is documented here http://hl7.org/fhir/StructureDefinition/Group
type Group struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Contained         json.RawMessage       `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool                 `bson:"active,omitempty" json:"active,omitempty"`
	Type              GroupType             `bson:"type" json:"type"`
	Actual            bool                  `bson:"actual" json:"actual"`
	Code              *CodeableConcept      `bson:"code,omitempty" json:"code,omitempty"`
	Name              *string               `bson:"name,omitempty" json:"name,omitempty"`
	Quantity          *int                  `bson:"quantity,omitempty" json:"quantity,omitempty"`
	ManagingEntity    *Reference            `bson:"managingEntity,omitempty" json:"managingEntity,omitempty"`
	Characteristic    []GroupCharacteristic `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	Member            []GroupMember         `bson:"member,omitempty" json:"member,omitempty"`
}
type GroupCharacteristic struct {
	Id                   *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code                 CodeableConcept `bson:"code" json:"code"`
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept"`
	ValueBoolean         bool            `bson:"valueBoolean" json:"valueBoolean"`
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity"`
	ValueRange           Range           `bson:"valueRange" json:"valueRange"`
	ValueReference       Reference       `bson:"valueReference" json:"valueReference"`
	Exclude              bool            `bson:"exclude" json:"exclude"`
	Period               *Period         `bson:"period,omitempty" json:"period,omitempty"`
}
type GroupMember struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Entity            Reference   `bson:"entity" json:"entity"`
	Period            *Period     `bson:"period,omitempty" json:"period,omitempty"`
	Inactive          *bool       `bson:"inactive,omitempty" json:"inactive,omitempty"`
}
type OtherGroup Group

// MarshalJSON marshals the given Group as JSON into a byte slice
func (r Group) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGroup
		ResourceType string `json:"resourceType"`
	}{
		OtherGroup:   OtherGroup(r),
		ResourceType: "Group",
	})
}

// UnmarshalGroup unmarshals a Group.
func UnmarshalGroup(b []byte) (Group, error) {
	var group Group
	if err := json.Unmarshal(b, &group); err != nil {
		return group, err
	}
	return group, nil
}
