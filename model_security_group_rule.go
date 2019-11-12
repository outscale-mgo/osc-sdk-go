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

// SecurityGroupRule Information about the security group rule.
type SecurityGroupRule struct {
	// The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.
	FromPortRange *int32 `json:"FromPortRange,omitempty"`
	// The IP protocol name (`tcp`, `udp`, `icmp`) or protocol number. By default, `-1`, which means all protocols.
	IpProtocol *string `json:"IpProtocol,omitempty"`
	// One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).
	IpRanges *[]string `json:"IpRanges,omitempty"`
	// Information about one or more members of a security group.
	SecurityGroupsMembers *[]SecurityGroupsMember `json:"SecurityGroupsMembers,omitempty"`
	// One or more service IDs to allow traffic from a Net to access the corresponding 3DS OUTSCALE services. For more information, see [ReadNetAccessPointServices](#readnetaccesspointservices).
	ServiceIds *[]string `json:"ServiceIds,omitempty"`
	// The end of the port range for the TCP and UDP protocols, or an ICMP type number.
	ToPortRange *int32 `json:"ToPortRange,omitempty"`
}

// GetFromPortRange returns the FromPortRange field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetFromPortRange() int32 {
	if o == nil || o.FromPortRange == nil {
		var ret int32
		return ret
	}
	return *o.FromPortRange
}

// GetFromPortRangeOk returns a tuple with the FromPortRange field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetFromPortRangeOk() (int32, bool) {
	if o == nil || o.FromPortRange == nil {
		var ret int32
		return ret, false
	}
	return *o.FromPortRange, true
}

// HasFromPortRange returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasFromPortRange() bool {
	if o != nil && o.FromPortRange != nil {
		return true
	}

	return false
}

// SetFromPortRange gets a reference to the given int32 and assigns it to the FromPortRange field.
func (o *SecurityGroupRule) SetFromPortRange(v int32) {
	o.FromPortRange = &v
}

// GetIpProtocol returns the IpProtocol field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetIpProtocol() string {
	if o == nil || o.IpProtocol == nil {
		var ret string
		return ret
	}
	return *o.IpProtocol
}

// GetIpProtocolOk returns a tuple with the IpProtocol field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetIpProtocolOk() (string, bool) {
	if o == nil || o.IpProtocol == nil {
		var ret string
		return ret, false
	}
	return *o.IpProtocol, true
}

// HasIpProtocol returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasIpProtocol() bool {
	if o != nil && o.IpProtocol != nil {
		return true
	}

	return false
}

// SetIpProtocol gets a reference to the given string and assigns it to the IpProtocol field.
func (o *SecurityGroupRule) SetIpProtocol(v string) {
	o.IpProtocol = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetIpRanges() []string {
	if o == nil || o.IpRanges == nil {
		var ret []string
		return ret
	}
	return *o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetIpRangesOk() ([]string, bool) {
	if o == nil || o.IpRanges == nil {
		var ret []string
		return ret, false
	}
	return *o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasIpRanges() bool {
	if o != nil && o.IpRanges != nil {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *SecurityGroupRule) SetIpRanges(v []string) {
	o.IpRanges = &v
}

// GetSecurityGroupsMembers returns the SecurityGroupsMembers field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetSecurityGroupsMembers() []SecurityGroupsMember {
	if o == nil || o.SecurityGroupsMembers == nil {
		var ret []SecurityGroupsMember
		return ret
	}
	return *o.SecurityGroupsMembers
}

// GetSecurityGroupsMembersOk returns a tuple with the SecurityGroupsMembers field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetSecurityGroupsMembersOk() ([]SecurityGroupsMember, bool) {
	if o == nil || o.SecurityGroupsMembers == nil {
		var ret []SecurityGroupsMember
		return ret, false
	}
	return *o.SecurityGroupsMembers, true
}

// HasSecurityGroupsMembers returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasSecurityGroupsMembers() bool {
	if o != nil && o.SecurityGroupsMembers != nil {
		return true
	}

	return false
}

// SetSecurityGroupsMembers gets a reference to the given []SecurityGroupsMember and assigns it to the SecurityGroupsMembers field.
func (o *SecurityGroupRule) SetSecurityGroupsMembers(v []SecurityGroupsMember) {
	o.SecurityGroupsMembers = &v
}

// GetServiceIds returns the ServiceIds field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetServiceIds() []string {
	if o == nil || o.ServiceIds == nil {
		var ret []string
		return ret
	}
	return *o.ServiceIds
}

// GetServiceIdsOk returns a tuple with the ServiceIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetServiceIdsOk() ([]string, bool) {
	if o == nil || o.ServiceIds == nil {
		var ret []string
		return ret, false
	}
	return *o.ServiceIds, true
}

// HasServiceIds returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasServiceIds() bool {
	if o != nil && o.ServiceIds != nil {
		return true
	}

	return false
}

// SetServiceIds gets a reference to the given []string and assigns it to the ServiceIds field.
func (o *SecurityGroupRule) SetServiceIds(v []string) {
	o.ServiceIds = &v
}

// GetToPortRange returns the ToPortRange field value if set, zero value otherwise.
func (o *SecurityGroupRule) GetToPortRange() int32 {
	if o == nil || o.ToPortRange == nil {
		var ret int32
		return ret
	}
	return *o.ToPortRange
}

// GetToPortRangeOk returns a tuple with the ToPortRange field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupRule) GetToPortRangeOk() (int32, bool) {
	if o == nil || o.ToPortRange == nil {
		var ret int32
		return ret, false
	}
	return *o.ToPortRange, true
}

// HasToPortRange returns a boolean if a field has been set.
func (o *SecurityGroupRule) HasToPortRange() bool {
	if o != nil && o.ToPortRange != nil {
		return true
	}

	return false
}

// SetToPortRange gets a reference to the given int32 and assigns it to the ToPortRange field.
func (o *SecurityGroupRule) SetToPortRange(v int32) {
	o.ToPortRange = &v
}

type NullableSecurityGroupRule struct {
	Value SecurityGroupRule
	ExplicitNull bool
}

func (v NullableSecurityGroupRule) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableSecurityGroupRule) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

