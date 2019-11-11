/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"bytes"
	"encoding/json"
)

// LinkNicToUpdate Information about the NIC attachment. If you are modifying the `DeleteOnVmDeletion` attribute, you must specify the ID of the NIC attachment.
type LinkNicToUpdate struct {
	// If `true`, the NIC is deleted when the VM is terminated.
	DeleteOnVmDeletion *bool `json:"DeleteOnVmDeletion,omitempty"`
	// The ID of the NIC attachment.
	LinkNicId *string `json:"LinkNicId,omitempty"`
}

// GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field value if set, zero value otherwise.
func (o *LinkNicToUpdate) GetDeleteOnVmDeletion() bool {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret
	}
	return *o.DeleteOnVmDeletion
}

// GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkNicToUpdate) GetDeleteOnVmDeletionOk() (bool, bool) {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret, false
	}
	return *o.DeleteOnVmDeletion, true
}

// HasDeleteOnVmDeletion returns a boolean if a field has been set.
func (o *LinkNicToUpdate) HasDeleteOnVmDeletion() bool {
	if o != nil && o.DeleteOnVmDeletion != nil {
		return true
	}

	return false
}

// SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.
func (o *LinkNicToUpdate) SetDeleteOnVmDeletion(v bool) {
	o.DeleteOnVmDeletion = &v
}

// GetLinkNicId returns the LinkNicId field value if set, zero value otherwise.
func (o *LinkNicToUpdate) GetLinkNicId() string {
	if o == nil || o.LinkNicId == nil {
		var ret string
		return ret
	}
	return *o.LinkNicId
}

// GetLinkNicIdOk returns a tuple with the LinkNicId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkNicToUpdate) GetLinkNicIdOk() (string, bool) {
	if o == nil || o.LinkNicId == nil {
		var ret string
		return ret, false
	}
	return *o.LinkNicId, true
}

// HasLinkNicId returns a boolean if a field has been set.
func (o *LinkNicToUpdate) HasLinkNicId() bool {
	if o != nil && o.LinkNicId != nil {
		return true
	}

	return false
}

// SetLinkNicId gets a reference to the given string and assigns it to the LinkNicId field.
func (o *LinkNicToUpdate) SetLinkNicId(v string) {
	o.LinkNicId = &v
}

type NullableLinkNicToUpdate struct {
	Value LinkNicToUpdate
	ExplicitNull bool
}

func (v NullableLinkNicToUpdate) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableLinkNicToUpdate) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

