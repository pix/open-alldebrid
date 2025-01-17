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

// checks if the UserLinksGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserLinksGet200ResponseData{}

// UserLinksGet200ResponseData struct for UserLinksGet200ResponseData
type UserLinksGet200ResponseData struct {
	Links []UserLinksGet200ResponseDataLinksInner `json:"links,omitempty"`
}

// NewUserLinksGet200ResponseData instantiates a new UserLinksGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLinksGet200ResponseData() *UserLinksGet200ResponseData {
	this := UserLinksGet200ResponseData{}
	return &this
}

// NewUserLinksGet200ResponseDataWithDefaults instantiates a new UserLinksGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLinksGet200ResponseDataWithDefaults() *UserLinksGet200ResponseData {
	this := UserLinksGet200ResponseData{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserLinksGet200ResponseData) GetLinks() []UserLinksGet200ResponseDataLinksInner {
	if o == nil || IsNil(o.Links) {
		var ret []UserLinksGet200ResponseDataLinksInner
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinksGet200ResponseData) GetLinksOk() ([]UserLinksGet200ResponseDataLinksInner, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserLinksGet200ResponseData) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []UserLinksGet200ResponseDataLinksInner and assigns it to the Links field.
func (o *UserLinksGet200ResponseData) SetLinks(v []UserLinksGet200ResponseDataLinksInner) {
	o.Links = v
}

func (o UserLinksGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserLinksGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableUserLinksGet200ResponseData struct {
	value *UserLinksGet200ResponseData
	isSet bool
}

func (v NullableUserLinksGet200ResponseData) Get() *UserLinksGet200ResponseData {
	return v.value
}

func (v *NullableUserLinksGet200ResponseData) Set(val *UserLinksGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLinksGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLinksGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLinksGet200ResponseData(val *UserLinksGet200ResponseData) *NullableUserLinksGet200ResponseData {
	return &NullableUserLinksGet200ResponseData{value: val, isSet: true}
}

func (v NullableUserLinksGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLinksGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


