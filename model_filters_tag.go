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

// FiltersTag One or more filters.
type FiltersTag struct {
	// The keys of the tags that are assigned to the resources. You can use this filter alongside the `Values` filter. In that case, you filter the resources corresponding to each tag, regardless of the other filter.
	Keys *[]string `json:"Keys,omitempty"`
	// The IDs of the resources with which the tags are associated.
	ResourceIds *[]string `json:"ResourceIds,omitempty"`
	// The resource type (`instance` \\| `image` \\| `volume` \\| `snapshot` \\| `public-ip` \\| `security-group` \\| `route-table` \\| `network-interface` \\| `vpc` \\| `subnet` \\| `network-link` \\| `vpc-endpoint` \\| `nat-gateway` \\| `internet-gateway` \\| `customer-gateway` \\| `vpn-gateway` \\| `vpn-connection` \\| `dhcp-options` \\| `task`).
	ResourceTypes *[]string `json:"ResourceTypes,omitempty"`
	// The values of the tags that are assigned to the resources. You can use this filter alongside the `TagKeys` filter. In that case, you filter the resources corresponding to each tag, regardless of the other filter.
	Values *[]string `json:"Values,omitempty"`
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *FiltersTag) GetKeys() []string {
	if o == nil || o.Keys == nil {
		var ret []string
		return ret
	}
	return *o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersTag) GetKeysOk() ([]string, bool) {
	if o == nil || o.Keys == nil {
		var ret []string
		return ret, false
	}
	return *o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *FiltersTag) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []string and assigns it to the Keys field.
func (o *FiltersTag) SetKeys(v []string) {
	o.Keys = &v
}

// GetResourceIds returns the ResourceIds field value if set, zero value otherwise.
func (o *FiltersTag) GetResourceIds() []string {
	if o == nil || o.ResourceIds == nil {
		var ret []string
		return ret
	}
	return *o.ResourceIds
}

// GetResourceIdsOk returns a tuple with the ResourceIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersTag) GetResourceIdsOk() ([]string, bool) {
	if o == nil || o.ResourceIds == nil {
		var ret []string
		return ret, false
	}
	return *o.ResourceIds, true
}

// HasResourceIds returns a boolean if a field has been set.
func (o *FiltersTag) HasResourceIds() bool {
	if o != nil && o.ResourceIds != nil {
		return true
	}

	return false
}

// SetResourceIds gets a reference to the given []string and assigns it to the ResourceIds field.
func (o *FiltersTag) SetResourceIds(v []string) {
	o.ResourceIds = &v
}

// GetResourceTypes returns the ResourceTypes field value if set, zero value otherwise.
func (o *FiltersTag) GetResourceTypes() []string {
	if o == nil || o.ResourceTypes == nil {
		var ret []string
		return ret
	}
	return *o.ResourceTypes
}

// GetResourceTypesOk returns a tuple with the ResourceTypes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersTag) GetResourceTypesOk() ([]string, bool) {
	if o == nil || o.ResourceTypes == nil {
		var ret []string
		return ret, false
	}
	return *o.ResourceTypes, true
}

// HasResourceTypes returns a boolean if a field has been set.
func (o *FiltersTag) HasResourceTypes() bool {
	if o != nil && o.ResourceTypes != nil {
		return true
	}

	return false
}

// SetResourceTypes gets a reference to the given []string and assigns it to the ResourceTypes field.
func (o *FiltersTag) SetResourceTypes(v []string) {
	o.ResourceTypes = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *FiltersTag) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersTag) GetValuesOk() ([]string, bool) {
	if o == nil || o.Values == nil {
		var ret []string
		return ret, false
	}
	return *o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *FiltersTag) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *FiltersTag) SetValues(v []string) {
	o.Values = &v
}

type NullableFiltersTag struct {
	Value        FiltersTag
	ExplicitNull bool
}

func (v NullableFiltersTag) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFiltersTag) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
