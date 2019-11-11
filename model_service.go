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

// Service Information about the service.
type Service struct {
	// The list of network prefixes used by the service, in CIDR notation.
	IpRanges *[]string `json:"IpRanges,omitempty"`
	// The ID of the service.
	ServiceId *string `json:"ServiceId,omitempty"`
	// The name of the prefix list, which identifies the 3DS OUTSCALE service it is associated with.
	ServiceName *string `json:"ServiceName,omitempty"`
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *Service) GetIpRanges() []string {
	if o == nil || o.IpRanges == nil {
		var ret []string
		return ret
	}
	return *o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetIpRangesOk() ([]string, bool) {
	if o == nil || o.IpRanges == nil {
		var ret []string
		return ret, false
	}
	return *o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *Service) HasIpRanges() bool {
	if o != nil && o.IpRanges != nil {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *Service) SetIpRanges(v []string) {
	o.IpRanges = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *Service) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetServiceIdOk() (string, bool) {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret, false
	}
	return *o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *Service) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *Service) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *Service) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetServiceNameOk() (string, bool) {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret, false
	}
	return *o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *Service) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *Service) SetServiceName(v string) {
	o.ServiceName = &v
}

type NullableService struct {
	Value Service
	ExplicitNull bool
}

func (v NullableService) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableService) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

