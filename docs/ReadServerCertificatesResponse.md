# ReadServerCertificatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContext** | Pointer to [**ResponseContext**](ResponseContext.md) |  | [optional] 
**ServerCertificates** | Pointer to [**[]ServerCertificate**](ServerCertificate.md) | Information about one or more server certificates. | [optional] 

## Methods

### GetResponseContext

`func (o *ReadServerCertificatesResponse) GetResponseContext() ResponseContext`

GetResponseContext returns the ResponseContext field if non-nil, zero value otherwise.

### GetResponseContextOk

`func (o *ReadServerCertificatesResponse) GetResponseContextOk() (ResponseContext, bool)`

GetResponseContextOk returns a tuple with the ResponseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContext

`func (o *ReadServerCertificatesResponse) HasResponseContext() bool`

HasResponseContext returns a boolean if a field has been set.

### SetResponseContext

`func (o *ReadServerCertificatesResponse) SetResponseContext(v ResponseContext)`

SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.

### GetServerCertificates

`func (o *ReadServerCertificatesResponse) GetServerCertificates() []ServerCertificate`

GetServerCertificates returns the ServerCertificates field if non-nil, zero value otherwise.

### GetServerCertificatesOk

`func (o *ReadServerCertificatesResponse) GetServerCertificatesOk() ([]ServerCertificate, bool)`

GetServerCertificatesOk returns a tuple with the ServerCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServerCertificates

`func (o *ReadServerCertificatesResponse) HasServerCertificates() bool`

HasServerCertificates returns a boolean if a field has been set.

### SetServerCertificates

`func (o *ReadServerCertificatesResponse) SetServerCertificates(v []ServerCertificate)`

SetServerCertificates gets a reference to the given []ServerCertificate and assigns it to the ServerCertificates field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


