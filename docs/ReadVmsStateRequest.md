# ReadVmsStateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllVms** | Pointer to **bool** | If &#x60;true&#x60;, includes the status of all VMs. By default or if set to &#x60;false&#x60;, only includes the status of running VMs. | [optional] 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Filters** | Pointer to [**FiltersVmsState**](FiltersVmsState.md) |  | [optional] 

## Methods

### GetAllVms

`func (o *ReadVmsStateRequest) GetAllVms() bool`

GetAllVms returns the AllVms field if non-nil, zero value otherwise.

### GetAllVmsOk

`func (o *ReadVmsStateRequest) GetAllVmsOk() (bool, bool)`

GetAllVmsOk returns a tuple with the AllVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllVms

`func (o *ReadVmsStateRequest) HasAllVms() bool`

HasAllVms returns a boolean if a field has been set.

### SetAllVms

`func (o *ReadVmsStateRequest) SetAllVms(v bool)`

SetAllVms gets a reference to the given bool and assigns it to the AllVms field.

### GetDryRun

`func (o *ReadVmsStateRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadVmsStateRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *ReadVmsStateRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *ReadVmsStateRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetFilters

`func (o *ReadVmsStateRequest) GetFilters() FiltersVmsState`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ReadVmsStateRequest) GetFiltersOk() (FiltersVmsState, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFilters

`func (o *ReadVmsStateRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFilters

`func (o *ReadVmsStateRequest) SetFilters(v FiltersVmsState)`

SetFilters gets a reference to the given FiltersVmsState and assigns it to the Filters field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


