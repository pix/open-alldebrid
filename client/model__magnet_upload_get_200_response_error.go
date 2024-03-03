/*
Alldebrid API

Welcome to the OpenAPI Alldebrid API v4 !<br /> You can use this API to access various Alldebrid services from custom applications or scripts.<br /> <br /> API is organized around REST,<br /> returns JSON-encoded responses and use standard HTTP response codes.<br /> <br /> All calls are to be made on the HTTPS endpoints.<br /> Some are public, others require to be authentificated with an apikey (see Authentication).<br /> <br /> You must also identify your apps or script with a meaningfull agent parameter.<br /> <br /> This API version is namespaced as v4, as such all endpoint start with /v4/,<br /> such like http://api.alldebrid.com/v4/ping?agent=apiShowcase.<br /> <br /> This API v4 should be the final version regarding general response format and errors (hopefully).<br />

API version: 4.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the MagnetUploadGet200ResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MagnetUploadGet200ResponseError{}

// MagnetUploadGet200ResponseError struct for MagnetUploadGet200ResponseError
type MagnetUploadGet200ResponseError struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewMagnetUploadGet200ResponseError instantiates a new MagnetUploadGet200ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMagnetUploadGet200ResponseError() *MagnetUploadGet200ResponseError {
	this := MagnetUploadGet200ResponseError{}
	return &this
}

// NewMagnetUploadGet200ResponseErrorWithDefaults instantiates a new MagnetUploadGet200ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMagnetUploadGet200ResponseErrorWithDefaults() *MagnetUploadGet200ResponseError {
	this := MagnetUploadGet200ResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *MagnetUploadGet200ResponseError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MagnetUploadGet200ResponseError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *MagnetUploadGet200ResponseError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *MagnetUploadGet200ResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *MagnetUploadGet200ResponseError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MagnetUploadGet200ResponseError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *MagnetUploadGet200ResponseError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *MagnetUploadGet200ResponseError) SetMessage(v string) {
	o.Message = &v
}

func (o MagnetUploadGet200ResponseError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MagnetUploadGet200ResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableMagnetUploadGet200ResponseError struct {
	value *MagnetUploadGet200ResponseError
	isSet bool
}

func (v NullableMagnetUploadGet200ResponseError) Get() *MagnetUploadGet200ResponseError {
	return v.value
}

func (v *NullableMagnetUploadGet200ResponseError) Set(val *MagnetUploadGet200ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableMagnetUploadGet200ResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableMagnetUploadGet200ResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMagnetUploadGet200ResponseError(val *MagnetUploadGet200ResponseError) *NullableMagnetUploadGet200ResponseError {
	return &NullableMagnetUploadGet200ResponseError{value: val, isSet: true}
}

func (v NullableMagnetUploadGet200ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMagnetUploadGet200ResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


