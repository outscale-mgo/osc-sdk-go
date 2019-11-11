# ReadVmsHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendVmHealth** | Pointer to [**[]BackendVmHealth**](BackendVmHealth.md) | Information about the health of one or more back-end VMs. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetBackendVmHealth

`func (o *ReadVmsHealthResponse) GetBackendVmHealth() []BackendVmHealth`

GetBackendVmHealth returns the BackendVmHealth field if non-nil, zero value otherwise.

### GetBackendVmHealthOk

`func (o *ReadVmsHealthResponse) GetBackendVmHealthOk() ([]BackendVmHealth, bool)`

GetBackendVmHealthOk returns a tuple with the BackendVmHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBackendVmHealth

`func (o *ReadVmsHealthResponse) HasBackendVmHealth() bool`

HasBackendVmHealth returns a boolean if a field has been set.

### SetBackendVmHealth

`func (o *ReadVmsHealthResponse) SetBackendVmHealth(v []BackendVmHealth)`

SetBackendVmHealth gets a reference to the given []BackendVmHealth and assigns it to the BackendVmHealth field.

### GetResponseContext

`func (o *ReadVmsHealthResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVmsHealthResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadVmsHealthResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadVmsHealthResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


