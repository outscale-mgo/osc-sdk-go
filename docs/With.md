# With

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **bool** | If set to &#x60;true&#x60;, the account ID is displayed in the logs. | [optional] 
**CallDuration** | Pointer to **bool** | If set to &#x60;true&#x60;, the duration of the call is displayed each log. | [optional] 
**QueryAccessKey** | Pointer to **bool** | If set to &#x60;true&#x60;, the API key used for the query is displayed each log. | [optional] 
**QueryApiName** | Pointer to **bool** | If set to &#x60;true&#x60;, the name of the API service used by the call is displayed in each log (&#x60;oapi&#x60; \\| &#x60;fcu&#x60; \\| &#x60;lbu&#x60; \\| &#x60;directlink&#x60; \\| &#x60;eim&#x60; \\| &#x60;icu&#x60;). | [optional] 
**QueryApiVersion** | Pointer to **bool** | If set to &#x60;true&#x60;, the version of the API service used by the call is displayed in each log. | [optional] 
**QueryCallName** | Pointer to **bool** | If set to &#x60;true&#x60;, the name of the call is displayed in each log. | [optional] 
**QueryDate** | Pointer to **bool** | If set to &#x60;true&#x60;, the date of the call is displayed in each log. | [optional] 
**QueryHeaderRaw** | Pointer to **bool** | If set to &#x60;true&#x60;, the query header RAW is displayed in each log. | [optional] 
**QueryHeaderSize** | Pointer to **bool** | If set to &#x60;true&#x60;, the query header size is displayed in each log. | [optional] 
**QueryIpAddress** | Pointer to **bool** | If set to &#x60;true&#x60;, the IP address used to make to query is displayed in each log. | [optional] 
**QueryPayloadRaw** | Pointer to **bool** | If set to &#x60;true&#x60;, the query payload raw is displayed in each log. | [optional] 
**QueryPayloadSize** | Pointer to **bool** | If set to &#x60;true&#x60;, the query payload size is displayed in each log. | [optional] 
**QueryUserAgent** | Pointer to **bool** | If set to &#x60;true&#x60;, the user agent used to make the HTTP request is displayed in each log. | [optional] 
**RequestId** | Pointer to **bool** | By default ot if set to &#x60;true&#x60;, the ID of the call is displayed in each log. | [optional] 
**ResponseSize** | Pointer to **bool** | If set to &#x60;true&#x60;, the size of the response (in bytes) is displayed in each log. | [optional] 
**ResponseStatusCode** | Pointer to **bool** | If set to &#x60;true&#x60;, the HTTP code provided by the response is displayed in each log. | [optional] 

## Methods

### GetAccountId

