# FiltersFlexibleGpu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | Indicates whether the fGPU is deleted when terminating the VM. | [optional] 
**FlexibleGpuIds** | Pointer to **[]string** | One or more IDs of fGPUs. | [optional] 
**ModelNames** | Pointer to **[]string** | One or more models of fGPUs. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs). | [optional] 
**States** | Pointer to **[]string** | The states of the fGPUs (&#x60;allocated&#x60; \\| &#x60;attaching&#x60; \\| &#x60;attached&#x60; \\| &#x60;detaching&#x60;). | [optional] 
**SubregionNames** | Pointer to **[]string** | The Subregions where the fGPUs are located. | [optional] 
**VmIds** | Pointer to **[]string** | One or more IDs of VMs. | [optional] 

## Methods

### GetDeleteOnVmDeletion

`func (o *FiltersFlexibleGpu) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *FiltersFlexibleGpu) GetDeleteOnVmDeletionOk() (bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeleteOnVmDeletion

`func (o *FiltersFlexibleGpu) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### SetDeleteOnVmDeletion

`func (o *FiltersFlexibleGpu) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.

### GetFlexibleGpuIds

`func (o *FiltersFlexibleGpu) GetFlexibleGpuIds() []string`

GetFlexibleGpuIds returns the FlexibleGpuIds field if non-nil, zero value otherwise.

### GetFlexibleGpuIdsOk

`func (o *FiltersFlexibleGpu) GetFlexibleGpuIdsOk() ([]string, bool)`

GetFlexibleGpuIdsOk returns a tuple with the FlexibleGpuIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFlexibleGpuIds

`func (o *FiltersFlexibleGpu) HasFlexibleGpuIds() bool`

HasFlexibleGpuIds returns a boolean if a field has been set.

### SetFlexibleGpuIds

`func (o *FiltersFlexibleGpu) SetFlexibleGpuIds(v []string)`

SetFlexibleGpuIds gets a reference to the given []string and assigns it to the FlexibleGpuIds field.

### GetModelNames

`func (o *FiltersFlexibleGpu) GetModelNames() []string`

GetModelNames returns the ModelNames field if non-nil, zero value otherwise.

### GetModelNamesOk

`func (o *FiltersFlexibleGpu) GetModelNamesOk() ([]string, bool)`

GetModelNamesOk returns a tuple with the ModelNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModelNames

`func (o *FiltersFlexibleGpu) HasModelNames() bool`

HasModelNames returns a boolean if a field has been set.

### SetModelNames

`func (o *FiltersFlexibleGpu) SetModelNames(v []string)`

SetModelNames gets a reference to the given []string and assigns it to the ModelNames field.

### GetStates

`func (o *FiltersFlexibleGpu) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersFlexibleGpu) GetStatesOk() ([]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStates

`func (o *FiltersFlexibleGpu) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStates

`func (o *FiltersFlexibleGpu) SetStates(v []string)`

SetStates gets a reference to the given []string and assigns it to the States field.

### GetSubregionNames

`func (o *FiltersFlexibleGpu) GetSubregionNames() []string`

GetSubregionNames returns the SubregionNames field if non-nil, zero value otherwise.

### GetSubregionNamesOk

`func (o *FiltersFlexibleGpu) GetSubregionNamesOk() ([]string, bool)`

GetSubregionNamesOk returns a tuple with the SubregionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubregionNames

`func (o *FiltersFlexibleGpu) HasSubregionNames() bool`

HasSubregionNames returns a boolean if a field has been set.

### SetSubregionNames

`func (o *FiltersFlexibleGpu) SetSubregionNames(v []string)`

SetSubregionNames gets a reference to the given []string and assigns it to the SubregionNames field.

### GetVmIds

`func (o *FiltersFlexibleGpu) GetVmIds() []string`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *FiltersFlexibleGpu) GetVmIdsOk() ([]string, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmIds

`func (o *FiltersFlexibleGpu) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.

### SetVmIds

`func (o *FiltersFlexibleGpu) SetVmIds(v []string)`

SetVmIds gets a reference to the given []string and assigns it to the VmIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


