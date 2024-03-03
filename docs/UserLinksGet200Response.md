# UserLinksGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**UserLinksGet200ResponseData**](UserLinksGet200ResponseData.md) |  | [optional] 
**Error** | Pointer to [**MagnetUploadGet200ResponseError**](MagnetUploadGet200ResponseError.md) |  | [optional] 

## Methods

### NewUserLinksGet200Response

`func NewUserLinksGet200Response() *UserLinksGet200Response`

NewUserLinksGet200Response instantiates a new UserLinksGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLinksGet200ResponseWithDefaults

`func NewUserLinksGet200ResponseWithDefaults() *UserLinksGet200Response`

NewUserLinksGet200ResponseWithDefaults instantiates a new UserLinksGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UserLinksGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserLinksGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserLinksGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserLinksGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *UserLinksGet200Response) GetData() UserLinksGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserLinksGet200Response) GetDataOk() (*UserLinksGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserLinksGet200Response) SetData(v UserLinksGet200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *UserLinksGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *UserLinksGet200Response) GetError() MagnetUploadGet200ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UserLinksGet200Response) GetErrorOk() (*MagnetUploadGet200ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UserLinksGet200Response) SetError(v MagnetUploadGet200ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *UserLinksGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


