# FiltersVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDates** | Pointer to **[]string** | The dates and times at which the volumes were created. | [optional] 
**LinkVolumeDeleteOnVmDeletion** | Pointer to **bool** | Indicates whether the volumes are deleted when terminating the VMs. | [optional] 
**LinkVolumeDeviceNames** | Pointer to **[]string** | The VM device names. | [optional] 
**LinkVolumeLinkDates** | Pointer to **[]string** | The dates and times at which the volumes were created. | [optional] 
**LinkVolumeLinkStates** | Pointer to **[]string** | The attachment states of the volumes (&#x60;attaching&#x60; \\| &#x60;detaching&#x60; \\| &#x60;attached&#x60; \\| &#x60;detached&#x60;). | [optional] 
**LinkVolumeVmIds** | Pointer to **[]string** | One or more IDs of VMs. | [optional] 
**SnapshotIds** | Pointer to **[]string** | The snapshots from which the volumes were created. | [optional] 
**SubregionNames** | Pointer to **[]string** | The names of the Subregions in which the volumes were created. | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the volumes. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the volumes. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the volumes, in the following format: \&quot;Filters\&quot;:{\&quot;Tags\&quot;:[\&quot;TAGKEY&#x3D;TAGVALUE\&quot;]}. | [optional] 
**VolumeIds** | Pointer to **[]string** | The IDs of the volumes. | [optional] 
**VolumeSizes** | Pointer to **[]int32** | The sizes of the volumes, in gibibytes (GiB). | [optional] 
**VolumeStates** | Pointer to **[]string** | The states of the volumes (&#x60;creating&#x60; \\| &#x60;available&#x60; \\| &#x60;in-use&#x60; \\| &#x60;deleting&#x60; \\| &#x60;error&#x60;). | [optional] 
**VolumeTypes** | Pointer to **[]string** | The types of the volumes (&#x60;standard&#x60; \\| &#x60;gp2&#x60; \\| &#x60;io1&#x60;). | [optional] 

## Methods

### GetCreationDates

`func (o *FiltersVolume) GetCreationDates() []string`

GetCreationDates returns the CreationDates field if non-nil, zero value otherwise.

### GetCreationDatesOk

`func (o *FiltersVolume) GetCreationDatesOk() ([]string, bool)`

GetCreationDatesOk returns a tuple with the CreationDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreationDates

`func (o *FiltersVolume) HasCreationDates() bool`

HasCreationDates returns a boolean if a field has been set.

### SetCreationDates

`func (o *FiltersVolume) SetCreationDates(v []string)`

SetCreationDates gets a reference to the given []string and assigns it to the CreationDates field.

### GetLinkVolumeDeleteOnVmDeletion

`func (o *FiltersVolume) GetLinkVolumeDeleteOnVmDeletion() bool`

GetLinkVolumeDeleteOnVmDeletion returns the LinkVolumeDeleteOnVmDeletion field if non-nil, zero value otherwise.

### GetLinkVolumeDeleteOnVmDeletionOk

`func (o *FiltersVolume) GetLinkVolumeDeleteOnVmDeletionOk() (bool, bool)`

GetLinkVolumeDeleteOnVmDeletionOk returns a tuple with the LinkVolumeDeleteOnVmDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkVolumeDeleteOnVmDeletion

`func (o *FiltersVolume) HasLinkVolumeDeleteOnVmDeletion() bool`

HasLinkVolumeDeleteOnVmDeletion returns a boolean if a field has been set.

### SetLinkVolumeDeleteOnVmDeletion

`func (o *FiltersVolume) SetLinkVolumeDeleteOnVmDeletion(v bool)`

SetLinkVolumeDeleteOnVmDeletion gets a reference to the given bool and assigns it to the LinkVolumeDeleteOnVmDeletion field.

### GetLinkVolumeDeviceNames

`func (o *FiltersVolume) GetLinkVolumeDeviceNames() []string`

GetLinkVolumeDeviceNames returns the LinkVolumeDeviceNames field if non-nil, zero value otherwise.

### GetLinkVolumeDeviceNamesOk

`func (o *FiltersVolume) GetLinkVolumeDeviceNamesOk() ([]string, bool)`

GetLinkVolumeDeviceNamesOk returns a tuple with the LinkVolumeDeviceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkVolumeDeviceNames

`func (o *FiltersVolume) HasLinkVolumeDeviceNames() bool`

HasLinkVolumeDeviceNames returns a boolean if a field has been set.

### SetLinkVolumeDeviceNames

`func (o *FiltersVolume) SetLinkVolumeDeviceNames(v []string)`

SetLinkVolumeDeviceNames gets a reference to the given []string and assigns it to the LinkVolumeDeviceNames field.

### GetLinkVolumeLinkDates

`func (o *FiltersVolume) GetLinkVolumeLinkDates() []string`

GetLinkVolumeLinkDates returns the LinkVolumeLinkDates field if non-nil, zero value otherwise.

### GetLinkVolumeLinkDatesOk

`func (o *FiltersVolume) GetLinkVolumeLinkDatesOk() ([]string, bool)`

GetLinkVolumeLinkDatesOk returns a tuple with the LinkVolumeLinkDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkVolumeLinkDates

`func (o *FiltersVolume) HasLinkVolumeLinkDates() bool`

HasLinkVolumeLinkDates returns a boolean if a field has been set.

### SetLinkVolumeLinkDates

`func (o *FiltersVolume) SetLinkVolumeLinkDates(v []string)`

SetLinkVolumeLinkDates gets a reference to the given []string and assigns it to the LinkVolumeLinkDates field.

### GetLinkVolumeLinkStates

`func (o *FiltersVolume) GetLinkVolumeLinkStates() []string`

GetLinkVolumeLinkStates returns the LinkVolumeLinkStates field if non-nil, zero value otherwise.

### GetLinkVolumeLinkStatesOk

`func (o *FiltersVolume) GetLinkVolumeLinkStatesOk() ([]string, bool)`

GetLinkVolumeLinkStatesOk returns a tuple with the LinkVolumeLinkStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkVolumeLinkStates

`func (o *FiltersVolume) HasLinkVolumeLinkStates() bool`

HasLinkVolumeLinkStates returns a boolean if a field has been set.

### SetLinkVolumeLinkStates

`func (o *FiltersVolume) SetLinkVolumeLinkStates(v []string)`

SetLinkVolumeLinkStates gets a reference to the given []string and assigns it to the LinkVolumeLinkStates field.

### GetLinkVolumeVmIds

`func (o *FiltersVolume) GetLinkVolumeVmIds() []string`

GetLinkVolumeVmIds returns the LinkVolumeVmIds field if non-nil, zero value otherwise.

### GetLinkVolumeVmIdsOk

`func (o *FiltersVolume) GetLinkVolumeVmIdsOk() ([]string, bool)`

GetLinkVolumeVmIdsOk returns a tuple with the LinkVolumeVmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkVolumeVmIds

`func (o *FiltersVolume) HasLinkVolumeVmIds() bool`

HasLinkVolumeVmIds returns a boolean if a field has been set.

### SetLinkVolumeVmIds

`func (o *FiltersVolume) SetLinkVolumeVmIds(v []string)`

SetLinkVolumeVmIds gets a reference to the given []string and assigns it to the LinkVolumeVmIds field.

### GetSnapshotIds

`func (o *FiltersVolume) GetSnapshotIds() []string`

GetSnapshotIds returns the SnapshotIds field if non-nil, zero value otherwise.

### GetSnapshotIdsOk

`func (o *FiltersVolume) GetSnapshotIdsOk() ([]string, bool)`

GetSnapshotIdsOk returns a tuple with the SnapshotIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSnapshotIds

`func (o *FiltersVolume) HasSnapshotIds() bool`

HasSnapshotIds returns a boolean if a field has been set.

### SetSnapshotIds

`func (o *FiltersVolume) SetSnapshotIds(v []string)`

SetSnapshotIds gets a reference to the given []string and assigns it to the SnapshotIds field.

### GetSubregionNames

`func (o *FiltersVolume) GetSubregionNames() []string`

GetSubregionNames returns the SubregionNames field if non-nil, zero value otherwise.

### GetSubregionNamesOk

`func (o *FiltersVolume) GetSubregionNamesOk() ([]string, bool)`

GetSubregionNamesOk returns a tuple with the SubregionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubregionNames

`func (o *FiltersVolume) HasSubregionNames() bool`

HasSubregionNames returns a boolean if a field has been set.

### SetSubregionNames

`func (o *FiltersVolume) SetSubregionNames(v []string)`

SetSubregionNames gets a reference to the given []string and assigns it to the SubregionNames field.

### GetTagKeys

`func (o *FiltersVolume) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersVolume) GetTagKeysOk() ([]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagKeys

`func (o *FiltersVolume) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### SetTagKeys

`func (o *FiltersVolume) SetTagKeys(v []string)`

SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.

### GetTagValues

`func (o *FiltersVolume) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersVolume) GetTagValuesOk() ([]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagValues

`func (o *FiltersVolume) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### SetTagValues

`func (o *FiltersVolume) SetTagValues(v []string)`

SetTagValues gets a reference to the given []string and assigns it to the TagValues field.

### GetTags

`func (o *FiltersVolume) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersVolume) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *FiltersVolume) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *FiltersVolume) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetVolumeIds

`func (o *FiltersVolume) GetVolumeIds() []string`

GetVolumeIds returns the VolumeIds field if non-nil, zero value otherwise.

### GetVolumeIdsOk

`func (o *FiltersVolume) GetVolumeIdsOk() ([]string, bool)`

GetVolumeIdsOk returns a tuple with the VolumeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeIds

`func (o *FiltersVolume) HasVolumeIds() bool`

HasVolumeIds returns a boolean if a field has been set.

### SetVolumeIds

`func (o *FiltersVolume) SetVolumeIds(v []string)`

SetVolumeIds gets a reference to the given []string and assigns it to the VolumeIds field.

### GetVolumeSizes

`func (o *FiltersVolume) GetVolumeSizes() []int32`

GetVolumeSizes returns the VolumeSizes field if non-nil, zero value otherwise.

### GetVolumeSizesOk

`func (o *FiltersVolume) GetVolumeSizesOk() ([]int32, bool)`

GetVolumeSizesOk returns a tuple with the VolumeSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeSizes

`func (o *FiltersVolume) HasVolumeSizes() bool`

HasVolumeSizes returns a boolean if a field has been set.

### SetVolumeSizes

`func (o *FiltersVolume) SetVolumeSizes(v []int32)`

SetVolumeSizes gets a reference to the given []int32 and assigns it to the VolumeSizes field.

### GetVolumeStates

`func (o *FiltersVolume) GetVolumeStates() []string`

GetVolumeStates returns the VolumeStates field if non-nil, zero value otherwise.

### GetVolumeStatesOk

`func (o *FiltersVolume) GetVolumeStatesOk() ([]string, bool)`

GetVolumeStatesOk returns a tuple with the VolumeStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeStates

`func (o *FiltersVolume) HasVolumeStates() bool`

HasVolumeStates returns a boolean if a field has been set.

### SetVolumeStates

`func (o *FiltersVolume) SetVolumeStates(v []string)`

SetVolumeStates gets a reference to the given []string and assigns it to the VolumeStates field.

### GetVolumeTypes

`func (o *FiltersVolume) GetVolumeTypes() []string`

GetVolumeTypes returns the VolumeTypes field if non-nil, zero value otherwise.

### GetVolumeTypesOk

`func (o *FiltersVolume) GetVolumeTypesOk() ([]string, bool)`

GetVolumeTypesOk returns a tuple with the VolumeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeTypes

`func (o *FiltersVolume) HasVolumeTypes() bool`

HasVolumeTypes returns a boolean if a field has been set.

### SetVolumeTypes

`func (o *FiltersVolume) SetVolumeTypes(v []string)`

SetVolumeTypes gets a reference to the given []string and assigns it to the VolumeTypes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


