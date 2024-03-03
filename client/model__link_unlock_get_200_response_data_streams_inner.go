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

// checks if the LinkUnlockGet200ResponseDataStreamsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkUnlockGet200ResponseDataStreamsInner{}

// LinkUnlockGet200ResponseDataStreamsInner struct for LinkUnlockGet200ResponseDataStreamsInner
type LinkUnlockGet200ResponseDataStreamsInner struct {
	Id *string `json:"id,omitempty"`
	Ext *string `json:"ext,omitempty"`
	Quality *string `json:"quality,omitempty"`
	Filesize *int64 `json:"filesize,omitempty"`
	Proto *string `json:"proto,omitempty"`
	Name *string `json:"name,omitempty"`
	Link *string `json:"link,omitempty"`
	Tb *float32 `json:"tb,omitempty"`
	Abr *int64 `json:"abr,omitempty"`
}

// NewLinkUnlockGet200ResponseDataStreamsInner instantiates a new LinkUnlockGet200ResponseDataStreamsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkUnlockGet200ResponseDataStreamsInner() *LinkUnlockGet200ResponseDataStreamsInner {
	this := LinkUnlockGet200ResponseDataStreamsInner{}
	return &this
}

// NewLinkUnlockGet200ResponseDataStreamsInnerWithDefaults instantiates a new LinkUnlockGet200ResponseDataStreamsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkUnlockGet200ResponseDataStreamsInnerWithDefaults() *LinkUnlockGet200ResponseDataStreamsInner {
	this := LinkUnlockGet200ResponseDataStreamsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LinkUnlockGet200ResponseDataStreamsInner) SetId(v string) {
	o.Id = &v
}

// GetExt returns the Ext field value if set, zero value otherwise.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetExt() string {
	if o == nil || IsNil(o.Ext) {
		var ret string
		return ret
	}
	return *o.Ext
}

// GetExtOk returns a tuple with the Ext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetExtOk() (*string, bool) {
	if o == nil || IsNil(o.Ext) {
		return nil, false
	}
	return o.Ext, true
}

// HasExt returns a boolean if a field has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) HasExt() bool {
	if o != nil && !IsNil(o.Ext) {
		return true
	}

	return false
}

// SetExt gets a reference to the given string and assigns it to the Ext field.
func (o *LinkUnlockGet200ResponseDataStreamsInner) SetExt(v string) {
	o.Ext = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetQuality() string {
	if o == nil || IsNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetQualityOk() (*string, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *LinkUnlockGet200ResponseDataStreamsInner) SetQuality(v string) {
	o.Quality = &v
}

// GetFilesize returns the Filesize field value if set, zero value otherwise.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetFilesize() int64 {
	if o == nil || IsNil(o.Filesize) {
		var ret int64
		return ret
	}
	return *o.Filesize
}

// GetFilesizeOk returns a tuple with the Filesize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetFilesizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Filesize) {
		return nil, false
	}
	return o.Filesize, true
}

// HasFilesize returns a boolean if a field has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) HasFilesize() bool {
	if o != nil && !IsNil(o.Filesize) {
		return true
	}

	return false
}

// SetFilesize gets a reference to the given int64 and assigns it to the Filesize field.
func (o *LinkUnlockGet200ResponseDataStreamsInner) SetFilesize(v int64) {
	o.Filesize = &v
}

// GetProto returns the Proto field value if set, zero value otherwise.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetProto() string {
	if o == nil || IsNil(o.Proto) {
		var ret string
		return ret
	}
	return *o.Proto
}

// GetProtoOk returns a tuple with the Proto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetProtoOk() (*string, bool) {
	if o == nil || IsNil(o.Proto) {
		return nil, false
	}
	return o.Proto, true
}

// HasProto returns a boolean if a field has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) HasProto() bool {
	if o != nil && !IsNil(o.Proto) {
		return true
	}

	return false
}

// SetProto gets a reference to the given string and assigns it to the Proto field.
func (o *LinkUnlockGet200ResponseDataStreamsInner) SetProto(v string) {
	o.Proto = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LinkUnlockGet200ResponseDataStreamsInner) SetName(v string) {
	o.Name = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *LinkUnlockGet200ResponseDataStreamsInner) SetLink(v string) {
	o.Link = &v
}

// GetTb returns the Tb field value if set, zero value otherwise.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetTb() float32 {
	if o == nil || IsNil(o.Tb) {
		var ret float32
		return ret
	}
	return *o.Tb
}

// GetTbOk returns a tuple with the Tb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetTbOk() (*float32, bool) {
	if o == nil || IsNil(o.Tb) {
		return nil, false
	}
	return o.Tb, true
}

// HasTb returns a boolean if a field has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) HasTb() bool {
	if o != nil && !IsNil(o.Tb) {
		return true
	}

	return false
}

// SetTb gets a reference to the given float32 and assigns it to the Tb field.
func (o *LinkUnlockGet200ResponseDataStreamsInner) SetTb(v float32) {
	o.Tb = &v
}

// GetAbr returns the Abr field value if set, zero value otherwise.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetAbr() int64 {
	if o == nil || IsNil(o.Abr) {
		var ret int64
		return ret
	}
	return *o.Abr
}

// GetAbrOk returns a tuple with the Abr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) GetAbrOk() (*int64, bool) {
	if o == nil || IsNil(o.Abr) {
		return nil, false
	}
	return o.Abr, true
}

// HasAbr returns a boolean if a field has been set.
func (o *LinkUnlockGet200ResponseDataStreamsInner) HasAbr() bool {
	if o != nil && !IsNil(o.Abr) {
		return true
	}

	return false
}

// SetAbr gets a reference to the given int64 and assigns it to the Abr field.
func (o *LinkUnlockGet200ResponseDataStreamsInner) SetAbr(v int64) {
	o.Abr = &v
}

func (o LinkUnlockGet200ResponseDataStreamsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkUnlockGet200ResponseDataStreamsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ext) {
		toSerialize["ext"] = o.Ext
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if !IsNil(o.Filesize) {
		toSerialize["filesize"] = o.Filesize
	}
	if !IsNil(o.Proto) {
		toSerialize["proto"] = o.Proto
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Tb) {
		toSerialize["tb"] = o.Tb
	}
	if !IsNil(o.Abr) {
		toSerialize["abr"] = o.Abr
	}
	return toSerialize, nil
}

type NullableLinkUnlockGet200ResponseDataStreamsInner struct {
	value *LinkUnlockGet200ResponseDataStreamsInner
	isSet bool
}

func (v NullableLinkUnlockGet200ResponseDataStreamsInner) Get() *LinkUnlockGet200ResponseDataStreamsInner {
	return v.value
}

func (v *NullableLinkUnlockGet200ResponseDataStreamsInner) Set(val *LinkUnlockGet200ResponseDataStreamsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkUnlockGet200ResponseDataStreamsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkUnlockGet200ResponseDataStreamsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkUnlockGet200ResponseDataStreamsInner(val *LinkUnlockGet200ResponseDataStreamsInner) *NullableLinkUnlockGet200ResponseDataStreamsInner {
	return &NullableLinkUnlockGet200ResponseDataStreamsInner{value: val, isSet: true}
}

func (v NullableLinkUnlockGet200ResponseDataStreamsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkUnlockGet200ResponseDataStreamsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


