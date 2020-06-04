# ListenerRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The type of action for the rule (always &#x60;forward&#x60;). | [optional] 
**HostNamePattern** | Pointer to **string** | A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?]. | [optional] 
**ListenerId** | Pointer to **int64** | The ID of the listener. | [optional] 
**ListenerRuleId** | Pointer to **int64** | The ID of the listener rule. | [optional] 
**ListenerRuleName** | Pointer to **string** | A human-readable name for the listener rule. | [optional] 
**PathPattern** | Pointer to **string** | A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~\&quot;&#39;@:+?]. | [optional] 
**Priority** | Pointer to **int64** | The priority level of the listener rule, between &#x60;1&#x60; and &#x60;19999&#x60; both included. Each rule must have a unique priority level. Otherwise, an error is returned. | [optional] 
**VmIds** | Pointer to **[]string** | The IDs of the backend VMs. | [optional] 

## Methods

### GetAction

`func (o *ListenerRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ListenerRule) GetActionOk() (string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAction

`func (o *ListenerRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetAction

`func (o *ListenerRule) SetAction(v string)`

SetAction gets a reference to the given string and assigns it to the Action field.

### GetHostNamePattern

`func (o *ListenerRule) GetHostNamePattern() string`

GetHostNamePattern returns the HostNamePattern field if non-nil, zero value otherwise.

### GetHostNamePatternOk

`func (o *ListenerRule) GetHostNamePatternOk() (string, bool)`

GetHostNamePatternOk returns a tuple with the HostNamePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHostNamePattern

`func (o *ListenerRule) HasHostNamePattern() bool`

HasHostNamePattern returns a boolean if a field has been set.

### SetHostNamePattern

`func (o *ListenerRule) SetHostNamePattern(v string)`

SetHostNamePattern gets a reference to the given string and assigns it to the HostNamePattern field.

### GetListenerId

`func (o *ListenerRule) GetListenerId() int64`

GetListenerId returns the ListenerId field if non-nil, zero value otherwise.

### GetListenerIdOk

`func (o *ListenerRule) GetListenerIdOk() (int64, bool)`

GetListenerIdOk returns a tuple with the ListenerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasListenerId

`func (o *ListenerRule) HasListenerId() bool`

HasListenerId returns a boolean if a field has been set.

### SetListenerId

`func (o *ListenerRule) SetListenerId(v int64)`

SetListenerId gets a reference to the given int64 and assigns it to the ListenerId field.

### GetListenerRuleId

`func (o *ListenerRule) GetListenerRuleId() int64`

GetListenerRuleId returns the ListenerRuleId field if non-nil, zero value otherwise.

### GetListenerRuleIdOk

`func (o *ListenerRule) GetListenerRuleIdOk() (int64, bool)`

GetListenerRuleIdOk returns a tuple with the ListenerRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasListenerRuleId

`func (o *ListenerRule) HasListenerRuleId() bool`

HasListenerRuleId returns a boolean if a field has been set.

### SetListenerRuleId

`func (o *ListenerRule) SetListenerRuleId(v int64)`

SetListenerRuleId gets a reference to the given int64 and assigns it to the ListenerRuleId field.

### GetListenerRuleName

`func (o *ListenerRule) GetListenerRuleName() string`

GetListenerRuleName returns the ListenerRuleName field if non-nil, zero value otherwise.

### GetListenerRuleNameOk

`func (o *ListenerRule) GetListenerRuleNameOk() (string, bool)`

GetListenerRuleNameOk returns a tuple with the ListenerRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasListenerRuleName

`func (o *ListenerRule) HasListenerRuleName() bool`

HasListenerRuleName returns a boolean if a field has been set.

### SetListenerRuleName

`func (o *ListenerRule) SetListenerRuleName(v string)`

SetListenerRuleName gets a reference to the given string and assigns it to the ListenerRuleName field.

### GetPathPattern

`func (o *ListenerRule) GetPathPattern() string`

GetPathPattern returns the PathPattern field if non-nil, zero value otherwise.

### GetPathPatternOk

`func (o *ListenerRule) GetPathPatternOk() (string, bool)`

GetPathPatternOk returns a tuple with the PathPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPathPattern

`func (o *ListenerRule) HasPathPattern() bool`

HasPathPattern returns a boolean if a field has been set.

### SetPathPattern

`func (o *ListenerRule) SetPathPattern(v string)`

SetPathPattern gets a reference to the given string and assigns it to the PathPattern field.

### GetPriority

`func (o *ListenerRule) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ListenerRule) GetPriorityOk() (int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPriority

`func (o *ListenerRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriority

`func (o *ListenerRule) SetPriority(v int64)`

SetPriority gets a reference to the given int64 and assigns it to the Priority field.

### GetVmIds

`func (o *ListenerRule) GetVmIds() []string`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *ListenerRule) GetVmIdsOk() ([]string, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmIds

`func (o *ListenerRule) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.

### SetVmIds

`func (o *ListenerRule) SetVmIds(v []string)`

SetVmIds gets a reference to the given []string and assigns it to the VmIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


