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

// MedicationKnowledge is documented here http://hl7.org/fhir/StructureDefinition/MedicationKnowledge
type MedicationKnowledge struct {
	Id                         *string                                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta                       *Meta                                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules              *string                                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                   *string                                         `bson:"language,omitempty" json:"language,omitempty"`
	Text                       *Narrative                                      `bson:"text,omitempty" json:"text,omitempty"`
	Contained                  []BaseResource                                  `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension                  []Extension                                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension          []Extension                                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code                       *CodeableConcept                                `bson:"code,omitempty" json:"code,omitempty"`
	Status                     *string                                         `bson:"status,omitempty" json:"status,omitempty"`
	Manufacturer               *Reference                                      `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	DoseForm                   *CodeableConcept                                `bson:"doseForm,omitempty" json:"doseForm,omitempty"`
	Amount                     *Quantity                                       `bson:"amount,omitempty" json:"amount,omitempty"`
	Synonym                    []string                                        `bson:"synonym,omitempty" json:"synonym,omitempty"`
	RelatedMedicationKnowledge []MedicationKnowledgeRelatedMedicationKnowledge `bson:"relatedMedicationKnowledge,omitempty" json:"relatedMedicationKnowledge,omitempty"`
	AssociatedMedication       []Reference                                     `bson:"associatedMedication,omitempty" json:"associatedMedication,omitempty"`
	ProductType                []CodeableConcept                               `bson:"productType,omitempty" json:"productType,omitempty"`
	Monograph                  []MedicationKnowledgeMonograph                  `bson:"monograph,omitempty" json:"monograph,omitempty"`
	Ingredient                 []MedicationKnowledgeIngredient                 `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	PreparationInstruction     *string                                         `bson:"preparationInstruction,omitempty" json:"preparationInstruction,omitempty"`
	IntendedRoute              []CodeableConcept                               `bson:"intendedRoute,omitempty" json:"intendedRoute,omitempty"`
	Cost                       []MedicationKnowledgeCost                       `bson:"cost,omitempty" json:"cost,omitempty"`
	MonitoringProgram          []MedicationKnowledgeMonitoringProgram          `bson:"monitoringProgram,omitempty" json:"monitoringProgram,omitempty"`
	AdministrationGuidelines   []MedicationKnowledgeAdministrationGuidelines   `bson:"administrationGuidelines,omitempty" json:"administrationGuidelines,omitempty"`
	MedicineClassification     []MedicationKnowledgeMedicineClassification     `bson:"medicineClassification,omitempty" json:"medicineClassification,omitempty"`
	Packaging                  *MedicationKnowledgePackaging                   `bson:"packaging,omitempty" json:"packaging,omitempty"`
	DrugCharacteristic         []MedicationKnowledgeDrugCharacteristic         `bson:"drugCharacteristic,omitempty" json:"drugCharacteristic,omitempty"`
	Contraindication           []Reference                                     `bson:"contraindication,omitempty" json:"contraindication,omitempty"`
	Regulatory                 []MedicationKnowledgeRegulatory                 `bson:"regulatory,omitempty" json:"regulatory,omitempty"`
	Kinetics                   []MedicationKnowledgeKinetics                   `bson:"kinetics,omitempty" json:"kinetics,omitempty"`
}
type MedicationKnowledgeRelatedMedicationKnowledge struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
	Reference         []Reference     `bson:"reference" json:"reference"`
}
type MedicationKnowledgeMonograph struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Source            *Reference       `bson:"source,omitempty" json:"source,omitempty"`
}
type MedicationKnowledgeIngredient struct {
	Id                  *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ItemCodeableConcept *CodeableConcept `bson:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `bson:"itemReference,omitempty" json:"itemReference,omitempty"`
	IsActive            *bool            `bson:"isActive,omitempty" json:"isActive,omitempty"`
	Strength            *Ratio           `bson:"strength,omitempty" json:"strength,omitempty"`
}
type MedicationKnowledgeCost struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
	Source            *string         `bson:"source,omitempty" json:"source,omitempty"`
	Cost              Money           `bson:"cost" json:"cost"`
}
type MedicationKnowledgeMonitoringProgram struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Name              *string          `bson:"name,omitempty" json:"name,omitempty"`
}
type MedicationKnowledgeAdministrationGuidelines struct {
	Id                        *string                                                             `bson:"id,omitempty" json:"id,omitempty"`
	Extension                 []Extension                                                         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []Extension                                                         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Dosage                    []MedicationKnowledgeAdministrationGuidelinesDosage                 `bson:"dosage,omitempty" json:"dosage,omitempty"`
	IndicationCodeableConcept *CodeableConcept                                                    `bson:"indicationCodeableConcept,omitempty" json:"indicationCodeableConcept,omitempty"`
	IndicationReference       *Reference                                                          `bson:"indicationReference,omitempty" json:"indicationReference,omitempty"`
	PatientCharacteristics    []MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics `bson:"patientCharacteristics,omitempty" json:"patientCharacteristics,omitempty"`
}
type MedicationKnowledgeAdministrationGuidelinesDosage struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
	Dosage            []Dosage        `bson:"dosage" json:"dosage"`
}
type MedicationKnowledgeAdministrationGuidelinesPatientCharacteristics struct {
	Id                            *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension                     []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension             []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	CharacteristicCodeableConcept *CodeableConcept `bson:"characteristicCodeableConcept,omitempty" json:"characteristicCodeableConcept,omitempty"`
	CharacteristicQuantity        *Quantity        `bson:"characteristicQuantity,omitempty" json:"characteristicQuantity,omitempty"`
	Value                         []string         `bson:"value,omitempty" json:"value,omitempty"`
}
type MedicationKnowledgeMedicineClassification struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `bson:"type" json:"type"`
	Classification    []CodeableConcept `bson:"classification,omitempty" json:"classification,omitempty"`
}
type MedicationKnowledgePackaging struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Quantity          *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
}
type MedicationKnowledgeDrugCharacteristic struct {
	Id                   *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type                 *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueString          *string          `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueBase64Binary    *string          `bson:"valueBase64Binary,omitempty" json:"valueBase64Binary,omitempty"`
}
type MedicationKnowledgeRegulatory struct {
	Id                  *string                                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension                                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	RegulatoryAuthority Reference                                   `bson:"regulatoryAuthority" json:"regulatoryAuthority"`
	Substitution        []MedicationKnowledgeRegulatorySubstitution `bson:"substitution,omitempty" json:"substitution,omitempty"`
	Schedule            []MedicationKnowledgeRegulatorySchedule     `bson:"schedule,omitempty" json:"schedule,omitempty"`
	MaxDispense         *MedicationKnowledgeRegulatoryMaxDispense   `bson:"maxDispense,omitempty" json:"maxDispense,omitempty"`
}
type MedicationKnowledgeRegulatorySubstitution struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
	Allowed           bool            `bson:"allowed" json:"allowed"`
}
type MedicationKnowledgeRegulatorySchedule struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Schedule          CodeableConcept `bson:"schedule" json:"schedule"`
}
type MedicationKnowledgeRegulatoryMaxDispense struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Quantity          Quantity    `bson:"quantity" json:"quantity"`
	Period            *Duration   `bson:"period,omitempty" json:"period,omitempty"`
}
type MedicationKnowledgeKinetics struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	AreaUnderCurve    []Quantity  `bson:"areaUnderCurve,omitempty" json:"areaUnderCurve,omitempty"`
	LethalDose50      []Quantity  `bson:"lethalDose50,omitempty" json:"lethalDose50,omitempty"`
	HalfLifePeriod    *Duration   `bson:"halfLifePeriod,omitempty" json:"halfLifePeriod,omitempty"`
}
type OtherMedicationKnowledge MedicationKnowledge

// MarshalJSON marshals the given MedicationKnowledge as JSON into a byte slice
func (r MedicationKnowledge) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationKnowledge
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationKnowledge: OtherMedicationKnowledge(r),
		ResourceType:             "MedicationKnowledge",
	})
}

// UnmarshalMedicationKnowledge unmarshals a MedicationKnowledge.
func UnmarshalMedicationKnowledge(b []byte) (MedicationKnowledge, error) {
	var medicationKnowledge MedicationKnowledge
	if err := json.Unmarshal(b, &medicationKnowledge); err != nil {
		return medicationKnowledge, err
	}
	return medicationKnowledge, nil
}
