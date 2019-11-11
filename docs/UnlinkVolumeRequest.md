# UnlinkVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**ForceUnlink** | Pointer to **bool** | Forces the detachment of the volume in case of previous failure. Important: This action may damage your data or file systems. | [optional] 
**VolumeId** | Pointer to **string** | The ID of the volume you want to detach. | 

## Methods

### GetDryRun

`func (o *UnlinkVolumeRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UnlinkVolumeRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UnlinkVolumeRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UnlinkVolumeRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetForceUnlink

`func (o *UnlinkVolumeRequest) GetForceUnlink() bool`

GetForceUnlink returns the ForceUnlink field if non-nil, zero value otherwise.

### GetForceUnlinkOk

`func (o *UnlinkVolumeRequest) GetForceUnlinkOk() (bool, bool)`

GetForceUnlinkOk returns a tuple with the ForceUnlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasForceUnlink

`func (o *UnlinkVolumeRequest) HasForceUnlink() bool`

HasForceUnlink returns a boolean if a field has been set.

### SetForceUnlink

`func (o *UnlinkVolumeRequest) SetForceUnlink(v bool)`

SetForceUnlink gets a reference to the given bool and assigns it to the ForceUnlink field.

### GetVolumeId

`func (o *UnlinkVolumeRequest) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *UnlinkVolumeRequest) GetVolumeIdOk() (string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeId

`func (o *UnlinkVolumeRequest) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### SetVolumeId

`func (o *UnlinkVolumeRequest) SetVolumeId(v string)`

SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


