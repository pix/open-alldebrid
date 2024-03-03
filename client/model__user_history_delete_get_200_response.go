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

// checks if the UserHistoryDeleteGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserHistoryDeleteGet200Response{}

// UserHistoryDeleteGet200Response struct for UserHistoryDeleteGet200Response
type UserHistoryDeleteGet200Response struct {
	Status *string `json:"status,omitempty"`
	Data *UserHistoryDeleteGet200ResponseData `json:"data,omitempty"`
	Error *MagnetUploadGet200ResponseError `json:"error,omitempty"`
}

// NewUserHistoryDeleteGet200Response instantiates a new UserHistoryDeleteGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserHistoryDeleteGet200Response() *UserHistoryDeleteGet200Response {
	this := UserHistoryDeleteGet200Response{}
	return &this
}

// NewUserHistoryDeleteGet200ResponseWithDefaults instantiates a new UserHistoryDeleteGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserHistoryDeleteGet200ResponseWithDefaults() *UserHistoryDeleteGet200Response {
	this := UserHistoryDeleteGet200Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserHistoryDeleteGet200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserHistoryDeleteGet200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserHistoryDeleteGet200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UserHistoryDeleteGet200Response) SetStatus(v string) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UserHistoryDeleteGet200Response) GetData() UserHistoryDeleteGet200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret UserHistoryDeleteGet200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserHistoryDeleteGet200Response) GetDataOk() (*UserHistoryDeleteGet200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UserHistoryDeleteGet200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given UserHistoryDeleteGet200ResponseData and assigns it to the Data field.
func (o *UserHistoryDeleteGet200Response) SetData(v UserHistoryDeleteGet200ResponseData) {
	o.Data = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *UserHistoryDeleteGet200Response) GetError() MagnetUploadGet200ResponseError {
	if o == nil || IsNil(o.Error) {
		var ret MagnetUploadGet200ResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserHistoryDeleteGet200Response) GetErrorOk() (*MagnetUploadGet200ResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *UserHistoryDeleteGet200Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given MagnetUploadGet200ResponseError and assigns it to the Error field.
func (o *UserHistoryDeleteGet200Response) SetError(v MagnetUploadGet200ResponseError) {
	o.Error = &v
}

func (o UserHistoryDeleteGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserHistoryDeleteGet200Response) ToMap() (map[string]interface{}, error) {
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

type NullableUserHistoryDeleteGet200Response struct {
	value *UserHistoryDeleteGet200Response
	isSet bool
}

func (v NullableUserHistoryDeleteGet200Response) Get() *UserHistoryDeleteGet200Response {
	return v.value
}

func (v *NullableUserHistoryDeleteGet200Response) Set(val *UserHistoryDeleteGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUserHistoryDeleteGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUserHistoryDeleteGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserHistoryDeleteGet200Response(val *UserHistoryDeleteGet200Response) *NullableUserHistoryDeleteGet200Response {
	return &NullableUserHistoryDeleteGet200Response{value: val, isSet: true}
}

func (v NullableUserHistoryDeleteGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserHistoryDeleteGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

