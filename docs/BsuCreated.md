# BsuCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | Set to &#x60;true&#x60; by default, which means that the volume is deleted when the VM is terminated. If set to &#x60;false&#x60;, the volume is not deleted when the VM is terminated. | [optional] 
**LinkDate** | Pointer to **string** | The time and date of attachment of the volume to the VM. | [optional] 
**State** | Pointer to **string** | The state of the volume. | [optional] 
**VolumeId** | Pointer to **string** | The ID of the volume. | [optional] 

## Methods

### GetDeleteOnVmDeletion

`func (o *BsuCreated) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *BsuCreated) GetDeleteOnVmDeletionOk() (bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeleteOnVmDeletion

`func (o *BsuCreated) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### SetDeleteOnVmDeletion

`func (o *BsuCreated) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.

### GetLinkDate

`func (o *BsuCreated) GetLinkDate() string`

GetLinkDate returns the LinkDate field if non-nil, zero value otherwise.

### GetLinkDateOk

`func (o *BsuCreated) GetLinkDateOk() (string, bool)`

GetLinkDateOk returns a tuple with the LinkDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkDate

`func (o *BsuCreated) HasLinkDate() bool`

HasLinkDate returns a boolean if a field has been set.

### SetLinkDate

`func (o *BsuCreated) SetLinkDate(v string)`

SetLinkDate gets a reference to the given string and assigns it to the LinkDate field.

### GetState

`func (o *BsuCreated) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BsuCreated) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *BsuCreated) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *BsuCreated) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetVolumeId

`func (o *BsuCreated) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *BsuCreated) GetVolumeIdOk() (string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeId

`func (o *BsuCreated) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### SetVolumeId

`func (o *BsuCreated) SetVolumeId(v string)`

SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


