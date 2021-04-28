# JobExportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | Array of job IDs. | [optional] 
**Types** | Pointer to **[]string** | Array of job types. | [optional] 
**Names** | Pointer to **[]string** | Array of job names. Wildcard characters are supported. | [optional] 

## Methods

### NewJobExportSpec

`func NewJobExportSpec() *JobExportSpec`

NewJobExportSpec instantiates a new JobExportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobExportSpecWithDefaults

`func NewJobExportSpecWithDefaults() *JobExportSpec`

NewJobExportSpecWithDefaults instantiates a new JobExportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *JobExportSpec) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *JobExportSpec) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *JobExportSpec) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *JobExportSpec) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetTypes

`func (o *JobExportSpec) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *JobExportSpec) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *JobExportSpec) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *JobExportSpec) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetNames

`func (o *JobExportSpec) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *JobExportSpec) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *JobExportSpec) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *JobExportSpec) HasNames() bool`

HasNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


