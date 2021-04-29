# CredentialsImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialsName** | **string** | User name, account name or access key. | 
**CredentialsTag** | Pointer to **string** | Tag used to identify the credentials record. | [optional] 

## Methods

### NewCredentialsImportModel

`func NewCredentialsImportModel(credentialsName string, ) *CredentialsImportModel`

NewCredentialsImportModel instantiates a new CredentialsImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsImportModelWithDefaults

`func NewCredentialsImportModelWithDefaults() *CredentialsImportModel`

NewCredentialsImportModelWithDefaults instantiates a new CredentialsImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialsName

`func (o *CredentialsImportModel) GetCredentialsName() string`

GetCredentialsName returns the CredentialsName field if non-nil, zero value otherwise.

### GetCredentialsNameOk

`func (o *CredentialsImportModel) GetCredentialsNameOk() (*string, bool)`

GetCredentialsNameOk returns a tuple with the CredentialsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsName

`func (o *CredentialsImportModel) SetCredentialsName(v string)`

SetCredentialsName sets CredentialsName field to given value.


### GetCredentialsTag

`func (o *CredentialsImportModel) GetCredentialsTag() string`

GetCredentialsTag returns the CredentialsTag field if non-nil, zero value otherwise.

### GetCredentialsTagOk

`func (o *CredentialsImportModel) GetCredentialsTagOk() (*string, bool)`

GetCredentialsTagOk returns a tuple with the CredentialsTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsTag

`func (o *CredentialsImportModel) SetCredentialsTag(v string)`

SetCredentialsTag sets CredentialsTag field to given value.

### HasCredentialsTag

`func (o *CredentialsImportModel) HasCredentialsTag() bool`

HasCredentialsTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


