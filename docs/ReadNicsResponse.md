# ReadNicsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nics** | Pointer to [**[]Nic**](Nic.md) | Information about one or more NICs. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetNics

`func (o *ReadNicsResponse) GetNics() []Nic`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *ReadNicsResponse) GetNicsOk() ([]Nic, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNics

`func (o *ReadNicsResponse) HasNics() bool`

HasNics returns a boolean if a field has been set.

### SetNics

`func (o *ReadNicsResponse) SetNics(v []Nic)`

SetNics gets a reference to the given []Nic and assigns it to the Nics field.

### GetResponseContext

`func (o *ReadNicsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadNicsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadNicsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadNicsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


