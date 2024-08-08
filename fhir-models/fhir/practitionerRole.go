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

// PractitionerRole is documented here http://hl7.org/fhir/StructureDefinition/PractitionerRole
type PractitionerRole struct {
	Id                     *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	Contained              []BaseResource                  `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension              []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier             []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active                 *bool                           `bson:"active,omitempty" json:"active,omitempty"`
	Period                 *Period                         `bson:"period,omitempty" json:"period,omitempty"`
	Practitioner           *Reference                      `bson:"practitioner,omitempty" json:"practitioner,omitempty"`
	Organization           *Reference                      `bson:"organization,omitempty" json:"organization,omitempty"`
	Code                   []CodeableConcept               `bson:"code,omitempty" json:"code,omitempty"`
	Specialty              []CodeableConcept               `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Location               []Reference                     `bson:"location,omitempty" json:"location,omitempty"`
	HealthcareService      []Reference                     `bson:"healthcareService,omitempty" json:"healthcareService,omitempty"`
	Telecom                []ContactPoint                  `bson:"telecom,omitempty" json:"telecom,omitempty"`
	AvailableTime          []PractitionerRoleAvailableTime `bson:"availableTime,omitempty" json:"availableTime,omitempty"`
	NotAvailable           []PractitionerRoleNotAvailable  `bson:"notAvailable,omitempty" json:"notAvailable,omitempty"`
	AvailabilityExceptions *string                         `bson:"availabilityExceptions,omitempty" json:"availabilityExceptions,omitempty"`
	Endpoint               []Reference                     `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}
type PractitionerRoleAvailableTime struct {
	Id                 *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DaysOfWeek         []DaysOfWeek `bson:"daysOfWeek,omitempty" json:"daysOfWeek,omitempty"`
	AllDay             *bool        `bson:"allDay,omitempty" json:"allDay,omitempty"`
	AvailableStartTime *string      `bson:"availableStartTime,omitempty" json:"availableStartTime,omitempty"`
	AvailableEndTime   *string      `bson:"availableEndTime,omitempty" json:"availableEndTime,omitempty"`
}
type PractitionerRoleNotAvailable struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       string      `bson:"description" json:"description"`
	During            *Period     `bson:"during,omitempty" json:"during,omitempty"`
}
type OtherPractitionerRole PractitionerRole

// MarshalJSON marshals the given PractitionerRole as JSON into a byte slice
func (r PractitionerRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPractitionerRole
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitionerRole: OtherPractitionerRole(r),
		ResourceType:          "PractitionerRole",
	})
}

// UnmarshalPractitionerRole unmarshals a PractitionerRole.
func UnmarshalPractitionerRole(b []byte) (PractitionerRole, error) {
	var practitionerRole PractitionerRole
	if err := json.Unmarshal(b, &practitionerRole); err != nil {
		return practitionerRole, err
	}
	return practitionerRole, nil
}
