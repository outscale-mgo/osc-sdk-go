# BlockDeviceMappingVmUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bsu** | Pointer to [**BsuToUpdateVm**](BsuToUpdateVm.md) |  | [optional] 
**DeviceName** | Pointer to **string** | The name of the device. | [optional] 
**NoDevice** | Pointer to **string** | Removes the device which is included in the block device mapping of the OMI. | [optional] 
**VirtualDeviceName** | Pointer to **string** | The name of the virtual device (ephemeralN). | [optional] 

## Methods

### GetBsu

`func (o *BlockDeviceMappingVmUpdate) GetBsu() BsuToUpdateVm`

GetBsu returns the Bsu field if non-nil, zero value otherwise.

### GetBsuOk

`func (o *BlockDeviceMappingVmUpdate) GetBsuOk() (BsuToUpdateVm, bool)`

GetBsuOk returns a tuple with the Bsu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBsu

`func (o *BlockDeviceMappingVmUpdate) HasBsu() bool`

HasBsu returns a boolean if a field has been set.

### SetBsu

`func (o *BlockDeviceMappingVmUpdate) SetBsu(v BsuToUpdateVm)`

SetBsu gets a reference to the given BsuToUpdateVm and assigns it to the Bsu field.

### GetDeviceName

`func (o *BlockDeviceMappingVmUpdate) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *BlockDeviceMappingVmUpdate) GetDeviceNameOk() (string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeviceName

`func (o *BlockDeviceMappingVmUpdate) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceName

`func (o *BlockDeviceMappingVmUpdate) SetDeviceName(v string)`

SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.

### GetNoDevice

`func (o *BlockDeviceMappingVmUpdate) GetNoDevice() string`

GetNoDevice returns the NoDevice field if non-nil, zero value otherwise.

### GetNoDeviceOk

`func (o *BlockDeviceMappingVmUpdate) GetNoDeviceOk() (string, bool)`

GetNoDeviceOk returns a tuple with the NoDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNoDevice

`func (o *BlockDeviceMappingVmUpdate) HasNoDevice() bool`

HasNoDevice returns a boolean if a field has been set.

### SetNoDevice

`func (o *BlockDeviceMappingVmUpdate) SetNoDevice(v string)`

SetNoDevice gets a reference to the given string and assigns it to the NoDevice field.

### GetVirtualDeviceName

`func (o *BlockDeviceMappingVmUpdate) GetVirtualDeviceName() string`

GetVirtualDeviceName returns the VirtualDeviceName field if non-nil, zero value otherwise.

### GetVirtualDeviceNameOk

`func (o *BlockDeviceMappingVmUpdate) GetVirtualDeviceNameOk() (string, bool)`

GetVirtualDeviceNameOk returns a tuple with the VirtualDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVirtualDeviceName

`func (o *BlockDeviceMappingVmUpdate) HasVirtualDeviceName() bool`

HasVirtualDeviceName returns a boolean if a field has been set.

### SetVirtualDeviceName

`func (o *BlockDeviceMappingVmUpdate) SetVirtualDeviceName(v string)`

SetVirtualDeviceName gets a reference to the given string and assigns it to the VirtualDeviceName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


