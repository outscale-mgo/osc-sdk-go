# UpdateNicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A new description for the NIC. | [optional] 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**LinkNic** | Pointer to [**LinkNicToUpdate**](LinkNicToUpdate.md) |  | [optional] 
**NicId** | Pointer to **string** | The ID of the NIC you want to modify. | 
**SecurityGroupIds** | Pointer to **[]string** | One or more IDs of security groups for the NIC.&lt;br /&gt; You must specify at least one group, even if you use the default security group in the Net. | [optional] 

## Methods

### GetDescription

`func (o *UpdateNicRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNicRequest) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *UpdateNicRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *UpdateNicRequest) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetDryRun

`func (o *UpdateNicRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateNicRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UpdateNicRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UpdateNicRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetLinkNic

`func (o *UpdateNicRequest) GetLinkNic() LinkNicToUpdate`

GetLinkNic returns the LinkNic field if non-nil, zero value otherwise.

### GetLinkNicOk

`func (o *UpdateNicRequest) GetLinkNicOk() (LinkNicToUpdate, bool)`

GetLinkNicOk returns a tuple with the LinkNic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkNic

`func (o *UpdateNicRequest) HasLinkNic() bool`

HasLinkNic returns a boolean if a field has been set.

### SetLinkNic

`func (o *UpdateNicRequest) SetLinkNic(v LinkNicToUpdate)`

SetLinkNic gets a reference to the given LinkNicToUpdate and assigns it to the LinkNic field.

### GetNicId

`func (o *UpdateNicRequest) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *UpdateNicRequest) GetNicIdOk() (string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNicId

`func (o *UpdateNicRequest) HasNicId() bool`

HasNicId returns a boolean if a field has been set.

### SetNicId

`func (o *UpdateNicRequest) SetNicId(v string)`

SetNicId gets a reference to the given string and assigns it to the NicId field.

### GetSecurityGroupIds

`func (o *UpdateNicRequest) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *UpdateNicRequest) GetSecurityGroupIdsOk() ([]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroupIds

`func (o *UpdateNicRequest) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### SetSecurityGroupIds

`func (o *UpdateNicRequest) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds gets a reference to the given []string and assigns it to the SecurityGroupIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


