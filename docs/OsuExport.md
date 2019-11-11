# OsuExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskImageFormat** | Pointer to **string** | The format of the export disk (&#x60;qcow2&#x60; \\| &#x60;vdi&#x60; \\| &#x60;vmdk&#x60;). | 
**OsuApiKey** | Pointer to [**OsuApiKey**](OsuApiKey.md) |  | [optional] 
**OsuBucket** | Pointer to **string** | The name of the OSU bucket you want to export the object to. | 
**OsuManifestUrl** | Pointer to **string** | The URL of the manifest file. | [optional] 
**OsuPrefix** | Pointer to **string** | The prefix for the key of the OSU object. This key follows this format: &#x60;prefix + object_export_task_id + &#39;.&#39; + disk_image_format&#x60;. | [optional] 

## Methods

### GetDiskImageFormat

`func (o *OsuExport) GetDiskImageFormat() string`

GetDiskImageFormat returns the DiskImageFormat field if non-nil, zero value otherwise.

### GetDiskImageFormatOk

`func (o *OsuExport) GetDiskImageFormatOk() (string, bool)`

GetDiskImageFormatOk returns a tuple with the DiskImageFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiskImageFormat

`func (o *OsuExport) HasDiskImageFormat() bool`

HasDiskImageFormat returns a boolean if a field has been set.

### SetDiskImageFormat

`func (o *OsuExport) SetDiskImageFormat(v string)`

SetDiskImageFormat gets a reference to the given string and assigns it to the DiskImageFormat field.

### GetOsuApiKey

`func (o *OsuExport) GetOsuApiKey() OsuApiKey`

GetOsuApiKey returns the OsuApiKey field if non-nil, zero value otherwise.

### GetOsuApiKeyOk

`func (o *OsuExport) GetOsuApiKeyOk() (OsuApiKey, bool)`

GetOsuApiKeyOk returns a tuple with the OsuApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsuApiKey

`func (o *OsuExport) HasOsuApiKey() bool`

HasOsuApiKey returns a boolean if a field has been set.

### SetOsuApiKey

`func (o *OsuExport) SetOsuApiKey(v OsuApiKey)`

SetOsuApiKey gets a reference to the given OsuApiKey and assigns it to the OsuApiKey field.

### GetOsuBucket

`func (o *OsuExport) GetOsuBucket() string`

GetOsuBucket returns the OsuBucket field if non-nil, zero value otherwise.

### GetOsuBucketOk

`func (o *OsuExport) GetOsuBucketOk() (string, bool)`

GetOsuBucketOk returns a tuple with the OsuBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsuBucket

`func (o *OsuExport) HasOsuBucket() bool`

HasOsuBucket returns a boolean if a field has been set.

### SetOsuBucket

`func (o *OsuExport) SetOsuBucket(v string)`

SetOsuBucket gets a reference to the given string and assigns it to the OsuBucket field.

### GetOsuManifestUrl

`func (o *OsuExport) GetOsuManifestUrl() string`

GetOsuManifestUrl returns the OsuManifestUrl field if non-nil, zero value otherwise.

### GetOsuManifestUrlOk

`func (o *OsuExport) GetOsuManifestUrlOk() (string, bool)`

GetOsuManifestUrlOk returns a tuple with the OsuManifestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsuManifestUrl

`func (o *OsuExport) HasOsuManifestUrl() bool`

HasOsuManifestUrl returns a boolean if a field has been set.

### SetOsuManifestUrl

`func (o *OsuExport) SetOsuManifestUrl(v string)`

SetOsuManifestUrl gets a reference to the given string and assigns it to the OsuManifestUrl field.

### GetOsuPrefix

`func (o *OsuExport) GetOsuPrefix() string`

GetOsuPrefix returns the OsuPrefix field if non-nil, zero value otherwise.

### GetOsuPrefixOk

`func (o *OsuExport) GetOsuPrefixOk() (string, bool)`

GetOsuPrefixOk returns a tuple with the OsuPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOsuPrefix

`func (o *OsuExport) HasOsuPrefix() bool`

HasOsuPrefix returns a boolean if a field has been set.

### SetOsuPrefix

`func (o *OsuExport) SetOsuPrefix(v string)`

SetOsuPrefix gets a reference to the given string and assigns it to the OsuPrefix field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


