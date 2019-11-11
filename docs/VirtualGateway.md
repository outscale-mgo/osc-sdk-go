# VirtualGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | Pointer to **string** | The type of VPN connection supported by the virtual gateway (only &#x60;ipsec.1&#x60; is supported). | [optional] 
**NetToVirtualGatewayLinks** | Pointer to [**[]NetToVirtualGatewayLink**](NetToVirtualGatewayLink.md) | The Net to which the virtual gateway is attached. | [optional] 
**State** | Pointer to **string** | The state of the virtual gateway (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60;). | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the virtual gateway. | [optional] 
**VirtualGatewayId** | Pointer to **string** | The ID of the virtual gateway. | [optional] 

## Methods

### GetConnectionType

`func (o *VirtualGateway) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *VirtualGateway) GetConnectionTypeOk() (string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectionType

`func (o *VirtualGateway) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### SetConnectionType

`func (o *VirtualGateway) SetConnectionType(v string)`

SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.

### GetNetToVirtualGatewayLinks

`func (o *VirtualGateway) GetNetToVirtualGatewayLinks() []NetToVirtualGatewayLink`

GetNetToVirtualGatewayLinks returns the NetToVirtualGatewayLinks field if non-nil, zero value otherwise.

### GetNetToVirtualGatewayLinksOk

`func (o *VirtualGateway) GetNetToVirtualGatewayLinksOk() ([]NetToVirtualGatewayLink, bool)`

GetNetToVirtualGatewayLinksOk returns a tuple with the NetToVirtualGatewayLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetToVirtualGatewayLinks

`func (o *VirtualGateway) HasNetToVirtualGatewayLinks() bool`

HasNetToVirtualGatewayLinks returns a boolean if a field has been set.

### SetNetToVirtualGatewayLinks

`func (o *VirtualGateway) SetNetToVirtualGatewayLinks(v []NetToVirtualGatewayLink)`

SetNetToVirtualGatewayLinks gets a reference to the given []NetToVirtualGatewayLink and assigns it to the NetToVirtualGatewayLinks field.

### GetState

`func (o *VirtualGateway) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VirtualGateway) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *VirtualGateway) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *VirtualGateway) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetTags

`func (o *VirtualGateway) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualGateway) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *VirtualGateway) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *VirtualGateway) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.

### GetVirtualGatewayId

`func (o *VirtualGateway) GetVirtualGatewayId() string`

GetVirtualGatewayId returns the VirtualGatewayId field if non-nil, zero value otherwise.

### GetVirtualGatewayIdOk

`func (o *VirtualGateway) GetVirtualGatewayIdOk() (string, bool)`

GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVirtualGatewayId

`func (o *VirtualGateway) HasVirtualGatewayId() bool`

HasVirtualGatewayId returns a boolean if a field has been set.

### SetVirtualGatewayId

`func (o *VirtualGateway) SetVirtualGatewayId(v string)`

SetVirtualGatewayId gets a reference to the given string and assigns it to the VirtualGatewayId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


