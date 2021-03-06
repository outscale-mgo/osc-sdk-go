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

// UpdateAccountRequest struct for UpdateAccountRequest
type UpdateAccountRequest struct {
	// The new city of the account owner.
	City *string `json:"City,omitempty"`
	// The new name of the company for the account.
	CompanyName *string `json:"CompanyName,omitempty"`
	// The new country of the account owner.
	Country *string `json:"Country,omitempty"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The new email address for the account.
	Email *string `json:"Email,omitempty"`
	// The new first name of the account owner.
	FirstName *string `json:"FirstName,omitempty"`
	// The new job title of the account owner.
	JobTitle *string `json:"JobTitle,omitempty"`
	// The new last name of the account owner.
	LastName *string `json:"LastName,omitempty"`
	// The new mobile phone number of the account owner.
	MobileNumber *string `json:"MobileNumber,omitempty"`
	// The new landline phone number of the account owner.
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// The new state/province of the account owner.
	StateProvince *string `json:"StateProvince,omitempty"`
	// The new value added tax (VAT) number for the account.
	VatNumber *string `json:"VatNumber,omitempty"`
	// The new ZIP code of the city.
	ZipCode *string `json:"ZipCode,omitempty"`
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetCityOk() (string, bool) {
	if o == nil || o.City == nil {
		var ret string
		return ret, false
	}
	return *o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *UpdateAccountRequest) SetCity(v string) {
	o.City = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetCompanyName() string {
	if o == nil || o.CompanyName == nil {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetCompanyNameOk() (string, bool) {
	if o == nil || o.CompanyName == nil {
		var ret string
		return ret, false
	}
	return *o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasCompanyName() bool {
	if o != nil && o.CompanyName != nil {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *UpdateAccountRequest) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetCountryOk() (string, bool) {
	if o == nil || o.Country == nil {
		var ret string
		return ret, false
	}
	return *o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *UpdateAccountRequest) SetCountry(v string) {
	o.Country = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateAccountRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetEmailOk() (string, bool) {
	if o == nil || o.Email == nil {
		var ret string
		return ret, false
	}
	return *o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateAccountRequest) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetFirstNameOk() (string, bool) {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret, false
	}
	return *o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UpdateAccountRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetJobTitle() string {
	if o == nil || o.JobTitle == nil {
		var ret string
		return ret
	}
	return *o.JobTitle
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetJobTitleOk() (string, bool) {
	if o == nil || o.JobTitle == nil {
		var ret string
		return ret, false
	}
	return *o.JobTitle, true
}

// HasJobTitle returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasJobTitle() bool {
	if o != nil && o.JobTitle != nil {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.
func (o *UpdateAccountRequest) SetJobTitle(v string) {
	o.JobTitle = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetLastNameOk() (string, bool) {
	if o == nil || o.LastName == nil {
		var ret string
		return ret, false
	}
	return *o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UpdateAccountRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetMobileNumber() string {
	if o == nil || o.MobileNumber == nil {
		var ret string
		return ret
	}
	return *o.MobileNumber
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetMobileNumberOk() (string, bool) {
	if o == nil || o.MobileNumber == nil {
		var ret string
		return ret, false
	}
	return *o.MobileNumber, true
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasMobileNumber() bool {
	if o != nil && o.MobileNumber != nil {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given string and assigns it to the MobileNumber field.
func (o *UpdateAccountRequest) SetMobileNumber(v string) {
	o.MobileNumber = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetPhoneNumberOk() (string, bool) {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret, false
	}
	return *o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *UpdateAccountRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetStateProvince returns the StateProvince field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetStateProvince() string {
	if o == nil || o.StateProvince == nil {
		var ret string
		return ret
	}
	return *o.StateProvince
}

// GetStateProvinceOk returns a tuple with the StateProvince field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetStateProvinceOk() (string, bool) {
	if o == nil || o.StateProvince == nil {
		var ret string
		return ret, false
	}
	return *o.StateProvince, true
}

// HasStateProvince returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasStateProvince() bool {
	if o != nil && o.StateProvince != nil {
		return true
	}

	return false
}

// SetStateProvince gets a reference to the given string and assigns it to the StateProvince field.
func (o *UpdateAccountRequest) SetStateProvince(v string) {
	o.StateProvince = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetVatNumber() string {
	if o == nil || o.VatNumber == nil {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetVatNumberOk() (string, bool) {
	if o == nil || o.VatNumber == nil {
		var ret string
		return ret, false
	}
	return *o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasVatNumber() bool {
	if o != nil && o.VatNumber != nil {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *UpdateAccountRequest) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *UpdateAccountRequest) GetZipCode() string {
	if o == nil || o.ZipCode == nil {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetZipCodeOk() (string, bool) {
	if o == nil || o.ZipCode == nil {
		var ret string
		return ret, false
	}
	return *o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *UpdateAccountRequest) HasZipCode() bool {
	if o != nil && o.ZipCode != nil {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *UpdateAccountRequest) SetZipCode(v string) {
	o.ZipCode = &v
}

type NullableUpdateAccountRequest struct {
	Value        UpdateAccountRequest
	ExplicitNull bool
}

func (v NullableUpdateAccountRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateAccountRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
