# FiltersImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAliases** | Pointer to **[]string** | The account aliases of the owners of the OMIs. | [optional] 
**AccountIds** | Pointer to **[]string** | The account IDs of the owners of the OMIs. By default, all the OMIs for which you have launch permissions are described. | [optional] 
**Architectures** | Pointer to **[]string** | The architectures of the OMIs (&#x60;i386&#x60; \\| &#x60;x86_64&#x60;). | [optional] 
**BlockDeviceMappingDeleteOnVmDeletion** | Pointer to **bool** | Indicates whether the block device mapping is deleted when terminating the VM. | [optional] 
**BlockDeviceMappingDeviceNames** | Pointer to **[]string** | The device names for the volumes. | [optional] 
**BlockDeviceMappingSnapshotIds** | Pointer to **[]string** | The IDs of the snapshots used to create the volumes. | [optional] 
**BlockDeviceMappingVolumeSizes** | Pointer to **[]int64** | The sizes of the volumes, in gibibytes (GiB). | [optional] 
**BlockDeviceMappingVolumeTypes** | Pointer to **[]string** | The types of volumes (&#x60;standard&#x60; \\| &#x60;gp2&#x60; \\| &#x60;io1&#x60;). | [optional] 
**Descriptions** | Pointer to **[]string** | The descriptions of the OMIs, provided when they were created. | [optional] 
**FileLocations** | Pointer to **[]string** | The locations where the OMI files are stored on Object Storage Unit (OSU). | [optional] 
**ImageIds** | Pointer to **[]string** | The IDs of the OMIs. | [optional] 
**ImageNames** | Pointer to **[]string** | The names of the OMIs, provided when they were created. | [optional] 
**PermissionsToLaunchAccountIds** | Pointer to **[]string** | The account IDs of the users who have launch permissions for the OMIs. | [optional] 
**PermissionsToLaunchGlobalPermission** | Pointer to **bool** | If &#x60;true&#x60;, lists all public OMIs. If &#x60;false&#x60;, lists all private OMIs. | [optional] 
**RootDeviceNames** | Pointer to **[]string** | The device names of the root devices (for example, &#x60;/dev/sda1&#x60;). | [optional] 
**RootDeviceTypes** | Pointer to **[]string** | The types of root device used by the OMIs (always &#x60;ebs&#x60;). | [optional] 
**States** | Pointer to **[]string** | The states of the OMIs. | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the OMIs. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the OMIs. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the OMIs, in the following format: \&quot;Filters\&quot;:{\&quot;Tags\&quot;:[\&quot;TAGKEY&#x3D;TAGVALUE\&quot;]}. | [optional] 
**VirtualizationTypes** | Pointer to **[]string** | The virtualization types (always &#x60;hvm&#x60;). | [optional] 

## Methods

### GetAccountAliases

`func (o *FiltersImage) GetAccountAliases() []string`

GetAccountAliases returns the AccountAliases field if non-nil, zero value otherwise.

### GetAccountAliasesOk

`func (o *FiltersImage) GetAccountAliasesOk() ([]string, bool)`

GetAccountAliasesOk returns a tuple with the AccountAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountAliases

`func (o *FiltersImage) HasAccountAliases() bool`

HasAccountAliases returns a boolean if a field has been set.

### SetAccountAliases

`func (o *FiltersImage) SetAccountAliases(v []string)`

SetAccountAliases gets a reference to the given []string and assigns it to the AccountAliases field.

### GetAccountIds

`func (o *FiltersImage) GetAccountIds() []string`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *FiltersImage) GetAccountIdsOk() ([]string, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountIds

`func (o *FiltersImage) HasAccountIds() bool`

HasAccountIds returns a boolean if a field has been set.

### SetAccountIds

`func (o *FiltersImage) SetAccountIds(v []string)`

SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.

### GetArchitectures

`func (o *FiltersImage) GetArchitectures() []string`

GetArchitectures returns the Architectures field if non-nil, zero value otherwise.

### GetArchitecturesOk

`func (o *FiltersImage) GetArchitecturesOk() ([]string, bool)`

GetArchitecturesOk returns a tuple with the Architectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasArchitectures

`func (o *FiltersImage) HasArchitectures() bool`

HasArchitectures returns a boolean if a field has been set.

### SetArchitectures

`func (o *FiltersImage) SetArchitectures(v []string)`

SetArchitectures gets a reference to the given []string and assigns it to the Architectures field.

### GetBlockDeviceMappingDeleteOnVmDeletion

`func (o *FiltersImage) GetBlockDeviceMappingDeleteOnVmDeletion() bool`

GetBlockDeviceMappingDeleteOnVmDeletion returns the BlockDeviceMappingDeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetBlockDeviceMappingDeleteOnVmDeletionOk

`func (o *FiltersImage) GetBlockDeviceMappingDeleteOnVmDeletionOk() (bool, bool)`

GetBlockDeviceMappingDeleteOnVmDeletionOk returns a tuple with the BlockDeviceMappingDeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBlockDeviceMappingDeleteOnVmDeletion

`func (o *FiltersImage) HasBlockDeviceMappingDeleteOnVmDeletion() bool`

HasBlockDeviceMappingDeleteOnVmDeletion returns a boolean if a field has been set.

### SetBlockDeviceMappingDeleteOnVmDeletion

`func (o *FiltersImage) SetBlockDeviceMappingDeleteOnVmDeletion(v bool)`

SetBlockDeviceMappingDeleteOnVmDeletion gets a reference to the given bool and assigns it to the BlockDeviceMappingDeleteOnVmDeletion field.

### GetBlockDeviceMappingDeviceNames

`func (o *FiltersImage) GetBlockDeviceMappingDeviceNames() []string`

GetBlockDeviceMappingDeviceNames returns the BlockDeviceMappingDeviceNames field if non-nil, zero value otherwise.

### GetBlockDeviceMappingDeviceNamesOk

`func (o *FiltersImage) GetBlockDeviceMappingDeviceNamesOk() ([]string, bool)`

GetBlockDeviceMappingDeviceNamesOk returns a tuple with the BlockDeviceMappingDeviceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBlockDeviceMappingDeviceNames

`func (o *FiltersImage) HasBlockDeviceMappingDeviceNames() bool`

HasBlockDeviceMappingDeviceNames returns a boolean if a field has been set.

### SetBlockDeviceMappingDeviceNames

`func (o *FiltersImage) SetBlockDeviceMappingDeviceNames(v []string)`

SetBlockDeviceMappingDeviceNames gets a reference to the given []string and assigns it to the BlockDeviceMappingDeviceNames field.

### GetBlockDeviceMappingSnapshotIds

`func (o *FiltersImage) GetBlockDeviceMappingSnapshotIds() []string`

GetBlockDeviceMappingSnapshotIds returns the BlockDeviceMappingSnapshotIds field if non-nil, zero value otherwise.

### GetBlockDeviceMappingSnapshotIdsOk

`func (o *FiltersImage) GetBlockDeviceMappingSnapshotIdsOk() ([]string, bool)`

GetBlockDeviceMappingSnapshotIdsOk returns a tuple with the BlockDeviceMappingSnapshotIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBlockDeviceMappingSnapshotIds

`func (o *FiltersImage) HasBlockDeviceMappingSnapshotIds() bool`

HasBlockDeviceMappingSnapshotIds returns a boolean if a field has been set.

### SetBlockDeviceMappingSnapshotIds

`func (o *FiltersImage) SetBlockDeviceMappingSnapshotIds(v []string)`

SetBlockDeviceMappingSnapshotIds gets a reference to the given []string and assigns it to the BlockDeviceMappingSnapshotIds field.

### GetBlockDeviceMappingVolumeSizes

`func (o *FiltersImage) GetBlockDeviceMappingVolumeSizes() []int64`

GetBlockDeviceMappingVolumeSizes returns the BlockDeviceMappingVolumeSizes field if non-nil, zero value otherwise.

### GetBlockDeviceMappingVolumeSizesOk

`func (o *FiltersImage) GetBlockDeviceMappingVolumeSizesOk() ([]int64, bool)`

GetBlockDeviceMappingVolumeSizesOk returns a tuple with the BlockDeviceMappingVolumeSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBlockDeviceMappingVolumeSizes

`func (o *FiltersImage) HasBlockDeviceMappingVolumeSizes() bool`

HasBlockDeviceMappingVolumeSizes returns a boolean if a field has been set.

### SetBlockDeviceMappingVolumeSizes

`func (o *FiltersImage) SetBlockDeviceMappingVolumeSizes(v []int64)`

SetBlockDeviceMappingVolumeSizes gets a reference to the given []int64 and assigns it to the BlockDeviceMappingVolumeSizes field.

### GetBlockDeviceMappingVolumeTypes

`func (o *FiltersImage) GetBlockDeviceMappingVolumeTypes() []string`

GetBlockDeviceMappingVolumeTypes returns the BlockDeviceMappingVolumeTypes field if non-nil, zero value otherwise.

### GetBlockDeviceMappingVolumeTypesOk

`func (o *FiltersImage) GetBlockDeviceMappingVolumeTypesOk() ([]string, bool)`

GetBlockDeviceMappingVolumeTypesOk returns a tuple with the BlockDeviceMappingVolumeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBlockDeviceMappingVolumeTypes

`func (o *FiltersImage) HasBlockDeviceMappingVolumeTypes() bool`

HasBlockDeviceMappingVolumeTypes returns a boolean if a field has been set.

### SetBlockDeviceMappingVolumeTypes

`func (o *FiltersImage) SetBlockDeviceMappingVolumeTypes(v []string)`

SetBlockDeviceMappingVolumeTypes gets a reference to the given []string and assigns it to the BlockDeviceMappingVolumeTypes field.

### GetDescriptions

`func (o *FiltersImage) GetDescriptions() []string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *FiltersImage) GetDescriptionsOk() ([]string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescriptions

`func (o *FiltersImage) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### SetDescriptions

`func (o *FiltersImage) SetDescriptions(v []string)`

SetDescriptions gets a reference to the given []string and assigns it to the Descriptions field.

### GetFileLocations

`func (o *FiltersImage) GetFileLocations() []string`

GetFileLocations returns the FileLocations field if non-nil, zero value otherwise.

### GetFileLocationsOk

`func (o *FiltersImage) GetFileLocationsOk() ([]string, bool)`

GetFileLocationsOk returns a tuple with the FileLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileLocations

`func (o *FiltersImage) HasFileLocations() bool`

HasFileLocations returns a boolean if a field has been set.

### SetFileLocations

`func (o *FiltersImage) SetFileLocations(v []string)`

SetFileLocations gets a reference to the given []string and assigns it to the FileLocations field.

### GetImageIds

`func (o *FiltersImage) GetImageIds() []string`

GetImageIds returns the ImageIds field if non-nil, zero value otherwise.

### GetImageIdsOk

`func (o *FiltersImage) GetImageIdsOk() ([]string, bool)`

GetImageIdsOk returns a tuple with the ImageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImageIds

`func (o *FiltersImage) HasImageIds() bool`

HasImageIds returns a boolean if a field has been set.

### SetImageIds

`func (o *FiltersImage) SetImageIds(v []string)`

SetImageIds gets a reference to the given []string and assigns it to the ImageIds field.

### GetImageNames

`func (o *FiltersImage) GetImageNames() []string`

GetImageNames returns the ImageNames field if non-nil, zero value otherwise.

### GetImageNamesOk

`func (o *FiltersImage) GetImageNamesOk() ([]string, bool)`

GetImageNamesOk returns a tuple with the ImageNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImageNames

`func (o *FiltersImage) HasImageNames() bool`

HasImageNames returns a boolean if a field has been set.

### SetImageNames

`func (o *FiltersImage) SetImageNames(v []string)`

SetImageNames gets a reference to the given []string and assigns it to the ImageNames field.

### GetPermissionsToLaunchAccountIds

`func (o *FiltersImage) GetPermissionsToLaunchAccountIds() []string`

GetPermissionsToLaunchAccountIds returns the PermissionsToLaunchAccountIds field if non-nil, zero value otherwise.

### GetPermissionsToLaunchAccountIdsOk

`func (o *FiltersImage) GetPermissionsToLaunchAccountIdsOk() ([]string, bool)`

GetPermissionsToLaunchAccountIdsOk returns a tuple with the PermissionsToLaunchAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPermissionsToLaunchAccountIds

`func (o *FiltersImage) HasPermissionsToLaunchAccountIds() bool`

HasPermissionsToLaunchAccountIds returns a boolean if a field has been set.

### SetPermissionsToLaunchAccountIds

`func (o *FiltersImage) SetPermissionsToLaunchAccountIds(v []string)`

SetPermissionsToLaunchAccountIds gets a reference to the given []string and assigns it to the PermissionsToLaunchAccountIds field.

### GetPermissionsToLaunchGlobalPermission

`func (o *FiltersImage) GetPermissionsToLaunchGlobalPermission() bool`

GetPermissionsToLaunchGlobalPermission returns the PermissionsToLaunchGlobalPermission field if non-nil, zero value otherwise.

### GetPermissionsToLaunchGlobalPermissionOk

`func (o *FiltersImage) GetPermissionsToLaunchGlobalPermissionOk() (bool, bool)`

GetPermissionsToLaunchGlobalPermissionOk returns a tuple with the PermissionsToLaunchGlobalPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPermissionsToLaunchGlobalPermission

`func (o *FiltersImage) HasPermissionsToLaunchGlobalPermission() bool`

HasPermissionsToLaunchGlobalPermission returns a boolean if a field has been set.

### SetPermissionsToLaunchGlobalPermission

`func (o *FiltersImage) SetPermissionsToLaunchGlobalPermission(v bool)`

SetPermissionsToLaunchGlobalPermission gets a reference to the given bool and assigns it to the PermissionsToLaunchGlobalPermission field.

### GetRootDeviceNames

`func (o *FiltersImage) GetRootDeviceNames() []string`

GetRootDeviceNames returns the RootDeviceNames field if non-nil, zero value otherwise.

### GetRootDeviceNamesOk

`func (o *FiltersImage) GetRootDeviceNamesOk() ([]string, bool)`

GetRootDeviceNamesOk returns a tuple with the RootDeviceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRootDeviceNames

`func (o *FiltersImage) HasRootDeviceNames() bool`

HasRootDeviceNames returns a boolean if a field has been set.

### SetRootDeviceNames

`func (o *FiltersImage) SetRootDeviceNames(v []string)`

SetRootDeviceNames gets a reference to the given []string and assigns it to the RootDeviceNames field.

### GetRootDeviceTypes

`func (o *FiltersImage) GetRootDeviceTypes() []string`

GetRootDeviceTypes returns the RootDeviceTypes field if non-nil, zero value otherwise.

### GetRootDeviceTypesOk

`func (o *FiltersImage) GetRootDeviceTypesOk() ([]string, bool)`

GetRootDeviceTypesOk returns a tuple with the RootDeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRootDeviceTypes

`func (o *FiltersImage) HasRootDeviceTypes() bool`

HasRootDeviceTypes returns a boolean if a field has been set.

### SetRootDeviceTypes

`func (o *FiltersImage) SetRootDeviceTypes(v []string)`

SetRootDeviceTypes gets a reference to the given []string and assigns it to the RootDeviceTypes field.

### GetStates

`func (o *FiltersImage) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersImage) GetStatesOk() ([]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStates

`func (o *FiltersImage) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStates

`func (o *FiltersImage) SetStates(v []string)`

SetStates gets a reference to the given []string and assigns it to the States field.

### GetTagKeys

`func (o *FiltersImage) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersImage) GetTagKeysOk() ([]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagKeys

`func (o *FiltersImage) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### SetTagKeys

`func (o *FiltersImage) SetTagKeys(v []string)`

SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.

### GetTagValues

`func (o *FiltersImage) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersImage) GetTagValuesOk() ([]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagValues

`func (o *FiltersImage) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### SetTagValues

`func (o *FiltersImage) SetTagValues(v []string)`

SetTagValues gets a reference to the given []string and assigns it to the TagValues field.

### GetTags

`func (o *FiltersImage) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersImage) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *FiltersImage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *FiltersImage) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetVirtualizationTypes

`func (o *FiltersImage) GetVirtualizationTypes() []string`

GetVirtualizationTypes returns the VirtualizationTypes field if non-nil, zero value otherwise.

### GetVirtualizationTypesOk

`func (o *FiltersImage) GetVirtualizationTypesOk() ([]string, bool)`

GetVirtualizationTypesOk returns a tuple with the VirtualizationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVirtualizationTypes

`func (o *FiltersImage) HasVirtualizationTypes() bool`

HasVirtualizationTypes returns a boolean if a field has been set.

### SetVirtualizationTypes

`func (o *FiltersImage) SetVirtualizationTypes(v []string)`

SetVirtualizationTypes gets a reference to the given []string and assigns it to the VirtualizationTypes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


