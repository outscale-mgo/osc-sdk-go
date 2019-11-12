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

// ReadVmTypesRequest struct for ReadVmTypesRequest
type ReadVmTypesRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	Filters *FiltersVmType `json:"Filters,omitempty"`
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadVmTypesRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmTypesRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadVmTypesRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadVmTypesRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ReadVmTypesRequest) GetFilters() FiltersVmType {
	if o == nil || o.Filters == nil {
		var ret FiltersVmType
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadVmTypesRequest) GetFiltersOk() (FiltersVmType, bool) {
	if o == nil || o.Filters == nil {
		var ret FiltersVmType
		return ret, false
	}
	return *o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ReadVmTypesRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given FiltersVmType and assigns it to the Filters field.
func (o *ReadVmTypesRequest) SetFilters(v FiltersVmType) {
	o.Filters = &v
}

type NullableReadVmTypesRequest struct {
	Value ReadVmTypesRequest
	ExplicitNull bool
}

func (v NullableReadVmTypesRequest) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableReadVmTypesRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

