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

// checks if the UserGet200ResponseDataUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGet200ResponseDataUser{}

// UserGet200ResponseDataUser struct for UserGet200ResponseDataUser
type UserGet200ResponseDataUser struct {
	Username *string `json:"username,omitempty"`
	Email *string `json:"email,omitempty"`
	IsPremium *bool `json:"isPremium,omitempty"`
	IsSubscribed *bool `json:"isSubscribed,omitempty"`
	IsTrial *bool `json:"isTrial,omitempty"`
	PremiumUntil *int64 `json:"premiumUntil,omitempty"`
	Lang *string `json:"lang,omitempty"`
	PreferedDomain *string `json:"preferedDomain,omitempty"`
	FidelityPoints *int64 `json:"fidelityPoints,omitempty"`
	LimitedHostersQuotas *map[string]int64 `json:"limitedHostersQuotas,omitempty"`
	Notifications []string `json:"notifications,omitempty"`
}

// NewUserGet200ResponseDataUser instantiates a new UserGet200ResponseDataUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGet200ResponseDataUser() *UserGet200ResponseDataUser {
	this := UserGet200ResponseDataUser{}
	return &this
}

// NewUserGet200ResponseDataUserWithDefaults instantiates a new UserGet200ResponseDataUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGet200ResponseDataUserWithDefaults() *UserGet200ResponseDataUser {
	this := UserGet200ResponseDataUser{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserGet200ResponseDataUser) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGet200ResponseDataUser) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserGet200ResponseDataUser) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserGet200ResponseDataUser) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserGet200ResponseDataUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGet200ResponseDataUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserGet200ResponseDataUser) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserGet200ResponseDataUser) SetEmail(v string) {
	o.Email = &v
}

// GetIsPremium returns the IsPremium field value if set, zero value otherwise.
func (o *UserGet200ResponseDataUser) GetIsPremium() bool {
	if o == nil || IsNil(o.IsPremium) {
		var ret bool
		return ret
	}
	return *o.IsPremium
}

// GetIsPremiumOk returns a tuple with the IsPremium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGet200ResponseDataUser) GetIsPremiumOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPremium) {
		return nil, false
	}
	return o.IsPremium, true
}

// HasIsPremium returns a boolean if a field has been set.
func (o *UserGet200ResponseDataUser) HasIsPremium() bool {
	if o != nil && !IsNil(o.IsPremium) {
		return true
	}

	return false
}

// SetIsPremium gets a reference to the given bool and assigns it to the IsPremium field.
func (o *UserGet200ResponseDataUser) SetIsPremium(v bool) {
	o.IsPremium = &v
}

// GetIsSubscribed returns the IsSubscribed field value if set, zero value otherwise.
func (o *UserGet200ResponseDataUser) GetIsSubscribed() bool {
	if o == nil || IsNil(o.IsSubscribed) {
		var ret bool
		return ret
	}
	return *o.IsSubscribed
}

// GetIsSubscribedOk returns a tuple with the IsSubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGet200ResponseDataUser) GetIsSubscribedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSubscribed) {
		return nil, false
	}
	return o.IsSubscribed, true
}

// HasIsSubscribed returns a boolean if a field has been set.
func (o *UserGet200ResponseDataUser) HasIsSubscribed() bool {
	if o != nil && !IsNil(o.IsSubscribed) {
		return true
	}

	return false
}

// SetIsSubscribed gets a reference to the given bool and assigns it to the IsSubscribed field.
func (o *UserGet200ResponseDataUser) SetIsSubscribed(v bool) {
	o.IsSubscribed = &v
}

// GetIsTrial returns the IsTrial field value if set, zero value otherwise.
func (o *UserGet200ResponseDataUser) GetIsTrial() bool {
	if o == nil || IsNil(o.IsTrial) {
		var ret bool
		return ret
	}
	return *o.IsTrial
}

// GetIsTrialOk returns a tuple with the IsTrial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGet200ResponseDataUser) GetIsTrialOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTrial) {
		return nil, false
	}
	return o.IsTrial, true
}

// HasIsTrial returns a boolean if a field has been set.
func (o *UserGet200ResponseDataUser) HasIsTrial() bool {
	if o != nil && !IsNil(o.IsTrial) {
		return true
	}

	return false
}

// SetIsTrial gets a reference to the given bool and assigns it to the IsTrial field.
func (o *UserGet200ResponseDataUser) SetIsTrial(v bool) {
	o.IsTrial = &v
}

// GetPremiumUntil returns the PremiumUntil field value if set, zero value otherwise.
func (o *UserGet200ResponseDataUser) GetPremiumUntil() int64 {
	if o == nil || IsNil(o.PremiumUntil) {
		var ret int64
		return ret
	}
	return *o.PremiumUntil
}

// GetPremiumUntilOk returns a tuple with the PremiumUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGet200ResponseDataUser) GetPremiumUntilOk() (*int64, bool) {
	if o == nil || IsNil(o.PremiumUntil) {
		return nil, false
	}
	return o.PremiumUntil, true
}

// HasPremiumUntil returns a boolean if a field has been set.
func (o *UserGet200ResponseDataUser) HasPremiumUntil() bool {
	if o != nil && !IsNil(o.PremiumUntil) {
		return true
	}

	return false
}

