# LinkUnlockGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**LinkUnlockGet200ResponseData**](LinkUnlockGet200ResponseData.md) |  | [optional] 
**Error** | Pointer to [**LinkUnlockGet200ResponseError**](LinkUnlockGet200ResponseError.md) |  | [optional] 

## Methods

### NewLinkUnlockGet200Response

`func NewLinkUnlockGet200Response() *LinkUnlockGet200Response`

NewLinkUnlockGet200Response instantiates a new LinkUnlockGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkUnlockGet200ResponseWithDefaults

`func NewLinkUnlockGet200ResponseWithDefaults() *LinkUnlockGet200Response`

NewLinkUnlockGet200ResponseWithDefaults instantiates a new LinkUnlockGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *LinkUnlockGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LinkUnlockGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LinkUnlockGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LinkUnlockGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *LinkUnlockGet200Response) GetData() LinkUnlockGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LinkUnlockGet200Response) GetDataOk() (*LinkUnlockGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LinkUnlockGet200Response) SetData(v LinkUnlockGet200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *LinkUnlockGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *LinkUnlockGet200Response) GetError() LinkUnlockGet200ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *LinkUnlockGet200Response) GetErrorOk() (*LinkUnlockGet200ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *LinkUnlockGet200Response) SetError(v LinkUnlockGet200ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *LinkUnlockGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


