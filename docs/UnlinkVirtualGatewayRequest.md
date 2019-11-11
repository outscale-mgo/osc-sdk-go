# UnlinkVirtualGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net from which you want to detach the virtual gateway. | 
**VirtualGatewayId** | Pointer to **string** | The ID of the virtual gateway. | 

## Methods

### GetDryRun

`func (o *UnlinkVirtualGatewayRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UnlinkVirtualGatewayRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UnlinkVirtualGatewayRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UnlinkVirtualGatewayRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetNetId

`func (o *UnlinkVirtualGatewayRequest) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *UnlinkVirtualGatewayRequest) GetNetIdOk() (string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetId

`func (o *UnlinkVirtualGatewayRequest) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### SetNetId

`func (o *UnlinkVirtualGatewayRequest) SetNetId(v string)`

SetNetId gets a reference to the given string and assigns it to the NetId field.

### GetVirtualGatewayId

`func (o *UnlinkVirtualGatewayRequest) GetVirtualGatewayId() string`

GetVirtualGatewayId returns the VirtualGatewayId field if non-nil, zero value otherwise.

### GetVirtualGatewayIdOk

`func (o *UnlinkVirtualGatewayRequest) GetVirtualGatewayIdOk() (string, bool)`

GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVirtualGatewayId

`func (o *UnlinkVirtualGatewayRequest) HasVirtualGatewayId() bool`

HasVirtualGatewayId returns a boolean if a field has been set.

### SetVirtualGatewayId

`func (o *UnlinkVirtualGatewayRequest) SetVirtualGatewayId(v string)`

SetVirtualGatewayId gets a reference to the given string and assigns it to the VirtualGatewayId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


