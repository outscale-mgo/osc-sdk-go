# ReadNetAccessPointsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetAccessPoints** | Pointer to [**[]NetAccessPoint**](NetAccessPoint.md) | One or more Net access points. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetNetAccessPoints

`func (o *ReadNetAccessPointsResponse) GetNetAccessPoints() []NetAccessPoint`

GetNetAccessPoints returns the NetAccessPoints field if non-nil, zero value otherwise.

### GetNetAccessPointsOk

`func (o *ReadNetAccessPointsResponse) GetNetAccessPointsOk() ([]NetAccessPoint, bool)`

GetNetAccessPointsOk returns a tuple with the NetAccessPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetAccessPoints

`func (o *ReadNetAccessPointsResponse) HasNetAccessPoints() bool`

HasNetAccessPoints returns a boolean if a field has been set.

### SetNetAccessPoints

`func (o *ReadNetAccessPointsResponse) SetNetAccessPoints(v []NetAccessPoint)`

SetNetAccessPoints gets a reference to the given []NetAccessPoint and assigns it to the NetAccessPoints field.

### GetResponseContext

`func (o *ReadNetAccessPointsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadNetAccessPointsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadNetAccessPointsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadNetAccessPointsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


