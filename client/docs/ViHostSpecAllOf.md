# ViHostSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** | Port used to communicate with the server. | [optional] 
**CertificateThumbprint** | Pointer to **string** | [Optional] Certificate thumbprint used to verify the server identity. For details on how to get the thumbprint, see [Get TLS Certificate or SSH Fingerprint](#operation/GetConnectionCertificate).  | [optional] 

## Methods

### NewViHostSpecAllOf

`func NewViHostSpecAllOf() *ViHostSpecAllOf`

NewViHostSpecAllOf instantiates a new ViHostSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViHostSpecAllOfWithDefaults

`func NewViHostSpecAllOfWithDefaults() *ViHostSpecAllOf`

NewViHostSpecAllOfWithDefaults instantiates a new ViHostSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *ViHostSpecAllOf) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ViHostSpecAllOf) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ViHostSpecAllOf) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ViHostSpecAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCertificateThumbprint

`func (o *ViHostSpecAllOf) GetCertificateThumbprint() string`

GetCertificateThumbprint returns the CertificateThumbprint field if non-nil, zero value otherwise.

### GetCertificateThumbprintOk

`func (o *ViHostSpecAllOf) GetCertificateThumbprintOk() (*string, bool)`

GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateThumbprint

`func (o *ViHostSpecAllOf) SetCertificateThumbprint(v string)`

SetCertificateThumbprint sets CertificateThumbprint field to given value.

### HasCertificateThumbprint

`func (o *ViHostSpecAllOf) HasCertificateThumbprint() bool`

HasCertificateThumbprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


