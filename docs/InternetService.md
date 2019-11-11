# InternetService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternetServiceId** | Pointer to **string** | The ID of the Internet service. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net attached to the Internet service. | [optional] 
**State** | Pointer to **string** | The state of the attachment of the Net to the Internet service (always &#x60;available&#x60;). | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the Internet service. | [optional] 

## Methods

### GetInternetServiceId

`func (o *InternetService) GetInternetServiceId() string`

GetInternetServiceId returns the InternetServiceId field if non-nil, zero value otherwise.

### GetInternetServiceIdOk

`func (o *InternetService) GetInternetServiceIdOk() (string, bool)`

GetInternetServiceIdOk returns a tuple with the InternetServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInternetServiceId

`func (o *InternetService) HasInternetServiceId() bool`

HasInternetServiceId returns a boolean if a field has been set.

### SetInternetServiceId

`func (o *InternetService) SetInternetServiceId(v string)`

SetInternetServiceId gets a reference to the given string and assigns it to the InternetServiceId field.

### GetNetId

`func (o *InternetService) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *InternetService) GetNetIdOk() (string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetId

`func (o *InternetService) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### SetNetId

`func (o *InternetService) SetNetId(v string)`

SetNetId gets a reference to the given string and assigns it to the NetId field.

### GetState

`func (o *InternetService) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InternetService) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *InternetService) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *InternetService) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetTags

`func (o *InternetService) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InternetService) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *InternetService) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *InternetService) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


