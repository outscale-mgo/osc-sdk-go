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

// ProductType Information about the product type.
type ProductType struct {
	// The description of the product type.
	Description *string `json:"Description,omitempty"`
	// The ID of the product type.
	ProductTypeId *string `json:"ProductTypeId,omitempty"`
	// The vendor of the product type.
	Vendor *string `json:"Vendor,omitempty"`
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductType) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ProductType) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductType) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProductType) SetDescription(v string) {
	o.Description = &v
}

// GetProductTypeId returns the ProductTypeId field value if set, zero value otherwise.
func (o *ProductType) GetProductTypeId() string {
	if o == nil || o.ProductTypeId == nil {
		var ret string
		return ret
	}
	return *o.ProductTypeId
}

// GetProductTypeIdOk returns a tuple with the ProductTypeId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ProductType) GetProductTypeIdOk() (string, bool) {
	if o == nil || o.ProductTypeId == nil {
		var ret string
		return ret, false
	}
	return *o.ProductTypeId, true
}

// HasProductTypeId returns a boolean if a field has been set.
func (o *ProductType) HasProductTypeId() bool {
	if o != nil && o.ProductTypeId != nil {
		return true
	}

	return false
}

// SetProductTypeId gets a reference to the given string and assigns it to the ProductTypeId field.
func (o *ProductType) SetProductTypeId(v string) {
	o.ProductTypeId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *ProductType) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ProductType) GetVendorOk() (string, bool) {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret, false
	}
	return *o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *ProductType) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *ProductType) SetVendor(v string) {
	o.Vendor = &v
}

type NullableProductType struct {
	Value ProductType
	ExplicitNull bool
}

func (v NullableProductType) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableProductType) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}

