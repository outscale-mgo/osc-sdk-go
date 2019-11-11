# ReadVpnConnectionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VpnConnections** | Pointer to [**[]VpnConnection**](VpnConnection.md) | Information about one or more VPN connections. | [optional] 

## Methods

### GetResponseContext

`func (o *ReadVpnConnectionsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVpnConnectionsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadVpnConnectionsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadVpnConnectionsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.

### GetVpnConnections

`func (o *ReadVpnConnectionsResponse) GetVpnConnections() []VpnConnection`

GetVpnConnections returns the VpnConnections field if non-nil, zero value otherwise.

### GetVpnConnectionsOk

`func (o *ReadVpnConnectionsResponse) GetVpnConnectionsOk() ([]VpnConnection, bool)`

GetVpnConnectionsOk returns a tuple with the VpnConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVpnConnections

`func (o *ReadVpnConnectionsResponse) HasVpnConnections() bool`

HasVpnConnections returns a boolean if a field has been set.

### SetVpnConnections

`func (o *ReadVpnConnectionsResponse) SetVpnConnections(v []VpnConnection)`

SetVpnConnections gets a reference to the given []VpnConnection and assigns it to the VpnConnections field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


