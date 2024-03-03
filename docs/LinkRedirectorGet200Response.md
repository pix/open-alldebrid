# LinkRedirectorGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**LinkRedirectorGet200ResponseData**](LinkRedirectorGet200ResponseData.md) |  | [optional] 
**Error** | Pointer to [**LinkRedirectorGet200ResponseError**](LinkRedirectorGet200ResponseError.md) |  | [optional] 

## Methods

### NewLinkRedirectorGet200Response

`func NewLinkRedirectorGet200Response() *LinkRedirectorGet200Response`

NewLinkRedirectorGet200Response instantiates a new LinkRedirectorGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkRedirectorGet200ResponseWithDefaults

`func NewLinkRedirectorGet200ResponseWithDefaults() *LinkRedirectorGet200Response`

NewLinkRedirectorGet200ResponseWithDefaults instantiates a new LinkRedirectorGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *LinkRedirectorGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LinkRedirectorGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LinkRedirectorGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LinkRedirectorGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *LinkRedirectorGet200Response) GetData() LinkRedirectorGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LinkRedirectorGet200Response) GetDataOk() (*LinkRedirectorGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LinkRedirectorGet200Response) SetData(v LinkRedirectorGet200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *LinkRedirectorGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *LinkRedirectorGet200Response) GetError() LinkRedirectorGet200ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *LinkRedirectorGet200Response) GetErrorOk() (*LinkRedirectorGet200ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *LinkRedirectorGet200Response) SetError(v LinkRedirectorGet200ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *LinkRedirectorGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


