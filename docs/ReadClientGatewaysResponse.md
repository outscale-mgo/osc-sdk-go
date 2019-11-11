# ReadClientGatewaysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientGateways** | Pointer to [**[]ClientGateway**](ClientGateway.md) | Information about one or more client gateways. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetClientGateways

`func (o *ReadClientGatewaysResponse) GetClientGateways() []ClientGateway`

GetClientGateways returns the ClientGateways field if non-nil, zero value otherwise.

### GetClientGatewaysOk

`func (o *ReadClientGatewaysResponse) GetClientGatewaysOk() ([]ClientGateway, bool)`

GetClientGatewaysOk returns a tuple with the ClientGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientGateways

`func (o *ReadClientGatewaysResponse) HasClientGateways() bool`

HasClientGateways returns a boolean if a field has been set.

### SetClientGateways

`func (o *ReadClientGatewaysResponse) SetClientGateways(v []ClientGateway)`

SetClientGateways gets a reference to the given []ClientGateway and assigns it to the ClientGateways field.

### GetResponseContext

`func (o *ReadClientGatewaysResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadClientGatewaysResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadClientGatewaysResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadClientGatewaysResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


