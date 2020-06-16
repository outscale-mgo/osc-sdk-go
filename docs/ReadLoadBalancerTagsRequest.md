# ReadLoadBalancerTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerNames** | Pointer to **[]string** | One or more load balancer names. | 

## Methods

### GetDryRun

`func (o *ReadLoadBalancerTagsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadLoadBalancerTagsRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *ReadLoadBalancerTagsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *ReadLoadBalancerTagsRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetLoadBalancerNames

`func (o *ReadLoadBalancerTagsRequest) GetLoadBalancerNames() []string`

GetLoadBalancerNames returns the LoadBalancerNames field if non-nil, zero value otherwise.

### GetLoadBalancerNamesOk

`func (o *ReadLoadBalancerTagsRequest) GetLoadBalancerNamesOk() ([]string, bool)`

GetLoadBalancerNamesOk returns a tuple with the LoadBalancerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerNames

`func (o *ReadLoadBalancerTagsRequest) HasLoadBalancerNames() bool`

HasLoadBalancerNames returns a boolean if a field has been set.

### SetLoadBalancerNames

`func (o *ReadLoadBalancerTagsRequest) SetLoadBalancerNames(v []string)`

SetLoadBalancerNames gets a reference to the given []string and assigns it to the LoadBalancerNames field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


