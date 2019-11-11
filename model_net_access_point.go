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

// NetAccessPoint Information about the Net access point.
type NetAccessPoint struct {
	// The ID of the Net access point.
	NetAccessPointId *string `json:"NetAccessPointId,omitempty"`
	// The ID of the Net with which the Net access point is associated.
	NetId *string `json:"NetId,omitempty"`
	// The ID of the route tables associated with the Net access point.
	RouteTableIds *[]string `json:"RouteTableIds,omitempty"`
	// The name of the prefix list corresponding to the service with which the Net access point is associated.
	ServiceName *string `json:"ServiceName,omitempty"`
	// The state of the Net access point (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	State *string `json:"State,omitempty"`
	// One or more tags associated with the Net access point.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
}

// GetNetAccessPointId returns the NetAccessPointId field value if set, zero value otherwise.
func (o *NetAccessPoint) GetNetAccessPointId() string {
	if o == nil || o.NetAccessPointId == nil {
		var ret string
		return ret
	}
	return *o.NetAccessPointId
}

// GetNetAccessPointIdOk returns a tuple with the NetAccessPointId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NetAccessPoint) GetNetAccessPointIdOk() (string, bool) {
	if o == nil || o.NetAccessPointId == nil {
		var ret string
		return ret, false
	}
	return *o.NetAccessPointId, true
}

// HasNetAccessPointId returns a boolean if a field has been set.
func (o *NetAccessPoint) HasNetAccessPointId() bool {
	if o != nil && o.NetAccessPointId != nil {
		return true
	}

	return false
}

// SetNetAccessPointId gets a reference to the given string and assigns it to the NetAccessPointId field.
func (o *NetAccessPoint) SetNetAccessPointId(v string) {
	o.NetAccessPointId = &v
}

// GetNetId returns the NetId field value if set, zero value otherwise.
func (o *NetAccessPoint) GetNetId() string {
	if o == nil || o.NetId == nil {
		var ret string
		return ret
	}
	return *o.NetId
}

// GetNetIdOk returns a tuple with the NetId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NetAccessPoint) GetNetIdOk() (string, bool) {
	if o == nil || o.NetId == nil {
		var ret string
		return ret, false
	}
	return *o.NetId, true
}

// HasNetId returns a boolean if a field has been set.
func (o *NetAccessPoint) HasNetId() bool {
	if o != nil && o.NetId != nil {
		return true
	}

	return false
}

// SetNetId gets a reference to the given string and assigns it to the NetId field.
func (o *NetAccessPoint) SetNetId(v string) {
	o.NetId = &v
}

// GetRouteTableIds returns the RouteTableIds field value if set, zero value otherwise.
func (o *NetAccessPoint) GetRouteTableIds() []string {
	if o == nil || o.RouteTableIds == nil {
		var ret []string
		return ret
	}
	return *o.RouteTableIds
}

// GetRouteTableIdsOk returns a tuple with the RouteTableIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NetAccessPoint) GetRouteTableIdsOk() ([]string, bool) {
	if o == nil || o.RouteTableIds == nil {
		var ret []string
		return ret, false
	}
	return *o.RouteTableIds, true
}

// HasRouteTableIds returns a boolean if a field has been set.
func (o *NetAccessPoint) HasRouteTableIds() bool {
	if o != nil && o.RouteTableIds != nil {
		return true
	}

	return false
}

// SetRouteTableIds gets a reference to the given []string and assigns it to the RouteTableIds field.
func (o *NetAccessPoint) SetRouteTableIds(v []string) {
	o.RouteTableIds = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *NetAccessPoint) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NetAccessPoint) GetServiceNameOk() (string, bool) {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret, false
	}
	return *o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *NetAccessPoint) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *NetAccessPoint) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NetAccessPoint) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NetAccessPoint) GetStateOk() (string, bool) {
	if o == nil || o.State == nil {
		var ret string
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NetAccessPoint) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NetAccessPoint) SetState(v string) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *NetAccessPoint) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NetAccessPoint) GetTagsOk() ([]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *NetAccessPoint) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *NetAccessPoint) SetTags(v []ResourceTag) {
	o.Tags = &v
}

type NullableNetAccessPoint struct {
	Value NetAccessPoint
	ExplicitNull bool
}

func (v NullableNetAccessPoint) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableNetAccessPoint) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

