# ReadApiLogsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Filters** | Pointer to [**FiltersApiLog**](FiltersApiLog.md) |  | [optional] 
**NextPageToken** | Pointer to **string** | The token to request the next page of results. | [optional] 
**ResultsPerPage** | Pointer to **int64** | The maximum number of items returned in a single page. By default, 100. | [optional] 
**With** | Pointer to [**With**](With.md) |  | [optional] 

## Methods

### GetDryRun

`func (o *ReadApiLogsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadApiLogsRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *ReadApiLogsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *ReadApiLogsRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetFilters

`func (o *ReadApiLogsRequest) GetFilters() FiltersApiLog`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReadApiLogsRequest) GetFiltersOk() (FiltersApiLog, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFilters

`func (o *ReadApiLogsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFilters

`func (o *ReadApiLogsRequest) SetFilters(v FiltersApiLog)`

SetFilters gets a reference to the given FiltersApiLog and assigns it to the Filters field.

### GetNextPageToken

`func (o *ReadApiLogsRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ReadApiLogsRequest) GetNextPageTokenOk() (string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNextPageToken

`func (o *ReadApiLogsRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageToken

`func (o *ReadApiLogsRequest) SetNextPageToken(v string)`

SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.

### GetResultsPerPage

`func (o *ReadApiLogsRequest) GetResultsPerPage() int64`

GetResultsPerPage returns the ResultsPerPage field if non-nil, zero value otherwise.

### GetResultsPerPageOk

`func (o *ReadApiLogsRequest) GetResultsPerPageOk() (int64, bool)`

GetResultsPerPageOk returns a tuple with the ResultsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResultsPerPage

`func (o *ReadApiLogsRequest) HasResultsPerPage() bool`

HasResultsPerPage returns a boolean if a field has been set.

### SetResultsPerPage

`func (o *ReadApiLogsRequest) SetResultsPerPage(v int64)`

SetResultsPerPage gets a reference to the given int64 and assigns it to the ResultsPerPage field.

### GetWith

`func (o *ReadApiLogsRequest) GetWith() With`

GetWith returns the With field if non-nil, zero value otherwise.

### GetWithOk

`func (o *ReadApiLogsRequest) GetWithOk() (With, bool)`

GetWithOk returns a tuple with the With field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWith

`func (o *ReadApiLogsRequest) HasWith() bool`

HasWith returns a boolean if a field has been set.

### SetWith

`func (o *ReadApiLogsRequest) SetWith(v With)`

SetWith gets a reference to the given With and assigns it to the With field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


