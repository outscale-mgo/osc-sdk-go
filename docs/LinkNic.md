# LinkNic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | If &#x60;true&#x60;, the volume is deleted when the VM is terminated. | [optional] 
**DeviceNumber** | Pointer to **int64** | The device index for the NIC attachment (between 1 and 7, both included). | [optional] 
**LinkNicId** | Pointer to **string** | The ID of the NIC to attach. | [optional] 
**State** | Pointer to **string** | The state of the attachment (&#x60;attaching&#x60; \\| &#x60;attached&#x60; \\| &#x60;detaching&#x60; \\| &#x60;detached&#x60;). | [optional] 
**VmAccountId** | Pointer to **string** | The account ID of the owner of the VM. | [optional] 
**VmId** | Pointer to **string** | The ID of the VM. | [optional] 

## Methods

### GetDeleteOnVmDeletion

`func (o *LinkNic) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *LinkNic) GetDeleteOnVmDeletionOk() (bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeleteOnVmDeletion

`func (o *LinkNic) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### SetDeleteOnVmDeletion

`func (o *LinkNic) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.

### GetDeviceNumber

`func (o *LinkNic) GetDeviceNumber() int64`

GetDeviceNumber returns the DeviceNumber field if non-nil, zero value otherwise.

### GetDeviceNumberOk

`func (o *LinkNic) GetDeviceNumberOk() (int64, bool)`

GetDeviceNumberOk returns a tuple with the DeviceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceNumber

`func (o *LinkNic) HasDeviceNumber() bool`

HasDeviceNumber returns a boolean if a field has been set.

### SetDeviceNumber

`func (o *LinkNic) SetDeviceNumber(v int64)`

SetDeviceNumber gets a reference to the given int64 and assigns it to the DeviceNumber field.

### GetLinkNicId

`func (o *LinkNic) GetLinkNicId() string`

GetLinkNicId returns the LinkNicId field if non-nil, zero value otherwise.

### GetLinkNicIdOk

`func (o *LinkNic) GetLinkNicIdOk() (string, bool)`

GetLinkNicIdOk returns a tuple with the LinkNicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkNicId

`func (o *LinkNic) HasLinkNicId() bool`

HasLinkNicId returns a boolean if a field has been set.

### SetLinkNicId

`func (o *LinkNic) SetLinkNicId(v string)`

SetLinkNicId gets a reference to the given string and assigns it to the LinkNicId field.

### GetState

`func (o *LinkNic) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LinkNic) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *LinkNic) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *LinkNic) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetVmAccountId

`func (o *LinkNic) GetVmAccountId() string`

GetVmAccountId returns the VmAccountId field if non-nil, zero value otherwise.

### GetVmAccountIdOk

`func (o *LinkNic) GetVmAccountIdOk() (string, bool)`

GetVmAccountIdOk returns a tuple with the VmAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmAccountId

`func (o *LinkNic) HasVmAccountId() bool`

HasVmAccountId returns a boolean if a field has been set.

### SetVmAccountId

`func (o *LinkNic) SetVmAccountId(v string)`

SetVmAccountId gets a reference to the given string and assigns it to the VmAccountId field.

### GetVmId

`func (o *LinkNic) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *LinkNic) GetVmIdOk() (string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmId

`func (o *LinkNic) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### SetVmId

`func (o *LinkNic) SetVmId(v string)`

SetVmId gets a reference to the given string and assigns it to the VmId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


