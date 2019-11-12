/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oscgo

import (
	"bytes"
	"encoding/json"
)

// Image Information about the OMI.
type Image struct {
	// The account alias of the owner of the OMI.
	AccountAlias *string `json:"AccountAlias,omitempty"`
	// The account ID of the owner of the OMI.
	AccountId *string `json:"AccountId,omitempty"`
	// The architecture of the OMI (by default, `i386`).
	Architecture *string `json:"Architecture,omitempty"`
	// One or more block device mappings.
	BlockDeviceMappings *[]BlockDeviceMappingImage `json:"BlockDeviceMappings,omitempty"`
	// The date and time at which the OMI was created.
	CreationDate *string `json:"CreationDate,omitempty"`
	// The description of the OMI.
	Description *string `json:"Description,omitempty"`
	// The location where the OMI file is stored on Object Storage Unit (OSU).
	FileLocation *string `json:"FileLocation,omitempty"`
	// The ID of the OMI.
	ImageId *string `json:"ImageId,omitempty"`
	// The name of the OMI.
	ImageName *string `json:"ImageName,omitempty"`
	// The type of the OMI.
	ImageType *string `json:"ImageType,omitempty"`
	PermissionsToLaunch *PermissionsOnResource `json:"PermissionsToLaunch,omitempty"`
	// The product code associated with the OMI (`0001` Linux/Unix \\| `0002` Windows \\| `0004` Linux/Oracle \\| `0005` Windows 10).
	ProductCodes *[]string `json:"ProductCodes,omitempty"`
	// The name of the root device.
	RootDeviceName *string `json:"RootDeviceName,omitempty"`
	// The type of root device used by the OMI (always `bsu`).
	RootDeviceType *string `json:"RootDeviceType,omitempty"`
	// The state of the OMI.
	State *string `json:"State,omitempty"`
	StateComment *StateComment `json:"StateComment,omitempty"`
	// One or more tags associated with the OMI.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
}

// GetAccountAlias returns the AccountAlias field value if set, zero value otherwise.
func (o *Image) GetAccountAlias() string {
	if o == nil || o.AccountAlias == nil {
		var ret string
		return ret
	}
	return *o.AccountAlias
}

// GetAccountAliasOk returns a tuple with the AccountAlias field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetAccountAliasOk() (string, bool) {
	if o == nil || o.AccountAlias == nil {
		var ret string
		return ret, false
	}
	return *o.AccountAlias, true
}

// HasAccountAlias returns a boolean if a field has been set.
func (o *Image) HasAccountAlias() bool {
	if o != nil && o.AccountAlias != nil {
		return true
	}

	return false
}

// SetAccountAlias gets a reference to the given string and assigns it to the AccountAlias field.
func (o *Image) SetAccountAlias(v string) {
	o.AccountAlias = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Image) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetAccountIdOk() (string, bool) {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret, false
	}
	return *o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Image) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Image) SetAccountId(v string) {
	o.AccountId = &v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *Image) GetArchitecture() string {
	if o == nil || o.Architecture == nil {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetArchitectureOk() (string, bool) {
	if o == nil || o.Architecture == nil {
		var ret string
		return ret, false
	}
	return *o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *Image) HasArchitecture() bool {
	if o != nil && o.Architecture != nil {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *Image) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetBlockDeviceMappings returns the BlockDeviceMappings field value if set, zero value otherwise.
func (o *Image) GetBlockDeviceMappings() []BlockDeviceMappingImage {
	if o == nil || o.BlockDeviceMappings == nil {
		var ret []BlockDeviceMappingImage
		return ret
	}
	return *o.BlockDeviceMappings
}

// GetBlockDeviceMappingsOk returns a tuple with the BlockDeviceMappings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetBlockDeviceMappingsOk() ([]BlockDeviceMappingImage, bool) {
	if o == nil || o.BlockDeviceMappings == nil {
		var ret []BlockDeviceMappingImage
		return ret, false
	}
	return *o.BlockDeviceMappings, true
}

// HasBlockDeviceMappings returns a boolean if a field has been set.
func (o *Image) HasBlockDeviceMappings() bool {
	if o != nil && o.BlockDeviceMappings != nil {
		return true
	}

	return false
}

// SetBlockDeviceMappings gets a reference to the given []BlockDeviceMappingImage and assigns it to the BlockDeviceMappings field.
func (o *Image) SetBlockDeviceMappings(v []BlockDeviceMappingImage) {
	o.BlockDeviceMappings = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *Image) GetCreationDate() string {
	if o == nil || o.CreationDate == nil {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetCreationDateOk() (string, bool) {
	if o == nil || o.CreationDate == nil {
		var ret string
		return ret, false
	}
	return *o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *Image) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *Image) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Image) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Image) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Image) SetDescription(v string) {
	o.Description = &v
}

