# CreateRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationIpRange** | Pointer to **string** | The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24). | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**GatewayId** | Pointer to **string** | The ID of an Internet service or virtual gateway attached to your Net. | [optional] 
**NatServiceId** | Pointer to **string** | The ID of a NAT service. | [optional] 
**NetPeeringId** | Pointer to **string** | The ID of a Net peering connection. | [optional] 
**NicId** | Pointer to **string** | The ID of a NIC. | [optional] 
**RouteTableId** | Pointer to **string** | The ID of the route table for which you want to create a route. | 
**VmId** | Pointer to **string** | The ID of a NAT VM in your Net (attached to exactly one NIC). | [optional] 

## Methods

### GetDestinationIpRange

`func (o *CreateRouteRequest) GetDestinationIpRange() string`

GetDestinationIpRange returns the DestinationIpRange field if non-nil, zero value otherwise.

### GetDestinationIpRangeOk

`func (o *CreateRouteRequest) GetDestinationIpRangeOk() (string, bool)`

GetDestinationIpRangeOk returns a tuple with the DestinationIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDestinationIpRange

`func (o *CreateRouteRequest) HasDestinationIpRange() bool`

HasDestinationIpRange returns a boolean if a field has been set.

### SetDestinationIpRange

`func (o *CreateRouteRequest) SetDestinationIpRange(v string)`

SetDestinationIpRange gets a reference to the given string and assigns it to the DestinationIpRange field.

### GetDryRun

`func (o *CreateRouteRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateRouteRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateRouteRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateRouteRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetGatewayId

`func (o *CreateRouteRequest) GetGatewayId() string`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *CreateRouteRequest) GetGatewayIdOk() (string, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGatewayId

`func (o *CreateRouteRequest) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### SetGatewayId

`func (o *CreateRouteRequest) SetGatewayId(v string)`

SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.

### GetNatServiceId

`func (o *CreateRouteRequest) GetNatServiceId() string`

GetNatServiceId returns the NatServiceId field if non-nil, zero value otherwise.

### GetNatServiceIdOk

`func (o *CreateRouteRequest) GetNatServiceIdOk() (string, bool)`

GetNatServiceIdOk returns a tuple with the NatServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNatServiceId

`func (o *CreateRouteRequest) HasNatServiceId() bool`

HasNatServiceId returns a boolean if a field has been set.

### SetNatServiceId

`func (o *CreateRouteRequest) SetNatServiceId(v string)`

SetNatServiceId gets a reference to the given string and assigns it to the NatServiceId field.

### GetNetPeeringId

`func (o *CreateRouteRequest) GetNetPeeringId() string`

GetNetPeeringId returns the NetPeeringId field if non-nil, zero value otherwise.

### GetNetPeeringIdOk

`func (o *CreateRouteRequest) GetNetPeeringIdOk() (string, bool)`

GetNetPeeringIdOk returns a tuple with the NetPeeringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetPeeringId

`func (o *CreateRouteRequest) HasNetPeeringId() bool`

HasNetPeeringId returns a boolean if a field has been set.

### SetNetPeeringId

`func (o *CreateRouteRequest) SetNetPeeringId(v string)`

SetNetPeeringId gets a reference to the given string and assigns it to the NetPeeringId field.

### GetNicId

`func (o *CreateRouteRequest) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *CreateRouteRequest) GetNicIdOk() (string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNicId

`func (o *CreateRouteRequest) HasNicId() bool`

HasNicId returns a boolean if a field has been set.

### SetNicId

`func (o *CreateRouteRequest) SetNicId(v string)`

SetNicId gets a reference to the given string and assigns it to the NicId field.

### GetRouteTableId

`func (o *CreateRouteRequest) GetRouteTableId() string`

GetRouteTableId returns the RouteTableId field if non-nil, zero value otherwise.

### GetRouteTableIdOk

`func (o *CreateRouteRequest) GetRouteTableIdOk() (string, bool)`

GetRouteTableIdOk returns a tuple with the RouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteTableId

`func (o *CreateRouteRequest) HasRouteTableId() bool`

HasRouteTableId returns a boolean if a field has been set.

### SetRouteTableId

`func (o *CreateRouteRequest) SetRouteTableId(v string)`

SetRouteTableId gets a reference to the given string and assigns it to the RouteTableId field.

### GetVmId

`func (o *CreateRouteRequest) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *CreateRouteRequest) GetVmIdOk() (string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmId

`func (o *CreateRouteRequest) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### SetVmId

`func (o *CreateRouteRequest) SetVmId(v string)`

SetVmId gets a reference to the given string and assigns it to the VmId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


