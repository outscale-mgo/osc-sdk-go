# DeleteNatServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**NatServiceId** | Pointer to **string** | The ID of the NAT service you want to delete. | 

## Methods

### GetDryRun

`func (o *DeleteNatServiceRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteNatServiceRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteNatServiceRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteNatServiceRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetNatServiceId

`func (o *DeleteNatServiceRequest) GetNatServiceId() string`

GetNatServiceId returns the NatServiceId field if non-nil, zero value otherwise.

### GetNatServiceIdOk

`func (o *DeleteNatServiceRequest) GetNatServiceIdOk() (string, bool)`

GetNatServiceIdOk returns a tuple with the NatServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNatServiceId

`func (o *DeleteNatServiceRequest) HasNatServiceId() bool`

HasNatServiceId returns a boolean if a field has been set.

### SetNatServiceId

`func (o *DeleteNatServiceRequest) SetNatServiceId(v string)`

SetNatServiceId gets a reference to the given string and assigns it to the NatServiceId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


