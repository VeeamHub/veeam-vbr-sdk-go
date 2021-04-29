# RepositoryExportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | Array of backup repository IDs. | [optional] 
**Types** | Pointer to [**[]ERepositoryType**](ERepositoryType.md) | Array of backup repository types. | [optional] 
**Names** | Pointer to **[]string** | Array of repository names. Wildcard characters are supported. | [optional] 

## Methods

### NewRepositoryExportSpec

`func NewRepositoryExportSpec() *RepositoryExportSpec`

NewRepositoryExportSpec instantiates a new RepositoryExportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryExportSpecWithDefaults

`func NewRepositoryExportSpecWithDefaults() *RepositoryExportSpec`

NewRepositoryExportSpecWithDefaults instantiates a new RepositoryExportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *RepositoryExportSpec) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *RepositoryExportSpec) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *RepositoryExportSpec) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *RepositoryExportSpec) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetTypes

`func (o *RepositoryExportSpec) GetTypes() []ERepositoryType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *RepositoryExportSpec) GetTypesOk() (*[]ERepositoryType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *RepositoryExportSpec) SetTypes(v []ERepositoryType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *RepositoryExportSpec) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetNames

`func (o *RepositoryExportSpec) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *RepositoryExportSpec) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *RepositoryExportSpec) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *RepositoryExportSpec) HasNames() bool`

HasNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


