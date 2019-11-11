# ReadRouteTablesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**RouteTables** | Pointer to [**[]RouteTable**](RouteTable.md) | Information about one or more route tables. | [optional] 

## Methods

### GetResponseContext

`func (o *ReadRouteTablesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadRouteTablesResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadRouteTablesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadRouteTablesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.

### GetRouteTables

`func (o *ReadRouteTablesResponse) GetRouteTables() []RouteTable`

GetRouteTables returns the RouteTables field if non-nil, zero value otherwise.

### GetRouteTablesOk

`func (o *ReadRouteTablesResponse) GetRouteTablesOk() ([]RouteTable, bool)`

GetRouteTablesOk returns a tuple with the RouteTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteTables

`func (o *ReadRouteTablesResponse) HasRouteTables() bool`

HasRouteTables returns a boolean if a field has been set.

### SetRouteTables

`func (o *ReadRouteTablesResponse) SetRouteTables(v []RouteTable)`

SetRouteTables gets a reference to the given []RouteTable and assigns it to the RouteTables field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


