# UnlinkPublicIpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**LinkPublicIpId** | Pointer to **string** | (Required in a Net) The ID representing the association of the EIP with the VM or the NIC. | [optional] 
**PublicIp** | Pointer to **string** | The External IP address. In the public Cloud, this parameter is required. | [optional] 

## Methods

### GetDryRun

`func (o *UnlinkPublicIpRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UnlinkPublicIpRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UnlinkPublicIpRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UnlinkPublicIpRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetLinkPublicIpId

`func (o *UnlinkPublicIpRequest) GetLinkPublicIpId() string`

GetLinkPublicIpId returns the LinkPublicIpId field if non-nil, zero value otherwise.

### GetLinkPublicIpIdOk

`func (o *UnlinkPublicIpRequest) GetLinkPublicIpIdOk() (string, bool)`

GetLinkPublicIpIdOk returns a tuple with the LinkPublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkPublicIpId

`func (o *UnlinkPublicIpRequest) HasLinkPublicIpId() bool`

HasLinkPublicIpId returns a boolean if a field has been set.

### SetLinkPublicIpId

`func (o *UnlinkPublicIpRequest) SetLinkPublicIpId(v string)`

SetLinkPublicIpId gets a reference to the given string and assigns it to the LinkPublicIpId field.

### GetPublicIp

`func (o *UnlinkPublicIpRequest) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *UnlinkPublicIpRequest) GetPublicIpOk() (string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicIp

`func (o *UnlinkPublicIpRequest) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### SetPublicIp

`func (o *UnlinkPublicIpRequest) SetPublicIp(v string)`

SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


