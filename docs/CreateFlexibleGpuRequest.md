# CreateFlexibleGpuRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | If &#x60;true&#x60;, the fGPU is deleted when the VM is terminated. | [optional] [default to false]
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**ModelName** | Pointer to **string** | The model of GPU you want to allocate. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs). | 
**SubregionName** | Pointer to **string** | The Subregion in which you want to create the fGPU. | 

## Methods

### GetDeleteOnVmDeletion

`func (o *CreateFlexibleGpuRequest) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *CreateFlexibleGpuRequest) GetDeleteOnVmDeletionOk() (bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeleteOnVmDeletion

`func (o *CreateFlexibleGpuRequest) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### SetDeleteOnVmDeletion

`func (o *CreateFlexibleGpuRequest) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.

### GetDryRun

`func (o *CreateFlexibleGpuRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateFlexibleGpuRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateFlexibleGpuRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateFlexibleGpuRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetModelName

`func (o *CreateFlexibleGpuRequest) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *CreateFlexibleGpuRequest) GetModelNameOk() (string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModelName

`func (o *CreateFlexibleGpuRequest) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### SetModelName

`func (o *CreateFlexibleGpuRequest) SetModelName(v string)`

SetModelName gets a reference to the given string and assigns it to the ModelName field.

### GetSubregionName

`func (o *CreateFlexibleGpuRequest) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *CreateFlexibleGpuRequest) GetSubregionNameOk() (string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubregionName

`func (o *CreateFlexibleGpuRequest) HasSubregionName() bool`

HasSubregionName returns a boolean if a field has been set.

### SetSubregionName

`func (o *CreateFlexibleGpuRequest) SetSubregionName(v string)`

SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


