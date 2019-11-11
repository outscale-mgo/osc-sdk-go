# DeleteNetPeeringRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**NetPeeringId** | Pointer to **string** | The ID of the Net peering connection you want to delete. | 

## Methods

### GetDryRun

`func (o *DeleteNetPeeringRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteNetPeeringRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteNetPeeringRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteNetPeeringRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetNetPeeringId

`func (o *DeleteNetPeeringRequest) GetNetPeeringId() string`

GetNetPeeringId returns the NetPeeringId field if non-nil, zero value otherwise.

### GetNetPeeringIdOk

`func (o *DeleteNetPeeringRequest) GetNetPeeringIdOk() (string, bool)`

GetNetPeeringIdOk returns a tuple with the NetPeeringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetPeeringId

`func (o *DeleteNetPeeringRequest) HasNetPeeringId() bool`

HasNetPeeringId returns a boolean if a field has been set.

### SetNetPeeringId

`func (o *DeleteNetPeeringRequest) SetNetPeeringId(v string)`

SetNetPeeringId gets a reference to the given string and assigns it to the NetPeeringId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


