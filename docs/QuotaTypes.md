# QuotaTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuotaType** | Pointer to **string** | The resource ID if it is a resource-specific quota, &#x60;global&#x60; if it is not. | [optional] 
**Quotas** | Pointer to [**[]Quota**](Quota.md) | One or more quotas associated with the user. | [optional] 

## Methods

### GetQuotaType

`func (o *QuotaTypes) GetQuotaType() string`

GetQuotaType returns the QuotaType field if non-nil, zero value otherwise.

### GetQuotaTypeOk

`func (o *QuotaTypes) GetQuotaTypeOk() (string, bool)`

GetQuotaTypeOk returns a tuple with the QuotaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuotaType

`func (o *QuotaTypes) HasQuotaType() bool`

HasQuotaType returns a boolean if a field has been set.

### SetQuotaType

`func (o *QuotaTypes) SetQuotaType(v string)`

SetQuotaType gets a reference to the given string and assigns it to the QuotaType field.

### GetQuotas

`func (o *QuotaTypes) GetQuotas() []Quota`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *QuotaTypes) GetQuotasOk() ([]Quota, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuotas

`func (o *QuotaTypes) HasQuotas() bool`

HasQuotas returns a boolean if a field has been set.

### SetQuotas

`func (o *QuotaTypes) SetQuotas(v []Quota)`

SetQuotas gets a reference to the given []Quota and assigns it to the Quotas field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


