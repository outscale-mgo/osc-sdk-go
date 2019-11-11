# CreateLoadBalancerListenersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Listeners** | Pointer to [**[]ListenerForCreation**](ListenerForCreation.md) | One or more listeners for the load balancer. | 
**LoadBalancerName** | Pointer to **string** | The name of the load balancer for which you want to create listeners. | 

## Methods

### GetDryRun

`func (o *CreateLoadBalancerListenersRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateLoadBalancerListenersRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateLoadBalancerListenersRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateLoadBalancerListenersRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetListeners

`func (o *CreateLoadBalancerListenersRequest) GetListeners() []ListenerForCreation`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *CreateLoadBalancerListenersRequest) GetListenersOk() ([]ListenerForCreation, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasListeners

`func (o *CreateLoadBalancerListenersRequest) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### SetListeners

`func (o *CreateLoadBalancerListenersRequest) SetListeners(v []ListenerForCreation)`

SetListeners gets a reference to the given []ListenerForCreation and assigns it to the Listeners field.

### GetLoadBalancerName

`func (o *CreateLoadBalancerListenersRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *CreateLoadBalancerListenersRequest) GetLoadBalancerNameOk() (string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerName

`func (o *CreateLoadBalancerListenersRequest) HasLoadBalancerName() bool`

HasLoadBalancerName returns a boolean if a field has been set.

### SetLoadBalancerName

`func (o *CreateLoadBalancerListenersRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName gets a reference to the given string and assigns it to the LoadBalancerName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


