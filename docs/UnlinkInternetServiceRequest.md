# UnlinkInternetServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**InternetServiceId** | Pointer to **string** | The ID of the Internet service you want to detach. | 
**NetId** | Pointer to **string** | The ID of the Net from which you want to detach the Internet service. | 

## Methods

### GetDryRun

`func (o *UnlinkInternetServiceRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UnlinkInternetServiceRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UnlinkInternetServiceRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UnlinkInternetServiceRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetInternetServiceId

`func (o *UnlinkInternetServiceRequest) GetInternetServiceId() string`

GetInternetServiceId returns the InternetServiceId field if non-nil, zero value otherwise.

### GetInternetServiceIdOk

`func (o *UnlinkInternetServiceRequest) GetInternetServiceIdOk() (string, bool)`

GetInternetServiceIdOk returns a tuple with the InternetServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInternetServiceId

`func (o *UnlinkInternetServiceRequest) HasInternetServiceId() bool`

HasInternetServiceId returns a boolean if a field has been set.

### SetInternetServiceId

`func (o *UnlinkInternetServiceRequest) SetInternetServiceId(v string)`

SetInternetServiceId gets a reference to the given string and assigns it to the InternetServiceId field.

### GetNetId

`func (o *UnlinkInternetServiceRequest) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *UnlinkInternetServiceRequest) GetNetIdOk() (string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetId

`func (o *UnlinkInternetServiceRequest) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### SetNetId

`func (o *UnlinkInternetServiceRequest) SetNetId(v string)`

SetNetId gets a reference to the given string and assigns it to the NetId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


