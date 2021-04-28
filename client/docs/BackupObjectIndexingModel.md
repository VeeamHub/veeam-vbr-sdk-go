# BackupObjectIndexingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuestFSIndexingMode** | [**EGuestFSIndexingMode**](EGuestFSIndexingMode.md) |  | 
**IndexingList** | Pointer to **[]string** | Array of folders. Environmental variables and full paths to folders can be used. | [optional] 

## Methods

### NewBackupObjectIndexingModel

`func NewBackupObjectIndexingModel(guestFSIndexingMode EGuestFSIndexingMode, ) *BackupObjectIndexingModel`

NewBackupObjectIndexingModel instantiates a new BackupObjectIndexingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupObjectIndexingModelWithDefaults

`func NewBackupObjectIndexingModelWithDefaults() *BackupObjectIndexingModel`

NewBackupObjectIndexingModelWithDefaults instantiates a new BackupObjectIndexingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuestFSIndexingMode

`func (o *BackupObjectIndexingModel) GetGuestFSIndexingMode() EGuestFSIndexingMode`

GetGuestFSIndexingMode returns the GuestFSIndexingMode field if non-nil, zero value otherwise.

### GetGuestFSIndexingModeOk

`func (o *BackupObjectIndexingModel) GetGuestFSIndexingModeOk() (*EGuestFSIndexingMode, bool)`

GetGuestFSIndexingModeOk returns a tuple with the GuestFSIndexingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestFSIndexingMode

`func (o *BackupObjectIndexingModel) SetGuestFSIndexingMode(v EGuestFSIndexingMode)`

SetGuestFSIndexingMode sets GuestFSIndexingMode field to given value.


### GetIndexingList

`func (o *BackupObjectIndexingModel) GetIndexingList() []string`

GetIndexingList returns the IndexingList field if non-nil, zero value otherwise.

### GetIndexingListOk

`func (o *BackupObjectIndexingModel) GetIndexingListOk() (*[]string, bool)`

GetIndexingListOk returns a tuple with the IndexingList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingList

`func (o *BackupObjectIndexingModel) SetIndexingList(v []string)`

SetIndexingList sets IndexingList field to given value.

### HasIndexingList

`func (o *BackupObjectIndexingModel) HasIndexingList() bool`

HasIndexingList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


