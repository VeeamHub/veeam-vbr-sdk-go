# ObjectRestorePointModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**PlatformName** | Pointer to [**EPlatformType**](EPlatformType.md) |  | [optional] 
**PlatformId** | **string** | ID of a platform on which the object was created. | 
**CreationTime** | **time.Time** | Date and time when the restore point was created. | 
**BackupId** | **string** | ID of a backup that contains the restore point. | 
**AllowedOperations** | [**[]EObjectRestorePointOperation**](EObjectRestorePointOperation.md) | Array of operations allowed for the restore point. | 

## Methods

### NewObjectRestorePointModel

`func NewObjectRestorePointModel(id string, name string, platformId string, creationTime time.Time, backupId string, allowedOperations []EObjectRestorePointOperation, ) *ObjectRestorePointModel`

NewObjectRestorePointModel instantiates a new ObjectRestorePointModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectRestorePointModelWithDefaults

`func NewObjectRestorePointModelWithDefaults() *ObjectRestorePointModel`

NewObjectRestorePointModelWithDefaults instantiates a new ObjectRestorePointModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectRestorePointModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectRestorePointModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectRestorePointModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ObjectRestorePointModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectRestorePointModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectRestorePointModel) SetName(v string)`

SetName sets Name field to given value.


### GetPlatformName

`func (o *ObjectRestorePointModel) GetPlatformName() EPlatformType`

GetPlatformName returns the PlatformName field if non-nil, zero value otherwise.

### GetPlatformNameOk

`func (o *ObjectRestorePointModel) GetPlatformNameOk() (*EPlatformType, bool)`

GetPlatformNameOk returns a tuple with the PlatformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformName

`func (o *ObjectRestorePointModel) SetPlatformName(v EPlatformType)`

SetPlatformName sets PlatformName field to given value.

### HasPlatformName

`func (o *ObjectRestorePointModel) HasPlatformName() bool`

HasPlatformName returns a boolean if a field has been set.

### GetPlatformId

`func (o *ObjectRestorePointModel) GetPlatformId() string`

GetPlatformId returns the PlatformId field if non-nil, zero value otherwise.

### GetPlatformIdOk

`func (o *ObjectRestorePointModel) GetPlatformIdOk() (*string, bool)`

GetPlatformIdOk returns a tuple with the PlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformId

`func (o *ObjectRestorePointModel) SetPlatformId(v string)`

SetPlatformId sets PlatformId field to given value.


### GetCreationTime

`func (o *ObjectRestorePointModel) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ObjectRestorePointModel) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ObjectRestorePointModel) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetBackupId

`func (o *ObjectRestorePointModel) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *ObjectRestorePointModel) GetBackupIdOk() (*string, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *ObjectRestorePointModel) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.


### GetAllowedOperations

`func (o *ObjectRestorePointModel) GetAllowedOperations() []EObjectRestorePointOperation`

GetAllowedOperations returns the AllowedOperations field if non-nil, zero value otherwise.

### GetAllowedOperationsOk

`func (o *ObjectRestorePointModel) GetAllowedOperationsOk() (*[]EObjectRestorePointOperation, bool)`

GetAllowedOperationsOk returns a tuple with the AllowedOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOperations

`func (o *ObjectRestorePointModel) SetAllowedOperations(v []EObjectRestorePointOperation)`

SetAllowedOperations sets AllowedOperations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


