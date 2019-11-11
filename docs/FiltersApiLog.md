# FiltersApiLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryAccessKeys** | Pointer to **[]string** | One or more API keys used for the query. | [optional] 
**QueryApiNames** | Pointer to **[]string** | The name of one or more API services used for the query. | [optional] 
**QueryCallNames** | Pointer to **[]string** | The name of one or more calls. | [optional] 
**QueryDateAfter** | Pointer to **string** | The logs of the queries made after the date and time you specify (in ISO 8601 base or extended format). | [optional] 
**QueryDateBefore** | Pointer to **string** | The logs of the queries made before the date and time you specify (in ISO 8601 base or extended format). | [optional] 
**QueryIpAddresses** | Pointer to **[]string** | One or more IP addresses used for the query. | [optional] 
**QueryUserAgents** | Pointer to **[]string** | One or more user agents used for the HTTP request. | [optional] 
**RequestIds** | Pointer to **[]string** | One or more request IDs. | [optional] 
**ResponseStatusCodes** | Pointer to **[]int32** | One or more HTTP codes provided by the responses. | [optional] 

## Methods

### GetQueryAccessKeys

`func (o *FiltersApiLog) GetQueryAccessKeys() []string`

GetQueryAccessKeys returns the QueryAccessKeys field if non-nil, zero value otherwise.

### GetQueryAccessKeysOk

`func (o *FiltersApiLog) GetQueryAccessKeysOk() ([]string, bool)`

GetQueryAccessKeysOk returns a tuple with the QueryAccessKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryAccessKeys

`func (o *FiltersApiLog) HasQueryAccessKeys() bool`

HasQueryAccessKeys returns a boolean if a field has been set.

### SetQueryAccessKeys

`func (o *FiltersApiLog) SetQueryAccessKeys(v []string)`

SetQueryAccessKeys gets a reference to the given []string and assigns it to the QueryAccessKeys field.

### GetQueryApiNames

`func (o *FiltersApiLog) GetQueryApiNames() []string`

GetQueryApiNames returns the QueryApiNames field if non-nil, zero value otherwise.

### GetQueryApiNamesOk

`func (o *FiltersApiLog) GetQueryApiNamesOk() ([]string, bool)`

GetQueryApiNamesOk returns a tuple with the QueryApiNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryApiNames

`func (o *FiltersApiLog) HasQueryApiNames() bool`

HasQueryApiNames returns a boolean if a field has been set.

### SetQueryApiNames

`func (o *FiltersApiLog) SetQueryApiNames(v []string)`

SetQueryApiNames gets a reference to the given []string and assigns it to the QueryApiNames field.

### GetQueryCallNames

`func (o *FiltersApiLog) GetQueryCallNames() []string`

GetQueryCallNames returns the QueryCallNames field if non-nil, zero value otherwise.

### GetQueryCallNamesOk

`func (o *FiltersApiLog) GetQueryCallNamesOk() ([]string, bool)`

GetQueryCallNamesOk returns a tuple with the QueryCallNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryCallNames

`func (o *FiltersApiLog) HasQueryCallNames() bool`

HasQueryCallNames returns a boolean if a field has been set.

### SetQueryCallNames

`func (o *FiltersApiLog) SetQueryCallNames(v []string)`

SetQueryCallNames gets a reference to the given []string and assigns it to the QueryCallNames field.

### GetQueryDateAfter

`func (o *FiltersApiLog) GetQueryDateAfter() string`

GetQueryDateAfter returns the QueryDateAfter field if non-nil, zero value otherwise.

### GetQueryDateAfterOk

`func (o *FiltersApiLog) GetQueryDateAfterOk() (string, bool)`

GetQueryDateAfterOk returns a tuple with the QueryDateAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryDateAfter

`func (o *FiltersApiLog) HasQueryDateAfter() bool`

HasQueryDateAfter returns a boolean if a field has been set.

### SetQueryDateAfter

