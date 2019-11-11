# BlockDeviceMappingImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bsu** | Pointer to [**BsuToCreate**](BsuToCreate.md) |  | [optional] 
**DeviceName** | Pointer to **string** | The name of the device. | [optional] 
**VirtualDeviceName** | Pointer to **string** | The name of the virtual device (ephemeralN). | [optional] 

## Methods

### GetBsu

`func (o *BlockDeviceMappingImage) GetBsu() BsuToCreate`

GetBsu returns the Bsu field if non-nil, zero value otherwise.

### GetBsuOk

`func (o *BlockDeviceMappingImage) GetBsuOk() (BsuToCreate, bool)`

GetBsuOk returns a tuple with the Bsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBsu

`func (o *BlockDeviceMappingImage) HasBsu() bool`

HasBsu returns a boolean if a field has been set.

### SetBsu

`func (o *BlockDeviceMappingImage) SetBsu(v BsuToCreate)`

SetBsu gets a reference to the given BsuToCreate and assigns it to the Bsu field.

### GetDeviceName

`func (o *BlockDeviceMappingImage) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *BlockDeviceMappingImage) GetDeviceNameOk() (string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceName

`func (o *BlockDeviceMappingImage) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceName

`func (o *BlockDeviceMappingImage) SetDeviceName(v string)`

SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.

### GetVirtualDeviceName

`func (o *BlockDeviceMappingImage) GetVirtualDeviceName() string`

GetVirtualDeviceName returns the VirtualDeviceName field if non-nil, zero value otherwise.

### GetVirtualDeviceNameOk

`func (o *BlockDeviceMappingImage) GetVirtualDeviceNameOk() (string, bool)`

GetVirtualDeviceNameOk returns a tuple with the VirtualDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVirtualDeviceName

`func (o *BlockDeviceMappingImage) HasVirtualDeviceName() bool`

HasVirtualDeviceName returns a boolean if a field has been set.

### SetVirtualDeviceName

`func (o *BlockDeviceMappingImage) SetVirtualDeviceName(v string)`

SetVirtualDeviceName gets a reference to the given string and assigns it to the VirtualDeviceName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


