# CreateKeypairRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**KeypairName** | Pointer to **string** | A unique name for the keypair, with a maximum length of 255 [ASCII printable characters](https://en.wikipedia.org/wiki/ASCII#Printable_characters). | 
**PublicKey** | Pointer to **string** | The public key. It must be base64-encoded. | [optional] 

## Methods

### GetDryRun

`func (o *CreateKeypairRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateKeypairRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateKeypairRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateKeypairRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetKeypairName

`func (o *CreateKeypairRequest) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *CreateKeypairRequest) GetKeypairNameOk() (string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeypairName

`func (o *CreateKeypairRequest) HasKeypairName() bool`

HasKeypairName returns a boolean if a field has been set.

### SetKeypairName

`func (o *CreateKeypairRequest) SetKeypairName(v string)`

SetKeypairName gets a reference to the given string and assigns it to the KeypairName field.

### GetPublicKey

`func (o *CreateKeypairRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CreateKeypairRequest) GetPublicKeyOk() (string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicKey

`func (o *CreateKeypairRequest) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKey

`func (o *CreateKeypairRequest) SetPublicKey(v string)`

SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


