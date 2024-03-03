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

// checks if the MagnetStatusGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MagnetStatusGet200ResponseData{}

// MagnetStatusGet200ResponseData struct for MagnetStatusGet200ResponseData
type MagnetStatusGet200ResponseData struct {
	Magnets []MagnetStatusGet200ResponseDataMagnetsInner `json:"magnets,omitempty"`
}

// NewMagnetStatusGet200ResponseData instantiates a new MagnetStatusGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMagnetStatusGet200ResponseData() *MagnetStatusGet200ResponseData {
	this := MagnetStatusGet200ResponseData{}
	return &this
}

// NewMagnetStatusGet200ResponseDataWithDefaults instantiates a new MagnetStatusGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMagnetStatusGet200ResponseDataWithDefaults() *MagnetStatusGet200ResponseData {
	this := MagnetStatusGet200ResponseData{}
	return &this
}

// GetMagnets returns the Magnets field value if set, zero value otherwise.
func (o *MagnetStatusGet200ResponseData) GetMagnets() []MagnetStatusGet200ResponseDataMagnetsInner {
	if o == nil || IsNil(o.Magnets) {
		var ret []MagnetStatusGet200ResponseDataMagnetsInner
		return ret
	}
	return o.Magnets
}

// GetMagnetsOk returns a tuple with the Magnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MagnetStatusGet200ResponseData) GetMagnetsOk() ([]MagnetStatusGet200ResponseDataMagnetsInner, bool) {
	if o == nil || IsNil(o.Magnets) {
		return nil, false
	}
	return o.Magnets, true
}

// HasMagnets returns a boolean if a field has been set.
func (o *MagnetStatusGet200ResponseData) HasMagnets() bool {
	if o != nil && !IsNil(o.Magnets) {
		return true
	}

	return false
}

// SetMagnets gets a reference to the given []MagnetStatusGet200ResponseDataMagnetsInner and assigns it to the Magnets field.
func (o *MagnetStatusGet200ResponseData) SetMagnets(v []MagnetStatusGet200ResponseDataMagnetsInner) {
	o.Magnets = v
}

func (o MagnetStatusGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MagnetStatusGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Magnets) {
		toSerialize["magnets"] = o.Magnets
	}
	return toSerialize, nil
}

type NullableMagnetStatusGet200ResponseData struct {
	value *MagnetStatusGet200ResponseData
	isSet bool
}

func (v NullableMagnetStatusGet200ResponseData) Get() *MagnetStatusGet200ResponseData {
	return v.value
}

func (v *NullableMagnetStatusGet200ResponseData) Set(val *MagnetStatusGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMagnetStatusGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMagnetStatusGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMagnetStatusGet200ResponseData(val *MagnetStatusGet200ResponseData) *NullableMagnetStatusGet200ResponseData {
	return &NullableMagnetStatusGet200ResponseData{value: val, isSet: true}
}

func (v NullableMagnetStatusGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMagnetStatusGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


