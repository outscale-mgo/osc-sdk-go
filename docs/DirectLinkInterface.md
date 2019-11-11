# DirectLinkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAsn** | Pointer to **int32** | The BGP (Border Gateway Protocol) ASN (Autonomous System Number) on the customer&#39;s side of the DirectLink interface. | 
**BgpKey** | Pointer to **string** | The BGP authentication key. | [optional] 
**ClientPrivateIp** | Pointer to **string** | The IP address on the customer&#39;s side of the DirectLink interface. | [optional] 
**DirectLinkInterfaceName** | Pointer to **string** | The name of the DirectLink interface. | 
**OutscalePrivateIp** | Pointer to **string** | The IP address on 3DS OUTSCALE&#39;s side of the DirectLink interface. | [optional] 
**VirtualGatewayId** | Pointer to **string** | The ID of the target virtual gateway. | 
**Vlan** | Pointer to **int32** | The VLAN number associated with the DirectLink interface. | 

## Methods

### GetBgpAsn

`func (o *DirectLinkInterface) GetBgpAsn() int32`

GetBgpAsn returns the BgpAsn field if non-nil, zero value otherwise.

### GetBgpAsnOk

`func (o *DirectLinkInterface) GetBgpAsnOk() (int32, bool)`

GetBgpAsnOk returns a tuple with the BgpAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBgpAsn

`func (o *DirectLinkInterface) HasBgpAsn() bool`

HasBgpAsn returns a boolean if a field has been set.

### SetBgpAsn

`func (o *DirectLinkInterface) SetBgpAsn(v int32)`

SetBgpAsn gets a reference to the given int32 and assigns it to the BgpAsn field.

### GetBgpKey

`func (o *DirectLinkInterface) GetBgpKey() string`

GetBgpKey returns the BgpKey field if non-nil, zero value otherwise.

### GetBgpKeyOk

`func (o *DirectLinkInterface) GetBgpKeyOk() (string, bool)`

GetBgpKeyOk returns a tuple with the BgpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBgpKey

`func (o *DirectLinkInterface) HasBgpKey() bool`

HasBgpKey returns a boolean if a field has been set.

### SetBgpKey

`func (o *DirectLinkInterface) SetBgpKey(v string)`

SetBgpKey gets a reference to the given string and assigns it to the BgpKey field.

### GetClientPrivateIp

`func (o *DirectLinkInterface) GetClientPrivateIp() string`

GetClientPrivateIp returns the ClientPrivateIp field if non-nil, zero value otherwise.

### GetClientPrivateIpOk

`func (o *DirectLinkInterface) GetClientPrivateIpOk() (string, bool)`

GetClientPrivateIpOk returns a tuple with the ClientPrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientPrivateIp

`func (o *DirectLinkInterface) HasClientPrivateIp() bool`

HasClientPrivateIp returns a boolean if a field has been set.

### SetClientPrivateIp

`func (o *DirectLinkInterface) SetClientPrivateIp(v string)`

SetClientPrivateIp gets a reference to the given string and assigns it to the ClientPrivateIp field.

### GetDirectLinkInterfaceName

`func (o *DirectLinkInterface) GetDirectLinkInterfaceName() string`

GetDirectLinkInterfaceName returns the DirectLinkInterfaceName field if non-nil, zero value otherwise.

### GetDirectLinkInterfaceNameOk

`func (o *DirectLinkInterface) GetDirectLinkInterfaceNameOk() (string, bool)`

GetDirectLinkInterfaceNameOk returns a tuple with the DirectLinkInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDirectLinkInterfaceName

`func (o *DirectLinkInterface) HasDirectLinkInterfaceName() bool`

HasDirectLinkInterfaceName returns a boolean if a field has been set.

### SetDirectLinkInterfaceName

`func (o *DirectLinkInterface) SetDirectLinkInterfaceName(v string)`

SetDirectLinkInterfaceName gets a reference to the given string and assigns it to the DirectLinkInterfaceName field.

### GetOutscalePrivateIp

`func (o *DirectLinkInterface) GetOutscalePrivateIp() string`

GetOutscalePrivateIp returns the OutscalePrivateIp field if non-nil, zero value otherwise.

### GetOutscalePrivateIpOk

`func (o *DirectLinkInterface) GetOutscalePrivateIpOk() (string, bool)`

GetOutscalePrivateIpOk returns a tuple with the OutscalePrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOutscalePrivateIp

`func (o *DirectLinkInterface) HasOutscalePrivateIp() bool`

HasOutscalePrivateIp returns a boolean if a field has been set.

### SetOutscalePrivateIp

`func (o *DirectLinkInterface) SetOutscalePrivateIp(v string)`

SetOutscalePrivateIp gets a reference to the given string and assigns it to the OutscalePrivateIp field.

### GetVirtualGatewayId

`func (o *DirectLinkInterface) GetVirtualGatewayId() string`

GetVirtualGatewayId returns the VirtualGatewayId field if non-nil, zero value otherwise.

### GetVirtualGatewayIdOk

`func (o *DirectLinkInterface) GetVirtualGatewayIdOk() (string, bool)`

GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVirtualGatewayId

`func (o *DirectLinkInterface) HasVirtualGatewayId() bool`

HasVirtualGatewayId returns a boolean if a field has been set.

### SetVirtualGatewayId

`func (o *DirectLinkInterface) SetVirtualGatewayId(v string)`

SetVirtualGatewayId gets a reference to the given string and assigns it to the VirtualGatewayId field.

### GetVlan

`func (o *DirectLinkInterface) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *DirectLinkInterface) GetVlanOk() (int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVlan

`func (o *DirectLinkInterface) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlan

`func (o *DirectLinkInterface) SetVlan(v int32)`

SetVlan gets a reference to the given int32 and assigns it to the Vlan field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


