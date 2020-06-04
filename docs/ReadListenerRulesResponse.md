# ReadListenerRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListenerRules** | Pointer to [**[]ListenerRule**](ListenerRule.md) | The list of the rules to describe. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetListenerRules

`func (o *ReadListenerRulesResponse) GetListenerRules() []ListenerRule`

GetListenerRules returns the ListenerRules field if non-nil, zero value otherwise.

### GetListenerRulesOk

`func (o *ReadListenerRulesResponse) GetListenerRulesOk() ([]ListenerRule, bool)`

GetListenerRulesOk returns a tuple with the ListenerRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasListenerRules

`func (o *ReadListenerRulesResponse) HasListenerRules() bool`

HasListenerRules returns a boolean if a field has been set.

### SetListenerRules

`func (o *ReadListenerRulesResponse) SetListenerRules(v []ListenerRule)`

SetListenerRules gets a reference to the given []ListenerRule and assigns it to the ListenerRules field.

### GetResponseContext

`func (o *ReadListenerRulesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadListenerRulesResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadListenerRulesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadListenerRulesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


