/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.1
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oscgo

import (
	"bytes"
	"encoding/json"
)

// LinkNicRequest struct for LinkNicRequest
type LinkNicRequest struct {
	// The index of the VM device for the NIC attachment (between 1 and 7, both included).
	DeviceNumber int64 `json:"DeviceNumber"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the NIC you want to attach.
	NicId string `json:"NicId"`
	// The ID of the VM to which you want to attach the NIC.
	VmId string `json:"VmId"`
}

// GetDeviceNumber returns the DeviceNumber field value
func (o *LinkNicRequest) GetDeviceNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DeviceNumber
}

// SetDeviceNumber sets field value
func (o *LinkNicRequest) SetDeviceNumber(v int64) {
	o.DeviceNumber = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *LinkNicRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkNicRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *LinkNicRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *LinkNicRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNicId returns the NicId field value
func (o *LinkNicRequest) GetNicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NicId
}

// SetNicId sets field value
func (o *LinkNicRequest) SetNicId(v string) {
	o.NicId = v
}

// GetVmId returns the VmId field value
func (o *LinkNicRequest) GetVmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VmId
}

// SetVmId sets field value
func (o *LinkNicRequest) SetVmId(v string) {
	o.VmId = v
}

type NullableLinkNicRequest struct {
	Value        LinkNicRequest
	ExplicitNull bool
}

func (v NullableLinkNicRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLinkNicRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
