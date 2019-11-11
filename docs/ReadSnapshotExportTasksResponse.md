# ReadSnapshotExportTasksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**SnapshotExportTasks** | Pointer to [**[]SnapshotExportTask**](SnapshotExportTask.md) | Information about one or more snapshot export tasks. | [optional] 

## Methods

### GetResponseContext

`func (o *ReadSnapshotExportTasksResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadSnapshotExportTasksResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadSnapshotExportTasksResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadSnapshotExportTasksResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.

### GetSnapshotExportTasks

`func (o *ReadSnapshotExportTasksResponse) GetSnapshotExportTasks() []SnapshotExportTask`

GetSnapshotExportTasks returns the SnapshotExportTasks field if non-nil, zero value otherwise.

### GetSnapshotExportTasksOk

`func (o *ReadSnapshotExportTasksResponse) GetSnapshotExportTasksOk() ([]SnapshotExportTask, bool)`

GetSnapshotExportTasksOk returns a tuple with the SnapshotExportTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSnapshotExportTasks

`func (o *ReadSnapshotExportTasksResponse) HasSnapshotExportTasks() bool`

HasSnapshotExportTasks returns a boolean if a field has been set.

### SetSnapshotExportTasks

`func (o *ReadSnapshotExportTasksResponse) SetSnapshotExportTasks(v []SnapshotExportTask)`

SetSnapshotExportTasks gets a reference to the given []SnapshotExportTask and assigns it to the SnapshotExportTasks field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


