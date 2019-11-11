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

// UpdateLoadBalancerResponse struct for UpdateLoadBalancerResponse
type UpdateLoadBalancerResponse struct {
	LoadBalancer *LoadBalancer `json:"LoadBalancer,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetLoadBalancer returns the LoadBalancer field value if set, zero value otherwise.
func (o *UpdateLoadBalancerResponse) GetLoadBalancer() LoadBalancer {
	if o == nil || o.LoadBalancer == nil {
		var ret LoadBalancer
		return ret
	}
	return *o.LoadBalancer
}

// GetLoadBalancerOk returns a tuple with the LoadBalancer field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerResponse) GetLoadBalancerOk() (LoadBalancer, bool) {
	if o == nil || o.LoadBalancer == nil {
		var ret LoadBalancer
		return ret, false
	}
	return *o.LoadBalancer, true
}

// HasLoadBalancer returns a boolean if a field has been set.
func (o *UpdateLoadBalancerResponse) HasLoadBalancer() bool {
	if o != nil && o.LoadBalancer != nil {
		return true
	}

	return false
}

// SetLoadBalancer gets a reference to the given LoadBalancer and assigns it to the LoadBalancer field.
func (o *UpdateLoadBalancerResponse) SetLoadBalancer(v LoadBalancer) {
	o.LoadBalancer = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *UpdateLoadBalancerResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoadBalancerResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *UpdateLoadBalancerResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *UpdateLoadBalancerResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableUpdateLoadBalancerResponse struct {
	Value UpdateLoadBalancerResponse
	ExplicitNull bool
}

func (v NullableUpdateLoadBalancerResponse) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableUpdateLoadBalancerResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

