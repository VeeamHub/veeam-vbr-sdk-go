# CertificateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Thumbprint** | **string** | Thumbprint of the certificate. | 
**SerialNumber** | **string** | Serial number of the certificate. | 
**KeyAlgorithm** | **string** | Key algorithm of the certificate. | 
**KeySize** | **string** | Key size of the certificate. | 
**Subject** | **string** | Subject of the certificate. | 
**IssuedTo** | **string** | Acquirer of the certificate. | 
**IssuedBy** | **string** | Issuer of the certificate. | 
**ValidFrom** | **time.Time** | Date and time the certificate is valid from. | 
**ValidBy** | **time.Time** | Expiration date and time of the certificate. | 

## Methods

### NewCertificateModel

`func NewCertificateModel(thumbprint string, serialNumber string, keyAlgorithm string, keySize string, subject string, issuedTo string, issuedBy string, validFrom time.Time, validBy time.Time, ) *CertificateModel`

NewCertificateModel instantiates a new CertificateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateModelWithDefaults

`func NewCertificateModelWithDefaults() *CertificateModel`

NewCertificateModelWithDefaults instantiates a new CertificateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThumbprint

`func (o *CertificateModel) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *CertificateModel) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *CertificateModel) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.


### GetSerialNumber

`func (o *CertificateModel) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CertificateModel) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CertificateModel) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetKeyAlgorithm

`func (o *CertificateModel) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *CertificateModel) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *CertificateModel) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.


### GetKeySize

`func (o *CertificateModel) GetKeySize() string`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *CertificateModel) GetKeySizeOk() (*string, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *CertificateModel) SetKeySize(v string)`

SetKeySize sets KeySize field to given value.


### GetSubject

`func (o *CertificateModel) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CertificateModel) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CertificateModel) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetIssuedTo

`func (o *CertificateModel) GetIssuedTo() string`

GetIssuedTo returns the IssuedTo field if non-nil, zero value otherwise.

### GetIssuedToOk

`func (o *CertificateModel) GetIssuedToOk() (*string, bool)`

GetIssuedToOk returns a tuple with the IssuedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTo

`func (o *CertificateModel) SetIssuedTo(v string)`

SetIssuedTo sets IssuedTo field to given value.


### GetIssuedBy

`func (o *CertificateModel) GetIssuedBy() string`

GetIssuedBy returns the IssuedBy field if non-nil, zero value otherwise.

### GetIssuedByOk

`func (o *CertificateModel) GetIssuedByOk() (*string, bool)`

GetIssuedByOk returns a tuple with the IssuedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedBy

`func (o *CertificateModel) SetIssuedBy(v string)`

SetIssuedBy sets IssuedBy field to given value.


### GetValidFrom

`func (o *CertificateModel) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *CertificateModel) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *CertificateModel) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.


### GetValidBy

`func (o *CertificateModel) GetValidBy() time.Time`

GetValidBy returns the ValidBy field if non-nil, zero value otherwise.

### GetValidByOk

`func (o *CertificateModel) GetValidByOk() (*time.Time, bool)`

GetValidByOk returns a tuple with the ValidBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidBy

`func (o *CertificateModel) SetValidBy(v time.Time)`

SetValidBy sets ValidBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


