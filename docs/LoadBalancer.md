# LoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLog** | Pointer to [**AccessLog**](AccessLog.md) |  | [optional] 
**ApplicationStickyCookiePolicies** | Pointer to [**[]ApplicationStickyCookiePolicy**](ApplicationStickyCookiePolicy.md) | The stickiness policies defined for the load balancer. | [optional] 
**BackendVmIds** | Pointer to **[]string** | One or more IDs of back-end VMs for the load balancer. | [optional] 
**DnsName** | Pointer to **string** | The DNS name of the load balancer. | [optional] 
**HealthCheck** | Pointer to [**HealthCheck**](HealthCheck.md) |  | [optional] 
**Listeners** | Pointer to [**[]Listener**](Listener.md) | The listeners for the load balancer. | [optional] 
**LoadBalancerName** | Pointer to **string** | The name of the load balancer. | [optional] 
**LoadBalancerStickyCookiePolicies** | Pointer to [**[]LoadBalancerStickyCookiePolicy**](LoadBalancerStickyCookiePolicy.md) | The policies defined for the load balancer. | [optional] 
**LoadBalancerType** | Pointer to **string** | The type of load balancer. Valid only for load balancers in a Net.&lt;br /&gt; If &#x60;LoadBalancerType&#x60; is &#x60;internet-facing&#x60;, the load balancer has a public DNS name that resolves to a public IP address.&lt;br /&gt; If &#x60;LoadBalancerType&#x60; is &#x60;internal&#x60;, the load balancer has a public DNS name that resolves to a private IP address. | [optional] 
**NetId** | Pointer to **string** | The ID of the Net for the load balancer. | [optional] 
**SecurityGroups** | Pointer to **[]string** | One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net. | [optional] 
**SourceSecurityGroup** | Pointer to [**SourceSecurityGroup**](SourceSecurityGroup.md) |  | [optional] 
**Subnets** | Pointer to **[]string** | The IDs of the Subnets for the load balancer. | [optional] 
**SubregionNames** | Pointer to **[]string** | One or more names of Subregions for the load balancer. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the load balancer. | [optional] 

## Methods

### GetAccessLog

`func (o *LoadBalancer) GetAccessLog() AccessLog`

GetAccessLog returns the AccessLog field if non-nil, zero value otherwise.

### GetAccessLogOk

`func (o *LoadBalancer) GetAccessLogOk() (AccessLog, bool)`

GetAccessLogOk returns a tuple with the AccessLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccessLog

`func (o *LoadBalancer) HasAccessLog() bool`

HasAccessLog returns a boolean if a field has been set.

### SetAccessLog

`func (o *LoadBalancer) SetAccessLog(v AccessLog)`

SetAccessLog gets a reference to the given AccessLog and assigns it to the AccessLog field.

### GetApplicationStickyCookiePolicies

`func (o *LoadBalancer) GetApplicationStickyCookiePolicies() []ApplicationStickyCookiePolicy`

GetApplicationStickyCookiePolicies returns the ApplicationStickyCookiePolicies field if non-nil, zero value otherwise.

### GetApplicationStickyCookiePoliciesOk

`func (o *LoadBalancer) GetApplicationStickyCookiePoliciesOk() ([]ApplicationStickyCookiePolicy, bool)`

GetApplicationStickyCookiePoliciesOk returns a tuple with the ApplicationStickyCookiePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationStickyCookiePolicies

`func (o *LoadBalancer) HasApplicationStickyCookiePolicies() bool`

HasApplicationStickyCookiePolicies returns a boolean if a field has been set.

### SetApplicationStickyCookiePolicies

`func (o *LoadBalancer) SetApplicationStickyCookiePolicies(v []ApplicationStickyCookiePolicy)`

SetApplicationStickyCookiePolicies gets a reference to the given []ApplicationStickyCookiePolicy and assigns it to the ApplicationStickyCookiePolicies field.

### GetBackendVmIds

`func (o *LoadBalancer) GetBackendVmIds() []string`

GetBackendVmIds returns the BackendVmIds field if non-nil, zero value otherwise.

### GetBackendVmIdsOk

`func (o *LoadBalancer) GetBackendVmIdsOk() ([]string, bool)`

GetBackendVmIdsOk returns a tuple with the BackendVmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBackendVmIds

`func (o *LoadBalancer) HasBackendVmIds() bool`

HasBackendVmIds returns a boolean if a field has been set.

### SetBackendVmIds

`func (o *LoadBalancer) SetBackendVmIds(v []string)`

SetBackendVmIds gets a reference to the given []string and assigns it to the BackendVmIds field.

### GetDnsName

`func (o *LoadBalancer) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *LoadBalancer) GetDnsNameOk() (string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDnsName

`func (o *LoadBalancer) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### SetDnsName

`func (o *LoadBalancer) SetDnsName(v string)`

SetDnsName gets a reference to the given string and assigns it to the DnsName field.

### GetHealthCheck

`func (o *LoadBalancer) GetHealthCheck() HealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *LoadBalancer) GetHealthCheckOk() (HealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHealthCheck

`func (o *LoadBalancer) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### SetHealthCheck

`func (o *LoadBalancer) SetHealthCheck(v HealthCheck)`

SetHealthCheck gets a reference to the given HealthCheck and assigns it to the HealthCheck field.

### GetListeners

