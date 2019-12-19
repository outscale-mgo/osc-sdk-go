# FiltersSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAliases** | Pointer to **[]string** | The account aliases of the owners of the snapshots. | [optional] 
**AccountIds** | Pointer to **[]string** | The account IDs of the owners of the snapshots. | [optional] 
**Descriptions** | Pointer to **[]string** | The descriptions of the snapshots. | [optional] 
**PermissionsToCreateVolumeAccountIds** | Pointer to **[]string** | The account IDs of one or more users who have permissions to create volumes. | [optional] 
**PermissionsToCreateVolumeGlobalPermission** | Pointer to **bool** | If &#x60;true&#x60;, lists all public volumes. If &#x60;false&#x60;, lists all private volumes. | [optional] 
**Progresses** | Pointer to **[]int64** | The progresses of the snapshots, as a percentage. | [optional] 
**SnapshotIds** | Pointer to **[]string** | The IDs of the snapshots. | [optional] 
**States** | Pointer to **[]string** | The states of the snapshots (&#x60;in-queue&#x60; \\| &#x60;pending&#x60; \\| &#x60;completed&#x60;). | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the snapshots. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the snapshots. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the snapshots, in the following format: \&quot;Filters\&quot;:{\&quot;Tags\&quot;:[\&quot;TAGKEY&#x3D;TAGVALUE\&quot;]}. | [optional] 
**VolumeIds** | Pointer to **[]string** | The IDs of the volumes used to create the snapshots. | [optional] 
**VolumeSizes** | Pointer to **[]int64** | The sizes of the volumes used to create the snapshots, in gibibytes (GiB). | [optional] 

## Methods

### GetAccountAliases

`func (o *FiltersSnapshot) GetAccountAliases() []string`

GetAccountAliases returns the AccountAliases field if non-nil, zero value otherwise.

### GetAccountAliasesOk

`func (o *FiltersSnapshot) GetAccountAliasesOk() ([]string, bool)`

GetAccountAliasesOk returns a tuple with the AccountAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountAliases

`func (o *FiltersSnapshot) HasAccountAliases() bool`

HasAccountAliases returns a boolean if a field has been set.

### SetAccountAliases

`func (o *FiltersSnapshot) SetAccountAliases(v []string)`

SetAccountAliases gets a reference to the given []string and assigns it to the AccountAliases field.

### GetAccountIds

`func (o *FiltersSnapshot) GetAccountIds() []string`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *FiltersSnapshot) GetAccountIdsOk() ([]string, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountIds

`func (o *FiltersSnapshot) HasAccountIds() bool`

HasAccountIds returns a boolean if a field has been set.

### SetAccountIds

`func (o *FiltersSnapshot) SetAccountIds(v []string)`

SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.

### GetDescriptions

`func (o *FiltersSnapshot) GetDescriptions() []string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *FiltersSnapshot) GetDescriptionsOk() ([]string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescriptions

`func (o *FiltersSnapshot) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### SetDescriptions

`func (o *FiltersSnapshot) SetDescriptions(v []string)`

SetDescriptions gets a reference to the given []string and assigns it to the Descriptions field.

### GetPermissionsToCreateVolumeAccountIds

`func (o *FiltersSnapshot) GetPermissionsToCreateVolumeAccountIds() []string`

GetPermissionsToCreateVolumeAccountIds returns the PermissionsToCreateVolumeAccountIds field if non-nil, zero value otherwise.

### GetPermissionsToCreateVolumeAccountIdsOk

`func (o *FiltersSnapshot) GetPermissionsToCreateVolumeAccountIdsOk() ([]string, bool)`

GetPermissionsToCreateVolumeAccountIdsOk returns a tuple with the PermissionsToCreateVolumeAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPermissionsToCreateVolumeAccountIds

`func (o *FiltersSnapshot) HasPermissionsToCreateVolumeAccountIds() bool`

HasPermissionsToCreateVolumeAccountIds returns a boolean if a field has been set.

### SetPermissionsToCreateVolumeAccountIds

`func (o *FiltersSnapshot) SetPermissionsToCreateVolumeAccountIds(v []string)`

SetPermissionsToCreateVolumeAccountIds gets a reference to the given []string and assigns it to the PermissionsToCreateVolumeAccountIds field.

### GetPermissionsToCreateVolumeGlobalPermission

`func (o *FiltersSnapshot) GetPermissionsToCreateVolumeGlobalPermission() bool`

GetPermissionsToCreateVolumeGlobalPermission returns the PermissionsToCreateVolumeGlobalPermission field if non-nil, zero value otherwise.

### GetPermissionsToCreateVolumeGlobalPermissionOk

`func (o *FiltersSnapshot) GetPermissionsToCreateVolumeGlobalPermissionOk() (bool, bool)`

GetPermissionsToCreateVolumeGlobalPermissionOk returns a tuple with the PermissionsToCreateVolumeGlobalPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPermissionsToCreateVolumeGlobalPermission

`func (o *FiltersSnapshot) HasPermissionsToCreateVolumeGlobalPermission() bool`

HasPermissionsToCreateVolumeGlobalPermission returns a boolean if a field has been set.

### SetPermissionsToCreateVolumeGlobalPermission

`func (o *FiltersSnapshot) SetPermissionsToCreateVolumeGlobalPermission(v bool)`

SetPermissionsToCreateVolumeGlobalPermission gets a reference to the given bool and assigns it to the PermissionsToCreateVolumeGlobalPermission field.

### GetProgresses

`func (o *FiltersSnapshot) GetProgresses() []int64`

GetProgresses returns the Progresses field if non-nil, zero value otherwise.

### GetProgressesOk

`func (o *FiltersSnapshot) GetProgressesOk() ([]int64, bool)`

GetProgressesOk returns a tuple with the Progresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgresses

`func (o *FiltersSnapshot) HasProgresses() bool`

HasProgresses returns a boolean if a field has been set.

### SetProgresses

`func (o *FiltersSnapshot) SetProgresses(v []int64)`

SetProgresses gets a reference to the given []int64 and assigns it to the Progresses field.

### GetSnapshotIds

`func (o *FiltersSnapshot) GetSnapshotIds() []string`

GetSnapshotIds returns the SnapshotIds field if non-nil, zero value otherwise.

### GetSnapshotIdsOk

`func (o *FiltersSnapshot) GetSnapshotIdsOk() ([]string, bool)`

GetSnapshotIdsOk returns a tuple with the SnapshotIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSnapshotIds

`func (o *FiltersSnapshot) HasSnapshotIds() bool`

HasSnapshotIds returns a boolean if a field has been set.

### SetSnapshotIds

`func (o *FiltersSnapshot) SetSnapshotIds(v []string)`

SetSnapshotIds gets a reference to the given []string and assigns it to the SnapshotIds field.

### GetStates

`func (o *FiltersSnapshot) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersSnapshot) GetStatesOk() ([]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStates

`func (o *FiltersSnapshot) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStates

`func (o *FiltersSnapshot) SetStates(v []string)`

SetStates gets a reference to the given []string and assigns it to the States field.

### GetTagKeys

`func (o *FiltersSnapshot) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersSnapshot) GetTagKeysOk() ([]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagKeys

`func (o *FiltersSnapshot) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### SetTagKeys

`func (o *FiltersSnapshot) SetTagKeys(v []string)`

SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.

### GetTagValues

`func (o *FiltersSnapshot) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersSnapshot) GetTagValuesOk() ([]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagValues

`func (o *FiltersSnapshot) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### SetTagValues

`func (o *FiltersSnapshot) SetTagValues(v []string)`

SetTagValues gets a reference to the given []string and assigns it to the TagValues field.

### GetTags

`func (o *FiltersSnapshot) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersSnapshot) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *FiltersSnapshot) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *FiltersSnapshot) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetVolumeIds

`func (o *FiltersSnapshot) GetVolumeIds() []string`

GetVolumeIds returns the VolumeIds field if non-nil, zero value otherwise.

### GetVolumeIdsOk

`func (o *FiltersSnapshot) GetVolumeIdsOk() ([]string, bool)`

GetVolumeIdsOk returns a tuple with the VolumeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeIds

`func (o *FiltersSnapshot) HasVolumeIds() bool`

HasVolumeIds returns a boolean if a field has been set.

### SetVolumeIds

`func (o *FiltersSnapshot) SetVolumeIds(v []string)`

SetVolumeIds gets a reference to the given []string and assigns it to the VolumeIds field.

### GetVolumeSizes

`func (o *FiltersSnapshot) GetVolumeSizes() []int64`

GetVolumeSizes returns the VolumeSizes field if non-nil, zero value otherwise.

### GetVolumeSizesOk

`func (o *FiltersSnapshot) GetVolumeSizesOk() ([]int64, bool)`

GetVolumeSizesOk returns a tuple with the VolumeSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeSizes

`func (o *FiltersSnapshot) HasVolumeSizes() bool`

HasVolumeSizes returns a boolean if a field has been set.

### SetVolumeSizes

`func (o *FiltersSnapshot) SetVolumeSizes(v []int64)`

SetVolumeSizes gets a reference to the given []int64 and assigns it to the VolumeSizes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


