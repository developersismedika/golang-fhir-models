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

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// DocumentRelationshipType is documented here http://hl7.org/fhir/ValueSet/document-relationship-type
type DocumentRelationshipType int

const (
	DocumentRelationshipTypeReplaces DocumentRelationshipType = iota
	DocumentRelationshipTypeTransforms
	DocumentRelationshipTypeSigns
	DocumentRelationshipTypeAppends
)

func (code DocumentRelationshipType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DocumentRelationshipType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "replaces":
		*code = DocumentRelationshipTypeReplaces
	case "transforms":
		*code = DocumentRelationshipTypeTransforms
	case "signs":
		*code = DocumentRelationshipTypeSigns
	case "appends":
		*code = DocumentRelationshipTypeAppends
	default:
		return fmt.Errorf("unknown DocumentRelationshipType code `%s`", s)
	}
	return nil
}
func (code DocumentRelationshipType) String() string {
	return code.Code()
}
func (code DocumentRelationshipType) Code() string {
	switch code {
	case DocumentRelationshipTypeReplaces:
		return "replaces"
	case DocumentRelationshipTypeTransforms:
		return "transforms"
	case DocumentRelationshipTypeSigns:
		return "signs"
	case DocumentRelationshipTypeAppends:
		return "appends"
	}
	return "<unknown>"
}
func (code DocumentRelationshipType) Display() string {
	switch code {
	case DocumentRelationshipTypeReplaces:
		return "Replaces"
	case DocumentRelationshipTypeTransforms:
		return "Transforms"
	case DocumentRelationshipTypeSigns:
		return "Signs"
	case DocumentRelationshipTypeAppends:
		return "Appends"
	}
	return "<unknown>"
}
func (code DocumentRelationshipType) Definition() string {
	switch code {
	case DocumentRelationshipTypeReplaces:
		return "This document logically replaces or supersedes the target document."
	case DocumentRelationshipTypeTransforms:
		return "This document was generated by transforming the target document (e.g. format or language conversion)."
	case DocumentRelationshipTypeSigns:
		return "This document is a signature of the target document."
	case DocumentRelationshipTypeAppends:
		return "This document adds additional information to the target document."
	}
	return "<unknown>"
}
