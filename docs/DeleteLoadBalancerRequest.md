# DeleteLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerName** | Pointer to **string** | The name of the load balancer you want to delete. | 

## Methods

### GetDryRun

`func (o *DeleteLoadBalancerRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteLoadBalancerRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteLoadBalancerRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteLoadBalancerRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetLoadBalancerName

`func (o *DeleteLoadBalancerRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *DeleteLoadBalancerRequest) GetLoadBalancerNameOk() (string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerName

`func (o *DeleteLoadBalancerRequest) HasLoadBalancerName() bool`

HasLoadBalancerName returns a boolean if a field has been set.

### SetLoadBalancerName

`func (o *DeleteLoadBalancerRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName gets a reference to the given string and assigns it to the LoadBalancerName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


