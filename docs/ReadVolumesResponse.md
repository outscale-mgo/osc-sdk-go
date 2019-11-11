# ReadVolumesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Volumes** | Pointer to [**[]Volume**](Volume.md) | Information about one or more volumes. | [optional] 

## Methods

### GetResponseContext

`func (o *ReadVolumesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadVolumesResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadVolumesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadVolumesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.

### GetVolumes

`func (o *ReadVolumesResponse) GetVolumes() []Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ReadVolumesResponse) GetVolumesOk() ([]Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumes

`func (o *ReadVolumesResponse) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumes

`func (o *ReadVolumesResponse) SetVolumes(v []Volume)`

SetVolumes gets a reference to the given []Volume and assigns it to the Volumes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


