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

// CreateSnapshotResponse struct for CreateSnapshotResponse
type CreateSnapshotResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
	Snapshot *Snapshot `json:"Snapshot,omitempty"`
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateSnapshotResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotResponse) GetResponseContextOk() (ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret, false
	}
	return *o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateSnapshotResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateSnapshotResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *CreateSnapshotResponse) GetSnapshot() Snapshot {
	if o == nil || o.Snapshot == nil {
		var ret Snapshot
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotResponse) GetSnapshotOk() (Snapshot, bool) {
	if o == nil || o.Snapshot == nil {
		var ret Snapshot
		return ret, false
	}
	return *o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *CreateSnapshotResponse) HasSnapshot() bool {
	if o != nil && o.Snapshot != nil {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given Snapshot and assigns it to the Snapshot field.
func (o *CreateSnapshotResponse) SetSnapshot(v Snapshot) {
	o.Snapshot = &v
}

type NullableCreateSnapshotResponse struct {
	Value CreateSnapshotResponse
	ExplicitNull bool
}

func (v NullableCreateSnapshotResponse) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableCreateSnapshotResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

