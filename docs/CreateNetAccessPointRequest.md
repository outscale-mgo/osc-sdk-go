# CreateNetAccessPointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net. | 
**RouteTableIds** | Pointer to **[]string** | One or more IDs of route tables to use for the connection. | [optional] 
**ServiceName** | Pointer to **string** | The prefix list name corresponding to the service (for example, &#x60;com.outscale.eu-west-2.osu&#x60; for OSU). | 

## Methods

### GetDryRun

`func (o *CreateNetAccessPointRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateNetAccessPointRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateNetAccessPointRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateNetAccessPointRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetNetId

`func (o *CreateNetAccessPointRequest) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *CreateNetAccessPointRequest) GetNetIdOk() (string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetId

`func (o *CreateNetAccessPointRequest) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### SetNetId

`func (o *CreateNetAccessPointRequest) SetNetId(v string)`

SetNetId gets a reference to the given string and assigns it to the NetId field.

### GetRouteTableIds

`func (o *CreateNetAccessPointRequest) GetRouteTableIds() []string`

GetRouteTableIds returns the RouteTableIds field if non-nil, zero value otherwise.

### GetRouteTableIdsOk

`func (o *CreateNetAccessPointRequest) GetRouteTableIdsOk() ([]string, bool)`

GetRouteTableIdsOk returns a tuple with the RouteTableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteTableIds

`func (o *CreateNetAccessPointRequest) HasRouteTableIds() bool`

HasRouteTableIds returns a boolean if a field has been set.

### SetRouteTableIds

`func (o *CreateNetAccessPointRequest) SetRouteTableIds(v []string)`

SetRouteTableIds gets a reference to the given []string and assigns it to the RouteTableIds field.

### GetServiceName

`func (o *CreateNetAccessPointRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *CreateNetAccessPointRequest) GetServiceNameOk() (string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceName

`func (o *CreateNetAccessPointRequest) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceName

`func (o *CreateNetAccessPointRequest) SetServiceName(v string)`

SetServiceName gets a reference to the given string and assigns it to the ServiceName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


