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

// LinkPublicIpLightForVm Information about the EIP associated with the NIC.
type LinkPublicIpLightForVm struct {
	// The name of the public DNS.
	PublicDnsName *string `json:"PublicDnsName,omitempty"`
	// The External IP address (EIP) associated with the NIC.
	PublicIp *string `json:"PublicIp,omitempty"`
	// The account ID of the owner of the EIP.
	PublicIpAccountId *string `json:"PublicIpAccountId,omitempty"`
}

// GetPublicDnsName returns the PublicDnsName field value if set, zero value otherwise.
func (o *LinkPublicIpLightForVm) GetPublicDnsName() string {
	if o == nil || o.PublicDnsName == nil {
		var ret string
		return ret
	}
	return *o.PublicDnsName
}

// GetPublicDnsNameOk returns a tuple with the PublicDnsName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIpLightForVm) GetPublicDnsNameOk() (string, bool) {
	if o == nil || o.PublicDnsName == nil {
		var ret string
		return ret, false
	}
	return *o.PublicDnsName, true
}

// HasPublicDnsName returns a boolean if a field has been set.
func (o *LinkPublicIpLightForVm) HasPublicDnsName() bool {
	if o != nil && o.PublicDnsName != nil {
		return true
	}

	return false
}

// SetPublicDnsName gets a reference to the given string and assigns it to the PublicDnsName field.
func (o *LinkPublicIpLightForVm) SetPublicDnsName(v string) {
	o.PublicDnsName = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *LinkPublicIpLightForVm) GetPublicIp() string {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIpLightForVm) GetPublicIpOk() (string, bool) {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret, false
	}
	return *o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *LinkPublicIpLightForVm) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *LinkPublicIpLightForVm) SetPublicIp(v string) {
	o.PublicIp = &v
}

// GetPublicIpAccountId returns the PublicIpAccountId field value if set, zero value otherwise.
func (o *LinkPublicIpLightForVm) GetPublicIpAccountId() string {
	if o == nil || o.PublicIpAccountId == nil {
		var ret string
		return ret
	}
	return *o.PublicIpAccountId
}

// GetPublicIpAccountIdOk returns a tuple with the PublicIpAccountId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LinkPublicIpLightForVm) GetPublicIpAccountIdOk() (string, bool) {
	if o == nil || o.PublicIpAccountId == nil {
		var ret string
		return ret, false
	}
	return *o.PublicIpAccountId, true
}

// HasPublicIpAccountId returns a boolean if a field has been set.
func (o *LinkPublicIpLightForVm) HasPublicIpAccountId() bool {
	if o != nil && o.PublicIpAccountId != nil {
		return true
	}

	return false
}

// SetPublicIpAccountId gets a reference to the given string and assigns it to the PublicIpAccountId field.
func (o *LinkPublicIpLightForVm) SetPublicIpAccountId(v string) {
	o.PublicIpAccountId = &v
}

type NullableLinkPublicIpLightForVm struct {
	Value        LinkPublicIpLightForVm
	ExplicitNull bool
}

func (v NullableLinkPublicIpLightForVm) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLinkPublicIpLightForVm) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
