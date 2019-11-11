# FiltersDhcpOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **bool** | If &#x60;true&#x60;, lists all default DHCP options set. If &#x60;false&#x60;, lists all non-default DHCP options set. | [optional] 
**DhcpOptionsSetIds** | Pointer to **[]string** | The IDs of the DHCP options sets. | [optional] 
**DomainNameServers** | Pointer to **[]string** | The domain name servers used for the DHCP options sets. | [optional] 
**DomainNames** | Pointer to **[]string** | The domain names used for the DHCP options sets. | [optional] 
**NtpServers** | Pointer to **[]string** | The Network Time Protocol (NTP) servers used for the DHCP options sets. | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the DHCP options sets. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the DHCP options sets. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the DHCP options sets, in the following format: \&quot;Filters\&quot;:{\&quot;Tags\&quot;:[\&quot;TAGKEY&#x3D;TAGVALUE\&quot;]}. | [optional] 

## Methods

### GetDefault

`func (o *FiltersDhcpOptions) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FiltersDhcpOptions) GetDefaultOk() (bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefault

`func (o *FiltersDhcpOptions) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefault

`func (o *FiltersDhcpOptions) SetDefault(v bool)`

SetDefault gets a reference to the given bool and assigns it to the Default field.

### GetDhcpOptionsSetIds

`func (o *FiltersDhcpOptions) GetDhcpOptionsSetIds() []string`

GetDhcpOptionsSetIds returns the DhcpOptionsSetIds field if non-nil, zero value otherwise.

### GetDhcpOptionsSetIdsOk

`func (o *FiltersDhcpOptions) GetDhcpOptionsSetIdsOk() ([]string, bool)`

GetDhcpOptionsSetIdsOk returns a tuple with the DhcpOptionsSetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDhcpOptionsSetIds

`func (o *FiltersDhcpOptions) HasDhcpOptionsSetIds() bool`

HasDhcpOptionsSetIds returns a boolean if a field has been set.

### SetDhcpOptionsSetIds

`func (o *FiltersDhcpOptions) SetDhcpOptionsSetIds(v []string)`

SetDhcpOptionsSetIds gets a reference to the given []string and assigns it to the DhcpOptionsSetIds field.

### GetDomainNameServers

`func (o *FiltersDhcpOptions) GetDomainNameServers() []string`

GetDomainNameServers returns the DomainNameServers field if non-nil, zero value otherwise.

### GetDomainNameServersOk

`func (o *FiltersDhcpOptions) GetDomainNameServersOk() ([]string, bool)`

GetDomainNameServersOk returns a tuple with the DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDomainNameServers

`func (o *FiltersDhcpOptions) HasDomainNameServers() bool`

HasDomainNameServers returns a boolean if a field has been set.

### SetDomainNameServers

`func (o *FiltersDhcpOptions) SetDomainNameServers(v []string)`

SetDomainNameServers gets a reference to the given []string and assigns it to the DomainNameServers field.

### GetDomainNames

`func (o *FiltersDhcpOptions) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *FiltersDhcpOptions) GetDomainNamesOk() ([]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDomainNames

`func (o *FiltersDhcpOptions) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### SetDomainNames

`func (o *FiltersDhcpOptions) SetDomainNames(v []string)`

SetDomainNames gets a reference to the given []string and assigns it to the DomainNames field.

### GetNtpServers

`func (o *FiltersDhcpOptions) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *FiltersDhcpOptions) GetNtpServersOk() ([]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNtpServers

`func (o *FiltersDhcpOptions) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### SetNtpServers

`func (o *FiltersDhcpOptions) SetNtpServers(v []string)`

SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.

### GetTagKeys

`func (o *FiltersDhcpOptions) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersDhcpOptions) GetTagKeysOk() ([]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagKeys

`func (o *FiltersDhcpOptions) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### SetTagKeys

`func (o *FiltersDhcpOptions) SetTagKeys(v []string)`

SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.

### GetTagValues

`func (o *FiltersDhcpOptions) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersDhcpOptions) GetTagValuesOk() ([]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagValues

`func (o *FiltersDhcpOptions) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### SetTagValues

`func (o *FiltersDhcpOptions) SetTagValues(v []string)`

SetTagValues gets a reference to the given []string and assigns it to the TagValues field.

### GetTags

`func (o *FiltersDhcpOptions) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersDhcpOptions) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *FiltersDhcpOptions) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *FiltersDhcpOptions) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


