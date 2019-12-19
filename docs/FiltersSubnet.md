# FiltersSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableIpsCounts** | Pointer to **[]int64** | The number of available IPs. | [optional] 
**IpRanges** | Pointer to **[]string** | The IP ranges in the Subnets, in CIDR notation (for example, 10.0.0.0/16). | [optional] 
**NetIds** | Pointer to **[]string** | The IDs of the Nets in which the Subnets are. | [optional] 
**States** | Pointer to **[]string** | The states of the Subnets (&#x60;pending&#x60; \\| &#x60;available&#x60;). | [optional] 
**SubnetIds** | Pointer to **[]string** | The IDs of the Subnets. | [optional] 
**SubregionNames** | Pointer to **[]string** | The names of the Subregions in which the Subnets are located. | [optional] 

## Methods

### GetAvailableIpsCounts

`func (o *FiltersSubnet) GetAvailableIpsCounts() []int64`

GetAvailableIpsCounts returns the AvailableIpsCounts field if non-nil, zero value otherwise.

### GetAvailableIpsCountsOk

`func (o *FiltersSubnet) GetAvailableIpsCountsOk() ([]int64, bool)`

GetAvailableIpsCountsOk returns a tuple with the AvailableIpsCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAvailableIpsCounts

`func (o *FiltersSubnet) HasAvailableIpsCounts() bool`

HasAvailableIpsCounts returns a boolean if a field has been set.

### SetAvailableIpsCounts

`func (o *FiltersSubnet) SetAvailableIpsCounts(v []int64)`

SetAvailableIpsCounts gets a reference to the given []int64 and assigns it to the AvailableIpsCounts field.

### GetIpRanges

`func (o *FiltersSubnet) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *FiltersSubnet) GetIpRangesOk() ([]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIpRanges

`func (o *FiltersSubnet) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### SetIpRanges

`func (o *FiltersSubnet) SetIpRanges(v []string)`

SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.

### GetNetIds

`func (o *FiltersSubnet) GetNetIds() []string`

GetNetIds returns the NetIds field if non-nil, zero value otherwise.

### GetNetIdsOk

`func (o *FiltersSubnet) GetNetIdsOk() ([]string, bool)`

GetNetIdsOk returns a tuple with the NetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetIds

`func (o *FiltersSubnet) HasNetIds() bool`

HasNetIds returns a boolean if a field has been set.

### SetNetIds

`func (o *FiltersSubnet) SetNetIds(v []string)`

SetNetIds gets a reference to the given []string and assigns it to the NetIds field.

### GetStates

`func (o *FiltersSubnet) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersSubnet) GetStatesOk() ([]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStates

`func (o *FiltersSubnet) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStates

`func (o *FiltersSubnet) SetStates(v []string)`

SetStates gets a reference to the given []string and assigns it to the States field.

### GetSubnetIds

`func (o *FiltersSubnet) GetSubnetIds() []string`

GetSubnetIds returns the SubnetIds field if non-nil, zero value otherwise.

### GetSubnetIdsOk

`func (o *FiltersSubnet) GetSubnetIdsOk() ([]string, bool)`

GetSubnetIdsOk returns a tuple with the SubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubnetIds

`func (o *FiltersSubnet) HasSubnetIds() bool`

HasSubnetIds returns a boolean if a field has been set.

### SetSubnetIds

`func (o *FiltersSubnet) SetSubnetIds(v []string)`

SetSubnetIds gets a reference to the given []string and assigns it to the SubnetIds field.

### GetSubregionNames

`func (o *FiltersSubnet) GetSubregionNames() []string`

GetSubregionNames returns the SubregionNames field if non-nil, zero value otherwise.

### GetSubregionNamesOk

`func (o *FiltersSubnet) GetSubregionNamesOk() ([]string, bool)`

GetSubregionNamesOk returns a tuple with the SubregionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubregionNames

`func (o *FiltersSubnet) HasSubregionNames() bool`

HasSubregionNames returns a boolean if a field has been set.

### SetSubregionNames

`func (o *FiltersSubnet) SetSubregionNames(v []string)`

SetSubregionNames gets a reference to the given []string and assigns it to the SubregionNames field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


