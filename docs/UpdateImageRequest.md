# UpdateImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**ImageId** | Pointer to **string** | The ID of the OMI you want to modify. | 
**PermissionsToLaunch** | Pointer to [**PermissionsOnResourceCreation**](PermissionsOnResourceCreation.md) |  | 

## Methods

### GetDryRun

`func (o *UpdateImageRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateImageRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UpdateImageRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UpdateImageRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetImageId

`func (o *UpdateImageRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *UpdateImageRequest) GetImageIdOk() (string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImageId

`func (o *UpdateImageRequest) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageId

`func (o *UpdateImageRequest) SetImageId(v string)`

SetImageId gets a reference to the given string and assigns it to the ImageId field.

### GetPermissionsToLaunch

`func (o *UpdateImageRequest) GetPermissionsToLaunch() PermissionsOnResourceCreation`

GetPermissionsToLaunch returns the PermissionsToLaunch field if non-nil, zero value otherwise.

### GetPermissionsToLaunchOk

`func (o *UpdateImageRequest) GetPermissionsToLaunchOk() (PermissionsOnResourceCreation, bool)`

GetPermissionsToLaunchOk returns a tuple with the PermissionsToLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPermissionsToLaunch

`func (o *UpdateImageRequest) HasPermissionsToLaunch() bool`

HasPermissionsToLaunch returns a boolean if a field has been set.

### SetPermissionsToLaunch

`func (o *UpdateImageRequest) SetPermissionsToLaunch(v PermissionsOnResourceCreation)`

SetPermissionsToLaunch gets a reference to the given PermissionsOnResourceCreation and assigns it to the PermissionsToLaunch field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


