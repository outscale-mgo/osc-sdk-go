# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAlias** | Pointer to **string** | The account alias of the owner of the OMI. | [optional] 
**AccountId** | Pointer to **string** | The account ID of the owner of the OMI. | [optional] 
**Architecture** | Pointer to **string** | The architecture of the OMI (by default, &#x60;i386&#x60;). | [optional] 
**BlockDeviceMappings** | Pointer to [**[]BlockDeviceMappingImage**](BlockDeviceMappingImage.md) | One or more block device mappings. | [optional] 
**CreationDate** | Pointer to **string** | The date and time at which the OMI was created. | [optional] 
**Description** | Pointer to **string** | The description of the OMI. | [optional] 
**FileLocation** | Pointer to **string** | The location where the OMI file is stored on Object Storage Unit (OSU). | [optional] 
**ImageId** | Pointer to **string** | The ID of the OMI. | [optional] 
**ImageName** | Pointer to **string** | The name of the OMI. | [optional] 
**ImageType** | Pointer to **string** | The type of the OMI. | [optional] 
**PermissionsToLaunch** | Pointer to [**PermissionsOnResource**](PermissionsOnResource.md) |  | [optional] 
**ProductCodes** | Pointer to **[]string** | The product code associated with the OMI (&#x60;0001&#x60; Linux/Unix \\| &#x60;0002&#x60; Windows \\| &#x60;0004&#x60; Linux/Oracle \\| &#x60;0005&#x60; Windows 10). | [optional] 
**RootDeviceName** | Pointer to **string** | The name of the root device. | [optional] 
**RootDeviceType** | Pointer to **string** | The type of root device used by the OMI (always &#x60;bsu&#x60;). | [optional] 
**State** | Pointer to **string** | The state of the OMI (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;failed&#x60;). | [optional] 
**StateComment** | Pointer to [**StateComment**](StateComment.md) |  | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the OMI. | [optional] 

## Methods

### GetAccountAlias

`func (o *Image) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *Image) GetAccountAliasOk() (string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountAlias

`func (o *Image) HasAccountAlias() bool`

HasAccountAlias returns a boolean if a field has been set.

### SetAccountAlias

`func (o *Image) SetAccountAlias(v string)`

SetAccountAlias gets a reference to the given string and assigns it to the AccountAlias field.

### GetAccountId

`func (o *Image) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Image) GetAccountIdOk() (string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *Image) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *Image) SetAccountId(v string)`

SetAccountId gets a reference to the given string and assigns it to the AccountId field.

### GetArchitecture

`func (o *Image) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *Image) GetArchitectureOk() (string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasArchitecture

`func (o *Image) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### SetArchitecture

`func (o *Image) SetArchitecture(v string)`

SetArchitecture gets a reference to the given string and assigns it to the Architecture field.

### GetBlockDeviceMappings

`func (o *Image) GetBlockDeviceMappings() []BlockDeviceMappingImage`

GetBlockDeviceMappings returns the BlockDeviceMappings field if non-nil, zero value otherwise.

### GetBlockDeviceMappingsOk

`func (o *Image) GetBlockDeviceMappingsOk() ([]BlockDeviceMappingImage, bool)`

GetBlockDeviceMappingsOk returns a tuple with the BlockDeviceMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBlockDeviceMappings

`func (o *Image) HasBlockDeviceMappings() bool`

HasBlockDeviceMappings returns a boolean if a field has been set.

### SetBlockDeviceMappings

`func (o *Image) SetBlockDeviceMappings(v []BlockDeviceMappingImage)`

SetBlockDeviceMappings gets a reference to the given []BlockDeviceMappingImage and assigns it to the BlockDeviceMappings field.

### GetCreationDate

`func (o *Image) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Image) GetCreationDateOk() (string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreationDate

`func (o *Image) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### SetCreationDate

`func (o *Image) SetCreationDate(v string)`

SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.

### GetDescription

`func (o *Image) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Image) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Image) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Image) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetFileLocation

`func (o *Image) GetFileLocation() string`

GetFileLocation returns the FileLocation field if non-nil, zero value otherwise.

### GetFileLocationOk

`func (o *Image) GetFileLocationOk() (string, bool)`

GetFileLocationOk returns a tuple with the FileLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFileLocation

`func (o *Image) HasFileLocation() bool`

