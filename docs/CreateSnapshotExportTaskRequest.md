# CreateSnapshotExportTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**OsuExport** | Pointer to [**OsuExport**](OsuExport.md) |  | 
**SnapshotId** | Pointer to **string** | The ID of the snapshot to export. | 

## Methods

### GetDryRun

`func (o *CreateSnapshotExportTaskRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateSnapshotExportTaskRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateSnapshotExportTaskRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateSnapshotExportTaskRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetOsuExport

`func (o *CreateSnapshotExportTaskRequest) GetOsuExport() OsuExport`

GetOsuExport returns the OsuExport field if non-nil, zero value otherwise.

### GetOsuExportOk

`func (o *CreateSnapshotExportTaskRequest) GetOsuExportOk() (OsuExport, bool)`

GetOsuExportOk returns a tuple with the OsuExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsuExport

`func (o *CreateSnapshotExportTaskRequest) HasOsuExport() bool`

HasOsuExport returns a boolean if a field has been set.

### SetOsuExport

`func (o *CreateSnapshotExportTaskRequest) SetOsuExport(v OsuExport)`

SetOsuExport gets a reference to the given OsuExport and assigns it to the OsuExport field.

### GetSnapshotId

`func (o *CreateSnapshotExportTaskRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *CreateSnapshotExportTaskRequest) GetSnapshotIdOk() (string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSnapshotId

`func (o *CreateSnapshotExportTaskRequest) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotId

`func (o *CreateSnapshotExportTaskRequest) SetSnapshotId(v string)`

SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


