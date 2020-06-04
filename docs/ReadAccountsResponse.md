# ReadAccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]Account**](Account.md) | The list of the accounts. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetAccounts

`func (o *ReadAccountsResponse) GetAccounts() []Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ReadAccountsResponse) GetAccountsOk() ([]Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccounts

`func (o *ReadAccountsResponse) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### SetAccounts

`func (o *ReadAccountsResponse) SetAccounts(v []Account)`

SetAccounts gets a reference to the given []Account and assigns it to the Accounts field.

### GetResponseContext

`func (o *ReadAccountsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadAccountsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadAccountsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadAccountsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


