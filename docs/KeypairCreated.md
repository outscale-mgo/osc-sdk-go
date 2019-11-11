# KeypairCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeypairFingerprint** | Pointer to **string** | If you create a keypair, the SHA-1 digest of the DER encoded private key.&lt;br /&gt; If you import a keypair, the MD5 public key fingerprint as specified in section 4 of RFC 4716. | [optional] 
**KeypairName** | Pointer to **string** | The name of the keypair. | [optional] 
**PrivateKey** | Pointer to **string** | The private key. | [optional] 

## Methods

### GetKeypairFingerprint

`func (o *KeypairCreated) GetKeypairFingerprint() string`

GetKeypairFingerprint returns the KeypairFingerprint field if non-nil, zero value otherwise.

### GetKeypairFingerprintOk

`func (o *KeypairCreated) GetKeypairFingerprintOk() (string, bool)`

GetKeypairFingerprintOk returns a tuple with the KeypairFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeypairFingerprint

`func (o *KeypairCreated) HasKeypairFingerprint() bool`

HasKeypairFingerprint returns a boolean if a field has been set.

### SetKeypairFingerprint

`func (o *KeypairCreated) SetKeypairFingerprint(v string)`

SetKeypairFingerprint gets a reference to the given string and assigns it to the KeypairFingerprint field.

### GetKeypairName

`func (o *KeypairCreated) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *KeypairCreated) GetKeypairNameOk() (string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeypairName

`func (o *KeypairCreated) HasKeypairName() bool`

HasKeypairName returns a boolean if a field has been set.

### SetKeypairName

`func (o *KeypairCreated) SetKeypairName(v string)`

SetKeypairName gets a reference to the given string and assigns it to the KeypairName field.

### GetPrivateKey

`func (o *KeypairCreated) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *KeypairCreated) GetPrivateKeyOk() (string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateKey

`func (o *KeypairCreated) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKey

`func (o *KeypairCreated) SetPrivateKey(v string)`

SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


