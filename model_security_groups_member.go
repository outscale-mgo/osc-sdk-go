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

// SecurityGroupsMember Information about the member of a security group.
type SecurityGroupsMember struct {
	// The account ID of a user.
	AccountId *string `json:"AccountId,omitempty"`
	// The ID of the security group.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty"`
	// (Public Cloud only) The name of the security group.
	SecurityGroupName *string `json:"SecurityGroupName,omitempty"`
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *SecurityGroupsMember) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupsMember) GetAccountIdOk() (string, bool) {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret, false
	}
	return *o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *SecurityGroupsMember) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *SecurityGroupsMember) SetAccountId(v string) {
	o.AccountId = &v
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise.
func (o *SecurityGroupsMember) GetSecurityGroupId() string {
	if o == nil || o.SecurityGroupId == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupsMember) GetSecurityGroupIdOk() (string, bool) {
	if o == nil || o.SecurityGroupId == nil {
		var ret string
		return ret, false
	}
	return *o.SecurityGroupId, true
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *SecurityGroupsMember) HasSecurityGroupId() bool {
	if o != nil && o.SecurityGroupId != nil {
		return true
	}

	return false
}

// SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.
func (o *SecurityGroupsMember) SetSecurityGroupId(v string) {
	o.SecurityGroupId = &v
}

// GetSecurityGroupName returns the SecurityGroupName field value if set, zero value otherwise.
func (o *SecurityGroupsMember) GetSecurityGroupName() string {
	if o == nil || o.SecurityGroupName == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupName
}

// GetSecurityGroupNameOk returns a tuple with the SecurityGroupName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupsMember) GetSecurityGroupNameOk() (string, bool) {
	if o == nil || o.SecurityGroupName == nil {
		var ret string
		return ret, false
	}
	return *o.SecurityGroupName, true
}

// HasSecurityGroupName returns a boolean if a field has been set.
func (o *SecurityGroupsMember) HasSecurityGroupName() bool {
	if o != nil && o.SecurityGroupName != nil {
		return true
	}

	return false
}

// SetSecurityGroupName gets a reference to the given string and assigns it to the SecurityGroupName field.
func (o *SecurityGroupsMember) SetSecurityGroupName(v string) {
	o.SecurityGroupName = &v
}

type NullableSecurityGroupsMember struct {
	Value        SecurityGroupsMember
	ExplicitNull bool
}

func (v NullableSecurityGroupsMember) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSecurityGroupsMember) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
