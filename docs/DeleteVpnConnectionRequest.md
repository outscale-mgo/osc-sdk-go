# DeleteVpnConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**VpnConnectionId** | Pointer to **string** | The ID of the VPN connection you want to delete. | 

## Methods

### GetDryRun

`func (o *DeleteVpnConnectionRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteVpnConnectionRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteVpnConnectionRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteVpnConnectionRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetVpnConnectionId

`func (o *DeleteVpnConnectionRequest) GetVpnConnectionId() string`

GetVpnConnectionId returns the VpnConnectionId field if non-nil, zero value otherwise.

### GetVpnConnectionIdOk

`func (o *DeleteVpnConnectionRequest) GetVpnConnectionIdOk() (string, bool)`

GetVpnConnectionIdOk returns a tuple with the VpnConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVpnConnectionId

`func (o *DeleteVpnConnectionRequest) HasVpnConnectionId() bool`

HasVpnConnectionId returns a boolean if a field has been set.

### SetVpnConnectionId

`func (o *DeleteVpnConnectionRequest) SetVpnConnectionId(v string)`

SetVpnConnectionId gets a reference to the given string and assigns it to the VpnConnectionId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


