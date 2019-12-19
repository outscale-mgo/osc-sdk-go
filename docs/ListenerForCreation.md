# ListenerForCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendPort** | Pointer to **int64** | The port on which the back-end VM is listening (between &#x60;1&#x60; and &#x60;65535&#x60;, both included). | 
**BackendProtocol** | Pointer to **string** | The protocol for routing traffic to back-end VMs (&#x60;HTTP&#x60; \\| &#x60;HTTPS&#x60; \\| &#x60;TCP&#x60; \\| &#x60;SSL&#x60; \\| &#x60;UDP&#x60;). | [optional] 
**LoadBalancerPort** | Pointer to **int64** | The port on which the load balancer is listening (between &#x60;1&#x60; and &#x60;65535&#x60;, both included). | 
**LoadBalancerProtocol** | Pointer to **string** | The routing protocol (&#x60;HTTP&#x60; \\| &#x60;HTTPS&#x60; \\| &#x60;TCP&#x60; \\| &#x60;SSL&#x60; \\| &#x60;UDP&#x60;). | 
**ServerCertificateId** | Pointer to **string** | The ID of the server certificate. | [optional] 

## Methods

### GetBackendPort

`func (o *ListenerForCreation) GetBackendPort() int64`

GetBackendPort returns the BackendPort field if non-nil, zero value otherwise.

### GetBackendPortOk

`func (o *ListenerForCreation) GetBackendPortOk() (int64, bool)`

GetBackendPortOk returns a tuple with the BackendPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBackendPort

`func (o *ListenerForCreation) HasBackendPort() bool`

HasBackendPort returns a boolean if a field has been set.

### SetBackendPort

`func (o *ListenerForCreation) SetBackendPort(v int64)`

SetBackendPort gets a reference to the given int64 and assigns it to the BackendPort field.

### GetBackendProtocol

`func (o *ListenerForCreation) GetBackendProtocol() string`

GetBackendProtocol returns the BackendProtocol field if non-nil, zero value otherwise.

### GetBackendProtocolOk

`func (o *ListenerForCreation) GetBackendProtocolOk() (string, bool)`

GetBackendProtocolOk returns a tuple with the BackendProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBackendProtocol

`func (o *ListenerForCreation) HasBackendProtocol() bool`

HasBackendProtocol returns a boolean if a field has been set.

### SetBackendProtocol

`func (o *ListenerForCreation) SetBackendProtocol(v string)`

SetBackendProtocol gets a reference to the given string and assigns it to the BackendProtocol field.

### GetLoadBalancerPort

`func (o *ListenerForCreation) GetLoadBalancerPort() int64`

GetLoadBalancerPort returns the LoadBalancerPort field if non-nil, zero value otherwise.

### GetLoadBalancerPortOk

`func (o *ListenerForCreation) GetLoadBalancerPortOk() (int64, bool)`

GetLoadBalancerPortOk returns a tuple with the LoadBalancerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerPort

`func (o *ListenerForCreation) HasLoadBalancerPort() bool`

HasLoadBalancerPort returns a boolean if a field has been set.

### SetLoadBalancerPort

`func (o *ListenerForCreation) SetLoadBalancerPort(v int64)`

SetLoadBalancerPort gets a reference to the given int64 and assigns it to the LoadBalancerPort field.

### GetLoadBalancerProtocol

`func (o *ListenerForCreation) GetLoadBalancerProtocol() string`

GetLoadBalancerProtocol returns the LoadBalancerProtocol field if non-nil, zero value otherwise.

### GetLoadBalancerProtocolOk

`func (o *ListenerForCreation) GetLoadBalancerProtocolOk() (string, bool)`

GetLoadBalancerProtocolOk returns a tuple with the LoadBalancerProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerProtocol

`func (o *ListenerForCreation) HasLoadBalancerProtocol() bool`

HasLoadBalancerProtocol returns a boolean if a field has been set.

### SetLoadBalancerProtocol

`func (o *ListenerForCreation) SetLoadBalancerProtocol(v string)`

SetLoadBalancerProtocol gets a reference to the given string and assigns it to the LoadBalancerProtocol field.

### GetServerCertificateId

`func (o *ListenerForCreation) GetServerCertificateId() string`

GetServerCertificateId returns the ServerCertificateId field if non-nil, zero value otherwise.

### GetServerCertificateIdOk

`func (o *ListenerForCreation) GetServerCertificateIdOk() (string, bool)`

GetServerCertificateIdOk returns a tuple with the ServerCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServerCertificateId

`func (o *ListenerForCreation) HasServerCertificateId() bool`

HasServerCertificateId returns a boolean if a field has been set.

### SetServerCertificateId

`func (o *ListenerForCreation) SetServerCertificateId(v string)`

SetServerCertificateId gets a reference to the given string and assigns it to the ServerCertificateId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


