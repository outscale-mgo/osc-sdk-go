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

// UpdateServerCertificateRequest struct for UpdateServerCertificateRequest
type UpdateServerCertificateRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The name of the server certificate you want to modify.
	Name *string `json:"Name,omitempty"`
	// A new name for the server certificate.
	NewName *string `json:"NewName,omitempty"`
	// A new path for the server certificate.
	NewPath *string `json:"NewPath,omitempty"`
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateServerCertificateRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServerCertificateRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateServerCertificateRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateServerCertificateRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateServerCertificateRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServerCertificateRequest) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateServerCertificateRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateServerCertificateRequest) SetName(v string) {
	o.Name = &v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *UpdateServerCertificateRequest) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServerCertificateRequest) GetNewNameOk() (string, bool) {
	if o == nil || o.NewName == nil {
		var ret string
		return ret, false
	}
	return *o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *UpdateServerCertificateRequest) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *UpdateServerCertificateRequest) SetNewName(v string) {
	o.NewName = &v
}

// GetNewPath returns the NewPath field value if set, zero value otherwise.
func (o *UpdateServerCertificateRequest) GetNewPath() string {
	if o == nil || o.NewPath == nil {
		var ret string
		return ret
	}
	return *o.NewPath
}

// GetNewPathOk returns a tuple with the NewPath field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServerCertificateRequest) GetNewPathOk() (string, bool) {
	if o == nil || o.NewPath == nil {
		var ret string
		return ret, false
	}
	return *o.NewPath, true
}

// HasNewPath returns a boolean if a field has been set.
func (o *UpdateServerCertificateRequest) HasNewPath() bool {
	if o != nil && o.NewPath != nil {
		return true
	}

	return false
}

// SetNewPath gets a reference to the given string and assigns it to the NewPath field.
func (o *UpdateServerCertificateRequest) SetNewPath(v string) {
	o.NewPath = &v
}

type NullableUpdateServerCertificateRequest struct {
	Value        UpdateServerCertificateRequest
	ExplicitNull bool
}

func (v NullableUpdateServerCertificateRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateServerCertificateRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}