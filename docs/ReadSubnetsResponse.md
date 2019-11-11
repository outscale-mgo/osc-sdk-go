# ReadSubnetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Subnets** | Pointer to [**[]Subnet**](Subnet.md) | Information about one or more Subnets. | [optional] 

## Methods

### GetResponseContext

`func (o *ReadSubnetsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadSubnetsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadSubnetsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadSubnetsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.

### GetSubnets

`func (o *ReadSubnetsResponse) GetSubnets() []Subnet`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *ReadSubnetsResponse) GetSubnetsOk() ([]Subnet, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubnets

`func (o *ReadSubnetsResponse) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnets

`func (o *ReadSubnetsResponse) SetSubnets(v []Subnet)`

SetSubnets gets a reference to the given []Subnet and assigns it to the Subnets field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


