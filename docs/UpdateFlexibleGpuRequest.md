# UpdateFlexibleGpuRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | If &#x60;true&#x60;, the fGPU is deleted when the VM is terminated. | [optional] 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**FlexibleGpuId** | Pointer to **string** | The ID of the fGPU you want to modify. | 

## Methods

### GetDeleteOnVmDeletion

`func (o *UpdateFlexibleGpuRequest) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *UpdateFlexibleGpuRequest) GetDeleteOnVmDeletionOk() (bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeleteOnVmDeletion

`func (o *UpdateFlexibleGpuRequest) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### SetDeleteOnVmDeletion

`func (o *UpdateFlexibleGpuRequest) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.

### GetDryRun

`func (o *UpdateFlexibleGpuRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateFlexibleGpuRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UpdateFlexibleGpuRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UpdateFlexibleGpuRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetFlexibleGpuId

`func (o *UpdateFlexibleGpuRequest) GetFlexibleGpuId() string`

GetFlexibleGpuId returns the FlexibleGpuId field if non-nil, zero value otherwise.

### GetFlexibleGpuIdOk

`func (o *UpdateFlexibleGpuRequest) GetFlexibleGpuIdOk() (string, bool)`

GetFlexibleGpuIdOk returns a tuple with the FlexibleGpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFlexibleGpuId

`func (o *UpdateFlexibleGpuRequest) HasFlexibleGpuId() bool`

HasFlexibleGpuId returns a boolean if a field has been set.

### SetFlexibleGpuId

`func (o *UpdateFlexibleGpuRequest) SetFlexibleGpuId(v string)`

SetFlexibleGpuId gets a reference to the given string and assigns it to the FlexibleGpuId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


