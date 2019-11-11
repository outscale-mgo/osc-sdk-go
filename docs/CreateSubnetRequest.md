# CreateSubnetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**IpRange** | Pointer to **string** | The IP range in the Subnet, in CIDR notation (for example, 10.0.0.0/16). | 
**NetId** | Pointer to **string** | The ID of the Net for which you want to create a Subnet. | 
**SubregionName** | Pointer to **string** | The name of the Subregion in which you want to create the Subnet. | [optional] 

## Methods

### GetDryRun

`func (o *CreateSubnetRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateSubnetRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateSubnetRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateSubnetRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetIpRange

`func (o *CreateSubnetRequest) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *CreateSubnetRequest) GetIpRangeOk() (string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIpRange

`func (o *CreateSubnetRequest) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### SetIpRange

`func (o *CreateSubnetRequest) SetIpRange(v string)`

SetIpRange gets a reference to the given string and assigns it to the IpRange field.

### GetNetId

`func (o *CreateSubnetRequest) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *CreateSubnetRequest) GetNetIdOk() (string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetId

`func (o *CreateSubnetRequest) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### SetNetId

`func (o *CreateSubnetRequest) SetNetId(v string)`

SetNetId gets a reference to the given string and assigns it to the NetId field.

### GetSubregionName

`func (o *CreateSubnetRequest) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *CreateSubnetRequest) GetSubregionNameOk() (string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubregionName

`func (o *CreateSubnetRequest) HasSubregionName() bool`

HasSubregionName returns a boolean if a field has been set.

### SetSubregionName

`func (o *CreateSubnetRequest) SetSubregionName(v string)`

SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


