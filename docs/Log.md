# Log

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID. | [optional] 
**CallDuration** | Pointer to **int64** | The duration of the call (in milliseconds). | [optional] 
**QueryAccessKey** | Pointer to **string** | The API key used for the query. | [optional] 
**QueryApiName** | Pointer to **string** | The name of the API service used by the call. | [optional] 
**QueryApiVersion** | Pointer to **string** | The version of the API service used by the call. | [optional] 
**QueryCallName** | Pointer to **string** | The name of the call. | [optional] 
**QueryDate** | Pointer to **string** | The date and time of the query (in ISO 8601 base or extended format). | [optional] 
**QueryHeaderRaw** | Pointer to **string** | The query header raw. | [optional] 
**QueryHeaderSize** | Pointer to **int64** | The query header size. | [optional] 
**QueryIpAddress** | Pointer to **string** | The IP address used for the query. | [optional] 
**QueryPayloadRaw** | Pointer to **string** | The query payload raw. | [optional] 
**QueryPayloadSize** | Pointer to **int64** | The query payload size. | [optional] 
**QueryUserAgent** | Pointer to **string** | The user agent used for the HTTP request. | [optional] 
**RequestId** | Pointer to **string** | The ID provided in the response. | [optional] 
**ResponseSize** | Pointer to **int64** | The size of the response (in bytes). | [optional] 
**ResponseStatusCode** | Pointer to **int64** | The HTTP code provided in the response. | [optional] 

## Methods

### GetAccountId

`func (o *Log) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Log) GetAccountIdOk() (string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *Log) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *Log) SetAccountId(v string)`

SetAccountId gets a reference to the given string and assigns it to the AccountId field.

### GetCallDuration

`func (o *Log) GetCallDuration() int64`

GetCallDuration returns the CallDuration field if non-nil, zero value otherwise.

### GetCallDurationOk

`func (o *Log) GetCallDurationOk() (int64, bool)`

GetCallDurationOk returns a tuple with the CallDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCallDuration

`func (o *Log) HasCallDuration() bool`

HasCallDuration returns a boolean if a field has been set.

### SetCallDuration

`func (o *Log) SetCallDuration(v int64)`

SetCallDuration gets a reference to the given int64 and assigns it to the CallDuration field.

### GetQueryAccessKey

`func (o *Log) GetQueryAccessKey() string`

GetQueryAccessKey returns the QueryAccessKey field if non-nil, zero value otherwise.

### GetQueryAccessKeyOk

`func (o *Log) GetQueryAccessKeyOk() (string, bool)`

GetQueryAccessKeyOk returns a tuple with the QueryAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryAccessKey

`func (o *Log) HasQueryAccessKey() bool`

HasQueryAccessKey returns a boolean if a field has been set.

### SetQueryAccessKey

`func (o *Log) SetQueryAccessKey(v string)`

SetQueryAccessKey gets a reference to the given string and assigns it to the QueryAccessKey field.

### GetQueryApiName

`func (o *Log) GetQueryApiName() string`

GetQueryApiName returns the QueryApiName field if non-nil, zero value otherwise.

### GetQueryApiNameOk

`func (o *Log) GetQueryApiNameOk() (string, bool)`

GetQueryApiNameOk returns a tuple with the QueryApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryApiName

`func (o *Log) HasQueryApiName() bool`

HasQueryApiName returns a boolean if a field has been set.

### SetQueryApiName

`func (o *Log) SetQueryApiName(v string)`

SetQueryApiName gets a reference to the given string and assigns it to the QueryApiName field.

### GetQueryApiVersion

`func (o *Log) GetQueryApiVersion() string`

GetQueryApiVersion returns the QueryApiVersion field if non-nil, zero value otherwise.

### GetQueryApiVersionOk

`func (o *Log) GetQueryApiVersionOk() (string, bool)`

GetQueryApiVersionOk returns a tuple with the QueryApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryApiVersion

`func (o *Log) HasQueryApiVersion() bool`

HasQueryApiVersion returns a boolean if a field has been set.

### SetQueryApiVersion

`func (o *Log) SetQueryApiVersion(v string)`

SetQueryApiVersion gets a reference to the given string and assigns it to the QueryApiVersion field.

### GetQueryCallName

`func (o *Log) GetQueryCallName() string`

GetQueryCallName returns the QueryCallName field if non-nil, zero value otherwise.

### GetQueryCallNameOk

`func (o *Log) GetQueryCallNameOk() (string, bool)`

GetQueryCallNameOk returns a tuple with the QueryCallName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryCallName

`func (o *Log) HasQueryCallName() bool`

HasQueryCallName returns a boolean if a field has been set.

### SetQueryCallName

`func (o *Log) SetQueryCallName(v string)`

SetQueryCallName gets a reference to the given string and assigns it to the QueryCallName field.

### GetQueryDate

`func (o *Log) GetQueryDate() string`

GetQueryDate returns the QueryDate field if non-nil, zero value otherwise.

### GetQueryDateOk

`func (o *Log) GetQueryDateOk() (string, bool)`

GetQueryDateOk returns a tuple with the QueryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryDate

`func (o *Log) HasQueryDate() bool`

HasQueryDate returns a boolean if a field has been set.

### SetQueryDate

`func (o *Log) SetQueryDate(v string)`

SetQueryDate gets a reference to the given string and assigns it to the QueryDate field.

### GetQueryHeaderRaw

