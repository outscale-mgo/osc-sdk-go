# FiltersVmType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BsuOptimized** | Pointer to **bool** | Indicates whether the VM is optimized for BSU I/O. | [optional] 
**MemorySizes** | Pointer to **[]float32** | The amounts of memory, in gibibytes (GiB). | [optional] 
**VcoreCounts** | Pointer to **[]int64** | The numbers of vCores. | [optional] 
**VmTypeNames** | Pointer to **[]string** | The names of the VM types. For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types). | [optional] 
**VolumeCounts** | Pointer to **[]int64** | The maximum number of ephemeral storage disks. | [optional] 
**VolumeSizes** | Pointer to **[]int64** | The size of one ephemeral storage disk, in gibibytes (GiB). | [optional] 

## Methods

### GetBsuOptimized

`func (o *FiltersVmType) GetBsuOptimized() bool`

GetBsuOptimized returns the BsuOptimized field if non-nil, zero value otherwise.

### GetBsuOptimizedOk

`func (o *FiltersVmType) GetBsuOptimizedOk() (bool, bool)`

GetBsuOptimizedOk returns a tuple with the BsuOptimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBsuOptimized

`func (o *FiltersVmType) HasBsuOptimized() bool`

HasBsuOptimized returns a boolean if a field has been set.

### SetBsuOptimized

`func (o *FiltersVmType) SetBsuOptimized(v bool)`

SetBsuOptimized gets a reference to the given bool and assigns it to the BsuOptimized field.

### GetMemorySizes

`func (o *FiltersVmType) GetMemorySizes() []float32`

GetMemorySizes returns the MemorySizes field if non-nil, zero value otherwise.

### GetMemorySizesOk

`func (o *FiltersVmType) GetMemorySizesOk() ([]float32, bool)`

GetMemorySizesOk returns a tuple with the MemorySizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMemorySizes

`func (o *FiltersVmType) HasMemorySizes() bool`

HasMemorySizes returns a boolean if a field has been set.

### SetMemorySizes

`func (o *FiltersVmType) SetMemorySizes(v []float32)`

SetMemorySizes gets a reference to the given []float32 and assigns it to the MemorySizes field.

### GetVcoreCounts

`func (o *FiltersVmType) GetVcoreCounts() []int64`

GetVcoreCounts returns the VcoreCounts field if non-nil, zero value otherwise.

### GetVcoreCountsOk

`func (o *FiltersVmType) GetVcoreCountsOk() ([]int64, bool)`

GetVcoreCountsOk returns a tuple with the VcoreCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVcoreCounts

`func (o *FiltersVmType) HasVcoreCounts() bool`

HasVcoreCounts returns a boolean if a field has been set.

### SetVcoreCounts

`func (o *FiltersVmType) SetVcoreCounts(v []int64)`

SetVcoreCounts gets a reference to the given []int64 and assigns it to the VcoreCounts field.

### GetVmTypeNames

`func (o *FiltersVmType) GetVmTypeNames() []string`

GetVmTypeNames returns the VmTypeNames field if non-nil, zero value otherwise.

### GetVmTypeNamesOk

`func (o *FiltersVmType) GetVmTypeNamesOk() ([]string, bool)`

GetVmTypeNamesOk returns a tuple with the VmTypeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmTypeNames

`func (o *FiltersVmType) HasVmTypeNames() bool`

HasVmTypeNames returns a boolean if a field has been set.

### SetVmTypeNames

`func (o *FiltersVmType) SetVmTypeNames(v []string)`

SetVmTypeNames gets a reference to the given []string and assigns it to the VmTypeNames field.

### GetVolumeCounts

`func (o *FiltersVmType) GetVolumeCounts() []int64`

GetVolumeCounts returns the VolumeCounts field if non-nil, zero value otherwise.

### GetVolumeCountsOk

`func (o *FiltersVmType) GetVolumeCountsOk() ([]int64, bool)`

GetVolumeCountsOk returns a tuple with the VolumeCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeCounts

`func (o *FiltersVmType) HasVolumeCounts() bool`

HasVolumeCounts returns a boolean if a field has been set.

### SetVolumeCounts

`func (o *FiltersVmType) SetVolumeCounts(v []int64)`

SetVolumeCounts gets a reference to the given []int64 and assigns it to the VolumeCounts field.

### GetVolumeSizes

`func (o *FiltersVmType) GetVolumeSizes() []int64`

GetVolumeSizes returns the VolumeSizes field if non-nil, zero value otherwise.

### GetVolumeSizesOk

`func (o *FiltersVmType) GetVolumeSizesOk() ([]int64, bool)`

GetVolumeSizesOk returns a tuple with the VolumeSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeSizes

`func (o *FiltersVmType) HasVolumeSizes() bool`

HasVolumeSizes returns a boolean if a field has been set.

### SetVolumeSizes

`func (o *FiltersVmType) SetVolumeSizes(v []int64)`

SetVolumeSizes gets a reference to the given []int64 and assigns it to the VolumeSizes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


