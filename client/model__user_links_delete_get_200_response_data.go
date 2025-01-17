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

// checks if the UserLinksDeleteGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserLinksDeleteGet200ResponseData{}

// UserLinksDeleteGet200ResponseData struct for UserLinksDeleteGet200ResponseData
type UserLinksDeleteGet200ResponseData struct {
	Message *string `json:"message,omitempty"`
}

// NewUserLinksDeleteGet200ResponseData instantiates a new UserLinksDeleteGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLinksDeleteGet200ResponseData() *UserLinksDeleteGet200ResponseData {
	this := UserLinksDeleteGet200ResponseData{}
	return &this
}

// NewUserLinksDeleteGet200ResponseDataWithDefaults instantiates a new UserLinksDeleteGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLinksDeleteGet200ResponseDataWithDefaults() *UserLinksDeleteGet200ResponseData {
	this := UserLinksDeleteGet200ResponseData{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UserLinksDeleteGet200ResponseData) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinksDeleteGet200ResponseData) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UserLinksDeleteGet200ResponseData) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UserLinksDeleteGet200ResponseData) SetMessage(v string) {
	o.Message = &v
}

func (o UserLinksDeleteGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserLinksDeleteGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableUserLinksDeleteGet200ResponseData struct {
	value *UserLinksDeleteGet200ResponseData
	isSet bool
}

func (v NullableUserLinksDeleteGet200ResponseData) Get() *UserLinksDeleteGet200ResponseData {
	return v.value
}

func (v *NullableUserLinksDeleteGet200ResponseData) Set(val *UserLinksDeleteGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLinksDeleteGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLinksDeleteGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLinksDeleteGet200ResponseData(val *UserLinksDeleteGet200ResponseData) *NullableUserLinksDeleteGet200ResponseData {
	return &NullableUserLinksDeleteGet200ResponseData{value: val, isSet: true}
}

func (v NullableUserLinksDeleteGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLinksDeleteGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


