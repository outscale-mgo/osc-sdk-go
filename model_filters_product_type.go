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

// FiltersProductType One or more filters.
type FiltersProductType struct {
	// The IDs of the product types.
	ProductTypeIds *[]string `json:"ProductTypeIds,omitempty"`
}

// GetProductTypeIds returns the ProductTypeIds field value if set, zero value otherwise.
func (o *FiltersProductType) GetProductTypeIds() []string {
	if o == nil || o.ProductTypeIds == nil {
		var ret []string
		return ret
	}
	return *o.ProductTypeIds
}

// GetProductTypeIdsOk returns a tuple with the ProductTypeIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersProductType) GetProductTypeIdsOk() ([]string, bool) {
	if o == nil || o.ProductTypeIds == nil {
		var ret []string
		return ret, false
	}
	return *o.ProductTypeIds, true
}

// HasProductTypeIds returns a boolean if a field has been set.
func (o *FiltersProductType) HasProductTypeIds() bool {
	if o != nil && o.ProductTypeIds != nil {
		return true
	}

	return false
}

// SetProductTypeIds gets a reference to the given []string and assigns it to the ProductTypeIds field.
func (o *FiltersProductType) SetProductTypeIds(v []string) {
	o.ProductTypeIds = &v
}

type NullableFiltersProductType struct {
	Value        FiltersProductType
	ExplicitNull bool
}

func (v NullableFiltersProductType) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFiltersProductType) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
