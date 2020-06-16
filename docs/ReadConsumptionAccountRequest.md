# ReadConsumptionAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**FromDate** | Pointer to **string** | The beginning of the time period, in ISO-8601 format with the date only (for example, &#x60;2017-06-14&#x60; or &#x60;2017-06-14T00:00:00Z&#x60;). | 
**ToDate** | Pointer to **string** | The end of the time period, in ISO-8601 format with the date only (for example, &#x60;2017-06-30&#x60; or &#x60;2017-06-30T00:00:00Z&#x60;). | 

## Methods

### GetDryRun

`func (o *ReadConsumptionAccountRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadConsumptionAccountRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *ReadConsumptionAccountRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *ReadConsumptionAccountRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetFromDate

`func (o *ReadConsumptionAccountRequest) GetFromDate() string`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *ReadConsumptionAccountRequest) GetFromDateOk() (string, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFromDate

`func (o *ReadConsumptionAccountRequest) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### SetFromDate

`func (o *ReadConsumptionAccountRequest) SetFromDate(v string)`

SetFromDate gets a reference to the given string and assigns it to the FromDate field.

### GetToDate

`func (o *ReadConsumptionAccountRequest) GetToDate() string`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *ReadConsumptionAccountRequest) GetToDateOk() (string, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasToDate

`func (o *ReadConsumptionAccountRequest) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### SetToDate

`func (o *ReadConsumptionAccountRequest) SetToDate(v string)`

SetToDate gets a reference to the given string and assigns it to the ToDate field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


