# CreateSecurityGroupRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Flow** | Pointer to **string** | The direction of the flow: &#x60;Inbound&#x60; or &#x60;Outbound&#x60;. You can specify &#x60;Outbound&#x60; for Nets only. | 
**FromPortRange** | Pointer to **int32** | The beginning of the port range for the TCP and UDP protocols, or an ICMP type number. | [optional] 
**IpProtocol** | Pointer to **string** | The IP protocol name (&#x60;tcp&#x60;, &#x60;udp&#x60;, &#x60;icmp&#x60;) or protocol number. By default, &#x60;-1&#x60;, which means all protocols. | [optional] 
**IpRange** | Pointer to **string** | The IP range for the security group rule, in CIDR notation (for example, 10.0.0.0/16). | [optional] 
**Rules** | Pointer to [**[]SecurityGroupRule**](SecurityGroupRule.md) | Information about the security group rule to create. | [optional] 
**SecurityGroupAccountIdToLink** | Pointer to **string** | The account ID of the owner of the security group for which you want to create a rule. | [optional] 
**SecurityGroupId** | Pointer to **string** | The ID of the security group for which you want to create a rule. | 
**SecurityGroupNameToLink** | Pointer to **string** | The ID of the source security group. If you are in the Public Cloud, you can also specify the name of the source security group. | [optional] 
**ToPortRange** | Pointer to **int32** | The end of the port range for the TCP and UDP protocols, or an ICMP type number. | [optional] 

## Methods

### GetDryRun

`func (o *CreateSecurityGroupRuleRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateSecurityGroupRuleRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateSecurityGroupRuleRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateSecurityGroupRuleRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetFlow

`func (o *CreateSecurityGroupRuleRequest) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *CreateSecurityGroupRuleRequest) GetFlowOk() (string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFlow

`func (o *CreateSecurityGroupRuleRequest) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### SetFlow

`func (o *CreateSecurityGroupRuleRequest) SetFlow(v string)`

SetFlow gets a reference to the given string and assigns it to the Flow field.

### GetFromPortRange

`func (o *CreateSecurityGroupRuleRequest) GetFromPortRange() int32`

GetFromPortRange returns the FromPortRange field if non-nil, zero value otherwise.

### GetFromPortRangeOk

`func (o *CreateSecurityGroupRuleRequest) GetFromPortRangeOk() (int32, bool)`

GetFromPortRangeOk returns a tuple with the FromPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFromPortRange

`func (o *CreateSecurityGroupRuleRequest) HasFromPortRange() bool`

HasFromPortRange returns a boolean if a field has been set.

### SetFromPortRange

`func (o *CreateSecurityGroupRuleRequest) SetFromPortRange(v int32)`

SetFromPortRange gets a reference to the given int32 and assigns it to the FromPortRange field.

### GetIpProtocol

`func (o *CreateSecurityGroupRuleRequest) GetIpProtocol() string`

GetIpProtocol returns the IpProtocol field if non-nil, zero value otherwise.

### GetIpProtocolOk

`func (o *CreateSecurityGroupRuleRequest) GetIpProtocolOk() (string, bool)`

GetIpProtocolOk returns a tuple with the IpProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIpProtocol

`func (o *CreateSecurityGroupRuleRequest) HasIpProtocol() bool`

HasIpProtocol returns a boolean if a field has been set.

### SetIpProtocol

`func (o *CreateSecurityGroupRuleRequest) SetIpProtocol(v string)`

SetIpProtocol gets a reference to the given string and assigns it to the IpProtocol field.

### GetIpRange

`func (o *CreateSecurityGroupRuleRequest) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *CreateSecurityGroupRuleRequest) GetIpRangeOk() (string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIpRange

`func (o *CreateSecurityGroupRuleRequest) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### SetIpRange

`func (o *CreateSecurityGroupRuleRequest) SetIpRange(v string)`

SetIpRange gets a reference to the given string and assigns it to the IpRange field.

### GetRules

`func (o *CreateSecurityGroupRuleRequest) GetRules() []SecurityGroupRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateSecurityGroupRuleRequest) GetRulesOk() ([]SecurityGroupRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRules

`func (o *CreateSecurityGroupRuleRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRules

`func (o *CreateSecurityGroupRuleRequest) SetRules(v []SecurityGroupRule)`

SetRules gets a reference to the given []SecurityGroupRule and assigns it to the Rules field.

### GetSecurityGroupAccountIdToLink

`func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupAccountIdToLink() string`

GetSecurityGroupAccountIdToLink returns the SecurityGroupAccountIdToLink field if non-nil, zero value otherwise.

### GetSecurityGroupAccountIdToLinkOk

`func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupAccountIdToLinkOk() (string, bool)`

GetSecurityGroupAccountIdToLinkOk returns a tuple with the SecurityGroupAccountIdToLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroupAccountIdToLink

`func (o *CreateSecurityGroupRuleRequest) HasSecurityGroupAccountIdToLink() bool`

HasSecurityGroupAccountIdToLink returns a boolean if a field has been set.

### SetSecurityGroupAccountIdToLink

`func (o *CreateSecurityGroupRuleRequest) SetSecurityGroupAccountIdToLink(v string)`

SetSecurityGroupAccountIdToLink gets a reference to the given string and assigns it to the SecurityGroupAccountIdToLink field.

### GetSecurityGroupId

`func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupIdOk() (string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroupId

`func (o *CreateSecurityGroupRuleRequest) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.

### SetSecurityGroupId

`func (o *CreateSecurityGroupRuleRequest) SetSecurityGroupId(v string)`

SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.

### GetSecurityGroupNameToLink

`func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupNameToLink() string`

GetSecurityGroupNameToLink returns the SecurityGroupNameToLink field if non-nil, zero value otherwise.

### GetSecurityGroupNameToLinkOk

`func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupNameToLinkOk() (string, bool)`

GetSecurityGroupNameToLinkOk returns a tuple with the SecurityGroupNameToLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroupNameToLink

`func (o *CreateSecurityGroupRuleRequest) HasSecurityGroupNameToLink() bool`

HasSecurityGroupNameToLink returns a boolean if a field has been set.

### SetSecurityGroupNameToLink

`func (o *CreateSecurityGroupRuleRequest) SetSecurityGroupNameToLink(v string)`

SetSecurityGroupNameToLink gets a reference to the given string and assigns it to the SecurityGroupNameToLink field.

### GetToPortRange

`func (o *CreateSecurityGroupRuleRequest) GetToPortRange() int32`

GetToPortRange returns the ToPortRange field if non-nil, zero value otherwise.

### GetToPortRangeOk

`func (o *CreateSecurityGroupRuleRequest) GetToPortRangeOk() (int32, bool)`

GetToPortRangeOk returns a tuple with the ToPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToPortRange

`func (o *CreateSecurityGroupRuleRequest) HasToPortRange() bool`

HasToPortRange returns a boolean if a field has been set.

### SetToPortRange

`func (o *CreateSecurityGroupRuleRequest) SetToPortRange(v int32)`

SetToPortRange gets a reference to the given int32 and assigns it to the ToPortRange field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


