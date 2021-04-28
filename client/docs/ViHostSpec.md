# ViHostSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** | Port used to communicate with the server. | [optional] 
**CertificateThumbprint** | Pointer to **string** | [Optional] Certificate thumbprint used to verify the server identity. For details on how to get the thumbprint, see [Get TLS Certificate or SSH Fingerprint](#operation/GetConnectionCertificate).  | [optional] 

## Methods

### NewViHostSpec

`func NewViHostSpec() *ViHostSpec`

NewViHostSpec instantiates a new ViHostSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViHostSpecWithDefaults

`func NewViHostSpecWithDefaults() *ViHostSpec`

NewViHostSpecWithDefaults instantiates a new ViHostSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *ViHostSpec) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ViHostSpec) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ViHostSpec) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ViHostSpec) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCertificateThumbprint

`func (o *ViHostSpec) GetCertificateThumbprint() string`

GetCertificateThumbprint returns the CertificateThumbprint field if non-nil, zero value otherwise.

### GetCertificateThumbprintOk

`func (o *ViHostSpec) GetCertificateThumbprintOk() (*string, bool)`

GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateThumbprint

`func (o *ViHostSpec) SetCertificateThumbprint(v string)`

SetCertificateThumbprint sets CertificateThumbprint field to given value.

### HasCertificateThumbprint

`func (o *ViHostSpec) HasCertificateThumbprint() bool`

HasCertificateThumbprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


