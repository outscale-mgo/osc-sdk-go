# CreateImageExportTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**ImageId** | Pointer to **string** | The ID of the OMI to export. | 
**OsuExport** | Pointer to [**OsuExport**](OsuExport.md) |  | 

## Methods

### GetDryRun

`func (o *CreateImageExportTaskRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateImageExportTaskRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateImageExportTaskRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateImageExportTaskRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetImageId

`func (o *CreateImageExportTaskRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateImageExportTaskRequest) GetImageIdOk() (string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImageId

`func (o *CreateImageExportTaskRequest) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageId

`func (o *CreateImageExportTaskRequest) SetImageId(v string)`

SetImageId gets a reference to the given string and assigns it to the ImageId field.

### GetOsuExport

`func (o *CreateImageExportTaskRequest) GetOsuExport() OsuExport`

GetOsuExport returns the OsuExport field if non-nil, zero value otherwise.

### GetOsuExportOk

`func (o *CreateImageExportTaskRequest) GetOsuExportOk() (OsuExport, bool)`

GetOsuExportOk returns a tuple with the OsuExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsuExport

`func (o *CreateImageExportTaskRequest) HasOsuExport() bool`

HasOsuExport returns a boolean if a field has been set.

### SetOsuExport

`func (o *CreateImageExportTaskRequest) SetOsuExport(v OsuExport)`

SetOsuExport gets a reference to the given OsuExport and assigns it to the OsuExport field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


