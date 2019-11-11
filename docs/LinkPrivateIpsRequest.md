# LinkPrivateIpsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRelink** | Pointer to **bool** | If &#x60;true&#x60;, allows an IP address that is already assigned to another NIC in the same Subnet to be assigned to the NIC you specified. | [optional] 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**NicId** | Pointer to **string** | The ID of the NIC. | 
**PrivateIps** | Pointer to **[]string** | The secondary private IP address or addresses you want to assign to the NIC within the IP address range of the Subnet. | [optional] 
**SecondaryPrivateIpCount** | Pointer to **int32** | The number of secondary private IP addresses to assign to the NIC. | [optional] 

## Methods

### GetAllowRelink

`func (o *LinkPrivateIpsRequest) GetAllowRelink() bool`

GetAllowRelink returns the AllowRelink field if non-nil, zero value otherwise.

### GetAllowRelinkOk

`func (o *LinkPrivateIpsRequest) GetAllowRelinkOk() (bool, bool)`

GetAllowRelinkOk returns a tuple with the AllowRelink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowRelink

`func (o *LinkPrivateIpsRequest) HasAllowRelink() bool`

HasAllowRelink returns a boolean if a field has been set.

### SetAllowRelink

`func (o *LinkPrivateIpsRequest) SetAllowRelink(v bool)`

SetAllowRelink gets a reference to the given bool and assigns it to the AllowRelink field.

### GetDryRun

`func (o *LinkPrivateIpsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *LinkPrivateIpsRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *LinkPrivateIpsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *LinkPrivateIpsRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetNicId

`func (o *LinkPrivateIpsRequest) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *LinkPrivateIpsRequest) GetNicIdOk() (string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNicId

`func (o *LinkPrivateIpsRequest) HasNicId() bool`

HasNicId returns a boolean if a field has been set.

### SetNicId

`func (o *LinkPrivateIpsRequest) SetNicId(v string)`

SetNicId gets a reference to the given string and assigns it to the NicId field.

### GetPrivateIps

`func (o *LinkPrivateIpsRequest) GetPrivateIps() []string`

GetPrivateIps returns the PrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsOk

`func (o *LinkPrivateIpsRequest) GetPrivateIpsOk() ([]string, bool)`

GetPrivateIpsOk returns a tuple with the PrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateIps

`func (o *LinkPrivateIpsRequest) HasPrivateIps() bool`

HasPrivateIps returns a boolean if a field has been set.

### SetPrivateIps

`func (o *LinkPrivateIpsRequest) SetPrivateIps(v []string)`

SetPrivateIps gets a reference to the given []string and assigns it to the PrivateIps field.

### GetSecondaryPrivateIpCount

`func (o *LinkPrivateIpsRequest) GetSecondaryPrivateIpCount() int32`

GetSecondaryPrivateIpCount returns the SecondaryPrivateIpCount field if non-nil, zero value otherwise.

### GetSecondaryPrivateIpCountOk

`func (o *LinkPrivateIpsRequest) GetSecondaryPrivateIpCountOk() (int32, bool)`

GetSecondaryPrivateIpCountOk returns a tuple with the SecondaryPrivateIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecondaryPrivateIpCount

`func (o *LinkPrivateIpsRequest) HasSecondaryPrivateIpCount() bool`

HasSecondaryPrivateIpCount returns a boolean if a field has been set.

### SetSecondaryPrivateIpCount

`func (o *LinkPrivateIpsRequest) SetSecondaryPrivateIpCount(v int32)`

SetSecondaryPrivateIpCount gets a reference to the given int32 and assigns it to the SecondaryPrivateIpCount field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


