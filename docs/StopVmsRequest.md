# StopVmsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**ForceStop** | Pointer to **bool** | Forces the VM to stop. | [optional] 
**VmIds** | Pointer to **[]string** | One or more IDs of VMs. | 

## Methods

### GetDryRun

`func (o *StopVmsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *StopVmsRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *StopVmsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *StopVmsRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetForceStop

`func (o *StopVmsRequest) GetForceStop() bool`

GetForceStop returns the ForceStop field if non-nil, zero value otherwise.

### GetForceStopOk

`func (o *StopVmsRequest) GetForceStopOk() (bool, bool)`

GetForceStopOk returns a tuple with the ForceStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasForceStop

`func (o *StopVmsRequest) HasForceStop() bool`

HasForceStop returns a boolean if a field has been set.

### SetForceStop

`func (o *StopVmsRequest) SetForceStop(v bool)`

SetForceStop gets a reference to the given bool and assigns it to the ForceStop field.

### GetVmIds

`func (o *StopVmsRequest) GetVmIds() []string`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *StopVmsRequest) GetVmIdsOk() ([]string, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmIds

`func (o *StopVmsRequest) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.

### SetVmIds

`func (o *StopVmsRequest) SetVmIds(v []string)`

SetVmIds gets a reference to the given []string and assigns it to the VmIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


