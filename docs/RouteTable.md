# RouteTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkRouteTables** | Pointer to [**[]LinkRouteTable**](LinkRouteTable.md) | One or more associations between the route table and Subnets. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net for the route table. | [optional] 
**RoutePropagatingVirtualGateways** | Pointer to [**[]RoutePropagatingVirtualGateway**](RoutePropagatingVirtualGateway.md) | Information about virtual gateways propagating routes. | [optional] 
**RouteTableId** | Pointer to **string** | The ID of the route table. | [optional] 
**Routes** | Pointer to [**[]Route**](Route.md) | One or more routes in the route table. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the route table. | [optional] 

## Methods

### GetLinkRouteTables

`func (o *RouteTable) GetLinkRouteTables() []LinkRouteTable`

GetLinkRouteTables returns the LinkRouteTables field if non-nil, zero value otherwise.

### GetLinkRouteTablesOk

`func (o *RouteTable) GetLinkRouteTablesOk() ([]LinkRouteTable, bool)`

GetLinkRouteTablesOk returns a tuple with the LinkRouteTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkRouteTables

`func (o *RouteTable) HasLinkRouteTables() bool`

HasLinkRouteTables returns a boolean if a field has been set.

### SetLinkRouteTables

`func (o *RouteTable) SetLinkRouteTables(v []LinkRouteTable)`

SetLinkRouteTables gets a reference to the given []LinkRouteTable and assigns it to the LinkRouteTables field.

### GetNetId

`func (o *RouteTable) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *RouteTable) GetNetIdOk() (string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetId

`func (o *RouteTable) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### SetNetId

`func (o *RouteTable) SetNetId(v string)`

SetNetId gets a reference to the given string and assigns it to the NetId field.

### GetRoutePropagatingVirtualGateways

`func (o *RouteTable) GetRoutePropagatingVirtualGateways() []RoutePropagatingVirtualGateway`

GetRoutePropagatingVirtualGateways returns the RoutePropagatingVirtualGateways field if non-nil, zero value otherwise.

### GetRoutePropagatingVirtualGatewaysOk

`func (o *RouteTable) GetRoutePropagatingVirtualGatewaysOk() ([]RoutePropagatingVirtualGateway, bool)`

GetRoutePropagatingVirtualGatewaysOk returns a tuple with the RoutePropagatingVirtualGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoutePropagatingVirtualGateways

`func (o *RouteTable) HasRoutePropagatingVirtualGateways() bool`

HasRoutePropagatingVirtualGateways returns a boolean if a field has been set.

### SetRoutePropagatingVirtualGateways

`func (o *RouteTable) SetRoutePropagatingVirtualGateways(v []RoutePropagatingVirtualGateway)`

SetRoutePropagatingVirtualGateways gets a reference to the given []RoutePropagatingVirtualGateway and assigns it to the RoutePropagatingVirtualGateways field.

### GetRouteTableId

`func (o *RouteTable) GetRouteTableId() string`

GetRouteTableId returns the RouteTableId field if non-nil, zero value otherwise.

### GetRouteTableIdOk

`func (o *RouteTable) GetRouteTableIdOk() (string, bool)`

GetRouteTableIdOk returns a tuple with the RouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteTableId

`func (o *RouteTable) HasRouteTableId() bool`

HasRouteTableId returns a boolean if a field has been set.

### SetRouteTableId

`func (o *RouteTable) SetRouteTableId(v string)`

SetRouteTableId gets a reference to the given string and assigns it to the RouteTableId field.

### GetRoutes

`func (o *RouteTable) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *RouteTable) GetRoutesOk() ([]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoutes

`func (o *RouteTable) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### SetRoutes

`func (o *RouteTable) SetRoutes(v []Route)`

SetRoutes gets a reference to the given []Route and assigns it to the Routes field.

### GetTags

`func (o *RouteTable) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RouteTable) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *RouteTable) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *RouteTable) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


