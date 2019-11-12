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

// ReadLoadBalancersResponse struct for ReadLoadBalancersResponse
type ReadLoadBalancersResponse struct {
	// Information about one or more load balancers.
	LoadBalancers *[]LoadBalancer `json:"LoadBalancers,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetLoadBalancers returns the LoadBalancers field value if set, zero value otherwise.
func (o *ReadLoadBalancersResponse) GetLoadBalancers() []LoadBalancer {
	if o == nil || o.LoadBalancers == nil {
		var ret []LoadBalancer
		return ret
	}
	return *o.LoadBalancers
}

// GetLoadBalancersOk returns a tuple with the LoadBalancers field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadLoadBalancersResponse) GetLoadBalancersOk() ([]LoadBalancer, bool) {
	if o == nil || o.LoadBalancers == nil {
		var ret []LoadBalancer
		return ret, false
	}
	return *o.LoadBalancers, true
}

// HasLoadBalancers returns a boolean if a field has been set.
func (o *ReadLoadBalancersResponse) HasLoadBalancers() bool {
	if o != nil && o.LoadBalancers != nil {
		return true
	}

	return false
}

// SetLoadBalancers gets a reference to the given []LoadBalancer and assigns it to the LoadBalancers field.
func (o *ReadLoadBalancersResponse) SetLoadBalancers(v []LoadBalancer) {
	o.LoadBalancers = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadLoadBalancersResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadLoadBalancersResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadLoadBalancersResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadLoadBalancersResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableReadLoadBalancersResponse struct {
	Value ReadLoadBalancersResponse
	ExplicitNull bool
}

func (v NullableReadLoadBalancersResponse) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableReadLoadBalancersResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

