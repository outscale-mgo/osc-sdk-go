# DeleteListenerRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**ListenerRuleName** | Pointer to **string** | The name of the rule you want to delete. | 

## Methods

### GetDryRun

`func (o *DeleteListenerRuleRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *DeleteListenerRuleRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *DeleteListenerRuleRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *DeleteListenerRuleRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetListenerRuleName

`func (o *DeleteListenerRuleRequest) GetListenerRuleName() string`

GetListenerRuleName returns the ListenerRuleName field if non-nil, zero value otherwise.

### GetListenerRuleNameOk

`func (o *DeleteListenerRuleRequest) GetListenerRuleNameOk() (string, bool)`

GetListenerRuleNameOk returns a tuple with the ListenerRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasListenerRuleName

`func (o *DeleteListenerRuleRequest) HasListenerRuleName() bool`

HasListenerRuleName returns a boolean if a field has been set.

### SetListenerRuleName

`func (o *DeleteListenerRuleRequest) SetListenerRuleName(v string)`

SetListenerRuleName gets a reference to the given string and assigns it to the ListenerRuleName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


