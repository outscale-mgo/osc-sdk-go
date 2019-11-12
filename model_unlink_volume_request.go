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

// UnlinkVolumeRequest struct for UnlinkVolumeRequest
type UnlinkVolumeRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// Forces the detachment of the volume in case of previous failure. Important: This action may damage your data or file systems.
	ForceUnlink *bool `json:"ForceUnlink,omitempty"`
	// The ID of the volume you want to detach.
	VolumeId string `json:"VolumeId"`
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UnlinkVolumeRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkVolumeRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UnlinkVolumeRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UnlinkVolumeRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetForceUnlink returns the ForceUnlink field value if set, zero value otherwise.
func (o *UnlinkVolumeRequest) GetForceUnlink() bool {
	if o == nil || o.ForceUnlink == nil {
		var ret bool
		return ret
	}
	return *o.ForceUnlink
}

// GetForceUnlinkOk returns a tuple with the ForceUnlink field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkVolumeRequest) GetForceUnlinkOk() (bool, bool) {
	if o == nil || o.ForceUnlink == nil {
		var ret bool
		return ret, false
	}
	return *o.ForceUnlink, true
}

// HasForceUnlink returns a boolean if a field has been set.
func (o *UnlinkVolumeRequest) HasForceUnlink() bool {
	if o != nil && o.ForceUnlink != nil {
		return true
	}

	return false
}

// SetForceUnlink gets a reference to the given bool and assigns it to the ForceUnlink field.
func (o *UnlinkVolumeRequest) SetForceUnlink(v bool) {
	o.ForceUnlink = &v
}

// GetVolumeId returns the VolumeId field value
func (o *UnlinkVolumeRequest) GetVolumeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeId
}

// SetVolumeId sets field value
func (o *UnlinkVolumeRequest) SetVolumeId(v string) {
	o.VolumeId = v
}

type NullableUnlinkVolumeRequest struct {
	Value UnlinkVolumeRequest
	ExplicitNull bool
}

func (v NullableUnlinkVolumeRequest) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableUnlinkVolumeRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

