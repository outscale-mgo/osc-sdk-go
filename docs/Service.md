# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpRanges** | Pointer to **[]string** | The list of network prefixes used by the service, in CIDR notation. | [optional] 
**ServiceId** | Pointer to **string** | The ID of the service. | [optional] 
**ServiceName** | Pointer to **string** | The name of the prefix list, which identifies the 3DS OUTSCALE service it is associated with. | [optional] 

## Methods

### GetIpRanges

`func (o *Service) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *Service) GetIpRangesOk() ([]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIpRanges

`func (o *Service) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### SetIpRanges

`func (o *Service) SetIpRanges(v []string)`

SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.

### GetServiceId

`func (o *Service) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Service) GetServiceIdOk() (string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceId

`func (o *Service) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### SetServiceId

`func (o *Service) SetServiceId(v string)`

SetServiceId gets a reference to the given string and assigns it to the ServiceId field.

### GetServiceName

`func (o *Service) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Service) GetServiceNameOk() (string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServiceName

`func (o *Service) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceName

`func (o *Service) SetServiceName(v string)`

SetServiceName gets a reference to the given string and assigns it to the ServiceName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


