# DhcpOptionsSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **bool** | If &#x60;true&#x60;, the DHCP options set is a default one. If &#x60;false&#x60;, it is not. | [optional] 
**DhcpOptionsName** | Pointer to **string** | The name of the DHCP options set. | [optional] 
**DhcpOptionsSetId** | Pointer to **string** | The ID of the DHCP options set. | [optional] 
**DomainName** | Pointer to **string** | The domain name. | [optional] 
**DomainNameServers** | Pointer to **[]string** | One or more IP addresses for the domain name servers. | [optional] 
**NtpServers** | Pointer to **[]string** | One or more IP addresses for the NTP servers. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the DHCP options set. | [optional] 

## Methods

### GetDefault

`func (o *DhcpOptionsSet) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DhcpOptionsSet) GetDefaultOk() (bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefault

`func (o *DhcpOptionsSet) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefault

`func (o *DhcpOptionsSet) SetDefault(v bool)`

SetDefault gets a reference to the given bool and assigns it to the Default field.

### GetDhcpOptionsName

`func (o *DhcpOptionsSet) GetDhcpOptionsName() string`

GetDhcpOptionsName returns the DhcpOptionsName field if non-nil, zero value otherwise.

### GetDhcpOptionsNameOk

`func (o *DhcpOptionsSet) GetDhcpOptionsNameOk() (string, bool)`

GetDhcpOptionsNameOk returns a tuple with the DhcpOptionsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDhcpOptionsName

`func (o *DhcpOptionsSet) HasDhcpOptionsName() bool`

HasDhcpOptionsName returns a boolean if a field has been set.

### SetDhcpOptionsName

`func (o *DhcpOptionsSet) SetDhcpOptionsName(v string)`

SetDhcpOptionsName gets a reference to the given string and assigns it to the DhcpOptionsName field.

### GetDhcpOptionsSetId

`func (o *DhcpOptionsSet) GetDhcpOptionsSetId() string`

GetDhcpOptionsSetId returns the DhcpOptionsSetId field if non-nil, zero value otherwise.

### GetDhcpOptionsSetIdOk

`func (o *DhcpOptionsSet) GetDhcpOptionsSetIdOk() (string, bool)`

GetDhcpOptionsSetIdOk returns a tuple with the DhcpOptionsSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDhcpOptionsSetId

`func (o *DhcpOptionsSet) HasDhcpOptionsSetId() bool`

HasDhcpOptionsSetId returns a boolean if a field has been set.

### SetDhcpOptionsSetId

`func (o *DhcpOptionsSet) SetDhcpOptionsSetId(v string)`

SetDhcpOptionsSetId gets a reference to the given string and assigns it to the DhcpOptionsSetId field.

### GetDomainName

`func (o *DhcpOptionsSet) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DhcpOptionsSet) GetDomainNameOk() (string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDomainName

`func (o *DhcpOptionsSet) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainName

`func (o *DhcpOptionsSet) SetDomainName(v string)`

SetDomainName gets a reference to the given string and assigns it to the DomainName field.

### GetDomainNameServers

`func (o *DhcpOptionsSet) GetDomainNameServers() []string`

GetDomainNameServers returns the DomainNameServers field if non-nil, zero value otherwise.

### GetDomainNameServersOk

`func (o *DhcpOptionsSet) GetDomainNameServersOk() ([]string, bool)`

GetDomainNameServersOk returns a tuple with the DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDomainNameServers

`func (o *DhcpOptionsSet) HasDomainNameServers() bool`

HasDomainNameServers returns a boolean if a field has been set.

### SetDomainNameServers

`func (o *DhcpOptionsSet) SetDomainNameServers(v []string)`

SetDomainNameServers gets a reference to the given []string and assigns it to the DomainNameServers field.

### GetNtpServers

`func (o *DhcpOptionsSet) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *DhcpOptionsSet) GetNtpServersOk() ([]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNtpServers

`func (o *DhcpOptionsSet) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### SetNtpServers

`func (o *DhcpOptionsSet) SetNtpServers(v []string)`

SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.

### GetTags

`func (o *DhcpOptionsSet) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DhcpOptionsSet) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *DhcpOptionsSet) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *DhcpOptionsSet) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


