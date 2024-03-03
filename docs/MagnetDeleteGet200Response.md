# MagnetDeleteGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**MagnetDeleteGet200ResponseData**](MagnetDeleteGet200ResponseData.md) |  | [optional] 
**Error** | Pointer to [**MagnetDeleteGet200ResponseError**](MagnetDeleteGet200ResponseError.md) |  | [optional] 

## Methods

### NewMagnetDeleteGet200Response

`func NewMagnetDeleteGet200Response() *MagnetDeleteGet200Response`

NewMagnetDeleteGet200Response instantiates a new MagnetDeleteGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMagnetDeleteGet200ResponseWithDefaults

`func NewMagnetDeleteGet200ResponseWithDefaults() *MagnetDeleteGet200Response`

NewMagnetDeleteGet200ResponseWithDefaults instantiates a new MagnetDeleteGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MagnetDeleteGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MagnetDeleteGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MagnetDeleteGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MagnetDeleteGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *MagnetDeleteGet200Response) GetData() MagnetDeleteGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MagnetDeleteGet200Response) GetDataOk() (*MagnetDeleteGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MagnetDeleteGet200Response) SetData(v MagnetDeleteGet200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *MagnetDeleteGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *MagnetDeleteGet200Response) GetError() MagnetDeleteGet200ResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MagnetDeleteGet200Response) GetErrorOk() (*MagnetDeleteGet200ResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MagnetDeleteGet200Response) SetError(v MagnetDeleteGet200ResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *MagnetDeleteGet200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


