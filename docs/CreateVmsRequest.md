# CreateVmsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockDeviceMappings** | Pointer to [**[]BlockDeviceMappingVmCreation**](BlockDeviceMappingVmCreation.md) | One or more block device mappings. | [optional] 
**BootOnCreation** | Pointer to **bool** | By default or if &#x60;true&#x60;, the VM is started on creation. If &#x60;false&#x60;, the VM is stopped on creation. | [optional] 
**BsuOptimized** | Pointer to **bool** | If &#x60;true&#x60;, the VM is created with optimized BSU I/O. | [optional] 
**ClientToken** | Pointer to **string** | A unique identifier which enables you to manage the idempotency. | [optional] 
**DeletionProtection** | Pointer to **bool** | If &#x60;true&#x60;, you cannot terminate the VM using Cockpit, the CLI or the API. If &#x60;false&#x60;, you can. | [optional] 
**DryRun** | Pointer to **bool** | If &#x60;true&#x60;, checks whether you have the required permissions to perform the action. | [optional] 
**ImageId** | Pointer to **string** | The ID of the OMI used to create the VM. You can find the list of OMIs by calling the [ReadImages](#readimages) method. | 
**KeypairName** | Pointer to **string** | The name of the keypair. | [optional] 
**MaxVmsCount** | Pointer to **int64** | The maximum number of VMs you want to create. If all the VMs cannot be created, the largest possible number of VMs above MinVmsCount is created. | [optional] 
**MinVmsCount** | Pointer to **int64** | The minimum number of VMs you want to create. If this number of VMs cannot be created, no VMs are created. | [optional] 
**Nics** | Pointer to [**[]NicForVmCreation**](NicForVmCreation.md) | One or more NICs. If you specify this parameter, you must define one NIC as the primary network interface of the VM with &#x60;0&#x60; as its device number. | [optional] 
**Performance** | Pointer to **string** | The performance of the VM (&#x60;standard&#x60; \\| &#x60;high&#x60; \\|  &#x60;highest&#x60;). | [optional] [default to PERFORMANCE_HIGH]
**Placement** | Pointer to [**Placement**](Placement.md) |  | [optional] 
**PrivateIps** | Pointer to **[]string** | One or more private IP addresses of the VM. | [optional] 
**SecurityGroupIds** | Pointer to **[]string** | One or more IDs of security group for the VMs. | [optional] 
**SecurityGroups** | Pointer to **[]string** | One or more names of security groups for the VMs. | [optional] 
**SubnetId** | Pointer to **string** | The ID of the Subnet in which you want to create the VM. | [optional] 
**UserData** | Pointer to **string** | Data or script used to add a specific configuration to the VM. It must be base64-encoded. | [optional] 
**VmInitiatedShutdownBehavior** | Pointer to **string** | The VM behavior when you stop it. By default or if set to &#x60;stop&#x60;, the VM stops. If set to &#x60;restart&#x60;, the VM stops then automatically restarts. If set to &#x60;terminate&#x60;, the VM stops and is terminated. | [optional] 
**VmType** | Pointer to **string** | The type of VM (&#x60;tinav2.c1r2&#x60; by default).&lt;br /&gt; For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types). | [optional] 

## Methods

### GetBlockDeviceMappings

`func (o *CreateVmsRequest) GetBlockDeviceMappings() []BlockDeviceMappingVmCreation`

GetBlockDeviceMappings returns the BlockDeviceMappings field if non-nil, zero value otherwise.

### GetBlockDeviceMappingsOk

`func (o *CreateVmsRequest) GetBlockDeviceMappingsOk() ([]BlockDeviceMappingVmCreation, bool)`

GetBlockDeviceMappingsOk returns a tuple with the BlockDeviceMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBlockDeviceMappings

`func (o *CreateVmsRequest) HasBlockDeviceMappings() bool`

HasBlockDeviceMappings returns a boolean if a field has been set.

### SetBlockDeviceMappings

`func (o *CreateVmsRequest) SetBlockDeviceMappings(v []BlockDeviceMappingVmCreation)`

SetBlockDeviceMappings gets a reference to the given []BlockDeviceMappingVmCreation and assigns it to the BlockDeviceMappings field.

### GetBootOnCreation

`func (o *CreateVmsRequest) GetBootOnCreation() bool`

GetBootOnCreation returns the BootOnCreation field if non-nil, zero value otherwise.

### GetBootOnCreationOk

`func (o *CreateVmsRequest) GetBootOnCreationOk() (bool, bool)`

GetBootOnCreationOk returns a tuple with the BootOnCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBootOnCreation

`func (o *CreateVmsRequest) HasBootOnCreation() bool`

HasBootOnCreation returns a boolean if a field has been set.

### SetBootOnCreation

`func (o *CreateVmsRequest) SetBootOnCreation(v bool)`

SetBootOnCreation gets a reference to the given bool and assigns it to the BootOnCreation field.

### GetBsuOptimized

`func (o *CreateVmsRequest) GetBsuOptimized() bool`

GetBsuOptimized returns the BsuOptimized field if non-nil, zero value otherwise.

### GetBsuOptimizedOk

`func (o *CreateVmsRequest) GetBsuOptimizedOk() (bool, bool)`

GetBsuOptimizedOk returns a tuple with the BsuOptimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBsuOptimized

`func (o *CreateVmsRequest) HasBsuOptimized() bool`

HasBsuOptimized returns a boolean if a field has been set.

### SetBsuOptimized

`func (o *CreateVmsRequest) SetBsuOptimized(v bool)`

SetBsuOptimized gets a reference to the given bool and assigns it to the BsuOptimized field.

### GetClientToken

`func (o *CreateVmsRequest) GetClientToken() string`

GetClientToken returns the ClientToken field if non-nil, zero value otherwise.

### GetClientTokenOk

`func (o *CreateVmsRequest) GetClientTokenOk() (string, bool)`

GetClientTokenOk returns a tuple with the ClientToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientToken

`func (o *CreateVmsRequest) HasClientToken() bool`

HasClientToken returns a boolean if a field has been set.

### SetClientToken

`func (o *CreateVmsRequest) SetClientToken(v string)`

SetClientToken gets a reference to the given string and assigns it to the ClientToken field.

### GetDeletionProtection

`func (o *CreateVmsRequest) GetDeletionProtection() bool`

GetDeletionProtection returns the DeletionProtection field if non-nil, zero value otherwise.

### GetDeletionProtectionOk

`func (o *CreateVmsRequest) GetDeletionProtectionOk() (bool, bool)`

GetDeletionProtectionOk returns a tuple with the DeletionProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeletionProtection

`func (o *CreateVmsRequest) HasDeletionProtection() bool`

HasDeletionProtection returns a boolean if a field has been set.

### SetDeletionProtection

`func (o *CreateVmsRequest) SetDeletionProtection(v bool)`

SetDeletionProtection gets a reference to the given bool and assigns it to the DeletionProtection field.

### GetDryRun

`func (o *CreateVmsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateVmsRequest) GetDryRunOk() (bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDryRun

`func (o *CreateVmsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### SetDryRun

`func (o *CreateVmsRequest) SetDryRun(v bool)`

SetDryRun gets a reference to the given bool and assigns it to the DryRun field.

### GetImageId

`func (o *CreateVmsRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateVmsRequest) GetImageIdOk() (string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImageId

`func (o *CreateVmsRequest) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageId

`func (o *CreateVmsRequest) SetImageId(v string)`

SetImageId gets a reference to the given string and assigns it to the ImageId field.

### GetKeypairName

`func (o *CreateVmsRequest) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *CreateVmsRequest) GetKeypairNameOk() (string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeypairName

`func (o *CreateVmsRequest) HasKeypairName() bool`

HasKeypairName returns a boolean if a field has been set.

### SetKeypairName

`func (o *CreateVmsRequest) SetKeypairName(v string)`

SetKeypairName gets a reference to the given string and assigns it to the KeypairName field.

### GetMaxVmsCount

`func (o *CreateVmsRequest) GetMaxVmsCount() int64`

GetMaxVmsCount returns the MaxVmsCount field if non-nil, zero value otherwise.

### GetMaxVmsCountOk

`func (o *CreateVmsRequest) GetMaxVmsCountOk() (int64, bool)`

GetMaxVmsCountOk returns a tuple with the MaxVmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMaxVmsCount

`func (o *CreateVmsRequest) HasMaxVmsCount() bool`

HasMaxVmsCount returns a boolean if a field has been set.

### SetMaxVmsCount

`func (o *CreateVmsRequest) SetMaxVmsCount(v int64)`

SetMaxVmsCount gets a reference to the given int64 and assigns it to the MaxVmsCount field.

### GetMinVmsCount

`func (o *CreateVmsRequest) GetMinVmsCount() int64`

GetMinVmsCount returns the MinVmsCount field if non-nil, zero value otherwise.

### GetMinVmsCountOk

`func (o *CreateVmsRequest) GetMinVmsCountOk() (int64, bool)`

GetMinVmsCountOk returns a tuple with the MinVmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMinVmsCount

`func (o *CreateVmsRequest) HasMinVmsCount() bool`

HasMinVmsCount returns a boolean if a field has been set.

### SetMinVmsCount

`func (o *CreateVmsRequest) SetMinVmsCount(v int64)`

SetMinVmsCount gets a reference to the given int64 and assigns it to the MinVmsCount field.

### GetNics

`func (o *CreateVmsRequest) GetNics() []NicForVmCreation`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *CreateVmsRequest) GetNicsOk() ([]NicForVmCreation, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNics

`func (o *CreateVmsRequest) HasNics() bool`

HasNics returns a boolean if a field has been set.

### SetNics

`func (o *CreateVmsRequest) SetNics(v []NicForVmCreation)`

SetNics gets a reference to the given []NicForVmCreation and assigns it to the Nics field.

### GetPerformance

`func (o *CreateVmsRequest) GetPerformance() string`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *CreateVmsRequest) GetPerformanceOk() (string, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPerformance

`func (o *CreateVmsRequest) HasPerformance() bool`

HasPerformance returns a boolean if a field has been set.

### SetPerformance

`func (o *CreateVmsRequest) SetPerformance(v string)`

SetPerformance gets a reference to the given string and assigns it to the Performance field.

### GetPlacement

`func (o *CreateVmsRequest) GetPlacement() Placement`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *CreateVmsRequest) GetPlacementOk() (Placement, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlacement

`func (o *CreateVmsRequest) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacement

`func (o *CreateVmsRequest) SetPlacement(v Placement)`

SetPlacement gets a reference to the given Placement and assigns it to the Placement field.

### GetPrivateIps

`func (o *CreateVmsRequest) GetPrivateIps() []string`

GetPrivateIps returns the PrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsOk

`func (o *CreateVmsRequest) GetPrivateIpsOk() ([]string, bool)`

GetPrivateIpsOk returns a tuple with the PrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateIps

`func (o *CreateVmsRequest) HasPrivateIps() bool`

HasPrivateIps returns a boolean if a field has been set.

### SetPrivateIps

`func (o *CreateVmsRequest) SetPrivateIps(v []string)`

SetPrivateIps gets a reference to the given []string and assigns it to the PrivateIps field.

### GetSecurityGroupIds

`func (o *CreateVmsRequest) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *CreateVmsRequest) GetSecurityGroupIdsOk() ([]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroupIds

`func (o *CreateVmsRequest) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### SetSecurityGroupIds

`func (o *CreateVmsRequest) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds gets a reference to the given []string and assigns it to the SecurityGroupIds field.

### GetSecurityGroups

`func (o *CreateVmsRequest) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CreateVmsRequest) GetSecurityGroupsOk() ([]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroups

`func (o *CreateVmsRequest) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroups

`func (o *CreateVmsRequest) SetSecurityGroups(v []string)`

SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.

### GetSubnetId

`func (o *CreateVmsRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateVmsRequest) GetSubnetIdOk() (string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubnetId

`func (o *CreateVmsRequest) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetId

`func (o *CreateVmsRequest) SetSubnetId(v string)`

SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.

### GetUserData

`func (o *CreateVmsRequest) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreateVmsRequest) GetUserDataOk() (string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserData

`func (o *CreateVmsRequest) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserData

`func (o *CreateVmsRequest) SetUserData(v string)`

SetUserData gets a reference to the given string and assigns it to the UserData field.

### GetVmInitiatedShutdownBehavior

`func (o *CreateVmsRequest) GetVmInitiatedShutdownBehavior() string`

GetVmInitiatedShutdownBehavior returns the VmInitiatedShutdownBehavior field if non-nil, zero value otherwise.

### GetVmInitiatedShutdownBehaviorOk

`func (o *CreateVmsRequest) GetVmInitiatedShutdownBehaviorOk() (string, bool)`

GetVmInitiatedShutdownBehaviorOk returns a tuple with the VmInitiatedShutdownBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmInitiatedShutdownBehavior

`func (o *CreateVmsRequest) HasVmInitiatedShutdownBehavior() bool`

HasVmInitiatedShutdownBehavior returns a boolean if a field has been set.

### SetVmInitiatedShutdownBehavior

`func (o *CreateVmsRequest) SetVmInitiatedShutdownBehavior(v string)`

SetVmInitiatedShutdownBehavior gets a reference to the given string and assigns it to the VmInitiatedShutdownBehavior field.

### GetVmType

`func (o *CreateVmsRequest) GetVmType() string`

GetVmType returns the VmType field if non-nil, zero value otherwise.

### GetVmTypeOk

`func (o *CreateVmsRequest) GetVmTypeOk() (string, bool)`

GetVmTypeOk returns a tuple with the VmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmType

`func (o *CreateVmsRequest) HasVmType() bool`

HasVmType returns a boolean if a field has been set.

### SetVmType

`func (o *CreateVmsRequest) SetVmType(v string)`

SetVmType gets a reference to the given string and assigns it to the VmType field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


