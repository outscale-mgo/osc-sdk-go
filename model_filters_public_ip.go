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

// FiltersPublicIp One or more filters.
type FiltersPublicIp struct {
	// The IDs representing the associations of EIPs with VMs or NICs.
	LinkPublicIpIds *[]string `json:"LinkPublicIpIds,omitempty"`
	// The account IDs of the owners of the NICs.
	NicAccountIds *[]string `json:"NicAccountIds,omitempty"`
	// The IDs of the NICs.
	NicIds *[]string `json:"NicIds,omitempty"`
	// Whether the EIPs are for use in the public Cloud or in a Net.
	Placements *[]string `json:"Placements,omitempty"`
	// The private IP addresses associated with the EIPs.
	PrivateIps *[]string `json:"PrivateIps,omitempty"`
	// The IDs of the External IP addresses (EIPs).
	PublicIpIds *[]string `json:"PublicIpIds,omitempty"`
	// The External IP addresses (EIPs).
	PublicIps *[]string `json:"PublicIps,omitempty"`
	// The keys of the tags associated with the EIPs.
	TagKeys *[]string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the EIPs.
	TagValues *[]string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the EIPs, in the following format: \"Filters\":{\"Tags\":[\"TAGKEY=TAGVALUE\"]}.
	Tags *[]string `json:"Tags,omitempty"`
	// The IDs of the VMs.
	VmIds *[]string `json:"VmIds,omitempty"`
}

// GetLinkPublicIpIds returns the LinkPublicIpIds field value if set, zero value otherwise.
func (o *FiltersPublicIp) GetLinkPublicIpIds() []string {
	if o == nil || o.LinkPublicIpIds == nil {
		var ret []string
		return ret
	}
	return *o.LinkPublicIpIds
}

// GetLinkPublicIpIdsOk returns a tuple with the LinkPublicIpIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersPublicIp) GetLinkPublicIpIdsOk() ([]string, bool) {
	if o == nil || o.LinkPublicIpIds == nil {
		var ret []string
		return ret, false
	}
	return *o.LinkPublicIpIds, true
}

// HasLinkPublicIpIds returns a boolean if a field has been set.
func (o *FiltersPublicIp) HasLinkPublicIpIds() bool {
	if o != nil && o.LinkPublicIpIds != nil {
		return true
	}

	return false
}

// SetLinkPublicIpIds gets a reference to the given []string and assigns it to the LinkPublicIpIds field.
func (o *FiltersPublicIp) SetLinkPublicIpIds(v []string) {
	o.LinkPublicIpIds = &v
}

// GetNicAccountIds returns the NicAccountIds field value if set, zero value otherwise.
func (o *FiltersPublicIp) GetNicAccountIds() []string {
	if o == nil || o.NicAccountIds == nil {
		var ret []string
		return ret
	}
	return *o.NicAccountIds
}

// GetNicAccountIdsOk returns a tuple with the NicAccountIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersPublicIp) GetNicAccountIdsOk() ([]string, bool) {
	if o == nil || o.NicAccountIds == nil {
		var ret []string
		return ret, false
	}
	return *o.NicAccountIds, true
}

// HasNicAccountIds returns a boolean if a field has been set.
func (o *FiltersPublicIp) HasNicAccountIds() bool {
	if o != nil && o.NicAccountIds != nil {
		return true
	}

	return false
}

// SetNicAccountIds gets a reference to the given []string and assigns it to the NicAccountIds field.
func (o *FiltersPublicIp) SetNicAccountIds(v []string) {
	o.NicAccountIds = &v
}

// GetNicIds returns the NicIds field value if set, zero value otherwise.
func (o *FiltersPublicIp) GetNicIds() []string {
	if o == nil || o.NicIds == nil {
		var ret []string
		return ret
	}
	return *o.NicIds
}

// GetNicIdsOk returns a tuple with the NicIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersPublicIp) GetNicIdsOk() ([]string, bool) {
	if o == nil || o.NicIds == nil {
		var ret []string
		return ret, false
	}
	return *o.NicIds, true
}

// HasNicIds returns a boolean if a field has been set.
func (o *FiltersPublicIp) HasNicIds() bool {
	if o != nil && o.NicIds != nil {
		return true
	}

	return false
}

// SetNicIds gets a reference to the given []string and assigns it to the NicIds field.
func (o *FiltersPublicIp) SetNicIds(v []string) {
	o.NicIds = &v
}

// GetPlacements returns the Placements field value if set, zero value otherwise.
func (o *FiltersPublicIp) GetPlacements() []string {
	if o == nil || o.Placements == nil {
		var ret []string
		return ret
	}
	return *o.Placements
}

// GetPlacementsOk returns a tuple with the Placements field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersPublicIp) GetPlacementsOk() ([]string, bool) {
	if o == nil || o.Placements == nil {
		var ret []string
		return ret, false
	}
	return *o.Placements, true
}

// HasPlacements returns a boolean if a field has been set.
func (o *FiltersPublicIp) HasPlacements() bool {
	if o != nil && o.Placements != nil {
		return true
	}

	return false
}

// SetPlacements gets a reference to the given []string and assigns it to the Placements field.
func (o *FiltersPublicIp) SetPlacements(v []string) {
	o.Placements = &v
}

