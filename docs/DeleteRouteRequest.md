# DeleteRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationIpRange** | Pointer to **string** | The exact IP range for the route. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**RouteTableId** | Pointer to **string** | The ID of the route table from which you want to delete a route. | 

## Methods

### GetDestinationIpRange

`func (o *DeleteRouteRequest) GetDestinationIpRange() string`

GetDestinationIpRange returns the DestinationIpRange field if non-nil, zero value otherwise.

### GetDestinationIpRangeOk

`func (o *DeleteRouteRequest) GetDestinationIpRangeOk() (string, bool)`

GetDestinationIpRangeOk returns a tuple with the DestinationIpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDestinationIpRange

`func (o *DeleteRouteRequest) HasDestinationIpRange() bool`

HasDestinationIpRange returns a boolean if a field has been set.

### SetDestinationIpRange

`func (o *DeleteRouteRequest) SetDestinationIpRange(v string)`

SetDestinationIpRange gets a reference to the given string and assigns it to the DestinationIpRange field.

### GetDryRun

`func (o *DeleteRouteRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteRouteRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteRouteRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteRouteRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetRouteTableId

`func (o *DeleteRouteRequest) GetRouteTableId() string`

GetRouteTableId returns the RouteTableId field if non-nil, zero value otherwise.

### GetRouteTableIdOk

`func (o *DeleteRouteRequest) GetRouteTableIdOk() (string, bool)`

GetRouteTableIdOk returns a tuple with the RouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteTableId

`func (o *DeleteRouteRequest) HasRouteTableId() bool`

HasRouteTableId returns a boolean if a field has been set.

### SetRouteTableId

`func (o *DeleteRouteRequest) SetRouteTableId(v string)`

SetRouteTableId gets a reference to the given string and assigns it to the RouteTableId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


