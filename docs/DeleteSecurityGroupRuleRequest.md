# DeleteSecurityGroupRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Flow** | Pointer to **string** | The direction of the flow: &#x60;Inbound&#x60; or &#x60;Outbound&#x60;. You can specify &#x60;Outbound&#x60; for Nets only. | 
**FromPortRange** | Pointer to **int64** | The beginning of the port range for the TCP and UDP protocols, or an ICMP type number. | [optional] 
**IpProtocol** | Pointer to **string** | The IP protocol name (&#x60;tcp&#x60;, &#x60;udp&#x60;, &#x60;icmp&#x60;) or protocol number. By default, &#x60;-1&#x60;, which means all protocols. | [optional] 
**IpRange** | Pointer to **string** | The IP range for the security group rule, in CIDR notation (for example, 10.0.0.0/16). | [optional] 
**Rules** | Pointer to [**[]SecurityGroupRule**](SecurityGroupRule.md) | One or more rules you want to delete from the security group. | [optional] 
**SecurityGroupAccountIdToUnlink** | Pointer to **string** | The account ID of the owner of the security group you want to delete a rule from. | [optional] 
**SecurityGroupId** | Pointer to **string** | The ID of the security group you want to delete a rule from. | 
**SecurityGroupNameToUnlink** | Pointer to **string** | The ID of the source security group. If you are in the Public Cloud, you can also specify the name of the source security group. | [optional] 
**ToPortRange** | Pointer to **int64** | The end of the port range for the TCP and UDP protocols, or an ICMP type number. | [optional] 

## Methods

### GetDryRun

`func (o *DeleteSecurityGroupRuleRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteSecurityGroupRuleRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteSecurityGroupRuleRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteSecurityGroupRuleRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetFlow

`func (o *DeleteSecurityGroupRuleRequest) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *DeleteSecurityGroupRuleRequest) GetFlowOk() (string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFlow

`func (o *DeleteSecurityGroupRuleRequest) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### SetFlow

`func (o *DeleteSecurityGroupRuleRequest) SetFlow(v string)`

SetFlow gets a reference to the given string and assigns it to the Flow field.

### GetFromPortRange

`func (o *DeleteSecurityGroupRuleRequest) GetFromPortRange() int64`

GetFromPortRange returns the FromPortRange field if non-nil, zero value otherwise.

### GetFromPortRangeOk

`func (o *DeleteSecurityGroupRuleRequest) GetFromPortRangeOk() (int64, bool)`

GetFromPortRangeOk returns a tuple with the FromPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFromPortRange

`func (o *DeleteSecurityGroupRuleRequest) HasFromPortRange() bool`

HasFromPortRange returns a boolean if a field has been set.

### SetFromPortRange

`func (o *DeleteSecurityGroupRuleRequest) SetFromPortRange(v int64)`

SetFromPortRange gets a reference to the given int64 and assigns it to the FromPortRange field.

### GetIpProtocol

`func (o *DeleteSecurityGroupRuleRequest) GetIpProtocol() string`

GetIpProtocol returns the IpProtocol field if non-nil, zero value otherwise.

### GetIpProtocolOk

`func (o *DeleteSecurityGroupRuleRequest) GetIpProtocolOk() (string, bool)`

GetIpProtocolOk returns a tuple with the IpProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIpProtocol

`func (o *DeleteSecurityGroupRuleRequest) HasIpProtocol() bool`

HasIpProtocol returns a boolean if a field has been set.

### SetIpProtocol

`func (o *DeleteSecurityGroupRuleRequest) SetIpProtocol(v string)`

SetIpProtocol gets a reference to the given string and assigns it to the IpProtocol field.

### GetIpRange

`func (o *DeleteSecurityGroupRuleRequest) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *DeleteSecurityGroupRuleRequest) GetIpRangeOk() (string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIpRange

`func (o *DeleteSecurityGroupRuleRequest) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### SetIpRange

`func (o *DeleteSecurityGroupRuleRequest) SetIpRange(v string)`

SetIpRange gets a reference to the given string and assigns it to the IpRange field.

### GetRules

`func (o *DeleteSecurityGroupRuleRequest) GetRules() []SecurityGroupRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DeleteSecurityGroupRuleRequest) GetRulesOk() ([]SecurityGroupRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRules

`func (o *DeleteSecurityGroupRuleRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRules

`func (o *DeleteSecurityGroupRuleRequest) SetRules(v []SecurityGroupRule)`

SetRules gets a reference to the given []SecurityGroupRule and assigns it to the Rules field.

### GetSecurityGroupAccountIdToUnlink

`func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupAccountIdToUnlink() string`

GetSecurityGroupAccountIdToUnlink returns the SecurityGroupAccountIdToUnlink field if non-nil, zero value otherwise.

### GetSecurityGroupAccountIdToUnlinkOk

`func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupAccountIdToUnlinkOk() (string, bool)`

GetSecurityGroupAccountIdToUnlinkOk returns a tuple with the SecurityGroupAccountIdToUnlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroupAccountIdToUnlink

`func (o *DeleteSecurityGroupRuleRequest) HasSecurityGroupAccountIdToUnlink() bool`

HasSecurityGroupAccountIdToUnlink returns a boolean if a field has been set.

### SetSecurityGroupAccountIdToUnlink

`func (o *DeleteSecurityGroupRuleRequest) SetSecurityGroupAccountIdToUnlink(v string)`

SetSecurityGroupAccountIdToUnlink gets a reference to the given string and assigns it to the SecurityGroupAccountIdToUnlink field.

### GetSecurityGroupId

`func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupIdOk() (string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroupId

`func (o *DeleteSecurityGroupRuleRequest) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### SetSecurityGroupId

`func (o *DeleteSecurityGroupRuleRequest) SetSecurityGroupId(v string)`

SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.

### GetSecurityGroupNameToUnlink

`func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupNameToUnlink() string`

GetSecurityGroupNameToUnlink returns the SecurityGroupNameToUnlink field if non-nil, zero value otherwise.

### GetSecurityGroupNameToUnlinkOk

`func (o *DeleteSecurityGroupRuleRequest) GetSecurityGroupNameToUnlinkOk() (string, bool)`

GetSecurityGroupNameToUnlinkOk returns a tuple with the SecurityGroupNameToUnlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroupNameToUnlink

`func (o *DeleteSecurityGroupRuleRequest) HasSecurityGroupNameToUnlink() bool`

HasSecurityGroupNameToUnlink returns a boolean if a field has been set.

### SetSecurityGroupNameToUnlink

`func (o *DeleteSecurityGroupRuleRequest) SetSecurityGroupNameToUnlink(v string)`

SetSecurityGroupNameToUnlink gets a reference to the given string and assigns it to the SecurityGroupNameToUnlink field.

### GetToPortRange

`func (o *DeleteSecurityGroupRuleRequest) GetToPortRange() int64`

GetToPortRange returns the ToPortRange field if non-nil, zero value otherwise.

### GetToPortRangeOk

`func (o *DeleteSecurityGroupRuleRequest) GetToPortRangeOk() (int64, bool)`

GetToPortRangeOk returns a tuple with the ToPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToPortRange

`func (o *DeleteSecurityGroupRuleRequest) HasToPortRange() bool`

HasToPortRange returns a boolean if a field has been set.

### SetToPortRange

`func (o *DeleteSecurityGroupRuleRequest) SetToPortRange(v int64)`

SetToPortRange gets a reference to the given int64 and assigns it to the ToPortRange field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


