# BackupWindowsScriptModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreFreezeScript** | Pointer to **string** | Path to a pre-freeze script. | [optional] 
**PostThawScript** | Pointer to **string** | Path to a post-thaw script. | [optional] 

## Methods

### NewBackupWindowsScriptModel

`func NewBackupWindowsScriptModel() *BackupWindowsScriptModel`

NewBackupWindowsScriptModel instantiates a new BackupWindowsScriptModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupWindowsScriptModelWithDefaults

`func NewBackupWindowsScriptModelWithDefaults() *BackupWindowsScriptModel`

NewBackupWindowsScriptModelWithDefaults instantiates a new BackupWindowsScriptModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreFreezeScript

`func (o *BackupWindowsScriptModel) GetPreFreezeScript() string`

GetPreFreezeScript returns the PreFreezeScript field if non-nil, zero value otherwise.

### GetPreFreezeScriptOk

`func (o *BackupWindowsScriptModel) GetPreFreezeScriptOk() (*string, bool)`

GetPreFreezeScriptOk returns a tuple with the PreFreezeScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreFreezeScript

`func (o *BackupWindowsScriptModel) SetPreFreezeScript(v string)`

SetPreFreezeScript sets PreFreezeScript field to given value.

### HasPreFreezeScript

`func (o *BackupWindowsScriptModel) HasPreFreezeScript() bool`

HasPreFreezeScript returns a boolean if a field has been set.

### GetPostThawScript

`func (o *BackupWindowsScriptModel) GetPostThawScript() string`

GetPostThawScript returns the PostThawScript field if non-nil, zero value otherwise.

### GetPostThawScriptOk

`func (o *BackupWindowsScriptModel) GetPostThawScriptOk() (*string, bool)`

GetPostThawScriptOk returns a tuple with the PostThawScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostThawScript

`func (o *BackupWindowsScriptModel) SetPostThawScript(v string)`

SetPostThawScript sets PostThawScript field to given value.

### HasPostThawScript

`func (o *BackupWindowsScriptModel) HasPostThawScript() bool`

HasPostThawScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


