# RegisterVmsInLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendVmIds** | Pointer to **[]string** | One or more IDs of back-end VMs.&lt;br /&gt; Specifying the same ID several times has no effect as each back-end VM has equal weight. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerName** | Pointer to **string** | The name of the load balancer. | 

## Methods

### GetBackendVmIds

`func (o *RegisterVmsInLoadBalancerRequest) GetBackendVmIds() []string`

GetBackendVmIds returns the BackendVmIds field if non-nil, zero value otherwise.

### GetBackendVmIdsOk

`func (o *RegisterVmsInLoadBalancerRequest) GetBackendVmIdsOk() ([]string, bool)`

GetBackendVmIdsOk returns a tuple with the BackendVmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBackendVmIds

`func (o *RegisterVmsInLoadBalancerRequest) HasBackendVmIds() bool`

HasBackendVmIds returns a boolean if a field has been set.

### SetBackendVmIds

`func (o *RegisterVmsInLoadBalancerRequest) SetBackendVmIds(v []string)`

SetBackendVmIds gets a reference to the given []string and assigns it to the BackendVmIds field.

### GetDryRun

`func (o *RegisterVmsInLoadBalancerRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *RegisterVmsInLoadBalancerRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *RegisterVmsInLoadBalancerRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *RegisterVmsInLoadBalancerRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetLoadBalancerName

`func (o *RegisterVmsInLoadBalancerRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *RegisterVmsInLoadBalancerRequest) GetLoadBalancerNameOk() (string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerName

`func (o *RegisterVmsInLoadBalancerRequest) HasLoadBalancerName() bool`

HasLoadBalancerName returns a boolean if a field has been set.

### SetLoadBalancerName

`func (o *RegisterVmsInLoadBalancerRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName gets a reference to the given string and assigns it to the LoadBalancerName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


