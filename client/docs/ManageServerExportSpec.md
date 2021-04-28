# ManageServerExportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | Array of server IDs. | [optional] 
**Types** | Pointer to [**[]EManagedServerType**](EManagedServerType.md) | Array of server types. | [optional] 
**Names** | Pointer to **[]string** | Array of server names. Wildcard characters are supported. | [optional] 

## Methods

### NewManageServerExportSpec

`func NewManageServerExportSpec() *ManageServerExportSpec`

NewManageServerExportSpec instantiates a new ManageServerExportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageServerExportSpecWithDefaults

`func NewManageServerExportSpecWithDefaults() *ManageServerExportSpec`

NewManageServerExportSpecWithDefaults instantiates a new ManageServerExportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ManageServerExportSpec) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ManageServerExportSpec) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ManageServerExportSpec) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ManageServerExportSpec) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetTypes

`func (o *ManageServerExportSpec) GetTypes() []EManagedServerType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *ManageServerExportSpec) GetTypesOk() (*[]EManagedServerType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *ManageServerExportSpec) SetTypes(v []EManagedServerType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *ManageServerExportSpec) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetNames

`func (o *ManageServerExportSpec) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *ManageServerExportSpec) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *ManageServerExportSpec) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *ManageServerExportSpec) HasNames() bool`

HasNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


