# UpdateSubnetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**MapPublicIpOnLaunch** | Pointer to **bool** | If &#x60;true&#x60;, a public IP address is assigned to the network interface cards (NICs) created in the specified Subnet. | 
**SubnetId** | Pointer to **string** | The ID of the Subnet. | 

## Methods

### GetDryRun

`func (o *UpdateSubnetRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateSubnetRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UpdateSubnetRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UpdateSubnetRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetMapPublicIpOnLaunch

`func (o *UpdateSubnetRequest) GetMapPublicIpOnLaunch() bool`

GetMapPublicIpOnLaunch returns the MapPublicIpOnLaunch field if non-nil, zero value otherwise.

### GetMapPublicIpOnLaunchOk

`func (o *UpdateSubnetRequest) GetMapPublicIpOnLaunchOk() (bool, bool)`

GetMapPublicIpOnLaunchOk returns a tuple with the MapPublicIpOnLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMapPublicIpOnLaunch

`func (o *UpdateSubnetRequest) HasMapPublicIpOnLaunch() bool`

HasMapPublicIpOnLaunch returns a boolean if a field has been set.

### SetMapPublicIpOnLaunch

`func (o *UpdateSubnetRequest) SetMapPublicIpOnLaunch(v bool)`

SetMapPublicIpOnLaunch gets a reference to the given bool and assigns it to the MapPublicIpOnLaunch field.

### GetSubnetId

`func (o *UpdateSubnetRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *UpdateSubnetRequest) GetSubnetIdOk() (string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubnetId

`func (o *UpdateSubnetRequest) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetId

`func (o *UpdateSubnetRequest) SetSubnetId(v string)`

SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


