# SecurityGroupRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPortRange** | Pointer to **int32** | The beginning of the port range for the TCP and UDP protocols, or an ICMP type number. | [optional] 
**IpProtocol** | Pointer to **string** | The IP protocol name (&#x60;tcp&#x60;, &#x60;udp&#x60;, &#x60;icmp&#x60;) or protocol number. By default, &#x60;-1&#x60;, which means all protocols. | [optional] 
**IpRanges** | Pointer to **[]string** | One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16). | [optional] 
**SecurityGroupsMembers** | Pointer to [**[]SecurityGroupsMember**](SecurityGroupsMember.md) | Information about one or more members of a security group. | [optional] 
**ServiceIds** | Pointer to **[]string** | One or more service IDs to allow traffic from a Net to access the corresponding 3DS OUTSCALE services. For more information, see [ReadNetAccessPointServices](#readnetaccesspointservices). | [optional] 
**ToPortRange** | Pointer to **int32** | The end of the port range for the TCP and UDP protocols, or an ICMP type number. | [optional] 

## Methods

### GetFromPortRange

`func (o *SecurityGroupRule) GetFromPortRange() int32`

GetFromPortRange returns the FromPortRange field if non-nil, zero value otherwise.

### GetFromPortRangeOk

`func (o *SecurityGroupRule) GetFromPortRangeOk() (int32, bool)`

GetFromPortRangeOk returns a tuple with the FromPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFromPortRange

`func (o *SecurityGroupRule) HasFromPortRange() bool`

HasFromPortRange returns a boolean if a field has been set.

### SetFromPortRange

`func (o *SecurityGroupRule) SetFromPortRange(v int32)`

SetFromPortRange gets a reference to the given int32 and assigns it to the FromPortRange field.

### GetIpProtocol

`func (o *SecurityGroupRule) GetIpProtocol() string`

GetIpProtocol returns the IpProtocol field if non-nil, zero value otherwise.

### GetIpProtocolOk

`func (o *SecurityGroupRule) GetIpProtocolOk() (string, bool)`

GetIpProtocolOk returns a tuple with the IpProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIpProtocol

`func (o *SecurityGroupRule) HasIpProtocol() bool`

HasIpProtocol returns a boolean if a field has been set.

### SetIpProtocol

`func (o *SecurityGroupRule) SetIpProtocol(v string)`

SetIpProtocol gets a reference to the given string and assigns it to the IpProtocol field.

### GetIpRanges

`func (o *SecurityGroupRule) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *SecurityGroupRule) GetIpRangesOk() ([]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIpRanges

`func (o *SecurityGroupRule) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### SetIpRanges

`func (o *SecurityGroupRule) SetIpRanges(v []string)`

SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.

### GetSecurityGroupsMembers

`func (o *SecurityGroupRule) GetSecurityGroupsMembers() []SecurityGroupsMember`

GetSecurityGroupsMembers returns the SecurityGroupsMembers field if non-nil, zero value otherwise.

### GetSecurityGroupsMembersOk

`func (o *SecurityGroupRule) GetSecurityGroupsMembersOk() ([]SecurityGroupsMember, bool)`

GetSecurityGroupsMembersOk returns a tuple with the SecurityGroupsMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroupsMembers

`func (o *SecurityGroupRule) HasSecurityGroupsMembers() bool`

HasSecurityGroupsMembers returns a boolean if a field has been set.

### SetSecurityGroupsMembers

`func (o *SecurityGroupRule) SetSecurityGroupsMembers(v []SecurityGroupsMember)`

SetSecurityGroupsMembers gets a reference to the given []SecurityGroupsMember and assigns it to the SecurityGroupsMembers field.

### GetServiceIds

`func (o *SecurityGroupRule) GetServiceIds() []string`

GetServiceIds returns the ServiceIds field if non-nil, zero value otherwise.

### GetServiceIdsOk

`func (o *SecurityGroupRule) GetServiceIdsOk() ([]string, bool)`

GetServiceIdsOk returns a tuple with the ServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceIds

`func (o *SecurityGroupRule) HasServiceIds() bool`

HasServiceIds returns a boolean if a field has been set.

### SetServiceIds

`func (o *SecurityGroupRule) SetServiceIds(v []string)`

SetServiceIds gets a reference to the given []string and assigns it to the ServiceIds field.

### GetToPortRange

`func (o *SecurityGroupRule) GetToPortRange() int32`

GetToPortRange returns the ToPortRange field if non-nil, zero value otherwise.

### GetToPortRangeOk

`func (o *SecurityGroupRule) GetToPortRangeOk() (int32, bool)`

GetToPortRangeOk returns a tuple with the ToPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToPortRange

`func (o *SecurityGroupRule) HasToPortRange() bool`

HasToPortRange returns a boolean if a field has been set.

### SetToPortRange

`func (o *SecurityGroupRule) SetToPortRange(v int32)`

SetToPortRange gets a reference to the given int32 and assigns it to the ToPortRange field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


