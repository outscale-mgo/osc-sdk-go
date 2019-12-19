# ClientGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAsn** | Pointer to **int64** | An unsigned 32-bits Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find out the path to your client gateway through the Internet network. | [optional] 
**ClientGatewayId** | Pointer to **string** | The ID of the client gateway. | [optional] 
**ConnectionType** | Pointer to **string** | The type of communication tunnel used by the client gateway (only &#x60;ipsec.1&#x60; is supported). | [optional] 
**PublicIp** | Pointer to **string** | The public IPv4 address of the client gateway (must be a fixed address into a NATed network). | [optional] 
**State** | Pointer to **string** | The state of the client gateway (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the client gateway. | [optional] 

## Methods

### GetBgpAsn

`func (o *ClientGateway) GetBgpAsn() int64`

GetBgpAsn returns the BgpAsn field if non-nil, zero value otherwise.

### GetBgpAsnOk

`func (o *ClientGateway) GetBgpAsnOk() (int64, bool)`

GetBgpAsnOk returns a tuple with the BgpAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBgpAsn

`func (o *ClientGateway) HasBgpAsn() bool`

HasBgpAsn returns a boolean if a field has been set.

### SetBgpAsn

`func (o *ClientGateway) SetBgpAsn(v int64)`

SetBgpAsn gets a reference to the given int64 and assigns it to the BgpAsn field.

### GetClientGatewayId

`func (o *ClientGateway) GetClientGatewayId() string`

GetClientGatewayId returns the ClientGatewayId field if non-nil, zero value otherwise.

### GetClientGatewayIdOk

`func (o *ClientGateway) GetClientGatewayIdOk() (string, bool)`

GetClientGatewayIdOk returns a tuple with the ClientGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientGatewayId

`func (o *ClientGateway) HasClientGatewayId() bool`

HasClientGatewayId returns a boolean if a field has been set.

### SetClientGatewayId

`func (o *ClientGateway) SetClientGatewayId(v string)`

SetClientGatewayId gets a reference to the given string and assigns it to the ClientGatewayId field.

### GetConnectionType

`func (o *ClientGateway) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *ClientGateway) GetConnectionTypeOk() (string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectionType

`func (o *ClientGateway) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### SetConnectionType

`func (o *ClientGateway) SetConnectionType(v string)`

SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.

### GetPublicIp

`func (o *ClientGateway) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *ClientGateway) GetPublicIpOk() (string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicIp

`func (o *ClientGateway) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### SetPublicIp

`func (o *ClientGateway) SetPublicIp(v string)`

SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.

### GetState

`func (o *ClientGateway) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ClientGateway) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *ClientGateway) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *ClientGateway) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetTags

`func (o *ClientGateway) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ClientGateway) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *ClientGateway) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *ClientGateway) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


