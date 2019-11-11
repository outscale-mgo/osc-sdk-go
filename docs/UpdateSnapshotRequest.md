# UpdateSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**PermissionsToCreateVolume** | Pointer to [**PermissionsOnResourceCreation**](PermissionsOnResourceCreation.md) |  | 
**SnapshotId** | Pointer to **string** | The ID of the snapshot. | 

## Methods

### GetDryRun

`func (o *UpdateSnapshotRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateSnapshotRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UpdateSnapshotRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UpdateSnapshotRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetPermissionsToCreateVolume

`func (o *UpdateSnapshotRequest) GetPermissionsToCreateVolume() PermissionsOnResourceCreation`

GetPermissionsToCreateVolume returns the PermissionsToCreateVolume field if non-nil, zero value otherwise.

### GetPermissionsToCreateVolumeOk

`func (o *UpdateSnapshotRequest) GetPermissionsToCreateVolumeOk() (PermissionsOnResourceCreation, bool)`

GetPermissionsToCreateVolumeOk returns a tuple with the PermissionsToCreateVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPermissionsToCreateVolume

`func (o *UpdateSnapshotRequest) HasPermissionsToCreateVolume() bool`

HasPermissionsToCreateVolume returns a boolean if a field has been set.

### SetPermissionsToCreateVolume

`func (o *UpdateSnapshotRequest) SetPermissionsToCreateVolume(v PermissionsOnResourceCreation)`

SetPermissionsToCreateVolume gets a reference to the given PermissionsOnResourceCreation and assigns it to the PermissionsToCreateVolume field.

### GetSnapshotId

`func (o *UpdateSnapshotRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *UpdateSnapshotRequest) GetSnapshotIdOk() (string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSnapshotId

`func (o *UpdateSnapshotRequest) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotId

`func (o *UpdateSnapshotRequest) SetSnapshotId(v string)`

SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


