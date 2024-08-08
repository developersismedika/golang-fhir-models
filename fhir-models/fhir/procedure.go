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

// Procedure is documented here http://hl7.org/fhir/StructureDefinition/Procedure
type Procedure struct {
	Id                    *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Contained             []BaseResource         `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension             []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical []string               `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string               `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	BasedOn               []Reference            `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf                []Reference            `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                EventStatus            `bson:"status" json:"status"`
	StatusReason          *CodeableConcept       `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Category              *CodeableConcept       `bson:"category,omitempty" json:"category,omitempty"`
	Code                  *CodeableConcept       `bson:"code,omitempty" json:"code,omitempty"`
	Subject               Reference              `bson:"subject" json:"subject"`
	Encounter             *Reference             `bson:"encounter,omitempty" json:"encounter,omitempty"`
	PerformedDateTime     *string                `bson:"performedDateTime,omitempty" json:"performedDateTime,omitempty"`
	PerformedPeriod       *Period                `bson:"performedPeriod,omitempty" json:"performedPeriod,omitempty"`
	PerformedString       *string                `bson:"performedString,omitempty" json:"performedString,omitempty"`
	PerformedAge          *Age                   `bson:"performedAge,omitempty" json:"performedAge,omitempty"`
	PerformedRange        *Range                 `bson:"performedRange,omitempty" json:"performedRange,omitempty"`
	Recorder              *Reference             `bson:"recorder,omitempty" json:"recorder,omitempty"`
	Asserter              *Reference             `bson:"asserter,omitempty" json:"asserter,omitempty"`
	Performer             []ProcedurePerformer   `bson:"performer,omitempty" json:"performer,omitempty"`
	Location              *Reference             `bson:"location,omitempty" json:"location,omitempty"`
	ReasonCode            []CodeableConcept      `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference       []Reference            `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	BodySite              []CodeableConcept      `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Outcome               *CodeableConcept       `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Report                []Reference            `bson:"report,omitempty" json:"report,omitempty"`
	Complication          []CodeableConcept      `bson:"complication,omitempty" json:"complication,omitempty"`
	ComplicationDetail    []Reference            `bson:"complicationDetail,omitempty" json:"complicationDetail,omitempty"`
	FollowUp              []CodeableConcept      `bson:"followUp,omitempty" json:"followUp,omitempty"`
	Note                  []Annotation           `bson:"note,omitempty" json:"note,omitempty"`
	FocalDevice           []ProcedureFocalDevice `bson:"focalDevice,omitempty" json:"focalDevice,omitempty"`
	UsedReference         []Reference            `bson:"usedReference,omitempty" json:"usedReference,omitempty"`
	UsedCode              []CodeableConcept      `bson:"usedCode,omitempty" json:"usedCode,omitempty"`
}
type ProcedurePerformer struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	Actor             Reference        `bson:"actor" json:"actor"`
	OnBehalfOf        *Reference       `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}
type ProcedureFocalDevice struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Action            *CodeableConcept `bson:"action,omitempty" json:"action,omitempty"`
	Manipulated       Reference        `bson:"manipulated" json:"manipulated"`
}
type OtherProcedure Procedure

// MarshalJSON marshals the given Procedure as JSON into a byte slice
func (r Procedure) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProcedure
		ResourceType string `json:"resourceType"`
	}{
		OtherProcedure: OtherProcedure(r),
		ResourceType:   "Procedure",
	})
}

// UnmarshalProcedure unmarshals a Procedure.
func UnmarshalProcedure(b []byte) (Procedure, error) {
	var procedure Procedure
	if err := json.Unmarshal(b, &procedure); err != nil {
		return procedure, err
	}
	return procedure, nil
}
