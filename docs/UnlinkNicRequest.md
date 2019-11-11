# UnlinkNicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**LinkNicId** | Pointer to **string** | The ID of the attachment operation. | 

## Methods

### GetDryRun

`func (o *UnlinkNicRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UnlinkNicRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UnlinkNicRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UnlinkNicRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetLinkNicId

`func (o *UnlinkNicRequest) GetLinkNicId() string`

GetLinkNicId returns the LinkNicId field if non-nil, zero value otherwise.

### GetLinkNicIdOk

`func (o *UnlinkNicRequest) GetLinkNicIdOk() (string, bool)`

GetLinkNicIdOk returns a tuple with the LinkNicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkNicId

`func (o *UnlinkNicRequest) HasLinkNicId() bool`

HasLinkNicId returns a boolean if a field has been set.

### SetLinkNicId

`func (o *UnlinkNicRequest) SetLinkNicId(v string)`

SetLinkNicId gets a reference to the given string and assigns it to the LinkNicId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


