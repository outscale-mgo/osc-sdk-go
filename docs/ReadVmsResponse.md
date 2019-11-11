# ReadVmsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Vms** | Pointer to [**[]Vm**](Vm.md) | Information about one or more VMs. | [optional] 

## Methods

### GetResponseContext

`func (o *ReadVmsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVmsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadVmsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadVmsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.

### GetVms

`func (o *ReadVmsResponse) GetVms() []Vm`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *ReadVmsResponse) GetVmsOk() ([]Vm, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVms

`func (o *ReadVmsResponse) HasVms() bool`

HasVms returns a boolean if a field has been set.

### SetVms

`func (o *ReadVmsResponse) SetVms(v []Vm)`

SetVms gets a reference to the given []Vm and assigns it to the Vms field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


