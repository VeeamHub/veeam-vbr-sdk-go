# ManagedServerSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Full DNS name or IP address of the server. | 
**Description** | **string** | Description of the server. | 
**Type** | [**EManagedServerType**](EManagedServerType.md) |  | 
**CredentialsId** | **string** | ID of the credentials used to connect to the server. | 
**NetworkSettings** | Pointer to [**WindowsHostPortsModel**](WindowsHostPortsModel.md) |  | [optional] 
**Port** | Pointer to **int32** | Port used to communicate with the server. | [optional] 
**CertificateThumbprint** | Pointer to **string** | Certificate thumbprint used to verify the server identity. For details on how to get the thumbprint, see [Request TLS Certificate or SSH Fingerprint](#tag/Connection/operation/GetConnectionCertificate). | [optional] 

## Methods

### NewManagedServerSpec

`func NewManagedServerSpec(name string, description string, type_ EManagedServerType, credentialsId string, ) *ManagedServerSpec`

NewManagedServerSpec instantiates a new ManagedServerSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedServerSpecWithDefaults

`func NewManagedServerSpecWithDefaults() *ManagedServerSpec`

NewManagedServerSpecWithDefaults instantiates a new ManagedServerSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ManagedServerSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedServerSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedServerSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ManagedServerSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManagedServerSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManagedServerSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ManagedServerSpec) GetType() EManagedServerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagedServerSpec) GetTypeOk() (*EManagedServerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagedServerSpec) SetType(v EManagedServerType)`

SetType sets Type field to given value.


### GetCredentialsId

`func (o *ManagedServerSpec) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *ManagedServerSpec) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *ManagedServerSpec) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetNetworkSettings

`func (o *ManagedServerSpec) GetNetworkSettings() WindowsHostPortsModel`

GetNetworkSettings returns the NetworkSettings field if non-nil, zero value otherwise.

### GetNetworkSettingsOk

`func (o *ManagedServerSpec) GetNetworkSettingsOk() (*WindowsHostPortsModel, bool)`

GetNetworkSettingsOk returns a tuple with the NetworkSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSettings

`func (o *ManagedServerSpec) SetNetworkSettings(v WindowsHostPortsModel)`

SetNetworkSettings sets NetworkSettings field to given value.

### HasNetworkSettings

`func (o *ManagedServerSpec) HasNetworkSettings() bool`

HasNetworkSettings returns a boolean if a field has been set.

### GetPort

`func (o *ManagedServerSpec) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ManagedServerSpec) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ManagedServerSpec) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ManagedServerSpec) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCertificateThumbprint

`func (o *ManagedServerSpec) GetCertificateThumbprint() string`

GetCertificateThumbprint returns the CertificateThumbprint field if non-nil, zero value otherwise.

### GetCertificateThumbprintOk

`func (o *ManagedServerSpec) GetCertificateThumbprintOk() (*string, bool)`

GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateThumbprint

`func (o *ManagedServerSpec) SetCertificateThumbprint(v string)`

SetCertificateThumbprint sets CertificateThumbprint field to given value.

### HasCertificateThumbprint

`func (o *ManagedServerSpec) HasCertificateThumbprint() bool`

HasCertificateThumbprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


