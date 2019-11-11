# ReadVmsStateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VmStates** | Pointer to [**[]VmStates**](VmStates.md) | Information about one or more VM states. | [optional] 

## Methods

### GetResponseContext

`func (o *ReadVmsStateResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVmsStateResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadVmsStateResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadVmsStateResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.

### GetVmStates

`func (o *ReadVmsStateResponse) GetVmStates() []VmStates`

GetVmStates returns the VmStates field if non-nil, zero value otherwise.

### GetVmStatesOk

`func (o *ReadVmsStateResponse) GetVmStatesOk() ([]VmStates, bool)`

GetVmStatesOk returns a tuple with the VmStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmStates

`func (o *ReadVmsStateResponse) HasVmStates() bool`

HasVmStates returns a boolean if a field has been set.

### SetVmStates

`func (o *ReadVmsStateResponse) SetVmStates(v []VmStates)`

SetVmStates gets a reference to the given []VmStates and assigns it to the VmStates field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


