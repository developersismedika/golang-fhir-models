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

// CapabilityStatementKind is documented here http://hl7.org/fhir/ValueSet/capability-statement-kind
type CapabilityStatementKind int

const (
	CapabilityStatementKindInstance CapabilityStatementKind = iota
	CapabilityStatementKindCapability
	CapabilityStatementKindRequirements
)

func (code CapabilityStatementKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CapabilityStatementKind) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "instance":
		*code = CapabilityStatementKindInstance
	case "capability":
		*code = CapabilityStatementKindCapability
	case "requirements":
		*code = CapabilityStatementKindRequirements
	default:
		return fmt.Errorf("unknown CapabilityStatementKind code `%s`", s)
	}
	return nil
}
func (code CapabilityStatementKind) String() string {
	return code.Code()
}
func (code CapabilityStatementKind) Code() string {
	switch code {
	case CapabilityStatementKindInstance:
		return "instance"
	case CapabilityStatementKindCapability:
		return "capability"
	case CapabilityStatementKindRequirements:
		return "requirements"
	}
	return "<unknown>"
}
func (code CapabilityStatementKind) Display() string {
	switch code {
	case CapabilityStatementKindInstance:
		return "Instance"
	case CapabilityStatementKindCapability:
		return "Capability"
	case CapabilityStatementKindRequirements:
		return "Requirements"
	}
	return "<unknown>"
}
func (code CapabilityStatementKind) Definition() string {
	switch code {
	case CapabilityStatementKindInstance:
		return "The CapabilityStatement instance represents the present capabilities of a specific system instance.  This is the kind returned by /metadata for a FHIR server end-point."
	case CapabilityStatementKindCapability:
		return "The CapabilityStatement instance represents the capabilities of a system or piece of software, independent of a particular installation."
	case CapabilityStatementKindRequirements:
		return "The CapabilityStatement instance represents a set of requirements for other systems to meet; e.g. as part of an implementation guide or 'request for proposal'."
	}
	return "<unknown>"
}
