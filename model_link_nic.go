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

// LinkNic Information about the NIC attachment.
type LinkNic struct {
	// If `true`, the volume is deleted when the VM is terminated.
	DeleteOnVmDeletion *bool `json:"DeleteOnVmDeletion,omitempty"`
	// The device index for the NIC attachment (between 1 and 7, both included).
	DeviceNumber *int32 `json:"DeviceNumber,omitempty"`
	// The ID of the NIC to attach.
	LinkNicId *string `json:"LinkNicId,omitempty"`
	// The state of the attachment (`attaching` \\| `attached` \\| `detaching` \\| `detached`).
	State *string `json:"State,omitempty"`
	// The account ID of the owner of the VM.
	VmAccountId *string `json:"VmAccountId,omitempty"`
	// The ID of the VM.
	VmId *string `json:"VmId,omitempty"`
}

// GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field value if set, zero value otherwise.
func (o *LinkNic) GetDeleteOnVmDeletion() bool {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret
	}
	return *o.DeleteOnVmDeletion
}

// GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkNic) GetDeleteOnVmDeletionOk() (bool, bool) {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret, false
	}
	return *o.DeleteOnVmDeletion, true
}

// HasDeleteOnVmDeletion returns a boolean if a field has been set.
func (o *LinkNic) HasDeleteOnVmDeletion() bool {
	if o != nil && o.DeleteOnVmDeletion != nil {
		return true
	}

	return false
}

// SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.
func (o *LinkNic) SetDeleteOnVmDeletion(v bool) {
	o.DeleteOnVmDeletion = &v
}

// GetDeviceNumber returns the DeviceNumber field value if set, zero value otherwise.
func (o *LinkNic) GetDeviceNumber() int32 {
	if o == nil || o.DeviceNumber == nil {
		var ret int32
		return ret
	}
	return *o.DeviceNumber
}

// GetDeviceNumberOk returns a tuple with the DeviceNumber field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkNic) GetDeviceNumberOk() (int32, bool) {
	if o == nil || o.DeviceNumber == nil {
		var ret int32
		return ret, false
	}
	return *o.DeviceNumber, true
}

// HasDeviceNumber returns a boolean if a field has been set.
func (o *LinkNic) HasDeviceNumber() bool {
	if o != nil && o.DeviceNumber != nil {
		return true
	}

	return false
}

// SetDeviceNumber gets a reference to the given int32 and assigns it to the DeviceNumber field.
func (o *LinkNic) SetDeviceNumber(v int32) {
	o.DeviceNumber = &v
}

// GetLinkNicId returns the LinkNicId field value if set, zero value otherwise.
func (o *LinkNic) GetLinkNicId() string {
	if o == nil || o.LinkNicId == nil {
		var ret string
		return ret
	}
	return *o.LinkNicId
}

// GetLinkNicIdOk returns a tuple with the LinkNicId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkNic) GetLinkNicIdOk() (string, bool) {
	if o == nil || o.LinkNicId == nil {
		var ret string
		return ret, false
	}
	return *o.LinkNicId, true
}

// HasLinkNicId returns a boolean if a field has been set.
func (o *LinkNic) HasLinkNicId() bool {
	if o != nil && o.LinkNicId != nil {
		return true
	}

	return false
}

// SetLinkNicId gets a reference to the given string and assigns it to the LinkNicId field.
func (o *LinkNic) SetLinkNicId(v string) {
	o.LinkNicId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *LinkNic) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkNic) GetStateOk() (string, bool) {
	if o == nil || o.State == nil {
		var ret string
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *LinkNic) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *LinkNic) SetState(v string) {
	o.State = &v
}

// GetVmAccountId returns the VmAccountId field value if set, zero value otherwise.
func (o *LinkNic) GetVmAccountId() string {
	if o == nil || o.VmAccountId == nil {
		var ret string
		return ret
	}
	return *o.VmAccountId
}

// GetVmAccountIdOk returns a tuple with the VmAccountId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkNic) GetVmAccountIdOk() (string, bool) {
	if o == nil || o.VmAccountId == nil {
		var ret string
		return ret, false
	}
	return *o.VmAccountId, true
}

// HasVmAccountId returns a boolean if a field has been set.
func (o *LinkNic) HasVmAccountId() bool {
	if o != nil && o.VmAccountId != nil {
		return true
	}

	return false
}

// SetVmAccountId gets a reference to the given string and assigns it to the VmAccountId field.
func (o *LinkNic) SetVmAccountId(v string) {
	o.VmAccountId = &v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *LinkNic) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkNic) GetVmIdOk() (string, bool) {
	if o == nil || o.VmId == nil {
		var ret string
		return ret, false
	}
	return *o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *LinkNic) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *LinkNic) SetVmId(v string) {
	o.VmId = &v
}

type NullableLinkNic struct {
	Value LinkNic
	ExplicitNull bool
}

func (v NullableLinkNic) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableLinkNic) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

