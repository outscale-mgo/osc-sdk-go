# CreateNetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**IpRange** | Pointer to **string** | The IP range for the Net, in CIDR notation (for example, 10.0.0.0/16). | 
**Tenancy** | Pointer to **string** | The tenancy options for the VMs (&#x60;default&#x60; if a VM created in a Net can be launched with any tenancy, &#x60;dedicated&#x60; if it can be launched with dedicated tenancy VMs running on single-tenant hardware). | [optional] 

## Methods

### GetDryRun

`func (o *CreateNetRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateNetRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateNetRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateNetRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetIpRange

`func (o *CreateNetRequest) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *CreateNetRequest) GetIpRangeOk() (string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIpRange

`func (o *CreateNetRequest) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### SetIpRange

`func (o *CreateNetRequest) SetIpRange(v string)`

SetIpRange gets a reference to the given string and assigns it to the IpRange field.

### GetTenancy

`func (o *CreateNetRequest) GetTenancy() string`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *CreateNetRequest) GetTenancyOk() (string, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTenancy

`func (o *CreateNetRequest) HasTenancy() bool`

HasTenancy returns a boolean if a field has been set.

### SetTenancy

`func (o *CreateNetRequest) SetTenancy(v string)`

SetTenancy gets a reference to the given string and assigns it to the Tenancy field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


