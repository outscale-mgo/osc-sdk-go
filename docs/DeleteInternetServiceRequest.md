# DeleteInternetServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**InternetServiceId** | Pointer to **string** | The ID of the Internet service you want to delete. | 

## Methods

### GetDryRun

`func (o *DeleteInternetServiceRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteInternetServiceRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteInternetServiceRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteInternetServiceRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetInternetServiceId

`func (o *DeleteInternetServiceRequest) GetInternetServiceId() string`

GetInternetServiceId returns the InternetServiceId field if non-nil, zero value otherwise.

### GetInternetServiceIdOk

`func (o *DeleteInternetServiceRequest) GetInternetServiceIdOk() (string, bool)`

GetInternetServiceIdOk returns a tuple with the InternetServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInternetServiceId

`func (o *DeleteInternetServiceRequest) HasInternetServiceId() bool`

HasInternetServiceId returns a boolean if a field has been set.

### SetInternetServiceId

`func (o *DeleteInternetServiceRequest) SetInternetServiceId(v string)`

SetInternetServiceId gets a reference to the given string and assigns it to the InternetServiceId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


