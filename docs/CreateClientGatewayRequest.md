# CreateClientGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpAsn** | Pointer to **int64** | An unsigned 32-bits Autonomous System Number (ASN) used by the Border Gateway Protocol (BGP) to find out the path to your client gateway through the Internet network. The long must be within the [0;4294967295] range. By default, 65000. | 
**ConnectionType** | Pointer to **string** | The communication protocol used to establish tunnel with your client gateway (only &#x60;ipsec.1&#x60; is supported). | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**PublicIp** | Pointer to **string** | The public fixed IPv4 address of your client gateway. | 

## Methods

### GetBgpAsn

`func (o *CreateClientGatewayRequest) GetBgpAsn() int64`

GetBgpAsn returns the BgpAsn field if non-nil, zero value otherwise.

### GetBgpAsnOk

`func (o *CreateClientGatewayRequest) GetBgpAsnOk() (int64, bool)`

GetBgpAsnOk returns a tuple with the BgpAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBgpAsn

`func (o *CreateClientGatewayRequest) HasBgpAsn() bool`

HasBgpAsn returns a boolean if a field has been set.

### SetBgpAsn

`func (o *CreateClientGatewayRequest) SetBgpAsn(v int64)`

SetBgpAsn gets a reference to the given int64 and assigns it to the BgpAsn field.

### GetConnectionType

`func (o *CreateClientGatewayRequest) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CreateClientGatewayRequest) GetConnectionTypeOk() (string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectionType

`func (o *CreateClientGatewayRequest) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### SetConnectionType

`func (o *CreateClientGatewayRequest) SetConnectionType(v string)`

SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.

### GetDryRun

`func (o *CreateClientGatewayRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateClientGatewayRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateClientGatewayRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateClientGatewayRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetPublicIp

`func (o *CreateClientGatewayRequest) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *CreateClientGatewayRequest) GetPublicIpOk() (string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicIp

`func (o *CreateClientGatewayRequest) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### SetPublicIp

`func (o *CreateClientGatewayRequest) SetPublicIp(v string)`

SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


