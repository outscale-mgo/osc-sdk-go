# CreateDirectLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bandwidth** | Pointer to **string** | The bandwidth of the DirectLink (&#x60;1GBps&#x60; \\| &#x60;10GBps&#x60;). | 
**DirectLinkName** | Pointer to **string** | The name of the DirectLink. | 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**Location** | Pointer to **string** | The code of the requested location for the DirectLink, returned by the [ReadLocations](#readlocations) method. | 

## Methods

### GetBandwidth

`func (o *CreateDirectLinkRequest) GetBandwidth() string`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *CreateDirectLinkRequest) GetBandwidthOk() (string, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBandwidth

`func (o *CreateDirectLinkRequest) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### SetBandwidth

`func (o *CreateDirectLinkRequest) SetBandwidth(v string)`

SetBandwidth gets a reference to the given string and assigns it to the Bandwidth field.

### GetDirectLinkName

`func (o *CreateDirectLinkRequest) GetDirectLinkName() string`

GetDirectLinkName returns the DirectLinkName field if non-nil, zero value otherwise.

### GetDirectLinkNameOk

`func (o *CreateDirectLinkRequest) GetDirectLinkNameOk() (string, bool)`

GetDirectLinkNameOk returns a tuple with the DirectLinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDirectLinkName

`func (o *CreateDirectLinkRequest) HasDirectLinkName() bool`

HasDirectLinkName returns a boolean if a field has been set.

### SetDirectLinkName

`func (o *CreateDirectLinkRequest) SetDirectLinkName(v string)`

SetDirectLinkName gets a reference to the given string and assigns it to the DirectLinkName field.

### GetDryRun

`func (o *CreateDirectLinkRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateDirectLinkRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateDirectLinkRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateDirectLinkRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetLocation

`func (o *CreateDirectLinkRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateDirectLinkRequest) GetLocationOk() (string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLocation

`func (o *CreateDirectLinkRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocation

`func (o *CreateDirectLinkRequest) SetLocation(v string)`

SetLocation gets a reference to the given string and assigns it to the Location field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