`func (o *FiltersApiLog) SetQueryDateAfter(v string)`

SetQueryDateAfter gets a reference to the given string and assigns it to the QueryDateAfter field.

### GetQueryDateBefore

`func (o *FiltersApiLog) GetQueryDateBefore() string`

GetQueryDateBefore returns the QueryDateBefore field if non-nil, zero value otherwise.

### GetQueryDateBeforeOk

`func (o *FiltersApiLog) GetQueryDateBeforeOk() (string, bool)`

GetQueryDateBeforeOk returns a tuple with the QueryDateBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryDateBefore

`func (o *FiltersApiLog) HasQueryDateBefore() bool`

HasQueryDateBefore returns a boolean if a field has been set.

### SetQueryDateBefore

`func (o *FiltersApiLog) SetQueryDateBefore(v string)`

SetQueryDateBefore gets a reference to the given string and assigns it to the QueryDateBefore field.

### GetQueryIpAddresses

`func (o *FiltersApiLog) GetQueryIpAddresses() []string`

GetQueryIpAddresses returns the QueryIpAddresses field if non-nil, zero value otherwise.

### GetQueryIpAddressesOk

`func (o *FiltersApiLog) GetQueryIpAddressesOk() ([]string, bool)`

GetQueryIpAddressesOk returns a tuple with the QueryIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryIpAddresses

`func (o *FiltersApiLog) HasQueryIpAddresses() bool`

HasQueryIpAddresses returns a boolean if a field has been set.

### SetQueryIpAddresses

`func (o *FiltersApiLog) SetQueryIpAddresses(v []string)`

SetQueryIpAddresses gets a reference to the given []string and assigns it to the QueryIpAddresses field.

### GetQueryUserAgents

`func (o *FiltersApiLog) GetQueryUserAgents() []string`

GetQueryUserAgents returns the QueryUserAgents field if non-nil, zero value otherwise.

### GetQueryUserAgentsOk

`func (o *FiltersApiLog) GetQueryUserAgentsOk() ([]string, bool)`

GetQueryUserAgentsOk returns a tuple with the QueryUserAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryUserAgents

`func (o *FiltersApiLog) HasQueryUserAgents() bool`

HasQueryUserAgents returns a boolean if a field has been set.

### SetQueryUserAgents

`func (o *FiltersApiLog) SetQueryUserAgents(v []string)`

SetQueryUserAgents gets a reference to the given []string and assigns it to the QueryUserAgents field.

### GetRequestIds

`func (o *FiltersApiLog) GetRequestIds() []string`

GetRequestIds returns the RequestIds field if non-nil, zero value otherwise.

### GetRequestIdsOk

`func (o *FiltersApiLog) GetRequestIdsOk() ([]string, bool)`

GetRequestIdsOk returns a tuple with the RequestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequestIds

`func (o *FiltersApiLog) HasRequestIds() bool`

HasRequestIds returns a boolean if a field has been set.

### SetRequestIds

`func (o *FiltersApiLog) SetRequestIds(v []string)`

SetRequestIds gets a reference to the given []string and assigns it to the RequestIds field.

### GetResponseStatusCodes

`func (o *FiltersApiLog) GetResponseStatusCodes() []int32`

GetResponseStatusCodes returns the ResponseStatusCodes field if non-nil, zero value otherwise.

### GetResponseStatusCodesOk

`func (o *FiltersApiLog) GetResponseStatusCodesOk() ([]int32, bool)`

GetResponseStatusCodesOk returns a tuple with the ResponseStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseStatusCodes

`func (o *FiltersApiLog) HasResponseStatusCodes() bool`

HasResponseStatusCodes returns a boolean if a field has been set.

### SetResponseStatusCodes

`func (o *FiltersApiLog) SetResponseStatusCodes(v []int32)`

SetResponseStatusCodes gets a reference to the given []int32 and assigns it to the ResponseStatusCodes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


