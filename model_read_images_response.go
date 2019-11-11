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

// ReadImagesResponse struct for ReadImagesResponse
type ReadImagesResponse struct {
	// Information about one or more OMIs.
	Images *[]Image `json:"Images,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *ReadImagesResponse) GetImages() []Image {
	if o == nil || o.Images == nil {
		var ret []Image
		return ret
	}
	return *o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadImagesResponse) GetImagesOk() ([]Image, bool) {
	if o == nil || o.Images == nil {
		var ret []Image
		return ret, false
	}
	return *o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *ReadImagesResponse) HasImages() bool {
	if o != nil && o.Images != nil {
		return true
	}

	return false
}

// SetImages gets a reference to the given []Image and assigns it to the Images field.
func (o *ReadImagesResponse) SetImages(v []Image) {
	o.Images = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadImagesResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadImagesResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadImagesResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadImagesResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableReadImagesResponse struct {
	Value ReadImagesResponse
	ExplicitNull bool
}

func (v NullableReadImagesResponse) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableReadImagesResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

