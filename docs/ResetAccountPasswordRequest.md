# ResetAccountPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Password** | Pointer to **string** | The new password for the account. | 
**Token** | Pointer to **string** | The token you received at the email address provided for the account. | 

## Methods

### GetDryRun

`func (o *ResetAccountPasswordRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ResetAccountPasswordRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *ResetAccountPasswordRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *ResetAccountPasswordRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetPassword

`func (o *ResetAccountPasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ResetAccountPasswordRequest) GetPasswordOk() (string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPassword

`func (o *ResetAccountPasswordRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPassword

`func (o *ResetAccountPasswordRequest) SetPassword(v string)`

SetPassword gets a reference to the given string and assigns it to the Password field.

### GetToken

`func (o *ResetAccountPasswordRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ResetAccountPasswordRequest) GetTokenOk() (string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToken

`func (o *ResetAccountPasswordRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetToken

`func (o *ResetAccountPasswordRequest) SetToken(v string)`

SetToken gets a reference to the given string and assigns it to the Token field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


