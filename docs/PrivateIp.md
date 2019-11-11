# PrivateIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPrimary** | Pointer to **bool** | If &#x60;true&#x60;, the IP address is the primary private IP address of the NIC. | [optional] 
**LinkPublicIp** | Pointer to [**LinkPublicIp**](LinkPublicIp.md) |  | [optional] 
**PrivateDnsName** | Pointer to **string** | The name of the private DNS. | [optional] 
**PrivateIp** | Pointer to **string** | The private IP address of the NIC. | [optional] 

## Methods

### GetIsPrimary

`func (o *PrivateIp) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *PrivateIp) GetIsPrimaryOk() (bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsPrimary

`func (o *PrivateIp) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### SetIsPrimary

`func (o *PrivateIp) SetIsPrimary(v bool)`

SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.

### GetLinkPublicIp

`func (o *PrivateIp) GetLinkPublicIp() LinkPublicIp`

GetLinkPublicIp returns the LinkPublicIp field if non-nil, zero value otherwise.

### GetLinkPublicIpOk

`func (o *PrivateIp) GetLinkPublicIpOk() (LinkPublicIp, bool)`

GetLinkPublicIpOk returns a tuple with the LinkPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLinkPublicIp

`func (o *PrivateIp) HasLinkPublicIp() bool`

HasLinkPublicIp returns a boolean if a field has been set.

### SetLinkPublicIp

`func (o *PrivateIp) SetLinkPublicIp(v LinkPublicIp)`

SetLinkPublicIp gets a reference to the given LinkPublicIp and assigns it to the LinkPublicIp field.

### GetPrivateDnsName

`func (o *PrivateIp) GetPrivateDnsName() string`

GetPrivateDnsName returns the PrivateDnsName field if non-nil, zero value otherwise.

### GetPrivateDnsNameOk

`func (o *PrivateIp) GetPrivateDnsNameOk() (string, bool)`

GetPrivateDnsNameOk returns a tuple with the PrivateDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateDnsName

`func (o *PrivateIp) HasPrivateDnsName() bool`

HasPrivateDnsName returns a boolean if a field has been set.

### SetPrivateDnsName

`func (o *PrivateIp) SetPrivateDnsName(v string)`

SetPrivateDnsName gets a reference to the given string and assigns it to the PrivateDnsName field.

### GetPrivateIp

`func (o *PrivateIp) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *PrivateIp) GetPrivateIpOk() (string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateIp

`func (o *PrivateIp) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIp

`func (o *PrivateIp) SetPrivateIp(v string)`

SetPrivateIp gets a reference to the given string and assigns it to the PrivateIp field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


