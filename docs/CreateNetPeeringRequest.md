# CreateNetPeeringRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccepterNetId** | Pointer to **string** | The ID of the Net you want to connect with. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**SourceNetId** | Pointer to **string** | The ID of the Net you send the peering request from. | 

## Methods

### GetAccepterNetId

`func (o *CreateNetPeeringRequest) GetAccepterNetId() string`

GetAccepterNetId returns the AccepterNetId field if non-nil, zero value otherwise.

### GetAccepterNetIdOk

`func (o *CreateNetPeeringRequest) GetAccepterNetIdOk() (string, bool)`

GetAccepterNetIdOk returns a tuple with the AccepterNetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccepterNetId

`func (o *CreateNetPeeringRequest) HasAccepterNetId() bool`

HasAccepterNetId returns a boolean if a field has been set.

### SetAccepterNetId

`func (o *CreateNetPeeringRequest) SetAccepterNetId(v string)`

SetAccepterNetId gets a reference to the given string and assigns it to the AccepterNetId field.

### GetDryRun

`func (o *CreateNetPeeringRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateNetPeeringRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateNetPeeringRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateNetPeeringRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetSourceNetId

`func (o *CreateNetPeeringRequest) GetSourceNetId() string`

GetSourceNetId returns the SourceNetId field if non-nil, zero value otherwise.

### GetSourceNetIdOk

`func (o *CreateNetPeeringRequest) GetSourceNetIdOk() (string, bool)`

GetSourceNetIdOk returns a tuple with the SourceNetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceNetId

`func (o *CreateNetPeeringRequest) HasSourceNetId() bool`

HasSourceNetId returns a boolean if a field has been set.

### SetSourceNetId

`func (o *CreateNetPeeringRequest) SetSourceNetId(v string)`

SetSourceNetId gets a reference to the given string and assigns it to the SourceNetId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


