# DeleteClientGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientGatewayId** | Pointer to **string** | The ID of the client gateway you want to delete. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 

## Methods

### GetClientGatewayId

`func (o *DeleteClientGatewayRequest) GetClientGatewayId() string`

GetClientGatewayId returns the ClientGatewayId field if non-nil, zero value otherwise.

### GetClientGatewayIdOk

`func (o *DeleteClientGatewayRequest) GetClientGatewayIdOk() (string, bool)`

GetClientGatewayIdOk returns a tuple with the ClientGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientGatewayId

`func (o *DeleteClientGatewayRequest) HasClientGatewayId() bool`

HasClientGatewayId returns a boolean if a field has been set.

### SetClientGatewayId

`func (o *DeleteClientGatewayRequest) SetClientGatewayId(v string)`

SetClientGatewayId gets a reference to the given string and assigns it to the ClientGatewayId field.

### GetDryRun

`func (o *DeleteClientGatewayRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteClientGatewayRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteClientGatewayRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteClientGatewayRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