// GetPrivateIps returns the PrivateIps field value if set, zero value otherwise.
func (o *FiltersPublicIp) GetPrivateIps() []string {
	if o == nil || o.PrivateIps == nil {
		var ret []string
		return ret
	}
	return *o.PrivateIps
}

// GetPrivateIpsOk returns a tuple with the PrivateIps field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersPublicIp) GetPrivateIpsOk() ([]string, bool) {
	if o == nil || o.PrivateIps == nil {
		var ret []string
		return ret, false
	}
	return *o.PrivateIps, true
}

// HasPrivateIps returns a boolean if a field has been set.
func (o *FiltersPublicIp) HasPrivateIps() bool {
	if o != nil && o.PrivateIps != nil {
		return true
	}

	return false
}

// SetPrivateIps gets a reference to the given []string and assigns it to the PrivateIps field.
func (o *FiltersPublicIp) SetPrivateIps(v []string) {
	o.PrivateIps = &v
}

// GetPublicIpIds returns the PublicIpIds field value if set, zero value otherwise.
func (o *FiltersPublicIp) GetPublicIpIds() []string {
	if o == nil || o.PublicIpIds == nil {
		var ret []string
		return ret
	}
	return *o.PublicIpIds
}

// GetPublicIpIdsOk returns a tuple with the PublicIpIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersPublicIp) GetPublicIpIdsOk() ([]string, bool) {
	if o == nil || o.PublicIpIds == nil {
		var ret []string
		return ret, false
	}
	return *o.PublicIpIds, true
}

// HasPublicIpIds returns a boolean if a field has been set.
func (o *FiltersPublicIp) HasPublicIpIds() bool {
	if o != nil && o.PublicIpIds != nil {
		return true
	}

	return false
}

// SetPublicIpIds gets a reference to the given []string and assigns it to the PublicIpIds field.
func (o *FiltersPublicIp) SetPublicIpIds(v []string) {
	o.PublicIpIds = &v
}

// GetPublicIps returns the PublicIps field value if set, zero value otherwise.
func (o *FiltersPublicIp) GetPublicIps() []string {
	if o == nil || o.PublicIps == nil {
		var ret []string
		return ret
	}
	return *o.PublicIps
}

// GetPublicIpsOk returns a tuple with the PublicIps field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersPublicIp) GetPublicIpsOk() ([]string, bool) {
	if o == nil || o.PublicIps == nil {
		var ret []string
		return ret, false
	}
	return *o.PublicIps, true
}

// HasPublicIps returns a boolean if a field has been set.
func (o *FiltersPublicIp) HasPublicIps() bool {
	if o != nil && o.PublicIps != nil {
		return true
	}

	return false
}

// SetPublicIps gets a reference to the given []string and assigns it to the PublicIps field.
func (o *FiltersPublicIp) SetPublicIps(v []string) {
	o.PublicIps = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *FiltersPublicIp) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return *o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersPublicIp) GetTagKeysOk() ([]string, bool) {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret, false
	}
	return *o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *FiltersPublicIp) HasTagKeys() bool {
	if o != nil && o.TagKeys != nil {
		return true
	}

	return false
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *FiltersPublicIp) SetTagKeys(v []string) {
	o.TagKeys = &v
}

// GetTagValues returns the TagValues field value if set, zero value otherwise.
func (o *FiltersPublicIp) GetTagValues() []string {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret
	}
	return *o.TagValues
}

// GetTagValuesOk returns a tuple with the TagValues field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersPublicIp) GetTagValuesOk() ([]string, bool) {
	if o == nil || o.TagValues == nil {
		var ret []string
		return ret, false
	}
	return *o.TagValues, true
}

// HasTagValues returns a boolean if a field has been set.
func (o *FiltersPublicIp) HasTagValues() bool {
	if o != nil && o.TagValues != nil {
		return true
	}

	return false
}

// SetTagValues gets a reference to the given []string and assigns it to the TagValues field.
func (o *FiltersPublicIp) SetTagValues(v []string) {
	o.TagValues = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FiltersPublicIp) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersPublicIp) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FiltersPublicIp) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FiltersPublicIp) SetTags(v []string) {
	o.Tags = &v
}

// GetVmIds returns the VmIds field value if set, zero value otherwise.
func (o *FiltersPublicIp) GetVmIds() []string {
	if o == nil || o.VmIds == nil {
		var ret []string
		return ret
	}
	return *o.VmIds
}

// GetVmIdsOk returns a tuple with the VmIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersPublicIp) GetVmIdsOk() ([]string, bool) {
	if o == nil || o.VmIds == nil {
		var ret []string
		return ret, false
	}
	return *o.VmIds, true
}

// HasVmIds returns a boolean if a field has been set.
func (o *FiltersPublicIp) HasVmIds() bool {
	if o != nil && o.VmIds != nil {
		return true
	}

	return false
}

// SetVmIds gets a reference to the given []string and assigns it to the VmIds field.
func (o *FiltersPublicIp) SetVmIds(v []string) {
	o.VmIds = &v
}

type NullableFiltersPublicIp struct {
	Value        FiltersPublicIp
	ExplicitNull bool
}

func (v NullableFiltersPublicIp) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFiltersPublicIp) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
