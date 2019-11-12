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

// UpdateFlexibleGpuResponse struct for UpdateFlexibleGpuResponse
type UpdateFlexibleGpuResponse struct {
	FlexibleGpu *FlexibleGpu `json:"FlexibleGpu,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetFlexibleGpu returns the FlexibleGpu field value if set, zero value otherwise.
func (o *UpdateFlexibleGpuResponse) GetFlexibleGpu() FlexibleGpu {
	if o == nil || o.FlexibleGpu == nil {
		var ret FlexibleGpu
		return ret
	}
	return *o.FlexibleGpu
}

// GetFlexibleGpuOk returns a tuple with the FlexibleGpu field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFlexibleGpuResponse) GetFlexibleGpuOk() (FlexibleGpu, bool) {
	if o == nil || o.FlexibleGpu == nil {
		var ret FlexibleGpu
		return ret, false
	}
	return *o.FlexibleGpu, true
}

// HasFlexibleGpu returns a boolean if a field has been set.
func (o *UpdateFlexibleGpuResponse) HasFlexibleGpu() bool {
	if o != nil && o.FlexibleGpu != nil {
		return true
	}

	return false
}

// SetFlexibleGpu gets a reference to the given FlexibleGpu and assigns it to the FlexibleGpu field.
func (o *UpdateFlexibleGpuResponse) SetFlexibleGpu(v FlexibleGpu) {
	o.FlexibleGpu = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateFlexibleGpuResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFlexibleGpuResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateFlexibleGpuResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateFlexibleGpuResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableUpdateFlexibleGpuResponse struct {
	Value UpdateFlexibleGpuResponse
	ExplicitNull bool
}

func (v NullableUpdateFlexibleGpuResponse) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableUpdateFlexibleGpuResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

