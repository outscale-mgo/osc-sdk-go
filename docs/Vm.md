# Vm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** | The architecture of the VM (&#x60;i386&#x60; \\| &#x60;x86_64&#x60;). | [optional] 
**BlockDeviceMappings** | Pointer to [**[]BlockDeviceMappingCreated**](BlockDeviceMappingCreated.md) | The block device mapping of the VM. | [optional] 
**BsuOptimized** | Pointer to **bool** | If &#x60;true&#x60;, the VM is optimized for BSU I/O. | [optional] 
**ClientToken** | Pointer to **string** | The idempotency token provided when launching the VM. | [optional] 
**DeletionProtection** | Pointer to **bool** | If &#x60;true&#x60;, you cannot terminate the VM using Cockpit, the CLI or the API. If &#x60;false&#x60;, you can. | [optional] 
**Hypervisor** | Pointer to **string** | The hypervisor type of the VMs (&#x60;ovm&#x60; \\| &#x60;xen&#x60;). | [optional] 
**ImageId** | Pointer to **string** | The ID of the OMI used to create the VM. | [optional] 
**IsSourceDestChecked** | Pointer to **bool** | (Net only) If &#x60;true&#x60;, the source/destination check is enabled. If &#x60;false&#x60;, it is disabled. This value must be &#x60;false&#x60; for a NAT VM to perform network address translation (NAT) in a Net. | [optional] 
**KeypairName** | Pointer to **string** | The name of the keypair used when launching the VM. | [optional] 
**LaunchNumber** | Pointer to **int64** | The number for the VM when launching a group of several VMs (for example, 0, 1, 2, and so on). | [optional] 
**NetId** | Pointer to **string** | The ID of the Net in which the VM is running. | [optional] 
**Nics** | Pointer to [**[]NicLight**](NicLight.md) | The network interface cards (NICs) the VMs are attached to. | [optional] 
**OsFamily** | Pointer to **string** | Indicates the operating system (OS) of the VM. | [optional] 
**Performance** | Pointer to **string** | The performance of the VM (&#x60;standard&#x60; \\| &#x60;high&#x60; \\|  &#x60;highest&#x60;). | [optional] 
**Placement** | Pointer to [**Placement**](Placement.md) |  | [optional] 
**PrivateDnsName** | Pointer to **string** | The name of the private DNS. | [optional] 
**PrivateIp** | Pointer to **string** | The primary private IP address of the VM. | [optional] 
**ProductCodes** | Pointer to **[]string** | The product code associated with the OMI used to create the VM (&#x60;0001&#x60; Linux/Unix \\| &#x60;0002&#x60; Windows \\| &#x60;0004&#x60; Linux/Oracle \\| &#x60;0005&#x60; Windows 10). | [optional] 
**PublicDnsName** | Pointer to **string** | The name of the public DNS. | [optional] 
**PublicIp** | Pointer to **string** | The public IP address of the VM. | [optional] 
**ReservationId** | Pointer to **string** | The reservation ID of the VM. | [optional] 
**RootDeviceName** | Pointer to **string** | The name of the root device for the VM (for example, /dev/vda1). | [optional] 
**RootDeviceType** | Pointer to **string** | The type of root device used by the VM (always &#x60;bsu&#x60;). | [optional] 
**SecurityGroups** | Pointer to [**[]SecurityGroupLight**](SecurityGroupLight.md) | One or more security groups associated with the VM. | [optional] 
**State** | Pointer to **string** | The state of the VM (&#x60;pending&#x60; \\| &#x60;running&#x60; \\| &#x60;stopping&#x60; \\| &#x60;stopped&#x60; \\| &#x60;shutting-down&#x60; \\| &#x60;terminated&#x60; \\| &#x60;quarantine&#x60;). | [optional] 
**StateReason** | Pointer to **string** | The reason explaining the current state of the VM. | [optional] 
**SubnetId** | Pointer to **string** | The ID of the Subnet for the VM. | [optional] 
**Tags** | Pointer to [**[]ResourceTag**](ResourceTag.md) | One or more tags associated with the VM. | [optional] 
**UserData** | Pointer to **string** | The Base64-encoded MIME user data. | [optional] 
**VmId** | Pointer to **string** | The ID of the VM. | [optional] 
**VmInitiatedShutdownBehavior** | Pointer to **string** | The VM behavior when you stop it. By default or if set to &#x60;stop&#x60;, the VM stops. If set to &#x60;restart&#x60;, the VM stops then automatically restarts. If set to &#x60;terminate&#x60;, the VM stops and is deleted. | [optional] 
**VmType** | Pointer to **string** | The type of VM. For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types). | [optional] 

## Methods

### GetArchitecture

`func (o *Vm) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *Vm) GetArchitectureOk() (string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasArchitecture

`func (o *Vm) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### SetArchitecture

`func (o *Vm) SetArchitecture(v string)`

SetArchitecture gets a reference to the given string and assigns it to the Architecture field.

### GetBlockDeviceMappings

`func (o *Vm) GetBlockDeviceMappings() []BlockDeviceMappingCreated`

GetBlockDeviceMappings returns the BlockDeviceMappings field if non-nil, zero value otherwise.

### GetBlockDeviceMappingsOk

`func (o *Vm) GetBlockDeviceMappingsOk() ([]BlockDeviceMappingCreated, bool)`

GetBlockDeviceMappingsOk returns a tuple with the BlockDeviceMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBlockDeviceMappings

`func (o *Vm) HasBlockDeviceMappings() bool`

HasBlockDeviceMappings returns a boolean if a field has been set.

### SetBlockDeviceMappings

`func (o *Vm) SetBlockDeviceMappings(v []BlockDeviceMappingCreated)`

SetBlockDeviceMappings gets a reference to the given []BlockDeviceMappingCreated and assigns it to the BlockDeviceMappings field.

### GetBsuOptimized

`func (o *Vm) GetBsuOptimized() bool`

GetBsuOptimized returns the BsuOptimized field if non-nil, zero value otherwise.

### GetBsuOptimizedOk

`func (o *Vm) GetBsuOptimizedOk() (bool, bool)`

GetBsuOptimizedOk returns a tuple with the BsuOptimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBsuOptimized

`func (o *Vm) HasBsuOptimized() bool`

HasBsuOptimized returns a boolean if a field has been set.

### SetBsuOptimized

`func (o *Vm) SetBsuOptimized(v bool)`

SetBsuOptimized gets a reference to the given bool and assigns it to the BsuOptimized field.

### GetClientToken

`func (o *Vm) GetClientToken() string`

GetClientToken returns the ClientToken field if non-nil, zero value otherwise.

### GetClientTokenOk

`func (o *Vm) GetClientTokenOk() (string, bool)`

GetClientTokenOk returns a tuple with the ClientToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientToken

`func (o *Vm) HasClientToken() bool`

HasClientToken returns a boolean if a field has been set.

### SetClientToken

`func (o *Vm) SetClientToken(v string)`

SetClientToken gets a reference to the given string and assigns it to the ClientToken field.

### GetDeletionProtection

`func (o *Vm) GetDeletionProtection() bool`

GetDeletionProtection returns the DeletionProtection field if non-nil, zero value otherwise.

### GetDeletionProtectionOk

`func (o *Vm) GetDeletionProtectionOk() (bool, bool)`

GetDeletionProtectionOk returns a tuple with the DeletionProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeletionProtection

`func (o *Vm) HasDeletionProtection() bool`

HasDeletionProtection returns a boolean if a field has been set.

### SetDeletionProtection

`func (o *Vm) SetDeletionProtection(v bool)`

SetDeletionProtection gets a reference to the given bool and assigns it to the DeletionProtection field.

### GetHypervisor

`func (o *Vm) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *Vm) GetHypervisorOk() (string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHypervisor

`func (o *Vm) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### SetHypervisor

`func (o *Vm) SetHypervisor(v string)`

SetHypervisor gets a reference to the given string and assigns it to the Hypervisor field.

### GetImageId

`func (o *Vm) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *Vm) GetImageIdOk() (string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImageId

`func (o *Vm) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageId

`func (o *Vm) SetImageId(v string)`

SetImageId gets a reference to the given string and assigns it to the ImageId field.

### GetIsSourceDestChecked

`func (o *Vm) GetIsSourceDestChecked() bool`

GetIsSourceDestChecked returns the IsSourceDestChecked field if non-nil, zero value otherwise.

### GetIsSourceDestCheckedOk

`func (o *Vm) GetIsSourceDestCheckedOk() (bool, bool)`

GetIsSourceDestCheckedOk returns a tuple with the IsSourceDestChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsSourceDestChecked

`func (o *Vm) HasIsSourceDestChecked() bool`

HasIsSourceDestChecked returns a boolean if a field has been set.

### SetIsSourceDestChecked

`func (o *Vm) SetIsSourceDestChecked(v bool)`

SetIsSourceDestChecked gets a reference to the given bool and assigns it to the IsSourceDestChecked field.

### GetKeypairName

`func (o *Vm) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *Vm) GetKeypairNameOk() (string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKeypairName

`func (o *Vm) HasKeypairName() bool`

HasKeypairName returns a boolean if a field has been set.

### SetKeypairName

`func (o *Vm) SetKeypairName(v string)`

SetKeypairName gets a reference to the given string and assigns it to the KeypairName field.

### GetLaunchNumber

`func (o *Vm) GetLaunchNumber() int64`

GetLaunchNumber returns the LaunchNumber field if non-nil, zero value otherwise.

### GetLaunchNumberOk

`func (o *Vm) GetLaunchNumberOk() (int64, bool)`

GetLaunchNumberOk returns a tuple with the LaunchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLaunchNumber

`func (o *Vm) HasLaunchNumber() bool`

HasLaunchNumber returns a boolean if a field has been set.

### SetLaunchNumber

`func (o *Vm) SetLaunchNumber(v int64)`

SetLaunchNumber gets a reference to the given int64 and assigns it to the LaunchNumber field.

### GetNetId

`func (o *Vm) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *Vm) GetNetIdOk() (string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNetId

`func (o *Vm) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### SetNetId

`func (o *Vm) SetNetId(v string)`

SetNetId gets a reference to the given string and assigns it to the NetId field.

### GetNics

`func (o *Vm) GetNics() []NicLight`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *Vm) GetNicsOk() ([]NicLight, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNics

`func (o *Vm) HasNics() bool`

HasNics returns a boolean if a field has been set.

### SetNics

`func (o *Vm) SetNics(v []NicLight)`

SetNics gets a reference to the given []NicLight and assigns it to the Nics field.

### GetOsFamily

`func (o *Vm) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *Vm) GetOsFamilyOk() (string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsFamily

`func (o *Vm) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### SetOsFamily

`func (o *Vm) SetOsFamily(v string)`

SetOsFamily gets a reference to the given string and assigns it to the OsFamily field.

### GetPerformance

`func (o *Vm) GetPerformance() string`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *Vm) GetPerformanceOk() (string, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPerformance

`func (o *Vm) HasPerformance() bool`

HasPerformance returns a boolean if a field has been set.

### SetPerformance

`func (o *Vm) SetPerformance(v string)`

SetPerformance gets a reference to the given string and assigns it to the Performance field.

### GetPlacement

`func (o *Vm) GetPlacement() Placement`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *Vm) GetPlacementOk() (Placement, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPlacement

`func (o *Vm) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacement

`func (o *Vm) SetPlacement(v Placement)`

SetPlacement gets a reference to the given Placement and assigns it to the Placement field.

### GetPrivateDnsName

`func (o *Vm) GetPrivateDnsName() string`

GetPrivateDnsName returns the PrivateDnsName field if non-nil, zero value otherwise.

### GetPrivateDnsNameOk

`func (o *Vm) GetPrivateDnsNameOk() (string, bool)`

GetPrivateDnsNameOk returns a tuple with the PrivateDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateDnsName

`func (o *Vm) HasPrivateDnsName() bool`

HasPrivateDnsName returns a boolean if a field has been set.

### SetPrivateDnsName

`func (o *Vm) SetPrivateDnsName(v string)`

SetPrivateDnsName gets a reference to the given string and assigns it to the PrivateDnsName field.

### GetPrivateIp

`func (o *Vm) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *Vm) GetPrivateIpOk() (string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateIp

`func (o *Vm) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIp

`func (o *Vm) SetPrivateIp(v string)`

SetPrivateIp gets a reference to the given string and assigns it to the PrivateIp field.

### GetProductCodes

`func (o *Vm) GetProductCodes() []string`

GetProductCodes returns the ProductCodes field if non-nil, zero value otherwise.

### GetProductCodesOk

`func (o *Vm) GetProductCodesOk() ([]string, bool)`

GetProductCodesOk returns a tuple with the ProductCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProductCodes

`func (o *Vm) HasProductCodes() bool`

HasProductCodes returns a boolean if a field has been set.

### SetProductCodes

`func (o *Vm) SetProductCodes(v []string)`

SetProductCodes gets a reference to the given []string and assigns it to the ProductCodes field.

### GetPublicDnsName

`func (o *Vm) GetPublicDnsName() string`

GetPublicDnsName returns the PublicDnsName field if non-nil, zero value otherwise.

### GetPublicDnsNameOk

`func (o *Vm) GetPublicDnsNameOk() (string, bool)`

GetPublicDnsNameOk returns a tuple with the PublicDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicDnsName

`func (o *Vm) HasPublicDnsName() bool`

HasPublicDnsName returns a boolean if a field has been set.

### SetPublicDnsName

`func (o *Vm) SetPublicDnsName(v string)`

SetPublicDnsName gets a reference to the given string and assigns it to the PublicDnsName field.

### GetPublicIp

`func (o *Vm) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *Vm) GetPublicIpOk() (string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPublicIp

`func (o *Vm) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### SetPublicIp

`func (o *Vm) SetPublicIp(v string)`

SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.

### GetReservationId

`func (o *Vm) GetReservationId() string`

GetReservationId returns the ReservationId field if non-nil, zero value otherwise.

### GetReservationIdOk

`func (o *Vm) GetReservationIdOk() (string, bool)`

GetReservationIdOk returns a tuple with the ReservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReservationId

`func (o *Vm) HasReservationId() bool`

HasReservationId returns a boolean if a field has been set.

### SetReservationId

`func (o *Vm) SetReservationId(v string)`

SetReservationId gets a reference to the given string and assigns it to the ReservationId field.

### GetRootDeviceName

`func (o *Vm) GetRootDeviceName() string`

GetRootDeviceName returns the RootDeviceName field if non-nil, zero value otherwise.

### GetRootDeviceNameOk

`func (o *Vm) GetRootDeviceNameOk() (string, bool)`

GetRootDeviceNameOk returns a tuple with the RootDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRootDeviceName

`func (o *Vm) HasRootDeviceName() bool`

HasRootDeviceName returns a boolean if a field has been set.

### SetRootDeviceName

`func (o *Vm) SetRootDeviceName(v string)`

SetRootDeviceName gets a reference to the given string and assigns it to the RootDeviceName field.

### GetRootDeviceType

`func (o *Vm) GetRootDeviceType() string`

GetRootDeviceType returns the RootDeviceType field if non-nil, zero value otherwise.

### GetRootDeviceTypeOk

`func (o *Vm) GetRootDeviceTypeOk() (string, bool)`

GetRootDeviceTypeOk returns a tuple with the RootDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRootDeviceType

`func (o *Vm) HasRootDeviceType() bool`

HasRootDeviceType returns a boolean if a field has been set.

### SetRootDeviceType

`func (o *Vm) SetRootDeviceType(v string)`

SetRootDeviceType gets a reference to the given string and assigns it to the RootDeviceType field.

### GetSecurityGroups

`func (o *Vm) GetSecurityGroups() []SecurityGroupLight`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *Vm) GetSecurityGroupsOk() ([]SecurityGroupLight, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSecurityGroups

`func (o *Vm) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroups

`func (o *Vm) SetSecurityGroups(v []SecurityGroupLight)`

SetSecurityGroups gets a reference to the given []SecurityGroupLight and assigns it to the SecurityGroups field.

### GetState

`func (o *Vm) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Vm) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *Vm) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *Vm) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetStateReason

`func (o *Vm) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *Vm) GetStateReasonOk() (string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStateReason

`func (o *Vm) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### SetStateReason

`func (o *Vm) SetStateReason(v string)`

SetStateReason gets a reference to the given string and assigns it to the StateReason field.

### GetSubnetId

`func (o *Vm) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *Vm) GetSubnetIdOk() (string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubnetId

`func (o *Vm) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetId

`func (o *Vm) SetSubnetId(v string)`

SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.

### GetTags

`func (o *Vm) GetTags() []ResourceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Vm) GetTagsOk() ([]ResourceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTags

`func (o *Vm) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTags

`func (o *Vm) SetTags(v []ResourceTag)`

SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.

### GetUserData

`func (o *Vm) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *Vm) GetUserDataOk() (string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUserData

`func (o *Vm) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserData

`func (o *Vm) SetUserData(v string)`

SetUserData gets a reference to the given string and assigns it to the UserData field.

### GetVmId

`func (o *Vm) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *Vm) GetVmIdOk() (string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmId

`func (o *Vm) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### SetVmId

`func (o *Vm) SetVmId(v string)`

SetVmId gets a reference to the given string and assigns it to the VmId field.

### GetVmInitiatedShutdownBehavior

`func (o *Vm) GetVmInitiatedShutdownBehavior() string`

GetVmInitiatedShutdownBehavior returns the VmInitiatedShutdownBehavior field if non-nil, zero value otherwise.

### GetVmInitiatedShutdownBehaviorOk

`func (o *Vm) GetVmInitiatedShutdownBehaviorOk() (string, bool)`

GetVmInitiatedShutdownBehaviorOk returns a tuple with the VmInitiatedShutdownBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmInitiatedShutdownBehavior

`func (o *Vm) HasVmInitiatedShutdownBehavior() bool`

HasVmInitiatedShutdownBehavior returns a boolean if a field has been set.

### SetVmInitiatedShutdownBehavior

`func (o *Vm) SetVmInitiatedShutdownBehavior(v string)`

SetVmInitiatedShutdownBehavior gets a reference to the given string and assigns it to the VmInitiatedShutdownBehavior field.

### GetVmType

`func (o *Vm) GetVmType() string`

GetVmType returns the VmType field if non-nil, zero value otherwise.

### GetVmTypeOk

`func (o *Vm) GetVmTypeOk() (string, bool)`

GetVmTypeOk returns a tuple with the VmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVmType

`func (o *Vm) HasVmType() bool`

HasVmType returns a boolean if a field has been set.

### SetVmType

`func (o *Vm) SetVmType(v string)`

SetVmType gets a reference to the given string and assigns it to the VmType field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


