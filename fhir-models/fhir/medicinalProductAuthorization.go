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

// MedicinalProductAuthorization is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductAuthorization
type MedicinalProductAuthorization struct {
	Id                          *string                                                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta                        *Meta                                                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules               *string                                                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                    *string                                                    `bson:"language,omitempty" json:"language,omitempty"`
	Text                        *Narrative                                                 `bson:"text,omitempty" json:"text,omitempty"`
	Contained                   []BaseResource                                             `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension                   []Extension                                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension           []Extension                                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier                  []Identifier                                               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject                     *Reference                                                 `bson:"subject,omitempty" json:"subject,omitempty"`
	Country                     []CodeableConcept                                          `bson:"country,omitempty" json:"country,omitempty"`
	Jurisdiction                []CodeableConcept                                          `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Status                      *CodeableConcept                                           `bson:"status,omitempty" json:"status,omitempty"`
	StatusDate                  *string                                                    `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	RestoreDate                 *string                                                    `bson:"restoreDate,omitempty" json:"restoreDate,omitempty"`
	ValidityPeriod              *Period                                                    `bson:"validityPeriod,omitempty" json:"validityPeriod,omitempty"`
	DataExclusivityPeriod       *Period                                                    `bson:"dataExclusivityPeriod,omitempty" json:"dataExclusivityPeriod,omitempty"`
	DateOfFirstAuthorization    *string                                                    `bson:"dateOfFirstAuthorization,omitempty" json:"dateOfFirstAuthorization,omitempty"`
	InternationalBirthDate      *string                                                    `bson:"internationalBirthDate,omitempty" json:"internationalBirthDate,omitempty"`
	LegalBasis                  *CodeableConcept                                           `bson:"legalBasis,omitempty" json:"legalBasis,omitempty"`
	JurisdictionalAuthorization []MedicinalProductAuthorizationJurisdictionalAuthorization `bson:"jurisdictionalAuthorization,omitempty" json:"jurisdictionalAuthorization,omitempty"`
	Holder                      *Reference                                                 `bson:"holder,omitempty" json:"holder,omitempty"`
	Regulator                   *Reference                                                 `bson:"regulator,omitempty" json:"regulator,omitempty"`
	Procedure                   *MedicinalProductAuthorizationProcedure                    `bson:"procedure,omitempty" json:"procedure,omitempty"`
}
type MedicinalProductAuthorizationJurisdictionalAuthorization struct {
	Id                  *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Country             *CodeableConcept  `bson:"country,omitempty" json:"country,omitempty"`
	Jurisdiction        []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	LegalStatusOfSupply *CodeableConcept  `bson:"legalStatusOfSupply,omitempty" json:"legalStatusOfSupply,omitempty"`
	ValidityPeriod      *Period           `bson:"validityPeriod,omitempty" json:"validityPeriod,omitempty"`
}
type MedicinalProductAuthorizationProcedure struct {
	Id                *string                                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier                              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type              CodeableConcept                          `bson:"type" json:"type"`
	DatePeriod        *Period                                  `bson:"datePeriod,omitempty" json:"datePeriod,omitempty"`
	DateDateTime      *string                                  `bson:"dateDateTime,omitempty" json:"dateDateTime,omitempty"`
	Application       []MedicinalProductAuthorizationProcedure `bson:"application,omitempty" json:"application,omitempty"`
}
type OtherMedicinalProductAuthorization MedicinalProductAuthorization

// MarshalJSON marshals the given MedicinalProductAuthorization as JSON into a byte slice
func (r MedicinalProductAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductAuthorization
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductAuthorization: OtherMedicinalProductAuthorization(r),
		ResourceType:                       "MedicinalProductAuthorization",
	})
}

// UnmarshalMedicinalProductAuthorization unmarshals a MedicinalProductAuthorization.
func UnmarshalMedicinalProductAuthorization(b []byte) (MedicinalProductAuthorization, error) {
	var medicinalProductAuthorization MedicinalProductAuthorization
	if err := json.Unmarshal(b, &medicinalProductAuthorization); err != nil {
		return medicinalProductAuthorization, err
	}
	return medicinalProductAuthorization, nil
}
