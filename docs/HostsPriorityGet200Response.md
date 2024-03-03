# HostsPriorityGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**HostsPriorityGet200ResponseData**](HostsPriorityGet200ResponseData.md) |  | [optional] 
**Error** | Pointer to [**HostsGet200ResponseError**](HostsGet200ResponseError.md) |  | [optional] 

## Methods

### NewHostsPriorityGet200Response

`func NewHostsPriorityGet200Response() *HostsPriorityGet200Response`

NewHostsPriorityGet200Response instantiates a new HostsPriorityGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostsPriorityGet200ResponseWithDefaults

`func NewHostsPriorityGet200ResponseWithDefaults() *HostsPriorityGet200Response`

NewHostsPriorityGet200ResponseWithDefaults instantiates a new HostsPriorityGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HostsPriorityGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostsPriorityGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostsPriorityGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HostsPriorityGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *HostsPriorityGet200Response) GetData() HostsPriorityGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HostsPriorityGet200Response) GetDataOk() (*HostsPriorityGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HostsPriorityGet200Response) SetData(v HostsPriorityGet200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *HostsPriorityGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *HostsPriorityGet200Response) GetError() HostsGet200ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HostsPriorityGet200Response) GetErrorOk() (*HostsGet200ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HostsPriorityGet200Response) SetError(v HostsGet200ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *HostsPriorityGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


