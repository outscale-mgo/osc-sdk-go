# StartVmsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Vms** | Pointer to [**[]VmState**](VmState.md) | Information about one or more started VMs. | [optional] 

## Methods

### GetResponseContext

`func (o *StartVmsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *StartVmsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *StartVmsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *StartVmsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.

### GetVms

`func (o *StartVmsResponse) GetVms() []VmState`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *StartVmsResponse) GetVmsOk() ([]VmState, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVms

`func (o *StartVmsResponse) HasVms() bool`

HasVms returns a boolean if a field has been set.

### SetVms

`func (o *StartVmsResponse) SetVms(v []VmState)`

SetVms gets a reference to the given []VmState and assigns it to the Vms field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


