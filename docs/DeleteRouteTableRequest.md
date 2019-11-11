# DeleteRouteTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**RouteTableId** | Pointer to **string** | The ID of the route table you want to delete. | 

## Methods

### GetDryRun

`func (o *DeleteRouteTableRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteRouteTableRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteRouteTableRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteRouteTableRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetRouteTableId

`func (o *DeleteRouteTableRequest) GetRouteTableId() string`

GetRouteTableId returns the RouteTableId field if non-nil, zero value otherwise.

### GetRouteTableIdOk

`func (o *DeleteRouteTableRequest) GetRouteTableIdOk() (string, bool)`

GetRouteTableIdOk returns a tuple with the RouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteTableId

`func (o *DeleteRouteTableRequest) HasRouteTableId() bool`

HasRouteTableId returns a boolean if a field has been set.

### SetRouteTableId

`func (o *DeleteRouteTableRequest) SetRouteTableId(v string)`

SetRouteTableId gets a reference to the given string and assigns it to the RouteTableId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


