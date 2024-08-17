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

// MedicinalProductIngredient is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductIngredient
type MedicinalProductIngredient struct {
	Id                  *string                                        `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                                          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                                        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                                        `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative                                     `bson:"text,omitempty" json:"text,omitempty"`
	Contained           []BaseResource                                 `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension           []Extension                                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          *Identifier                                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Role                CodeableConcept                                `bson:"role" json:"role"`
	AllergenicIndicator *bool                                          `bson:"allergenicIndicator,omitempty" json:"allergenicIndicator,omitempty"`
	Manufacturer        []Reference                                    `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	SpecifiedSubstance  []MedicinalProductIngredientSpecifiedSubstance `bson:"specifiedSubstance,omitempty" json:"specifiedSubstance,omitempty"`
	Substance           *MedicinalProductIngredientSubstance           `bson:"substance,omitempty" json:"substance,omitempty"`
}
type MedicinalProductIngredientSpecifiedSubstance struct {
	Id                *string                                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                        `bson:"code" json:"code"`
	Group             CodeableConcept                                        `bson:"group" json:"group"`
	Confidentiality   *CodeableConcept                                       `bson:"confidentiality,omitempty" json:"confidentiality,omitempty"`
	Strength          []MedicinalProductIngredientSpecifiedSubstanceStrength `bson:"strength,omitempty" json:"strength,omitempty"`
}
type MedicinalProductIngredientSpecifiedSubstanceStrength struct {
	Id                    *string                                                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []Extension                                                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                                                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Presentation          Ratio                                                                   `bson:"presentation" json:"presentation"`
	PresentationLowLimit  *Ratio                                                                  `bson:"presentationLowLimit,omitempty" json:"presentationLowLimit,omitempty"`
	Concentration         *Ratio                                                                  `bson:"concentration,omitempty" json:"concentration,omitempty"`
	ConcentrationLowLimit *Ratio                                                                  `bson:"concentrationLowLimit,omitempty" json:"concentrationLowLimit,omitempty"`
	MeasurementPoint      *string                                                                 `bson:"measurementPoint,omitempty" json:"measurementPoint,omitempty"`
	Country               []CodeableConcept                                                       `bson:"country,omitempty" json:"country,omitempty"`
	ReferenceStrength     []MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength `bson:"referenceStrength,omitempty" json:"referenceStrength,omitempty"`
}
type MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Substance         *CodeableConcept  `bson:"substance,omitempty" json:"substance,omitempty"`
	Strength          Ratio             `bson:"strength" json:"strength"`
	StrengthLowLimit  *Ratio            `bson:"strengthLowLimit,omitempty" json:"strengthLowLimit,omitempty"`
	MeasurementPoint  *string           `bson:"measurementPoint,omitempty" json:"measurementPoint,omitempty"`
	Country           []CodeableConcept `bson:"country,omitempty" json:"country,omitempty"`
}
type MedicinalProductIngredientSubstance struct {
	Id                *string                                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                        `bson:"code" json:"code"`
	Strength          []MedicinalProductIngredientSpecifiedSubstanceStrength `bson:"strength,omitempty" json:"strength,omitempty"`
}
type OtherMedicinalProductIngredient MedicinalProductIngredient

// MarshalJSON marshals the given MedicinalProductIngredient as JSON into a byte slice
func (r MedicinalProductIngredient) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ResourceType string `json:"resourceType"`
		OtherMedicinalProductIngredient
	}{
		OtherMedicinalProductIngredient: OtherMedicinalProductIngredient(r),
		ResourceType:                    "MedicinalProductIngredient",
	})
}

// UnmarshalMedicinalProductIngredient unmarshals a MedicinalProductIngredient.
func UnmarshalMedicinalProductIngredient(b []byte) (MedicinalProductIngredient, error) {
	var medicinalProductIngredient MedicinalProductIngredient
	if err := json.Unmarshal(b, &medicinalProductIngredient); err != nil {
		return medicinalProductIngredient, err
	}
	return medicinalProductIngredient, nil
}
