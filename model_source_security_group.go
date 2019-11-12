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

// SourceSecurityGroup Information about the source security group of the load balancer, which you can use as part of your inbound rules for your registered VMs.<br /> To only allow traffic from load balancers, add a security group rule that specifies this source security group as the inbound source.
type SourceSecurityGroup struct {
	// The account ID of the owner of the security group.
	SecurityGroupAccountId *string `json:"SecurityGroupAccountId,omitempty"`
	// (Public Cloud only) The name of the security group.
	SecurityGroupName *string `json:"SecurityGroupName,omitempty"`
}

// GetSecurityGroupAccountId returns the SecurityGroupAccountId field value if set, zero value otherwise.
func (o *SourceSecurityGroup) GetSecurityGroupAccountId() string {
	if o == nil || o.SecurityGroupAccountId == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupAccountId
}

// GetSecurityGroupAccountIdOk returns a tuple with the SecurityGroupAccountId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SourceSecurityGroup) GetSecurityGroupAccountIdOk() (string, bool) {
	if o == nil || o.SecurityGroupAccountId == nil {
		var ret string
		return ret, false
	}
	return *o.SecurityGroupAccountId, true
}

// HasSecurityGroupAccountId returns a boolean if a field has been set.
func (o *SourceSecurityGroup) HasSecurityGroupAccountId() bool {
	if o != nil && o.SecurityGroupAccountId != nil {
		return true
	}

	return false
}

// SetSecurityGroupAccountId gets a reference to the given string and assigns it to the SecurityGroupAccountId field.
func (o *SourceSecurityGroup) SetSecurityGroupAccountId(v string) {
	o.SecurityGroupAccountId = &v
}

// GetSecurityGroupName returns the SecurityGroupName field value if set, zero value otherwise.
func (o *SourceSecurityGroup) GetSecurityGroupName() string {
	if o == nil || o.SecurityGroupName == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupName
}

// GetSecurityGroupNameOk returns a tuple with the SecurityGroupName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SourceSecurityGroup) GetSecurityGroupNameOk() (string, bool) {
	if o == nil || o.SecurityGroupName == nil {
		var ret string
		return ret, false
	}
	return *o.SecurityGroupName, true
}

// HasSecurityGroupName returns a boolean if a field has been set.
func (o *SourceSecurityGroup) HasSecurityGroupName() bool {
	if o != nil && o.SecurityGroupName != nil {
		return true
	}

	return false
}

// SetSecurityGroupName gets a reference to the given string and assigns it to the SecurityGroupName field.
func (o *SourceSecurityGroup) SetSecurityGroupName(v string) {
	o.SecurityGroupName = &v
}

type NullableSourceSecurityGroup struct {
	Value SourceSecurityGroup
	ExplicitNull bool
}

func (v NullableSourceSecurityGroup) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableSourceSecurityGroup) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

