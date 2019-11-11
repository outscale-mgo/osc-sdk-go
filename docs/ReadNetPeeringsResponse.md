# ReadNetPeeringsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetPeerings** | Pointer to [**[]NetPeering**](NetPeering.md) | Information about one or more Net peering connections. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetNetPeerings

`func (o *ReadNetPeeringsResponse) GetNetPeerings() []NetPeering`

GetNetPeerings returns the NetPeerings field if non-nil, zero value otherwise.

### GetNetPeeringsOk

`func (o *ReadNetPeeringsResponse) GetNetPeeringsOk() ([]NetPeering, bool)`

GetNetPeeringsOk returns a tuple with the NetPeerings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetPeerings

`func (o *ReadNetPeeringsResponse) HasNetPeerings() bool`

HasNetPeerings returns a boolean if a field has been set.

### SetNetPeerings

`func (o *ReadNetPeeringsResponse) SetNetPeerings(v []NetPeering)`

SetNetPeerings gets a reference to the given []NetPeering and assigns it to the NetPeerings field.

### GetResponseContext

`func (o *ReadNetPeeringsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadNetPeeringsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadNetPeeringsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadNetPeeringsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