`func (o *LoadBalancer) GetListeners() []Listener`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *LoadBalancer) GetListenersOk() ([]Listener, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasListeners

`func (o *LoadBalancer) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### SetListeners

`func (o *LoadBalancer) SetListeners(v []Listener)`

SetListeners gets a reference to the given []Listener and assigns it to the Listeners field.

### GetLoadBalancerName

`func (o *LoadBalancer) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *LoadBalancer) GetLoadBalancerNameOk() (string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerName

`func (o *LoadBalancer) HasLoadBalancerName() bool`

HasLoadBalancerName returns a boolean if a field has been set.

### SetLoadBalancerName

`func (o *LoadBalancer) SetLoadBalancerName(v string)`

SetLoadBalancerName gets a reference to the given string and assigns it to the LoadBalancerName field.

### GetLoadBalancerStickyCookiePolicies

`func (o *LoadBalancer) GetLoadBalancerStickyCookiePolicies() []LoadBalancerStickyCookiePolicy`

GetLoadBalancerStickyCookiePolicies returns the LoadBalancerStickyCookiePolicies field if non-nil, zero value otherwise.

### GetLoadBalancerStickyCookiePoliciesOk

`func (o *LoadBalancer) GetLoadBalancerStickyCookiePoliciesOk() ([]LoadBalancerStickyCookiePolicy, bool)`

GetLoadBalancerStickyCookiePoliciesOk returns a tuple with the LoadBalancerStickyCookiePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerStickyCookiePolicies

`func (o *LoadBalancer) HasLoadBalancerStickyCookiePolicies() bool`

HasLoadBalancerStickyCookiePolicies returns a boolean if a field has been set.

### SetLoadBalancerStickyCookiePolicies

`func (o *LoadBalancer) SetLoadBalancerStickyCookiePolicies(v []LoadBalancerStickyCookiePolicy)`

SetLoadBalancerStickyCookiePolicies gets a reference to the given []LoadBalancerStickyCookiePolicy and assigns it to the LoadBalancerStickyCookiePolicies field.

### GetLoadBalancerType

`func (o *LoadBalancer) GetLoadBalancerType() string`

GetLoadBalancerType returns the LoadBalancerType field if non-nil, zero value otherwise.

### GetLoadBalancerTypeOk

`func (o *LoadBalancer) GetLoadBalancerTypeOk() (string, bool)`

GetLoadBalancerTypeOk returns a tuple with the LoadBalancerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoadBalancerType

`func (o *LoadBalancer) HasLoadBalancerType() bool`

HasLoadBalancerType returns a boolean if a field has been set.

### SetLoadBalancerType

`func (o *LoadBalancer) SetLoadBalancerType(v string)`

SetLoadBalancerType gets a reference to the given string and assigns it to the LoadBalancerType field.

### GetNetId

`func (o *LoadBalancer) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *LoadBalancer) GetNetIdOk() (string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetId

`func (o *LoadBalancer) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### SetNetId

`func (o *LoadBalancer) SetNetId(v string)`

SetNetId gets a reference to the given string and assigns it to the NetId field.

### GetSecurityGroups

`func (o *LoadBalancer) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *LoadBalancer) GetSecurityGroupsOk() ([]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroups

`func (o *LoadBalancer) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroups

`func (o *LoadBalancer) SetSecurityGroups(v []string)`

SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.

### GetSourceSecurityGroup

`func (o *LoadBalancer) GetSourceSecurityGroup() SourceSecurityGroup`

GetSourceSecurityGroup returns the SourceSecurityGroup field if non-nil, zero value otherwise.

### GetSourceSecurityGroupOk

`func (o *LoadBalancer) GetSourceSecurityGroupOk() (SourceSecurityGroup, bool)`

GetSourceSecurityGroupOk returns a tuple with the SourceSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceSecurityGroup

`func (o *LoadBalancer) HasSourceSecurityGroup() bool`

HasSourceSecurityGroup returns a boolean if a field has been set.

### SetSourceSecurityGroup

`func (o *LoadBalancer) SetSourceSecurityGroup(v SourceSecurityGroup)`

SetSourceSecurityGroup gets a reference to the given SourceSecurityGroup and assigns it to the SourceSecurityGroup field.

### GetSubnets

`func (o *LoadBalancer) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *LoadBalancer) GetSubnetsOk() ([]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubnets

`func (o *LoadBalancer) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnets

`func (o *LoadBalancer) SetSubnets(v []string)`

SetSubnets gets a reference to the given []string and assigns it to the Subnets field.

### GetSubregionNames

`func (o *LoadBalancer) GetSubregionNames() []string`

GetSubregionNames returns the SubregionNames field if non-nil, zero value otherwise.

### GetSubregionNamesOk

`func (o *LoadBalancer) GetSubregionNamesOk() ([]string, bool)`

GetSubregionNamesOk returns a tuple with the SubregionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubregionNames

`func (o *LoadBalancer) HasSubregionNames() bool`

HasSubregionNames returns a boolean if a field has been set.

### SetSubregionNames

`func (o *LoadBalancer) SetSubregionNames(v []string)`

SetSubregionNames gets a reference to the given []string and assigns it to the SubregionNames field.

### GetTags

`func (o *LoadBalancer) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LoadBalancer) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *LoadBalancer) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *LoadBalancer) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


