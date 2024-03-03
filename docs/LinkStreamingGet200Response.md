# LinkStreamingGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**LinkStreamingGet200ResponseData**](LinkStreamingGet200ResponseData.md) |  | [optional] 
**Error** | Pointer to [**LinkUnlockGet200ResponseError**](LinkUnlockGet200ResponseError.md) |  | [optional] 

## Methods

### NewLinkStreamingGet200Response

`func NewLinkStreamingGet200Response() *LinkStreamingGet200Response`

NewLinkStreamingGet200Response instantiates a new LinkStreamingGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkStreamingGet200ResponseWithDefaults

`func NewLinkStreamingGet200ResponseWithDefaults() *LinkStreamingGet200Response`

NewLinkStreamingGet200ResponseWithDefaults instantiates a new LinkStreamingGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *LinkStreamingGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LinkStreamingGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LinkStreamingGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LinkStreamingGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *LinkStreamingGet200Response) GetData() LinkStreamingGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LinkStreamingGet200Response) GetDataOk() (*LinkStreamingGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LinkStreamingGet200Response) SetData(v LinkStreamingGet200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *LinkStreamingGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *LinkStreamingGet200Response) GetError() LinkUnlockGet200ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *LinkStreamingGet200Response) GetErrorOk() (*LinkUnlockGet200ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *LinkStreamingGet200Response) SetError(v LinkUnlockGet200ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *LinkStreamingGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


