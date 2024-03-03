# UserGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**UserGet200ResponseData**](UserGet200ResponseData.md) |  | [optional] 
**Error** | Pointer to [**HostsGet200ResponseError**](HostsGet200ResponseError.md) |  | [optional] 

## Methods

### NewUserGet200Response

`func NewUserGet200Response() *UserGet200Response`

NewUserGet200Response instantiates a new UserGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGet200ResponseWithDefaults

`func NewUserGet200ResponseWithDefaults() *UserGet200Response`

NewUserGet200ResponseWithDefaults instantiates a new UserGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UserGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *UserGet200Response) GetData() UserGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserGet200Response) GetDataOk() (*UserGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserGet200Response) SetData(v UserGet200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *UserGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *UserGet200Response) GetError() HostsGet200ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UserGet200Response) GetErrorOk() (*HostsGet200ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UserGet200Response) SetError(v HostsGet200ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *UserGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


