# VpnConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientGatewayConfiguration** | Pointer to **string** | The configuration to apply to the client gateway to establish the VPN connection, in XML format. | [optional] 
**ClientGatewayId** | Pointer to **string** | The ID of the client gateway used on the client end of the connection. | [optional] 
**ConnectionType** | Pointer to **string** | The type of VPN connection (always &#x60;ipsec.1&#x60;). | [optional] 
**Routes** | Pointer to [**[]RouteLight**](RouteLight.md) | Information about one or more static routes associated with the VPN connection, if any. | [optional] 
**State** | Pointer to **string** | The state of the VPN connection (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**StaticRoutesOnly** | Pointer to **bool** | If &#x60;false&#x60;, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If &#x60;true&#x60;, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](#createvpnconnectionroute) and [DeleteVpnConnectionRoute](#deletevpnconnectionroute). | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the VPN connection. | [optional] 
**VirtualGatewayId** | Pointer to **string** | The ID of the virtual gateway used on the 3DS OUTSCALE end of the connection. | [optional] 
**VpnConnectionId** | Pointer to **string** | The ID of the VPN connection. | [optional] 

## Methods

### GetClientGatewayConfiguration

`func (o *VpnConnection) GetClientGatewayConfiguration() string`

GetClientGatewayConfiguration returns the ClientGatewayConfiguration field if non-nil, zero value otherwise.

### GetClientGatewayConfigurationOk

`func (o *VpnConnection) GetClientGatewayConfigurationOk() (string, bool)`

GetClientGatewayConfigurationOk returns a tuple with the ClientGatewayConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientGatewayConfiguration

`func (o *VpnConnection) HasClientGatewayConfiguration() bool`

HasClientGatewayConfiguration returns a boolean if a field has been set.

### SetClientGatewayConfiguration

`func (o *VpnConnection) SetClientGatewayConfiguration(v string)`

SetClientGatewayConfiguration gets a reference to the given string and assigns it to the ClientGatewayConfiguration field.

### GetClientGatewayId

`func (o *VpnConnection) GetClientGatewayId() string`

GetClientGatewayId returns the ClientGatewayId field if non-nil, zero value otherwise.

### GetClientGatewayIdOk

`func (o *VpnConnection) GetClientGatewayIdOk() (string, bool)`

GetClientGatewayIdOk returns a tuple with the ClientGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientGatewayId

`func (o *VpnConnection) HasClientGatewayId() bool`

HasClientGatewayId returns a boolean if a field has been set.

### SetClientGatewayId

`func (o *VpnConnection) SetClientGatewayId(v string)`

SetClientGatewayId gets a reference to the given string and assigns it to the ClientGatewayId field.

### GetConnectionType

`func (o *VpnConnection) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *VpnConnection) GetConnectionTypeOk() (string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectionType

`func (o *VpnConnection) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### SetConnectionType

`func (o *VpnConnection) SetConnectionType(v string)`

SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.

### GetRoutes

`func (o *VpnConnection) GetRoutes() []RouteLight`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *VpnConnection) GetRoutesOk() ([]RouteLight, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoutes

`func (o *VpnConnection) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### SetRoutes

`func (o *VpnConnection) SetRoutes(v []RouteLight)`

SetRoutes gets a reference to the given []RouteLight and assigns it to the Routes field.

### GetState

`func (o *VpnConnection) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VpnConnection) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *VpnConnection) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *VpnConnection) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetStaticRoutesOnly

`func (o *VpnConnection) GetStaticRoutesOnly() bool`

GetStaticRoutesOnly returns the StaticRoutesOnly field if non-nil, zero value otherwise.

### GetStaticRoutesOnlyOk

`func (o *VpnConnection) GetStaticRoutesOnlyOk() (bool, bool)`

GetStaticRoutesOnlyOk returns a tuple with the StaticRoutesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStaticRoutesOnly

`func (o *VpnConnection) HasStaticRoutesOnly() bool`

HasStaticRoutesOnly returns a boolean if a field has been set.

### SetStaticRoutesOnly

`func (o *VpnConnection) SetStaticRoutesOnly(v bool)`

SetStaticRoutesOnly gets a reference to the given bool and assigns it to the StaticRoutesOnly field.

### GetTags

`func (o *VpnConnection) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VpnConnection) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *VpnConnection) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *VpnConnection) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.

### GetVirtualGatewayId

`func (o *VpnConnection) GetVirtualGatewayId() string`

GetVirtualGatewayId returns the VirtualGatewayId field if non-nil, zero value otherwise.

### GetVirtualGatewayIdOk

`func (o *VpnConnection) GetVirtualGatewayIdOk() (string, bool)`

GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVirtualGatewayId

`func (o *VpnConnection) HasVirtualGatewayId() bool`

HasVirtualGatewayId returns a boolean if a field has been set.

### SetVirtualGatewayId

`func (o *VpnConnection) SetVirtualGatewayId(v string)`

SetVirtualGatewayId gets a reference to the given string and assigns it to the VirtualGatewayId field.

### GetVpnConnectionId

`func (o *VpnConnection) GetVpnConnectionId() string`

GetVpnConnectionId returns the VpnConnectionId field if non-nil, zero value otherwise.

### GetVpnConnectionIdOk

`func (o *VpnConnection) GetVpnConnectionIdOk() (string, bool)`

GetVpnConnectionIdOk returns a tuple with the VpnConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVpnConnectionId

`func (o *VpnConnection) HasVpnConnectionId() bool`

HasVpnConnectionId returns a boolean if a field has been set.

### SetVpnConnectionId

`func (o *VpnConnection) SetVpnConnectionId(v string)`

SetVpnConnectionId gets a reference to the given string and assigns it to the VpnConnectionId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


