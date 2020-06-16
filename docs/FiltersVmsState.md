# FiltersVmsState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubregionNames** | Pointer to **[]string** | The names of the Subregions of the VMs. | [optional] 
**VmIds** | Pointer to **[]string** | One or more IDs of VMs. | [optional] 
**VmStates** | Pointer to **[]string** | The states of the VMs (&#x60;pending&#x60; \\| &#x60;running&#x60; \\| &#x60;stopping&#x60; \\| &#x60;stopped&#x60; \\| &#x60;shutting-down&#x60; \\| &#x60;terminated&#x60; \\| &#x60;quarantine&#x60;). | [optional] 

## Methods

### GetSubregionNames

`func (o *FiltersVmsState) GetSubregionNames() []string`

GetSubregionNames returns the SubregionNames field if non-nil, zero value otherwise.

### GetSubregionNamesOk

`func (o *FiltersVmsState) GetSubregionNamesOk() ([]string, bool)`

GetSubregionNamesOk returns a tuple with the SubregionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubregionNames

`func (o *FiltersVmsState) HasSubregionNames() bool`

HasSubregionNames returns a boolean if a field has been set.

### SetSubregionNames

`func (o *FiltersVmsState) SetSubregionNames(v []string)`

SetSubregionNames gets a reference to the given []string and assigns it to the SubregionNames field.

### GetVmIds

`func (o *FiltersVmsState) GetVmIds() []string`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *FiltersVmsState) GetVmIdsOk() ([]string, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmIds

`func (o *FiltersVmsState) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.

### SetVmIds

`func (o *FiltersVmsState) SetVmIds(v []string)`

SetVmIds gets a reference to the given []string and assigns it to the VmIds field.

### GetVmStates

`func (o *FiltersVmsState) GetVmStates() []string`

GetVmStates returns the VmStates field if non-nil, zero value otherwise.

### GetVmStatesOk

`func (o *FiltersVmsState) GetVmStatesOk() ([]string, bool)`

GetVmStatesOk returns a tuple with the VmStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmStates

`func (o *FiltersVmsState) HasVmStates() bool`

HasVmStates returns a boolean if a field has been set.

### SetVmStates

`func (o *FiltersVmsState) SetVmStates(v []string)`

SetVmStates gets a reference to the given []string and assigns it to the VmStates field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


