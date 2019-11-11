# ReadSubregionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Subregions** | Pointer to [**[]Subregion**](Subregion.md) | Information about one or more Subregions. | [optional] 

## Methods

### GetResponseContext

`func (o *ReadSubregionsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadSubregionsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadSubregionsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadSubregionsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.

### GetSubregions

`func (o *ReadSubregionsResponse) GetSubregions() []Subregion`

GetSubregions returns the Subregions field if non-nil, zero value otherwise.

### GetSubregionsOk

`func (o *ReadSubregionsResponse) GetSubregionsOk() ([]Subregion, bool)`

GetSubregionsOk returns a tuple with the Subregions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubregions

`func (o *ReadSubregionsResponse) HasSubregions() bool`

HasSubregions returns a boolean if a field has been set.

### SetSubregions

`func (o *ReadSubregionsResponse) SetSubregions(v []Subregion)`

SetSubregions gets a reference to the given []Subregion and assigns it to the Subregions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


