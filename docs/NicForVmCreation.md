# NicForVmCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | If &#x60;true&#x60;, the NIC is deleted when the VM is terminated. You can specify &#x60;true&#x60; only if you create a NIC when creating a VM. | [optional] 
**Description** | Pointer to **string** | The description of the NIC, if you are creating a NIC when creating the VM. | [optional] 
**DeviceNumber** | Pointer to **int64** | The index of the VM device for the NIC attachment (between 0 and 7, both included). This parameter is required if you create a NIC when creating the VM. | [optional] 
**NicId** | Pointer to **string** | The ID of the NIC, if you are attaching an existing NIC when creating a VM. | [optional] 
**PrivateIps** | Pointer to [**[]PrivateIpLight**](PrivateIpLight.md) | One or more private IP addresses to assign to the NIC, if you create a NIC when creating a VM. Only one private IP address can be the primary private IP address. | [optional] 
**SecondaryPrivateIpCount** | Pointer to **int64** | The number of secondary private IP addresses, if you create a NIC when creating a VM. This parameter cannot be specified if you specified more than one private IP address in the &#x60;PrivateIps&#x60; parameter. | [optional] 
**SecurityGroupIds** | Pointer to **[]string** | One or more IDs of security groups for the NIC, if you acreate a NIC when creating a VM. | [optional] 
**SubnetId** | Pointer to **string** | The ID of the Subnet for the NIC, if you create a NIC when creating a VM. | [optional] 

## Methods

### GetDeleteOnVmDeletion

`func (o *NicForVmCreation) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *NicForVmCreation) GetDeleteOnVmDeletionOk() (bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeleteOnVmDeletion

`func (o *NicForVmCreation) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### SetDeleteOnVmDeletion

`func (o *NicForVmCreation) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.

### GetDescription

`func (o *NicForVmCreation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NicForVmCreation) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *NicForVmCreation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *NicForVmCreation) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetDeviceNumber

`func (o *NicForVmCreation) GetDeviceNumber() int64`

GetDeviceNumber returns the DeviceNumber field if non-nil, zero value otherwise.

### GetDeviceNumberOk

`func (o *NicForVmCreation) GetDeviceNumberOk() (int64, bool)`

GetDeviceNumberOk returns a tuple with the DeviceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceNumber

`func (o *NicForVmCreation) HasDeviceNumber() bool`

HasDeviceNumber returns a boolean if a field has been set.

### SetDeviceNumber

`func (o *NicForVmCreation) SetDeviceNumber(v int64)`

SetDeviceNumber gets a reference to the given int64 and assigns it to the DeviceNumber field.

### GetNicId

`func (o *NicForVmCreation) GetNicId() string`

GetNicId returns the NicId field if non-nil, zero value otherwise.

### GetNicIdOk

`func (o *NicForVmCreation) GetNicIdOk() (string, bool)`

GetNicIdOk returns a tuple with the NicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNicId

`func (o *NicForVmCreation) HasNicId() bool`

HasNicId returns a boolean if a field has been set.

### SetNicId

`func (o *NicForVmCreation) SetNicId(v string)`

SetNicId gets a reference to the given string and assigns it to the NicId field.

### GetPrivateIps

`func (o *NicForVmCreation) GetPrivateIps() []PrivateIpLight`

GetPrivateIps returns the PrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsOk

`func (o *NicForVmCreation) GetPrivateIpsOk() ([]PrivateIpLight, bool)`

GetPrivateIpsOk returns a tuple with the PrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateIps

`func (o *NicForVmCreation) HasPrivateIps() bool`

HasPrivateIps returns a boolean if a field has been set.

### SetPrivateIps

`func (o *NicForVmCreation) SetPrivateIps(v []PrivateIpLight)`

SetPrivateIps gets a reference to the given []PrivateIpLight and assigns it to the PrivateIps field.

### GetSecondaryPrivateIpCount

`func (o *NicForVmCreation) GetSecondaryPrivateIpCount() int64`

GetSecondaryPrivateIpCount returns the SecondaryPrivateIpCount field if non-nil, zero value otherwise.

### GetSecondaryPrivateIpCountOk

`func (o *NicForVmCreation) GetSecondaryPrivateIpCountOk() (int64, bool)`

GetSecondaryPrivateIpCountOk returns a tuple with the SecondaryPrivateIpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecondaryPrivateIpCount

`func (o *NicForVmCreation) HasSecondaryPrivateIpCount() bool`

HasSecondaryPrivateIpCount returns a boolean if a field has been set.

### SetSecondaryPrivateIpCount

`func (o *NicForVmCreation) SetSecondaryPrivateIpCount(v int64)`

SetSecondaryPrivateIpCount gets a reference to the given int64 and assigns it to the SecondaryPrivateIpCount field.

### GetSecurityGroupIds

`func (o *NicForVmCreation) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *NicForVmCreation) GetSecurityGroupIdsOk() ([]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroupIds

`func (o *NicForVmCreation) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### SetSecurityGroupIds

`func (o *NicForVmCreation) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds gets a reference to the given []string and assigns it to the SecurityGroupIds field.

### GetSubnetId

`func (o *NicForVmCreation) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *NicForVmCreation) GetSubnetIdOk() (string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubnetId

`func (o *NicForVmCreation) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetId

`func (o *NicForVmCreation) SetSubnetId(v string)`

SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


