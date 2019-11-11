/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"bytes"
	"encoding/json"
)

// ReadNetsResponse struct for ReadNetsResponse
type ReadNetsResponse struct {
	// Information about the described Nets.
	Nets *[]Net `json:"Nets,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetNets returns the Nets field value if set, zero value otherwise.
func (o *ReadNetsResponse) GetNets() []Net {
	if o == nil || o.Nets == nil {
		var ret []Net
		return ret
	}
	return *o.Nets
}

// GetNetsOk returns a tuple with the Nets field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadNetsResponse) GetNetsOk() ([]Net, bool) {
	if o == nil || o.Nets == nil {
		var ret []Net
		return ret, false
	}
	return *o.Nets, true
}

// HasNets returns a boolean if a field has been set.
func (o *ReadNetsResponse) HasNets() bool {
	if o != nil && o.Nets != nil {
		return true
	}

	return false
}

// SetNets gets a reference to the given []Net and assigns it to the Nets field.
func (o *ReadNetsResponse) SetNets(v []Net) {
	o.Nets = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadNetsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadNetsResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadNetsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadNetsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableReadNetsResponse struct {
	Value ReadNetsResponse
	ExplicitNull bool
}

func (v NullableReadNetsResponse) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableReadNetsResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
