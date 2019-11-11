# DeletePublicIpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**PublicIp** | Pointer to **string** | The EIP. In the public Cloud, this parameter is required. | [optional] 
**PublicIpId** | Pointer to **string** | The ID representing the association of the EIP with the VM or the NIC. In a Net, this parameter is required. | [optional] 

## Methods

### GetDryRun

`func (o *DeletePublicIpRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeletePublicIpRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeletePublicIpRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeletePublicIpRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetPublicIp

`func (o *DeletePublicIpRequest) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *DeletePublicIpRequest) GetPublicIpOk() (string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicIp

`func (o *DeletePublicIpRequest) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### SetPublicIp

`func (o *DeletePublicIpRequest) SetPublicIp(v string)`

SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.

### GetPublicIpId

`func (o *DeletePublicIpRequest) GetPublicIpId() string`

GetPublicIpId returns the PublicIpId field if non-nil, zero value otherwise.

### GetPublicIpIdOk

`func (o *DeletePublicIpRequest) GetPublicIpIdOk() (string, bool)`

GetPublicIpIdOk returns a tuple with the PublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicIpId

`func (o *DeletePublicIpRequest) HasPublicIpId() bool`

HasPublicIpId returns a boolean if a field has been set.

### SetPublicIpId

`func (o *DeletePublicIpRequest) SetPublicIpId(v string)`

SetPublicIpId gets a reference to the given string and assigns it to the PublicIpId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


