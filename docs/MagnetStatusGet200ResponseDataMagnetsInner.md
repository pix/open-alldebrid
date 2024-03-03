# MagnetStatusGet200ResponseDataMagnetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int64** |  | [optional] 
**Downloaded** | Pointer to **int64** |  | [optional] 
**Uploaded** | Pointer to **int64** |  | [optional] 
**Seeders** | Pointer to **int64** |  | [optional] 
**DownloadSpeed** | Pointer to **int64** |  | [optional] 
**UploadSpeed** | Pointer to **int64** |  | [optional] 
**UploadDate** | Pointer to **int64** |  | [optional] 
**CompletionDate** | Pointer to **int64** |  | [optional] 
**Links** | Pointer to [**[]MagnetStatusGet200ResponseDataMagnetsInnerLinksInner**](MagnetStatusGet200ResponseDataMagnetsInnerLinksInner.md) |  | [optional] 

## Methods

### NewMagnetStatusGet200ResponseDataMagnetsInner

`func NewMagnetStatusGet200ResponseDataMagnetsInner() *MagnetStatusGet200ResponseDataMagnetsInner`

NewMagnetStatusGet200ResponseDataMagnetsInner instantiates a new MagnetStatusGet200ResponseDataMagnetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMagnetStatusGet200ResponseDataMagnetsInnerWithDefaults

`func NewMagnetStatusGet200ResponseDataMagnetsInnerWithDefaults() *MagnetStatusGet200ResponseDataMagnetsInner`

NewMagnetStatusGet200ResponseDataMagnetsInnerWithDefaults instantiates a new MagnetStatusGet200ResponseDataMagnetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFilename

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetSize

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetDownloaded

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetDownloaded() int64`

GetDownloaded returns the Downloaded field if non-nil, zero value otherwise.

### GetDownloadedOk

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetDownloadedOk() (*int64, bool)`

GetDownloadedOk returns a tuple with the Downloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloaded

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) SetDownloaded(v int64)`

SetDownloaded sets Downloaded field to given value.

### HasDownloaded

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) HasDownloaded() bool`

HasDownloaded returns a boolean if a field has been set.

### GetUploaded

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetUploaded() int64`

GetUploaded returns the Uploaded field if non-nil, zero value otherwise.

### GetUploadedOk

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetUploadedOk() (*int64, bool)`

GetUploadedOk returns a tuple with the Uploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaded

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) SetUploaded(v int64)`

SetUploaded sets Uploaded field to given value.

### HasUploaded

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) HasUploaded() bool`

HasUploaded returns a boolean if a field has been set.

### GetSeeders

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetSeeders() int64`

GetSeeders returns the Seeders field if non-nil, zero value otherwise.

### GetSeedersOk

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetSeedersOk() (*int64, bool)`

GetSeedersOk returns a tuple with the Seeders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeders

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) SetSeeders(v int64)`

SetSeeders sets Seeders field to given value.

### HasSeeders

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) HasSeeders() bool`

HasSeeders returns a boolean if a field has been set.

### GetDownloadSpeed

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetDownloadSpeed() int64`

GetDownloadSpeed returns the DownloadSpeed field if non-nil, zero value otherwise.

### GetDownloadSpeedOk

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetDownloadSpeedOk() (*int64, bool)`

GetDownloadSpeedOk returns a tuple with the DownloadSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadSpeed

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) SetDownloadSpeed(v int64)`

SetDownloadSpeed sets DownloadSpeed field to given value.

### HasDownloadSpeed

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) HasDownloadSpeed() bool`

HasDownloadSpeed returns a boolean if a field has been set.

### GetUploadSpeed

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetUploadSpeed() int64`

GetUploadSpeed returns the UploadSpeed field if non-nil, zero value otherwise.

### GetUploadSpeedOk

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetUploadSpeedOk() (*int64, bool)`

GetUploadSpeedOk returns a tuple with the UploadSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadSpeed

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) SetUploadSpeed(v int64)`

SetUploadSpeed sets UploadSpeed field to given value.

### HasUploadSpeed

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) HasUploadSpeed() bool`

HasUploadSpeed returns a boolean if a field has been set.

### GetUploadDate

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetUploadDate() int64`

GetUploadDate returns the UploadDate field if non-nil, zero value otherwise.

### GetUploadDateOk

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetUploadDateOk() (*int64, bool)`

GetUploadDateOk returns a tuple with the UploadDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadDate

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) SetUploadDate(v int64)`

SetUploadDate sets UploadDate field to given value.

### HasUploadDate

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) HasUploadDate() bool`

HasUploadDate returns a boolean if a field has been set.

### GetCompletionDate

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetCompletionDate() int64`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetCompletionDateOk() (*int64, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) SetCompletionDate(v int64)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### GetLinks

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetLinks() []MagnetStatusGet200ResponseDataMagnetsInnerLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) GetLinksOk() (*[]MagnetStatusGet200ResponseDataMagnetsInnerLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) SetLinks(v []MagnetStatusGet200ResponseDataMagnetsInnerLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MagnetStatusGet200ResponseDataMagnetsInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


