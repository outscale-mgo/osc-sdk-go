# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationMethod** | Pointer to **string** | The method used to create the route. | [optional] 
**DestinationIpRange** | Pointer to **string** | The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24). | [optional] 
**DestinationServiceId** | Pointer to **string** | The ID of the 3DS OUTSCALE service. | [optional] 
**GatewayId** | Pointer to **string** | The ID of the Internet service or virtual gateway attached to the Net. | [optional] 
**NatServiceId** | Pointer to **string** | The ID of a NAT service attached to the Net. | [optional] 
**NetAccessPointId** | Pointer to **string** | The ID of the Net access point. | [optional] 
**NetPeeringId** | Pointer to **string** | The ID of the Net peering connection. | [optional] 
**NicId** | Pointer to **string** | The ID of the NIC. | [optional] 
**State** | Pointer to **string** | The state of a route in the route table (&#x60;active&#x60; \\| &#x60;blackhole&#x60;). The &#x60;blackhole&#x60; state indicates that the target of the route is not available. | [optional] 
**VmAccountId** | Pointer to **string** | The account ID of the owner of the VM. | [optional] 
**VmId** | Pointer to **string** | The ID of a VM specified in a route in the table. | [optional] 

## Methods

### GetCreationMethod

`func (o *Route) GetCreationMethod() string`

GetCreationMethod returns the CreationMethod field if non-nil, zero value otherwise.

### GetCreationMethodOk

`func (o *Route) GetCreationMethodOk() (string, bool)`

GetCreationMethodOk returns a tuple with the CreationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreationMethod

`func (o *Route) HasCreationMethod() bool`

HasCreationMethod returns a boolean if a field has been set.

### SetCreationMethod

`func (o *Route) SetCreationMethod(v string)`

SetCreationMethod gets a reference to the given string and assigns it to the CreationMethod field.

### GetDestinationIpRange

`func (o *Route) GetDestinationIpRange() string`

GetDestinationIpRange returns the DestinationIpRange field if non-nil, zero value otherwise.

### GetDestinationIpRangeOk

`func (o *Route) GetDestinationIpRangeOk() (string, bool)`

GetDestinationIpRangeOk returns a tuple with the DestinationIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDestinationIpRange

`func (o *Route) HasDestinationIpRange() bool`

HasDestinationIpRange returns a boolean if a field has been set.

### SetDestinationIpRange

`func (o *Route) SetDestinationIpRange(v string)`

SetDestinationIpRange gets a reference to the given string and assigns it to the DestinationIpRange field.

### GetDestinationServiceId

`func (o *Route) GetDestinationServiceId() string`

GetDestinationServiceId returns the DestinationServiceId field if non-nil, zero value otherwise.

### GetDestinationServiceIdOk

`func (o *Route) GetDestinationServiceIdOk() (string, bool)`

GetDestinationServiceIdOk returns a tuple with the DestinationServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDestinationServiceId

`func (o *Route) HasDestinationServiceId() bool`

HasDestinationServiceId returns a boolean if a field has been set.

### SetDestinationServiceId

`func (o *Route) SetDestinationServiceId(v string)`

SetDestinationServiceId gets a reference to the given string and assigns it to the DestinationServiceId field.

### GetGatewayId

`func (o *Route) GetGatewayId() string`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *Route) GetGatewayIdOk() (string, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGatewayId

`func (o *Route) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### SetGatewayId

`func (o *Route) SetGatewayId(v string)`

SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.

### GetNatServiceId

`func (o *Route) GetNatServiceId() string`

GetNatServiceId returns the NatServiceId field if non-nil, zero value otherwise.

### GetNatServiceIdOk

`func (o *Route) GetNatServiceIdOk() (string, bool)`

GetNatServiceIdOk returns a tuple with the NatServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNatServiceId

`func (o *Route) HasNatServiceId() bool`

HasNatServiceId returns a boolean if a field has been set.

### SetNatServiceId

`func (o *Route) SetNatServiceId(v string)`

SetNatServiceId gets a reference to the given string and assigns it to the NatServiceId field.

### GetNetAccessPointId

`func (o *Route) GetNetAccessPointId() string`

GetNetAccessPointId returns the NetAccessPointId field if non-nil, zero value otherwise.

### GetNetAccessPointIdOk

`func (o *Route) GetNetAccessPointIdOk() (string, bool)`

GetNetAccessPointIdOk returns a tuple with the NetAccessPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetAccessPointId

`func (o *Route) HasNetAccessPointId() bool`

HasNetAccessPointId returns a boolean if a field has been set.

### SetNetAccessPointId

`func (o *Route) SetNetAccessPointId(v string)`

SetNetAccessPointId gets a reference to the given string and assigns it to the NetAccessPointId field.

### GetNetPeeringId

`func (o *Route) GetNetPeeringId() string`

GetNetPeeringId returns the NetPeeringId field if non-nil, zero value otherwise.

### GetNetPeeringIdOk

`func (o *Route) GetNetPeeringIdOk() (string, bool)`

GetNetPeeringIdOk returns a tuple with the NetPeeringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetPeeringId

`func (o *Route) HasNetPeeringId() bool`

HasNetPeeringId returns a boolean if a field has been set.

### SetNetPeeringId

`func (o *Route) SetNetPeeringId(v string)`

SetNetPeeringId gets a reference to the given string and assigns it to the NetPeeringId field.

### GetNicId

`func (o *Route) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *Route) GetNicIdOk() (string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNicId

`func (o *Route) HasNicId() bool`

HasNicId returns a boolean if a field has been set.

### SetNicId

`func (o *Route) SetNicId(v string)`

SetNicId gets a reference to the given string and assigns it to the NicId field.

### GetState

`func (o *Route) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Route) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *Route) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *Route) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetVmAccountId

`func (o *Route) GetVmAccountId() string`

GetVmAccountId returns the VmAccountId field if non-nil, zero value otherwise.

### GetVmAccountIdOk

`func (o *Route) GetVmAccountIdOk() (string, bool)`

GetVmAccountIdOk returns a tuple with the VmAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmAccountId

`func (o *Route) HasVmAccountId() bool`

HasVmAccountId returns a boolean if a field has been set.

### SetVmAccountId

`func (o *Route) SetVmAccountId(v string)`

SetVmAccountId gets a reference to the given string and assigns it to the VmAccountId field.

### GetVmId

`func (o *Route) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *Route) GetVmIdOk() (string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmId

`func (o *Route) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### SetVmId

`func (o *Route) SetVmId(v string)`

SetVmId gets a reference to the given string and assigns it to the VmId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


