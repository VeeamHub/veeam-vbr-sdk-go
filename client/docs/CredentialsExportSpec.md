# CredentialsExportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowHiddenCreds** | Pointer to **bool** | If *true*, service credentials are exported. | [optional] 
**Ids** | Pointer to **[]string** | Array of credentials IDs. | [optional] 
**Types** | Pointer to [**[]ECredentialsType**](ECredentialsType.md) | Array of credentials types. | [optional] 
**Names** | Pointer to **[]string** | Array of credentials user names. Wildcard characters are supported. | [optional] 

## Methods

### NewCredentialsExportSpec

`func NewCredentialsExportSpec() *CredentialsExportSpec`

NewCredentialsExportSpec instantiates a new CredentialsExportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsExportSpecWithDefaults

`func NewCredentialsExportSpecWithDefaults() *CredentialsExportSpec`

NewCredentialsExportSpecWithDefaults instantiates a new CredentialsExportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowHiddenCreds

`func (o *CredentialsExportSpec) GetShowHiddenCreds() bool`

GetShowHiddenCreds returns the ShowHiddenCreds field if non-nil, zero value otherwise.

### GetShowHiddenCredsOk

`func (o *CredentialsExportSpec) GetShowHiddenCredsOk() (*bool, bool)`

GetShowHiddenCredsOk returns a tuple with the ShowHiddenCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHiddenCreds

`func (o *CredentialsExportSpec) SetShowHiddenCreds(v bool)`

SetShowHiddenCreds sets ShowHiddenCreds field to given value.

### HasShowHiddenCreds

`func (o *CredentialsExportSpec) HasShowHiddenCreds() bool`

HasShowHiddenCreds returns a boolean if a field has been set.

### GetIds

`func (o *CredentialsExportSpec) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *CredentialsExportSpec) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *CredentialsExportSpec) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *CredentialsExportSpec) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetTypes

`func (o *CredentialsExportSpec) GetTypes() []ECredentialsType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *CredentialsExportSpec) GetTypesOk() (*[]ECredentialsType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *CredentialsExportSpec) SetTypes(v []ECredentialsType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *CredentialsExportSpec) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetNames

`func (o *CredentialsExportSpec) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *CredentialsExportSpec) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *CredentialsExportSpec) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *CredentialsExportSpec) HasNames() bool`

HasNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


