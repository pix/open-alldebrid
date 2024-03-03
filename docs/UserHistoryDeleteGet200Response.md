# UserHistoryDeleteGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**UserHistoryDeleteGet200ResponseData**](UserHistoryDeleteGet200ResponseData.md) |  | [optional] 
**Error** | Pointer to [**MagnetUploadGet200ResponseError**](MagnetUploadGet200ResponseError.md) |  | [optional] 

## Methods

### NewUserHistoryDeleteGet200Response

`func NewUserHistoryDeleteGet200Response() *UserHistoryDeleteGet200Response`

NewUserHistoryDeleteGet200Response instantiates a new UserHistoryDeleteGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserHistoryDeleteGet200ResponseWithDefaults

`func NewUserHistoryDeleteGet200ResponseWithDefaults() *UserHistoryDeleteGet200Response`

NewUserHistoryDeleteGet200ResponseWithDefaults instantiates a new UserHistoryDeleteGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UserHistoryDeleteGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserHistoryDeleteGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserHistoryDeleteGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserHistoryDeleteGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *UserHistoryDeleteGet200Response) GetData() UserHistoryDeleteGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserHistoryDeleteGet200Response) GetDataOk() (*UserHistoryDeleteGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserHistoryDeleteGet200Response) SetData(v UserHistoryDeleteGet200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *UserHistoryDeleteGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *UserHistoryDeleteGet200Response) GetError() MagnetUploadGet200ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UserHistoryDeleteGet200Response) GetErrorOk() (*MagnetUploadGet200ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UserHistoryDeleteGet200Response) SetError(v MagnetUploadGet200ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *UserHistoryDeleteGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


