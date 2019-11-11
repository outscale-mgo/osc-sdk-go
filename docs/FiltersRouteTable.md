# FiltersRouteTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkRouteTableLinkRouteTableIds** | Pointer to **[]string** | The IDs of the associations between the route tables and the Subnets. | [optional] 
**LinkSubnetIds** | Pointer to **[]string** | The IDs of the Subnets involved in the associations. | [optional] 
**NetIds** | Pointer to **[]string** | The IDs of the Nets for the route tables. | [optional] 
**RouteCreationMethods** | Pointer to **[]string** | The methods used to create a route. | [optional] 
**RouteDestinationIpRanges** | Pointer to **[]string** | The IP ranges specified in routes in the tables. | [optional] 
**RouteGatewayIds** | Pointer to **[]string** | The IDs of the gateways specified in routes in the tables. | [optional] 
**RouteNatServiceIds** | Pointer to **[]string** | The IDs of the NAT services specified in routes in the tables. | [optional] 
**RouteNetPeeringIds** | Pointer to **[]string** | The IDs of the Net peering connections specified in routes in the tables. | [optional] 
**RouteStates** | Pointer to **[]string** | The states of routes in the route tables (&#x60;active&#x60; \\| &#x60;blackhole&#x60;). The &#x60;blackhole&#x60; state indicates that the target of the route is not available. | [optional] 
**RouteTableIds** | Pointer to **[]string** | The IDs of the route tables. | [optional] 
**RouteVmIds** | Pointer to **[]string** | The IDs of the VMs specified in routes in the tables. | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the route tables. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the route tables. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the route tables, in the following format: \&quot;Filters\&quot;:{\&quot;Tags\&quot;:[\&quot;TAGKEY&#x3D;TAGVALUE\&quot;]}. | [optional] 

## Methods

### GetLinkRouteTableLinkRouteTableIds

`func (o *FiltersRouteTable) GetLinkRouteTableLinkRouteTableIds() []string`

GetLinkRouteTableLinkRouteTableIds returns the LinkRouteTableLinkRouteTableIds field if non-nil, zero value otherwise.

### GetLinkRouteTableLinkRouteTableIdsOk

`func (o *FiltersRouteTable) GetLinkRouteTableLinkRouteTableIdsOk() ([]string, bool)`

GetLinkRouteTableLinkRouteTableIdsOk returns a tuple with the LinkRouteTableLinkRouteTableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkRouteTableLinkRouteTableIds

`func (o *FiltersRouteTable) HasLinkRouteTableLinkRouteTableIds() bool`

HasLinkRouteTableLinkRouteTableIds returns a boolean if a field has been set.

### SetLinkRouteTableLinkRouteTableIds

`func (o *FiltersRouteTable) SetLinkRouteTableLinkRouteTableIds(v []string)`

SetLinkRouteTableLinkRouteTableIds gets a reference to the given []string and assigns it to the LinkRouteTableLinkRouteTableIds field.

### GetLinkSubnetIds

`func (o *FiltersRouteTable) GetLinkSubnetIds() []string`

GetLinkSubnetIds returns the LinkSubnetIds field if non-nil, zero value otherwise.

### GetLinkSubnetIdsOk

`func (o *FiltersRouteTable) GetLinkSubnetIdsOk() ([]string, bool)`

GetLinkSubnetIdsOk returns a tuple with the LinkSubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkSubnetIds

`func (o *FiltersRouteTable) HasLinkSubnetIds() bool`

HasLinkSubnetIds returns a boolean if a field has been set.

### SetLinkSubnetIds

`func (o *FiltersRouteTable) SetLinkSubnetIds(v []string)`

SetLinkSubnetIds gets a reference to the given []string and assigns it to the LinkSubnetIds field.

### GetNetIds

`func (o *FiltersRouteTable) GetNetIds() []string`

GetNetIds returns the NetIds field if non-nil, zero value otherwise.

### GetNetIdsOk

`func (o *FiltersRouteTable) GetNetIdsOk() ([]string, bool)`

GetNetIdsOk returns a tuple with the NetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetIds

`func (o *FiltersRouteTable) HasNetIds() bool`

HasNetIds returns a boolean if a field has been set.

### SetNetIds

`func (o *FiltersRouteTable) SetNetIds(v []string)`

SetNetIds gets a reference to the given []string and assigns it to the NetIds field.

### GetRouteCreationMethods

`func (o *FiltersRouteTable) GetRouteCreationMethods() []string`

GetRouteCreationMethods returns the RouteCreationMethods field if non-nil, zero value otherwise.

### GetRouteCreationMethodsOk

`func (o *FiltersRouteTable) GetRouteCreationMethodsOk() ([]string, bool)`

GetRouteCreationMethodsOk returns a tuple with the RouteCreationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteCreationMethods

`func (o *FiltersRouteTable) HasRouteCreationMethods() bool`

HasRouteCreationMethods returns a boolean if a field has been set.

### SetRouteCreationMethods

`func (o *FiltersRouteTable) SetRouteCreationMethods(v []string)`

SetRouteCreationMethods gets a reference to the given []string and assigns it to the RouteCreationMethods field.

### GetRouteDestinationIpRanges

`func (o *FiltersRouteTable) GetRouteDestinationIpRanges() []string`

GetRouteDestinationIpRanges returns the RouteDestinationIpRanges field if non-nil, zero value otherwise.

### GetRouteDestinationIpRangesOk

`func (o *FiltersRouteTable) GetRouteDestinationIpRangesOk() ([]string, bool)`

GetRouteDestinationIpRangesOk returns a tuple with the RouteDestinationIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteDestinationIpRanges

`func (o *FiltersRouteTable) HasRouteDestinationIpRanges() bool`

HasRouteDestinationIpRanges returns a boolean if a field has been set.

### SetRouteDestinationIpRanges

`func (o *FiltersRouteTable) SetRouteDestinationIpRanges(v []string)`

SetRouteDestinationIpRanges gets a reference to the given []string and assigns it to the RouteDestinationIpRanges field.

### GetRouteGatewayIds

`func (o *FiltersRouteTable) GetRouteGatewayIds() []string`

GetRouteGatewayIds returns the RouteGatewayIds field if non-nil, zero value otherwise.

### GetRouteGatewayIdsOk

`func (o *FiltersRouteTable) GetRouteGatewayIdsOk() ([]string, bool)`

GetRouteGatewayIdsOk returns a tuple with the RouteGatewayIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteGatewayIds

`func (o *FiltersRouteTable) HasRouteGatewayIds() bool`

HasRouteGatewayIds returns a boolean if a field has been set.

### SetRouteGatewayIds

`func (o *FiltersRouteTable) SetRouteGatewayIds(v []string)`

SetRouteGatewayIds gets a reference to the given []string and assigns it to the RouteGatewayIds field.

### GetRouteNatServiceIds

`func (o *FiltersRouteTable) GetRouteNatServiceIds() []string`

GetRouteNatServiceIds returns the RouteNatServiceIds field if non-nil, zero value otherwise.

### GetRouteNatServiceIdsOk

`func (o *FiltersRouteTable) GetRouteNatServiceIdsOk() ([]string, bool)`

GetRouteNatServiceIdsOk returns a tuple with the RouteNatServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteNatServiceIds

`func (o *FiltersRouteTable) HasRouteNatServiceIds() bool`

HasRouteNatServiceIds returns a boolean if a field has been set.

### SetRouteNatServiceIds

`func (o *FiltersRouteTable) SetRouteNatServiceIds(v []string)`

SetRouteNatServiceIds gets a reference to the given []string and assigns it to the RouteNatServiceIds field.

### GetRouteNetPeeringIds

`func (o *FiltersRouteTable) GetRouteNetPeeringIds() []string`

GetRouteNetPeeringIds returns the RouteNetPeeringIds field if non-nil, zero value otherwise.

### GetRouteNetPeeringIdsOk

`func (o *FiltersRouteTable) GetRouteNetPeeringIdsOk() ([]string, bool)`

GetRouteNetPeeringIdsOk returns a tuple with the RouteNetPeeringIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteNetPeeringIds

`func (o *FiltersRouteTable) HasRouteNetPeeringIds() bool`

HasRouteNetPeeringIds returns a boolean if a field has been set.

### SetRouteNetPeeringIds

`func (o *FiltersRouteTable) SetRouteNetPeeringIds(v []string)`

SetRouteNetPeeringIds gets a reference to the given []string and assigns it to the RouteNetPeeringIds field.

### GetRouteStates

`func (o *FiltersRouteTable) GetRouteStates() []string`

GetRouteStates returns the RouteStates field if non-nil, zero value otherwise.

### GetRouteStatesOk

`func (o *FiltersRouteTable) GetRouteStatesOk() ([]string, bool)`

GetRouteStatesOk returns a tuple with the RouteStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteStates

`func (o *FiltersRouteTable) HasRouteStates() bool`

HasRouteStates returns a boolean if a field has been set.

### SetRouteStates

`func (o *FiltersRouteTable) SetRouteStates(v []string)`

SetRouteStates gets a reference to the given []string and assigns it to the RouteStates field.

### GetRouteTableIds

`func (o *FiltersRouteTable) GetRouteTableIds() []string`

GetRouteTableIds returns the RouteTableIds field if non-nil, zero value otherwise.

### GetRouteTableIdsOk

`func (o *FiltersRouteTable) GetRouteTableIdsOk() ([]string, bool)`

GetRouteTableIdsOk returns a tuple with the RouteTableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteTableIds

`func (o *FiltersRouteTable) HasRouteTableIds() bool`

HasRouteTableIds returns a boolean if a field has been set.

### SetRouteTableIds

`func (o *FiltersRouteTable) SetRouteTableIds(v []string)`

SetRouteTableIds gets a reference to the given []string and assigns it to the RouteTableIds field.

### GetRouteVmIds

`func (o *FiltersRouteTable) GetRouteVmIds() []string`

GetRouteVmIds returns the RouteVmIds field if non-nil, zero value otherwise.

### GetRouteVmIdsOk

`func (o *FiltersRouteTable) GetRouteVmIdsOk() ([]string, bool)`

GetRouteVmIdsOk returns a tuple with the RouteVmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteVmIds

`func (o *FiltersRouteTable) HasRouteVmIds() bool`

HasRouteVmIds returns a boolean if a field has been set.

### SetRouteVmIds

`func (o *FiltersRouteTable) SetRouteVmIds(v []string)`

SetRouteVmIds gets a reference to the given []string and assigns it to the RouteVmIds field.

### GetTagKeys

`func (o *FiltersRouteTable) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersRouteTable) GetTagKeysOk() ([]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagKeys

`func (o *FiltersRouteTable) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### SetTagKeys

`func (o *FiltersRouteTable) SetTagKeys(v []string)`

SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.

### GetTagValues

`func (o *FiltersRouteTable) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersRouteTable) GetTagValuesOk() ([]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagValues

`func (o *FiltersRouteTable) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### SetTagValues

`func (o *FiltersRouteTable) SetTagValues(v []string)`

SetTagValues gets a reference to the given []string and assigns it to the TagValues field.

### GetTags

`func (o *FiltersRouteTable) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersRouteTable) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *FiltersRouteTable) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *FiltersRouteTable) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


