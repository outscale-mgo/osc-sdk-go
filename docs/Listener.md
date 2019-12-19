# Listener

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendPort** | Pointer to **int64** | The port on which the back-end VM is listening (between &#x60;1&#x60; and &#x60;65535&#x60;, both included). | [optional] 
**BackendProtocol** | Pointer to **string** | The protocol for routing traffic to back-end VMs (&#x60;HTTP&#x60; \\| &#x60;HTTPS&#x60; \\| &#x60;TCP&#x60; \\| &#x60;SSL&#x60; \\| &#x60;UDP&#x60;). | [optional] 
**LoadBalancerPort** | Pointer to **int64** | The port on which the load balancer is listening (between 1 and &#x60;65535&#x60;, both included). | [optional] 
**LoadBalancerProtocol** | Pointer to **string** | The routing protocol (&#x60;HTTP&#x60; \\| &#x60;HTTPS&#x60; \\| &#x60;TCP&#x60; \\| &#x60;SSL&#x60; \\| &#x60;UDP&#x60;). | [optional] 
**PolicyNames** | Pointer to **[]string** | The names of the policies. If there are no policies enabled, the list is empty. | [optional] 
**ServerCertificateId** | Pointer to **string** | The ID of the server certificate. | [optional] 

## Methods

### GetBackendPort

`func (o *Listener) GetBackendPort() int64`

GetBackendPort returns the BackendPort field if non-nil, zero value otherwise.

### GetBackendPortOk

`func (o *Listener) GetBackendPortOk() (int64, bool)`

GetBackendPortOk returns a tuple with the BackendPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBackendPort

`func (o *Listener) HasBackendPort() bool`

HasBackendPort returns a boolean if a field has been set.

### SetBackendPort

`func (o *Listener) SetBackendPort(v int64)`

SetBackendPort gets a reference to the given int64 and assigns it to the BackendPort field.

### GetBackendProtocol

`func (o *Listener) GetBackendProtocol() string`

GetBackendProtocol returns the BackendProtocol field if non-nil, zero value otherwise.

### GetBackendProtocolOk

`func (o *Listener) GetBackendProtocolOk() (string, bool)`

GetBackendProtocolOk returns a tuple with the BackendProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBackendProtocol

`func (o *Listener) HasBackendProtocol() bool`

HasBackendProtocol returns a boolean if a field has been set.

### SetBackendProtocol

`func (o *Listener) SetBackendProtocol(v string)`

SetBackendProtocol gets a reference to the given string and assigns it to the BackendProtocol field.

### GetLoadBalancerPort

`func (o *Listener) GetLoadBalancerPort() int64`

GetLoadBalancerPort returns the LoadBalancerPort field if non-nil, zero value otherwise.

### GetLoadBalancerPortOk

`func (o *Listener) GetLoadBalancerPortOk() (int64, bool)`

GetLoadBalancerPortOk returns a tuple with the LoadBalancerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerPort

`func (o *Listener) HasLoadBalancerPort() bool`

HasLoadBalancerPort returns a boolean if a field has been set.

### SetLoadBalancerPort

`func (o *Listener) SetLoadBalancerPort(v int64)`

SetLoadBalancerPort gets a reference to the given int64 and assigns it to the LoadBalancerPort field.

### GetLoadBalancerProtocol

`func (o *Listener) GetLoadBalancerProtocol() string`

GetLoadBalancerProtocol returns the LoadBalancerProtocol field if non-nil, zero value otherwise.

### GetLoadBalancerProtocolOk

`func (o *Listener) GetLoadBalancerProtocolOk() (string, bool)`

GetLoadBalancerProtocolOk returns a tuple with the LoadBalancerProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerProtocol

`func (o *Listener) HasLoadBalancerProtocol() bool`

HasLoadBalancerProtocol returns a boolean if a field has been set.

### SetLoadBalancerProtocol

`func (o *Listener) SetLoadBalancerProtocol(v string)`

SetLoadBalancerProtocol gets a reference to the given string and assigns it to the LoadBalancerProtocol field.

### GetPolicyNames

`func (o *Listener) GetPolicyNames() []string`

GetPolicyNames returns the PolicyNames field if non-nil, zero value otherwise.

### GetPolicyNamesOk

`func (o *Listener) GetPolicyNamesOk() ([]string, bool)`

GetPolicyNamesOk returns a tuple with the PolicyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPolicyNames

`func (o *Listener) HasPolicyNames() bool`

HasPolicyNames returns a boolean if a field has been set.

### SetPolicyNames

`func (o *Listener) SetPolicyNames(v []string)`

SetPolicyNames gets a reference to the given []string and assigns it to the PolicyNames field.

### GetServerCertificateId

`func (o *Listener) GetServerCertificateId() string`

GetServerCertificateId returns the ServerCertificateId field if non-nil, zero value otherwise.

### GetServerCertificateIdOk

`func (o *Listener) GetServerCertificateIdOk() (string, bool)`

GetServerCertificateIdOk returns a tuple with the ServerCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServerCertificateId

`func (o *Listener) HasServerCertificateId() bool`

HasServerCertificateId returns a boolean if a field has been set.

### SetServerCertificateId

`func (o *Listener) SetServerCertificateId(v string)`

SetServerCertificateId gets a reference to the given string and assigns it to the ServerCertificateId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


