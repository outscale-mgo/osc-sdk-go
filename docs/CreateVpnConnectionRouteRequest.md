# CreateVpnConnectionRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationIpRange** | Pointer to **string** | The network prefix of the route, in CIDR notation (for example, 10.12.0.0/16). | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**VpnConnectionId** | Pointer to **string** | The ID of the target VPN connection of the static route. | 

## Methods

### GetDestinationIpRange

`func (o *CreateVpnConnectionRouteRequest) GetDestinationIpRange() string`

GetDestinationIpRange returns the DestinationIpRange field if non-nil, zero value otherwise.

### GetDestinationIpRangeOk

`func (o *CreateVpnConnectionRouteRequest) GetDestinationIpRangeOk() (string, bool)`

GetDestinationIpRangeOk returns a tuple with the DestinationIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDestinationIpRange

`func (o *CreateVpnConnectionRouteRequest) HasDestinationIpRange() bool`

HasDestinationIpRange returns a boolean if a field has been set.

### SetDestinationIpRange

`func (o *CreateVpnConnectionRouteRequest) SetDestinationIpRange(v string)`

SetDestinationIpRange gets a reference to the given string and assigns it to the DestinationIpRange field.

### GetDryRun

`func (o *CreateVpnConnectionRouteRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateVpnConnectionRouteRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateVpnConnectionRouteRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateVpnConnectionRouteRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetVpnConnectionId

`func (o *CreateVpnConnectionRouteRequest) GetVpnConnectionId() string`

GetVpnConnectionId returns the VpnConnectionId field if non-nil, zero value otherwise.

### GetVpnConnectionIdOk

`func (o *CreateVpnConnectionRouteRequest) GetVpnConnectionIdOk() (string, bool)`

GetVpnConnectionIdOk returns a tuple with the VpnConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVpnConnectionId

`func (o *CreateVpnConnectionRouteRequest) HasVpnConnectionId() bool`

HasVpnConnectionId returns a boolean if a field has been set.

### SetVpnConnectionId

`func (o *CreateVpnConnectionRouteRequest) SetVpnConnectionId(v string)`

SetVpnConnectionId gets a reference to the given string and assigns it to the VpnConnectionId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


