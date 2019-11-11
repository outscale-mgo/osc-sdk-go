# ReadNatServicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NatServices** | Pointer to [**[]NatService**](NatService.md) | Information about one or more NAT services. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetNatServices

`func (o *ReadNatServicesResponse) GetNatServices() []NatService`

GetNatServices returns the NatServices field if non-nil, zero value otherwise.

### GetNatServicesOk

`func (o *ReadNatServicesResponse) GetNatServicesOk() ([]NatService, bool)`

GetNatServicesOk returns a tuple with the NatServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNatServices

`func (o *ReadNatServicesResponse) HasNatServices() bool`

HasNatServices returns a boolean if a field has been set.

### SetNatServices

`func (o *ReadNatServicesResponse) SetNatServices(v []NatService)`

SetNatServices gets a reference to the given []NatService and assigns it to the NatServices field.

### GetResponseContext

`func (o *ReadNatServicesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadNatServicesResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadNatServicesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadNatServicesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