`func (o *Log) GetQueryHeaderRaw() string`

GetQueryHeaderRaw returns the QueryHeaderRaw field if non-nil, zero value otherwise.

### GetQueryHeaderRawOk

`func (o *Log) GetQueryHeaderRawOk() (string, bool)`

GetQueryHeaderRawOk returns a tuple with the QueryHeaderRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryHeaderRaw

`func (o *Log) HasQueryHeaderRaw() bool`

HasQueryHeaderRaw returns a boolean if a field has been set.

### SetQueryHeaderRaw

`func (o *Log) SetQueryHeaderRaw(v string)`

SetQueryHeaderRaw gets a reference to the given string and assigns it to the QueryHeaderRaw field.

### GetQueryHeaderSize

`func (o *Log) GetQueryHeaderSize() int64`

GetQueryHeaderSize returns the QueryHeaderSize field if non-nil, zero value otherwise.

### GetQueryHeaderSizeOk

`func (o *Log) GetQueryHeaderSizeOk() (int64, bool)`

GetQueryHeaderSizeOk returns a tuple with the QueryHeaderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryHeaderSize

`func (o *Log) HasQueryHeaderSize() bool`

HasQueryHeaderSize returns a boolean if a field has been set.

### SetQueryHeaderSize

`func (o *Log) SetQueryHeaderSize(v int64)`

SetQueryHeaderSize gets a reference to the given int64 and assigns it to the QueryHeaderSize field.

### GetQueryIpAddress

`func (o *Log) GetQueryIpAddress() string`

GetQueryIpAddress returns the QueryIpAddress field if non-nil, zero value otherwise.

### GetQueryIpAddressOk

`func (o *Log) GetQueryIpAddressOk() (string, bool)`

GetQueryIpAddressOk returns a tuple with the QueryIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryIpAddress

`func (o *Log) HasQueryIpAddress() bool`

HasQueryIpAddress returns a boolean if a field has been set.

### SetQueryIpAddress

`func (o *Log) SetQueryIpAddress(v string)`

SetQueryIpAddress gets a reference to the given string and assigns it to the QueryIpAddress field.

### GetQueryPayloadRaw

`func (o *Log) GetQueryPayloadRaw() string`

GetQueryPayloadRaw returns the QueryPayloadRaw field if non-nil, zero value otherwise.

### GetQueryPayloadRawOk

`func (o *Log) GetQueryPayloadRawOk() (string, bool)`

GetQueryPayloadRawOk returns a tuple with the QueryPayloadRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryPayloadRaw

`func (o *Log) HasQueryPayloadRaw() bool`

HasQueryPayloadRaw returns a boolean if a field has been set.

### SetQueryPayloadRaw

`func (o *Log) SetQueryPayloadRaw(v string)`

SetQueryPayloadRaw gets a reference to the given string and assigns it to the QueryPayloadRaw field.

### GetQueryPayloadSize

`func (o *Log) GetQueryPayloadSize() int64`

GetQueryPayloadSize returns the QueryPayloadSize field if non-nil, zero value otherwise.

### GetQueryPayloadSizeOk

`func (o *Log) GetQueryPayloadSizeOk() (int64, bool)`

GetQueryPayloadSizeOk returns a tuple with the QueryPayloadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryPayloadSize

`func (o *Log) HasQueryPayloadSize() bool`

HasQueryPayloadSize returns a boolean if a field has been set.

### SetQueryPayloadSize

`func (o *Log) SetQueryPayloadSize(v int64)`

SetQueryPayloadSize gets a reference to the given int64 and assigns it to the QueryPayloadSize field.

### GetQueryUserAgent

`func (o *Log) GetQueryUserAgent() string`

GetQueryUserAgent returns the QueryUserAgent field if non-nil, zero value otherwise.

### GetQueryUserAgentOk

`func (o *Log) GetQueryUserAgentOk() (string, bool)`

GetQueryUserAgentOk returns a tuple with the QueryUserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryUserAgent

`func (o *Log) HasQueryUserAgent() bool`

HasQueryUserAgent returns a boolean if a field has been set.

### SetQueryUserAgent

`func (o *Log) SetQueryUserAgent(v string)`

SetQueryUserAgent gets a reference to the given string and assigns it to the QueryUserAgent field.

### GetRequestId

`func (o *Log) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Log) GetRequestIdOk() (string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequestId

`func (o *Log) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestId

`func (o *Log) SetRequestId(v string)`

SetRequestId gets a reference to the given string and assigns it to the RequestId field.

### GetResponseSize

`func (o *Log) GetResponseSize() int64`

GetResponseSize returns the ResponseSize field if non-nil, zero value otherwise.

### GetResponseSizeOk

`func (o *Log) GetResponseSizeOk() (int64, bool)`

GetResponseSizeOk returns a tuple with the ResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseSize

`func (o *Log) HasResponseSize() bool`

HasResponseSize returns a boolean if a field has been set.

### SetResponseSize

`func (o *Log) SetResponseSize(v int64)`

SetResponseSize gets a reference to the given int64 and assigns it to the ResponseSize field.

### GetResponseStatusCode

`func (o *Log) GetResponseStatusCode() int64`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *Log) GetResponseStatusCodeOk() (int64, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseStatusCode

`func (o *Log) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### SetResponseStatusCode

`func (o *Log) SetResponseStatusCode(v int64)`

SetResponseStatusCode gets a reference to the given int64 and assigns it to the ResponseStatusCode field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


