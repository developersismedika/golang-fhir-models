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

// RiskEvidenceSynthesis is documented here http://hl7.org/fhir/StructureDefinition/RiskEvidenceSynthesis
type RiskEvidenceSynthesis struct {
	Id                *string                            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                         `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                            `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                            `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string                            `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                            `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus                  `bson:"status" json:"status"`
	Date              *string                            `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                            `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail                    `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                            `bson:"description,omitempty" json:"description,omitempty"`
	Note              []Annotation                       `bson:"note,omitempty" json:"note,omitempty"`
	UseContext        []UsageContext                     `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                  `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Copyright         *string                            `bson:"copyright,omitempty" json:"copyright,omitempty"`
	ApprovalDate      *string                            `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate    *string                            `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period                            `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Topic             []CodeableConcept                  `bson:"topic,omitempty" json:"topic,omitempty"`
	Author            []ContactDetail                    `bson:"author,omitempty" json:"author,omitempty"`
	Editor            []ContactDetail                    `bson:"editor,omitempty" json:"editor,omitempty"`
	Reviewer          []ContactDetail                    `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	Endorser          []ContactDetail                    `bson:"endorser,omitempty" json:"endorser,omitempty"`
	RelatedArtifact   []RelatedArtifact                  `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	SynthesisType     *CodeableConcept                   `bson:"synthesisType,omitempty" json:"synthesisType,omitempty"`
	StudyType         *CodeableConcept                   `bson:"studyType,omitempty" json:"studyType,omitempty"`
	Population        Reference                          `bson:"population" json:"population"`
	Exposure          *Reference                         `bson:"exposure,omitempty" json:"exposure,omitempty"`
	Outcome           Reference                          `bson:"outcome" json:"outcome"`
	SampleSize        *RiskEvidenceSynthesisSampleSize   `bson:"sampleSize,omitempty" json:"sampleSize,omitempty"`
	RiskEstimate      *RiskEvidenceSynthesisRiskEstimate `bson:"riskEstimate,omitempty" json:"riskEstimate,omitempty"`
	Certainty         []RiskEvidenceSynthesisCertainty   `bson:"certainty,omitempty" json:"certainty,omitempty"`
}
type RiskEvidenceSynthesisSampleSize struct {
	Id                   *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description          *string     `bson:"description,omitempty" json:"description,omitempty"`
	NumberOfStudies      *int        `bson:"numberOfStudies,omitempty" json:"numberOfStudies,omitempty"`
	NumberOfParticipants *int        `bson:"numberOfParticipants,omitempty" json:"numberOfParticipants,omitempty"`
}
type RiskEvidenceSynthesisRiskEstimate struct {
	Id                *string                                              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string                                              `bson:"description,omitempty" json:"description,omitempty"`
	Type              *CodeableConcept                                     `bson:"type,omitempty" json:"type,omitempty"`
	Value             *string                                              `bson:"value,omitempty" json:"value,omitempty"`
	UnitOfMeasure     *CodeableConcept                                     `bson:"unitOfMeasure,omitempty" json:"unitOfMeasure,omitempty"`
	DenominatorCount  *int                                                 `bson:"denominatorCount,omitempty" json:"denominatorCount,omitempty"`
	NumeratorCount    *int                                                 `bson:"numeratorCount,omitempty" json:"numeratorCount,omitempty"`
	PrecisionEstimate []RiskEvidenceSynthesisRiskEstimatePrecisionEstimate `bson:"precisionEstimate,omitempty" json:"precisionEstimate,omitempty"`
}
type RiskEvidenceSynthesisRiskEstimatePrecisionEstimate struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Level             *string          `bson:"level,omitempty" json:"level,omitempty"`
	From              *string          `bson:"from,omitempty" json:"from,omitempty"`
	To                *string          `bson:"to,omitempty" json:"to,omitempty"`
}
type RiskEvidenceSynthesisCertainty struct {
	Id                    *string                                               `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []Extension                                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Rating                []CodeableConcept                                     `bson:"rating,omitempty" json:"rating,omitempty"`
	Note                  []Annotation                                          `bson:"note,omitempty" json:"note,omitempty"`
	CertaintySubcomponent []RiskEvidenceSynthesisCertaintyCertaintySubcomponent `bson:"certaintySubcomponent,omitempty" json:"certaintySubcomponent,omitempty"`
}
type RiskEvidenceSynthesisCertaintyCertaintySubcomponent struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Rating            []CodeableConcept `bson:"rating,omitempty" json:"rating,omitempty"`
	Note              []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
}
type OtherRiskEvidenceSynthesis RiskEvidenceSynthesis

// MarshalJSON marshals the given RiskEvidenceSynthesis as JSON into a byte slice
func (r RiskEvidenceSynthesis) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRiskEvidenceSynthesis
		ResourceType string `json:"resourceType"`
	}{
		OtherRiskEvidenceSynthesis: OtherRiskEvidenceSynthesis(r),
		ResourceType:               "RiskEvidenceSynthesis",
	})
}

// UnmarshalRiskEvidenceSynthesis unmarshals a RiskEvidenceSynthesis.
func UnmarshalRiskEvidenceSynthesis(b []byte) (RiskEvidenceSynthesis, error) {
	var riskEvidenceSynthesis RiskEvidenceSynthesis
	if err := json.Unmarshal(b, &riskEvidenceSynthesis); err != nil {
		return riskEvidenceSynthesis, err
	}
	return riskEvidenceSynthesis, nil
}
