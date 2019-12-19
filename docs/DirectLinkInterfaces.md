# DirectLinkInterfaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of the owner of the DirectLink interface. | [optional] 
**BgpAsn** | Pointer to **int64** | The BGP (Border Gateway Protocol) ASN (Autonomous System Number) on the customer&#39;s side of the DirectLink interface. | [optional] 
**BgpKey** | Pointer to **string** | The BGP authentication key. | [optional] 
**ClientPrivateIp** | Pointer to **string** | The IP address on the customer&#39;s side of the DirectLink interface. | [optional] 
**DirectLinkId** | Pointer to **string** | The ID of the DirectLink. | [optional] 
**DirectLinkInterfaceId** | Pointer to **string** | The ID of the DirectLink interface. | [optional] 
**DirectLinkInterfaceName** | Pointer to **string** | The name of the DirectLink interface. | [optional] 
**InterfaceType** | Pointer to **string** | The type of the DirectLink interface (always &#x60;private&#x60;). | [optional] 
**Location** | Pointer to **string** | The datacenter where the DirectLink interface is located. | [optional] 
**OutscalePrivateIp** | Pointer to **string** | The IP address on 3DS OUTSCALE&#39;s side of the DirectLink interface. | [optional] 
**State** | Pointer to **string** | The state of the DirectLink interface (&#x60;pending&#x60; \\| &#x60;available&#x60; \\| &#x60;deleting&#x60; \\| &#x60;deleted&#x60; \\| &#x60;confirming&#x60; \\| &#x60;rejected&#x60; \\| &#x60;expired&#x60;). | [optional] 
**VirtualGatewayId** | Pointer to **string** | The ID of the target virtual gateway. | [optional] 
**Vlan** | Pointer to **int64** | The VLAN number associated with the DirectLink interface. | [optional] 

## Methods

### GetAccountId

`func (o *DirectLinkInterfaces) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *DirectLinkInterfaces) GetAccountIdOk() (string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *DirectLinkInterfaces) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *DirectLinkInterfaces) SetAccountId(v string)`

SetAccountId gets a reference to the given string and assigns it to the AccountId field.

### GetBgpAsn

`func (o *DirectLinkInterfaces) GetBgpAsn() int64`

GetBgpAsn returns the BgpAsn field if non-nil, zero value otherwise.

### GetBgpAsnOk

`func (o *DirectLinkInterfaces) GetBgpAsnOk() (int64, bool)`

GetBgpAsnOk returns a tuple with the BgpAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBgpAsn

`func (o *DirectLinkInterfaces) HasBgpAsn() bool`

HasBgpAsn returns a boolean if a field has been set.

### SetBgpAsn

`func (o *DirectLinkInterfaces) SetBgpAsn(v int64)`

SetBgpAsn gets a reference to the given int64 and assigns it to the BgpAsn field.

### GetBgpKey

`func (o *DirectLinkInterfaces) GetBgpKey() string`

GetBgpKey returns the BgpKey field if non-nil, zero value otherwise.

### GetBgpKeyOk

`func (o *DirectLinkInterfaces) GetBgpKeyOk() (string, bool)`

GetBgpKeyOk returns a tuple with the BgpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBgpKey

`func (o *DirectLinkInterfaces) HasBgpKey() bool`

HasBgpKey returns a boolean if a field has been set.

### SetBgpKey

`func (o *DirectLinkInterfaces) SetBgpKey(v string)`

SetBgpKey gets a reference to the given string and assigns it to the BgpKey field.

### GetClientPrivateIp

`func (o *DirectLinkInterfaces) GetClientPrivateIp() string`

GetClientPrivateIp returns the ClientPrivateIp field if non-nil, zero value otherwise.

### GetClientPrivateIpOk

`func (o *DirectLinkInterfaces) GetClientPrivateIpOk() (string, bool)`

GetClientPrivateIpOk returns a tuple with the ClientPrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientPrivateIp

`func (o *DirectLinkInterfaces) HasClientPrivateIp() bool`

HasClientPrivateIp returns a boolean if a field has been set.

### SetClientPrivateIp

`func (o *DirectLinkInterfaces) SetClientPrivateIp(v string)`

SetClientPrivateIp gets a reference to the given string and assigns it to the ClientPrivateIp field.

### GetDirectLinkId

`func (o *DirectLinkInterfaces) GetDirectLinkId() string`

GetDirectLinkId returns the DirectLinkId field if non-nil, zero value otherwise.

### GetDirectLinkIdOk

`func (o *DirectLinkInterfaces) GetDirectLinkIdOk() (string, bool)`

GetDirectLinkIdOk returns a tuple with the DirectLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDirectLinkId

`func (o *DirectLinkInterfaces) HasDirectLinkId() bool`

HasDirectLinkId returns a boolean if a field has been set.

### SetDirectLinkId

`func (o *DirectLinkInterfaces) SetDirectLinkId(v string)`

SetDirectLinkId gets a reference to the given string and assigns it to the DirectLinkId field.

### GetDirectLinkInterfaceId

`func (o *DirectLinkInterfaces) GetDirectLinkInterfaceId() string`

GetDirectLinkInterfaceId returns the DirectLinkInterfaceId field if non-nil, zero value otherwise.

### GetDirectLinkInterfaceIdOk

`func (o *DirectLinkInterfaces) GetDirectLinkInterfaceIdOk() (string, bool)`

GetDirectLinkInterfaceIdOk returns a tuple with the DirectLinkInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDirectLinkInterfaceId

`func (o *DirectLinkInterfaces) HasDirectLinkInterfaceId() bool`

HasDirectLinkInterfaceId returns a boolean if a field has been set.

### SetDirectLinkInterfaceId

`func (o *DirectLinkInterfaces) SetDirectLinkInterfaceId(v string)`

SetDirectLinkInterfaceId gets a reference to the given string and assigns it to the DirectLinkInterfaceId field.

### GetDirectLinkInterfaceName

`func (o *DirectLinkInterfaces) GetDirectLinkInterfaceName() string`

GetDirectLinkInterfaceName returns the DirectLinkInterfaceName field if non-nil, zero value otherwise.

### GetDirectLinkInterfaceNameOk

`func (o *DirectLinkInterfaces) GetDirectLinkInterfaceNameOk() (string, bool)`

GetDirectLinkInterfaceNameOk returns a tuple with the DirectLinkInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDirectLinkInterfaceName

`func (o *DirectLinkInterfaces) HasDirectLinkInterfaceName() bool`

HasDirectLinkInterfaceName returns a boolean if a field has been set.

### SetDirectLinkInterfaceName

`func (o *DirectLinkInterfaces) SetDirectLinkInterfaceName(v string)`

SetDirectLinkInterfaceName gets a reference to the given string and assigns it to the DirectLinkInterfaceName field.

### GetInterfaceType

`func (o *DirectLinkInterfaces) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *DirectLinkInterfaces) GetInterfaceTypeOk() (string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasInterfaceType

`func (o *DirectLinkInterfaces) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### SetInterfaceType

`func (o *DirectLinkInterfaces) SetInterfaceType(v string)`

SetInterfaceType gets a reference to the given string and assigns it to the InterfaceType field.

### GetLocation

`func (o *DirectLinkInterfaces) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DirectLinkInterfaces) GetLocationOk() (string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocation

`func (o *DirectLinkInterfaces) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocation

`func (o *DirectLinkInterfaces) SetLocation(v string)`

SetLocation gets a reference to the given string and assigns it to the Location field.

### GetOutscalePrivateIp

`func (o *DirectLinkInterfaces) GetOutscalePrivateIp() string`

GetOutscalePrivateIp returns the OutscalePrivateIp field if non-nil, zero value otherwise.

### GetOutscalePrivateIpOk

`func (o *DirectLinkInterfaces) GetOutscalePrivateIpOk() (string, bool)`

GetOutscalePrivateIpOk returns a tuple with the OutscalePrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOutscalePrivateIp

`func (o *DirectLinkInterfaces) HasOutscalePrivateIp() bool`

HasOutscalePrivateIp returns a boolean if a field has been set.

### SetOutscalePrivateIp

`func (o *DirectLinkInterfaces) SetOutscalePrivateIp(v string)`

SetOutscalePrivateIp gets a reference to the given string and assigns it to the OutscalePrivateIp field.

### GetState

`func (o *DirectLinkInterfaces) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DirectLinkInterfaces) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *DirectLinkInterfaces) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *DirectLinkInterfaces) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetVirtualGatewayId

`func (o *DirectLinkInterfaces) GetVirtualGatewayId() string`

GetVirtualGatewayId returns the VirtualGatewayId field if non-nil, zero value otherwise.

### GetVirtualGatewayIdOk

`func (o *DirectLinkInterfaces) GetVirtualGatewayIdOk() (string, bool)`

GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVirtualGatewayId

`func (o *DirectLinkInterfaces) HasVirtualGatewayId() bool`

HasVirtualGatewayId returns a boolean if a field has been set.

### SetVirtualGatewayId

`func (o *DirectLinkInterfaces) SetVirtualGatewayId(v string)`

SetVirtualGatewayId gets a reference to the given string and assigns it to the VirtualGatewayId field.

### GetVlan

`func (o *DirectLinkInterfaces) GetVlan() int64`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *DirectLinkInterfaces) GetVlanOk() (int64, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVlan

`func (o *DirectLinkInterfaces) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlan

`func (o *DirectLinkInterfaces) SetVlan(v int64)`

SetVlan gets a reference to the given int64 and assigns it to the Vlan field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


