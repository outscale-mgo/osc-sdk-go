# UpdateLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLog** | Pointer to [**AccessLog**](AccessLog.md) |  | [optional] 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**HealthCheck** | Pointer to [**HealthCheck**](HealthCheck.md) |  | [optional] 
**LoadBalancerName** | Pointer to **string** | The name of the load balancer. | 
**LoadBalancerPort** | Pointer to **int64** | The port on which the load balancer is listening (between &#x60;1&#x60; and &#x60;65535&#x60;, both included). | [optional] 
**PolicyNames** | Pointer to **[]string** | The list of policy names (must contain all the policies to be enabled). | [optional] 
**ServerCertificateId** | Pointer to **string** | The Outscale Resource Name (ORN) of the SSL certificate. | [optional] 

## Methods

### GetAccessLog

`func (o *UpdateLoadBalancerRequest) GetAccessLog() AccessLog`

GetAccessLog returns the AccessLog field if non-nil, zero value otherwise.

### GetAccessLogOk

`func (o *UpdateLoadBalancerRequest) GetAccessLogOk() (AccessLog, bool)`

GetAccessLogOk returns a tuple with the AccessLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccessLog

`func (o *UpdateLoadBalancerRequest) HasAccessLog() bool`

HasAccessLog returns a boolean if a field has been set.

### SetAccessLog

`func (o *UpdateLoadBalancerRequest) SetAccessLog(v AccessLog)`

SetAccessLog gets a reference to the given AccessLog and assigns it to the AccessLog field.

### GetDryRun

`func (o *UpdateLoadBalancerRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateLoadBalancerRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UpdateLoadBalancerRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UpdateLoadBalancerRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetHealthCheck

`func (o *UpdateLoadBalancerRequest) GetHealthCheck() HealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *UpdateLoadBalancerRequest) GetHealthCheckOk() (HealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHealthCheck

`func (o *UpdateLoadBalancerRequest) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### SetHealthCheck

`func (o *UpdateLoadBalancerRequest) SetHealthCheck(v HealthCheck)`

SetHealthCheck gets a reference to the given HealthCheck and assigns it to the HealthCheck field.

### GetLoadBalancerName

`func (o *UpdateLoadBalancerRequest) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *UpdateLoadBalancerRequest) GetLoadBalancerNameOk() (string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerName

`func (o *UpdateLoadBalancerRequest) HasLoadBalancerName() bool`

HasLoadBalancerName returns a boolean if a field has been set.

### SetLoadBalancerName

`func (o *UpdateLoadBalancerRequest) SetLoadBalancerName(v string)`

SetLoadBalancerName gets a reference to the given string and assigns it to the LoadBalancerName field.

### GetLoadBalancerPort

`func (o *UpdateLoadBalancerRequest) GetLoadBalancerPort() int64`

GetLoadBalancerPort returns the LoadBalancerPort field if non-nil, zero value otherwise.

### GetLoadBalancerPortOk

`func (o *UpdateLoadBalancerRequest) GetLoadBalancerPortOk() (int64, bool)`

GetLoadBalancerPortOk returns a tuple with the LoadBalancerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerPort

`func (o *UpdateLoadBalancerRequest) HasLoadBalancerPort() bool`

HasLoadBalancerPort returns a boolean if a field has been set.

### SetLoadBalancerPort

`func (o *UpdateLoadBalancerRequest) SetLoadBalancerPort(v int64)`

SetLoadBalancerPort gets a reference to the given int64 and assigns it to the LoadBalancerPort field.

### GetPolicyNames

`func (o *UpdateLoadBalancerRequest) GetPolicyNames() []string`

GetPolicyNames returns the PolicyNames field if non-nil, zero value otherwise.

### GetPolicyNamesOk

`func (o *UpdateLoadBalancerRequest) GetPolicyNamesOk() ([]string, bool)`

GetPolicyNamesOk returns a tuple with the PolicyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPolicyNames

`func (o *UpdateLoadBalancerRequest) HasPolicyNames() bool`

HasPolicyNames returns a boolean if a field has been set.

### SetPolicyNames

`func (o *UpdateLoadBalancerRequest) SetPolicyNames(v []string)`

SetPolicyNames gets a reference to the given []string and assigns it to the PolicyNames field.

### GetServerCertificateId

`func (o *UpdateLoadBalancerRequest) GetServerCertificateId() string`

GetServerCertificateId returns the ServerCertificateId field if non-nil, zero value otherwise.

### GetServerCertificateIdOk

`func (o *UpdateLoadBalancerRequest) GetServerCertificateIdOk() (string, bool)`

GetServerCertificateIdOk returns a tuple with the ServerCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServerCertificateId

`func (o *UpdateLoadBalancerRequest) HasServerCertificateId() bool`

HasServerCertificateId returns a boolean if a field has been set.

### SetServerCertificateId

`func (o *UpdateLoadBalancerRequest) SetServerCertificateId(v string)`

SetServerCertificateId gets a reference to the given string and assigns it to the ServerCertificateId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


