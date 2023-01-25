# GuestFileSystemIndexingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, file indexing is enabled. | 
**IndexingSettings** | Pointer to [**[]BackupIndexingSettingsModel**](BackupIndexingSettingsModel.md) | Array of VMs with guest OS file indexing options. | [optional] 

## Methods

### NewGuestFileSystemIndexingModel

`func NewGuestFileSystemIndexingModel(isEnabled bool, ) *GuestFileSystemIndexingModel`

NewGuestFileSystemIndexingModel instantiates a new GuestFileSystemIndexingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestFileSystemIndexingModelWithDefaults

`func NewGuestFileSystemIndexingModelWithDefaults() *GuestFileSystemIndexingModel`

NewGuestFileSystemIndexingModelWithDefaults instantiates a new GuestFileSystemIndexingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *GuestFileSystemIndexingModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *GuestFileSystemIndexingModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *GuestFileSystemIndexingModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetIndexingSettings

`func (o *GuestFileSystemIndexingModel) GetIndexingSettings() []BackupIndexingSettingsModel`

GetIndexingSettings returns the IndexingSettings field if non-nil, zero value otherwise.

### GetIndexingSettingsOk

`func (o *GuestFileSystemIndexingModel) GetIndexingSettingsOk() (*[]BackupIndexingSettingsModel, bool)`

GetIndexingSettingsOk returns a tuple with the IndexingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingSettings

`func (o *GuestFileSystemIndexingModel) SetIndexingSettings(v []BackupIndexingSettingsModel)`

SetIndexingSettings sets IndexingSettings field to given value.

### HasIndexingSettings

`func (o *GuestFileSystemIndexingModel) HasIndexingSettings() bool`

HasIndexingSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


