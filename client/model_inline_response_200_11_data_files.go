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

// InlineResponse20011DataFiles struct for InlineResponse20011DataFiles
type InlineResponse20011DataFiles struct {
	File *string `json:"file,omitempty"`
	Hash *string `json:"hash,omitempty"`
	Name *string `json:"name,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Ready *bool `json:"ready,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Error *InlineResponse20011DataError `json:"error,omitempty"`
}

// NewInlineResponse20011DataFiles instantiates a new InlineResponse20011DataFiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20011DataFiles() *InlineResponse20011DataFiles {
	this := InlineResponse20011DataFiles{}
	return &this
}

// NewInlineResponse20011DataFilesWithDefaults instantiates a new InlineResponse20011DataFiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20011DataFilesWithDefaults() *InlineResponse20011DataFiles {
	this := InlineResponse20011DataFiles{}
	return &this
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *InlineResponse20011DataFiles) GetFile() string {
	if o == nil || o.File == nil {
		var ret string
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011DataFiles) GetFileOk() (*string, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *InlineResponse20011DataFiles) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given string and assigns it to the File field.
func (o *InlineResponse20011DataFiles) SetFile(v string) {
	o.File = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *InlineResponse20011DataFiles) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011DataFiles) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *InlineResponse20011DataFiles) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *InlineResponse20011DataFiles) SetHash(v string) {
	o.Hash = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20011DataFiles) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011DataFiles) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20011DataFiles) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20011DataFiles) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *InlineResponse20011DataFiles) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011DataFiles) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *InlineResponse20011DataFiles) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *InlineResponse20011DataFiles) SetSize(v int64) {
	o.Size = &v
}

// GetReady returns the Ready field value if set, zero value otherwise.
func (o *InlineResponse20011DataFiles) GetReady() bool {
	if o == nil || o.Ready == nil {
		var ret bool
		return ret
	}
	return *o.Ready
}

// GetReadyOk returns a tuple with the Ready field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011DataFiles) GetReadyOk() (*bool, bool) {
	if o == nil || o.Ready == nil {
		return nil, false
	}
	return o.Ready, true
}

// HasReady returns a boolean if a field has been set.
func (o *InlineResponse20011DataFiles) HasReady() bool {
	if o != nil && o.Ready != nil {
		return true
	}

	return false
}

// SetReady gets a reference to the given bool and assigns it to the Ready field.
func (o *InlineResponse20011DataFiles) SetReady(v bool) {
	o.Ready = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20011DataFiles) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011DataFiles) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20011DataFiles) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *InlineResponse20011DataFiles) SetId(v int64) {
	o.Id = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InlineResponse20011DataFiles) GetError() InlineResponse20011DataError {
	if o == nil || o.Error == nil {
		var ret InlineResponse20011DataError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011DataFiles) GetErrorOk() (*InlineResponse20011DataError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InlineResponse20011DataFiles) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given InlineResponse20011DataError and assigns it to the Error field.
func (o *InlineResponse20011DataFiles) SetError(v InlineResponse20011DataError) {
	o.Error = &v
}

func (o InlineResponse20011DataFiles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Ready != nil {
		toSerialize["ready"] = o.Ready
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20011DataFiles struct {
	value *InlineResponse20011DataFiles
	isSet bool
}

func (v NullableInlineResponse20011DataFiles) Get() *InlineResponse20011DataFiles {
	return v.value
}

func (v *NullableInlineResponse20011DataFiles) Set(val *InlineResponse20011DataFiles) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20011DataFiles) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20011DataFiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20011DataFiles(val *InlineResponse20011DataFiles) *NullableInlineResponse20011DataFiles {
	return &NullableInlineResponse20011DataFiles{value: val, isSet: true}
}

func (v NullableInlineResponse20011DataFiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20011DataFiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

