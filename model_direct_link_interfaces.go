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

// DirectLinkInterfaces Information about the DirectLink interfaces.
type DirectLinkInterfaces struct {
	// The account ID of the owner of the DirectLink interface.
	AccountId *string `json:"AccountId,omitempty"`
	// The BGP (Border Gateway Protocol) ASN (Autonomous System Number) on the customer's side of the DirectLink interface.
	BgpAsn *int32 `json:"BgpAsn,omitempty"`
	// The BGP authentication key.
	BgpKey *string `json:"BgpKey,omitempty"`
	// The IP address on the customer's side of the DirectLink interface.
	ClientPrivateIp *string `json:"ClientPrivateIp,omitempty"`
	// The ID of the DirectLink.
	DirectLinkId *string `json:"DirectLinkId,omitempty"`
	// The ID of the DirectLink interface.
	DirectLinkInterfaceId *string `json:"DirectLinkInterfaceId,omitempty"`
	// The name of the DirectLink interface.
	DirectLinkInterfaceName *string `json:"DirectLinkInterfaceName,omitempty"`
	// The type of the DirectLink interface (always `private`).
	InterfaceType *string `json:"InterfaceType,omitempty"`
	// The datacenter where the DirectLink interface is located.
	Location *string `json:"Location,omitempty"`
	// The IP address on 3DS OUTSCALE's side of the DirectLink interface.
	OutscalePrivateIp *string `json:"OutscalePrivateIp,omitempty"`
	// The state of the DirectLink interface (`pending` \\| `available` \\| `deleting` \\| `deleted` \\| `confirming` \\| `rejected` \\| `expired`).
	State *string `json:"State,omitempty"`
	// The ID of the target virtual gateway.
	VirtualGatewayId *string `json:"VirtualGatewayId,omitempty"`
	// The VLAN number associated with the DirectLink interface.
	Vlan *int32 `json:"Vlan,omitempty"`
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *DirectLinkInterfaces) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DirectLinkInterfaces) GetAccountIdOk() (string, bool) {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret, false
	}
	return *o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *DirectLinkInterfaces) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *DirectLinkInterfaces) SetAccountId(v string) {
	o.AccountId = &v
}

// GetBgpAsn returns the BgpAsn field value if set, zero value otherwise.
func (o *DirectLinkInterfaces) GetBgpAsn() int32 {
	if o == nil || o.BgpAsn == nil {
		var ret int32
		return ret
	}
	return *o.BgpAsn
}

// GetBgpAsnOk returns a tuple with the BgpAsn field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DirectLinkInterfaces) GetBgpAsnOk() (int32, bool) {
	if o == nil || o.BgpAsn == nil {
		var ret int32
		return ret, false
	}
	return *o.BgpAsn, true
}

// HasBgpAsn returns a boolean if a field has been set.
func (o *DirectLinkInterfaces) HasBgpAsn() bool {
	if o != nil && o.BgpAsn != nil {
		return true
	}

	return false
}

// SetBgpAsn gets a reference to the given int32 and assigns it to the BgpAsn field.
func (o *DirectLinkInterfaces) SetBgpAsn(v int32) {
	o.BgpAsn = &v
}

// GetBgpKey returns the BgpKey field value if set, zero value otherwise.
func (o *DirectLinkInterfaces) GetBgpKey() string {
	if o == nil || o.BgpKey == nil {
		var ret string
		return ret
	}
	return *o.BgpKey
}

// GetBgpKeyOk returns a tuple with the BgpKey field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DirectLinkInterfaces) GetBgpKeyOk() (string, bool) {
	if o == nil || o.BgpKey == nil {
		var ret string
		return ret, false
	}
	return *o.BgpKey, true
}

// HasBgpKey returns a boolean if a field has been set.
func (o *DirectLinkInterfaces) HasBgpKey() bool {
	if o != nil && o.BgpKey != nil {
		return true
	}

	return false
}

// SetBgpKey gets a reference to the given string and assigns it to the BgpKey field.
func (o *DirectLinkInterfaces) SetBgpKey(v string) {
	o.BgpKey = &v
}

// GetClientPrivateIp returns the ClientPrivateIp field value if set, zero value otherwise.
func (o *DirectLinkInterfaces) GetClientPrivateIp() string {
	if o == nil || o.ClientPrivateIp == nil {
		var ret string
		return ret
	}
	return *o.ClientPrivateIp
}

