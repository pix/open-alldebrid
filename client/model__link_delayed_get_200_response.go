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

// checks if the LinkDelayedGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkDelayedGet200Response{}

// LinkDelayedGet200Response struct for LinkDelayedGet200Response
type LinkDelayedGet200Response struct {
	Status *string `json:"status,omitempty"`
	Data *LinkDelayedGet200ResponseData `json:"data,omitempty"`
	Error *LinkDelayedGet200ResponseError `json:"error,omitempty"`
}

// NewLinkDelayedGet200Response instantiates a new LinkDelayedGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkDelayedGet200Response() *LinkDelayedGet200Response {
	this := LinkDelayedGet200Response{}
	return &this
}

// NewLinkDelayedGet200ResponseWithDefaults instantiates a new LinkDelayedGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDelayedGet200ResponseWithDefaults() *LinkDelayedGet200Response {
	this := LinkDelayedGet200Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LinkDelayedGet200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDelayedGet200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LinkDelayedGet200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LinkDelayedGet200Response) SetStatus(v string) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LinkDelayedGet200Response) GetData() LinkDelayedGet200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret LinkDelayedGet200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDelayedGet200Response) GetDataOk() (*LinkDelayedGet200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LinkDelayedGet200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given LinkDelayedGet200ResponseData and assigns it to the Data field.
func (o *LinkDelayedGet200Response) SetData(v LinkDelayedGet200ResponseData) {
	o.Data = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *LinkDelayedGet200Response) GetError() LinkDelayedGet200ResponseError {
	if o == nil || IsNil(o.Error) {
		var ret LinkDelayedGet200ResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDelayedGet200Response) GetErrorOk() (*LinkDelayedGet200ResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *LinkDelayedGet200Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given LinkDelayedGet200ResponseError and assigns it to the Error field.
func (o *LinkDelayedGet200Response) SetError(v LinkDelayedGet200ResponseError) {
	o.Error = &v
}

func (o LinkDelayedGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkDelayedGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableLinkDelayedGet200Response struct {
	value *LinkDelayedGet200Response
	isSet bool
}

func (v NullableLinkDelayedGet200Response) Get() *LinkDelayedGet200Response {
	return v.value
}

func (v *NullableLinkDelayedGet200Response) Set(val *LinkDelayedGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDelayedGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDelayedGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDelayedGet200Response(val *LinkDelayedGet200Response) *NullableLinkDelayedGet200Response {
	return &NullableLinkDelayedGet200Response{value: val, isSet: true}
}

func (v NullableLinkDelayedGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDelayedGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


