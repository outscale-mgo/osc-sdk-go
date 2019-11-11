# SecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of a user that has been granted permission. | [optional] 
**Description** | Pointer to **string** | The description of the security group. | [optional] 
**InboundRules** | Pointer to [**[]SecurityGroupRule**](SecurityGroupRule.md) | The inbound rules associated with the security group. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net for the security group. | [optional] 
**OutboundRules** | Pointer to [**[]SecurityGroupRule**](SecurityGroupRule.md) | The outbound rules associated with the security group. | [optional] 
**SecurityGroupId** | Pointer to **string** | The ID of the security group. | [optional] 
**SecurityGroupName** | Pointer to **string** | (Public Cloud only) The name of the security group. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the security group. | [optional] 

## Methods

### GetAccountId

`func (o *SecurityGroup) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SecurityGroup) GetAccountIdOk() (string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *SecurityGroup) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *SecurityGroup) SetAccountId(v string)`

SetAccountId gets a reference to the given string and assigns it to the AccountId field.

### GetDescription

`func (o *SecurityGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityGroup) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *SecurityGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *SecurityGroup) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetInboundRules

`func (o *SecurityGroup) GetInboundRules() []SecurityGroupRule`

GetInboundRules returns the InboundRules field if non-nil, zero value otherwise.

### GetInboundRulesOk

`func (o *SecurityGroup) GetInboundRulesOk() ([]SecurityGroupRule, bool)`

GetInboundRulesOk returns a tuple with the InboundRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInboundRules

`func (o *SecurityGroup) HasInboundRules() bool`

HasInboundRules returns a boolean if a field has been set.

### SetInboundRules

`func (o *SecurityGroup) SetInboundRules(v []SecurityGroupRule)`

SetInboundRules gets a reference to the given []SecurityGroupRule and assigns it to the InboundRules field.

### GetNetId

`func (o *SecurityGroup) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *SecurityGroup) GetNetIdOk() (string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetId

`func (o *SecurityGroup) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### SetNetId

`func (o *SecurityGroup) SetNetId(v string)`

SetNetId gets a reference to the given string and assigns it to the NetId field.

### GetOutboundRules

`func (o *SecurityGroup) GetOutboundRules() []SecurityGroupRule`

GetOutboundRules returns the OutboundRules field if non-nil, zero value otherwise.

### GetOutboundRulesOk

`func (o *SecurityGroup) GetOutboundRulesOk() ([]SecurityGroupRule, bool)`

GetOutboundRulesOk returns a tuple with the OutboundRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOutboundRules

`func (o *SecurityGroup) HasOutboundRules() bool`

HasOutboundRules returns a boolean if a field has been set.

### SetOutboundRules

`func (o *SecurityGroup) SetOutboundRules(v []SecurityGroupRule)`

SetOutboundRules gets a reference to the given []SecurityGroupRule and assigns it to the OutboundRules field.

### GetSecurityGroupId

`func (o *SecurityGroup) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *SecurityGroup) GetSecurityGroupIdOk() (string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroupId

`func (o *SecurityGroup) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### SetSecurityGroupId

`func (o *SecurityGroup) SetSecurityGroupId(v string)`

SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.

### GetSecurityGroupName

`func (o *SecurityGroup) GetSecurityGroupName() string`

GetSecurityGroupName returns the SecurityGroupName field if non-nil, zero value otherwise.

### GetSecurityGroupNameOk

`func (o *SecurityGroup) GetSecurityGroupNameOk() (string, bool)`

GetSecurityGroupNameOk returns a tuple with the SecurityGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroupName

`func (o *SecurityGroup) HasSecurityGroupName() bool`

HasSecurityGroupName returns a boolean if a field has been set.

### SetSecurityGroupName

`func (o *SecurityGroup) SetSecurityGroupName(v string)`

SetSecurityGroupName gets a reference to the given string and assigns it to the SecurityGroupName field.

### GetTags

`func (o *SecurityGroup) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SecurityGroup) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *SecurityGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *SecurityGroup) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


