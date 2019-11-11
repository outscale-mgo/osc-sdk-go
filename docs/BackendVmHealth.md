# BackendVmHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the state of the back-end VM. | [optional] 
**State** | Pointer to **string** | The state of the back-end VM (&#x60;InService&#x60; \\| &#x60;OutOfService&#x60; \\| &#x60;Unknown&#x60;). | [optional] 
**StateReason** | Pointer to **string** | Information about the cause of &#x60;OutOfService&#x60; VMs.&lt;br /&gt; Specifically, whether the cause is Elastic Load Balancing or the VM (&#x60;ELB&#x60; \\| &#x60;Instance&#x60; \\| &#x60;N/A&#x60;). | [optional] 
**VmId** | Pointer to **string** | The ID of the back-end VM. | [optional] 

## Methods

### GetDescription

`func (o *BackendVmHealth) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BackendVmHealth) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *BackendVmHealth) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *BackendVmHealth) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetState

`func (o *BackendVmHealth) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BackendVmHealth) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *BackendVmHealth) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *BackendVmHealth) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetStateReason

`func (o *BackendVmHealth) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *BackendVmHealth) GetStateReasonOk() (string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStateReason

`func (o *BackendVmHealth) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### SetStateReason

`func (o *BackendVmHealth) SetStateReason(v string)`

SetStateReason gets a reference to the given string and assigns it to the StateReason field.

### GetVmId

`func (o *BackendVmHealth) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *BackendVmHealth) GetVmIdOk() (string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmId

`func (o *BackendVmHealth) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### SetVmId

`func (o *BackendVmHealth) SetVmId(v string)`

SetVmId gets a reference to the given string and assigns it to the VmId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


