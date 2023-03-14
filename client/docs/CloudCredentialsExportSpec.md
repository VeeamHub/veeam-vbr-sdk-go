# CloudCredentialsExportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | Array of cloud credentials IDs that you want to export. | [optional] 
**Types** | Pointer to [**[]ECloudCredentialsType**](ECloudCredentialsType.md) | Array of cloud credentials types that you want to export. | [optional] 
**Names** | Pointer to **[]string** | Array of cloud credentials user names. Wildcard characters are supported. | [optional] 

## Methods

### NewCloudCredentialsExportSpec

`func NewCloudCredentialsExportSpec() *CloudCredentialsExportSpec`

NewCloudCredentialsExportSpec instantiates a new CloudCredentialsExportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsExportSpecWithDefaults

`func NewCloudCredentialsExportSpecWithDefaults() *CloudCredentialsExportSpec`

NewCloudCredentialsExportSpecWithDefaults instantiates a new CloudCredentialsExportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *CloudCredentialsExportSpec) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *CloudCredentialsExportSpec) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *CloudCredentialsExportSpec) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *CloudCredentialsExportSpec) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetTypes

`func (o *CloudCredentialsExportSpec) GetTypes() []ECloudCredentialsType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *CloudCredentialsExportSpec) GetTypesOk() (*[]ECloudCredentialsType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *CloudCredentialsExportSpec) SetTypes(v []ECloudCredentialsType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *CloudCredentialsExportSpec) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetNames

`func (o *CloudCredentialsExportSpec) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *CloudCredentialsExportSpec) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *CloudCredentialsExportSpec) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *CloudCredentialsExportSpec) HasNames() bool`

HasNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


