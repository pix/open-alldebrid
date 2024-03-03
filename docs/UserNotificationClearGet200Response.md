# UserNotificationClearGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**UserNotificationClearGet200ResponseData**](UserNotificationClearGet200ResponseData.md) |  | [optional] 
**Error** | Pointer to [**HostsGet200ResponseError**](HostsGet200ResponseError.md) |  | [optional] 

## Methods

### NewUserNotificationClearGet200Response

`func NewUserNotificationClearGet200Response() *UserNotificationClearGet200Response`

NewUserNotificationClearGet200Response instantiates a new UserNotificationClearGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserNotificationClearGet200ResponseWithDefaults

`func NewUserNotificationClearGet200ResponseWithDefaults() *UserNotificationClearGet200Response`

NewUserNotificationClearGet200ResponseWithDefaults instantiates a new UserNotificationClearGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UserNotificationClearGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserNotificationClearGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserNotificationClearGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserNotificationClearGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *UserNotificationClearGet200Response) GetData() UserNotificationClearGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserNotificationClearGet200Response) GetDataOk() (*UserNotificationClearGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserNotificationClearGet200Response) SetData(v UserNotificationClearGet200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *UserNotificationClearGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *UserNotificationClearGet200Response) GetError() HostsGet200ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UserNotificationClearGet200Response) GetErrorOk() (*HostsGet200ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UserNotificationClearGet200Response) SetError(v HostsGet200ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *UserNotificationClearGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


