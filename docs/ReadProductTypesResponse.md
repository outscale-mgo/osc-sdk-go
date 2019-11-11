# ReadProductTypesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductTypes** | Pointer to [**[]ProductType**](ProductType.md) | Information about one or more product types. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetProductTypes

`func (o *ReadProductTypesResponse) GetProductTypes() []ProductType`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *ReadProductTypesResponse) GetProductTypesOk() ([]ProductType, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductTypes

`func (o *ReadProductTypesResponse) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.

### SetProductTypes

`func (o *ReadProductTypesResponse) SetProductTypes(v []ProductType)`

SetProductTypes gets a reference to the given []ProductType and assigns it to the ProductTypes field.

### GetResponseContext

`func (o *ReadProductTypesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadProductTypesResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadProductTypesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadProductTypesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


