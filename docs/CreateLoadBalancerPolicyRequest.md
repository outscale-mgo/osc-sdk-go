# CreateLoadBalancerPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CookieName** | Pointer to **string** | The name of the application cookie used for stickiness. This parameter is required if you create a stickiness policy based on an application-generated cookie. | [optional] 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**LoadBalancerName** | Pointer to **string** | The name of the load balancer for which you want to create a policy. | 
**PolicyName** | Pointer to **string** | The name of the policy. This name must be unique and consist of alphanumeric characters and dashes (-). | 
**PolicyType** | Pointer to **string** | The type of stickiness policy you want to create: &#x60;app&#x60; or &#x60;load_balancer&#x60;. | 

## Methods

### GetCookieName

`func (o *CreateLoadBalancerPolicyRequest) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *CreateLoadBalancerPolicyRequest) GetCookieNameOk() (string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCookieName

`func (o *CreateLoadBalancerPolicyRequest) HasCookieName() bool`

HasCookieName returns a boolean if a field has been set.

### SetCookieName

`func (o *CreateLoadBalancerPolicyRequest) SetCookieName(v string)`

SetCookieName gets a reference to the given string and assigns it to the CookieName field.

### GetDryRun

`func (o *CreateLoadBalancerPolicyRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateLoadBalancerPolicyRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateLoadBalancerPolicyRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateLoadBalancerPolicyRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetLoadBalancerName

`func (o *CreateLoadBalancerPolicyRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *CreateLoadBalancerPolicyRequest) GetLoadBalancerNameOk() (string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerName

`func (o *CreateLoadBalancerPolicyRequest) HasLoadBalancerName() bool`

HasLoadBalancerName returns a boolean if a field has been set.

### SetLoadBalancerName

`func (o *CreateLoadBalancerPolicyRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName gets a reference to the given string and assigns it to the LoadBalancerName field.

### GetPolicyName

`func (o *CreateLoadBalancerPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *CreateLoadBalancerPolicyRequest) GetPolicyNameOk() (string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPolicyName

`func (o *CreateLoadBalancerPolicyRequest) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyName

`func (o *CreateLoadBalancerPolicyRequest) SetPolicyName(v string)`

SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.

### GetPolicyType

`func (o *CreateLoadBalancerPolicyRequest) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *CreateLoadBalancerPolicyRequest) GetPolicyTypeOk() (string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPolicyType

`func (o *CreateLoadBalancerPolicyRequest) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### SetPolicyType

`func (o *CreateLoadBalancerPolicyRequest) SetPolicyType(v string)`

SetPolicyType gets a reference to the given string and assigns it to the PolicyType field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


