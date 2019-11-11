# ReadImageExportTasksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageExportTasks** | Pointer to [**[]ImageExportTask**](ImageExportTask.md) | Information about one or more image export tasks. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetImageExportTasks

`func (o *ReadImageExportTasksResponse) GetImageExportTasks() []ImageExportTask`

GetImageExportTasks returns the ImageExportTasks field if non-nil, zero value otherwise.

### GetImageExportTasksOk

`func (o *ReadImageExportTasksResponse) GetImageExportTasksOk() ([]ImageExportTask, bool)`

GetImageExportTasksOk returns a tuple with the ImageExportTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImageExportTasks

`func (o *ReadImageExportTasksResponse) HasImageExportTasks() bool`

HasImageExportTasks returns a boolean if a field has been set.

### SetImageExportTasks

`func (o *ReadImageExportTasksResponse) SetImageExportTasks(v []ImageExportTask)`

SetImageExportTasks gets a reference to the given []ImageExportTask and assigns it to the ImageExportTasks field.

### GetResponseContext

`func (o *ReadImageExportTasksResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadImageExportTasksResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadImageExportTasksResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadImageExportTasksResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


