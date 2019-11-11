# DeleteTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**ResourceIds** | Pointer to **[]string** | One or more resource IDs. | 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags to delete (if you set a tag value, only the tags matching exactly this value are deleted). | 

## Methods

### GetDryRun

`func (o *DeleteTagsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteTagsRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteTagsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteTagsRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetResourceIds

`func (o *DeleteTagsRequest) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *DeleteTagsRequest) GetResourceIdsOk() ([]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceIds

`func (o *DeleteTagsRequest) HasResourceIds() bool`

HasResourceIds returns a boolean if a field has been set.

### SetResourceIds

`func (o *DeleteTagsRequest) SetResourceIds(v []string)`

SetResourceIds gets a reference to the given []string and assigns it to the ResourceIds field.

### GetTags

`func (o *DeleteTagsRequest) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeleteTagsRequest) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *DeleteTagsRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *DeleteTagsRequest) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


