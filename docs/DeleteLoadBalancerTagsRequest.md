# DeleteLoadBalancerTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerNames** | Pointer to **[]string** | One or more load balancer names. | 
**Tags** | Pointer to [**[]ResourceLoadBalancerTag**](ResourceLoadBalancerTag.md) | One or more tags to delete from the load balancers. | 

## Methods

### GetDryRun

`func (o *DeleteLoadBalancerTagsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteLoadBalancerTagsRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteLoadBalancerTagsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteLoadBalancerTagsRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetLoadBalancerNames

`func (o *DeleteLoadBalancerTagsRequest) GetLoadBalancerNames() []string`

GetLoadBalancerNames returns the LoadBalancerNames field if non-nil, zero value otherwise.

### GetLoadBalancerNamesOk

`func (o *DeleteLoadBalancerTagsRequest) GetLoadBalancerNamesOk() ([]string, bool)`

GetLoadBalancerNamesOk returns a tuple with the LoadBalancerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerNames

`func (o *DeleteLoadBalancerTagsRequest) HasLoadBalancerNames() bool`

HasLoadBalancerNames returns a boolean if a field has been set.

### SetLoadBalancerNames

`func (o *DeleteLoadBalancerTagsRequest) SetLoadBalancerNames(v []string)`

SetLoadBalancerNames gets a reference to the given []string and assigns it to the LoadBalancerNames field.

### GetTags

`func (o *DeleteLoadBalancerTagsRequest) GetTags() []ResourceLoadBalancerTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeleteLoadBalancerTagsRequest) GetTagsOk() ([]ResourceLoadBalancerTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *DeleteLoadBalancerTagsRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *DeleteLoadBalancerTagsRequest) SetTags(v []ResourceLoadBalancerTag)`

SetTags gets a reference to the given []ResourceLoadBalancerTag and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


