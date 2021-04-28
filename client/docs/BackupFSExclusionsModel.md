# BackupFSExclusionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExclusionPolicy** | [**EBackupExclusionPolicy**](EBackupExclusionPolicy.md) |  | 
**ItemsList** | Pointer to **[]string** | Array of files and folders. Full paths to files and folders, environmental variables and file masks with the asterisk (*) and question mark (?) characters can be used. | [optional] 

## Methods

### NewBackupFSExclusionsModel

`func NewBackupFSExclusionsModel(exclusionPolicy EBackupExclusionPolicy, ) *BackupFSExclusionsModel`

NewBackupFSExclusionsModel instantiates a new BackupFSExclusionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupFSExclusionsModelWithDefaults

`func NewBackupFSExclusionsModelWithDefaults() *BackupFSExclusionsModel`

NewBackupFSExclusionsModelWithDefaults instantiates a new BackupFSExclusionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclusionPolicy

`func (o *BackupFSExclusionsModel) GetExclusionPolicy() EBackupExclusionPolicy`

GetExclusionPolicy returns the ExclusionPolicy field if non-nil, zero value otherwise.

### GetExclusionPolicyOk

`func (o *BackupFSExclusionsModel) GetExclusionPolicyOk() (*EBackupExclusionPolicy, bool)`

GetExclusionPolicyOk returns a tuple with the ExclusionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionPolicy

`func (o *BackupFSExclusionsModel) SetExclusionPolicy(v EBackupExclusionPolicy)`

SetExclusionPolicy sets ExclusionPolicy field to given value.


### GetItemsList

`func (o *BackupFSExclusionsModel) GetItemsList() []string`

GetItemsList returns the ItemsList field if non-nil, zero value otherwise.

### GetItemsListOk

`func (o *BackupFSExclusionsModel) GetItemsListOk() (*[]string, bool)`

GetItemsListOk returns a tuple with the ItemsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsList

`func (o *BackupFSExclusionsModel) SetItemsList(v []string)`

SetItemsList sets ItemsList field to given value.

### HasItemsList

`func (o *BackupFSExclusionsModel) HasItemsList() bool`

HasItemsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


