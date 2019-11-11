# UnlinkRouteTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**LinkRouteTableId** | Pointer to **string** | The ID of the association between the route table and the Subnet. | 

## Methods

### GetDryRun

`func (o *UnlinkRouteTableRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UnlinkRouteTableRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UnlinkRouteTableRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UnlinkRouteTableRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetLinkRouteTableId

`func (o *UnlinkRouteTableRequest) GetLinkRouteTableId() string`

GetLinkRouteTableId returns the LinkRouteTableId field if non-nil, zero value otherwise.

### GetLinkRouteTableIdOk

`func (o *UnlinkRouteTableRequest) GetLinkRouteTableIdOk() (string, bool)`

GetLinkRouteTableIdOk returns a tuple with the LinkRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkRouteTableId

`func (o *UnlinkRouteTableRequest) HasLinkRouteTableId() bool`

HasLinkRouteTableId returns a boolean if a field has been set.

### SetLinkRouteTableId

`func (o *UnlinkRouteTableRequest) SetLinkRouteTableId(v string)`

SetLinkRouteTableId gets a reference to the given string and assigns it to the LinkRouteTableId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


