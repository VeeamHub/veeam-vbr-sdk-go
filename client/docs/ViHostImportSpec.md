# ViHostImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Full DNS name or IP address of the server. | 
**Description** | **string** | Description of the server. | 
**Type** | [**EManagedServerType**](EManagedServerType.md) |  | 
**ViHostType** | [**EViHostType**](EViHostType.md) |  | 
**Credentials** | [**CredentialsImportModel**](CredentialsImportModel.md) |  | 
**Port** | Pointer to **int32** | Port used to communicate with the server. | [optional] 
**CertificateThumbprint** | Pointer to **string** | Certificate thumbprint used to verify the server identity. | [optional] 

## Methods

### NewViHostImportSpec

`func NewViHostImportSpec(name string, description string, type_ EManagedServerType, viHostType EViHostType, credentials CredentialsImportModel, ) *ViHostImportSpec`

NewViHostImportSpec instantiates a new ViHostImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViHostImportSpecWithDefaults

`func NewViHostImportSpecWithDefaults() *ViHostImportSpec`

NewViHostImportSpecWithDefaults instantiates a new ViHostImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ViHostImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViHostImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViHostImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ViHostImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ViHostImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ViHostImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ViHostImportSpec) GetType() EManagedServerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ViHostImportSpec) GetTypeOk() (*EManagedServerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ViHostImportSpec) SetType(v EManagedServerType)`

SetType sets Type field to given value.


### GetViHostType

`func (o *ViHostImportSpec) GetViHostType() EViHostType`

GetViHostType returns the ViHostType field if non-nil, zero value otherwise.

### GetViHostTypeOk

`func (o *ViHostImportSpec) GetViHostTypeOk() (*EViHostType, bool)`

GetViHostTypeOk returns a tuple with the ViHostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViHostType

`func (o *ViHostImportSpec) SetViHostType(v EViHostType)`

SetViHostType sets ViHostType field to given value.


### GetCredentials

`func (o *ViHostImportSpec) GetCredentials() CredentialsImportModel`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ViHostImportSpec) GetCredentialsOk() (*CredentialsImportModel, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ViHostImportSpec) SetCredentials(v CredentialsImportModel)`

SetCredentials sets Credentials field to given value.


### GetPort

`func (o *ViHostImportSpec) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ViHostImportSpec) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ViHostImportSpec) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ViHostImportSpec) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCertificateThumbprint

`func (o *ViHostImportSpec) GetCertificateThumbprint() string`

GetCertificateThumbprint returns the CertificateThumbprint field if non-nil, zero value otherwise.

### GetCertificateThumbprintOk

`func (o *ViHostImportSpec) GetCertificateThumbprintOk() (*string, bool)`

GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateThumbprint

`func (o *ViHostImportSpec) SetCertificateThumbprint(v string)`

SetCertificateThumbprint sets CertificateThumbprint field to given value.

### HasCertificateThumbprint

`func (o *ViHostImportSpec) HasCertificateThumbprint() bool`

HasCertificateThumbprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


