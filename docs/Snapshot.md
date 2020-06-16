# Snapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAlias** | Pointer to **string** | The account alias of the owner of the snapshot. | [optional] 
**AccountId** | Pointer to **string** | The account ID of the owner of the snapshot. | [optional] 
**Description** | Pointer to **string** | The description of the snapshot. | [optional] 
**PermissionsToCreateVolume** | Pointer to [**PermissionsOnResource**](PermissionsOnResource.md) |  | [optional] 
**Progress** | Pointer to **int64** | The progress of the snapshot, as a percentage. | [optional] 
**SnapshotId** | Pointer to **string** | The ID of the snapshot. | [optional] 
**State** | Pointer to **string** | The state of the snapshot (&#x60;in-queue&#x60; \\| &#x60;completed&#x60; \\| &#x60;error&#x60;). | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the snapshot. | [optional] 
**VolumeId** | Pointer to **string** | The ID of the volume used to create the snapshot. | [optional] 
**VolumeSize** | Pointer to **int64** | The size of the volume used to create the snapshot, in gibibytes (GiB). | [optional] 

## Methods

### GetAccountAlias

`func (o *Snapshot) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *Snapshot) GetAccountAliasOk() (string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountAlias

`func (o *Snapshot) HasAccountAlias() bool`

HasAccountAlias returns a boolean if a field has been set.

### SetAccountAlias

`func (o *Snapshot) SetAccountAlias(v string)`

SetAccountAlias gets a reference to the given string and assigns it to the AccountAlias field.

### GetAccountId

`func (o *Snapshot) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Snapshot) GetAccountIdOk() (string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *Snapshot) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *Snapshot) SetAccountId(v string)`

SetAccountId gets a reference to the given string and assigns it to the AccountId field.

### GetDescription

`func (o *Snapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Snapshot) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Snapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Snapshot) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetPermissionsToCreateVolume

`func (o *Snapshot) GetPermissionsToCreateVolume() PermissionsOnResource`

GetPermissionsToCreateVolume returns the PermissionsToCreateVolume field if non-nil, zero value otherwise.

### GetPermissionsToCreateVolumeOk

`func (o *Snapshot) GetPermissionsToCreateVolumeOk() (PermissionsOnResource, bool)`

GetPermissionsToCreateVolumeOk returns a tuple with the PermissionsToCreateVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPermissionsToCreateVolume

`func (o *Snapshot) HasPermissionsToCreateVolume() bool`

HasPermissionsToCreateVolume returns a boolean if a field has been set.

### SetPermissionsToCreateVolume

`func (o *Snapshot) SetPermissionsToCreateVolume(v PermissionsOnResource)`

SetPermissionsToCreateVolume gets a reference to the given PermissionsOnResource and assigns it to the PermissionsToCreateVolume field.

### GetProgress

`func (o *Snapshot) GetProgress() int64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Snapshot) GetProgressOk() (int64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProgress

`func (o *Snapshot) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgress

`func (o *Snapshot) SetProgress(v int64)`

SetProgress gets a reference to the given int64 and assigns it to the Progress field.

### GetSnapshotId

`func (o *Snapshot) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *Snapshot) GetSnapshotIdOk() (string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSnapshotId

`func (o *Snapshot) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotId

`func (o *Snapshot) SetSnapshotId(v string)`

SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.

### GetState

`func (o *Snapshot) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Snapshot) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *Snapshot) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *Snapshot) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetTags

`func (o *Snapshot) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Snapshot) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *Snapshot) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *Snapshot) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.

### GetVolumeId

`func (o *Snapshot) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *Snapshot) GetVolumeIdOk() (string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeId

`func (o *Snapshot) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### SetVolumeId

`func (o *Snapshot) SetVolumeId(v string)`

SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.

### GetVolumeSize

`func (o *Snapshot) GetVolumeSize() int64`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *Snapshot) GetVolumeSizeOk() (int64, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeSize

`func (o *Snapshot) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.

### SetVolumeSize

`func (o *Snapshot) SetVolumeSize(v int64)`

SetVolumeSize gets a reference to the given int64 and assigns it to the VolumeSize field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


