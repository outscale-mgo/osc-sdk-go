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

// ReadNatServicesResponse struct for ReadNatServicesResponse
type ReadNatServicesResponse struct {
	// Information about one or more NAT services.
	NatServices     *[]NatService    `json:"NatServices,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetNatServices returns the NatServices field value if set, zero value otherwise.
func (o *ReadNatServicesResponse) GetNatServices() []NatService {
	if o == nil || o.NatServices == nil {
		var ret []NatService
		return ret
	}
	return *o.NatServices
}

// GetNatServicesOk returns a tuple with the NatServices field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadNatServicesResponse) GetNatServicesOk() ([]NatService, bool) {
	if o == nil || o.NatServices == nil {
		var ret []NatService
		return ret, false
	}
	return *o.NatServices, true
}

// HasNatServices returns a boolean if a field has been set.
func (o *ReadNatServicesResponse) HasNatServices() bool {
	if o != nil && o.NatServices != nil {
		return true
	}

	return false
}

// SetNatServices gets a reference to the given []NatService and assigns it to the NatServices field.
func (o *ReadNatServicesResponse) SetNatServices(v []NatService) {
	o.NatServices = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadNatServicesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadNatServicesResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadNatServicesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadNatServicesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableReadNatServicesResponse struct {
	Value        ReadNatServicesResponse
	ExplicitNull bool
}

func (v NullableReadNatServicesResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableReadNatServicesResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