// GetFileLocation returns the FileLocation field value if set, zero value otherwise.
func (o *Image) GetFileLocation() string {
	if o == nil || o.FileLocation == nil {
		var ret string
		return ret
	}
	return *o.FileLocation
}

// GetFileLocationOk returns a tuple with the FileLocation field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetFileLocationOk() (string, bool) {
	if o == nil || o.FileLocation == nil {
		var ret string
		return ret, false
	}
	return *o.FileLocation, true
}

// HasFileLocation returns a boolean if a field has been set.
func (o *Image) HasFileLocation() bool {
	if o != nil && o.FileLocation != nil {
		return true
	}

	return false
}

// SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.
func (o *Image) SetFileLocation(v string) {
	o.FileLocation = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *Image) GetImageId() string {
	if o == nil || o.ImageId == nil {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetImageIdOk() (string, bool) {
	if o == nil || o.ImageId == nil {
		var ret string
		return ret, false
	}
	return *o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *Image) HasImageId() bool {
	if o != nil && o.ImageId != nil {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *Image) SetImageId(v string) {
	o.ImageId = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *Image) GetImageName() string {
	if o == nil || o.ImageName == nil {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetImageNameOk() (string, bool) {
	if o == nil || o.ImageName == nil {
		var ret string
		return ret, false
	}
	return *o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *Image) HasImageName() bool {
	if o != nil && o.ImageName != nil {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *Image) SetImageName(v string) {
	o.ImageName = &v
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *Image) GetImageType() string {
	if o == nil || o.ImageType == nil {
		var ret string
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetImageTypeOk() (string, bool) {
	if o == nil || o.ImageType == nil {
		var ret string
		return ret, false
	}
	return *o.ImageType, true
}

// HasImageType returns a boolean if a field has been set.
func (o *Image) HasImageType() bool {
	if o != nil && o.ImageType != nil {
		return true
	}

	return false
}

// SetImageType gets a reference to the given string and assigns it to the ImageType field.
func (o *Image) SetImageType(v string) {
	o.ImageType = &v
}

// GetPermissionsToLaunch returns the PermissionsToLaunch field value if set, zero value otherwise.
func (o *Image) GetPermissionsToLaunch() PermissionsOnResource {
	if o == nil || o.PermissionsToLaunch == nil {
		var ret PermissionsOnResource
		return ret
	}
	return *o.PermissionsToLaunch
}

// GetPermissionsToLaunchOk returns a tuple with the PermissionsToLaunch field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetPermissionsToLaunchOk() (PermissionsOnResource, bool) {
	if o == nil || o.PermissionsToLaunch == nil {
		var ret PermissionsOnResource
		return ret, false
	}
	return *o.PermissionsToLaunch, true
}

// HasPermissionsToLaunch returns a boolean if a field has been set.
func (o *Image) HasPermissionsToLaunch() bool {
	if o != nil && o.PermissionsToLaunch != nil {
		return true
	}

	return false
}

// SetPermissionsToLaunch gets a reference to the given PermissionsOnResource and assigns it to the PermissionsToLaunch field.
func (o *Image) SetPermissionsToLaunch(v PermissionsOnResource) {
	o.PermissionsToLaunch = &v
}

// GetProductCodes returns the ProductCodes field value if set, zero value otherwise.
func (o *Image) GetProductCodes() []string {
	if o == nil || o.ProductCodes == nil {
		var ret []string
		return ret
	}
	return *o.ProductCodes
}

// GetProductCodesOk returns a tuple with the ProductCodes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetProductCodesOk() ([]string, bool) {
	if o == nil || o.ProductCodes == nil {
		var ret []string
		return ret, false
	}
	return *o.ProductCodes, true
}

// HasProductCodes returns a boolean if a field has been set.
func (o *Image) HasProductCodes() bool {
	if o != nil && o.ProductCodes != nil {
		return true
	}

	return false
}

// SetProductCodes gets a reference to the given []string and assigns it to the ProductCodes field.
func (o *Image) SetProductCodes(v []string) {
	o.ProductCodes = &v
}

// GetRootDeviceName returns the RootDeviceName field value if set, zero value otherwise.
func (o *Image) GetRootDeviceName() string {
	if o == nil || o.RootDeviceName == nil {
		var ret string
		return ret
	}
	return *o.RootDeviceName
}

// GetRootDeviceNameOk returns a tuple with the RootDeviceName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetRootDeviceNameOk() (string, bool) {
	if o == nil || o.RootDeviceName == nil {
		var ret string
		return ret, false
	}
	return *o.RootDeviceName, true
}

// HasRootDeviceName returns a boolean if a field has been set.
func (o *Image) HasRootDeviceName() bool {
	if o != nil && o.RootDeviceName != nil {
		return true
	}

	return false
}

// SetRootDeviceName gets a reference to the given string and assigns it to the RootDeviceName field.
func (o *Image) SetRootDeviceName(v string) {
	o.RootDeviceName = &v
}

// GetRootDeviceType returns the RootDeviceType field value if set, zero value otherwise.
func (o *Image) GetRootDeviceType() string {
	if o == nil || o.RootDeviceType == nil {
		var ret string
		return ret
	}
	return *o.RootDeviceType
}

// GetRootDeviceTypeOk returns a tuple with the RootDeviceType field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetRootDeviceTypeOk() (string, bool) {
	if o == nil || o.RootDeviceType == nil {
		var ret string
		return ret, false
	}
	return *o.RootDeviceType, true
}

// HasRootDeviceType returns a boolean if a field has been set.
func (o *Image) HasRootDeviceType() bool {
	if o != nil && o.RootDeviceType != nil {
		return true
	}

	return false
}

// SetRootDeviceType gets a reference to the given string and assigns it to the RootDeviceType field.
func (o *Image) SetRootDeviceType(v string) {
	o.RootDeviceType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Image) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetStateOk() (string, bool) {
	if o == nil || o.State == nil {
		var ret string
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Image) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Image) SetState(v string) {
	o.State = &v
}

// GetStateComment returns the StateComment field value if set, zero value otherwise.
func (o *Image) GetStateComment() StateComment {
	if o == nil || o.StateComment == nil {
		var ret StateComment
		return ret
	}
	return *o.StateComment
}

// GetStateCommentOk returns a tuple with the StateComment field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetStateCommentOk() (StateComment, bool) {
	if o == nil || o.StateComment == nil {
		var ret StateComment
		return ret, false
	}
	return *o.StateComment, true
}

// HasStateComment returns a boolean if a field has been set.
func (o *Image) HasStateComment() bool {
	if o != nil && o.StateComment != nil {
		return true
	}

	return false
}

// SetStateComment gets a reference to the given StateComment and assigns it to the StateComment field.
func (o *Image) SetStateComment(v StateComment) {
	o.StateComment = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Image) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetTagsOk() ([]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Image) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *Image) SetTags(v []ResourceTag) {
	o.Tags = &v
}

type NullableImage struct {
	Value Image
	ExplicitNull bool
}

func (v NullableImage) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableImage) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

