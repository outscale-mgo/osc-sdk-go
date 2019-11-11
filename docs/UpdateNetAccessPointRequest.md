# UpdateNetAccessPointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddRouteTableIds** | Pointer to **[]string** | One or more IDs of route tables to associate with the specified Net access point. | [optional] 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**NetAccessPointId** | Pointer to **string** | The ID of the Net access point. | 
**RemoveRouteTableIds** | Pointer to **[]string** | One or more IDs of route tables to disassociate from the specified Net access point. | [optional] 

## Methods

### GetAddRouteTableIds

`func (o *UpdateNetAccessPointRequest) GetAddRouteTableIds() []string`

GetAddRouteTableIds returns the AddRouteTableIds field if non-nil, zero value otherwise.

### GetAddRouteTableIdsOk

`func (o *UpdateNetAccessPointRequest) GetAddRouteTableIdsOk() ([]string, bool)`

GetAddRouteTableIdsOk returns a tuple with the AddRouteTableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAddRouteTableIds

`func (o *UpdateNetAccessPointRequest) HasAddRouteTableIds() bool`

HasAddRouteTableIds returns a boolean if a field has been set.

### SetAddRouteTableIds

`func (o *UpdateNetAccessPointRequest) SetAddRouteTableIds(v []string)`

SetAddRouteTableIds gets a reference to the given []string and assigns it to the AddRouteTableIds field.

### GetDryRun

`func (o *UpdateNetAccessPointRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateNetAccessPointRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UpdateNetAccessPointRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UpdateNetAccessPointRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetNetAccessPointId

`func (o *UpdateNetAccessPointRequest) GetNetAccessPointId() string`

GetNetAccessPointId returns the NetAccessPointId field if non-nil, zero value otherwise.

### GetNetAccessPointIdOk

`func (o *UpdateNetAccessPointRequest) GetNetAccessPointIdOk() (string, bool)`

GetNetAccessPointIdOk returns a tuple with the NetAccessPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetAccessPointId

`func (o *UpdateNetAccessPointRequest) HasNetAccessPointId() bool`

HasNetAccessPointId returns a boolean if a field has been set.

### SetNetAccessPointId

`func (o *UpdateNetAccessPointRequest) SetNetAccessPointId(v string)`

SetNetAccessPointId gets a reference to the given string and assigns it to the NetAccessPointId field.

### GetRemoveRouteTableIds

`func (o *UpdateNetAccessPointRequest) GetRemoveRouteTableIds() []string`

GetRemoveRouteTableIds returns the RemoveRouteTableIds field if non-nil, zero value otherwise.

### GetRemoveRouteTableIdsOk

`func (o *UpdateNetAccessPointRequest) GetRemoveRouteTableIdsOk() ([]string, bool)`

GetRemoveRouteTableIdsOk returns a tuple with the RemoveRouteTableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRemoveRouteTableIds

`func (o *UpdateNetAccessPointRequest) HasRemoveRouteTableIds() bool`

HasRemoveRouteTableIds returns a boolean if a field has been set.

### SetRemoveRouteTableIds

`func (o *UpdateNetAccessPointRequest) SetRemoveRouteTableIds(v []string)`

SetRemoveRouteTableIds gets a reference to the given []string and assigns it to the RemoveRouteTableIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


