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

// FiltersService One or more filters.
type FiltersService struct {
	// The IDs of the services.
	ServiceIds *[]string `json:"ServiceIds,omitempty"`
	// The names of the prefix lists, which identify the 3DS OUTSCALE services they are associated with.
	ServiceNames *[]string `json:"ServiceNames,omitempty"`
}

// GetServiceIds returns the ServiceIds field value if set, zero value otherwise.
func (o *FiltersService) GetServiceIds() []string {
	if o == nil || o.ServiceIds == nil {
		var ret []string
		return ret
	}
	return *o.ServiceIds
}

// GetServiceIdsOk returns a tuple with the ServiceIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersService) GetServiceIdsOk() ([]string, bool) {
	if o == nil || o.ServiceIds == nil {
		var ret []string
		return ret, false
	}
	return *o.ServiceIds, true
}

// HasServiceIds returns a boolean if a field has been set.
func (o *FiltersService) HasServiceIds() bool {
	if o != nil && o.ServiceIds != nil {
		return true
	}

	return false
}

// SetServiceIds gets a reference to the given []string and assigns it to the ServiceIds field.
func (o *FiltersService) SetServiceIds(v []string) {
	o.ServiceIds = &v
}

// GetServiceNames returns the ServiceNames field value if set, zero value otherwise.
func (o *FiltersService) GetServiceNames() []string {
	if o == nil || o.ServiceNames == nil {
		var ret []string
		return ret
	}
	return *o.ServiceNames
}

// GetServiceNamesOk returns a tuple with the ServiceNames field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FiltersService) GetServiceNamesOk() ([]string, bool) {
	if o == nil || o.ServiceNames == nil {
		var ret []string
		return ret, false
	}
	return *o.ServiceNames, true
}

// HasServiceNames returns a boolean if a field has been set.
func (o *FiltersService) HasServiceNames() bool {
	if o != nil && o.ServiceNames != nil {
		return true
	}

	return false
}

// SetServiceNames gets a reference to the given []string and assigns it to the ServiceNames field.
func (o *FiltersService) SetServiceNames(v []string) {
	o.ServiceNames = &v
}

type NullableFiltersService struct {
	Value FiltersService
	ExplicitNull bool
}

func (v NullableFiltersService) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableFiltersService) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

