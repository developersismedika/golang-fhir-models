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

// ClaimProcessingCodes is documented here http://hl7.org/fhir/ValueSet/remittance-outcome
type ClaimProcessingCodes int

const (
	ClaimProcessingCodesQueued ClaimProcessingCodes = iota
	ClaimProcessingCodesComplete
	ClaimProcessingCodesError
	ClaimProcessingCodesPartial
)

func (code ClaimProcessingCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ClaimProcessingCodes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "queued":
		*code = ClaimProcessingCodesQueued
	case "complete":
		*code = ClaimProcessingCodesComplete
	case "error":
		*code = ClaimProcessingCodesError
	case "partial":
		*code = ClaimProcessingCodesPartial
	default:
		return fmt.Errorf("unknown ClaimProcessingCodes code `%s`", s)
	}
	return nil
}
func (code ClaimProcessingCodes) String() string {
	return code.Code()
}
func (code ClaimProcessingCodes) Code() string {
	switch code {
	case ClaimProcessingCodesQueued:
		return "queued"
	case ClaimProcessingCodesComplete:
		return "complete"
	case ClaimProcessingCodesError:
		return "error"
	case ClaimProcessingCodesPartial:
		return "partial"
	}
	return "<unknown>"
}
func (code ClaimProcessingCodes) Display() string {
	switch code {
	case ClaimProcessingCodesQueued:
		return "Queued"
	case ClaimProcessingCodesComplete:
		return "Processing Complete"
	case ClaimProcessingCodesError:
		return "Error"
	case ClaimProcessingCodesPartial:
		return "Partial Processing"
	}
	return "<unknown>"
}
func (code ClaimProcessingCodes) Definition() string {
	switch code {
	case ClaimProcessingCodesQueued:
		return "The Claim/Pre-authorization/Pre-determination has been received but processing has not begun."
	case ClaimProcessingCodesComplete:
		return "The processing has completed without errors"
	case ClaimProcessingCodesError:
		return "One or more errors have been detected in the Claim"
	case ClaimProcessingCodesPartial:
		return "No errors have been detected in the Claim and some of the adjudication has been performed."
	}
	return "<unknown>"
}
