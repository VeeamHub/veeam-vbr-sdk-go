# ConnectionCertificateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | Pointer to **string** | SSH key fingerprint used to verify the server identity. | [optional] 
**Certificate** | Pointer to [**CertificateModel**](CertificateModel.md) |  | [optional] 

## Methods

### NewConnectionCertificateModel

`func NewConnectionCertificateModel() *ConnectionCertificateModel`

NewConnectionCertificateModel instantiates a new ConnectionCertificateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionCertificateModelWithDefaults

`func NewConnectionCertificateModelWithDefaults() *ConnectionCertificateModel`

NewConnectionCertificateModelWithDefaults instantiates a new ConnectionCertificateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprint

`func (o *ConnectionCertificateModel) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ConnectionCertificateModel) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ConnectionCertificateModel) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ConnectionCertificateModel) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetCertificate

`func (o *ConnectionCertificateModel) GetCertificate() CertificateModel`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ConnectionCertificateModel) GetCertificateOk() (*CertificateModel, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ConnectionCertificateModel) SetCertificate(v CertificateModel)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ConnectionCertificateModel) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


