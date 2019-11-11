# DeleteSubnetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**SubnetId** | Pointer to **string** | The ID of the Subnet you want to delete. | 

## Methods

### GetDryRun

`func (o *DeleteSubnetRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteSubnetRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteSubnetRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteSubnetRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetSubnetId

`func (o *DeleteSubnetRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *DeleteSubnetRequest) GetSubnetIdOk() (string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubnetId

`func (o *DeleteSubnetRequest) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetId

`func (o *DeleteSubnetRequest) SetSubnetId(v string)`

SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


