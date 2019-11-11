# DeleteDirectLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectLinkId** | Pointer to **string** | The ID of the DirectLink you want to delete. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 

## Methods

### GetDirectLinkId

`func (o *DeleteDirectLinkRequest) GetDirectLinkId() string`

GetDirectLinkId returns the DirectLinkId field if non-nil, zero value otherwise.

### GetDirectLinkIdOk

`func (o *DeleteDirectLinkRequest) GetDirectLinkIdOk() (string, bool)`

GetDirectLinkIdOk returns a tuple with the DirectLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDirectLinkId

`func (o *DeleteDirectLinkRequest) HasDirectLinkId() bool`

HasDirectLinkId returns a boolean if a field has been set.

### SetDirectLinkId

`func (o *DeleteDirectLinkRequest) SetDirectLinkId(v string)`

SetDirectLinkId gets a reference to the given string and assigns it to the DirectLinkId field.

### GetDryRun

`func (o *DeleteDirectLinkRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteDirectLinkRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteDirectLinkRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteDirectLinkRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


