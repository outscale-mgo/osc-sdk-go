# UpdateNetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpOptionsSetId** | Pointer to **string** | The ID of the DHCP options set (or &#x60;default&#x60; if you want to associate the default one). | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net. | 

## Methods

### GetDhcpOptionsSetId

`func (o *UpdateNetRequest) GetDhcpOptionsSetId() string`

GetDhcpOptionsSetId returns the DhcpOptionsSetId field if non-nil, zero value otherwise.

### GetDhcpOptionsSetIdOk

`func (o *UpdateNetRequest) GetDhcpOptionsSetIdOk() (string, bool)`

GetDhcpOptionsSetIdOk returns a tuple with the DhcpOptionsSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDhcpOptionsSetId

`func (o *UpdateNetRequest) HasDhcpOptionsSetId() bool`

HasDhcpOptionsSetId returns a boolean if a field has been set.

### SetDhcpOptionsSetId

`func (o *UpdateNetRequest) SetDhcpOptionsSetId(v string)`

SetDhcpOptionsSetId gets a reference to the given string and assigns it to the DhcpOptionsSetId field.

### GetDryRun

`func (o *UpdateNetRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateNetRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UpdateNetRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UpdateNetRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetNetId

`func (o *UpdateNetRequest) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *UpdateNetRequest) GetNetIdOk() (string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetId

`func (o *UpdateNetRequest) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### SetNetId

`func (o *UpdateNetRequest) SetNetId(v string)`

SetNetId gets a reference to the given string and assigns it to the NetId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


