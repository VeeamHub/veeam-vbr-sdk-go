# BackupJobExclusionsTemplates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | If *true*, VM templates are excluded from the backup. | [optional] 
**ExcludeFromIncremental** | Pointer to **bool** | If *true*, VM templates are excluded from the incremental backup. | [optional] 

## Methods

### NewBackupJobExclusionsTemplates

`func NewBackupJobExclusionsTemplates() *BackupJobExclusionsTemplates`

NewBackupJobExclusionsTemplates instantiates a new BackupJobExclusionsTemplates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobExclusionsTemplatesWithDefaults

`func NewBackupJobExclusionsTemplatesWithDefaults() *BackupJobExclusionsTemplates`

NewBackupJobExclusionsTemplatesWithDefaults instantiates a new BackupJobExclusionsTemplates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *BackupJobExclusionsTemplates) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *BackupJobExclusionsTemplates) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *BackupJobExclusionsTemplates) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *BackupJobExclusionsTemplates) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetExcludeFromIncremental

`func (o *BackupJobExclusionsTemplates) GetExcludeFromIncremental() bool`

GetExcludeFromIncremental returns the ExcludeFromIncremental field if non-nil, zero value otherwise.

### GetExcludeFromIncrementalOk

`func (o *BackupJobExclusionsTemplates) GetExcludeFromIncrementalOk() (*bool, bool)`

GetExcludeFromIncrementalOk returns a tuple with the ExcludeFromIncremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromIncremental

`func (o *BackupJobExclusionsTemplates) SetExcludeFromIncremental(v bool)`

SetExcludeFromIncremental sets ExcludeFromIncremental field to given value.

### HasExcludeFromIncremental

`func (o *BackupJobExclusionsTemplates) HasExcludeFromIncremental() bool`

HasExcludeFromIncremental returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


