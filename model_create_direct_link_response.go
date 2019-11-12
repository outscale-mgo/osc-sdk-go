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

// CreateDirectLinkResponse struct for CreateDirectLinkResponse
type CreateDirectLinkResponse struct {
	DirectLink *DirectLink `json:"DirectLink,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetDirectLink returns the DirectLink field value if set, zero value otherwise.
func (o *CreateDirectLinkResponse) GetDirectLink() DirectLink {
	if o == nil || o.DirectLink == nil {
		var ret DirectLink
		return ret
	}
	return *o.DirectLink
}

// GetDirectLinkOk returns a tuple with the DirectLink field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateDirectLinkResponse) GetDirectLinkOk() (DirectLink, bool) {
	if o == nil || o.DirectLink == nil {
		var ret DirectLink
		return ret, false
	}
	return *o.DirectLink, true
}

// HasDirectLink returns a boolean if a field has been set.
func (o *CreateDirectLinkResponse) HasDirectLink() bool {
	if o != nil && o.DirectLink != nil {
		return true
	}

	return false
}

// SetDirectLink gets a reference to the given DirectLink and assigns it to the DirectLink field.
func (o *CreateDirectLinkResponse) SetDirectLink(v DirectLink) {
	o.DirectLink = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateDirectLinkResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateDirectLinkResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateDirectLinkResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateDirectLinkResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableCreateDirectLinkResponse struct {
	Value CreateDirectLinkResponse
	ExplicitNull bool
}

func (v NullableCreateDirectLinkResponse) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableCreateDirectLinkResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

