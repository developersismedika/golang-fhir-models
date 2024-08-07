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

// AppointmentStatus is documented here http://hl7.org/fhir/ValueSet/appointmentstatus
type AppointmentStatus int

const (
	AppointmentStatusProposed AppointmentStatus = iota
	AppointmentStatusPending
	AppointmentStatusBooked
	AppointmentStatusArrived
	AppointmentStatusFulfilled
	AppointmentStatusCancelled
	AppointmentStatusNoshow
	AppointmentStatusEnteredInError
	AppointmentStatusCheckedIn
	AppointmentStatusWaitlist
)

func (code AppointmentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AppointmentStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "proposed":
		*code = AppointmentStatusProposed
	case "pending":
		*code = AppointmentStatusPending
	case "booked":
		*code = AppointmentStatusBooked
	case "arrived":
		*code = AppointmentStatusArrived
	case "fulfilled":
		*code = AppointmentStatusFulfilled
	case "cancelled":
		*code = AppointmentStatusCancelled
	case "noshow":
		*code = AppointmentStatusNoshow
	case "entered-in-error":
		*code = AppointmentStatusEnteredInError
	case "checked-in":
		*code = AppointmentStatusCheckedIn
	case "waitlist":
		*code = AppointmentStatusWaitlist
	default:
		return fmt.Errorf("unknown AppointmentStatus code `%s`", s)
	}
	return nil
}
func (code AppointmentStatus) String() string {
	return code.Code()
}
func (code AppointmentStatus) Code() string {
	switch code {
	case AppointmentStatusProposed:
		return "proposed"
	case AppointmentStatusPending:
		return "pending"
	case AppointmentStatusBooked:
		return "booked"
	case AppointmentStatusArrived:
		return "arrived"
	case AppointmentStatusFulfilled:
		return "fulfilled"
	case AppointmentStatusCancelled:
		return "cancelled"
	case AppointmentStatusNoshow:
		return "noshow"
	case AppointmentStatusEnteredInError:
		return "entered-in-error"
	case AppointmentStatusCheckedIn:
		return "checked-in"
	case AppointmentStatusWaitlist:
		return "waitlist"
	}
	return "<unknown>"
}
func (code AppointmentStatus) Display() string {
	switch code {
	case AppointmentStatusProposed:
		return "Proposed"
	case AppointmentStatusPending:
		return "Pending"
	case AppointmentStatusBooked:
		return "Booked"
	case AppointmentStatusArrived:
		return "Arrived"
	case AppointmentStatusFulfilled:
		return "Fulfilled"
	case AppointmentStatusCancelled:
		return "Cancelled"
	case AppointmentStatusNoshow:
		return "No Show"
	case AppointmentStatusEnteredInError:
		return "Entered in error"
	case AppointmentStatusCheckedIn:
		return "Checked In"
	case AppointmentStatusWaitlist:
		return "Waitlisted"
	}
	return "<unknown>"
}
func (code AppointmentStatus) Definition() string {
	switch code {
	case AppointmentStatusProposed:
		return "None of the participant(s) have finalized their acceptance of the appointment request, and the start/end time might not be set yet."
	case AppointmentStatusPending:
		return "Some or all of the participant(s) have not finalized their acceptance of the appointment request."
	case AppointmentStatusBooked:
		return "All participant(s) have been considered and the appointment is confirmed to go ahead at the date/times specified."
	case AppointmentStatusArrived:
		return "The patient/patients has/have arrived and is/are waiting to be seen."
	case AppointmentStatusFulfilled:
		return "The planning stages of the appointment are now complete, the encounter resource will exist and will track further status changes. Note that an encounter may exist before the appointment status is fulfilled for many reasons."
	case AppointmentStatusCancelled:
		return "The appointment has been cancelled."
	case AppointmentStatusNoshow:
		return "Some or all of the participant(s) have not/did not appear for the appointment (usually the patient)."
	case AppointmentStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	case AppointmentStatusCheckedIn:
		return "When checked in, all pre-encounter administrative work is complete, and the encounter may begin. (where multiple patients are involved, they are all present)."
	case AppointmentStatusWaitlist:
		return "The appointment has been placed on a waitlist, to be scheduled/confirmed in the future when a slot/service is available.\nA specific time might or might not be pre-allocated."
	}
	return "<unknown>"
}
