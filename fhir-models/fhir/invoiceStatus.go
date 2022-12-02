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

// InvoiceStatus is documented here http://hl7.org/fhir/ValueSet/invoice-status
type InvoiceStatus int

const (
	InvoiceStatusDraft InvoiceStatus = iota
	InvoiceStatusIssued
	InvoiceStatusBalanced
	InvoiceStatusCancelled
	InvoiceStatusEnteredInError
)

func (code InvoiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *InvoiceStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "draft":
		*code = InvoiceStatusDraft
	case "issued":
		*code = InvoiceStatusIssued
	case "balanced":
		*code = InvoiceStatusBalanced
	case "cancelled":
		*code = InvoiceStatusCancelled
	case "entered-in-error":
		*code = InvoiceStatusEnteredInError
	default:
		return fmt.Errorf("unknown InvoiceStatus code `%s`", s)
	}
	return nil
}
func (code InvoiceStatus) String() string {
	return code.Code()
}
func (code InvoiceStatus) Code() string {
	switch code {
	case InvoiceStatusDraft:
		return "draft"
	case InvoiceStatusIssued:
		return "issued"
	case InvoiceStatusBalanced:
		return "balanced"
	case InvoiceStatusCancelled:
		return "cancelled"
	case InvoiceStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code InvoiceStatus) Display() string {
	switch code {
	case InvoiceStatusDraft:
		return "draft"
	case InvoiceStatusIssued:
		return "issued"
	case InvoiceStatusBalanced:
		return "balanced"
	case InvoiceStatusCancelled:
		return "cancelled"
	case InvoiceStatusEnteredInError:
		return "entered in error"
	}
	return "<unknown>"
}
func (code InvoiceStatus) Definition() string {
	switch code {
	case InvoiceStatusDraft:
		return "the invoice has been prepared but not yet finalized."
	case InvoiceStatusIssued:
		return "the invoice has been finalized and sent to the recipient."
	case InvoiceStatusBalanced:
		return "the invoice has been balaced / completely paid."
	case InvoiceStatusCancelled:
		return "the invoice was cancelled."
	case InvoiceStatusEnteredInError:
		return "the invoice was determined as entered in error before it was issued."
	}
	return "<unknown>"
}
