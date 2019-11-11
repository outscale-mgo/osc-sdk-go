# CreateVirtualGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | Pointer to **string** | The type of VPN connection supported by the virtual gateway (only &#x60;ipsec.1&#x60; is supported). | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 

## Methods

### GetConnectionType

`func (o *CreateVirtualGatewayRequest) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CreateVirtualGatewayRequest) GetConnectionTypeOk() (string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasConnectionType

`func (o *CreateVirtualGatewayRequest) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### SetConnectionType

`func (o *CreateVirtualGatewayRequest) SetConnectionType(v string)`

SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.

### GetDryRun

`func (o *CreateVirtualGatewayRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateVirtualGatewayRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateVirtualGatewayRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateVirtualGatewayRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


