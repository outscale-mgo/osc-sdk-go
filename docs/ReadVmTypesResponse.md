# ReadVmTypesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**VmTypes** | Pointer to [**[]VmType**](VmType.md) | Information about one or more VM types. | [optional] 

## Methods

### GetResponseContext

`func (o *ReadVmTypesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVmTypesResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadVmTypesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadVmTypesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.

### GetVmTypes

`func (o *ReadVmTypesResponse) GetVmTypes() []VmType`

GetVmTypes returns the VmTypes field if non-nil, zero value otherwise.

### GetVmTypesOk

`func (o *ReadVmTypesResponse) GetVmTypesOk() ([]VmType, bool)`

GetVmTypesOk returns a tuple with the VmTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmTypes

`func (o *ReadVmTypesResponse) HasVmTypes() bool`

HasVmTypes returns a boolean if a field has been set.

### SetVmTypes

`func (o *ReadVmTypesResponse) SetVmTypes(v []VmType)`

SetVmTypes gets a reference to the given []VmType and assigns it to the VmTypes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


