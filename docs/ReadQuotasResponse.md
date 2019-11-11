# ReadQuotasResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuotaTypes** | Pointer to [**[]QuotaTypes**](QuotaTypes.md) | Information about one or more quotas. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetQuotaTypes

`func (o *ReadQuotasResponse) GetQuotaTypes() []QuotaTypes`

GetQuotaTypes returns the QuotaTypes field if non-nil, zero value otherwise.

### GetQuotaTypesOk

`func (o *ReadQuotasResponse) GetQuotaTypesOk() ([]QuotaTypes, bool)`

GetQuotaTypesOk returns a tuple with the QuotaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuotaTypes

`func (o *ReadQuotasResponse) HasQuotaTypes() bool`

HasQuotaTypes returns a boolean if a field has been set.

### SetQuotaTypes

`func (o *ReadQuotasResponse) SetQuotaTypes(v []QuotaTypes)`

SetQuotaTypes gets a reference to the given []QuotaTypes and assigns it to the QuotaTypes field.

### GetResponseContext

`func (o *ReadQuotasResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadQuotasResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadQuotasResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadQuotasResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


