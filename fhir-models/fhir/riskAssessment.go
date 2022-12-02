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

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// RiskAssessment is documented here http://hl7.org/fhir/StructureDefinition/RiskAssessment
type RiskAssessment struct {
	Id                 *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                    `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative                 `bson:"text,omitempty" json:"text,omitempty"`
	Extension          []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn            *Reference                 `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Parent             *Reference                 `bson:"parent,omitempty" json:"parent,omitempty"`
	Status             ObservationStatus          `bson:"status" json:"status"`
	Method             *CodeableConcept           `bson:"method,omitempty" json:"method,omitempty"`
	Code               *CodeableConcept           `bson:"code,omitempty" json:"code,omitempty"`
	Subject            Reference                  `bson:"subject" json:"subject"`
	Encounter          *Reference                 `bson:"encounter,omitempty" json:"encounter,omitempty"`
	OccurrenceDateTime *string                    `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                    `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	Condition          *Reference                 `bson:"condition,omitempty" json:"condition,omitempty"`
	Performer          *Reference                 `bson:"performer,omitempty" json:"performer,omitempty"`
	ReasonCode         []CodeableConcept          `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference    []Reference                `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Basis              []Reference                `bson:"basis,omitempty" json:"basis,omitempty"`
	Prediction         []RiskAssessmentPrediction `bson:"prediction,omitempty" json:"prediction,omitempty"`
	Mitigation         *string                    `bson:"mitigation,omitempty" json:"mitigation,omitempty"`
	Note               []Annotation               `bson:"note,omitempty" json:"note,omitempty"`
}
type RiskAssessmentPrediction struct {
	Id                 *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Outcome            *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty"`
	ProbabilityDecimal *string          `bson:"probabilityDecimal,omitempty" json:"probabilityDecimal,omitempty"`
	ProbabilityRange   *Range           `bson:"probabilityRange,omitempty" json:"probabilityRange,omitempty"`
	QualitativeRisk    *CodeableConcept `bson:"qualitativeRisk,omitempty" json:"qualitativeRisk,omitempty"`
	RelativeRisk       *string          `bson:"relativeRisk,omitempty" json:"relativeRisk,omitempty"`
	WhenPeriod         *Period          `bson:"whenPeriod,omitempty" json:"whenPeriod,omitempty"`
	WhenRange          *Range           `bson:"whenRange,omitempty" json:"whenRange,omitempty"`
	Rationale          *string          `bson:"rationale,omitempty" json:"rationale,omitempty"`
}
type OtherRiskAssessment RiskAssessment

// MarshalJSON marshals the given RiskAssessment as JSON into a byte slice
func (r RiskAssessment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRiskAssessment
		ResourceType string `json:"resourceType"`
	}{
		OtherRiskAssessment: OtherRiskAssessment(r),
		ResourceType:        "RiskAssessment",
	})
}

// UnmarshalRiskAssessment unmarshals a RiskAssessment.
func UnmarshalRiskAssessment(b []byte) (RiskAssessment, error) {
	var riskAssessment RiskAssessment
	if err := json.Unmarshal(b, &riskAssessment); err != nil {
		return riskAssessment, err
	}
	return riskAssessment, nil
}
