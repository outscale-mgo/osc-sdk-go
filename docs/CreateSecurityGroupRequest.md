# CreateSecurityGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description for the security group, with a maximum length of 255 [ASCII printable characters](https://en.wikipedia.org/wiki/ASCII#Printable_characters). | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net for the security group. | [optional] 
**SecurityGroupName** | Pointer to **string** | (Public Cloud only) The name of the security group.&lt;br /&gt; This name must not start with &#x60;sg-&#x60;.&lt;/br&gt; This name must be unique and contain between 1 and 255 ASCII characters. Accented letters are not allowed. | 

## Methods

### GetDescription

`func (o *CreateSecurityGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSecurityGroupRequest) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *CreateSecurityGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *CreateSecurityGroupRequest) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetDryRun

`func (o *CreateSecurityGroupRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateSecurityGroupRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateSecurityGroupRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateSecurityGroupRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetNetId

`func (o *CreateSecurityGroupRequest) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *CreateSecurityGroupRequest) GetNetIdOk() (string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetId

`func (o *CreateSecurityGroupRequest) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### SetNetId

`func (o *CreateSecurityGroupRequest) SetNetId(v string)`

SetNetId gets a reference to the given string and assigns it to the NetId field.

### GetSecurityGroupName

`func (o *CreateSecurityGroupRequest) GetSecurityGroupName() string`

GetSecurityGroupName returns the SecurityGroupName field if non-nil, zero value otherwise.

### GetSecurityGroupNameOk

`func (o *CreateSecurityGroupRequest) GetSecurityGroupNameOk() (string, bool)`

GetSecurityGroupNameOk returns a tuple with the SecurityGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroupName

`func (o *CreateSecurityGroupRequest) HasSecurityGroupName() bool`

HasSecurityGroupName returns a boolean if a field has been set.

### SetSecurityGroupName

`func (o *CreateSecurityGroupRequest) SetSecurityGroupName(v string)`

SetSecurityGroupName gets a reference to the given string and assigns it to the SecurityGroupName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


