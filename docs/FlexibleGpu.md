# FlexibleGpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | If &#x60;true&#x60;, the fGPU is deleted when the VM is terminated. | [optional] 
**FlexibleGpuId** | Pointer to **string** | The ID of the fGPU. | [optional] 
**ModelName** | Pointer to **string** | The model of fGPU. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs). | [optional] 
**State** | Pointer to **string** | The state of the fGPU (&#x60;allocated&#x60; \\| &#x60;attaching&#x60; \\| &#x60;attached&#x60; \\| &#x60;detaching&#x60;). | [optional] 
**SubregionName** | Pointer to **string** | The Subregion where the fGPU is located. | [optional] 
**VmId** | Pointer to **string** | The ID of the VM the fGPU is attached to, if any. | [optional] 

## Methods

### GetDeleteOnVmDeletion

`func (o *FlexibleGpu) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *FlexibleGpu) GetDeleteOnVmDeletionOk() (bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeleteOnVmDeletion

`func (o *FlexibleGpu) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### SetDeleteOnVmDeletion

`func (o *FlexibleGpu) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.

### GetFlexibleGpuId

`func (o *FlexibleGpu) GetFlexibleGpuId() string`

GetFlexibleGpuId returns the FlexibleGpuId field if non-nil, zero value otherwise.

### GetFlexibleGpuIdOk

`func (o *FlexibleGpu) GetFlexibleGpuIdOk() (string, bool)`

GetFlexibleGpuIdOk returns a tuple with the FlexibleGpuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFlexibleGpuId

`func (o *FlexibleGpu) HasFlexibleGpuId() bool`

HasFlexibleGpuId returns a boolean if a field has been set.

### SetFlexibleGpuId

`func (o *FlexibleGpu) SetFlexibleGpuId(v string)`

SetFlexibleGpuId gets a reference to the given string and assigns it to the FlexibleGpuId field.

### GetModelName

`func (o *FlexibleGpu) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *FlexibleGpu) GetModelNameOk() (string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModelName

`func (o *FlexibleGpu) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### SetModelName

`func (o *FlexibleGpu) SetModelName(v string)`

SetModelName gets a reference to the given string and assigns it to the ModelName field.

### GetState

`func (o *FlexibleGpu) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FlexibleGpu) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *FlexibleGpu) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *FlexibleGpu) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetSubregionName

`func (o *FlexibleGpu) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *FlexibleGpu) GetSubregionNameOk() (string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubregionName

`func (o *FlexibleGpu) HasSubregionName() bool`

HasSubregionName returns a boolean if a field has been set.

### SetSubregionName

`func (o *FlexibleGpu) SetSubregionName(v string)`

SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.

### GetVmId

`func (o *FlexibleGpu) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *FlexibleGpu) GetVmIdOk() (string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmId

`func (o *FlexibleGpu) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### SetVmId

`func (o *FlexibleGpu) SetVmId(v string)`

SetVmId gets a reference to the given string and assigns it to the VmId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


