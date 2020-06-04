# ReadAccessKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeys** | Pointer to [**[]AccessKey**](AccessKey.md) | A list of access keys. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetAccessKeys

`func (o *ReadAccessKeysResponse) GetAccessKeys() []AccessKey`

GetAccessKeys returns the AccessKeys field if non-nil, zero value otherwise.

### GetAccessKeysOk

`func (o *ReadAccessKeysResponse) GetAccessKeysOk() ([]AccessKey, bool)`

GetAccessKeysOk returns a tuple with the AccessKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccessKeys

`func (o *ReadAccessKeysResponse) HasAccessKeys() bool`

HasAccessKeys returns a boolean if a field has been set.

### SetAccessKeys

`func (o *ReadAccessKeysResponse) SetAccessKeys(v []AccessKey)`

SetAccessKeys gets a reference to the given []AccessKey and assigns it to the AccessKeys field.

### GetResponseContext

`func (o *ReadAccessKeysResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadAccessKeysResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadAccessKeysResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadAccessKeysResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


