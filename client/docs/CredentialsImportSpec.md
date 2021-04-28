# CredentialsImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | User name. | 
**Password** | Pointer to **string** | Password. | [optional] 
**Description** | Pointer to **string** | Description of the credentials record. | [optional] 
**Tag** | **string** | Tag used to identify the credentials record. | 
**Type** | [**ECredentialsType**](ECredentialsType.md) |  | 
**LinuxAdditionalSettings** | Pointer to [**CredentialsLinuxSettingsImportModel**](CredentialsLinuxSettingsImportModel.md) |  | [optional] 

## Methods

### NewCredentialsImportSpec

`func NewCredentialsImportSpec(username string, tag string, type_ ECredentialsType, ) *CredentialsImportSpec`

NewCredentialsImportSpec instantiates a new CredentialsImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsImportSpecWithDefaults

`func NewCredentialsImportSpecWithDefaults() *CredentialsImportSpec`

NewCredentialsImportSpecWithDefaults instantiates a new CredentialsImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CredentialsImportSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CredentialsImportSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CredentialsImportSpec) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CredentialsImportSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CredentialsImportSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CredentialsImportSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CredentialsImportSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDescription

`func (o *CredentialsImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialsImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialsImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CredentialsImportSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTag

`func (o *CredentialsImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CredentialsImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CredentialsImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetType

`func (o *CredentialsImportSpec) GetType() ECredentialsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialsImportSpec) GetTypeOk() (*ECredentialsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialsImportSpec) SetType(v ECredentialsType)`

SetType sets Type field to given value.


### GetLinuxAdditionalSettings

`func (o *CredentialsImportSpec) GetLinuxAdditionalSettings() CredentialsLinuxSettingsImportModel`

GetLinuxAdditionalSettings returns the LinuxAdditionalSettings field if non-nil, zero value otherwise.

### GetLinuxAdditionalSettingsOk

`func (o *CredentialsImportSpec) GetLinuxAdditionalSettingsOk() (*CredentialsLinuxSettingsImportModel, bool)`

GetLinuxAdditionalSettingsOk returns a tuple with the LinuxAdditionalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxAdditionalSettings

`func (o *CredentialsImportSpec) SetLinuxAdditionalSettings(v CredentialsLinuxSettingsImportModel)`

SetLinuxAdditionalSettings sets LinuxAdditionalSettings field to given value.

### HasLinuxAdditionalSettings

`func (o *CredentialsImportSpec) HasLinuxAdditionalSettings() bool`

HasLinuxAdditionalSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


