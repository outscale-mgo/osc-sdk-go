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

// NicForVmCreation Information about the network interface card (NIC) when creating a virtual machine (VM).
type NicForVmCreation struct {
	// If `true`, the NIC is deleted when the VM is terminated. You can specify `true` only if you create a NIC when creating a VM.
	DeleteOnVmDeletion *bool `json:"DeleteOnVmDeletion,omitempty"`
	// The description of the NIC, if you are creating a NIC when creating the VM.
	Description *string `json:"Description,omitempty"`
	// The index of the VM device for the NIC attachment (between 0 and 7, both included). This parameter is required if you create a NIC when creating the VM.
	DeviceNumber *int64 `json:"DeviceNumber,omitempty"`
	// The ID of the NIC, if you are attaching an existing NIC when creating a VM.
	NicId *string `json:"NicId,omitempty"`
	// One or more private IP addresses to assign to the NIC, if you create a NIC when creating a VM. Only one private IP address can be the primary private IP address.
	PrivateIps *[]PrivateIpLight `json:"PrivateIps,omitempty"`
	// The number of secondary private IP addresses, if you create a NIC when creating a VM. This parameter cannot be specified if you specified more than one private IP address in the `PrivateIps` parameter.
	SecondaryPrivateIpCount *int64 `json:"SecondaryPrivateIpCount,omitempty"`
	// One or more IDs of security groups for the NIC, if you acreate a NIC when creating a VM.
	SecurityGroupIds *[]string `json:"SecurityGroupIds,omitempty"`
	// The ID of the Subnet for the NIC, if you create a NIC when creating a VM.
	SubnetId *string `json:"SubnetId,omitempty"`
}

// GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field value if set, zero value otherwise.
func (o *NicForVmCreation) GetDeleteOnVmDeletion() bool {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret
	}
	return *o.DeleteOnVmDeletion
}

// GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NicForVmCreation) GetDeleteOnVmDeletionOk() (bool, bool) {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret, false
	}
	return *o.DeleteOnVmDeletion, true
}

// HasDeleteOnVmDeletion returns a boolean if a field has been set.
func (o *NicForVmCreation) HasDeleteOnVmDeletion() bool {
	if o != nil && o.DeleteOnVmDeletion != nil {
		return true
	}

	return false
}

// SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.
func (o *NicForVmCreation) SetDeleteOnVmDeletion(v bool) {
	o.DeleteOnVmDeletion = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NicForVmCreation) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NicForVmCreation) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NicForVmCreation) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NicForVmCreation) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceNumber returns the DeviceNumber field value if set, zero value otherwise.
func (o *NicForVmCreation) GetDeviceNumber() int64 {
	if o == nil || o.DeviceNumber == nil {
		var ret int64
		return ret
	}
	return *o.DeviceNumber
}

// GetDeviceNumberOk returns a tuple with the DeviceNumber field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NicForVmCreation) GetDeviceNumberOk() (int64, bool) {
	if o == nil || o.DeviceNumber == nil {
		var ret int64
		return ret, false
	}
	return *o.DeviceNumber, true
}

// HasDeviceNumber returns a boolean if a field has been set.
func (o *NicForVmCreation) HasDeviceNumber() bool {
	if o != nil && o.DeviceNumber != nil {
		return true
	}

	return false
}

// SetDeviceNumber gets a reference to the given int64 and assigns it to the DeviceNumber field.
func (o *NicForVmCreation) SetDeviceNumber(v int64) {
	o.DeviceNumber = &v
}

// GetNicId returns the NicId field value if set, zero value otherwise.
func (o *NicForVmCreation) GetNicId() string {
	if o == nil || o.NicId == nil {
		var ret string
		return ret
	}
	return *o.NicId
}

// GetNicIdOk returns a tuple with the NicId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NicForVmCreation) GetNicIdOk() (string, bool) {
	if o == nil || o.NicId == nil {
		var ret string
		return ret, false
	}
	return *o.NicId, true
}

// HasNicId returns a boolean if a field has been set.
func (o *NicForVmCreation) HasNicId() bool {
	if o != nil && o.NicId != nil {
		return true
	}

	return false
}

