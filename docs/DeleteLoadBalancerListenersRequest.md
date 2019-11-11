# DeleteLoadBalancerListenersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerName** | Pointer to **string** | The name of the load balancer for which you want to delete listeners. | 
**LoadBalancerPorts** | Pointer to **[]int32** | One or more port numbers of the listeners you want to delete. | 

## Methods

### GetDryRun

`func (o *DeleteLoadBalancerListenersRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteLoadBalancerListenersRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteLoadBalancerListenersRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteLoadBalancerListenersRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetLoadBalancerName

`func (o *DeleteLoadBalancerListenersRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *DeleteLoadBalancerListenersRequest) GetLoadBalancerNameOk() (string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerName

`func (o *DeleteLoadBalancerListenersRequest) HasLoadBalancerName() bool`

HasLoadBalancerName returns a boolean if a field has been set.

### SetLoadBalancerName

`func (o *DeleteLoadBalancerListenersRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName gets a reference to the given string and assigns it to the LoadBalancerName field.

### GetLoadBalancerPorts

`func (o *DeleteLoadBalancerListenersRequest) GetLoadBalancerPorts() []int32`

GetLoadBalancerPorts returns the LoadBalancerPorts field if non-nil, zero value otherwise.

### GetLoadBalancerPortsOk

`func (o *DeleteLoadBalancerListenersRequest) GetLoadBalancerPortsOk() ([]int32, bool)`

GetLoadBalancerPortsOk returns a tuple with the LoadBalancerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerPorts

`func (o *DeleteLoadBalancerListenersRequest) HasLoadBalancerPorts() bool`

HasLoadBalancerPorts returns a boolean if a field has been set.

### SetLoadBalancerPorts

`func (o *DeleteLoadBalancerListenersRequest) SetLoadBalancerPorts(v []int32)`

SetLoadBalancerPorts gets a reference to the given []int32 and assigns it to the LoadBalancerPorts field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


