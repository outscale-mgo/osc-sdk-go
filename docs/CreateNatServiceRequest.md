# CreateNatServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**PublicIpId** | Pointer to **string** | The allocation ID of the EIP to associate with the NAT service.&lt;br /&gt; If the EIP is already associated with another resource, you must first disassociate it. | 
**SubnetId** | Pointer to **string** | The ID of the Subnet in which you want to create the NAT service. | 

## Methods

### GetDryRun

`func (o *CreateNatServiceRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateNatServiceRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateNatServiceRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateNatServiceRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetPublicIpId

`func (o *CreateNatServiceRequest) GetPublicIpId() string`

GetPublicIpId returns the PublicIpId field if non-nil, zero value otherwise.

### GetPublicIpIdOk

`func (o *CreateNatServiceRequest) GetPublicIpIdOk() (string, bool)`

GetPublicIpIdOk returns a tuple with the PublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicIpId

`func (o *CreateNatServiceRequest) HasPublicIpId() bool`

HasPublicIpId returns a boolean if a field has been set.

### SetPublicIpId

`func (o *CreateNatServiceRequest) SetPublicIpId(v string)`

SetPublicIpId gets a reference to the given string and assigns it to the PublicIpId field.

### GetSubnetId

`func (o *CreateNatServiceRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateNatServiceRequest) GetSubnetIdOk() (string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubnetId

`func (o *CreateNatServiceRequest) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetId

`func (o *CreateNatServiceRequest) SetSubnetId(v string)`

SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


