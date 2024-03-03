# MagnetUploadFilePost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**MagnetUploadFilePost200ResponseData**](MagnetUploadFilePost200ResponseData.md) |  | [optional] 
**Error** | Pointer to [**MagnetUploadGet200ResponseError**](MagnetUploadGet200ResponseError.md) |  | [optional] 

## Methods

### NewMagnetUploadFilePost200Response

`func NewMagnetUploadFilePost200Response() *MagnetUploadFilePost200Response`

NewMagnetUploadFilePost200Response instantiates a new MagnetUploadFilePost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMagnetUploadFilePost200ResponseWithDefaults

`func NewMagnetUploadFilePost200ResponseWithDefaults() *MagnetUploadFilePost200Response`

NewMagnetUploadFilePost200ResponseWithDefaults instantiates a new MagnetUploadFilePost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MagnetUploadFilePost200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MagnetUploadFilePost200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MagnetUploadFilePost200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MagnetUploadFilePost200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *MagnetUploadFilePost200Response) GetData() MagnetUploadFilePost200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MagnetUploadFilePost200Response) GetDataOk() (*MagnetUploadFilePost200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MagnetUploadFilePost200Response) SetData(v MagnetUploadFilePost200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *MagnetUploadFilePost200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *MagnetUploadFilePost200Response) GetError() MagnetUploadGet200ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MagnetUploadFilePost200Response) GetErrorOk() (*MagnetUploadGet200ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MagnetUploadFilePost200Response) SetError(v MagnetUploadGet200ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *MagnetUploadFilePost200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


