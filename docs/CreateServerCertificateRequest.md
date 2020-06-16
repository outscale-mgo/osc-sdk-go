# CreateServerCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | The PEM-encoded X509 certificate. | 
**Chain** | Pointer to **string** | The PEM-encoded intermediate certification authorities. | [optional] 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Name** | Pointer to **string** | A unique name for the certificate. Constraints: 1-128 alphanumeric characters, pluses (+), equals (&#x3D;), commas (,), periods (.), at signs (@), minuses (-), or underscores (_). | 
**Path** | Pointer to **string** | The path to the server certificate, set to a slash (/) if not specified. | [optional] 
**PrivateKey** | Pointer to **string** | The PEM-encoded private key matching the certificate. | 

## Methods

### GetBody

`func (o *CreateServerCertificateRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateServerCertificateRequest) GetBodyOk() (string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBody

`func (o *CreateServerCertificateRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBody

`func (o *CreateServerCertificateRequest) SetBody(v string)`

SetBody gets a reference to the given string and assigns it to the Body field.

### GetChain

`func (o *CreateServerCertificateRequest) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *CreateServerCertificateRequest) GetChainOk() (string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChain

`func (o *CreateServerCertificateRequest) HasChain() bool`

HasChain returns a boolean if a field has been set.

### SetChain

`func (o *CreateServerCertificateRequest) SetChain(v string)`

SetChain gets a reference to the given string and assigns it to the Chain field.

### GetDryRun

`func (o *CreateServerCertificateRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateServerCertificateRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateServerCertificateRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateServerCertificateRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetName

`func (o *CreateServerCertificateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServerCertificateRequest) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CreateServerCertificateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CreateServerCertificateRequest) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetPath

`func (o *CreateServerCertificateRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateServerCertificateRequest) GetPathOk() (string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPath

`func (o *CreateServerCertificateRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPath

`func (o *CreateServerCertificateRequest) SetPath(v string)`

SetPath gets a reference to the given string and assigns it to the Path field.

### GetPrivateKey

`func (o *CreateServerCertificateRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CreateServerCertificateRequest) GetPrivateKeyOk() (string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateKey

`func (o *CreateServerCertificateRequest) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKey

`func (o *CreateServerCertificateRequest) SetPrivateKey(v string)`

SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


