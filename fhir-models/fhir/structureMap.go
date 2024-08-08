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

// StructureMap is documented here http://hl7.org/fhir/StructureDefinition/StructureMap
type StructureMap struct {
	BaseResource

	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	Contained         []BaseResource          `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                  `bson:"url" json:"url"`
	Identifier        []Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                 `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                  `bson:"name" json:"name"`
	Title             *string                 `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus       `bson:"status" json:"status"`
	Experimental      *bool                   `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                 `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                 `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail         `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                 `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext          `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept       `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                 `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string                 `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Structure         []StructureMapStructure `bson:"structure,omitempty" json:"structure,omitempty"`
	Import            []string                `bson:"import,omitempty" json:"import,omitempty"`
	Group             []StructureMapGroup     `bson:"group" json:"group"`
}
type StructureMapStructure struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                `bson:"url" json:"url"`
	Mode              StructureMapModelMode `bson:"mode" json:"mode"`
	Alias             *string               `bson:"alias,omitempty" json:"alias,omitempty"`
	Documentation     *string               `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type StructureMapGroup struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                    `bson:"name" json:"name"`
	Extends           *string                   `bson:"extends,omitempty" json:"extends,omitempty"`
	TypeMode          StructureMapGroupTypeMode `bson:"typeMode" json:"typeMode"`
	Documentation     *string                   `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Input             []StructureMapGroupInput  `bson:"input" json:"input"`
	Rule              []StructureMapGroupRule   `bson:"rule" json:"rule"`
}
type StructureMapGroupInput struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                `bson:"name" json:"name"`
	Type              *string               `bson:"type,omitempty" json:"type,omitempty"`
	Mode              StructureMapInputMode `bson:"mode" json:"mode"`
	Documentation     *string               `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type StructureMapGroupRule struct {
	Id                *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                           `bson:"name" json:"name"`
	Source            []StructureMapGroupRuleSource    `bson:"source" json:"source"`
	Target            []StructureMapGroupRuleTarget    `bson:"target,omitempty" json:"target,omitempty"`
	Rule              []StructureMapGroupRule          `bson:"rule,omitempty" json:"rule,omitempty"`
	Dependent         []StructureMapGroupRuleDependent `bson:"dependent,omitempty" json:"dependent,omitempty"`
	Documentation     *string                          `bson:"documentation,omitempty" json:"documentation,omitempty"`
}
type StructureMapGroupRuleSource struct {
	Id                              *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension                       []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension               []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Context                         string                      `bson:"context" json:"context"`
	Min                             *int                        `bson:"min,omitempty" json:"min,omitempty"`
	Max                             *string                     `bson:"max,omitempty" json:"max,omitempty"`
	Type                            *string                     `bson:"type,omitempty" json:"type,omitempty"`
	DefaultValueBase64Binary        *string                     `bson:"defaultValueBase64Binary,omitempty" json:"defaultValueBase64Binary,omitempty"`
	DefaultValueBoolean             *bool                       `bson:"defaultValueBoolean,omitempty" json:"defaultValueBoolean,omitempty"`
	DefaultValueCanonical           *string                     `bson:"defaultValueCanonical,omitempty" json:"defaultValueCanonical,omitempty"`
	DefaultValueCode                *string                     `bson:"defaultValueCode,omitempty" json:"defaultValueCode,omitempty"`
	DefaultValueDate                *string                     `bson:"defaultValueDate,omitempty" json:"defaultValueDate,omitempty"`
	DefaultValueDateTime            *string                     `bson:"defaultValueDateTime,omitempty" json:"defaultValueDateTime,omitempty"`
	DefaultValueDecimal             *json.Number                `bson:"defaultValueDecimal,omitempty" json:"defaultValueDecimal,omitempty"`
	DefaultValueId                  *string                     `bson:"defaultValueId,omitempty" json:"defaultValueId,omitempty"`
	DefaultValueInstant             *string                     `bson:"defaultValueInstant,omitempty" json:"defaultValueInstant,omitempty"`
	DefaultValueInteger             *int                        `bson:"defaultValueInteger,omitempty" json:"defaultValueInteger,omitempty"`
	DefaultValueMarkdown            *string                     `bson:"defaultValueMarkdown,omitempty" json:"defaultValueMarkdown,omitempty"`
	DefaultValueOid                 *string                     `bson:"defaultValueOid,omitempty" json:"defaultValueOid,omitempty"`
	DefaultValuePositiveInt         *int                        `bson:"defaultValuePositiveInt,omitempty" json:"defaultValuePositiveInt,omitempty"`
	DefaultValueString              *string                     `bson:"defaultValueString,omitempty" json:"defaultValueString,omitempty"`
	DefaultValueTime                *string                     `bson:"defaultValueTime,omitempty" json:"defaultValueTime,omitempty"`
	DefaultValueUnsignedInt         *int                        `bson:"defaultValueUnsignedInt,omitempty" json:"defaultValueUnsignedInt,omitempty"`
	DefaultValueUri                 *string                     `bson:"defaultValueUri,omitempty" json:"defaultValueUri,omitempty"`
	DefaultValueUrl                 *string                     `bson:"defaultValueUrl,omitempty" json:"defaultValueUrl,omitempty"`
	DefaultValueUuid                *string                     `bson:"defaultValueUuid,omitempty" json:"defaultValueUuid,omitempty"`
	DefaultValueAddress             *Address                    `bson:"defaultValueAddress,omitempty" json:"defaultValueAddress,omitempty"`
	DefaultValueAge                 *Age                        `bson:"defaultValueAge,omitempty" json:"defaultValueAge,omitempty"`
	DefaultValueAnnotation          *Annotation                 `bson:"defaultValueAnnotation,omitempty" json:"defaultValueAnnotation,omitempty"`
	DefaultValueAttachment          *Attachment                 `bson:"defaultValueAttachment,omitempty" json:"defaultValueAttachment,omitempty"`
	DefaultValueCodeableConcept     *CodeableConcept            `bson:"defaultValueCodeableConcept,omitempty" json:"defaultValueCodeableConcept,omitempty"`
	DefaultValueCoding              *Coding                     `bson:"defaultValueCoding,omitempty" json:"defaultValueCoding,omitempty"`
	DefaultValueContactPoint        *ContactPoint               `bson:"defaultValueContactPoint,omitempty" json:"defaultValueContactPoint,omitempty"`
	DefaultValueCount               *Count                      `bson:"defaultValueCount,omitempty" json:"defaultValueCount,omitempty"`
	DefaultValueDistance            *Distance                   `bson:"defaultValueDistance,omitempty" json:"defaultValueDistance,omitempty"`
	DefaultValueDuration            *Duration                   `bson:"defaultValueDuration,omitempty" json:"defaultValueDuration,omitempty"`
	DefaultValueHumanName           *HumanName                  `bson:"defaultValueHumanName,omitempty" json:"defaultValueHumanName,omitempty"`
	DefaultValueIdentifier          *Identifier                 `bson:"defaultValueIdentifier,omitempty" json:"defaultValueIdentifier,omitempty"`
	DefaultValueMoney               *Money                      `bson:"defaultValueMoney,omitempty" json:"defaultValueMoney,omitempty"`
	DefaultValuePeriod              *Period                     `bson:"defaultValuePeriod,omitempty" json:"defaultValuePeriod,omitempty"`
	DefaultValueQuantity            *Quantity                   `bson:"defaultValueQuantity,omitempty" json:"defaultValueQuantity,omitempty"`
	DefaultValueRange               *Range                      `bson:"defaultValueRange,omitempty" json:"defaultValueRange,omitempty"`
	DefaultValueRatio               *Ratio                      `bson:"defaultValueRatio,omitempty" json:"defaultValueRatio,omitempty"`
	DefaultValueReference           *Reference                  `bson:"defaultValueReference,omitempty" json:"defaultValueReference,omitempty"`
	DefaultValueSampledData         *SampledData                `bson:"defaultValueSampledData,omitempty" json:"defaultValueSampledData,omitempty"`
	DefaultValueSignature           *Signature                  `bson:"defaultValueSignature,omitempty" json:"defaultValueSignature,omitempty"`
	DefaultValueTiming              *Timing                     `bson:"defaultValueTiming,omitempty" json:"defaultValueTiming,omitempty"`
	DefaultValueContactDetail       *ContactDetail              `bson:"defaultValueContactDetail,omitempty" json:"defaultValueContactDetail,omitempty"`
	DefaultValueContributor         *Contributor                `bson:"defaultValueContributor,omitempty" json:"defaultValueContributor,omitempty"`
	DefaultValueDataRequirement     *DataRequirement            `bson:"defaultValueDataRequirement,omitempty" json:"defaultValueDataRequirement,omitempty"`
	DefaultValueExpression          *Expression                 `bson:"defaultValueExpression,omitempty" json:"defaultValueExpression,omitempty"`
	DefaultValueParameterDefinition *ParameterDefinition        `bson:"defaultValueParameterDefinition,omitempty" json:"defaultValueParameterDefinition,omitempty"`
	DefaultValueRelatedArtifact     *RelatedArtifact            `bson:"defaultValueRelatedArtifact,omitempty" json:"defaultValueRelatedArtifact,omitempty"`
	DefaultValueTriggerDefinition   *TriggerDefinition          `bson:"defaultValueTriggerDefinition,omitempty" json:"defaultValueTriggerDefinition,omitempty"`
	DefaultValueUsageContext        *UsageContext               `bson:"defaultValueUsageContext,omitempty" json:"defaultValueUsageContext,omitempty"`
	DefaultValueDosage              *Dosage                     `bson:"defaultValueDosage,omitempty" json:"defaultValueDosage,omitempty"`
	DefaultValueMeta                *Meta                       `bson:"defaultValueMeta,omitempty" json:"defaultValueMeta,omitempty"`
	Element                         *string                     `bson:"element,omitempty" json:"element,omitempty"`
	ListMode                        *StructureMapSourceListMode `bson:"listMode,omitempty" json:"listMode,omitempty"`
	Variable                        *string                     `bson:"variable,omitempty" json:"variable,omitempty"`
	Condition                       *string                     `bson:"condition,omitempty" json:"condition,omitempty"`
	Check                           *string                     `bson:"check,omitempty" json:"check,omitempty"`
	LogMessage                      *string                     `bson:"logMessage,omitempty" json:"logMessage,omitempty"`
}
type StructureMapGroupRuleTarget struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Context           *string                                `bson:"context,omitempty" json:"context,omitempty"`
	ContextType       *StructureMapContextType               `bson:"contextType,omitempty" json:"contextType,omitempty"`
	Element           *string                                `bson:"element,omitempty" json:"element,omitempty"`
	Variable          *string                                `bson:"variable,omitempty" json:"variable,omitempty"`
	ListMode          []StructureMapTargetListMode           `bson:"listMode,omitempty" json:"listMode,omitempty"`
	ListRuleId        *string                                `bson:"listRuleId,omitempty" json:"listRuleId,omitempty"`
	Transform         *StructureMapTransform                 `bson:"transform,omitempty" json:"transform,omitempty"`
	Parameter         []StructureMapGroupRuleTargetParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
}
type StructureMapGroupRuleTargetParameter struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ValueId           string      `bson:"valueId" json:"valueId"`
	ValueString       string      `bson:"valueString" json:"valueString"`
	ValueBoolean      bool        `bson:"valueBoolean" json:"valueBoolean"`
	ValueInteger      int         `bson:"valueInteger" json:"valueInteger"`
	ValueDecimal      json.Number `bson:"valueDecimal" json:"valueDecimal"`
}
type StructureMapGroupRuleDependent struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name" json:"name"`
	Variable          []string    `bson:"variable" json:"variable"`
}
type OtherStructureMap StructureMap

// MarshalJSON marshals the given StructureMap as JSON into a byte slice
func (r StructureMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherStructureMap
		ResourceType string `json:"resourceType"`
	}{
		OtherStructureMap: OtherStructureMap(r),
		ResourceType:      "StructureMap",
	})
}

// UnmarshalStructureMap unmarshals a StructureMap.
func UnmarshalStructureMap(b []byte) (StructureMap, error) {
	var structureMap StructureMap
	if err := json.Unmarshal(b, &structureMap); err != nil {
		return structureMap, err
	}
	return structureMap, nil
}
