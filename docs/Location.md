# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The location code, to be set as the &#x60;Location&#x60; parameter of the *CreateDirectLink* method when creating a DirectLink. | [optional] 
**Name** | Pointer to **string** | The name and description of the location, corresponding to a datacenter. | [optional] 

## Methods

### GetCode

`func (o *Location) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Location) GetCodeOk() (string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCode

`func (o *Location) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCode

`func (o *Location) SetCode(v string)`

SetCode gets a reference to the given string and assigns it to the Code field.

### GetName

`func (o *Location) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Location) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Location) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Location) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


