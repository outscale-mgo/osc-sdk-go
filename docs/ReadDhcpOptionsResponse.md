# ReadDhcpOptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpOptionsSets** | Pointer to [**[]DhcpOptionsSet**](DhcpOptionsSet.md) | Information about one or more DHCP options sets. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetDhcpOptionsSets

`func (o *ReadDhcpOptionsResponse) GetDhcpOptionsSets() []DhcpOptionsSet`

GetDhcpOptionsSets returns the DhcpOptionsSets field if non-nil, zero value otherwise.

### GetDhcpOptionsSetsOk

`func (o *ReadDhcpOptionsResponse) GetDhcpOptionsSetsOk() ([]DhcpOptionsSet, bool)`

GetDhcpOptionsSetsOk returns a tuple with the DhcpOptionsSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDhcpOptionsSets

`func (o *ReadDhcpOptionsResponse) HasDhcpOptionsSets() bool`

HasDhcpOptionsSets returns a boolean if a field has been set.

### SetDhcpOptionsSets

`func (o *ReadDhcpOptionsResponse) SetDhcpOptionsSets(v []DhcpOptionsSet)`

SetDhcpOptionsSets gets a reference to the given []DhcpOptionsSet and assigns it to the DhcpOptionsSets field.

### GetResponseContext

`func (o *ReadDhcpOptionsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadDhcpOptionsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadDhcpOptionsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadDhcpOptionsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