`func (o *With) GetAccountId() bool`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *With) GetAccountIdOk() (bool, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *With) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *With) SetAccountId(v bool)`

SetAccountId gets a reference to the given bool and assigns it to the AccountId field.

### GetCallDuration

`func (o *With) GetCallDuration() bool`

GetCallDuration returns the CallDuration field if non-nil, zero value otherwise.

### GetCallDurationOk

`func (o *With) GetCallDurationOk() (bool, bool)`

GetCallDurationOk returns a tuple with the CallDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCallDuration

`func (o *With) HasCallDuration() bool`

HasCallDuration returns a boolean if a field has been set.

### SetCallDuration

`func (o *With) SetCallDuration(v bool)`

SetCallDuration gets a reference to the given bool and assigns it to the CallDuration field.

### GetQueryAccessKey

`func (o *With) GetQueryAccessKey() bool`

GetQueryAccessKey returns the QueryAccessKey field if non-nil, zero value otherwise.

### GetQueryAccessKeyOk

`func (o *With) GetQueryAccessKeyOk() (bool, bool)`

GetQueryAccessKeyOk returns a tuple with the QueryAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryAccessKey

`func (o *With) HasQueryAccessKey() bool`

HasQueryAccessKey returns a boolean if a field has been set.

### SetQueryAccessKey

`func (o *With) SetQueryAccessKey(v bool)`

SetQueryAccessKey gets a reference to the given bool and assigns it to the QueryAccessKey field.

### GetQueryApiName

`func (o *With) GetQueryApiName() bool`

GetQueryApiName returns the QueryApiName field if non-nil, zero value otherwise.

### GetQueryApiNameOk

`func (o *With) GetQueryApiNameOk() (bool, bool)`

GetQueryApiNameOk returns a tuple with the QueryApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryApiName

`func (o *With) HasQueryApiName() bool`

HasQueryApiName returns a boolean if a field has been set.

### SetQueryApiName

`func (o *With) SetQueryApiName(v bool)`

SetQueryApiName gets a reference to the given bool and assigns it to the QueryApiName field.

### GetQueryApiVersion

`func (o *With) GetQueryApiVersion() bool`

GetQueryApiVersion returns the QueryApiVersion field if non-nil, zero value otherwise.

### GetQueryApiVersionOk

`func (o *With) GetQueryApiVersionOk() (bool, bool)`

GetQueryApiVersionOk returns a tuple with the QueryApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryApiVersion

`func (o *With) HasQueryApiVersion() bool`

HasQueryApiVersion returns a boolean if a field has been set.

### SetQueryApiVersion

`func (o *With) SetQueryApiVersion(v bool)`

SetQueryApiVersion gets a reference to the given bool and assigns it to the QueryApiVersion field.

### GetQueryCallName

`func (o *With) GetQueryCallName() bool`

GetQueryCallName returns the QueryCallName field if non-nil, zero value otherwise.

### GetQueryCallNameOk

`func (o *With) GetQueryCallNameOk() (bool, bool)`

GetQueryCallNameOk returns a tuple with the QueryCallName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryCallName

`func (o *With) HasQueryCallName() bool`

HasQueryCallName returns a boolean if a field has been set.

### SetQueryCallName

`func (o *With) SetQueryCallName(v bool)`

SetQueryCallName gets a reference to the given bool and assigns it to the QueryCallName field.

### GetQueryDate

`func (o *With) GetQueryDate() bool`

GetQueryDate returns the QueryDate field if non-nil, zero value otherwise.

### GetQueryDateOk

`func (o *With) GetQueryDateOk() (bool, bool)`

GetQueryDateOk returns a tuple with the QueryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryDate

`func (o *With) HasQueryDate() bool`

HasQueryDate returns a boolean if a field has been set.

### SetQueryDate

`func (o *With) SetQueryDate(v bool)`

SetQueryDate gets a reference to the given bool and assigns it to the QueryDate field.

### GetQueryHeaderRaw

`func (o *With) GetQueryHeaderRaw() bool`

GetQueryHeaderRaw returns the QueryHeaderRaw field if non-nil, zero value otherwise.

### GetQueryHeaderRawOk

`func (o *With) GetQueryHeaderRawOk() (bool, bool)`

GetQueryHeaderRawOk returns a tuple with the QueryHeaderRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryHeaderRaw

`func (o *With) HasQueryHeaderRaw() bool`

HasQueryHeaderRaw returns a boolean if a field has been set.

### SetQueryHeaderRaw

`func (o *With) SetQueryHeaderRaw(v bool)`

SetQueryHeaderRaw gets a reference to the given bool and assigns it to the QueryHeaderRaw field.

### GetQueryHeaderSize

`func (o *With) GetQueryHeaderSize() bool`

GetQueryHeaderSize returns the QueryHeaderSize field if non-nil, zero value otherwise.

### GetQueryHeaderSizeOk

`func (o *With) GetQueryHeaderSizeOk() (bool, bool)`

GetQueryHeaderSizeOk returns a tuple with the QueryHeaderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryHeaderSize

`func (o *With) HasQueryHeaderSize() bool`

HasQueryHeaderSize returns a boolean if a field has been set.

### SetQueryHeaderSize

`func (o *With) SetQueryHeaderSize(v bool)`

SetQueryHeaderSize gets a reference to the given bool and assigns it to the QueryHeaderSize field.

### GetQueryIpAddress

`func (o *With) GetQueryIpAddress() bool`

GetQueryIpAddress returns the QueryIpAddress field if non-nil, zero value otherwise.

### GetQueryIpAddressOk

`func (o *With) GetQueryIpAddressOk() (bool, bool)`

GetQueryIpAddressOk returns a tuple with the QueryIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryIpAddress

`func (o *With) HasQueryIpAddress() bool`

HasQueryIpAddress returns a boolean if a field has been set.

### SetQueryIpAddress

`func (o *With) SetQueryIpAddress(v bool)`

SetQueryIpAddress gets a reference to the given bool and assigns it to the QueryIpAddress field.

### GetQueryPayloadRaw

`func (o *With) GetQueryPayloadRaw() bool`

GetQueryPayloadRaw returns the QueryPayloadRaw field if non-nil, zero value otherwise.

### GetQueryPayloadRawOk

`func (o *With) GetQueryPayloadRawOk() (bool, bool)`

GetQueryPayloadRawOk returns a tuple with the QueryPayloadRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryPayloadRaw

`func (o *With) HasQueryPayloadRaw() bool`

HasQueryPayloadRaw returns a boolean if a field has been set.

### SetQueryPayloadRaw

`func (o *With) SetQueryPayloadRaw(v bool)`

SetQueryPayloadRaw gets a reference to the given bool and assigns it to the QueryPayloadRaw field.

### GetQueryPayloadSize

`func (o *With) GetQueryPayloadSize() bool`

GetQueryPayloadSize returns the QueryPayloadSize field if non-nil, zero value otherwise.

### GetQueryPayloadSizeOk

`func (o *With) GetQueryPayloadSizeOk() (bool, bool)`

GetQueryPayloadSizeOk returns a tuple with the QueryPayloadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryPayloadSize

`func (o *With) HasQueryPayloadSize() bool`

HasQueryPayloadSize returns a boolean if a field has been set.

### SetQueryPayloadSize

`func (o *With) SetQueryPayloadSize(v bool)`

SetQueryPayloadSize gets a reference to the given bool and assigns it to the QueryPayloadSize field.

### GetQueryUserAgent

`func (o *With) GetQueryUserAgent() bool`

GetQueryUserAgent returns the QueryUserAgent field if non-nil, zero value otherwise.

### GetQueryUserAgentOk

`func (o *With) GetQueryUserAgentOk() (bool, bool)`

GetQueryUserAgentOk returns a tuple with the QueryUserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQueryUserAgent

`func (o *With) HasQueryUserAgent() bool`

HasQueryUserAgent returns a boolean if a field has been set.

### SetQueryUserAgent

`func (o *With) SetQueryUserAgent(v bool)`

SetQueryUserAgent gets a reference to the given bool and assigns it to the QueryUserAgent field.

### GetRequestId

`func (o *With) GetRequestId() bool`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *With) GetRequestIdOk() (bool, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequestId

`func (o *With) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestId

`func (o *With) SetRequestId(v bool)`

SetRequestId gets a reference to the given bool and assigns it to the RequestId field.

### GetResponseSize

`func (o *With) GetResponseSize() bool`

GetResponseSize returns the ResponseSize field if non-nil, zero value otherwise.

### GetResponseSizeOk

`func (o *With) GetResponseSizeOk() (bool, bool)`

GetResponseSizeOk returns a tuple with the ResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseSize

`func (o *With) HasResponseSize() bool`

HasResponseSize returns a boolean if a field has been set.

### SetResponseSize

`func (o *With) SetResponseSize(v bool)`

SetResponseSize gets a reference to the given bool and assigns it to the ResponseSize field.

### GetResponseStatusCode

`func (o *With) GetResponseStatusCode() bool`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *With) GetResponseStatusCodeOk() (bool, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseStatusCode

`func (o *With) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### SetResponseStatusCode

`func (o *With) SetResponseStatusCode(v bool)`

SetResponseStatusCode gets a reference to the given bool and assigns it to the ResponseStatusCode field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


