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

// checks if the HostsGet200ResponseDataStreamsValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostsGet200ResponseDataStreamsValue{}

// HostsGet200ResponseDataStreamsValue struct for HostsGet200ResponseDataStreamsValue
type HostsGet200ResponseDataStreamsValue struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Domains []string `json:"domains,omitempty"`
	Regexp *string `json:"regexp,omitempty"`
	Regexps []string `json:"regexps,omitempty"`
	Status *bool `json:"status,omitempty"`
}

// NewHostsGet200ResponseDataStreamsValue instantiates a new HostsGet200ResponseDataStreamsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostsGet200ResponseDataStreamsValue() *HostsGet200ResponseDataStreamsValue {
	this := HostsGet200ResponseDataStreamsValue{}
	return &this
}

// NewHostsGet200ResponseDataStreamsValueWithDefaults instantiates a new HostsGet200ResponseDataStreamsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostsGet200ResponseDataStreamsValueWithDefaults() *HostsGet200ResponseDataStreamsValue {
	this := HostsGet200ResponseDataStreamsValue{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HostsGet200ResponseDataStreamsValue) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostsGet200ResponseDataStreamsValue) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HostsGet200ResponseDataStreamsValue) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HostsGet200ResponseDataStreamsValue) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HostsGet200ResponseDataStreamsValue) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostsGet200ResponseDataStreamsValue) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HostsGet200ResponseDataStreamsValue) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HostsGet200ResponseDataStreamsValue) SetType(v string) {
	o.Type = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *HostsGet200ResponseDataStreamsValue) GetDomains() []string {
	if o == nil || IsNil(o.Domains) {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostsGet200ResponseDataStreamsValue) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *HostsGet200ResponseDataStreamsValue) HasDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *HostsGet200ResponseDataStreamsValue) SetDomains(v []string) {
	o.Domains = v
}

// GetRegexp returns the Regexp field value if set, zero value otherwise.
func (o *HostsGet200ResponseDataStreamsValue) GetRegexp() string {
	if o == nil || IsNil(o.Regexp) {
		var ret string
		return ret
	}
	return *o.Regexp
}

// GetRegexpOk returns a tuple with the Regexp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostsGet200ResponseDataStreamsValue) GetRegexpOk() (*string, bool) {
	if o == nil || IsNil(o.Regexp) {
		return nil, false
	}
	return o.Regexp, true
}

// HasRegexp returns a boolean if a field has been set.
func (o *HostsGet200ResponseDataStreamsValue) HasRegexp() bool {
	if o != nil && !IsNil(o.Regexp) {
		return true
	}

	return false
}

// SetRegexp gets a reference to the given string and assigns it to the Regexp field.
func (o *HostsGet200ResponseDataStreamsValue) SetRegexp(v string) {
	o.Regexp = &v
}

// GetRegexps returns the Regexps field value if set, zero value otherwise.
func (o *HostsGet200ResponseDataStreamsValue) GetRegexps() []string {
	if o == nil || IsNil(o.Regexps) {
		var ret []string
		return ret
	}
	return o.Regexps
}

// GetRegexpsOk returns a tuple with the Regexps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostsGet200ResponseDataStreamsValue) GetRegexpsOk() ([]string, bool) {
	if o == nil || IsNil(o.Regexps) {
		return nil, false
	}
	return o.Regexps, true
}

// HasRegexps returns a boolean if a field has been set.
func (o *HostsGet200ResponseDataStreamsValue) HasRegexps() bool {
	if o != nil && !IsNil(o.Regexps) {
		return true
	}

	return false
}

// SetRegexps gets a reference to the given []string and assigns it to the Regexps field.
func (o *HostsGet200ResponseDataStreamsValue) SetRegexps(v []string) {
	o.Regexps = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HostsGet200ResponseDataStreamsValue) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostsGet200ResponseDataStreamsValue) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HostsGet200ResponseDataStreamsValue) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *HostsGet200ResponseDataStreamsValue) SetStatus(v bool) {
	o.Status = &v
}

func (o HostsGet200ResponseDataStreamsValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostsGet200ResponseDataStreamsValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}
	if !IsNil(o.Regexp) {
		toSerialize["regexp"] = o.Regexp
	}
	if !IsNil(o.Regexps) {
		toSerialize["regexps"] = o.Regexps
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableHostsGet200ResponseDataStreamsValue struct {
	value *HostsGet200ResponseDataStreamsValue
	isSet bool
}

func (v NullableHostsGet200ResponseDataStreamsValue) Get() *HostsGet200ResponseDataStreamsValue {
	return v.value
}

func (v *NullableHostsGet200ResponseDataStreamsValue) Set(val *HostsGet200ResponseDataStreamsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableHostsGet200ResponseDataStreamsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableHostsGet200ResponseDataStreamsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostsGet200ResponseDataStreamsValue(val *HostsGet200ResponseDataStreamsValue) *NullableHostsGet200ResponseDataStreamsValue {
	return &NullableHostsGet200ResponseDataStreamsValue{value: val, isSet: true}
}

func (v NullableHostsGet200ResponseDataStreamsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostsGet200ResponseDataStreamsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