// SetNicId gets a reference to the given string and assigns it to the NicId field.
func (o *NicForVmCreation) SetNicId(v string) {
	o.NicId = &v
}

// GetPrivateIps returns the PrivateIps field value if set, zero value otherwise.
func (o *NicForVmCreation) GetPrivateIps() []PrivateIpLight {
	if o == nil || o.PrivateIps == nil {
		var ret []PrivateIpLight
		return ret
	}
	return *o.PrivateIps
}

// GetPrivateIpsOk returns a tuple with the PrivateIps field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NicForVmCreation) GetPrivateIpsOk() ([]PrivateIpLight, bool) {
	if o == nil || o.PrivateIps == nil {
		var ret []PrivateIpLight
		return ret, false
	}
	return *o.PrivateIps, true
}

// HasPrivateIps returns a boolean if a field has been set.
func (o *NicForVmCreation) HasPrivateIps() bool {
	if o != nil && o.PrivateIps != nil {
		return true
	}

	return false
}

// SetPrivateIps gets a reference to the given []PrivateIpLight and assigns it to the PrivateIps field.
func (o *NicForVmCreation) SetPrivateIps(v []PrivateIpLight) {
	o.PrivateIps = &v
}

// GetSecondaryPrivateIpCount returns the SecondaryPrivateIpCount field value if set, zero value otherwise.
func (o *NicForVmCreation) GetSecondaryPrivateIpCount() int64 {
	if o == nil || o.SecondaryPrivateIpCount == nil {
		var ret int64
		return ret
	}
	return *o.SecondaryPrivateIpCount
}

// GetSecondaryPrivateIpCountOk returns a tuple with the SecondaryPrivateIpCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NicForVmCreation) GetSecondaryPrivateIpCountOk() (int64, bool) {
	if o == nil || o.SecondaryPrivateIpCount == nil {
		var ret int64
		return ret, false
	}
	return *o.SecondaryPrivateIpCount, true
}

// HasSecondaryPrivateIpCount returns a boolean if a field has been set.
func (o *NicForVmCreation) HasSecondaryPrivateIpCount() bool {
	if o != nil && o.SecondaryPrivateIpCount != nil {
		return true
	}

	return false
}

// SetSecondaryPrivateIpCount gets a reference to the given int64 and assigns it to the SecondaryPrivateIpCount field.
func (o *NicForVmCreation) SetSecondaryPrivateIpCount(v int64) {
	o.SecondaryPrivateIpCount = &v
}

// GetSecurityGroupIds returns the SecurityGroupIds field value if set, zero value otherwise.
func (o *NicForVmCreation) GetSecurityGroupIds() []string {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroupIds
}

// GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NicForVmCreation) GetSecurityGroupIdsOk() ([]string, bool) {
	if o == nil || o.SecurityGroupIds == nil {
		var ret []string
		return ret, false
	}
	return *o.SecurityGroupIds, true
}

// HasSecurityGroupIds returns a boolean if a field has been set.
func (o *NicForVmCreation) HasSecurityGroupIds() bool {
	if o != nil && o.SecurityGroupIds != nil {
		return true
	}

	return false
}

// SetSecurityGroupIds gets a reference to the given []string and assigns it to the SecurityGroupIds field.
func (o *NicForVmCreation) SetSecurityGroupIds(v []string) {
	o.SecurityGroupIds = &v
}

// GetSubnetId returns the SubnetId field value if set, zero value otherwise.
func (o *NicForVmCreation) GetSubnetId() string {
	if o == nil || o.SubnetId == nil {
		var ret string
		return ret
	}
	return *o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NicForVmCreation) GetSubnetIdOk() (string, bool) {
	if o == nil || o.SubnetId == nil {
		var ret string
		return ret, false
	}
	return *o.SubnetId, true
}

// HasSubnetId returns a boolean if a field has been set.
func (o *NicForVmCreation) HasSubnetId() bool {
	if o != nil && o.SubnetId != nil {
		return true
	}

	return false
}

// SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.
func (o *NicForVmCreation) SetSubnetId(v string) {
	o.SubnetId = &v
}

type NullableNicForVmCreation struct {
	Value        NicForVmCreation
	ExplicitNull bool
}

func (v NullableNicForVmCreation) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNicForVmCreation) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
