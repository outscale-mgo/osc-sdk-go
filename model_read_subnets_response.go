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

// ReadSubnetsResponse struct for ReadSubnetsResponse
type ReadSubnetsResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// Information about one or more Subnets.
	Subnets *[]Subnet `json:"Subnets,omitempty"`
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadSubnetsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadSubnetsResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadSubnetsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadSubnetsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *ReadSubnetsResponse) GetSubnets() []Subnet {
	if o == nil || o.Subnets == nil {
		var ret []Subnet
		return ret
	}
	return *o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadSubnetsResponse) GetSubnetsOk() ([]Subnet, bool) {
	if o == nil || o.Subnets == nil {
		var ret []Subnet
		return ret, false
	}
	return *o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *ReadSubnetsResponse) HasSubnets() bool {
	if o != nil && o.Subnets != nil {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []Subnet and assigns it to the Subnets field.
func (o *ReadSubnetsResponse) SetSubnets(v []Subnet) {
	o.Subnets = &v
}

type NullableReadSubnetsResponse struct {
	Value ReadSubnetsResponse
	ExplicitNull bool
}

func (v NullableReadSubnetsResponse) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableReadSubnetsResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

