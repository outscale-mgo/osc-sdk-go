# CreateDhcpOptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** | Specify a domain name (for example, MyCompany.com). You can specify only one domain name. | [optional] 
**DomainNameServers** | Pointer to **[]string** | The IP addresses of domain name servers. If no IP addresses are specified, the &#x60;OutscaleProvidedDNS&#x60; value is set by default. | [optional] 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**NtpServers** | Pointer to **[]string** | The IP addresses of the Network Time Protocol (NTP) servers. | [optional] 

## Methods

### GetDomainName

`func (o *CreateDhcpOptionsRequest) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CreateDhcpOptionsRequest) GetDomainNameOk() (string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDomainName

`func (o *CreateDhcpOptionsRequest) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainName

`func (o *CreateDhcpOptionsRequest) SetDomainName(v string)`

SetDomainName gets a reference to the given string and assigns it to the DomainName field.

### GetDomainNameServers

`func (o *CreateDhcpOptionsRequest) GetDomainNameServers() []string`

GetDomainNameServers returns the DomainNameServers field if non-nil, zero value otherwise.

### GetDomainNameServersOk

`func (o *CreateDhcpOptionsRequest) GetDomainNameServersOk() ([]string, bool)`

GetDomainNameServersOk returns a tuple with the DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDomainNameServers

`func (o *CreateDhcpOptionsRequest) HasDomainNameServers() bool`

HasDomainNameServers returns a boolean if a field has been set.

### SetDomainNameServers

`func (o *CreateDhcpOptionsRequest) SetDomainNameServers(v []string)`

SetDomainNameServers gets a reference to the given []string and assigns it to the DomainNameServers field.

### GetDryRun

`func (o *CreateDhcpOptionsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateDhcpOptionsRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateDhcpOptionsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateDhcpOptionsRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetNtpServers

`func (o *CreateDhcpOptionsRequest) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *CreateDhcpOptionsRequest) GetNtpServersOk() ([]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNtpServers

`func (o *CreateDhcpOptionsRequest) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### SetNtpServers

`func (o *CreateDhcpOptionsRequest) SetNtpServers(v []string)`

SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


