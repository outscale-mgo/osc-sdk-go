# LinkNicToUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnVmDeletion** | Pointer to **bool** | If &#x60;true&#x60;, the NIC is deleted when the VM is terminated. | [optional] 
**LinkNicId** | Pointer to **string** | The ID of the NIC attachment. | [optional] 

## Methods

### GetDeleteOnVmDeletion

`func (o *LinkNicToUpdate) GetDeleteOnVmDeletion() bool`

GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetDeleteOnVmDeletionOk

`func (o *LinkNicToUpdate) GetDeleteOnVmDeletionOk() (bool, bool)`

GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeleteOnVmDeletion

`func (o *LinkNicToUpdate) HasDeleteOnVmDeletion() bool`

HasDeleteOnVmDeletion returns a boolean if a field has been set.

### SetDeleteOnVmDeletion

`func (o *LinkNicToUpdate) SetDeleteOnVmDeletion(v bool)`

SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.

### GetLinkNicId

`func (o *LinkNicToUpdate) GetLinkNicId() string`

GetLinkNicId returns the LinkNicId field if non-nil, zero value otherwise.

### GetLinkNicIdOk

`func (o *LinkNicToUpdate) GetLinkNicIdOk() (string, bool)`

GetLinkNicIdOk returns a tuple with the LinkNicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkNicId

`func (o *LinkNicToUpdate) HasLinkNicId() bool`

HasLinkNicId returns a boolean if a field has been set.

### SetLinkNicId

`func (o *LinkNicToUpdate) SetLinkNicId(v string)`

SetLinkNicId gets a reference to the given string and assigns it to the LinkNicId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


