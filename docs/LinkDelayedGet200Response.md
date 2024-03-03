# LinkDelayedGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**LinkDelayedGet200ResponseData**](LinkDelayedGet200ResponseData.md) |  | [optional] 
**Error** | Pointer to [**LinkDelayedGet200ResponseError**](LinkDelayedGet200ResponseError.md) |  | [optional] 

## Methods

### NewLinkDelayedGet200Response

`func NewLinkDelayedGet200Response() *LinkDelayedGet200Response`

NewLinkDelayedGet200Response instantiates a new LinkDelayedGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkDelayedGet200ResponseWithDefaults

`func NewLinkDelayedGet200ResponseWithDefaults() *LinkDelayedGet200Response`

NewLinkDelayedGet200ResponseWithDefaults instantiates a new LinkDelayedGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *LinkDelayedGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LinkDelayedGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LinkDelayedGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LinkDelayedGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *LinkDelayedGet200Response) GetData() LinkDelayedGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LinkDelayedGet200Response) GetDataOk() (*LinkDelayedGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LinkDelayedGet200Response) SetData(v LinkDelayedGet200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *LinkDelayedGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *LinkDelayedGet200Response) GetError() LinkDelayedGet200ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *LinkDelayedGet200Response) GetErrorOk() (*LinkDelayedGet200ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *LinkDelayedGet200Response) SetError(v LinkDelayedGet200ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *LinkDelayedGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


