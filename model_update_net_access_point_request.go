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

// UpdateNetAccessPointRequest struct for UpdateNetAccessPointRequest
type UpdateNetAccessPointRequest struct {
	// One or more IDs of route tables to associate with the specified Net access point.
	AddRouteTableIds *[]string `json:"AddRouteTableIds,omitempty"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the Net access point.
	NetAccessPointId string `json:"NetAccessPointId"`
	// One or more IDs of route tables to disassociate from the specified Net access point.
	RemoveRouteTableIds *[]string `json:"RemoveRouteTableIds,omitempty"`
}

// GetAddRouteTableIds returns the AddRouteTableIds field value if set, zero value otherwise.
func (o *UpdateNetAccessPointRequest) GetAddRouteTableIds() []string {
	if o == nil || o.AddRouteTableIds == nil {
		var ret []string
		return ret
	}
	return *o.AddRouteTableIds
}

// GetAddRouteTableIdsOk returns a tuple with the AddRouteTableIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetAccessPointRequest) GetAddRouteTableIdsOk() ([]string, bool) {
	if o == nil || o.AddRouteTableIds == nil {
		var ret []string
		return ret, false
	}
	return *o.AddRouteTableIds, true
}

// HasAddRouteTableIds returns a boolean if a field has been set.
func (o *UpdateNetAccessPointRequest) HasAddRouteTableIds() bool {
	if o != nil && o.AddRouteTableIds != nil {
		return true
	}

	return false
}

// SetAddRouteTableIds gets a reference to the given []string and assigns it to the AddRouteTableIds field.
func (o *UpdateNetAccessPointRequest) SetAddRouteTableIds(v []string) {
	o.AddRouteTableIds = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateNetAccessPointRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetAccessPointRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateNetAccessPointRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateNetAccessPointRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetNetAccessPointId returns the NetAccessPointId field value
func (o *UpdateNetAccessPointRequest) GetNetAccessPointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetAccessPointId
}

// SetNetAccessPointId sets field value
func (o *UpdateNetAccessPointRequest) SetNetAccessPointId(v string) {
	o.NetAccessPointId = v
}

// GetRemoveRouteTableIds returns the RemoveRouteTableIds field value if set, zero value otherwise.
func (o *UpdateNetAccessPointRequest) GetRemoveRouteTableIds() []string {
	if o == nil || o.RemoveRouteTableIds == nil {
		var ret []string
		return ret
	}
	return *o.RemoveRouteTableIds
}

// GetRemoveRouteTableIdsOk returns a tuple with the RemoveRouteTableIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetAccessPointRequest) GetRemoveRouteTableIdsOk() ([]string, bool) {
	if o == nil || o.RemoveRouteTableIds == nil {
		var ret []string
		return ret, false
	}
	return *o.RemoveRouteTableIds, true
}

// HasRemoveRouteTableIds returns a boolean if a field has been set.
func (o *UpdateNetAccessPointRequest) HasRemoveRouteTableIds() bool {
	if o != nil && o.RemoveRouteTableIds != nil {
		return true
	}

	return false
}

// SetRemoveRouteTableIds gets a reference to the given []string and assigns it to the RemoveRouteTableIds field.
func (o *UpdateNetAccessPointRequest) SetRemoveRouteTableIds(v []string) {
	o.RemoveRouteTableIds = &v
}

type NullableUpdateNetAccessPointRequest struct {
	Value UpdateNetAccessPointRequest
	ExplicitNull bool
}

func (v NullableUpdateNetAccessPointRequest) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableUpdateNetAccessPointRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

