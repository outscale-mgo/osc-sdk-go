# NetPeeringState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Additional information about the state of the Net peering connection. | [optional] 
**Name** | Pointer to **string** | The state of the Net peering connection (&#x60;pending-acceptance&#x60; \\| &#x60;active&#x60; \\| &#x60;rejected&#x60; \\| &#x60;failed&#x60; \\| &#x60;expired&#x60; \\| &#x60;deleted&#x60;). | [optional] 

## Methods

### GetMessage

`func (o *NetPeeringState) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NetPeeringState) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *NetPeeringState) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *NetPeeringState) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### GetName

`func (o *NetPeeringState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetPeeringState) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *NetPeeringState) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *NetPeeringState) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


