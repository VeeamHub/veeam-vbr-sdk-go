# BackupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the backup. | 
**JobId** | Pointer to **string** | ID of the job that created the backup. | [optional] 
**PolicyTag** | Pointer to **string** | Tag that identifies retention policy. | [optional] 
**Name** | **string** | Name of the job that created the backup. | 
**PlatformName** | [**EPlatformType**](EPlatformType.md) |  | 
**PlatformId** | **string** | ID of the platform of the backup resource. | 
**CreationTime** | **time.Time** | Date and time when the backup was created. | 
**RepositoryId** | **string** | ID of the backup repository where the backup is stored. | 

## Methods

### NewBackupModel

`func NewBackupModel(id string, name string, platformName EPlatformType, platformId string, creationTime time.Time, repositoryId string, ) *BackupModel`

NewBackupModel instantiates a new BackupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupModelWithDefaults

`func NewBackupModelWithDefaults() *BackupModel`

NewBackupModelWithDefaults instantiates a new BackupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupModel) SetId(v string)`

SetId sets Id field to given value.


### GetJobId

`func (o *BackupModel) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BackupModel) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BackupModel) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *BackupModel) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetPolicyTag

`func (o *BackupModel) GetPolicyTag() string`

GetPolicyTag returns the PolicyTag field if non-nil, zero value otherwise.

### GetPolicyTagOk

`func (o *BackupModel) GetPolicyTagOk() (*string, bool)`

GetPolicyTagOk returns a tuple with the PolicyTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTag

`func (o *BackupModel) SetPolicyTag(v string)`

SetPolicyTag sets PolicyTag field to given value.

### HasPolicyTag

`func (o *BackupModel) HasPolicyTag() bool`

HasPolicyTag returns a boolean if a field has been set.

### GetName

`func (o *BackupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupModel) SetName(v string)`

SetName sets Name field to given value.


### GetPlatformName

`func (o *BackupModel) GetPlatformName() EPlatformType`

GetPlatformName returns the PlatformName field if non-nil, zero value otherwise.

### GetPlatformNameOk

`func (o *BackupModel) GetPlatformNameOk() (*EPlatformType, bool)`

GetPlatformNameOk returns a tuple with the PlatformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformName

`func (o *BackupModel) SetPlatformName(v EPlatformType)`

SetPlatformName sets PlatformName field to given value.


### GetPlatformId

`func (o *BackupModel) GetPlatformId() string`

GetPlatformId returns the PlatformId field if non-nil, zero value otherwise.

### GetPlatformIdOk

`func (o *BackupModel) GetPlatformIdOk() (*string, bool)`

GetPlatformIdOk returns a tuple with the PlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformId

`func (o *BackupModel) SetPlatformId(v string)`

SetPlatformId sets PlatformId field to given value.


### GetCreationTime

`func (o *BackupModel) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BackupModel) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BackupModel) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetRepositoryId

`func (o *BackupModel) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *BackupModel) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *BackupModel) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


