# CheckAuthenticationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Login** | Pointer to **string** | The email address of the account. | [optional] 
**Password** | Pointer to **string** | The password of the account. | [optional] 

## Methods

### GetDryRun

`func (o *CheckAuthenticationRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CheckAuthenticationRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CheckAuthenticationRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CheckAuthenticationRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetLogin

`func (o *CheckAuthenticationRequest) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *CheckAuthenticationRequest) GetLoginOk() (string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogin

`func (o *CheckAuthenticationRequest) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### SetLogin

`func (o *CheckAuthenticationRequest) SetLogin(v string)`

SetLogin gets a reference to the given string and assigns it to the Login field.

### GetPassword

`func (o *CheckAuthenticationRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CheckAuthenticationRequest) GetPasswordOk() (string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPassword

`func (o *CheckAuthenticationRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPassword

`func (o *CheckAuthenticationRequest) SetPassword(v string)`

SetPassword gets a reference to the given string and assigns it to the Password field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


