# CreateImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** | The architecture of the OMI (by default, &#x60;i386&#x60;). | [optional] 
**BlockDeviceMappings** | Pointer to [**[]BlockDeviceMappingImage**](BlockDeviceMappingImage.md) | One or more block device mappings. | [optional] 
**Description** | Pointer to **string** | A description for the new OMI. | [optional] 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**FileLocation** | Pointer to **string** | The pre-signed URL of the OMI manifest file, or the full path to the OMI stored in an OSU bucket. If you specify this parameter, a copy of the OMI is created in your account. | [optional] 
**ImageName** | Pointer to **string** | A unique name for the new OMI.&lt;br /&gt; Constraints: 3-128 alphanumeric characters, underscores (_), spaces ( ), parentheses (()), slashes (/), periods (.), or dashes (-). | [optional] 
**NoReboot** | Pointer to **bool** | If &#x60;false&#x60;, the VM shuts down before creating the OMI and then reboots. If &#x60;true&#x60;, the VM does not. | [optional] 
**RootDeviceName** | Pointer to **string** | The name of the root device. | [optional] 
**SourceImageId** | Pointer to **string** | The ID of the OMI you want to copy. | [optional] 
**SourceRegionName** | Pointer to **string** | The name of the source Region, which must be the same as the Region of your account. | [optional] 
**VmId** | Pointer to **string** | The ID of the VM from which you want to create the OMI. | [optional] 

## Methods

### GetArchitecture

`func (o *CreateImageRequest) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *CreateImageRequest) GetArchitectureOk() (string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasArchitecture

`func (o *CreateImageRequest) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### SetArchitecture

`func (o *CreateImageRequest) SetArchitecture(v string)`

SetArchitecture gets a reference to the given string and assigns it to the Architecture field.

### GetBlockDeviceMappings

`func (o *CreateImageRequest) GetBlockDeviceMappings() []BlockDeviceMappingImage`

GetBlockDeviceMappings returns the BlockDeviceMappings field if non-nil, zero value otherwise.

### GetBlockDeviceMappingsOk

`func (o *CreateImageRequest) GetBlockDeviceMappingsOk() ([]BlockDeviceMappingImage, bool)`

GetBlockDeviceMappingsOk returns a tuple with the BlockDeviceMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBlockDeviceMappings

`func (o *CreateImageRequest) HasBlockDeviceMappings() bool`

HasBlockDeviceMappings returns a boolean if a field has been set.

### SetBlockDeviceMappings

`func (o *CreateImageRequest) SetBlockDeviceMappings(v []BlockDeviceMappingImage)`

SetBlockDeviceMappings gets a reference to the given []BlockDeviceMappingImage and assigns it to the BlockDeviceMappings field.

### GetDescription

`func (o *CreateImageRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateImageRequest) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *CreateImageRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *CreateImageRequest) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetDryRun

`func (o *CreateImageRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateImageRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateImageRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateImageRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetFileLocation

`func (o *CreateImageRequest) GetFileLocation() string`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *CreateImageRequest) GetFileLocationOk() (string, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileLocation

`func (o *CreateImageRequest) HasFileLocation() bool`

HasFileLocation returns a boolean if a field has been set.

### SetFileLocation

`func (o *CreateImageRequest) SetFileLocation(v string)`

SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.

### GetImageName

`func (o *CreateImageRequest) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *CreateImageRequest) GetImageNameOk() (string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImageName

`func (o *CreateImageRequest) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### SetImageName

`func (o *CreateImageRequest) SetImageName(v string)`

SetImageName gets a reference to the given string and assigns it to the ImageName field.

### GetNoReboot

`func (o *CreateImageRequest) GetNoReboot() bool`

GetNoReboot returns the NoReboot field if non-nil, zero value otherwise.

### GetNoRebootOk

`func (o *CreateImageRequest) GetNoRebootOk() (bool, bool)`

GetNoRebootOk returns a tuple with the NoReboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNoReboot

`func (o *CreateImageRequest) HasNoReboot() bool`

HasNoReboot returns a boolean if a field has been set.

### SetNoReboot

`func (o *CreateImageRequest) SetNoReboot(v bool)`

SetNoReboot gets a reference to the given bool and assigns it to the NoReboot field.

### GetRootDeviceName

`func (o *CreateImageRequest) GetRootDeviceName() string`

GetRootDeviceName returns the RootDeviceName field if non-nil, zero value otherwise.

### GetRootDeviceNameOk

`func (o *CreateImageRequest) GetRootDeviceNameOk() (string, bool)`

GetRootDeviceNameOk returns a tuple with the RootDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRootDeviceName

`func (o *CreateImageRequest) HasRootDeviceName() bool`

HasRootDeviceName returns a boolean if a field has been set.

### SetRootDeviceName

`func (o *CreateImageRequest) SetRootDeviceName(v string)`

SetRootDeviceName gets a reference to the given string and assigns it to the RootDeviceName field.

### GetSourceImageId

`func (o *CreateImageRequest) GetSourceImageId() string`

GetSourceImageId returns the SourceImageId field if non-nil, zero value otherwise.

### GetSourceImageIdOk

`func (o *CreateImageRequest) GetSourceImageIdOk() (string, bool)`

GetSourceImageIdOk returns a tuple with the SourceImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceImageId

`func (o *CreateImageRequest) HasSourceImageId() bool`

HasSourceImageId returns a boolean if a field has been set.

### SetSourceImageId

`func (o *CreateImageRequest) SetSourceImageId(v string)`

SetSourceImageId gets a reference to the given string and assigns it to the SourceImageId field.

### GetSourceRegionName

`func (o *CreateImageRequest) GetSourceRegionName() string`

GetSourceRegionName returns the SourceRegionName field if non-nil, zero value otherwise.

### GetSourceRegionNameOk

`func (o *CreateImageRequest) GetSourceRegionNameOk() (string, bool)`

GetSourceRegionNameOk returns a tuple with the SourceRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceRegionName

`func (o *CreateImageRequest) HasSourceRegionName() bool`

HasSourceRegionName returns a boolean if a field has been set.

### SetSourceRegionName

`func (o *CreateImageRequest) SetSourceRegionName(v string)`

SetSourceRegionName gets a reference to the given string and assigns it to the SourceRegionName field.

### GetVmId

`func (o *CreateImageRequest) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *CreateImageRequest) GetVmIdOk() (string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmId

`func (o *CreateImageRequest) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### SetVmId

`func (o *CreateImageRequest) SetVmId(v string)`

SetVmId gets a reference to the given string and assigns it to the VmId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


