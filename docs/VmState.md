# VmState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentState** | Pointer to **string** | The current state of the VM (&#x60;InService&#x60; \\| &#x60;OutOfService&#x60; \\| &#x60;Unknown&#x60;). | [optional] 
**PreviousState** | Pointer to **string** | The previous state of the VM (&#x60;InService&#x60; \\| &#x60;OutOfService&#x60; \\| &#x60;Unknown&#x60;). | [optional] 
**VmId** | Pointer to **string** | The ID of the VM. | [optional] 

## Methods

### GetCurrentState

`func (o *VmState) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *VmState) GetCurrentStateOk() (string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrentState

`func (o *VmState) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### SetCurrentState

`func (o *VmState) SetCurrentState(v string)`

SetCurrentState gets a reference to the given string and assigns it to the CurrentState field.

### GetPreviousState

`func (o *VmState) GetPreviousState() string`

GetPreviousState returns the PreviousState field if non-nil, zero value otherwise.

### GetPreviousStateOk

`func (o *VmState) GetPreviousStateOk() (string, bool)`

GetPreviousStateOk returns a tuple with the PreviousState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPreviousState

`func (o *VmState) HasPreviousState() bool`

HasPreviousState returns a boolean if a field has been set.

### SetPreviousState

`func (o *VmState) SetPreviousState(v string)`

SetPreviousState gets a reference to the given string and assigns it to the PreviousState field.

### GetVmId

`func (o *VmState) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *VmState) GetVmIdOk() (string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmId

`func (o *VmState) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### SetVmId

`func (o *VmState) SetVmId(v string)`

SetVmId gets a reference to the given string and assigns it to the VmId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


