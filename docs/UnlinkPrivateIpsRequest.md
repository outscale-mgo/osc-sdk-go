# UnlinkPrivateIpsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**NicId** | Pointer to **string** | The ID of the NIC. | 
**PrivateIps** | Pointer to **[]string** | One or more secondary private IP addresses you want to unassign from the NIC. | 

## Methods

### GetDryRun

`func (o *UnlinkPrivateIpsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UnlinkPrivateIpsRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UnlinkPrivateIpsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UnlinkPrivateIpsRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetNicId

`func (o *UnlinkPrivateIpsRequest) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *UnlinkPrivateIpsRequest) GetNicIdOk() (string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNicId

`func (o *UnlinkPrivateIpsRequest) HasNicId() bool`

HasNicId returns a boolean if a field has been set.

### SetNicId

`func (o *UnlinkPrivateIpsRequest) SetNicId(v string)`

SetNicId gets a reference to the given string and assigns it to the NicId field.

### GetPrivateIps

`func (o *UnlinkPrivateIpsRequest) GetPrivateIps() []string`

GetPrivateIps returns the PrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsOk

`func (o *UnlinkPrivateIpsRequest) GetPrivateIpsOk() ([]string, bool)`

GetPrivateIpsOk returns a tuple with the PrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateIps

`func (o *UnlinkPrivateIpsRequest) HasPrivateIps() bool`

HasPrivateIps returns a boolean if a field has been set.

### SetPrivateIps

`func (o *UnlinkPrivateIpsRequest) SetPrivateIps(v []string)`

SetPrivateIps gets a reference to the given []string and assigns it to the PrivateIps field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


