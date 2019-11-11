# NetPeering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccepterNet** | Pointer to [**AccepterNet**](AccepterNet.md) |  | [optional] 
**NetPeeringId** | Pointer to **string** | The ID of the Net peering connection. | [optional] 
**SourceNet** | Pointer to [**SourceNet**](SourceNet.md) |  | [optional] 
**State** | Pointer to [**NetPeeringState**](NetPeeringState.md) |  | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the Net peering connection. | [optional] 

## Methods

### GetAccepterNet

`func (o *NetPeering) GetAccepterNet() AccepterNet`

GetAccepterNet returns the AccepterNet field if non-nil, zero value otherwise.

### GetAccepterNetOk

`func (o *NetPeering) GetAccepterNetOk() (AccepterNet, bool)`

GetAccepterNetOk returns a tuple with the AccepterNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccepterNet

`func (o *NetPeering) HasAccepterNet() bool`

HasAccepterNet returns a boolean if a field has been set.

### SetAccepterNet

`func (o *NetPeering) SetAccepterNet(v AccepterNet)`

SetAccepterNet gets a reference to the given AccepterNet and assigns it to the AccepterNet field.

### GetNetPeeringId

`func (o *NetPeering) GetNetPeeringId() string`

GetNetPeeringId returns the NetPeeringId field if non-nil, zero value otherwise.

### GetNetPeeringIdOk

`func (o *NetPeering) GetNetPeeringIdOk() (string, bool)`

GetNetPeeringIdOk returns a tuple with the NetPeeringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetPeeringId

`func (o *NetPeering) HasNetPeeringId() bool`

HasNetPeeringId returns a boolean if a field has been set.

### SetNetPeeringId

`func (o *NetPeering) SetNetPeeringId(v string)`

SetNetPeeringId gets a reference to the given string and assigns it to the NetPeeringId field.

### GetSourceNet

`func (o *NetPeering) GetSourceNet() SourceNet`

GetSourceNet returns the SourceNet field if non-nil, zero value otherwise.

### GetSourceNetOk

`func (o *NetPeering) GetSourceNetOk() (SourceNet, bool)`

GetSourceNetOk returns a tuple with the SourceNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSourceNet

`func (o *NetPeering) HasSourceNet() bool`

HasSourceNet returns a boolean if a field has been set.

### SetSourceNet

`func (o *NetPeering) SetSourceNet(v SourceNet)`

SetSourceNet gets a reference to the given SourceNet and assigns it to the SourceNet field.

### GetState

`func (o *NetPeering) GetState() NetPeeringState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NetPeering) GetStateOk() (NetPeeringState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *NetPeering) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *NetPeering) SetState(v NetPeeringState)`

SetState gets a reference to the given NetPeeringState and assigns it to the State field.

### GetTags

`func (o *NetPeering) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NetPeering) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *NetPeering) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *NetPeering) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


