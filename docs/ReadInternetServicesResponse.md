# ReadInternetServicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternetServices** | Pointer to [**[]InternetService**](InternetService.md) | Information about one or more Internet services. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetInternetServices

`func (o *ReadInternetServicesResponse) GetInternetServices() []InternetService`

GetInternetServices returns the InternetServices field if non-nil, zero value otherwise.

### GetInternetServicesOk

`func (o *ReadInternetServicesResponse) GetInternetServicesOk() ([]InternetService, bool)`

GetInternetServicesOk returns a tuple with the InternetServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInternetServices

`func (o *ReadInternetServicesResponse) HasInternetServices() bool`

HasInternetServices returns a boolean if a field has been set.

### SetInternetServices

`func (o *ReadInternetServicesResponse) SetInternetServices(v []InternetService)`

SetInternetServices gets a reference to the given []InternetService and assigns it to the InternetServices field.

### GetResponseContext

`func (o *ReadInternetServicesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadInternetServicesResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadInternetServicesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadInternetServicesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