// GetClientPrivateIpOk returns a tuple with the ClientPrivateIp field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DirectLinkInterfaces) GetClientPrivateIpOk() (string, bool) {
	if o == nil || o.ClientPrivateIp == nil {
		var ret string
		return ret, false
	}
	return *o.ClientPrivateIp, true
}

// HasClientPrivateIp returns a boolean if a field has been set.
func (o *DirectLinkInterfaces) HasClientPrivateIp() bool {
	if o != nil && o.ClientPrivateIp != nil {
		return true
	}

	return false
}

// SetClientPrivateIp gets a reference to the given string and assigns it to the ClientPrivateIp field.
func (o *DirectLinkInterfaces) SetClientPrivateIp(v string) {
	o.ClientPrivateIp = &v
}

// GetDirectLinkId returns the DirectLinkId field value if set, zero value otherwise.
func (o *DirectLinkInterfaces) GetDirectLinkId() string {
	if o == nil || o.DirectLinkId == nil {
		var ret string
		return ret
	}
	return *o.DirectLinkId
}

// GetDirectLinkIdOk returns a tuple with the DirectLinkId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DirectLinkInterfaces) GetDirectLinkIdOk() (string, bool) {
	if o == nil || o.DirectLinkId == nil {
		var ret string
		return ret, false
	}
	return *o.DirectLinkId, true
}

// HasDirectLinkId returns a boolean if a field has been set.
func (o *DirectLinkInterfaces) HasDirectLinkId() bool {
	if o != nil && o.DirectLinkId != nil {
		return true
	}

	return false
}

// SetDirectLinkId gets a reference to the given string and assigns it to the DirectLinkId field.
func (o *DirectLinkInterfaces) SetDirectLinkId(v string) {
	o.DirectLinkId = &v
}

// GetDirectLinkInterfaceId returns the DirectLinkInterfaceId field value if set, zero value otherwise.
func (o *DirectLinkInterfaces) GetDirectLinkInterfaceId() string {
	if o == nil || o.DirectLinkInterfaceId == nil {
		var ret string
		return ret
	}
	return *o.DirectLinkInterfaceId
}

// GetDirectLinkInterfaceIdOk returns a tuple with the DirectLinkInterfaceId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DirectLinkInterfaces) GetDirectLinkInterfaceIdOk() (string, bool) {
	if o == nil || o.DirectLinkInterfaceId == nil {
		var ret string
		return ret, false
	}
	return *o.DirectLinkInterfaceId, true
}

// HasDirectLinkInterfaceId returns a boolean if a field has been set.
func (o *DirectLinkInterfaces) HasDirectLinkInterfaceId() bool {
	if o != nil && o.DirectLinkInterfaceId != nil {
		return true
	}

	return false
}

// SetDirectLinkInterfaceId gets a reference to the given string and assigns it to the DirectLinkInterfaceId field.
func (o *DirectLinkInterfaces) SetDirectLinkInterfaceId(v string) {
	o.DirectLinkInterfaceId = &v
}

// GetDirectLinkInterfaceName returns the DirectLinkInterfaceName field value if set, zero value otherwise.
func (o *DirectLinkInterfaces) GetDirectLinkInterfaceName() string {
	if o == nil || o.DirectLinkInterfaceName == nil {
		var ret string
		return ret
	}
	return *o.DirectLinkInterfaceName
}

// GetDirectLinkInterfaceNameOk returns a tuple with the DirectLinkInterfaceName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DirectLinkInterfaces) GetDirectLinkInterfaceNameOk() (string, bool) {
	if o == nil || o.DirectLinkInterfaceName == nil {
		var ret string
		return ret, false
	}
	return *o.DirectLinkInterfaceName, true
}

// HasDirectLinkInterfaceName returns a boolean if a field has been set.
func (o *DirectLinkInterfaces) HasDirectLinkInterfaceName() bool {
	if o != nil && o.DirectLinkInterfaceName != nil {
		return true
	}

	return false
}

// SetDirectLinkInterfaceName gets a reference to the given string and assigns it to the DirectLinkInterfaceName field.
func (o *DirectLinkInterfaces) SetDirectLinkInterfaceName(v string) {
	o.DirectLinkInterfaceName = &v
}

