# CertificateUploadSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | Base64-encoded string of the content of a PFX certificate file. | 
**FormatType** | [**ECertificateFileFormatType**](ECertificateFileFormatType.md) |  | 
**Password** | Pointer to **string** | Decryption password for the certificate file. | [optional] 

## Methods

### NewCertificateUploadSpec

`func NewCertificateUploadSpec(certificate string, formatType ECertificateFileFormatType, ) *CertificateUploadSpec`

NewCertificateUploadSpec instantiates a new CertificateUploadSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateUploadSpecWithDefaults

`func NewCertificateUploadSpecWithDefaults() *CertificateUploadSpec`

NewCertificateUploadSpecWithDefaults instantiates a new CertificateUploadSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *CertificateUploadSpec) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateUploadSpec) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateUploadSpec) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetFormatType

`func (o *CertificateUploadSpec) GetFormatType() ECertificateFileFormatType`

GetFormatType returns the FormatType field if non-nil, zero value otherwise.

### GetFormatTypeOk

`func (o *CertificateUploadSpec) GetFormatTypeOk() (*ECertificateFileFormatType, bool)`

GetFormatTypeOk returns a tuple with the FormatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatType

`func (o *CertificateUploadSpec) SetFormatType(v ECertificateFileFormatType)`

SetFormatType sets FormatType field to given value.


### GetPassword

`func (o *CertificateUploadSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CertificateUploadSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CertificateUploadSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CertificateUploadSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


