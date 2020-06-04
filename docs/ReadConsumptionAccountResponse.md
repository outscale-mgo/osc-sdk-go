# ReadConsumptionAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumptionEntries** | Pointer to [**[]ConsumptionEntry**](ConsumptionEntry.md) | Information about the resources consumed during the specified time period. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetConsumptionEntries

`func (o *ReadConsumptionAccountResponse) GetConsumptionEntries() []ConsumptionEntry`

GetConsumptionEntries returns the ConsumptionEntries field if non-nil, zero value otherwise.

### GetConsumptionEntriesOk

`func (o *ReadConsumptionAccountResponse) GetConsumptionEntriesOk() ([]ConsumptionEntry, bool)`

GetConsumptionEntriesOk returns a tuple with the ConsumptionEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConsumptionEntries

`func (o *ReadConsumptionAccountResponse) HasConsumptionEntries() bool`

HasConsumptionEntries returns a boolean if a field has been set.

### SetConsumptionEntries

`func (o *ReadConsumptionAccountResponse) SetConsumptionEntries(v []ConsumptionEntry)`

SetConsumptionEntries gets a reference to the given []ConsumptionEntry and assigns it to the ConsumptionEntries field.

### GetResponseContext

`func (o *ReadConsumptionAccountResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadConsumptionAccountResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadConsumptionAccountResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadConsumptionAccountResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