// GetInterfaceType returns the InterfaceType field value if set, zero value otherwise.
func (o *DirectLinkInterfaces) GetInterfaceType() string {
	if o == nil || o.InterfaceType == nil {
		var ret string
		return ret
	}
	return *o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DirectLinkInterfaces) GetInterfaceTypeOk() (string, bool) {
	if o == nil || o.InterfaceType == nil {
		var ret string
		return ret, false
	}
	return *o.InterfaceType, true
}

// HasInterfaceType returns a boolean if a field has been set.
func (o *DirectLinkInterfaces) HasInterfaceType() bool {
	if o != nil && o.InterfaceType != nil {
		return true
	}

	return false
}

// SetInterfaceType gets a reference to the given string and assigns it to the InterfaceType field.
func (o *DirectLinkInterfaces) SetInterfaceType(v string) {
	o.InterfaceType = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *DirectLinkInterfaces) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DirectLinkInterfaces) GetLocationOk() (string, bool) {
	if o == nil || o.Location == nil {
		var ret string
		return ret, false
	}
	return *o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *DirectLinkInterfaces) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *DirectLinkInterfaces) SetLocation(v string) {
	o.Location = &v
}

// GetOutscalePrivateIp returns the OutscalePrivateIp field value if set, zero value otherwise.
func (o *DirectLinkInterfaces) GetOutscalePrivateIp() string {
	if o == nil || o.OutscalePrivateIp == nil {
		var ret string
		return ret
	}
	return *o.OutscalePrivateIp
}

// GetOutscalePrivateIpOk returns a tuple with the OutscalePrivateIp field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DirectLinkInterfaces) GetOutscalePrivateIpOk() (string, bool) {
	if o == nil || o.OutscalePrivateIp == nil {
		var ret string
		return ret, false
	}
	return *o.OutscalePrivateIp, true
}

// HasOutscalePrivateIp returns a boolean if a field has been set.
func (o *DirectLinkInterfaces) HasOutscalePrivateIp() bool {
	if o != nil && o.OutscalePrivateIp != nil {
		return true
	}

	return false
}

// SetOutscalePrivateIp gets a reference to the given string and assigns it to the OutscalePrivateIp field.
func (o *DirectLinkInterfaces) SetOutscalePrivateIp(v string) {
	o.OutscalePrivateIp = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DirectLinkInterfaces) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DirectLinkInterfaces) GetStateOk() (string, bool) {
	if o == nil || o.State == nil {
		var ret string
		return ret, false
	}
	return *o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DirectLinkInterfaces) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DirectLinkInterfaces) SetState(v string) {
	o.State = &v
}

// GetVirtualGatewayId returns the VirtualGatewayId field value if set, zero value otherwise.
func (o *DirectLinkInterfaces) GetVirtualGatewayId() string {
	if o == nil || o.VirtualGatewayId == nil {
		var ret string
		return ret
	}
	return *o.VirtualGatewayId
}

// GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DirectLinkInterfaces) GetVirtualGatewayIdOk() (string, bool) {
	if o == nil || o.VirtualGatewayId == nil {
		var ret string
		return ret, false
	}
	return *o.VirtualGatewayId, true
}

// HasVirtualGatewayId returns a boolean if a field has been set.
func (o *DirectLinkInterfaces) HasVirtualGatewayId() bool {
	if o != nil && o.VirtualGatewayId != nil {
		return true
	}

	return false
}

// SetVirtualGatewayId gets a reference to the given string and assigns it to the VirtualGatewayId field.
func (o *DirectLinkInterfaces) SetVirtualGatewayId(v string) {
	o.VirtualGatewayId = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *DirectLinkInterfaces) GetVlan() int32 {
	if o == nil || o.Vlan == nil {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DirectLinkInterfaces) GetVlanOk() (int32, bool) {
	if o == nil || o.Vlan == nil {
		var ret int32
		return ret, false
	}
	return *o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *DirectLinkInterfaces) HasVlan() bool {
	if o != nil && o.Vlan != nil {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *DirectLinkInterfaces) SetVlan(v int32) {
	o.Vlan = &v
}

type NullableDirectLinkInterfaces struct {
	Value        DirectLinkInterfaces
	ExplicitNull bool
}

func (v NullableDirectLinkInterfaces) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableDirectLinkInterfaces) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
