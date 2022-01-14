/*
Alldebrid API

Welcome to the OpenAPI Alldebrid API v4 !<br /> You can use this API to access various Alldebrid services from custom applications or scripts.<br /> <br /> API is organized around REST,<br /> returns JSON-encoded responses and use standard HTTP response codes.<br /> <br /> All calls are to be made on the HTTPS endpoints.<br /> Some are public, others require to be authentificated with an apikey (see Authentication).<br /> <br /> You must also identify your apps or script with a meaningfull agent parameter.<br /> <br /> This API version is namespaced as v4, as such all endpoint start with /v4/,<br /> such like http://api.alldebrid.com/v4/ping?agent=apiShowcase.<br /> <br /> This API v4 should be the final version regarding general response format and errors (hopefully).<br />

API version: 4.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// ErrorCodes You can find all possible errors here, available in PHP array and JSON format if needed.<br /> Some errors will return an HTTP code 401 or 429, depending of the error.<br />
type ErrorCodes string

// List of ErrorCodes
const (
	GENERIC ErrorCodes = "GENERIC"
	_404 ErrorCodes = "404"
	AUTH_MISSING_AGENT ErrorCodes = "AUTH_MISSING_AGENT"
	AUTH_BAD_AGENT ErrorCodes = "AUTH_BAD_AGENT"
	AUTH_MISSING_APIKEY ErrorCodes = "AUTH_MISSING_APIKEY"
	AUTH_BAD_APIKEY ErrorCodes = "AUTH_BAD_APIKEY"
	AUTH_BLOCKED ErrorCodes = "AUTH_BLOCKED"
	AUTH_USER_BANNED ErrorCodes = "AUTH_USER_BANNED"
	LINK_IS_MISSING ErrorCodes = "LINK_IS_MISSING"
	LINK_HOST_NOT_SUPPORTED ErrorCodes = "LINK_HOST_NOT_SUPPORTED"
	LINK_DOWN ErrorCodes = "LINK_DOWN"
	LINK_PASS_PROTECTED ErrorCodes = "LINK_PASS_PROTECTED"
	LINK_HOST_UNAVAILABLE ErrorCodes = "LINK_HOST_UNAVAILABLE"
	LINK_TOO_MANY_DOWNLOADS ErrorCodes = "LINK_TOO_MANY_DOWNLOADS"
	LINK_HOST_FULL ErrorCodes = "LINK_HOST_FULL"
	LINK_HOST_LIMIT_REACHED ErrorCodes = "LINK_HOST_LIMIT_REACHED"
	LINK_ERROR ErrorCodes = "LINK_ERROR"
	REDIRECTOR_NOT_SUPPORTED ErrorCodes = "REDIRECTOR_NOT_SUPPORTED"
	REDIRECTOR_ERROR ErrorCodes = "REDIRECTOR_ERROR"
	STREAM_INVALID_GEN_ID ErrorCodes = "STREAM_INVALID_GEN_ID"
	STREAM_INVALID_STREAM_ID ErrorCodes = "STREAM_INVALID_STREAM_ID"
	DELAYED_INVALID_ID ErrorCodes = "DELAYED_INVALID_ID"
	FREE_TRIAL_LIMIT_REACHED ErrorCodes = "FREE_TRIAL_LIMIT_REACHED"
	MUST_BE_PREMIUM ErrorCodes = "MUST_BE_PREMIUM"
	MAGNET_INVALID_ID ErrorCodes = "MAGNET_INVALID_ID"
	MAGNET_INVALID_URI ErrorCodes = "MAGNET_INVALID_URI"
	MAGNET_INVALID_FILE ErrorCodes = "MAGNET_INVALID_FILE"
	MAGNET_FILE_UPLOAD_FAILED ErrorCodes = "MAGNET_FILE_UPLOAD_FAILED"
	MAGNET_NO_URI ErrorCodes = "MAGNET_NO_URI"
	MAGNET_PROCESSING ErrorCodes = "MAGNET_PROCESSING"
	MAGNET_TOO_MANY_ACTIVE ErrorCodes = "MAGNET_TOO_MANY_ACTIVE"
	MAGNET_MUST_BE_PREMIUM ErrorCodes = "MAGNET_MUST_BE_PREMIUM"
	MAGNET_NO_SERVER ErrorCodes = "MAGNET_NO_SERVER"
	MAGNET_TOO_LARGE ErrorCodes = "MAGNET_TOO_LARGE"
	PIN_ALREADY_AUTHED ErrorCodes = "PIN_ALREADY_AUTHED"
	PIN_EXPIRED ErrorCodes = "PIN_EXPIRED"
	PIN_INVALID ErrorCodes = "PIN_INVALID"
	USER_LINK_MISSING ErrorCodes = "USER_LINK_MISSING"
	USER_LINK_INVALID ErrorCodes = "USER_LINK_INVALID"
	NO_SERVER ErrorCodes = "NO_SERVER"
	MISSING_NOTIF_ENDPOINT ErrorCodes = "MISSING_NOTIF_ENDPOINT"
)

// All allowed values of ErrorCodes enum
var AllowedErrorCodesEnumValues = []ErrorCodes{
	"GENERIC",
	"404",
	"AUTH_MISSING_AGENT",
	"AUTH_BAD_AGENT",
	"AUTH_MISSING_APIKEY",
	"AUTH_BAD_APIKEY",
	"AUTH_BLOCKED",
	"AUTH_USER_BANNED",
	"LINK_IS_MISSING",
	"LINK_HOST_NOT_SUPPORTED",
	"LINK_DOWN",
	"LINK_PASS_PROTECTED",
	"LINK_HOST_UNAVAILABLE",
	"LINK_TOO_MANY_DOWNLOADS",
	"LINK_HOST_FULL",
	"LINK_HOST_LIMIT_REACHED",
	"LINK_ERROR",
	"REDIRECTOR_NOT_SUPPORTED",
	"REDIRECTOR_ERROR",
	"STREAM_INVALID_GEN_ID",
	"STREAM_INVALID_STREAM_ID",
	"DELAYED_INVALID_ID",
	"FREE_TRIAL_LIMIT_REACHED",
	"MUST_BE_PREMIUM",
	"MAGNET_INVALID_ID",
	"MAGNET_INVALID_URI",
	"MAGNET_INVALID_FILE",
	"MAGNET_FILE_UPLOAD_FAILED",
	"MAGNET_NO_URI",
	"MAGNET_PROCESSING",
	"MAGNET_TOO_MANY_ACTIVE",
	"MAGNET_MUST_BE_PREMIUM",
	"MAGNET_NO_SERVER",
	"MAGNET_TOO_LARGE",
	"PIN_ALREADY_AUTHED",
	"PIN_EXPIRED",
	"PIN_INVALID",
	"USER_LINK_MISSING",
	"USER_LINK_INVALID",
	"NO_SERVER",
	"MISSING_NOTIF_ENDPOINT",
}

func (v *ErrorCodes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorCodes(value)
	for _, existing := range AllowedErrorCodesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorCodes", value)
}

// NewErrorCodesFromValue returns a pointer to a valid ErrorCodes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewErrorCodesFromValue(v string) (*ErrorCodes, error) {
	ev := ErrorCodes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ErrorCodes: valid values are %v", v, AllowedErrorCodesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ErrorCodes) IsValid() bool {
	for _, existing := range AllowedErrorCodesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ErrorCodes value
func (v ErrorCodes) Ptr() *ErrorCodes {
	return &v
}

type NullableErrorCodes struct {
	value *ErrorCodes
	isSet bool
}

func (v NullableErrorCodes) Get() *ErrorCodes {
	return v.value
}

func (v *NullableErrorCodes) Set(val *ErrorCodes) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorCodes) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorCodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorCodes(val *ErrorCodes) *NullableErrorCodes {
	return &NullableErrorCodes{value: val, isSet: true}
}

func (v NullableErrorCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorCodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
