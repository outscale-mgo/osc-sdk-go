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

// CreateNicRequest struct for CreateNicRequest
type CreateNicRequest struct {
	// A description for the NIC.
	Description *string `json:"Description,omitempty"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The primary private IP address for the NIC.<br /><br />  This IP address must be within the IP address range of the Subnet that you specify with the `SubnetId` attribute.<br /> If you do not specify this attribute, a random private IP address is selected within the IP address range of the Subnet.
	PrivateIps *[]PrivateIpLight `json:"PrivateIps,omitempty"`
	// One or more IDs of security groups for the NIC.
	SecurityGroupIds *[]string `json:"SecurityGroupIds,omitempty"`
	// The ID of the Subnet in which you want to create the NIC.
	SubnetId string `json:"SubnetId"`
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateNicRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateNicRequest) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateNicRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateNicRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateNicRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateNicRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateNicRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateNicRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetPrivateIps returns the PrivateIps field value if set, zero value otherwise.
func (o *CreateNicRequest) GetPrivateIps() []PrivateIpLight {
	if o == nil || o.PrivateIps == nil {
		var ret []PrivateIpLight
		return ret
	}
	return *o.PrivateIps
}

// GetPrivateIpsOk returns a tuple with the PrivateIps field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateNicRequest) GetPrivateIpsOk() ([]PrivateIpLight, bool) {
	if o == nil || o.PrivateIps == nil {
		var ret []PrivateIpLight
		return ret, false
	}
	return *o.PrivateIps, true
}

// HasPrivateIps returns a boolean if a field has been set.
func (o *CreateNicRequest) HasPrivateIps() bool {
	if o != nil && o.PrivateIps != nil {
		return true
	}

	return false
}

// SetPrivateIps gets a reference to the given []PrivateIpLight and assigns it to the PrivateIps field.
func (o *CreateNicRequest) SetPrivateIps(v []PrivateIpLight) {
	o.PrivateIps = &v
}

// GetSecurityGroupIds returns the SecurityGroupIds field value if set, zero value otherwise.
func (o *CreateNicRequest) GetSecurityGroupIds() []string {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateNicRequest) GetSecurityGroupIdsOk() ([]string, bool) {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []string
		return ret, false
	}
	return *o.SecurityGroupIds, true
}

// HasSecurityGroupIds returns a boolean if a field has been set.
func (o *CreateNicRequest) HasSecurityGroupIds() bool {
	if o != nil && o.SecurityGroupIds != nil {
		return true
	}

	return false
}

// SetSecurityGroupIds gets a reference to the given []string and assigns it to the SecurityGroupIds field.
func (o *CreateNicRequest) SetSecurityGroupIds(v []string) {
	o.SecurityGroupIds = &v
}

// GetSubnetId returns the SubnetId field value
func (o *CreateNicRequest) GetSubnetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetId
}

// SetSubnetId sets field value
func (o *CreateNicRequest) SetSubnetId(v string) {
	o.SubnetId = v
}

type NullableCreateNicRequest struct {
	Value        CreateNicRequest
	ExplicitNull bool
}

func (v NullableCreateNicRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCreateNicRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
