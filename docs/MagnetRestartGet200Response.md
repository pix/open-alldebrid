# MagnetRestartGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**MagnetRestartGet200ResponseData**](MagnetRestartGet200ResponseData.md) |  | [optional] 
**Error** | Pointer to [**MagnetDeleteGet200ResponseError**](MagnetDeleteGet200ResponseError.md) |  | [optional] 

## Methods

### NewMagnetRestartGet200Response

`func NewMagnetRestartGet200Response() *MagnetRestartGet200Response`

NewMagnetRestartGet200Response instantiates a new MagnetRestartGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMagnetRestartGet200ResponseWithDefaults

`func NewMagnetRestartGet200ResponseWithDefaults() *MagnetRestartGet200Response`

NewMagnetRestartGet200ResponseWithDefaults instantiates a new MagnetRestartGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MagnetRestartGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MagnetRestartGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MagnetRestartGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MagnetRestartGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *MagnetRestartGet200Response) GetData() MagnetRestartGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MagnetRestartGet200Response) GetDataOk() (*MagnetRestartGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MagnetRestartGet200Response) SetData(v MagnetRestartGet200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *MagnetRestartGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *MagnetRestartGet200Response) GetError() MagnetDeleteGet200ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MagnetRestartGet200Response) GetErrorOk() (*MagnetDeleteGet200ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MagnetRestartGet200Response) SetError(v MagnetDeleteGet200ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *MagnetRestartGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


