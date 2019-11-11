# Net

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpOptionsSetId** | Pointer to **string** | The ID of the DHCP options set (or &#x60;default&#x60; if you want to associate the default one). | [optional] 
**IpRange** | Pointer to **string** | The IP range for the Net, in CIDR notation (for example, 10.0.0.0/16). | [optional] 
**NetId** | Pointer to **string** | The ID of the Net. | [optional] 
**State** | Pointer to **string** | The state of the Net (&#x60;pending&#x60; \\| &#x60;available&#x60;). | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the Net. | [optional] 
**Tenancy** | Pointer to **string** | The VM tenancy in a Net. | [optional] 

## Methods

### GetDhcpOptionsSetId

`func (o *Net) GetDhcpOptionsSetId() string`

GetDhcpOptionsSetId returns the DhcpOptionsSetId field if non-nil, zero value otherwise.

### GetDhcpOptionsSetIdOk

`func (o *Net) GetDhcpOptionsSetIdOk() (string, bool)`

GetDhcpOptionsSetIdOk returns a tuple with the DhcpOptionsSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDhcpOptionsSetId

`func (o *Net) HasDhcpOptionsSetId() bool`

HasDhcpOptionsSetId returns a boolean if a field has been set.

### SetDhcpOptionsSetId

`func (o *Net) SetDhcpOptionsSetId(v string)`

SetDhcpOptionsSetId gets a reference to the given string and assigns it to the DhcpOptionsSetId field.

### GetIpRange

`func (o *Net) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *Net) GetIpRangeOk() (string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIpRange

`func (o *Net) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### SetIpRange

`func (o *Net) SetIpRange(v string)`

SetIpRange gets a reference to the given string and assigns it to the IpRange field.

### GetNetId

`func (o *Net) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *Net) GetNetIdOk() (string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetId

`func (o *Net) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### SetNetId

`func (o *Net) SetNetId(v string)`

SetNetId gets a reference to the given string and assigns it to the NetId field.

### GetState

`func (o *Net) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Net) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *Net) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *Net) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetTags

`func (o *Net) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Net) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *Net) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *Net) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.

### GetTenancy

`func (o *Net) GetTenancy() string`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *Net) GetTenancyOk() (string, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTenancy

`func (o *Net) HasTenancy() bool`

HasTenancy returns a boolean if a field has been set.

### SetTenancy

`func (o *Net) SetTenancy(v string)`

SetTenancy gets a reference to the given string and assigns it to the Tenancy field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


