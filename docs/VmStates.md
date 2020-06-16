# VmStates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaintenanceEvents** | Pointer to [**[]MaintenanceEvent**](MaintenanceEvent.md) | One or more scheduled events associated with the VM. | [optional] 
**SubregionName** | Pointer to **string** | The name of the Subregion of the VM. | [optional] 
**VmId** | Pointer to **string** | The ID of the VM. | [optional] 
**VmState** | Pointer to **string** | The state of the VM (&#x60;pending&#x60; \\| &#x60;running&#x60; \\| &#x60;stopping&#x60; \\| &#x60;stopped&#x60; \\| &#x60;shutting-down&#x60; \\| &#x60;terminated&#x60; \\| &#x60;quarantine&#x60;). | [optional] 

## Methods

### GetMaintenanceEvents

`func (o *VmStates) GetMaintenanceEvents() []MaintenanceEvent`

GetMaintenanceEvents returns the MaintenanceEvents field if non-nil, zero value otherwise.

### GetMaintenanceEventsOk

`func (o *VmStates) GetMaintenanceEventsOk() ([]MaintenanceEvent, bool)`

GetMaintenanceEventsOk returns a tuple with the MaintenanceEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaintenanceEvents

`func (o *VmStates) HasMaintenanceEvents() bool`

HasMaintenanceEvents returns a boolean if a field has been set.

### SetMaintenanceEvents

`func (o *VmStates) SetMaintenanceEvents(v []MaintenanceEvent)`

SetMaintenanceEvents gets a reference to the given []MaintenanceEvent and assigns it to the MaintenanceEvents field.

### GetSubregionName

`func (o *VmStates) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *VmStates) GetSubregionNameOk() (string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubregionName

`func (o *VmStates) HasSubregionName() bool`

HasSubregionName returns a boolean if a field has been set.

### SetSubregionName

`func (o *VmStates) SetSubregionName(v string)`

SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.

### GetVmId

`func (o *VmStates) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *VmStates) GetVmIdOk() (string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmId

`func (o *VmStates) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### SetVmId

`func (o *VmStates) SetVmId(v string)`

SetVmId gets a reference to the given string and assigns it to the VmId field.

### GetVmState

`func (o *VmStates) GetVmState() string`

GetVmState returns the VmState field if non-nil, zero value otherwise.

### GetVmStateOk

`func (o *VmStates) GetVmStateOk() (string, bool)`

GetVmStateOk returns a tuple with the VmState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmState

`func (o *VmStates) HasVmState() bool`

HasVmState returns a boolean if a field has been set.

### SetVmState

`func (o *VmStates) SetVmState(v string)`

SetVmState gets a reference to the given string and assigns it to the VmState field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


