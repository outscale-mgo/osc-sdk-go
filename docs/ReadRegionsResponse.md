# ReadRegionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | Pointer to [**[]Region**](Region.md) | Information about one or more Regions. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetRegions

`func (o *ReadRegionsResponse) GetRegions() []Region`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ReadRegionsResponse) GetRegionsOk() ([]Region, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRegions

`func (o *ReadRegionsResponse) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### SetRegions

`func (o *ReadRegionsResponse) SetRegions(v []Region)`

SetRegions gets a reference to the given []Region and assigns it to the Regions field.

### GetResponseContext

`func (o *ReadRegionsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadRegionsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadRegionsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadRegionsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


