# CreateDirectLinkInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectLinkId** | Pointer to **string** | The ID of the existing DirectLink for which you want to create the DirectLink interface. | 
**DirectLinkInterface** | Pointer to [**DirectLinkInterface**](DirectLinkInterface.md) |  | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 

## Methods

### GetDirectLinkId

`func (o *CreateDirectLinkInterfaceRequest) GetDirectLinkId() string`

GetDirectLinkId returns the DirectLinkId field if non-nil, zero value otherwise.

### GetDirectLinkIdOk

`func (o *CreateDirectLinkInterfaceRequest) GetDirectLinkIdOk() (string, bool)`

GetDirectLinkIdOk returns a tuple with the DirectLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDirectLinkId

`func (o *CreateDirectLinkInterfaceRequest) HasDirectLinkId() bool`

HasDirectLinkId returns a boolean if a field has been set.

### SetDirectLinkId

`func (o *CreateDirectLinkInterfaceRequest) SetDirectLinkId(v string)`

SetDirectLinkId gets a reference to the given string and assigns it to the DirectLinkId field.

### GetDirectLinkInterface

`func (o *CreateDirectLinkInterfaceRequest) GetDirectLinkInterface() DirectLinkInterface`

GetDirectLinkInterface returns the DirectLinkInterface field if non-nil, zero value otherwise.

### GetDirectLinkInterfaceOk

`func (o *CreateDirectLinkInterfaceRequest) GetDirectLinkInterfaceOk() (DirectLinkInterface, bool)`

GetDirectLinkInterfaceOk returns a tuple with the DirectLinkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDirectLinkInterface

`func (o *CreateDirectLinkInterfaceRequest) HasDirectLinkInterface() bool`

HasDirectLinkInterface returns a boolean if a field has been set.

### SetDirectLinkInterface

`func (o *CreateDirectLinkInterfaceRequest) SetDirectLinkInterface(v DirectLinkInterface)`

SetDirectLinkInterface gets a reference to the given DirectLinkInterface and assigns it to the DirectLinkInterface field.

### GetDryRun

`func (o *CreateDirectLinkInterfaceRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateDirectLinkInterfaceRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateDirectLinkInterfaceRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateDirectLinkInterfaceRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


