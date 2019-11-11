# DeleteSecurityGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**SecurityGroupId** | Pointer to **string** | The ID of the security group you want to delete. | [optional] 
**SecurityGroupName** | Pointer to **string** | (Public Cloud only) The name of the security group. | [optional] 

## Methods

### GetDryRun

`func (o *DeleteSecurityGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteSecurityGroupRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteSecurityGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteSecurityGroupRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetSecurityGroupId

`func (o *DeleteSecurityGroupRequest) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *DeleteSecurityGroupRequest) GetSecurityGroupIdOk() (string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroupId

`func (o *DeleteSecurityGroupRequest) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### SetSecurityGroupId

`func (o *DeleteSecurityGroupRequest) SetSecurityGroupId(v string)`

SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.

### GetSecurityGroupName

`func (o *DeleteSecurityGroupRequest) GetSecurityGroupName() string`

GetSecurityGroupName returns the SecurityGroupName field if non-nil, zero value otherwise.

### GetSecurityGroupNameOk

`func (o *DeleteSecurityGroupRequest) GetSecurityGroupNameOk() (string, bool)`

GetSecurityGroupNameOk returns a tuple with the SecurityGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroupName

`func (o *DeleteSecurityGroupRequest) HasSecurityGroupName() bool`

HasSecurityGroupName returns a boolean if a field has been set.

### SetSecurityGroupName

`func (o *DeleteSecurityGroupRequest) SetSecurityGroupName(v string)`

SetSecurityGroupName gets a reference to the given string and assigns it to the SecurityGroupName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


