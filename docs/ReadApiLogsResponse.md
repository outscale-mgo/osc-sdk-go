# ReadApiLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]Log**](Log.md) | Information displayed in one or more API logs. | [optional] 
**NextPageToken** | Pointer to **string** | The token to request the next page of results. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetLogs

`func (o *ReadApiLogsResponse) GetLogs() []Log`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ReadApiLogsResponse) GetLogsOk() ([]Log, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogs

`func (o *ReadApiLogsResponse) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### SetLogs

`func (o *ReadApiLogsResponse) SetLogs(v []Log)`

SetLogs gets a reference to the given []Log and assigns it to the Logs field.

### GetNextPageToken

`func (o *ReadApiLogsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ReadApiLogsResponse) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *ReadApiLogsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *ReadApiLogsResponse) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.

### GetResponseContext

`func (o *ReadApiLogsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadApiLogsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadApiLogsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadApiLogsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


