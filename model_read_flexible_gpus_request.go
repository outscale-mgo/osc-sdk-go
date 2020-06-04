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

// ReadFlexibleGpusRequest struct for ReadFlexibleGpusRequest
type ReadFlexibleGpusRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun  *bool               `json:"DryRun,omitempty"`
	Filters *FiltersFlexibleGpu `json:"Filters,omitempty"`
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadFlexibleGpusRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadFlexibleGpusRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadFlexibleGpusRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadFlexibleGpusRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ReadFlexibleGpusRequest) GetFilters() FiltersFlexibleGpu {
	if o == nil || o.Filters == nil {
		var ret FiltersFlexibleGpu
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadFlexibleGpusRequest) GetFiltersOk() (FiltersFlexibleGpu, bool) {
	if o == nil || o.Filters == nil {
		var ret FiltersFlexibleGpu
		return ret, false
	}
	return *o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ReadFlexibleGpusRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given FiltersFlexibleGpu and assigns it to the Filters field.
func (o *ReadFlexibleGpusRequest) SetFilters(v FiltersFlexibleGpu) {
	o.Filters = &v
}

type NullableReadFlexibleGpusRequest struct {
	Value        ReadFlexibleGpusRequest
	ExplicitNull bool
}

func (v NullableReadFlexibleGpusRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableReadFlexibleGpusRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
