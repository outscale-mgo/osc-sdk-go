# DeleteAccessKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **string** | The ID of the access key you want to delete. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 

## Methods

### GetAccessKeyId

`func (o *DeleteAccessKeyRequest) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *DeleteAccessKeyRequest) GetAccessKeyIdOk() (string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccessKeyId

`func (o *DeleteAccessKeyRequest) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### SetAccessKeyId

`func (o *DeleteAccessKeyRequest) SetAccessKeyId(v string)`

SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.

### GetDryRun

`func (o *DeleteAccessKeyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteAccessKeyRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteAccessKeyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteAccessKeyRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


