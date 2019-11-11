# LinkFlexibleGpuRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**FlexibleGpuId** | Pointer to **string** | The ID of the fGPU you want to attach. | 
**VmId** | Pointer to **string** | The ID of the VM you want to attach the fGPU to. | 

## Methods

### GetDryRun

`func (o *LinkFlexibleGpuRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *LinkFlexibleGpuRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *LinkFlexibleGpuRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *LinkFlexibleGpuRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetFlexibleGpuId

`func (o *LinkFlexibleGpuRequest) GetFlexibleGpuId() string`

GetFlexibleGpuId returns the FlexibleGpuId field if non-nil, zero value otherwise.

### GetFlexibleGpuIdOk

`func (o *LinkFlexibleGpuRequest) GetFlexibleGpuIdOk() (string, bool)`

GetFlexibleGpuIdOk returns a tuple with the FlexibleGpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFlexibleGpuId

`func (o *LinkFlexibleGpuRequest) HasFlexibleGpuId() bool`

HasFlexibleGpuId returns a boolean if a field has been set.

### SetFlexibleGpuId

`func (o *LinkFlexibleGpuRequest) SetFlexibleGpuId(v string)`

SetFlexibleGpuId gets a reference to the given string and assigns it to the FlexibleGpuId field.

### GetVmId

`func (o *LinkFlexibleGpuRequest) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *LinkFlexibleGpuRequest) GetVmIdOk() (string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmId

`func (o *LinkFlexibleGpuRequest) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### SetVmId

`func (o *LinkFlexibleGpuRequest) SetVmId(v string)`

SetVmId gets a reference to the given string and assigns it to the VmId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


