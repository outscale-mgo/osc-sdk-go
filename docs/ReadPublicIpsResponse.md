# ReadPublicIpsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIps** | Pointer to [**[]PublicIp**](PublicIp.md) | Information about one or more EIPs. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetPublicIps

`func (o *ReadPublicIpsResponse) GetPublicIps() []PublicIp`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *ReadPublicIpsResponse) GetPublicIpsOk() ([]PublicIp, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicIps

`func (o *ReadPublicIpsResponse) HasPublicIps() bool`

HasPublicIps returns a boolean if a field has been set.

### SetPublicIps

`func (o *ReadPublicIpsResponse) SetPublicIps(v []PublicIp)`

SetPublicIps gets a reference to the given []PublicIp and assigns it to the PublicIps field.

### GetResponseContext

`func (o *ReadPublicIpsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadPublicIpsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadPublicIpsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadPublicIpsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


