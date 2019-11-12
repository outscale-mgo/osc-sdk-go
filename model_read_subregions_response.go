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

// ReadSubregionsResponse struct for ReadSubregionsResponse
type ReadSubregionsResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more Subregions.
	Subregions *[]Subregion `json:"Subregions,omitempty"`
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadSubregionsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadSubregionsResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadSubregionsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadSubregionsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetSubregions returns the Subregions field value if set, zero value otherwise.
func (o *ReadSubregionsResponse) GetSubregions() []Subregion {
	if o == nil || o.Subregions == nil {
		var ret []Subregion
		return ret
	}
	return *o.Subregions
}

// GetSubregionsOk returns a tuple with the Subregions field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadSubregionsResponse) GetSubregionsOk() ([]Subregion, bool) {
	if o == nil || o.Subregions == nil {
		var ret []Subregion
		return ret, false
	}
	return *o.Subregions, true
}

// HasSubregions returns a boolean if a field has been set.
func (o *ReadSubregionsResponse) HasSubregions() bool {
	if o != nil && o.Subregions != nil {
		return true
	}

	return false
}

// SetSubregions gets a reference to the given []Subregion and assigns it to the Subregions field.
func (o *ReadSubregionsResponse) SetSubregions(v []Subregion) {
	o.Subregions = &v
}

type NullableReadSubregionsResponse struct {
	Value ReadSubregionsResponse
	ExplicitNull bool
}

func (v NullableReadSubregionsResponse) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableReadSubregionsResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

