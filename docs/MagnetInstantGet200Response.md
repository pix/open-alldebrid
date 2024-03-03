# MagnetInstantGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**MagnetInstantGet200ResponseData**](MagnetInstantGet200ResponseData.md) |  | [optional] 
**Error** | Pointer to [**MagnetUploadGet200ResponseError**](MagnetUploadGet200ResponseError.md) |  | [optional] 

## Methods

### NewMagnetInstantGet200Response

`func NewMagnetInstantGet200Response() *MagnetInstantGet200Response`

NewMagnetInstantGet200Response instantiates a new MagnetInstantGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMagnetInstantGet200ResponseWithDefaults

`func NewMagnetInstantGet200ResponseWithDefaults() *MagnetInstantGet200Response`

NewMagnetInstantGet200ResponseWithDefaults instantiates a new MagnetInstantGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MagnetInstantGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MagnetInstantGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MagnetInstantGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MagnetInstantGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *MagnetInstantGet200Response) GetData() MagnetInstantGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MagnetInstantGet200Response) GetDataOk() (*MagnetInstantGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MagnetInstantGet200Response) SetData(v MagnetInstantGet200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *MagnetInstantGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *MagnetInstantGet200Response) GetError() MagnetUploadGet200ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MagnetInstantGet200Response) GetErrorOk() (*MagnetUploadGet200ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MagnetInstantGet200Response) SetError(v MagnetUploadGet200ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *MagnetInstantGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


