# UserGet200ResponseDataUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**IsPremium** | Pointer to **bool** |  | [optional] 
**IsSubscribed** | Pointer to **bool** |  | [optional] 
**IsTrial** | Pointer to **bool** |  | [optional] 
**PremiumUntil** | Pointer to **int64** |  | [optional] 
**Lang** | Pointer to **string** |  | [optional] 
**PreferedDomain** | Pointer to **string** |  | [optional] 
**FidelityPoints** | Pointer to **int64** |  | [optional] 
**LimitedHostersQuotas** | Pointer to **map[string]int64** |  | [optional] 
**Notifications** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUserGet200ResponseDataUser

`func NewUserGet200ResponseDataUser() *UserGet200ResponseDataUser`

NewUserGet200ResponseDataUser instantiates a new UserGet200ResponseDataUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGet200ResponseDataUserWithDefaults

`func NewUserGet200ResponseDataUserWithDefaults() *UserGet200ResponseDataUser`

NewUserGet200ResponseDataUserWithDefaults instantiates a new UserGet200ResponseDataUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserGet200ResponseDataUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserGet200ResponseDataUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserGet200ResponseDataUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserGet200ResponseDataUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *UserGet200ResponseDataUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserGet200ResponseDataUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserGet200ResponseDataUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserGet200ResponseDataUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIsPremium

`func (o *UserGet200ResponseDataUser) GetIsPremium() bool`

GetIsPremium returns the IsPremium field if non-nil, zero value otherwise.

### GetIsPremiumOk

`func (o *UserGet200ResponseDataUser) GetIsPremiumOk() (*bool, bool)`

GetIsPremiumOk returns a tuple with the IsPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPremium

`func (o *UserGet200ResponseDataUser) SetIsPremium(v bool)`

SetIsPremium sets IsPremium field to given value.

### HasIsPremium

`func (o *UserGet200ResponseDataUser) HasIsPremium() bool`

HasIsPremium returns a boolean if a field has been set.

### GetIsSubscribed

`func (o *UserGet200ResponseDataUser) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *UserGet200ResponseDataUser) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *UserGet200ResponseDataUser) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *UserGet200ResponseDataUser) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### GetIsTrial

`func (o *UserGet200ResponseDataUser) GetIsTrial() bool`

GetIsTrial returns the IsTrial field if non-nil, zero value otherwise.

### GetIsTrialOk

`func (o *UserGet200ResponseDataUser) GetIsTrialOk() (*bool, bool)`

GetIsTrialOk returns a tuple with the IsTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrial

`func (o *UserGet200ResponseDataUser) SetIsTrial(v bool)`

SetIsTrial sets IsTrial field to given value.

### HasIsTrial

`func (o *UserGet200ResponseDataUser) HasIsTrial() bool`

HasIsTrial returns a boolean if a field has been set.

### GetPremiumUntil

`func (o *UserGet200ResponseDataUser) GetPremiumUntil() int64`

GetPremiumUntil returns the PremiumUntil field if non-nil, zero value otherwise.

### GetPremiumUntilOk

`func (o *UserGet200ResponseDataUser) GetPremiumUntilOk() (*int64, bool)`

GetPremiumUntilOk returns a tuple with the PremiumUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumUntil

`func (o *UserGet200ResponseDataUser) SetPremiumUntil(v int64)`

SetPremiumUntil sets PremiumUntil field to given value.

### HasPremiumUntil

`func (o *UserGet200ResponseDataUser) HasPremiumUntil() bool`

HasPremiumUntil returns a boolean if a field has been set.

### GetLang

`func (o *UserGet200ResponseDataUser) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *UserGet200ResponseDataUser) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *UserGet200ResponseDataUser) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *UserGet200ResponseDataUser) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetPreferedDomain

`func (o *UserGet200ResponseDataUser) GetPreferedDomain() string`

GetPreferedDomain returns the PreferedDomain field if non-nil, zero value otherwise.

### GetPreferedDomainOk

`func (o *UserGet200ResponseDataUser) GetPreferedDomainOk() (*string, bool)`

GetPreferedDomainOk returns a tuple with the PreferedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferedDomain

`func (o *UserGet200ResponseDataUser) SetPreferedDomain(v string)`

SetPreferedDomain sets PreferedDomain field to given value.

### HasPreferedDomain

`func (o *UserGet200ResponseDataUser) HasPreferedDomain() bool`

HasPreferedDomain returns a boolean if a field has been set.

### GetFidelityPoints

`func (o *UserGet200ResponseDataUser) GetFidelityPoints() int64`

GetFidelityPoints returns the FidelityPoints field if non-nil, zero value otherwise.

### GetFidelityPointsOk

`func (o *UserGet200ResponseDataUser) GetFidelityPointsOk() (*int64, bool)`

GetFidelityPointsOk returns a tuple with the FidelityPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFidelityPoints

`func (o *UserGet200ResponseDataUser) SetFidelityPoints(v int64)`

SetFidelityPoints sets FidelityPoints field to given value.

### HasFidelityPoints

`func (o *UserGet200ResponseDataUser) HasFidelityPoints() bool`

HasFidelityPoints returns a boolean if a field has been set.

### GetLimitedHostersQuotas

`func (o *UserGet200ResponseDataUser) GetLimitedHostersQuotas() map[string]int64`

GetLimitedHostersQuotas returns the LimitedHostersQuotas field if non-nil, zero value otherwise.

### GetLimitedHostersQuotasOk

`func (o *UserGet200ResponseDataUser) GetLimitedHostersQuotasOk() (*map[string]int64, bool)`

GetLimitedHostersQuotasOk returns a tuple with the LimitedHostersQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedHostersQuotas

`func (o *UserGet200ResponseDataUser) SetLimitedHostersQuotas(v map[string]int64)`

SetLimitedHostersQuotas sets LimitedHostersQuotas field to given value.

### HasLimitedHostersQuotas

`func (o *UserGet200ResponseDataUser) HasLimitedHostersQuotas() bool`

HasLimitedHostersQuotas returns a boolean if a field has been set.

### GetNotifications

`func (o *UserGet200ResponseDataUser) GetNotifications() []string`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *UserGet200ResponseDataUser) GetNotificationsOk() (*[]string, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *UserGet200ResponseDataUser) SetNotifications(v []string)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *UserGet200ResponseDataUser) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


