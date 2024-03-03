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

// checks if the UserGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGet200ResponseData{}

// UserGet200ResponseData struct for UserGet200ResponseData
type UserGet200ResponseData struct {
	User *UserGet200ResponseDataUser `json:"user,omitempty"`
}

// NewUserGet200ResponseData instantiates a new UserGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGet200ResponseData() *UserGet200ResponseData {
	this := UserGet200ResponseData{}
	return &this
}

// NewUserGet200ResponseDataWithDefaults instantiates a new UserGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGet200ResponseDataWithDefaults() *UserGet200ResponseData {
	this := UserGet200ResponseData{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserGet200ResponseData) GetUser() UserGet200ResponseDataUser {
	if o == nil || IsNil(o.User) {
		var ret UserGet200ResponseDataUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGet200ResponseData) GetUserOk() (*UserGet200ResponseDataUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserGet200ResponseData) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserGet200ResponseDataUser and assigns it to the User field.
func (o *UserGet200ResponseData) SetUser(v UserGet200ResponseDataUser) {
	o.User = &v
}

func (o UserGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUserGet200ResponseData struct {
	value *UserGet200ResponseData
	isSet bool
}

func (v NullableUserGet200ResponseData) Get() *UserGet200ResponseData {
	return v.value
}

func (v *NullableUserGet200ResponseData) Set(val *UserGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGet200ResponseData(val *UserGet200ResponseData) *NullableUserGet200ResponseData {
	return &NullableUserGet200ResponseData{value: val, isSet: true}
}

func (v NullableUserGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

