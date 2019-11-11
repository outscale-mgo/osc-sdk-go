# NetAccessPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetAccessPointId** | Pointer to **string** | The ID of the Net access point. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net with which the Net access point is associated. | [optional] 
**RouteTableIds** | Pointer to **[]string** | The ID of the route tables associated with the Net access point. | [optional] 
**ServiceName** | Pointer to **string** | The name of the prefix list corresponding to the service with which the Net access point is associated. | [optional] 
**State** | Pointer to **string** | The state of the Net access point (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the Net access point. | [optional] 

## Methods

### GetNetAccessPointId

`func (o *NetAccessPoint) GetNetAccessPointId() string`

GetNetAccessPointId returns the NetAccessPointId field if non-nil, zero value otherwise.

### GetNetAccessPointIdOk

`func (o *NetAccessPoint) GetNetAccessPointIdOk() (string, bool)`

GetNetAccessPointIdOk returns a tuple with the NetAccessPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetAccessPointId

`func (o *NetAccessPoint) HasNetAccessPointId() bool`

HasNetAccessPointId returns a boolean if a field has been set.

### SetNetAccessPointId

`func (o *NetAccessPoint) SetNetAccessPointId(v string)`

SetNetAccessPointId gets a reference to the given string and assigns it to the NetAccessPointId field.

### GetNetId

`func (o *NetAccessPoint) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *NetAccessPoint) GetNetIdOk() (string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetId

`func (o *NetAccessPoint) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### SetNetId

`func (o *NetAccessPoint) SetNetId(v string)`

SetNetId gets a reference to the given string and assigns it to the NetId field.

### GetRouteTableIds

`func (o *NetAccessPoint) GetRouteTableIds() []string`

GetRouteTableIds returns the RouteTableIds field if non-nil, zero value otherwise.

### GetRouteTableIdsOk

`func (o *NetAccessPoint) GetRouteTableIdsOk() ([]string, bool)`

GetRouteTableIdsOk returns a tuple with the RouteTableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRouteTableIds

`func (o *NetAccessPoint) HasRouteTableIds() bool`

HasRouteTableIds returns a boolean if a field has been set.

### SetRouteTableIds

`func (o *NetAccessPoint) SetRouteTableIds(v []string)`

SetRouteTableIds gets a reference to the given []string and assigns it to the RouteTableIds field.

### GetServiceName

`func (o *NetAccessPoint) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *NetAccessPoint) GetServiceNameOk() (string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceName

`func (o *NetAccessPoint) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceName

`func (o *NetAccessPoint) SetServiceName(v string)`

SetServiceName gets a reference to the given string and assigns it to the ServiceName field.

### GetState

`func (o *NetAccessPoint) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NetAccessPoint) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *NetAccessPoint) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *NetAccessPoint) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetTags

`func (o *NetAccessPoint) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NetAccessPoint) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *NetAccessPoint) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *NetAccessPoint) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


