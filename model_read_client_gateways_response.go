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

// ReadClientGatewaysResponse struct for ReadClientGatewaysResponse
type ReadClientGatewaysResponse struct {
	// Information about one or more client gateways.
	ClientGateways  *[]ClientGateway `json:"ClientGateways,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetClientGateways returns the ClientGateways field value if set, zero value otherwise.
func (o *ReadClientGatewaysResponse) GetClientGateways() []ClientGateway {
	if o == nil || o.ClientGateways == nil {
		var ret []ClientGateway
		return ret
	}
	return *o.ClientGateways
}

// GetClientGatewaysOk returns a tuple with the ClientGateways field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadClientGatewaysResponse) GetClientGatewaysOk() ([]ClientGateway, bool) {
	if o == nil || o.ClientGateways == nil {
		var ret []ClientGateway
		return ret, false
	}
	return *o.ClientGateways, true
}

// HasClientGateways returns a boolean if a field has been set.
func (o *ReadClientGatewaysResponse) HasClientGateways() bool {
	if o != nil && o.ClientGateways != nil {
		return true
	}

	return false
}

// SetClientGateways gets a reference to the given []ClientGateway and assigns it to the ClientGateways field.
func (o *ReadClientGatewaysResponse) SetClientGateways(v []ClientGateway) {
	o.ClientGateways = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadClientGatewaysResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadClientGatewaysResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadClientGatewaysResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadClientGatewaysResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableReadClientGatewaysResponse struct {
	Value        ReadClientGatewaysResponse
	ExplicitNull bool
}

func (v NullableReadClientGatewaysResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableReadClientGatewaysResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
