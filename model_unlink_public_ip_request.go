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

// UnlinkPublicIpRequest struct for UnlinkPublicIpRequest
type UnlinkPublicIpRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// (Required in a Net) The ID representing the association of the EIP with the VM or the NIC.
	LinkPublicIpId *string `json:"LinkPublicIpId,omitempty"`
	// The External IP address. In the public Cloud, this parameter is required.
	PublicIp *string `json:"PublicIp,omitempty"`
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UnlinkPublicIpRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkPublicIpRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UnlinkPublicIpRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UnlinkPublicIpRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetLinkPublicIpId returns the LinkPublicIpId field value if set, zero value otherwise.
func (o *UnlinkPublicIpRequest) GetLinkPublicIpId() string {
	if o == nil || o.LinkPublicIpId == nil {
		var ret string
		return ret
	}
	return *o.LinkPublicIpId
}

// GetLinkPublicIpIdOk returns a tuple with the LinkPublicIpId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkPublicIpRequest) GetLinkPublicIpIdOk() (string, bool) {
	if o == nil || o.LinkPublicIpId == nil {
		var ret string
		return ret, false
	}
	return *o.LinkPublicIpId, true
}

// HasLinkPublicIpId returns a boolean if a field has been set.
func (o *UnlinkPublicIpRequest) HasLinkPublicIpId() bool {
	if o != nil && o.LinkPublicIpId != nil {
		return true
	}

	return false
}

// SetLinkPublicIpId gets a reference to the given string and assigns it to the LinkPublicIpId field.
func (o *UnlinkPublicIpRequest) SetLinkPublicIpId(v string) {
	o.LinkPublicIpId = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *UnlinkPublicIpRequest) GetPublicIp() string {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UnlinkPublicIpRequest) GetPublicIpOk() (string, bool) {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret, false
	}
	return *o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *UnlinkPublicIpRequest) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *UnlinkPublicIpRequest) SetPublicIp(v string) {
	o.PublicIp = &v
}

type NullableUnlinkPublicIpRequest struct {
	Value        UnlinkPublicIpRequest
	ExplicitNull bool
}

func (v NullableUnlinkPublicIpRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUnlinkPublicIpRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