// SetPremiumUntil gets a reference to the given int64 and assigns it to the PremiumUntil field.
func (o *UserGet200ResponseDataUser) SetPremiumUntil(v int64) {
	o.PremiumUntil = &v
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *UserGet200ResponseDataUser) GetLang() string {
	if o == nil || IsNil(o.Lang) {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGet200ResponseDataUser) GetLangOk() (*string, bool) {
	if o == nil || IsNil(o.Lang) {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *UserGet200ResponseDataUser) HasLang() bool {
	if o != nil && !IsNil(o.Lang) {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *UserGet200ResponseDataUser) SetLang(v string) {
	o.Lang = &v
}

// GetPreferedDomain returns the PreferedDomain field value if set, zero value otherwise.
func (o *UserGet200ResponseDataUser) GetPreferedDomain() string {
	if o == nil || IsNil(o.PreferedDomain) {
		var ret string
		return ret
	}
	return *o.PreferedDomain
}

// GetPreferedDomainOk returns a tuple with the PreferedDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGet200ResponseDataUser) GetPreferedDomainOk() (*string, bool) {
	if o == nil || IsNil(o.PreferedDomain) {
		return nil, false
	}
	return o.PreferedDomain, true
}

// HasPreferedDomain returns a boolean if a field has been set.
func (o *UserGet200ResponseDataUser) HasPreferedDomain() bool {
	if o != nil && !IsNil(o.PreferedDomain) {
		return true
	}

	return false
}

// SetPreferedDomain gets a reference to the given string and assigns it to the PreferedDomain field.
func (o *UserGet200ResponseDataUser) SetPreferedDomain(v string) {
	o.PreferedDomain = &v
}

// GetFidelityPoints returns the FidelityPoints field value if set, zero value otherwise.
func (o *UserGet200ResponseDataUser) GetFidelityPoints() int64 {
	if o == nil || IsNil(o.FidelityPoints) {
		var ret int64
		return ret
	}
	return *o.FidelityPoints
}

// GetFidelityPointsOk returns a tuple with the FidelityPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGet200ResponseDataUser) GetFidelityPointsOk() (*int64, bool) {
	if o == nil || IsNil(o.FidelityPoints) {
		return nil, false
	}
	return o.FidelityPoints, true
}

// HasFidelityPoints returns a boolean if a field has been set.
func (o *UserGet200ResponseDataUser) HasFidelityPoints() bool {
	if o != nil && !IsNil(o.FidelityPoints) {
		return true
	}

	return false
}

// SetFidelityPoints gets a reference to the given int64 and assigns it to the FidelityPoints field.
func (o *UserGet200ResponseDataUser) SetFidelityPoints(v int64) {
	o.FidelityPoints = &v
}

// GetLimitedHostersQuotas returns the LimitedHostersQuotas field value if set, zero value otherwise.
func (o *UserGet200ResponseDataUser) GetLimitedHostersQuotas() map[string]int64 {
	if o == nil || IsNil(o.LimitedHostersQuotas) {
		var ret map[string]int64
		return ret
	}
	return *o.LimitedHostersQuotas
}

// GetLimitedHostersQuotasOk returns a tuple with the LimitedHostersQuotas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGet200ResponseDataUser) GetLimitedHostersQuotasOk() (*map[string]int64, bool) {
	if o == nil || IsNil(o.LimitedHostersQuotas) {
		return nil, false
	}
	return o.LimitedHostersQuotas, true
}

// HasLimitedHostersQuotas returns a boolean if a field has been set.
func (o *UserGet200ResponseDataUser) HasLimitedHostersQuotas() bool {
	if o != nil && !IsNil(o.LimitedHostersQuotas) {
		return true
	}

	return false
}

// SetLimitedHostersQuotas gets a reference to the given map[string]int64 and assigns it to the LimitedHostersQuotas field.
func (o *UserGet200ResponseDataUser) SetLimitedHostersQuotas(v map[string]int64) {
	o.LimitedHostersQuotas = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *UserGet200ResponseDataUser) GetNotifications() []string {
	if o == nil || IsNil(o.Notifications) {
		var ret []string
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGet200ResponseDataUser) GetNotificationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *UserGet200ResponseDataUser) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []string and assigns it to the Notifications field.
func (o *UserGet200ResponseDataUser) SetNotifications(v []string) {
	o.Notifications = v
}

func (o UserGet200ResponseDataUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGet200ResponseDataUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.IsPremium) {
		toSerialize["isPremium"] = o.IsPremium
	}
	if !IsNil(o.IsSubscribed) {
		toSerialize["isSubscribed"] = o.IsSubscribed
	}
	if !IsNil(o.IsTrial) {
		toSerialize["isTrial"] = o.IsTrial
	}
	if !IsNil(o.PremiumUntil) {
		toSerialize["premiumUntil"] = o.PremiumUntil
	}
	if !IsNil(o.Lang) {
		toSerialize["lang"] = o.Lang
	}
	if !IsNil(o.PreferedDomain) {
		toSerialize["preferedDomain"] = o.PreferedDomain
	}
	if !IsNil(o.FidelityPoints) {
		toSerialize["fidelityPoints"] = o.FidelityPoints
	}
	if !IsNil(o.LimitedHostersQuotas) {
		toSerialize["limitedHostersQuotas"] = o.LimitedHostersQuotas
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	return toSerialize, nil
}

type NullableUserGet200ResponseDataUser struct {
	value *UserGet200ResponseDataUser
	isSet bool
}

func (v NullableUserGet200ResponseDataUser) Get() *UserGet200ResponseDataUser {
	return v.value
}

func (v *NullableUserGet200ResponseDataUser) Set(val *UserGet200ResponseDataUser) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGet200ResponseDataUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGet200ResponseDataUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGet200ResponseDataUser(val *UserGet200ResponseDataUser) *NullableUserGet200ResponseDataUser {
	return &NullableUserGet200ResponseDataUser{value: val, isSet: true}
}

func (v NullableUserGet200ResponseDataUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGet200ResponseDataUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


