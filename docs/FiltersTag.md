# FiltersTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to **[]string** | The keys of the tags that are assigned to the resources. You can use this filter alongside the &#x60;Values&#x60; filter. In that case, you filter the resources corresponding to each tag, regardless of the other filter. | [optional] 
**ResourceIds** | Pointer to **[]string** | The IDs of the resources with which the tags are associated. | [optional] 
**ResourceTypes** | Pointer to **[]string** | The resource type (&#x60;instance&#x60; \\| &#x60;image&#x60; \\| &#x60;volume&#x60; \\| &#x60;snapshot&#x60; \\| &#x60;public-ip&#x60; \\| &#x60;security-group&#x60; \\| &#x60;route-table&#x60; \\| &#x60;network-interface&#x60; \\| &#x60;vpc&#x60; \\| &#x60;subnet&#x60; \\| &#x60;network-link&#x60; \\| &#x60;vpc-endpoint&#x60; \\| &#x60;nat-gateway&#x60; \\| &#x60;internet-gateway&#x60; \\| &#x60;customer-gateway&#x60; \\| &#x60;vpn-gateway&#x60; \\| &#x60;vpn-connection&#x60; \\| &#x60;dhcp-options&#x60; \\| &#x60;task&#x60;). | [optional] 
**Values** | Pointer to **[]string** | The values of the tags that are assigned to the resources. You can use this filter alongside the &#x60;TagKeys&#x60; filter. In that case, you filter the resources corresponding to each tag, regardless of the other filter. | [optional] 

## Methods

### GetKeys

`func (o *FiltersTag) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *FiltersTag) GetKeysOk() ([]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeys

`func (o *FiltersTag) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### SetKeys

`func (o *FiltersTag) SetKeys(v []string)`

SetKeys gets a reference to the given []string and assigns it to the Keys field.

### GetResourceIds

`func (o *FiltersTag) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *FiltersTag) GetResourceIdsOk() ([]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceIds

`func (o *FiltersTag) HasResourceIds() bool`

HasResourceIds returns a boolean if a field has been set.

### SetResourceIds

`func (o *FiltersTag) SetResourceIds(v []string)`

SetResourceIds gets a reference to the given []string and assigns it to the ResourceIds field.

### GetResourceTypes

`func (o *FiltersTag) GetResourceTypes() []string`

GetResourceTypes returns the ResourceTypes field if non-nil, zero value otherwise.

### GetResourceTypesOk

`func (o *FiltersTag) GetResourceTypesOk() ([]string, bool)`

GetResourceTypesOk returns a tuple with the ResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResourceTypes

`func (o *FiltersTag) HasResourceTypes() bool`

HasResourceTypes returns a boolean if a field has been set.

### SetResourceTypes

`func (o *FiltersTag) SetResourceTypes(v []string)`

SetResourceTypes gets a reference to the given []string and assigns it to the ResourceTypes field.

### GetValues

`func (o *FiltersTag) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FiltersTag) GetValuesOk() ([]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValues

`func (o *FiltersTag) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValues

`func (o *FiltersTag) SetValues(v []string)`

SetValues gets a reference to the given []string and assigns it to the Values field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


