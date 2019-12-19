# Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iops** | Pointer to **int64** | The number of I/O operations per second (IOPS):&lt;br /&gt; - For &#x60;io1&#x60; volumes, the number of provisioned IOPS&lt;br /&gt; - For &#x60;gp2&#x60; volumes, the baseline performance of the volume | [optional] 
**LinkedVolumes** | Pointer to [**[]LinkedVolume**](LinkedVolume.md) | Information about your volume attachment. | [optional] 
**Size** | Pointer to **int64** | The size of the volume, in gibibytes (GiB). | [optional] 
**SnapshotId** | Pointer to **string** | The snapshot from which the volume was created. | [optional] 
**State** | Pointer to **string** | The state of the volume (&#x60;creating&#x60; \\| &#x60;available&#x60; \\| &#x60;in-use&#x60; \\| &#x60;deleting&#x60; \\| &#x60;error&#x60;). | [optional] 
**SubregionName** | Pointer to **string** | The Subregion in which the volume was created. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the volume. | [optional] 
**VolumeId** | Pointer to **string** | The ID of the volume. | [optional] 
**VolumeType** | Pointer to **string** | The type of the volume (&#x60;standard&#x60; \\| &#x60;gp2&#x60; \\| &#x60;io1&#x60;). | [optional] 

## Methods

### GetIops

`func (o *Volume) GetIops() int64`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *Volume) GetIopsOk() (int64, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIops

`func (o *Volume) HasIops() bool`

HasIops returns a boolean if a field has been set.

### SetIops

`func (o *Volume) SetIops(v int64)`

SetIops gets a reference to the given int64 and assigns it to the Iops field.

### GetLinkedVolumes

`func (o *Volume) GetLinkedVolumes() []LinkedVolume`

GetLinkedVolumes returns the LinkedVolumes field if non-nil, zero value otherwise.

### GetLinkedVolumesOk

`func (o *Volume) GetLinkedVolumesOk() ([]LinkedVolume, bool)`

GetLinkedVolumesOk returns a tuple with the LinkedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkedVolumes

`func (o *Volume) HasLinkedVolumes() bool`

HasLinkedVolumes returns a boolean if a field has been set.

### SetLinkedVolumes

`func (o *Volume) SetLinkedVolumes(v []LinkedVolume)`

SetLinkedVolumes gets a reference to the given []LinkedVolume and assigns it to the LinkedVolumes field.

### GetSize

`func (o *Volume) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Volume) GetSizeOk() (int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSize

`func (o *Volume) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSize

`func (o *Volume) SetSize(v int64)`

SetSize gets a reference to the given int64 and assigns it to the Size field.

### GetSnapshotId

`func (o *Volume) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *Volume) GetSnapshotIdOk() (string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSnapshotId

`func (o *Volume) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotId

`func (o *Volume) SetSnapshotId(v string)`

SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.

### GetState

`func (o *Volume) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Volume) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *Volume) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *Volume) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetSubregionName

`func (o *Volume) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *Volume) GetSubregionNameOk() (string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubregionName

`func (o *Volume) HasSubregionName() bool`

HasSubregionName returns a boolean if a field has been set.

### SetSubregionName

`func (o *Volume) SetSubregionName(v string)`

SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.

### GetTags

`func (o *Volume) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Volume) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *Volume) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *Volume) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.

### GetVolumeId

`func (o *Volume) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *Volume) GetVolumeIdOk() (string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeId

`func (o *Volume) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### SetVolumeId

`func (o *Volume) SetVolumeId(v string)`

SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.

### GetVolumeType

`func (o *Volume) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *Volume) GetVolumeTypeOk() (string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeType

`func (o *Volume) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeType

`func (o *Volume) SetVolumeType(v string)`

SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


