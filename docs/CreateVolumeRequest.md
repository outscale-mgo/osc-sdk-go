# CreateVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Iops** | Pointer to **int32** | The number of I/O operations per second (IOPS). This parameter must be specified only if you create an &#x60;io1&#x60; volume. The maximum number of IOPS allowed for &#x60;io1&#x60; volumes is &#x60;13000&#x60;. | [optional] 
**Size** | Pointer to **int32** | The size of the volume, in gibibytes (GiB). The maximum allowed size for a volume is 14,901 GiB. | [optional] 
**SnapshotId** | Pointer to **string** | The ID of the snapshot from which you want to create the volume. | [optional] 
**SubregionName** | Pointer to **string** | The Subregion in which you want to create the volume. | 
**VolumeType** | Pointer to **string** | The type of volume you want to create (&#x60;io1&#x60; \\| &#x60;gp2&#x60; \\| &#x60;standard&#x60;). If not specified, a &#x60;standard&#x60; volume is created.&lt;br /&gt; For more information about volume types, see [Volume Types and IOPS](https://wiki.outscale.net/display/EN/About+Volumes#AboutVolumes-VolumeTypesVolumeTypesandIOPS). | [optional] 

## Methods

### GetDryRun

`func (o *CreateVolumeRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateVolumeRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateVolumeRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateVolumeRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetIops

`func (o *CreateVolumeRequest) GetIops() int32`

GetIops returns the Iops field if non-nil, zero value otherwise.

### GetIopsOk

`func (o *CreateVolumeRequest) GetIopsOk() (int32, bool)`

GetIopsOk returns a tuple with the Iops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIops

`func (o *CreateVolumeRequest) HasIops() bool`

HasIops returns a boolean if a field has been set.

### SetIops

`func (o *CreateVolumeRequest) SetIops(v int32)`

SetIops gets a reference to the given int32 and assigns it to the Iops field.

### GetSize

`func (o *CreateVolumeRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateVolumeRequest) GetSizeOk() (int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSize

`func (o *CreateVolumeRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSize

`func (o *CreateVolumeRequest) SetSize(v int32)`

SetSize gets a reference to the given int32 and assigns it to the Size field.

### GetSnapshotId

`func (o *CreateVolumeRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *CreateVolumeRequest) GetSnapshotIdOk() (string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSnapshotId

`func (o *CreateVolumeRequest) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotId

`func (o *CreateVolumeRequest) SetSnapshotId(v string)`

SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.

### GetSubregionName

`func (o *CreateVolumeRequest) GetSubregionName() string`

GetSubregionName returns the SubregionName field if non-nil, zero value otherwise.

### GetSubregionNameOk

`func (o *CreateVolumeRequest) GetSubregionNameOk() (string, bool)`

GetSubregionNameOk returns a tuple with the SubregionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubregionName

`func (o *CreateVolumeRequest) HasSubregionName() bool`

HasSubregionName returns a boolean if a field has been set.

### SetSubregionName

`func (o *CreateVolumeRequest) SetSubregionName(v string)`

SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.

### GetVolumeType

`func (o *CreateVolumeRequest) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *CreateVolumeRequest) GetVolumeTypeOk() (string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeType

`func (o *CreateVolumeRequest) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeType

`func (o *CreateVolumeRequest) SetVolumeType(v string)`

SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


