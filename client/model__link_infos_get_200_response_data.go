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

// checks if the LinkInfosGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkInfosGet200ResponseData{}

// LinkInfosGet200ResponseData struct for LinkInfosGet200ResponseData
type LinkInfosGet200ResponseData struct {
	// Array of info objects
	Infos []LinkInfosGet200ResponseDataInfosInner `json:"infos,omitempty"`
}

// NewLinkInfosGet200ResponseData instantiates a new LinkInfosGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkInfosGet200ResponseData() *LinkInfosGet200ResponseData {
	this := LinkInfosGet200ResponseData{}
	return &this
}

// NewLinkInfosGet200ResponseDataWithDefaults instantiates a new LinkInfosGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkInfosGet200ResponseDataWithDefaults() *LinkInfosGet200ResponseData {
	this := LinkInfosGet200ResponseData{}
	return &this
}

// GetInfos returns the Infos field value if set, zero value otherwise.
func (o *LinkInfosGet200ResponseData) GetInfos() []LinkInfosGet200ResponseDataInfosInner {
	if o == nil || IsNil(o.Infos) {
		var ret []LinkInfosGet200ResponseDataInfosInner
		return ret
	}
	return o.Infos
}

// GetInfosOk returns a tuple with the Infos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkInfosGet200ResponseData) GetInfosOk() ([]LinkInfosGet200ResponseDataInfosInner, bool) {
	if o == nil || IsNil(o.Infos) {
		return nil, false
	}
	return o.Infos, true
}

// HasInfos returns a boolean if a field has been set.
func (o *LinkInfosGet200ResponseData) HasInfos() bool {
	if o != nil && !IsNil(o.Infos) {
		return true
	}

	return false
}

// SetInfos gets a reference to the given []LinkInfosGet200ResponseDataInfosInner and assigns it to the Infos field.
func (o *LinkInfosGet200ResponseData) SetInfos(v []LinkInfosGet200ResponseDataInfosInner) {
	o.Infos = v
}

func (o LinkInfosGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkInfosGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Infos) {
		toSerialize["infos"] = o.Infos
	}
	return toSerialize, nil
}

type NullableLinkInfosGet200ResponseData struct {
	value *LinkInfosGet200ResponseData
	isSet bool
}

func (v NullableLinkInfosGet200ResponseData) Get() *LinkInfosGet200ResponseData {
	return v.value
}

func (v *NullableLinkInfosGet200ResponseData) Set(val *LinkInfosGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkInfosGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkInfosGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkInfosGet200ResponseData(val *LinkInfosGet200ResponseData) *NullableLinkInfosGet200ResponseData {
	return &NullableLinkInfosGet200ResponseData{value: val, isSet: true}
}

func (v NullableLinkInfosGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkInfosGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


