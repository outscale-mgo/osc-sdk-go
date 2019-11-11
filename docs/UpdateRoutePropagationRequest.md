# UpdateRoutePropagationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Enable** | Pointer to **bool** | If &#x60;true&#x60;, a virtual gateway can propagate routes to a specified route table of a Net. If &#x60;false&#x60;, the propagation is disabled. | 
**RouteTableId** | Pointer to **string** | The ID of the route table. | 
**VirtualGatewayId** | Pointer to **string** | The ID of the virtual gateway. | 

## Methods

### GetDryRun

`func (o *UpdateRoutePropagationRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateRoutePropagationRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UpdateRoutePropagationRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UpdateRoutePropagationRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetEnable

`func (o *UpdateRoutePropagationRequest) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *UpdateRoutePropagationRequest) GetEnableOk() (bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnable

`func (o *UpdateRoutePropagationRequest) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnable

`func (o *UpdateRoutePropagationRequest) SetEnable(v bool)`

SetEnable gets a reference to the given bool and assigns it to the Enable field.

### GetRouteTableId

`func (o *UpdateRoutePropagationRequest) GetRouteTableId() string`

GetRouteTableId returns the RouteTableId field if non-nil, zero value otherwise.

### GetRouteTableIdOk

`func (o *UpdateRoutePropagationRequest) GetRouteTableIdOk() (string, bool)`

GetRouteTableIdOk returns a tuple with the RouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteTableId

`func (o *UpdateRoutePropagationRequest) HasRouteTableId() bool`

HasRouteTableId returns a boolean if a field has been set.

### SetRouteTableId

`func (o *UpdateRoutePropagationRequest) SetRouteTableId(v string)`

SetRouteTableId gets a reference to the given string and assigns it to the RouteTableId field.

### GetVirtualGatewayId

`func (o *UpdateRoutePropagationRequest) GetVirtualGatewayId() string`

GetVirtualGatewayId returns the VirtualGatewayId field if non-nil, zero value otherwise.

### GetVirtualGatewayIdOk

`func (o *UpdateRoutePropagationRequest) GetVirtualGatewayIdOk() (string, bool)`

GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVirtualGatewayId

`func (o *UpdateRoutePropagationRequest) HasVirtualGatewayId() bool`

HasVirtualGatewayId returns a boolean if a field has been set.

### SetVirtualGatewayId

`func (o *UpdateRoutePropagationRequest) SetVirtualGatewayId(v string)`

SetVirtualGatewayId gets a reference to the given string and assigns it to the VirtualGatewayId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


