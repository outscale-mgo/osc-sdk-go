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

// Subnet Information about the Subnet.
type Subnet struct {
	// The number of available IP addresses in the Subnets.
	AvailableIpsCount *int32 `json:"AvailableIpsCount,omitempty"`
	// The IP range in the Subnet, in CIDR notation (for example, 10.0.0.0/16).
	IpRange *string `json:"IpRange,omitempty"`
	// If `true`, a public IP address is assigned to the network interface cards (NICs) created in the specified Subnet.
	MapPublicIpOnLaunch *bool `json:"MapPublicIpOnLaunch,omitempty"`
	// The ID of the Net in which the Subnet is.
	NetId *string `json:"NetId,omitempty"`
	// The state of the Subnet (`pending` \\| `available`).
	State *string `json:"State,omitempty"`
	// The ID of the Subnet.
	SubnetId *string `json:"SubnetId,omitempty"`
	// The name of the Subregion in which the Subnet is located.
	SubregionName *string `json:"SubregionName,omitempty"`
	// One or more tags associated with the Subnet.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
}

// GetAvailableIpsCount returns the AvailableIpsCount field value if set, zero value otherwise.
func (o *Subnet) GetAvailableIpsCount() int32 {
	if o == nil || o.AvailableIpsCount == nil {
		var ret int32
		return ret
	}
	return *o.AvailableIpsCount
}

// GetAvailableIpsCountOk returns a tuple with the AvailableIpsCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Subnet) GetAvailableIpsCountOk() (int32, bool) {
	if o == nil || o.AvailableIpsCount == nil {
		var ret int32
		return ret, false
	}
	return *o.AvailableIpsCount, true
}

// HasAvailableIpsCount returns a boolean if a field has been set.
func (o *Subnet) HasAvailableIpsCount() bool {
	if o != nil && o.AvailableIpsCount != nil {
		return true
	}

	return false
}

// SetAvailableIpsCount gets a reference to the given int32 and assigns it to the AvailableIpsCount field.
func (o *Subnet) SetAvailableIpsCount(v int32) {
	o.AvailableIpsCount = &v
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *Subnet) GetIpRange() string {
	if o == nil || o.IpRange == nil {
		var ret string
		return ret
	}
	return *o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Subnet) GetIpRangeOk() (string, bool) {
	if o == nil || o.IpRange == nil {
		var ret string
		return ret, false
	}
	return *o.IpRange, true
}

// HasIpRange returns a boolean if a field has been set.
func (o *Subnet) HasIpRange() bool {
	if o != nil && o.IpRange != nil {
		return true
	}

	return false
}

// SetIpRange gets a reference to the given string and assigns it to the IpRange field.
func (o *Subnet) SetIpRange(v string) {
	o.IpRange = &v
}

// GetMapPublicIpOnLaunch returns the MapPublicIpOnLaunch field value if set, zero value otherwise.
func (o *Subnet) GetMapPublicIpOnLaunch() bool {
	if o == nil || o.MapPublicIpOnLaunch == nil {
		var ret bool
		return ret
	}
	return *o.MapPublicIpOnLaunch
}

// GetMapPublicIpOnLaunchOk returns a tuple with the MapPublicIpOnLaunch field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Subnet) GetMapPublicIpOnLaunchOk() (bool, bool) {
	if o == nil || o.MapPublicIpOnLaunch == nil {
		var ret bool
		return ret, false
	}
	return *o.MapPublicIpOnLaunch, true
}

// HasMapPublicIpOnLaunch returns a boolean if a field has been set.
func (o *Subnet) HasMapPublicIpOnLaunch() bool {
	if o != nil && o.MapPublicIpOnLaunch != nil {
		return true
	}

	return false
}

// SetMapPublicIpOnLaunch gets a reference to the given bool and assigns it to the MapPublicIpOnLaunch field.
func (o *Subnet) SetMapPublicIpOnLaunch(v bool) {
	o.MapPublicIpOnLaunch = &v
}

// GetNetId returns the NetId field value if set, zero value otherwise.
func (o *Subnet) GetNetId() string {
	if o == nil || o.NetId == nil {
		var ret string
		return ret
	}
	return *o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Subnet) GetNetIdOk() (string, bool) {
	if o == nil || o.NetId == nil {
		var ret string
		return ret, false
	}
	return *o.NetId, true
}

// HasNetId returns a boolean if a field has been set.
func (o *Subnet) HasNetId() bool {
	if o != nil && o.NetId != nil {
		return true
	}

	return false
}

// SetNetId gets a reference to the given string and assigns it to the NetId field.
func (o *Subnet) SetNetId(v string) {
	o.NetId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Subnet) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Subnet) GetStateOk() (string, bool) {
	if o == nil || o.State == nil {
		var ret string
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Subnet) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Subnet) SetState(v string) {
	o.State = &v
}

// GetSubnetId returns the SubnetId field value if set, zero value otherwise.
func (o *Subnet) GetSubnetId() string {
	if o == nil || o.SubnetId == nil {
		var ret string
		return ret
	}
	return *o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Subnet) GetSubnetIdOk() (string, bool) {
	if o == nil || o.SubnetId == nil {
		var ret string
		return ret, false
	}
	return *o.SubnetId, true
}

// HasSubnetId returns a boolean if a field has been set.
func (o *Subnet) HasSubnetId() bool {
	if o != nil && o.SubnetId != nil {
		return true
	}

	return false
}

// SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.
func (o *Subnet) SetSubnetId(v string) {
	o.SubnetId = &v
}

// GetSubregionName returns the SubregionName field value if set, zero value otherwise.
func (o *Subnet) GetSubregionName() string {
	if o == nil || o.SubregionName == nil {
		var ret string
		return ret
	}
	return *o.SubregionName
}

// GetSubregionNameOk returns a tuple with the SubregionName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Subnet) GetSubregionNameOk() (string, bool) {
	if o == nil || o.SubregionName == nil {
		var ret string
		return ret, false
	}
	return *o.SubregionName, true
}

// HasSubregionName returns a boolean if a field has been set.
func (o *Subnet) HasSubregionName() bool {
	if o != nil && o.SubregionName != nil {
		return true
	}

	return false
}

// SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.
func (o *Subnet) SetSubregionName(v string) {
	o.SubregionName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Subnet) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Subnet) GetTagsOk() ([]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Subnet) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *Subnet) SetTags(v []ResourceTag) {
	o.Tags = &v
}

type NullableSubnet struct {
	Value Subnet
	ExplicitNull bool
}

func (v NullableSubnet) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableSubnet) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

