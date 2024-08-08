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

// BiologicallyDerivedProduct is documented here http://hl7.org/fhir/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProduct struct {
	BaseResource

	Id                *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                              `bson:"text,omitempty" json:"text,omitempty"`
	Contained         []BaseResource                          `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ProductCategory   *BiologicallyDerivedProductCategory     `bson:"productCategory,omitempty" json:"productCategory,omitempty"`
	ProductCode       *CodeableConcept                        `bson:"productCode,omitempty" json:"productCode,omitempty"`
	Status            *BiologicallyDerivedProductStatus       `bson:"status,omitempty" json:"status,omitempty"`
	Request           []Reference                             `bson:"request,omitempty" json:"request,omitempty"`
	Quantity          *int                                    `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Parent            []Reference                             `bson:"parent,omitempty" json:"parent,omitempty"`
	Collection        *BiologicallyDerivedProductCollection   `bson:"collection,omitempty" json:"collection,omitempty"`
	Processing        []BiologicallyDerivedProductProcessing  `bson:"processing,omitempty" json:"processing,omitempty"`
	Manipulation      *BiologicallyDerivedProductManipulation `bson:"manipulation,omitempty" json:"manipulation,omitempty"`
	Storage           []BiologicallyDerivedProductStorage     `bson:"storage,omitempty" json:"storage,omitempty"`
}
type BiologicallyDerivedProductCollection struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Collector         *Reference  `bson:"collector,omitempty" json:"collector,omitempty"`
	Source            *Reference  `bson:"source,omitempty" json:"source,omitempty"`
	CollectedDateTime *string     `bson:"collectedDateTime,omitempty" json:"collectedDateTime,omitempty"`
	CollectedPeriod   *Period     `bson:"collectedPeriod,omitempty" json:"collectedPeriod,omitempty"`
}
type BiologicallyDerivedProductProcessing struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string          `bson:"description,omitempty" json:"description,omitempty"`
	Procedure         *CodeableConcept `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Additive          *Reference       `bson:"additive,omitempty" json:"additive,omitempty"`
	TimeDateTime      *string          `bson:"timeDateTime,omitempty" json:"timeDateTime,omitempty"`
	TimePeriod        *Period          `bson:"timePeriod,omitempty" json:"timePeriod,omitempty"`
}
type BiologicallyDerivedProductManipulation struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	TimeDateTime      *string     `bson:"timeDateTime,omitempty" json:"timeDateTime,omitempty"`
	TimePeriod        *Period     `bson:"timePeriod,omitempty" json:"timePeriod,omitempty"`
}
type BiologicallyDerivedProductStorage struct {
	Id                *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string                                 `bson:"description,omitempty" json:"description,omitempty"`
	Temperature       *json.Number                            `bson:"temperature,omitempty" json:"temperature,omitempty"`
	Scale             *BiologicallyDerivedProductStorageScale `bson:"scale,omitempty" json:"scale,omitempty"`
	Duration          *Period                                 `bson:"duration,omitempty" json:"duration,omitempty"`
}
type OtherBiologicallyDerivedProduct BiologicallyDerivedProduct

// MarshalJSON marshals the given BiologicallyDerivedProduct as JSON into a byte slice
func (r BiologicallyDerivedProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBiologicallyDerivedProduct
		ResourceType string `json:"resourceType"`
	}{
		OtherBiologicallyDerivedProduct: OtherBiologicallyDerivedProduct(r),
		ResourceType:                    "BiologicallyDerivedProduct",
	})
}

// UnmarshalBiologicallyDerivedProduct unmarshals a BiologicallyDerivedProduct.
func UnmarshalBiologicallyDerivedProduct(b []byte) (BiologicallyDerivedProduct, error) {
	var biologicallyDerivedProduct BiologicallyDerivedProduct
	if err := json.Unmarshal(b, &biologicallyDerivedProduct); err != nil {
		return biologicallyDerivedProduct, err
	}
	return biologicallyDerivedProduct, nil
}
