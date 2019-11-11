# ReadPublicIpRangesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIps** | Pointer to **[]string** | The list of public IPv4 addresses used in the Region, in CIDR notation. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetPublicIps

`func (o *ReadPublicIpRangesResponse) GetPublicIps() []string`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *ReadPublicIpRangesResponse) GetPublicIpsOk() ([]string, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicIps

`func (o *ReadPublicIpRangesResponse) HasPublicIps() bool`

HasPublicIps returns a boolean if a field has been set.

### SetPublicIps

`func (o *ReadPublicIpRangesResponse) SetPublicIps(v []string)`

SetPublicIps gets a reference to the given []string and assigns it to the PublicIps field.

### GetResponseContext

`func (o *ReadPublicIpRangesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadPublicIpRangesResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadPublicIpRangesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadPublicIpRangesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


