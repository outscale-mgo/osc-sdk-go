# ImageExportTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | If the OMI export task fails, an error message appears. | [optional] 
**ImageId** | Pointer to **string** | The ID of the OMI to be exported. | [optional] 
**OsuExport** | Pointer to [**OsuExport**](OsuExport.md) |  | [optional] 
**Progress** | Pointer to **int64** | The progress of the OMI export task, as a percentage. | [optional] 
**State** | Pointer to **string** | The state of the OMI export task (&#x60;pending/queued&#x60; \\| &#x60;pending&#x60; \\| &#x60;completed&#x60; \\| &#x60;failed&#x60; \\| &#x60;cancelled&#x60;). | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the image export task. | [optional] 
**TaskId** | Pointer to **string** | The ID of the OMI export task. | [optional] 

## Methods

### GetComment

`func (o *ImageExportTask) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ImageExportTask) GetCommentOk() (string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComment

`func (o *ImageExportTask) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetComment

`func (o *ImageExportTask) SetComment(v string)`

SetComment gets a reference to the given string and assigns it to the Comment field.

### GetImageId

`func (o *ImageExportTask) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ImageExportTask) GetImageIdOk() (string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImageId

`func (o *ImageExportTask) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageId

`func (o *ImageExportTask) SetImageId(v string)`

SetImageId gets a reference to the given string and assigns it to the ImageId field.

### GetOsuExport

`func (o *ImageExportTask) GetOsuExport() OsuExport`

GetOsuExport returns the OsuExport field if non-nil, zero value otherwise.

### GetOsuExportOk

`func (o *ImageExportTask) GetOsuExportOk() (OsuExport, bool)`

GetOsuExportOk returns a tuple with the OsuExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsuExport

`func (o *ImageExportTask) HasOsuExport() bool`

HasOsuExport returns a boolean if a field has been set.

### SetOsuExport

`func (o *ImageExportTask) SetOsuExport(v OsuExport)`

SetOsuExport gets a reference to the given OsuExport and assigns it to the OsuExport field.

### GetProgress

`func (o *ImageExportTask) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ImageExportTask) GetProgressOk() (int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgress

`func (o *ImageExportTask) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgress

`func (o *ImageExportTask) SetProgress(v int64)`

SetProgress gets a reference to the given int64 and assigns it to the Progress field.

### GetState

`func (o *ImageExportTask) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ImageExportTask) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *ImageExportTask) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *ImageExportTask) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetTags

`func (o *ImageExportTask) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImageExportTask) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *ImageExportTask) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *ImageExportTask) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.

### GetTaskId

`func (o *ImageExportTask) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ImageExportTask) GetTaskIdOk() (string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTaskId

`func (o *ImageExportTask) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskId

`func (o *ImageExportTask) SetTaskId(v string)`

SetTaskId gets a reference to the given string and assigns it to the TaskId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


