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

// ReadLocationsResponse struct for ReadLocationsResponse
type ReadLocationsResponse struct {
	// Information about one or more locations.
	Locations *[]Location `json:"Locations,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *ReadLocationsResponse) GetLocations() []Location {
	if o == nil || o.Locations == nil {
		var ret []Location
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadLocationsResponse) GetLocationsOk() ([]Location, bool) {
	if o == nil || o.Locations == nil {
		var ret []Location
		return ret, false
	}
	return *o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *ReadLocationsResponse) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []Location and assigns it to the Locations field.
func (o *ReadLocationsResponse) SetLocations(v []Location) {
	o.Locations = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadLocationsResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadLocationsResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadLocationsResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadLocationsResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

type NullableReadLocationsResponse struct {
	Value ReadLocationsResponse
	ExplicitNull bool
}

func (v NullableReadLocationsResponse) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableReadLocationsResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

