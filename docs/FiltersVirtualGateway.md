# FiltersVirtualGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionTypes** | Pointer to **[]string** | The types of the virtual gateways (only &#x60;ipsec.1&#x60; is supported). | [optional] 
**LinkNetIds** | Pointer to **[]string** | The IDs of the Nets the virtual gateways are attached to. | [optional] 
**LinkStates** | Pointer to **[]string** | The current states of the attachments between the virtual gateways and the Nets (&#x60;attaching&#x60; \\| &#x60;attached&#x60; \\| &#x60;detaching&#x60; \\| &#x60;detached&#x60;). | [optional] 
**States** | Pointer to **[]string** | The states of the virtual gateways (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**TagKeys** | Pointer to **[]string** | The keys of the tags associated with the virtual gateways. | [optional] 
**TagValues** | Pointer to **[]string** | The values of the tags associated with the virtual gateways. | [optional] 
**Tags** | Pointer to **[]string** | The key/value combination of the tags associated with the virtual gateways, in the following format: \&quot;Filters\&quot;:{\&quot;Tags\&quot;:[\&quot;TAGKEY&#x3D;TAGVALUE\&quot;]}. | [optional] 
**VirtualGatewayIds** | Pointer to **[]string** | The IDs of the virtual gateways. | [optional] 

## Methods

### GetConnectionTypes

`func (o *FiltersVirtualGateway) GetConnectionTypes() []string`

GetConnectionTypes returns the ConnectionTypes field if non-nil, zero value otherwise.

### GetConnectionTypesOk

`func (o *FiltersVirtualGateway) GetConnectionTypesOk() ([]string, bool)`

GetConnectionTypesOk returns a tuple with the ConnectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectionTypes

`func (o *FiltersVirtualGateway) HasConnectionTypes() bool`

HasConnectionTypes returns a boolean if a field has been set.

### SetConnectionTypes

`func (o *FiltersVirtualGateway) SetConnectionTypes(v []string)`

SetConnectionTypes gets a reference to the given []string and assigns it to the ConnectionTypes field.

### GetLinkNetIds

`func (o *FiltersVirtualGateway) GetLinkNetIds() []string`

GetLinkNetIds returns the LinkNetIds field if non-nil, zero value otherwise.

### GetLinkNetIdsOk

`func (o *FiltersVirtualGateway) GetLinkNetIdsOk() ([]string, bool)`

GetLinkNetIdsOk returns a tuple with the LinkNetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkNetIds

`func (o *FiltersVirtualGateway) HasLinkNetIds() bool`

HasLinkNetIds returns a boolean if a field has been set.

### SetLinkNetIds

`func (o *FiltersVirtualGateway) SetLinkNetIds(v []string)`

SetLinkNetIds gets a reference to the given []string and assigns it to the LinkNetIds field.

### GetLinkStates

`func (o *FiltersVirtualGateway) GetLinkStates() []string`

GetLinkStates returns the LinkStates field if non-nil, zero value otherwise.

### GetLinkStatesOk

`func (o *FiltersVirtualGateway) GetLinkStatesOk() ([]string, bool)`

GetLinkStatesOk returns a tuple with the LinkStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkStates

`func (o *FiltersVirtualGateway) HasLinkStates() bool`

HasLinkStates returns a boolean if a field has been set.

### SetLinkStates

`func (o *FiltersVirtualGateway) SetLinkStates(v []string)`

SetLinkStates gets a reference to the given []string and assigns it to the LinkStates field.

### GetStates

`func (o *FiltersVirtualGateway) GetStates() []string`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *FiltersVirtualGateway) GetStatesOk() ([]string, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStates

`func (o *FiltersVirtualGateway) HasStates() bool`

HasStates returns a boolean if a field has been set.

### SetStates

`func (o *FiltersVirtualGateway) SetStates(v []string)`

SetStates gets a reference to the given []string and assigns it to the States field.

### GetTagKeys

`func (o *FiltersVirtualGateway) GetTagKeys() []string`

GetTagKeys returns the TagKeys field if non-nil, zero value otherwise.

### GetTagKeysOk

`func (o *FiltersVirtualGateway) GetTagKeysOk() ([]string, bool)`

GetTagKeysOk returns a tuple with the TagKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagKeys

`func (o *FiltersVirtualGateway) HasTagKeys() bool`

HasTagKeys returns a boolean if a field has been set.

### SetTagKeys

`func (o *FiltersVirtualGateway) SetTagKeys(v []string)`

SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.

### GetTagValues

`func (o *FiltersVirtualGateway) GetTagValues() []string`

GetTagValues returns the TagValues field if non-nil, zero value otherwise.

### GetTagValuesOk

`func (o *FiltersVirtualGateway) GetTagValuesOk() ([]string, bool)`

GetTagValuesOk returns a tuple with the TagValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagValues

`func (o *FiltersVirtualGateway) HasTagValues() bool`

HasTagValues returns a boolean if a field has been set.

### SetTagValues

`func (o *FiltersVirtualGateway) SetTagValues(v []string)`

SetTagValues gets a reference to the given []string and assigns it to the TagValues field.

### GetTags

`func (o *FiltersVirtualGateway) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FiltersVirtualGateway) GetTagsOk() ([]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *FiltersVirtualGateway) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *FiltersVirtualGateway) SetTags(v []string)`

SetTags gets a reference to the given []string and assigns it to the Tags field.

### GetVirtualGatewayIds

`func (o *FiltersVirtualGateway) GetVirtualGatewayIds() []string`

GetVirtualGatewayIds returns the VirtualGatewayIds field if non-nil, zero value otherwise.

### GetVirtualGatewayIdsOk

`func (o *FiltersVirtualGateway) GetVirtualGatewayIdsOk() ([]string, bool)`

GetVirtualGatewayIdsOk returns a tuple with the VirtualGatewayIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVirtualGatewayIds

`func (o *FiltersVirtualGateway) HasVirtualGatewayIds() bool`

HasVirtualGatewayIds returns a boolean if a field has been set.

### SetVirtualGatewayIds

`func (o *FiltersVirtualGateway) SetVirtualGatewayIds(v []string)`

SetVirtualGatewayIds gets a reference to the given []string and assigns it to the VirtualGatewayIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


