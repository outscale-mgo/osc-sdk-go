# LinkVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | Pointer to **string** | The name of the device. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**VmId** | Pointer to **string** | The ID of the VM you want to attach the volume to. | 
**VolumeId** | Pointer to **string** | The ID of the volume you want to attach. | 

## Methods

### GetDeviceName

`func (o *LinkVolumeRequest) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *LinkVolumeRequest) GetDeviceNameOk() (string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceName

`func (o *LinkVolumeRequest) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceName

`func (o *LinkVolumeRequest) SetDeviceName(v string)`

SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.

### GetDryRun

`func (o *LinkVolumeRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *LinkVolumeRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *LinkVolumeRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *LinkVolumeRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetVmId

`func (o *LinkVolumeRequest) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *LinkVolumeRequest) GetVmIdOk() (string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmId

`func (o *LinkVolumeRequest) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### SetVmId

`func (o *LinkVolumeRequest) SetVmId(v string)`

SetVmId gets a reference to the given string and assigns it to the VmId field.

### GetVolumeId

`func (o *LinkVolumeRequest) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *LinkVolumeRequest) GetVolumeIdOk() (string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVolumeId

`func (o *LinkVolumeRequest) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### SetVolumeId

`func (o *LinkVolumeRequest) SetVolumeId(v string)`

SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


