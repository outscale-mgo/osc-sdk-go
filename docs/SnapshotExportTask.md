# SnapshotExportTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | If the snapshot export task fails, an error message appears. | [optional] 
**OsuExport** | Pointer to [**OsuExport**](OsuExport.md) |  | [optional] 
**Progress** | Pointer to **int64** | The progress of the snapshot export task, as a percentage. | [optional] 
**SnapshotId** | Pointer to **string** | The ID of the snapshot to be exported. | [optional] 
**State** | Pointer to **string** | The state of the snapshot export task (&#x60;pending&#x60; \\| &#x60;active&#x60; \\| &#x60;completed&#x60; \\| &#x60;failed&#x60;). | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the snapshot export task. | [optional] 
**TaskId** | Pointer to **string** | The ID of the snapshot export task. | [optional] 

## Methods

### GetComment

`func (o *SnapshotExportTask) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SnapshotExportTask) GetCommentOk() (string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasComment

`func (o *SnapshotExportTask) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetComment

`func (o *SnapshotExportTask) SetComment(v string)`

SetComment gets a reference to the given string and assigns it to the Comment field.

### GetOsuExport

`func (o *SnapshotExportTask) GetOsuExport() OsuExport`

GetOsuExport returns the OsuExport field if non-nil, zero value otherwise.

### GetOsuExportOk

`func (o *SnapshotExportTask) GetOsuExportOk() (OsuExport, bool)`

GetOsuExportOk returns a tuple with the OsuExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsuExport

`func (o *SnapshotExportTask) HasOsuExport() bool`

HasOsuExport returns a boolean if a field has been set.

### SetOsuExport

`func (o *SnapshotExportTask) SetOsuExport(v OsuExport)`

SetOsuExport gets a reference to the given OsuExport and assigns it to the OsuExport field.

### GetProgress

`func (o *SnapshotExportTask) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *SnapshotExportTask) GetProgressOk() (int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgress

`func (o *SnapshotExportTask) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgress

`func (o *SnapshotExportTask) SetProgress(v int64)`

SetProgress gets a reference to the given int64 and assigns it to the Progress field.

### GetSnapshotId

`func (o *SnapshotExportTask) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *SnapshotExportTask) GetSnapshotIdOk() (string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSnapshotId

`func (o *SnapshotExportTask) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotId

`func (o *SnapshotExportTask) SetSnapshotId(v string)`

SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.

### GetState

`func (o *SnapshotExportTask) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SnapshotExportTask) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *SnapshotExportTask) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *SnapshotExportTask) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetTags

`func (o *SnapshotExportTask) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SnapshotExportTask) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *SnapshotExportTask) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *SnapshotExportTask) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.

### GetTaskId

`func (o *SnapshotExportTask) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *SnapshotExportTask) GetTaskIdOk() (string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTaskId

`func (o *SnapshotExportTask) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskId

`func (o *SnapshotExportTask) SetTaskId(v string)`

SetTaskId gets a reference to the given string and assigns it to the TaskId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


