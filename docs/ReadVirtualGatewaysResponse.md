# ReadVirtualGatewaysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VirtualGateways** | Pointer to [**[]VirtualGateway**](VirtualGateway.md) | Information about one or more virtual gateways. | [optional] 

## Methods

### GetResponseContext

`func (o *ReadVirtualGatewaysResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVirtualGatewaysResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadVirtualGatewaysResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadVirtualGatewaysResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.

### GetVirtualGateways

`func (o *ReadVirtualGatewaysResponse) GetVirtualGateways() []VirtualGateway`

GetVirtualGateways returns the VirtualGateways field if non-nil, zero value otherwise.

### GetVirtualGatewaysOk

`func (o *ReadVirtualGatewaysResponse) GetVirtualGatewaysOk() ([]VirtualGateway, bool)`

GetVirtualGatewaysOk returns a tuple with the VirtualGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVirtualGateways

`func (o *ReadVirtualGatewaysResponse) HasVirtualGateways() bool`

HasVirtualGateways returns a boolean if a field has been set.

### SetVirtualGateways

`func (o *ReadVirtualGatewaysResponse) SetVirtualGateways(v []VirtualGateway)`

SetVirtualGateways gets a reference to the given []VirtualGateway and assigns it to the VirtualGateways field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


