# LinkNicLight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | If &#x60;true&#x60;, the volume is deleted when the VM is terminated. | [optional] 
**DeviceNumber** | Pointer to **int32** | The device index for the NIC attachment (between 1 and 7, both included). | [optional] 
**LinkNicId** | Pointer to **string** | The ID of the NIC to attach. | [optional] 
**State** | Pointer to **string** | The state of the attachment (&#x60;attaching&#x60; \\| &#x60;attached&#x60; \\| &#x60;detaching&#x60; \\| &#x60;detached&#x60;). | [optional] 

## Methods

### GetDeleteOnVmDeletion

`func (o *LinkNicLight) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *LinkNicLight) GetDeleteOnVmDeletionOk() (bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeleteOnVmDeletion

`func (o *LinkNicLight) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### SetDeleteOnVmDeletion

`func (o *LinkNicLight) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.

### GetDeviceNumber

`func (o *LinkNicLight) GetDeviceNumber() int32`

GetDeviceNumber returns the DeviceNumber field if non-nil, zero value otherwise.

### GetDeviceNumberOk

`func (o *LinkNicLight) GetDeviceNumberOk() (int32, bool)`

GetDeviceNumberOk returns a tuple with the DeviceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceNumber

`func (o *LinkNicLight) HasDeviceNumber() bool`

HasDeviceNumber returns a boolean if a field has been set.

### SetDeviceNumber

`func (o *LinkNicLight) SetDeviceNumber(v int32)`

SetDeviceNumber gets a reference to the given int32 and assigns it to the DeviceNumber field.

### GetLinkNicId

`func (o *LinkNicLight) GetLinkNicId() string`

GetLinkNicId returns the LinkNicId field if non-nil, zero value otherwise.

### GetLinkNicIdOk

`func (o *LinkNicLight) GetLinkNicIdOk() (string, bool)`

GetLinkNicIdOk returns a tuple with the LinkNicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkNicId

`func (o *LinkNicLight) HasLinkNicId() bool`

HasLinkNicId returns a boolean if a field has been set.

### SetLinkNicId

`func (o *LinkNicLight) SetLinkNicId(v string)`

SetLinkNicId gets a reference to the given string and assigns it to the LinkNicId field.

### GetState

`func (o *LinkNicLight) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LinkNicLight) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *LinkNicLight) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *LinkNicLight) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


