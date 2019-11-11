# ReadSnapshotsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**Snapshots** | Pointer to [**[]Snapshot**](Snapshot.md) | Information about one or more snapshots and their permissions. | [optional] 

## Methods

### GetResponseContext

`func (o *ReadSnapshotsResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadSnapshotsResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadSnapshotsResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadSnapshotsResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.

### GetSnapshots

`func (o *ReadSnapshotsResponse) GetSnapshots() []Snapshot`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *ReadSnapshotsResponse) GetSnapshotsOk() ([]Snapshot, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSnapshots

`func (o *ReadSnapshotsResponse) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### SetSnapshots

`func (o *ReadSnapshotsResponse) SetSnapshots(v []Snapshot)`

SetSnapshots gets a reference to the given []Snapshot and assigns it to the Snapshots field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


