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

// AppointmentResponse is documented here http://hl7.org/fhir/StructureDefinition/AppointmentResponse
type AppointmentResponse struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string             `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative          `bson:"text,omitempty" json:"text,omitempty"`
	Contained         []BaseResource      `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Appointment       Reference           `bson:"appointment" json:"appointment"`
	Start             *string             `bson:"start,omitempty" json:"start,omitempty"`
	End               *string             `bson:"end,omitempty" json:"end,omitempty"`
	ParticipantType   []CodeableConcept   `bson:"participantType,omitempty" json:"participantType,omitempty"`
	Actor             *Reference          `bson:"actor,omitempty" json:"actor,omitempty"`
	ParticipantStatus ParticipationStatus `bson:"participantStatus" json:"participantStatus"`
	Comment           *string             `bson:"comment,omitempty" json:"comment,omitempty"`
}
type OtherAppointmentResponse AppointmentResponse

// MarshalJSON marshals the given AppointmentResponse as JSON into a byte slice
func (r AppointmentResponse) MarshalJSON() ([]byte, error) {
	return jsonMarshal(struct {
		ResourceType string `json:"resourceType"`
		OtherAppointmentResponse
	}{
		OtherAppointmentResponse: OtherAppointmentResponse(r),
		ResourceType:             "AppointmentResponse",
	})
}

// UnmarshalAppointmentResponse unmarshals a AppointmentResponse.
func UnmarshalAppointmentResponse(b []byte) (AppointmentResponse, error) {
	var appointmentResponse AppointmentResponse
	if err := jsonUnmarshal(b, &appointmentResponse); err != nil {
		return appointmentResponse, err
	}
	return appointmentResponse, nil
}
