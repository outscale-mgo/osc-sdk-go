# ReadNetAccessPointServicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Services** | Pointer to [**[]Service**](Service.md) | The names of the services you can use for Net access points. | [optional] 

## Methods

### GetResponseContext

`func (o *ReadNetAccessPointServicesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadNetAccessPointServicesResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadNetAccessPointServicesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadNetAccessPointServicesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.

### GetServices

`func (o *ReadNetAccessPointServicesResponse) GetServices() []Service`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ReadNetAccessPointServicesResponse) GetServicesOk() ([]Service, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServices

`func (o *ReadNetAccessPointServicesResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServices

`func (o *ReadNetAccessPointServicesResponse) SetServices(v []Service)`

SetServices gets a reference to the given []Service and assigns it to the Services field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


