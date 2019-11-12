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

// ReadNicsResponse struct for ReadNicsResponse
type ReadNicsResponse struct {
	// Information about one or more NICs.
	Nics *[]Nic `json:"Nics,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetNics returns the Nics field value if set, zero value otherwise.
func (o *ReadNicsResponse) GetNics() []Nic {
	if o == nil || o.Nics == nil {
		var ret []Nic
		return ret
	}
	return *o.Nics
}

// GetNicsOk returns a tuple with the Nics field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadNicsResponse) GetNicsOk() ([]Nic, bool) {
	if o == nil || o.Nics == nil {
		var ret []Nic
		return ret, false
	}
	return *o.Nics, true
}

// HasNics returns a boolean if a field has been set.
func (o *ReadNicsResponse) HasNics() bool {
	if o != nil && o.Nics != nil {
		return true
	}

	return false
}

// SetNics gets a reference to the given []Nic and assigns it to the Nics field.
func (o *ReadNicsResponse) SetNics(v []Nic) {
	o.Nics = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadNicsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadNicsResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadNicsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadNicsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableReadNicsResponse struct {
	Value ReadNicsResponse
	ExplicitNull bool
}

func (v NullableReadNicsResponse) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableReadNicsResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

