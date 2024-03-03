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

// checks if the LinkDelayedGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkDelayedGet200ResponseData{}

// LinkDelayedGet200ResponseData struct for LinkDelayedGet200ResponseData
type LinkDelayedGet200ResponseData struct {
	Status *int64 `json:"status,omitempty"`
	TimeLeft *int64 `json:"time_left,omitempty"`
	Link *string `json:"link,omitempty"`
}

// NewLinkDelayedGet200ResponseData instantiates a new LinkDelayedGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkDelayedGet200ResponseData() *LinkDelayedGet200ResponseData {
	this := LinkDelayedGet200ResponseData{}
	return &this
}

// NewLinkDelayedGet200ResponseDataWithDefaults instantiates a new LinkDelayedGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDelayedGet200ResponseDataWithDefaults() *LinkDelayedGet200ResponseData {
	this := LinkDelayedGet200ResponseData{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LinkDelayedGet200ResponseData) GetStatus() int64 {
	if o == nil || IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDelayedGet200ResponseData) GetStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LinkDelayedGet200ResponseData) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *LinkDelayedGet200ResponseData) SetStatus(v int64) {
	o.Status = &v
}

// GetTimeLeft returns the TimeLeft field value if set, zero value otherwise.
func (o *LinkDelayedGet200ResponseData) GetTimeLeft() int64 {
	if o == nil || IsNil(o.TimeLeft) {
		var ret int64
		return ret
	}
	return *o.TimeLeft
}

// GetTimeLeftOk returns a tuple with the TimeLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDelayedGet200ResponseData) GetTimeLeftOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeLeft) {
		return nil, false
	}
	return o.TimeLeft, true
}

// HasTimeLeft returns a boolean if a field has been set.
func (o *LinkDelayedGet200ResponseData) HasTimeLeft() bool {
	if o != nil && !IsNil(o.TimeLeft) {
		return true
	}

	return false
}

// SetTimeLeft gets a reference to the given int64 and assigns it to the TimeLeft field.
func (o *LinkDelayedGet200ResponseData) SetTimeLeft(v int64) {
	o.TimeLeft = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *LinkDelayedGet200ResponseData) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDelayedGet200ResponseData) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *LinkDelayedGet200ResponseData) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *LinkDelayedGet200ResponseData) SetLink(v string) {
	o.Link = &v
}

func (o LinkDelayedGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkDelayedGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TimeLeft) {
		toSerialize["time_left"] = o.TimeLeft
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	return toSerialize, nil
}

type NullableLinkDelayedGet200ResponseData struct {
	value *LinkDelayedGet200ResponseData
	isSet bool
}

func (v NullableLinkDelayedGet200ResponseData) Get() *LinkDelayedGet200ResponseData {
	return v.value
}

func (v *NullableLinkDelayedGet200ResponseData) Set(val *LinkDelayedGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDelayedGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDelayedGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDelayedGet200ResponseData(val *LinkDelayedGet200ResponseData) *NullableLinkDelayedGet200ResponseData {
	return &NullableLinkDelayedGet200ResponseData{value: val, isSet: true}
}

func (v NullableLinkDelayedGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDelayedGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

