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
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/developersismedika/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// DetectedIssueSeverity is documented here http://hl7.org/fhir/ValueSet/detectedissue-severity
type DetectedIssueSeverity int

const (
	DetectedIssueSeverityHigh DetectedIssueSeverity = iota
	DetectedIssueSeverityModerate
	DetectedIssueSeverityLow
)

func (code DetectedIssueSeverity) MarshalJSON() ([]byte, error) {
	return jsonMarshal(code.Code())
}
func (code *DetectedIssueSeverity) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "high":
		*code = DetectedIssueSeverityHigh
	case "moderate":
		*code = DetectedIssueSeverityModerate
	case "low":
		*code = DetectedIssueSeverityLow
	default:
		return fmt.Errorf("unknown DetectedIssueSeverity code `%s`", s)
	}
	return nil
}
func (code DetectedIssueSeverity) String() string {
	return code.Code()
}
func (code DetectedIssueSeverity) Code() string {
	switch code {
	case DetectedIssueSeverityHigh:
		return "high"
	case DetectedIssueSeverityModerate:
		return "moderate"
	case DetectedIssueSeverityLow:
		return "low"
	}
	return "<unknown>"
}
func (code DetectedIssueSeverity) Display() string {
	switch code {
	case DetectedIssueSeverityHigh:
		return "High"
	case DetectedIssueSeverityModerate:
		return "Moderate"
	case DetectedIssueSeverityLow:
		return "Low"
	}
	return "<unknown>"
}
func (code DetectedIssueSeverity) Definition() string {
	switch code {
	case DetectedIssueSeverityHigh:
		return "Indicates the issue may be life-threatening or has the potential to cause permanent injury."
	case DetectedIssueSeverityModerate:
		return "Indicates the issue may result in noticeable adverse consequences but is unlikely to be life-threatening or cause permanent injury."
	case DetectedIssueSeverityLow:
		return "Indicates the issue may result in some adverse consequences but is unlikely to substantially affect the situation of the subject."
	}
	return "<unknown>"
}
