# BackupRepositoryImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the backup repository. | 
**Tag** | Pointer to **string** | Tag assigned to the backup repository. | [optional] 

## Methods

### NewBackupRepositoryImportModel

`func NewBackupRepositoryImportModel(name string, ) *BackupRepositoryImportModel`

NewBackupRepositoryImportModel instantiates a new BackupRepositoryImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRepositoryImportModelWithDefaults

`func NewBackupRepositoryImportModelWithDefaults() *BackupRepositoryImportModel`

NewBackupRepositoryImportModelWithDefaults instantiates a new BackupRepositoryImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BackupRepositoryImportModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupRepositoryImportModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupRepositoryImportModel) SetName(v string)`

SetName sets Name field to given value.


### GetTag

`func (o *BackupRepositoryImportModel) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *BackupRepositoryImportModel) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *BackupRepositoryImportModel) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *BackupRepositoryImportModel) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


