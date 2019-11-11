# Nic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of the owner of the NIC. | [optional] 
**Description** | Pointer to **string** | The description of the NIC. | [optional] 
**IsSourceDestChecked** | Pointer to **bool** | (Net only) If &#x60;true&#x60;, the source/destination check is enabled. If &#x60;false&#x60;, it is disabled. This value must be &#x60;false&#x60; for a NAT VM to perform network address translation (NAT) in a Net. | [optional] 
**LinkNic** | Pointer to [**LinkNic**](LinkNic.md) |  | [optional] 
**LinkPublicIp** | Pointer to [**LinkPublicIp**](LinkPublicIp.md) |  | [optional] 
**MacAddress** | Pointer to **string** | The Media Access Control (MAC) address of the NIC. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net for the NIC. | [optional] 
**NicId** | Pointer to **string** | The ID of the NIC. | [optional] 
**PrivateDnsName** | Pointer to **string** | The name of the private DNS. | [optional] 
**PrivateIps** | Pointer to [**[]PrivateIp**](PrivateIp.md) | The private IP addresses of the NIC. | [optional] 
**SecurityGroups** | Pointer to [**[]SecurityGroupLight**](SecurityGroupLight.md) | One or more IDs of security groups for the NIC. | [optional] 
**State** | Pointer to **string** | The state of the NIC (&#x60;available&#x60; \\| &#x60;attaching&#x60; \\| &#x60;in-use&#x60; \\| &#x60;detaching&#x60;). | [optional] 
**SubnetId** | Pointer to **string** | The ID of the Subnet. | [optional] 
**SubregionName** | Pointer to **string** | The Subregion in which the NIC is located. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the NIC. | [optional] 

## Methods

### GetAccountId

`func (o *Nic) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Nic) GetAccountIdOk() (string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *Nic) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *Nic) SetAccountId(v string)`

SetAccountId gets a reference to the given string and assigns it to the AccountId field.

### GetDescription

`func (o *Nic) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Nic) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Nic) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Nic) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetIsSourceDestChecked

`func (o *Nic) GetIsSourceDestChecked() bool`

GetIsSourceDestChecked returns the IsSourceDestChecked field if non-nil, zero value otherwise.

### GetIsSourceDestCheckedOk

`func (o *Nic) GetIsSourceDestCheckedOk() (bool, bool)`

GetIsSourceDestCheckedOk returns a tuple with the IsSourceDestChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsSourceDestChecked

`func (o *Nic) HasIsSourceDestChecked() bool`

HasIsSourceDestChecked returns a boolean if a field has been set.

### SetIsSourceDestChecked

`func (o *Nic) SetIsSourceDestChecked(v bool)`

SetIsSourceDestChecked gets a reference to the given bool and assigns it to the IsSourceDestChecked field.

### GetLinkNic

`func (o *Nic) GetLinkNic() LinkNic`

GetLinkNic returns the LinkNic field if non-nil, zero value otherwise.

### GetLinkNicOk

`func (o *Nic) GetLinkNicOk() (LinkNic, bool)`

GetLinkNicOk returns a tuple with the LinkNic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkNic

`func (o *Nic) HasLinkNic() bool`

HasLinkNic returns a boolean if a field has been set.

### SetLinkNic

`func (o *Nic) SetLinkNic(v LinkNic)`

SetLinkNic gets a reference to the given LinkNic and assigns it to the LinkNic field.

### GetLinkPublicIp

`func (o *Nic) GetLinkPublicIp() LinkPublicIp`

GetLinkPublicIp returns the LinkPublicIp field if non-nil, zero value otherwise.

### GetLinkPublicIpOk

`func (o *Nic) GetLinkPublicIpOk() (LinkPublicIp, bool)`

GetLinkPublicIpOk returns a tuple with the LinkPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkPublicIp

`func (o *Nic) HasLinkPublicIp() bool`

HasLinkPublicIp returns a boolean if a field has been set.

### SetLinkPublicIp

`func (o *Nic) SetLinkPublicIp(v LinkPublicIp)`

SetLinkPublicIp gets a reference to the given LinkPublicIp and assigns it to the LinkPublicIp field.

### GetMacAddress

`func (o *Nic) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *Nic) GetMacAddressOk() (string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMacAddress

`func (o *Nic) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddress

`func (o *Nic) SetMacAddress(v string)`

SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.

### GetNetId

`func (o *Nic) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *Nic) GetNetIdOk() (string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetId

`func (o *Nic) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### SetNetId

`func (o *Nic) SetNetId(v string)`

SetNetId gets a reference to the given string and assigns it to the NetId field.

### GetNicId

`func (o *Nic) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *Nic) GetNicIdOk() (string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNicId

`func (o *Nic) HasNicId() bool`

HasNicId returns a boolean if a field has been set.

### SetNicId

`func (o *Nic) SetNicId(v string)`

SetNicId gets a reference to the given string and assigns it to the NicId field.

### GetPrivateDnsName

`func (o *Nic) GetPrivateDnsName() string`

GetPrivateDnsName returns the PrivateDnsName field if non-nil, zero value otherwise.

### GetPrivateDnsNameOk

`func (o *Nic) GetPrivateDnsNameOk() (string, bool)`

GetPrivateDnsNameOk returns a tuple with the PrivateDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateDnsName

`func (o *Nic) HasPrivateDnsName() bool`

HasPrivateDnsName returns a boolean if a field has been set.

### SetPrivateDnsName

`func (o *Nic) SetPrivateDnsName(v string)`

SetPrivateDnsName gets a reference to the given string and assigns it to the PrivateDnsName field.

### GetPrivateIps

`func (o *Nic) GetPrivateIps() []PrivateIp`

GetPrivateIps returns the PrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsOk

`func (o *Nic) GetPrivateIpsOk() ([]PrivateIp, bool)`

GetPrivateIpsOk returns a tuple with the PrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateIps

`func (o *Nic) HasPrivateIps() bool`

HasPrivateIps returns a boolean if a field has been set.

### SetPrivateIps

`func (o *Nic) SetPrivateIps(v []PrivateIp)`

SetPrivateIps gets a reference to the given []PrivateIp and assigns it to the PrivateIps field.

### GetSecurityGroups

`func (o *Nic) GetSecurityGroups() []SecurityGroupLight`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *Nic) GetSecurityGroupsOk() ([]SecurityGroupLight, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroups

`func (o *Nic) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroups

`func (o *Nic) SetSecurityGroups(v []SecurityGroupLight)`

SetSecurityGroups gets a reference to the given []SecurityGroupLight and assigns it to the SecurityGroups field.

### GetState

`func (o *Nic) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Nic) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *Nic) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *Nic) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetSubnetId

`func (o *Nic) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *Nic) GetSubnetIdOk() (string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubnetId

`func (o *Nic) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetId

`func (o *Nic) SetSubnetId(v string)`

SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.

### GetSubregionName

`func (o *Nic) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *Nic) GetSubregionNameOk() (string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubregionName

`func (o *Nic) HasSubregionName() bool`

HasSubregionName returns a boolean if a field has been set.

### SetSubregionName

`func (o *Nic) SetSubregionName(v string)`

SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.

### GetTags

`func (o *Nic) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Nic) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *Nic) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *Nic) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


