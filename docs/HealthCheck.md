# HealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckInterval** | Pointer to **int64** | The number of seconds between two pings (between &#x60;5&#x60; and &#x60;600&#x60; both included). | 
**HealthyThreshold** | Pointer to **int64** | The number of consecutive successful pings before considering the VM as healthy (between &#x60;2&#x60; and &#x60;10&#x60; both included). | 
**Path** | Pointer to **string** | The path for HTTP or HTTPS requests. | 
**Port** | Pointer to **int64** | The port number (between &#x60;1&#x60; and &#x60;65535&#x60;, both included). | 
**Protocol** | Pointer to **string** | The protocol for the URL of the VM (&#x60;HTTP&#x60; \\| &#x60;HTTPS&#x60; \\| &#x60;TCP&#x60; \\| &#x60;SSL&#x60; \\| &#x60;UDP&#x60;). | 
**Timeout** | Pointer to **int64** | The maximum waiting time for a response before considering the VM as unhealthy, in seconds (between &#x60;2&#x60; and &#x60;60&#x60; both included). | 
**UnhealthyThreshold** | Pointer to **int64** | The number of consecutive failed pings before considering the VM as unhealthy (between &#x60;2&#x60; and &#x60;10&#x60; both included). | 

## Methods

### GetCheckInterval

`func (o *HealthCheck) GetCheckInterval() int64`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *HealthCheck) GetCheckIntervalOk() (int64, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCheckInterval

`func (o *HealthCheck) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### SetCheckInterval

`func (o *HealthCheck) SetCheckInterval(v int64)`

SetCheckInterval gets a reference to the given int64 and assigns it to the CheckInterval field.

### GetHealthyThreshold

`func (o *HealthCheck) GetHealthyThreshold() int64`

GetHealthyThreshold returns the HealthyThreshold field if non-nil, zero value otherwise.

### GetHealthyThresholdOk

`func (o *HealthCheck) GetHealthyThresholdOk() (int64, bool)`

GetHealthyThresholdOk returns a tuple with the HealthyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHealthyThreshold

`func (o *HealthCheck) HasHealthyThreshold() bool`

HasHealthyThreshold returns a boolean if a field has been set.

### SetHealthyThreshold

`func (o *HealthCheck) SetHealthyThreshold(v int64)`

SetHealthyThreshold gets a reference to the given int64 and assigns it to the HealthyThreshold field.

### GetPath

`func (o *HealthCheck) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HealthCheck) GetPathOk() (string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPath

`func (o *HealthCheck) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPath

`func (o *HealthCheck) SetPath(v string)`

SetPath gets a reference to the given string and assigns it to the Path field.

### GetPort

`func (o *HealthCheck) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HealthCheck) GetPortOk() (int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPort

`func (o *HealthCheck) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPort

`func (o *HealthCheck) SetPort(v int64)`

SetPort gets a reference to the given int64 and assigns it to the Port field.

### GetProtocol

`func (o *HealthCheck) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HealthCheck) GetProtocolOk() (string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtocol

`func (o *HealthCheck) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocol

`func (o *HealthCheck) SetProtocol(v string)`

SetProtocol gets a reference to the given string and assigns it to the Protocol field.

### GetTimeout

`func (o *HealthCheck) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *HealthCheck) GetTimeoutOk() (int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimeout

`func (o *HealthCheck) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeout

`func (o *HealthCheck) SetTimeout(v int64)`

SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.

### GetUnhealthyThreshold

`func (o *HealthCheck) GetUnhealthyThreshold() int64`

GetUnhealthyThreshold returns the UnhealthyThreshold field if non-nil, zero value otherwise.

### GetUnhealthyThresholdOk

`func (o *HealthCheck) GetUnhealthyThresholdOk() (int64, bool)`

GetUnhealthyThresholdOk returns a tuple with the UnhealthyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUnhealthyThreshold

`func (o *HealthCheck) HasUnhealthyThreshold() bool`

HasUnhealthyThreshold returns a boolean if a field has been set.

### SetUnhealthyThreshold

`func (o *HealthCheck) SetUnhealthyThreshold(v int64)`

SetUnhealthyThreshold gets a reference to the given int64 and assigns it to the UnhealthyThreshold field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


