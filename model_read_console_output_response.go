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

// ReadConsoleOutputResponse struct for ReadConsoleOutputResponse
type ReadConsoleOutputResponse struct {
	// The Base64-encoded output of the console. If a command line tool is used, the output is decoded by the tool.
	ConsoleOutput   *string          `json:"ConsoleOutput,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	// The ID of the VM.
	VmId *string `json:"VmId,omitempty"`
}

// GetConsoleOutput returns the ConsoleOutput field value if set, zero value otherwise.
func (o *ReadConsoleOutputResponse) GetConsoleOutput() string {
	if o == nil || o.ConsoleOutput == nil {
		var ret string
		return ret
	}
	return *o.ConsoleOutput
}

// GetConsoleOutputOk returns a tuple with the ConsoleOutput field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadConsoleOutputResponse) GetConsoleOutputOk() (string, bool) {
	if o == nil || o.ConsoleOutput == nil {
		var ret string
		return ret, false
	}
	return *o.ConsoleOutput, true
}

// HasConsoleOutput returns a boolean if a field has been set.
func (o *ReadConsoleOutputResponse) HasConsoleOutput() bool {
	if o != nil && o.ConsoleOutput != nil {
		return true
	}

	return false
}

// SetConsoleOutput gets a reference to the given string and assigns it to the ConsoleOutput field.
func (o *ReadConsoleOutputResponse) SetConsoleOutput(v string) {
	o.ConsoleOutput = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *ReadConsoleOutputResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadConsoleOutputResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *ReadConsoleOutputResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *ReadConsoleOutputResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *ReadConsoleOutputResponse) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ReadConsoleOutputResponse) GetVmIdOk() (string, bool) {
	if o == nil || o.VmId == nil {
		var ret string
		return ret, false
	}
	return *o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *ReadConsoleOutputResponse) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *ReadConsoleOutputResponse) SetVmId(v string) {
	o.VmId = &v
}

type NullableReadConsoleOutputResponse struct {
	Value        ReadConsoleOutputResponse
	ExplicitNull bool
}

func (v NullableReadConsoleOutputResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableReadConsoleOutputResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
