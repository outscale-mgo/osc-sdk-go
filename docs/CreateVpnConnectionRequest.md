# CreateVpnConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientGatewayId** | Pointer to **string** | The ID of the client gateway. | 
**ConnectionType** | Pointer to **string** | The type of VPN connection (only &#x60;ipsec.1&#x60; is supported). | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**StaticRoutesOnly** | Pointer to **bool** | If &#x60;false&#x60;, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If &#x60;true&#x60;, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](#createvpnconnectionroute) and [DeleteVpnConnectionRoute](#deletevpnconnectionroute). | [optional] 
**VirtualGatewayId** | Pointer to **string** | The ID of the virtual gateway. | 

## Methods

### GetClientGatewayId

`func (o *CreateVpnConnectionRequest) GetClientGatewayId() string`

GetClientGatewayId returns the ClientGatewayId field if non-nil, zero value otherwise.

### GetClientGatewayIdOk

`func (o *CreateVpnConnectionRequest) GetClientGatewayIdOk() (string, bool)`

GetClientGatewayIdOk returns a tuple with the ClientGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientGatewayId

`func (o *CreateVpnConnectionRequest) HasClientGatewayId() bool`

HasClientGatewayId returns a boolean if a field has been set.

### SetClientGatewayId

`func (o *CreateVpnConnectionRequest) SetClientGatewayId(v string)`

SetClientGatewayId gets a reference to the given string and assigns it to the ClientGatewayId field.

### GetConnectionType

`func (o *CreateVpnConnectionRequest) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CreateVpnConnectionRequest) GetConnectionTypeOk() (string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectionType

`func (o *CreateVpnConnectionRequest) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### SetConnectionType

`func (o *CreateVpnConnectionRequest) SetConnectionType(v string)`

SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.

### GetDryRun

`func (o *CreateVpnConnectionRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateVpnConnectionRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateVpnConnectionRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateVpnConnectionRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetStaticRoutesOnly

`func (o *CreateVpnConnectionRequest) GetStaticRoutesOnly() bool`

GetStaticRoutesOnly returns the StaticRoutesOnly field if non-nil, zero value otherwise.

### GetStaticRoutesOnlyOk

`func (o *CreateVpnConnectionRequest) GetStaticRoutesOnlyOk() (bool, bool)`

GetStaticRoutesOnlyOk returns a tuple with the StaticRoutesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStaticRoutesOnly

`func (o *CreateVpnConnectionRequest) HasStaticRoutesOnly() bool`

HasStaticRoutesOnly returns a boolean if a field has been set.

### SetStaticRoutesOnly

`func (o *CreateVpnConnectionRequest) SetStaticRoutesOnly(v bool)`

SetStaticRoutesOnly gets a reference to the given bool and assigns it to the StaticRoutesOnly field.

### GetVirtualGatewayId

`func (o *CreateVpnConnectionRequest) GetVirtualGatewayId() string`

GetVirtualGatewayId returns the VirtualGatewayId field if non-nil, zero value otherwise.

### GetVirtualGatewayIdOk

`func (o *CreateVpnConnectionRequest) GetVirtualGatewayIdOk() (string, bool)`

GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVirtualGatewayId

`func (o *CreateVpnConnectionRequest) HasVirtualGatewayId() bool`

HasVirtualGatewayId returns a boolean if a field has been set.

### SetVirtualGatewayId

`func (o *CreateVpnConnectionRequest) SetVirtualGatewayId(v string)`

SetVirtualGatewayId gets a reference to the given string and assigns it to the VirtualGatewayId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


