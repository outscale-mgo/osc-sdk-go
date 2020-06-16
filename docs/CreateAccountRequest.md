# CreateAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | The city of the account owner. | 
**CompanyName** | Pointer to **string** | The name of the company for the account. | 
**Country** | Pointer to **string** | The country of the account owner. | 
**CustomerId** | Pointer to **string** | The ID of the customer. It must be 8 digits (by default, &#x60;12345678&#x60;). | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Email** | Pointer to **string** | The email address for the account. | 
**FirstName** | Pointer to **string** | The first name of the account owner. | 
**JobTitle** | Pointer to **string** | The job title of the account owner. | [optional] 
**LastName** | Pointer to **string** | The last name of the account owner. | 
**MobileNumber** | Pointer to **string** | The mobile phone number of the account owner. | [optional] 
**PhoneNumber** | Pointer to **string** | The landline phone number of the account owner. | [optional] 
**StateProvince** | Pointer to **string** | The state/province of the account. | [optional] 
**VatNumber** | Pointer to **string** | The value added tax (VAT) number for the account. | [optional] 
**ZipCode** | Pointer to **string** | The ZIP code of the city. | 

## Methods

### GetCity

`func (o *CreateAccountRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CreateAccountRequest) GetCityOk() (string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCity

`func (o *CreateAccountRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCity

`func (o *CreateAccountRequest) SetCity(v string)`

SetCity gets a reference to the given string and assigns it to the City field.

### GetCompanyName

`func (o *CreateAccountRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CreateAccountRequest) GetCompanyNameOk() (string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCompanyName

`func (o *CreateAccountRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyName

`func (o *CreateAccountRequest) SetCompanyName(v string)`

SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.

### GetCountry

`func (o *CreateAccountRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreateAccountRequest) GetCountryOk() (string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCountry

`func (o *CreateAccountRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountry

`func (o *CreateAccountRequest) SetCountry(v string)`

SetCountry gets a reference to the given string and assigns it to the Country field.

### GetCustomerId

`func (o *CreateAccountRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateAccountRequest) GetCustomerIdOk() (string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerId

`func (o *CreateAccountRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerId

`func (o *CreateAccountRequest) SetCustomerId(v string)`

SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.

### GetDryRun

`func (o *CreateAccountRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateAccountRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateAccountRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateAccountRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetEmail

`func (o *CreateAccountRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateAccountRequest) GetEmailOk() (string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEmail

`func (o *CreateAccountRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmail

`func (o *CreateAccountRequest) SetEmail(v string)`

SetEmail gets a reference to the given string and assigns it to the Email field.

### GetFirstName

`func (o *CreateAccountRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CreateAccountRequest) GetFirstNameOk() (string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirstName

`func (o *CreateAccountRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstName

`func (o *CreateAccountRequest) SetFirstName(v string)`

SetFirstName gets a reference to the given string and assigns it to the FirstName field.

### GetJobTitle

`func (o *CreateAccountRequest) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *CreateAccountRequest) GetJobTitleOk() (string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasJobTitle

`func (o *CreateAccountRequest) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitle

`func (o *CreateAccountRequest) SetJobTitle(v string)`

SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.

### GetLastName

`func (o *CreateAccountRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CreateAccountRequest) GetLastNameOk() (string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastName

`func (o *CreateAccountRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastName

`func (o *CreateAccountRequest) SetLastName(v string)`

SetLastName gets a reference to the given string and assigns it to the LastName field.

### GetMobileNumber

`func (o *CreateAccountRequest) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *CreateAccountRequest) GetMobileNumberOk() (string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMobileNumber

`func (o *CreateAccountRequest) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### SetMobileNumber

`func (o *CreateAccountRequest) SetMobileNumber(v string)`

SetMobileNumber gets a reference to the given string and assigns it to the MobileNumber field.

### GetPhoneNumber

`func (o *CreateAccountRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CreateAccountRequest) GetPhoneNumberOk() (string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPhoneNumber

`func (o *CreateAccountRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumber

`func (o *CreateAccountRequest) SetPhoneNumber(v string)`

SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.

### GetStateProvince

`func (o *CreateAccountRequest) GetStateProvince() string`

GetStateProvince returns the StateProvince field if non-nil, zero value otherwise.

### GetStateProvinceOk

`func (o *CreateAccountRequest) GetStateProvinceOk() (string, bool)`

GetStateProvinceOk returns a tuple with the StateProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStateProvince

`func (o *CreateAccountRequest) HasStateProvince() bool`

HasStateProvince returns a boolean if a field has been set.

### SetStateProvince

`func (o *CreateAccountRequest) SetStateProvince(v string)`

SetStateProvince gets a reference to the given string and assigns it to the StateProvince field.

### GetVatNumber

`func (o *CreateAccountRequest) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *CreateAccountRequest) GetVatNumberOk() (string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVatNumber

`func (o *CreateAccountRequest) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### SetVatNumber

`func (o *CreateAccountRequest) SetVatNumber(v string)`

SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.

### GetZipCode

`func (o *CreateAccountRequest) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *CreateAccountRequest) GetZipCodeOk() (string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasZipCode

`func (o *CreateAccountRequest) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCode

`func (o *CreateAccountRequest) SetZipCode(v string)`

SetZipCode gets a reference to the given string and assigns it to the ZipCode field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


