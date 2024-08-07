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

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/developersismedika/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// EvidenceVariableType is documented here http://hl7.org/fhir/ValueSet/variable-type
type EvidenceVariableType int

const (
	EvidenceVariableTypeDichotomous EvidenceVariableType = iota
	EvidenceVariableTypeContinuous
	EvidenceVariableTypeDescriptive
)

func (code EvidenceVariableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *EvidenceVariableType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "dichotomous":
		*code = EvidenceVariableTypeDichotomous
	case "continuous":
		*code = EvidenceVariableTypeContinuous
	case "descriptive":
		*code = EvidenceVariableTypeDescriptive
	default:
		return fmt.Errorf("unknown EvidenceVariableType code `%s`", s)
	}
	return nil
}
func (code EvidenceVariableType) String() string {
	return code.Code()
}
func (code EvidenceVariableType) Code() string {
	switch code {
	case EvidenceVariableTypeDichotomous:
		return "dichotomous"
	case EvidenceVariableTypeContinuous:
		return "continuous"
	case EvidenceVariableTypeDescriptive:
		return "descriptive"
	}
	return "<unknown>"
}
func (code EvidenceVariableType) Display() string {
	switch code {
	case EvidenceVariableTypeDichotomous:
		return "Dichotomous"
	case EvidenceVariableTypeContinuous:
		return "Continuous"
	case EvidenceVariableTypeDescriptive:
		return "Descriptive"
	}
	return "<unknown>"
}
func (code EvidenceVariableType) Definition() string {
	switch code {
	case EvidenceVariableTypeDichotomous:
		return "The variable is dichotomous, such as present or absent."
	case EvidenceVariableTypeContinuous:
		return "The variable is a continuous result such as a quantity."
	case EvidenceVariableTypeDescriptive:
		return "The variable is described narratively rather than quantitatively."
	}
	return "<unknown>"
}
