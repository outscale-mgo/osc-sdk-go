# DeregisterVmsInLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendVmIds** | Pointer to **[]string** | One or more IDs of back-end VMs. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerName** | Pointer to **string** | The name of the load balancer. | 

## Methods

### GetBackendVmIds

`func (o *DeregisterVmsInLoadBalancerRequest) GetBackendVmIds() []string`

GetBackendVmIds returns the BackendVmIds field if non-nil, zero value otherwise.

### GetBackendVmIdsOk

`func (o *DeregisterVmsInLoadBalancerRequest) GetBackendVmIdsOk() ([]string, bool)`

GetBackendVmIdsOk returns a tuple with the BackendVmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBackendVmIds

`func (o *DeregisterVmsInLoadBalancerRequest) HasBackendVmIds() bool`

HasBackendVmIds returns a boolean if a field has been set.

### SetBackendVmIds

`func (o *DeregisterVmsInLoadBalancerRequest) SetBackendVmIds(v []string)`

SetBackendVmIds gets a reference to the given []string and assigns it to the BackendVmIds field.

### GetDryRun

`func (o *DeregisterVmsInLoadBalancerRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeregisterVmsInLoadBalancerRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeregisterVmsInLoadBalancerRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeregisterVmsInLoadBalancerRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetLoadBalancerName

`func (o *DeregisterVmsInLoadBalancerRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *DeregisterVmsInLoadBalancerRequest) GetLoadBalancerNameOk() (string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerName

`func (o *DeregisterVmsInLoadBalancerRequest) HasLoadBalancerName() bool`

HasLoadBalancerName returns a boolean if a field has been set.

### SetLoadBalancerName

`func (o *DeregisterVmsInLoadBalancerRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName gets a reference to the given string and assigns it to the LoadBalancerName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


