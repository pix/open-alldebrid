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

// checks if the UserLinksGet200ResponseDataLinksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserLinksGet200ResponseDataLinksInner{}

// UserLinksGet200ResponseDataLinksInner struct for UserLinksGet200ResponseDataLinksInner
type UserLinksGet200ResponseDataLinksInner struct {
	Link *string `json:"link,omitempty"`
	Filename *string `json:"filename,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Date *int64 `json:"date,omitempty"`
	Host *string `json:"host,omitempty"`
}

// NewUserLinksGet200ResponseDataLinksInner instantiates a new UserLinksGet200ResponseDataLinksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLinksGet200ResponseDataLinksInner() *UserLinksGet200ResponseDataLinksInner {
	this := UserLinksGet200ResponseDataLinksInner{}
	return &this
}

// NewUserLinksGet200ResponseDataLinksInnerWithDefaults instantiates a new UserLinksGet200ResponseDataLinksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLinksGet200ResponseDataLinksInnerWithDefaults() *UserLinksGet200ResponseDataLinksInner {
	this := UserLinksGet200ResponseDataLinksInner{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *UserLinksGet200ResponseDataLinksInner) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinksGet200ResponseDataLinksInner) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *UserLinksGet200ResponseDataLinksInner) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *UserLinksGet200ResponseDataLinksInner) SetLink(v string) {
	o.Link = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *UserLinksGet200ResponseDataLinksInner) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinksGet200ResponseDataLinksInner) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *UserLinksGet200ResponseDataLinksInner) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *UserLinksGet200ResponseDataLinksInner) SetFilename(v string) {
	o.Filename = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *UserLinksGet200ResponseDataLinksInner) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinksGet200ResponseDataLinksInner) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *UserLinksGet200ResponseDataLinksInner) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *UserLinksGet200ResponseDataLinksInner) SetSize(v int64) {
	o.Size = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *UserLinksGet200ResponseDataLinksInner) GetDate() int64 {
	if o == nil || IsNil(o.Date) {
		var ret int64
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinksGet200ResponseDataLinksInner) GetDateOk() (*int64, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *UserLinksGet200ResponseDataLinksInner) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given int64 and assigns it to the Date field.
func (o *UserLinksGet200ResponseDataLinksInner) SetDate(v int64) {
	o.Date = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *UserLinksGet200ResponseDataLinksInner) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinksGet200ResponseDataLinksInner) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *UserLinksGet200ResponseDataLinksInner) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *UserLinksGet200ResponseDataLinksInner) SetHost(v string) {
	o.Host = &v
}

func (o UserLinksGet200ResponseDataLinksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserLinksGet200ResponseDataLinksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	return toSerialize, nil
}

type NullableUserLinksGet200ResponseDataLinksInner struct {
	value *UserLinksGet200ResponseDataLinksInner
	isSet bool
}

func (v NullableUserLinksGet200ResponseDataLinksInner) Get() *UserLinksGet200ResponseDataLinksInner {
	return v.value
}

func (v *NullableUserLinksGet200ResponseDataLinksInner) Set(val *UserLinksGet200ResponseDataLinksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLinksGet200ResponseDataLinksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLinksGet200ResponseDataLinksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLinksGet200ResponseDataLinksInner(val *UserLinksGet200ResponseDataLinksInner) *NullableUserLinksGet200ResponseDataLinksInner {
	return &NullableUserLinksGet200ResponseDataLinksInner{value: val, isSet: true}
}

func (v NullableUserLinksGet200ResponseDataLinksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLinksGet200ResponseDataLinksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


