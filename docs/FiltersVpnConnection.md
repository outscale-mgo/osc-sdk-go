# FiltersVpnConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAsns** | Pointer to **[]int64** | The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections. | [optional] 
**ClientGatewayIds** | Pointer to **[]string** | The IDs of the client gateways. | [optional] 
**ConnectionTypes** | Pointer to **[]string** | The types of the VPN connections (only &#x60;ipsec.1&#x60; is supported). | [optional] 
**RouteDestinationIpRanges** | Pointer to **[]string** | The destination IP ranges. | [optional] 
**States** | Pointer to **[]string** | The states of the VPN connections (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**StaticRoutesOnly** | Pointer to **bool** | If &#x60;false&#x60;, the VPN connection uses dynamic routing with Border Gateway Protocol (BGP). If &#x60;true&#x60;, routing is controlled using static routes. For more information about how to create and delete static routes, see [CreateVpnConnectionRoute](#createvpnconnectionroute) and [DeleteVpnConnectionRoute](#deletevpnconnectionroute). | [optional] 
**VirtualGatewayIds** | Pointer to **[]string** | The IDs of the virtual gateways. | [optional] 
**VpnConnectionIds** | Pointer to **[]string** | The IDs of the VPN connections. | [optional] 

## Methods

### GetBgpAsns

`func (o *FiltersVpnConnection) GetBgpAsns() []int64`

GetBgpAsns returns the BgpAsns field if non-nil, zero value otherwise.

### GetBgpAsnsOk

`func (o *FiltersVpnConnection) GetBgpAsnsOk() ([]int64, bool)`

GetBgpAsnsOk returns a tuple with the BgpAsns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBgpAsns

`func (o *FiltersVpnConnection) HasBgpAsns() bool`

HasBgpAsns returns a boolean if a field has been set.

### SetBgpAsns

`func (o *FiltersVpnConnection) SetBgpAsns(v []int64)`

SetBgpAsns gets a reference to the given []int64 and assigns it to the BgpAsns field.

### GetClientGatewayIds

`func (o *FiltersVpnConnection) GetClientGatewayIds() []string`

GetClientGatewayIds returns the ClientGatewayIds field if non-nil, zero value otherwise.

### GetClientGatewayIdsOk

`func (o *FiltersVpnConnection) GetClientGatewayIdsOk() ([]string, bool)`

GetClientGatewayIdsOk returns a tuple with the ClientGatewayIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientGatewayIds

`func (o *FiltersVpnConnection) HasClientGatewayIds() bool`

HasClientGatewayIds returns a boolean if a field has been set.

### SetClientGatewayIds

`func (o *FiltersVpnConnection) SetClientGatewayIds(v []string)`

SetClientGatewayIds gets a reference to the given []string and assigns it to the ClientGatewayIds field.

### GetConnectionTypes

`func (o *FiltersVpnConnection) GetConnectionTypes() []string`

GetConnectionTypes returns the ConnectionTypes field if non-nil, zero value otherwise.

### GetConnectionTypesOk

`func (o *FiltersVpnConnection) GetConnectionTypesOk() ([]string, bool)`

GetConnectionTypesOk returns a tuple with the ConnectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectionTypes

`func (o *FiltersVpnConnection) HasConnectionTypes() bool`

HasConnectionTypes returns a boolean if a field has been set.

### SetConnectionTypes

`func (o *FiltersVpnConnection) SetConnectionTypes(v []string)`

SetConnectionTypes gets a reference to the given []string and assigns it to the ConnectionTypes field.

### GetRouteDestinationIpRanges

`func (o *FiltersVpnConnection) GetRouteDestinationIpRanges() []string`

GetRouteDestinationIpRanges returns the RouteDestinationIpRanges field if non-nil, zero value otherwise.

### GetRouteDestinationIpRangesOk

`func (o *FiltersVpnConnection) GetRouteDestinationIpRangesOk() ([]string, bool)`

GetRouteDestinationIpRangesOk returns a tuple with the RouteDestinationIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteDestinationIpRanges

`func (o *FiltersVpnConnection) HasRouteDestinationIpRanges() bool`

HasRouteDestinationIpRanges returns a boolean if a field has been set.

### SetRouteDestinationIpRanges

`func (o *FiltersVpnConnection) SetRouteDestinationIpRanges(v []string)`

SetRouteDestinationIpRanges gets a reference to the given []string and assigns it to the RouteDestinationIpRanges field.

### GetStates

`func (o *FiltersVpnConnection) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersVpnConnection) GetStatesOk() ([]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStates

`func (o *FiltersVpnConnection) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStates

`func (o *FiltersVpnConnection) SetStates(v []string)`

SetStates gets a reference to the given []string and assigns it to the States field.

### GetStaticRoutesOnly

`func (o *FiltersVpnConnection) GetStaticRoutesOnly() bool`

GetStaticRoutesOnly returns the StaticRoutesOnly field if non-nil, zero value otherwise.

### GetStaticRoutesOnlyOk

`func (o *FiltersVpnConnection) GetStaticRoutesOnlyOk() (bool, bool)`

GetStaticRoutesOnlyOk returns a tuple with the StaticRoutesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStaticRoutesOnly

`func (o *FiltersVpnConnection) HasStaticRoutesOnly() bool`

HasStaticRoutesOnly returns a boolean if a field has been set.

### SetStaticRoutesOnly

`func (o *FiltersVpnConnection) SetStaticRoutesOnly(v bool)`

SetStaticRoutesOnly gets a reference to the given bool and assigns it to the StaticRoutesOnly field.

### GetVirtualGatewayIds

`func (o *FiltersVpnConnection) GetVirtualGatewayIds() []string`

GetVirtualGatewayIds returns the VirtualGatewayIds field if non-nil, zero value otherwise.

### GetVirtualGatewayIdsOk

`func (o *FiltersVpnConnection) GetVirtualGatewayIdsOk() ([]string, bool)`

GetVirtualGatewayIdsOk returns a tuple with the VirtualGatewayIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVirtualGatewayIds

`func (o *FiltersVpnConnection) HasVirtualGatewayIds() bool`

HasVirtualGatewayIds returns a boolean if a field has been set.

### SetVirtualGatewayIds

`func (o *FiltersVpnConnection) SetVirtualGatewayIds(v []string)`

SetVirtualGatewayIds gets a reference to the given []string and assigns it to the VirtualGatewayIds field.

### GetVpnConnectionIds

`func (o *FiltersVpnConnection) GetVpnConnectionIds() []string`

GetVpnConnectionIds returns the VpnConnectionIds field if non-nil, zero value otherwise.

### GetVpnConnectionIdsOk

`func (o *FiltersVpnConnection) GetVpnConnectionIdsOk() ([]string, bool)`

GetVpnConnectionIdsOk returns a tuple with the VpnConnectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVpnConnectionIds

`func (o *FiltersVpnConnection) HasVpnConnectionIds() bool`

HasVpnConnectionIds returns a boolean if a field has been set.

### SetVpnConnectionIds

`func (o *FiltersVpnConnection) SetVpnConnectionIds(v []string)`

SetVpnConnectionIds gets a reference to the given []string and assigns it to the VpnConnectionIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


