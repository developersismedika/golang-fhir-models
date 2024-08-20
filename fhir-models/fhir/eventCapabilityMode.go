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

// EventCapabilityMode is documented here http://hl7.org/fhir/ValueSet/event-capability-mode
type EventCapabilityMode int

const (
	EventCapabilityModeSender EventCapabilityMode = iota
	EventCapabilityModeReceiver
)

func (code EventCapabilityMode) MarshalJSON() ([]byte, error) {
	return jsonMarshal(code.Code())
}
func (code *EventCapabilityMode) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "sender":
		*code = EventCapabilityModeSender
	case "receiver":
		*code = EventCapabilityModeReceiver
	default:
		return fmt.Errorf("unknown EventCapabilityMode code `%s`", s)
	}
	return nil
}
func (code EventCapabilityMode) String() string {
	return code.Code()
}
func (code EventCapabilityMode) Code() string {
	switch code {
	case EventCapabilityModeSender:
		return "sender"
	case EventCapabilityModeReceiver:
		return "receiver"
	}
	return "<unknown>"
}
func (code EventCapabilityMode) Display() string {
	switch code {
	case EventCapabilityModeSender:
		return "Sender"
	case EventCapabilityModeReceiver:
		return "Receiver"
	}
	return "<unknown>"
}
func (code EventCapabilityMode) Definition() string {
	switch code {
	case EventCapabilityModeSender:
		return "The application sends requests and receives responses."
	case EventCapabilityModeReceiver:
		return "The application receives requests and sends responses."
	}
	return "<unknown>"
}
