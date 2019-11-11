# LinkPublicIpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkPublicIpId** | Pointer to **string** | (Net only) The ID representing the association of the EIP with the VM or the NIC. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetLinkPublicIpId

`func (o *LinkPublicIpResponse) GetLinkPublicIpId() string`

GetLinkPublicIpId returns the LinkPublicIpId field if non-nil, zero value otherwise.

### GetLinkPublicIpIdOk

`func (o *LinkPublicIpResponse) GetLinkPublicIpIdOk() (string, bool)`

GetLinkPublicIpIdOk returns a tuple with the LinkPublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkPublicIpId

`func (o *LinkPublicIpResponse) HasLinkPublicIpId() bool`

HasLinkPublicIpId returns a boolean if a field has been set.

### SetLinkPublicIpId

`func (o *LinkPublicIpResponse) SetLinkPublicIpId(v string)`

SetLinkPublicIpId gets a reference to the given string and assigns it to the LinkPublicIpId field.

### GetResponseContext

`func (o *LinkPublicIpResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *LinkPublicIpResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *LinkPublicIpResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *LinkPublicIpResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


