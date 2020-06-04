# ReadLoadBalancerTagsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Tags** | Pointer to [**[]LoadBalancerTag**](LoadBalancerTag.md) | Information about one or more load balancer tags. | [optional] 

## Methods

### GetResponseContext

`func (o *ReadLoadBalancerTagsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadLoadBalancerTagsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadLoadBalancerTagsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadLoadBalancerTagsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.

### GetTags

`func (o *ReadLoadBalancerTagsResponse) GetTags() []LoadBalancerTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReadLoadBalancerTagsResponse) GetTagsOk() ([]LoadBalancerTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *ReadLoadBalancerTagsResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *ReadLoadBalancerTagsResponse) SetTags(v []LoadBalancerTag)`

SetTags gets a reference to the given []LoadBalancerTag and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


