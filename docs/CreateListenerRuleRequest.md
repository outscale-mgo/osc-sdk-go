# CreateListenerRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Listener** | Pointer to [**LoadBalancerLight**](LoadBalancerLight.md) |  | 
**ListenerRule** | Pointer to [**ListenerRuleForCreation**](ListenerRuleForCreation.md) |  | 
**VmIds** | Pointer to **[]string** | The IDs of the backend VMs. | 

## Methods

### GetDryRun

`func (o *CreateListenerRuleRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateListenerRuleRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateListenerRuleRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateListenerRuleRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetListener

`func (o *CreateListenerRuleRequest) GetListener() LoadBalancerLight`

GetListener returns the Listener field if non-nil, zero value otherwise.

### GetListenerOk

`func (o *CreateListenerRuleRequest) GetListenerOk() (LoadBalancerLight, bool)`

GetListenerOk returns a tuple with the Listener field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasListener

`func (o *CreateListenerRuleRequest) HasListener() bool`

HasListener returns a boolean if a field has been set.

### SetListener

`func (o *CreateListenerRuleRequest) SetListener(v LoadBalancerLight)`

SetListener gets a reference to the given LoadBalancerLight and assigns it to the Listener field.

### GetListenerRule

`func (o *CreateListenerRuleRequest) GetListenerRule() ListenerRuleForCreation`

GetListenerRule returns the ListenerRule field if non-nil, zero value otherwise.

### GetListenerRuleOk

`func (o *CreateListenerRuleRequest) GetListenerRuleOk() (ListenerRuleForCreation, bool)`

GetListenerRuleOk returns a tuple with the ListenerRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasListenerRule

`func (o *CreateListenerRuleRequest) HasListenerRule() bool`

HasListenerRule returns a boolean if a field has been set.

### SetListenerRule

`func (o *CreateListenerRuleRequest) SetListenerRule(v ListenerRuleForCreation)`

SetListenerRule gets a reference to the given ListenerRuleForCreation and assigns it to the ListenerRule field.

### GetVmIds

`func (o *CreateListenerRuleRequest) GetVmIds() []string`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *CreateListenerRuleRequest) GetVmIdsOk() ([]string, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmIds

`func (o *CreateListenerRuleRequest) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.

### SetVmIds

`func (o *CreateListenerRuleRequest) SetVmIds(v []string)`

SetVmIds gets a reference to the given []string and assigns it to the VmIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


