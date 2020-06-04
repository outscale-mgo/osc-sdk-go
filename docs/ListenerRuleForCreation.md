# ListenerRuleForCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The type of action for the rule (always &#x60;forward&#x60;). | [optional] 
**HostNamePattern** | Pointer to **string** | A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].  | [optional] 
**ListenerRuleId** | Pointer to **string** | The ID of the listener. | [optional] 
**ListenerRuleName** | Pointer to **string** | A human-readable name for the listener rule. | [optional] 
**PathPattern** | Pointer to **string** | A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~\&quot;&#39;@:+?]. | [optional] 
**Priority** | Pointer to **int64** | The priority level of the listener rule, between &#x60;1&#x60; and &#x60;19999&#x60; both included. Each rule must have a unique priority level. Otherwise, an error is returned. | 

## Methods

### GetAction

`func (o *ListenerRuleForCreation) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ListenerRuleForCreation) GetActionOk() (string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAction

`func (o *ListenerRuleForCreation) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetAction

`func (o *ListenerRuleForCreation) SetAction(v string)`

SetAction gets a reference to the given string and assigns it to the Action field.

### GetHostNamePattern

`func (o *ListenerRuleForCreation) GetHostNamePattern() string`

GetHostNamePattern returns the HostNamePattern field if non-nil, zero value otherwise.

### GetHostNamePatternOk

`func (o *ListenerRuleForCreation) GetHostNamePatternOk() (string, bool)`

GetHostNamePatternOk returns a tuple with the HostNamePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHostNamePattern

`func (o *ListenerRuleForCreation) HasHostNamePattern() bool`

HasHostNamePattern returns a boolean if a field has been set.

### SetHostNamePattern

`func (o *ListenerRuleForCreation) SetHostNamePattern(v string)`

SetHostNamePattern gets a reference to the given string and assigns it to the HostNamePattern field.

### GetListenerRuleId

`func (o *ListenerRuleForCreation) GetListenerRuleId() string`

GetListenerRuleId returns the ListenerRuleId field if non-nil, zero value otherwise.

### GetListenerRuleIdOk

`func (o *ListenerRuleForCreation) GetListenerRuleIdOk() (string, bool)`

GetListenerRuleIdOk returns a tuple with the ListenerRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasListenerRuleId

`func (o *ListenerRuleForCreation) HasListenerRuleId() bool`

HasListenerRuleId returns a boolean if a field has been set.

### SetListenerRuleId

`func (o *ListenerRuleForCreation) SetListenerRuleId(v string)`

SetListenerRuleId gets a reference to the given string and assigns it to the ListenerRuleId field.

### GetListenerRuleName

`func (o *ListenerRuleForCreation) GetListenerRuleName() string`

GetListenerRuleName returns the ListenerRuleName field if non-nil, zero value otherwise.

### GetListenerRuleNameOk

`func (o *ListenerRuleForCreation) GetListenerRuleNameOk() (string, bool)`

GetListenerRuleNameOk returns a tuple with the ListenerRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasListenerRuleName

`func (o *ListenerRuleForCreation) HasListenerRuleName() bool`

HasListenerRuleName returns a boolean if a field has been set.

### SetListenerRuleName

`func (o *ListenerRuleForCreation) SetListenerRuleName(v string)`

SetListenerRuleName gets a reference to the given string and assigns it to the ListenerRuleName field.

### GetPathPattern

`func (o *ListenerRuleForCreation) GetPathPattern() string`

GetPathPattern returns the PathPattern field if non-nil, zero value otherwise.

### GetPathPatternOk

`func (o *ListenerRuleForCreation) GetPathPatternOk() (string, bool)`

GetPathPatternOk returns a tuple with the PathPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPathPattern

`func (o *ListenerRuleForCreation) HasPathPattern() bool`

HasPathPattern returns a boolean if a field has been set.

### SetPathPattern

`func (o *ListenerRuleForCreation) SetPathPattern(v string)`

SetPathPattern gets a reference to the given string and assigns it to the PathPattern field.

### GetPriority

`func (o *ListenerRuleForCreation) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ListenerRuleForCreation) GetPriorityOk() (int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPriority

`func (o *ListenerRuleForCreation) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriority

`func (o *ListenerRuleForCreation) SetPriority(v int64)`

SetPriority gets a reference to the given int64 and assigns it to the Priority field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