HasFileLocation returns a boolean if a field has been set.

### SetFileLocation

`func (o *Image) SetFileLocation(v string)`

SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.

### GetImageId

`func (o *Image) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *Image) GetImageIdOk() (string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImageId

`func (o *Image) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageId

`func (o *Image) SetImageId(v string)`

SetImageId gets a reference to the given string and assigns it to the ImageId field.

### GetImageName

`func (o *Image) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *Image) GetImageNameOk() (string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImageName

`func (o *Image) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### SetImageName

`func (o *Image) SetImageName(v string)`

SetImageName gets a reference to the given string and assigns it to the ImageName field.

### GetImageType

`func (o *Image) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *Image) GetImageTypeOk() (string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImageType

`func (o *Image) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### SetImageType

`func (o *Image) SetImageType(v string)`

SetImageType gets a reference to the given string and assigns it to the ImageType field.

### GetPermissionsToLaunch

`func (o *Image) GetPermissionsToLaunch() PermissionsOnResource`

GetPermissionsToLaunch returns the PermissionsToLaunch field if non-nil, zero value otherwise.

### GetPermissionsToLaunchOk

`func (o *Image) GetPermissionsToLaunchOk() (PermissionsOnResource, bool)`

GetPermissionsToLaunchOk returns a tuple with the PermissionsToLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPermissionsToLaunch

`func (o *Image) HasPermissionsToLaunch() bool`

HasPermissionsToLaunch returns a boolean if a field has been set.

### SetPermissionsToLaunch

`func (o *Image) SetPermissionsToLaunch(v PermissionsOnResource)`

SetPermissionsToLaunch gets a reference to the given PermissionsOnResource and assigns it to the PermissionsToLaunch field.

### GetProductCodes

`func (o *Image) GetProductCodes() []string`

GetProductCodes returns the ProductCodes field if non-nil, zero value otherwise.

### GetProductCodesOk

`func (o *Image) GetProductCodesOk() ([]string, bool)`

GetProductCodesOk returns a tuple with the ProductCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductCodes

`func (o *Image) HasProductCodes() bool`

HasProductCodes returns a boolean if a field has been set.

### SetProductCodes

`func (o *Image) SetProductCodes(v []string)`

SetProductCodes gets a reference to the given []string and assigns it to the ProductCodes field.

### GetRootDeviceName

`func (o *Image) GetRootDeviceName() string`

GetRootDeviceName returns the RootDeviceName field if non-nil, zero value otherwise.

### GetRootDeviceNameOk

`func (o *Image) GetRootDeviceNameOk() (string, bool)`

GetRootDeviceNameOk returns a tuple with the RootDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRootDeviceName

`func (o *Image) HasRootDeviceName() bool`

HasRootDeviceName returns a boolean if a field has been set.

### SetRootDeviceName

`func (o *Image) SetRootDeviceName(v string)`

SetRootDeviceName gets a reference to the given string and assigns it to the RootDeviceName field.

### GetRootDeviceType

`func (o *Image) GetRootDeviceType() string`

GetRootDeviceType returns the RootDeviceType field if non-nil, zero value otherwise.

### GetRootDeviceTypeOk

`func (o *Image) GetRootDeviceTypeOk() (string, bool)`

GetRootDeviceTypeOk returns a tuple with the RootDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRootDeviceType

`func (o *Image) HasRootDeviceType() bool`

HasRootDeviceType returns a boolean if a field has been set.

### SetRootDeviceType

`func (o *Image) SetRootDeviceType(v string)`

SetRootDeviceType gets a reference to the given string and assigns it to the RootDeviceType field.

### GetState

`func (o *Image) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Image) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *Image) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *Image) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetStateComment

`func (o *Image) GetStateComment() StateComment`

GetStateComment returns the StateComment field if non-nil, zero value otherwise.

### GetStateCommentOk

`func (o *Image) GetStateCommentOk() (StateComment, bool)`

GetStateCommentOk returns a tuple with the StateComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStateComment

`func (o *Image) HasStateComment() bool`

HasStateComment returns a boolean if a field has been set.

### SetStateComment

`func (o *Image) SetStateComment(v StateComment)`

SetStateComment gets a reference to the given StateComment and assigns it to the StateComment field.

### GetTags

`func (o *Image) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Image) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *Image) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *Image) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


