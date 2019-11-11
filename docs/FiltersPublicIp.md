# FiltersPublicIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkPublicIpIds** | Pointer to **[]string** | The IDs representing the associations of EIPs with VMs or NICs. | [optional] 
**NicAccountIds** | Pointer to **[]string** | The account IDs of the owners of the NICs. | [optional] 
**NicIds** | Pointer to **[]string** | The IDs of the NICs. | [optional] 
**Placements** | Pointer to **[]string** | Whether the EIPs are for use in the public Cloud or in a Net. | [optional] 
**PrivateIps** | Pointer to **[]string** | The private IP addresses associated with the EIPs. | [optional] 
**PublicIpIds** | Pointer to **[]string** | The IDs of the External IP addresses (EIPs). | [optional] 
**PublicIps** | Pointer to **[]string** | The External IP addresses (EIPs). | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the EIPs. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the EIPs. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the EIPs, in the following format: \&quot;Filters\&quot;:{\&quot;Tags\&quot;:[\&quot;TAGKEY&#x3D;TAGVALUE\&quot;]}. | [optional] 
**VmIds** | Pointer to **[]string** | The IDs of the VMs. | [optional] 

## Methods

### GetLinkPublicIpIds

`func (o *FiltersPublicIp) GetLinkPublicIpIds() []string`

GetLinkPublicIpIds returns the LinkPublicIpIds field if non-nil, zero value otherwise.

### GetLinkPublicIpIdsOk

`func (o *FiltersPublicIp) GetLinkPublicIpIdsOk() ([]string, bool)`

GetLinkPublicIpIdsOk returns a tuple with the LinkPublicIpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkPublicIpIds

`func (o *FiltersPublicIp) HasLinkPublicIpIds() bool`

HasLinkPublicIpIds returns a boolean if a field has been set.

### SetLinkPublicIpIds

`func (o *FiltersPublicIp) SetLinkPublicIpIds(v []string)`

SetLinkPublicIpIds gets a reference to the given []string and assigns it to the LinkPublicIpIds field.

### GetNicAccountIds

`func (o *FiltersPublicIp) GetNicAccountIds() []string`

GetNicAccountIds returns the NicAccountIds field if non-nil, zero value otherwise.

### GetNicAccountIdsOk

`func (o *FiltersPublicIp) GetNicAccountIdsOk() ([]string, bool)`

GetNicAccountIdsOk returns a tuple with the NicAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNicAccountIds

`func (o *FiltersPublicIp) HasNicAccountIds() bool`

HasNicAccountIds returns a boolean if a field has been set.

### SetNicAccountIds

`func (o *FiltersPublicIp) SetNicAccountIds(v []string)`

SetNicAccountIds gets a reference to the given []string and assigns it to the NicAccountIds field.

### GetNicIds

`func (o *FiltersPublicIp) GetNicIds() []string`

GetNicIds returns the NicIds field if non-nil, zero value otherwise.

### GetNicIdsOk

`func (o *FiltersPublicIp) GetNicIdsOk() ([]string, bool)`

GetNicIdsOk returns a tuple with the NicIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNicIds

`func (o *FiltersPublicIp) HasNicIds() bool`

HasNicIds returns a boolean if a field has been set.

### SetNicIds

`func (o *FiltersPublicIp) SetNicIds(v []string)`

SetNicIds gets a reference to the given []string and assigns it to the NicIds field.

### GetPlacements

`func (o *FiltersPublicIp) GetPlacements() []string`

GetPlacements returns the Placements field if non-nil, zero value otherwise.

### GetPlacementsOk

`func (o *FiltersPublicIp) GetPlacementsOk() ([]string, bool)`

GetPlacementsOk returns a tuple with the Placements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlacements

`func (o *FiltersPublicIp) HasPlacements() bool`

HasPlacements returns a boolean if a field has been set.

### SetPlacements

`func (o *FiltersPublicIp) SetPlacements(v []string)`

SetPlacements gets a reference to the given []string and assigns it to the Placements field.

### GetPrivateIps

`func (o *FiltersPublicIp) GetPrivateIps() []string`

GetPrivateIps returns the PrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsOk

`func (o *FiltersPublicIp) GetPrivateIpsOk() ([]string, bool)`

GetPrivateIpsOk returns a tuple with the PrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateIps

`func (o *FiltersPublicIp) HasPrivateIps() bool`

HasPrivateIps returns a boolean if a field has been set.

### SetPrivateIps

`func (o *FiltersPublicIp) SetPrivateIps(v []string)`

SetPrivateIps gets a reference to the given []string and assigns it to the PrivateIps field.

### GetPublicIpIds

`func (o *FiltersPublicIp) GetPublicIpIds() []string`

GetPublicIpIds returns the PublicIpIds field if non-nil, zero value otherwise.

### GetPublicIpIdsOk

`func (o *FiltersPublicIp) GetPublicIpIdsOk() ([]string, bool)`

GetPublicIpIdsOk returns a tuple with the PublicIpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicIpIds

`func (o *FiltersPublicIp) HasPublicIpIds() bool`

HasPublicIpIds returns a boolean if a field has been set.

### SetPublicIpIds

`func (o *FiltersPublicIp) SetPublicIpIds(v []string)`

SetPublicIpIds gets a reference to the given []string and assigns it to the PublicIpIds field.

### GetPublicIps

`func (o *FiltersPublicIp) GetPublicIps() []string`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *FiltersPublicIp) GetPublicIpsOk() ([]string, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicIps

`func (o *FiltersPublicIp) HasPublicIps() bool`

HasPublicIps returns a boolean if a field has been set.

### SetPublicIps

`func (o *FiltersPublicIp) SetPublicIps(v []string)`

SetPublicIps gets a reference to the given []string and assigns it to the PublicIps field.

### GetTagKeys

`func (o *FiltersPublicIp) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersPublicIp) GetTagKeysOk() ([]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagKeys

`func (o *FiltersPublicIp) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### SetTagKeys

`func (o *FiltersPublicIp) SetTagKeys(v []string)`

SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.

### GetTagValues

`func (o *FiltersPublicIp) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersPublicIp) GetTagValuesOk() ([]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagValues

`func (o *FiltersPublicIp) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### SetTagValues

`func (o *FiltersPublicIp) SetTagValues(v []string)`

SetTagValues gets a reference to the given []string and assigns it to the TagValues field.

### GetTags

`func (o *FiltersPublicIp) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersPublicIp) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *FiltersPublicIp) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *FiltersPublicIp) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetVmIds

`func (o *FiltersPublicIp) GetVmIds() []string`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *FiltersPublicIp) GetVmIdsOk() ([]string, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmIds

`func (o *FiltersPublicIp) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.

### SetVmIds

`func (o *FiltersPublicIp) SetVmIds(v []string)`

SetVmIds gets a reference to the given []string and assigns it to the VmIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

