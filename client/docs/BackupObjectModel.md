# BackupObjectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the object. | 
**Name** | Pointer to **string** | Name of the object. | [optional] 
**Type** | Pointer to **string** | Type of the object. | [optional] 
**PlatformName** | [**EPlatformType**](EPlatformType.md) |  | 
**PlatformId** | Pointer to **string** | Id of the platform where the object was created. | [optional] 
**RestorePointsCount** | Pointer to **int32** | Number of restore points. | [optional] 

## Methods

### NewBackupObjectModel

`func NewBackupObjectModel(id string, platformName EPlatformType, ) *BackupObjectModel`

NewBackupObjectModel instantiates a new BackupObjectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupObjectModelWithDefaults

`func NewBackupObjectModelWithDefaults() *BackupObjectModel`

NewBackupObjectModelWithDefaults instantiates a new BackupObjectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupObjectModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupObjectModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupObjectModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BackupObjectModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupObjectModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupObjectModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BackupObjectModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *BackupObjectModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupObjectModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupObjectModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BackupObjectModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPlatformName

`func (o *BackupObjectModel) GetPlatformName() EPlatformType`

GetPlatformName returns the PlatformName field if non-nil, zero value otherwise.

### GetPlatformNameOk

`func (o *BackupObjectModel) GetPlatformNameOk() (*EPlatformType, bool)`

GetPlatformNameOk returns a tuple with the PlatformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformName

`func (o *BackupObjectModel) SetPlatformName(v EPlatformType)`

SetPlatformName sets PlatformName field to given value.


### GetPlatformId

`func (o *BackupObjectModel) GetPlatformId() string`

GetPlatformId returns the PlatformId field if non-nil, zero value otherwise.

### GetPlatformIdOk

`func (o *BackupObjectModel) GetPlatformIdOk() (*string, bool)`

GetPlatformIdOk returns a tuple with the PlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformId

`func (o *BackupObjectModel) SetPlatformId(v string)`

SetPlatformId sets PlatformId field to given value.

### HasPlatformId

`func (o *BackupObjectModel) HasPlatformId() bool`

HasPlatformId returns a boolean if a field has been set.

### GetRestorePointsCount

`func (o *BackupObjectModel) GetRestorePointsCount() int32`

GetRestorePointsCount returns the RestorePointsCount field if non-nil, zero value otherwise.

### GetRestorePointsCountOk

`func (o *BackupObjectModel) GetRestorePointsCountOk() (*int32, bool)`

GetRestorePointsCountOk returns a tuple with the RestorePointsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorePointsCount

`func (o *BackupObjectModel) SetRestorePointsCount(v int32)`

SetRestorePointsCount sets RestorePointsCount field to given value.

### HasRestorePointsCount

`func (o *BackupObjectModel) HasRestorePointsCount() bool`

HasRestorePointsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


