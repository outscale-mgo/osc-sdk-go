# ReadDirectLinksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectLinks** | Pointer to [**[]DirectLink**](DirectLink.md) | Information about one or more DirectLinks. | [optional] 
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 

## Methods

### GetDirectLinks

`func (o *ReadDirectLinksResponse) GetDirectLinks() []DirectLink`

GetDirectLinks returns the DirectLinks field if non-nil, zero value otherwise.

### GetDirectLinksOk

`func (o *ReadDirectLinksResponse) GetDirectLinksOk() ([]DirectLink, bool)`

GetDirectLinksOk returns a tuple with the DirectLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDirectLinks

`func (o *ReadDirectLinksResponse) HasDirectLinks() bool`

HasDirectLinks returns a boolean if a field has been set.

### SetDirectLinks

`func (o *ReadDirectLinksResponse) SetDirectLinks(v []DirectLink)`

SetDirectLinks gets a reference to the given []DirectLink and assigns it to the DirectLinks field.

### GetResponseContext

`func (o *ReadDirectLinksResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadDirectLinksResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadDirectLinksResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadDirectLinksResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


