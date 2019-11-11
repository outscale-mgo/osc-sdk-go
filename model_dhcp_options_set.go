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

// DhcpOptionsSet Information about the DHCP options set.
type DhcpOptionsSet struct {
	// If `true`, the DHCP options set is a default one. If `false`, it is not.
	Default *bool `json:"Default,omitempty"`
	// The name of the DHCP options set.
	DhcpOptionsName *string `json:"DhcpOptionsName,omitempty"`
	// The ID of the DHCP options set.
	DhcpOptionsSetId *string `json:"DhcpOptionsSetId,omitempty"`
	// The domain name.
	DomainName *string `json:"DomainName,omitempty"`
	// One or more IP addresses for the domain name servers.
	DomainNameServers *[]string `json:"DomainNameServers,omitempty"`
	// One or more IP addresses for the NTP servers.
	NtpServers *[]string `json:"NtpServers,omitempty"`
	// One or more tags associated with the DHCP options set.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *DhcpOptionsSet) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DhcpOptionsSet) GetDefaultOk() (bool, bool) {
	if o == nil || o.Default == nil {
		var ret bool
		return ret, false
	}
	return *o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *DhcpOptionsSet) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *DhcpOptionsSet) SetDefault(v bool) {
	o.Default = &v
}

// GetDhcpOptionsName returns the DhcpOptionsName field value if set, zero value otherwise.
func (o *DhcpOptionsSet) GetDhcpOptionsName() string {
	if o == nil || o.DhcpOptionsName == nil {
		var ret string
		return ret
	}
	return *o.DhcpOptionsName
}

// GetDhcpOptionsNameOk returns a tuple with the DhcpOptionsName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DhcpOptionsSet) GetDhcpOptionsNameOk() (string, bool) {
	if o == nil || o.DhcpOptionsName == nil {
		var ret string
		return ret, false
	}
	return *o.DhcpOptionsName, true
}

// HasDhcpOptionsName returns a boolean if a field has been set.
func (o *DhcpOptionsSet) HasDhcpOptionsName() bool {
	if o != nil && o.DhcpOptionsName != nil {
		return true
	}

	return false
}

// SetDhcpOptionsName gets a reference to the given string and assigns it to the DhcpOptionsName field.
func (o *DhcpOptionsSet) SetDhcpOptionsName(v string) {
	o.DhcpOptionsName = &v
}

// GetDhcpOptionsSetId returns the DhcpOptionsSetId field value if set, zero value otherwise.
func (o *DhcpOptionsSet) GetDhcpOptionsSetId() string {
	if o == nil || o.DhcpOptionsSetId == nil {
		var ret string
		return ret
	}
	return *o.DhcpOptionsSetId
}

// GetDhcpOptionsSetIdOk returns a tuple with the DhcpOptionsSetId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DhcpOptionsSet) GetDhcpOptionsSetIdOk() (string, bool) {
	if o == nil || o.DhcpOptionsSetId == nil {
		var ret string
		return ret, false
	}
	return *o.DhcpOptionsSetId, true
}

// HasDhcpOptionsSetId returns a boolean if a field has been set.
func (o *DhcpOptionsSet) HasDhcpOptionsSetId() bool {
	if o != nil && o.DhcpOptionsSetId != nil {
		return true
	}

	return false
}

// SetDhcpOptionsSetId gets a reference to the given string and assigns it to the DhcpOptionsSetId field.
func (o *DhcpOptionsSet) SetDhcpOptionsSetId(v string) {
	o.DhcpOptionsSetId = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *DhcpOptionsSet) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DhcpOptionsSet) GetDomainNameOk() (string, bool) {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret, false
	}
	return *o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *DhcpOptionsSet) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *DhcpOptionsSet) SetDomainName(v string) {
	o.DomainName = &v
}

// GetDomainNameServers returns the DomainNameServers field value if set, zero value otherwise.
func (o *DhcpOptionsSet) GetDomainNameServers() []string {
	if o == nil || o.DomainNameServers == nil {
		var ret []string
		return ret
	}
	return *o.DomainNameServers
}

// GetDomainNameServersOk returns a tuple with the DomainNameServers field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DhcpOptionsSet) GetDomainNameServersOk() ([]string, bool) {
	if o == nil || o.DomainNameServers == nil {
		var ret []string
		return ret, false
	}
	return *o.DomainNameServers, true
}

// HasDomainNameServers returns a boolean if a field has been set.
func (o *DhcpOptionsSet) HasDomainNameServers() bool {
	if o != nil && o.DomainNameServers != nil {
		return true
	}

	return false
}

// SetDomainNameServers gets a reference to the given []string and assigns it to the DomainNameServers field.
func (o *DhcpOptionsSet) SetDomainNameServers(v []string) {
	o.DomainNameServers = &v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise.
func (o *DhcpOptionsSet) GetNtpServers() []string {
	if o == nil || o.NtpServers == nil {
		var ret []string
		return ret
	}
	return *o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DhcpOptionsSet) GetNtpServersOk() ([]string, bool) {
	if o == nil || o.NtpServers == nil {
		var ret []string
		return ret, false
	}
	return *o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *DhcpOptionsSet) HasNtpServers() bool {
	if o != nil && o.NtpServers != nil {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *DhcpOptionsSet) SetNtpServers(v []string) {
	o.NtpServers = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DhcpOptionsSet) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DhcpOptionsSet) GetTagsOk() ([]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret, false
	}
	return *o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DhcpOptionsSet) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *DhcpOptionsSet) SetTags(v []ResourceTag) {
	o.Tags = &v
}

type NullableDhcpOptionsSet struct {
	Value DhcpOptionsSet
	ExplicitNull bool
}

func (v NullableDhcpOptionsSet) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableDhcpOptionsSet) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

