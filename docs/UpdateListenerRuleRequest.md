# UpdateListenerRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**HostPattern** | Pointer to **string** | A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?]. | [optional] 
**ListenerRuleName** | Pointer to **string** | The name of the listener rule. | 
**PathPattern** | Pointer to **string** | A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~\&quot;&#39;@:+?]. | [optional] 

## Methods

### GetDryRun

`func (o *UpdateListenerRuleRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *UpdateListenerRuleRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *UpdateListenerRuleRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *UpdateListenerRuleRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetHostPattern

`func (o *UpdateListenerRuleRequest) GetHostPattern() string`

GetHostPattern returns the HostPattern field if non-nil, zero value otherwise.

### GetHostPatternOk

`func (o *UpdateListenerRuleRequest) GetHostPatternOk() (string, bool)`

GetHostPatternOk returns a tuple with the HostPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHostPattern

`func (o *UpdateListenerRuleRequest) HasHostPattern() bool`

HasHostPattern returns a boolean if a field has been set.

### SetHostPattern

`func (o *UpdateListenerRuleRequest) SetHostPattern(v string)`

SetHostPattern gets a reference to the given string and assigns it to the HostPattern field.

### GetListenerRuleName

`func (o *UpdateListenerRuleRequest) GetListenerRuleName() string`

GetListenerRuleName returns the ListenerRuleName field if non-nil, zero value otherwise.

### GetListenerRuleNameOk

`func (o *UpdateListenerRuleRequest) GetListenerRuleNameOk() (string, bool)`

GetListenerRuleNameOk returns a tuple with the ListenerRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasListenerRuleName

`func (o *UpdateListenerRuleRequest) HasListenerRuleName() bool`

HasListenerRuleName returns a boolean if a field has been set.

### SetListenerRuleName

`func (o *UpdateListenerRuleRequest) SetListenerRuleName(v string)`

SetListenerRuleName gets a reference to the given string and assigns it to the ListenerRuleName field.

### GetPathPattern

`func (o *UpdateListenerRuleRequest) GetPathPattern() string`

GetPathPattern returns the PathPattern field if non-nil, zero value otherwise.

### GetPathPatternOk

`func (o *UpdateListenerRuleRequest) GetPathPatternOk() (string, bool)`

GetPathPatternOk returns a tuple with the PathPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPathPattern

`func (o *UpdateListenerRuleRequest) HasPathPattern() bool`

HasPathPattern returns a boolean if a field has been set.

### SetPathPattern

`func (o *UpdateListenerRuleRequest) SetPathPattern(v string)`

SetPathPattern gets a reference to the given string and assigns it to the PathPattern field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


