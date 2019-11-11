# PermissionsOnResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIds** | Pointer to **[]string** | The account ID of one or more users who have permissions for the resource. | [optional] 
**GlobalPermission** | Pointer to **bool** | If &#x60;true&#x60;, the resource is public. If &#x60;false&#x60;, the resource is private. | [optional] 

## Methods

### GetAccountIds

`func (o *PermissionsOnResource) GetAccountIds() []string`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *PermissionsOnResource) GetAccountIdsOk() ([]string, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountIds

`func (o *PermissionsOnResource) HasAccountIds() bool`

HasAccountIds returns a boolean if a field has been set.

### SetAccountIds

`func (o *PermissionsOnResource) SetAccountIds(v []string)`

SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.

### GetGlobalPermission

`func (o *PermissionsOnResource) GetGlobalPermission() bool`

GetGlobalPermission returns the GlobalPermission field if non-nil, zero value otherwise.

### GetGlobalPermissionOk

`func (o *PermissionsOnResource) GetGlobalPermissionOk() (bool, bool)`

GetGlobalPermissionOk returns a tuple with the GlobalPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasGlobalPermission

`func (o *PermissionsOnResource) HasGlobalPermission() bool`

HasGlobalPermission returns a boolean if a field has been set.

### SetGlobalPermission

`func (o *PermissionsOnResource) SetGlobalPermission(v bool)`

SetGlobalPermission gets a reference to the given bool and assigns it to the GlobalPermission field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


