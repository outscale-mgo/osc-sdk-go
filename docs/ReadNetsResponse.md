# ReadNetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nets** | Pointer to [**[]Net**](Net.md) | Information about the described Nets. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetNets

`func (o *ReadNetsResponse) GetNets() []Net`

GetNets returns the Nets field if non-nil, zero value otherwise.

### GetNetsOk

`func (o *ReadNetsResponse) GetNetsOk() ([]Net, bool)`

GetNetsOk returns a tuple with the Nets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNets

`func (o *ReadNetsResponse) HasNets() bool`

HasNets returns a boolean if a field has been set.

### SetNets

`func (o *ReadNetsResponse) SetNets(v []Net)`

SetNets gets a reference to the given []Net and assigns it to the Nets field.

### GetResponseContext

`func (o *ReadNetsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadNetsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadNetsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadNetsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


