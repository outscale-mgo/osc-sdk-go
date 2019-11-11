# ReadKeypairsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keypairs** | Pointer to [**[]Keypair**](Keypair.md) | Information about one or more keypairs. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetKeypairs

`func (o *ReadKeypairsResponse) GetKeypairs() []Keypair`

GetKeypairs returns the Keypairs field if non-nil, zero value otherwise.

### GetKeypairsOk

`func (o *ReadKeypairsResponse) GetKeypairsOk() ([]Keypair, bool)`

GetKeypairsOk returns a tuple with the Keypairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeypairs

`func (o *ReadKeypairsResponse) HasKeypairs() bool`

HasKeypairs returns a boolean if a field has been set.

### SetKeypairs

`func (o *ReadKeypairsResponse) SetKeypairs(v []Keypair)`

SetKeypairs gets a reference to the given []Keypair and assigns it to the Keypairs field.

### GetResponseContext

`func (o *ReadKeypairsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadKeypairsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadKeypairsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadKeypairsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


