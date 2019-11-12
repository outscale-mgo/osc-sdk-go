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

// ReadDhcpOptionsResponse struct for ReadDhcpOptionsResponse
type ReadDhcpOptionsResponse struct {
	// Information about one or more DHCP options sets.
	DhcpOptionsSets *[]DhcpOptionsSet `json:"DhcpOptionsSets,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetDhcpOptionsSets returns the DhcpOptionsSets field value if set, zero value otherwise.
func (o *ReadDhcpOptionsResponse) GetDhcpOptionsSets() []DhcpOptionsSet {
	if o == nil || o.DhcpOptionsSets == nil {
		var ret []DhcpOptionsSet
		return ret
	}
	return *o.DhcpOptionsSets
}

// GetDhcpOptionsSetsOk returns a tuple with the DhcpOptionsSets field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadDhcpOptionsResponse) GetDhcpOptionsSetsOk() ([]DhcpOptionsSet, bool) {
	if o == nil || o.DhcpOptionsSets == nil {
		var ret []DhcpOptionsSet
		return ret, false
	}
	return *o.DhcpOptionsSets, true
}

// HasDhcpOptionsSets returns a boolean if a field has been set.
func (o *ReadDhcpOptionsResponse) HasDhcpOptionsSets() bool {
	if o != nil && o.DhcpOptionsSets != nil {
		return true
	}

	return false
}

// SetDhcpOptionsSets gets a reference to the given []DhcpOptionsSet and assigns it to the DhcpOptionsSets field.
func (o *ReadDhcpOptionsResponse) SetDhcpOptionsSets(v []DhcpOptionsSet) {
	o.DhcpOptionsSets = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadDhcpOptionsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadDhcpOptionsResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadDhcpOptionsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadDhcpOptionsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableReadDhcpOptionsResponse struct {
	Value ReadDhcpOptionsResponse
	ExplicitNull bool
}

func (v NullableReadDhcpOptionsResponse) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableReadDhcpOptionsResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

