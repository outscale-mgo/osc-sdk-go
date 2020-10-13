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

// HealthCheck Information about the health check configuration.
type HealthCheck struct {
	// The number of seconds between two pings (between `5` and `600` both included).
	CheckInterval int64 `json:"CheckInterval"`
	// The number of consecutive successful pings before considering the VM as healthy (between `2` and `10` both included).
	HealthyThreshold int64 `json:"HealthyThreshold"`
	// The path for HTTP or HTTPS requests.
	Path *string `json:"Path,omitempty"`
	// The port number (between `1` and `65535`, both included).
	Port int64 `json:"Port"`
	// The protocol for the URL of the VM (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL` \\| `UDP`).
	Protocol string `json:"Protocol"`
	// The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between `2` and `60` both included).
	Timeout int64 `json:"Timeout"`
	// The number of consecutive failed pings before considering the VM as unhealthy (between `2` and `10` both included).
	UnhealthyThreshold int64 `json:"UnhealthyThreshold"`
}

// GetCheckInterval returns the CheckInterval field value
func (o *HealthCheck) GetCheckInterval() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CheckInterval
}

// SetCheckInterval sets field value
func (o *HealthCheck) SetCheckInterval(v int64) {
	o.CheckInterval = v
}

// GetHealthyThreshold returns the HealthyThreshold field value
func (o *HealthCheck) GetHealthyThreshold() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.HealthyThreshold
}

// SetHealthyThreshold sets field value
func (o *HealthCheck) SetHealthyThreshold(v int64) {
	o.HealthyThreshold = v
}

// GetPath returns the Path field value
func (o *HealthCheck) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// SetPath sets field value
func (o *HealthCheck) SetPath(v string) {
	o.Path = v
}

// GetPort returns the Port field value
func (o *HealthCheck) GetPort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Port
}

// SetPort sets field value
func (o *HealthCheck) SetPort(v int64) {
	o.Port = v
}

// GetProtocol returns the Protocol field value
func (o *HealthCheck) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// SetProtocol sets field value
func (o *HealthCheck) SetProtocol(v string) {
	o.Protocol = v
}

// GetTimeout returns the Timeout field value
func (o *HealthCheck) GetTimeout() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timeout
}

// SetTimeout sets field value
func (o *HealthCheck) SetTimeout(v int64) {
	o.Timeout = v
}

// GetUnhealthyThreshold returns the UnhealthyThreshold field value
func (o *HealthCheck) GetUnhealthyThreshold() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UnhealthyThreshold
}

// SetUnhealthyThreshold sets field value
func (o *HealthCheck) SetUnhealthyThreshold(v int64) {
	o.UnhealthyThreshold = v
}

type NullableHealthCheck struct {
	Value        HealthCheck
	ExplicitNull bool
}

func (v NullableHealthCheck) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableHealthCheck) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
