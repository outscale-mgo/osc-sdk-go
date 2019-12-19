# FlexibleGpuCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Generations** | Pointer to **string** | The generations of VMs that the fGPU is compatible with. | [optional] 
**MaxCpu** | Pointer to **int64** | The maximum number of VM vCores that the fGPU is compatible with. | [optional] 
**MaxRam** | Pointer to **int64** | The maximum amount of VM memory that the fGPU is compatible with. | [optional] 
**ModelName** | Pointer to **string** | The model of fGPU. | [optional] 
**VRam** | Pointer to **int64** | The amount of video RAM (VRAM) of the fGPU. | [optional] 

## Methods

### GetGenerations

`func (o *FlexibleGpuCatalog) GetGenerations() string`

GetGenerations returns the Generations field if non-nil, zero value otherwise.

### GetGenerationsOk

`func (o *FlexibleGpuCatalog) GetGenerationsOk() (string, bool)`

GetGenerationsOk returns a tuple with the Generations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGenerations

`func (o *FlexibleGpuCatalog) HasGenerations() bool`

HasGenerations returns a boolean if a field has been set.

### SetGenerations

`func (o *FlexibleGpuCatalog) SetGenerations(v string)`

SetGenerations gets a reference to the given string and assigns it to the Generations field.

### GetMaxCpu

`func (o *FlexibleGpuCatalog) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *FlexibleGpuCatalog) GetMaxCpuOk() (int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxCpu

`func (o *FlexibleGpuCatalog) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### SetMaxCpu

`func (o *FlexibleGpuCatalog) SetMaxCpu(v int64)`

SetMaxCpu gets a reference to the given int64 and assigns it to the MaxCpu field.

### GetMaxRam

`func (o *FlexibleGpuCatalog) GetMaxRam() int64`

GetMaxRam returns the MaxRam field if non-nil, zero value otherwise.

### GetMaxRamOk

`func (o *FlexibleGpuCatalog) GetMaxRamOk() (int64, bool)`

GetMaxRamOk returns a tuple with the MaxRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxRam

`func (o *FlexibleGpuCatalog) HasMaxRam() bool`

HasMaxRam returns a boolean if a field has been set.

### SetMaxRam

`func (o *FlexibleGpuCatalog) SetMaxRam(v int64)`

SetMaxRam gets a reference to the given int64 and assigns it to the MaxRam field.

### GetModelName

`func (o *FlexibleGpuCatalog) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *FlexibleGpuCatalog) GetModelNameOk() (string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModelName

`func (o *FlexibleGpuCatalog) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### SetModelName

`func (o *FlexibleGpuCatalog) SetModelName(v string)`

SetModelName gets a reference to the given string and assigns it to the ModelName field.

### GetVRam

`func (o *FlexibleGpuCatalog) GetVRam() int64`

GetVRam returns the VRam field if non-nil, zero value otherwise.

### GetVRamOk

`func (o *FlexibleGpuCatalog) GetVRamOk() (int64, bool)`

GetVRamOk returns a tuple with the VRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVRam

`func (o *FlexibleGpuCatalog) HasVRam() bool`

HasVRam returns a boolean if a field has been set.

### SetVRam

`func (o *FlexibleGpuCatalog) SetVRam(v int64)`

SetVRam gets a reference to the given int64 and assigns it to the VRam field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


