# BackupJobGuestProcessingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppAwareProcessing** | [**BackupApplicationAwareProcessingModel**](BackupApplicationAwareProcessingModel.md) |  | 
**GuestFSIndexing** | [**GuestFileSystemIndexingModel**](GuestFileSystemIndexingModel.md) |  | 
**GuestInteractionProxies** | Pointer to [**GuestInteractionProxiesSettingsModel**](GuestInteractionProxiesSettingsModel.md) |  | [optional] 
**GuestCredentials** | Pointer to [**GuestOsCredentialsModel**](GuestOsCredentialsModel.md) |  | [optional] 

## Methods

### NewBackupJobGuestProcessingModel

`func NewBackupJobGuestProcessingModel(appAwareProcessing BackupApplicationAwareProcessingModel, guestFSIndexing GuestFileSystemIndexingModel, ) *BackupJobGuestProcessingModel`

NewBackupJobGuestProcessingModel instantiates a new BackupJobGuestProcessingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobGuestProcessingModelWithDefaults

`func NewBackupJobGuestProcessingModelWithDefaults() *BackupJobGuestProcessingModel`

NewBackupJobGuestProcessingModelWithDefaults instantiates a new BackupJobGuestProcessingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppAwareProcessing

`func (o *BackupJobGuestProcessingModel) GetAppAwareProcessing() BackupApplicationAwareProcessingModel`

GetAppAwareProcessing returns the AppAwareProcessing field if non-nil, zero value otherwise.

### GetAppAwareProcessingOk

`func (o *BackupJobGuestProcessingModel) GetAppAwareProcessingOk() (*BackupApplicationAwareProcessingModel, bool)`

GetAppAwareProcessingOk returns a tuple with the AppAwareProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAwareProcessing

`func (o *BackupJobGuestProcessingModel) SetAppAwareProcessing(v BackupApplicationAwareProcessingModel)`

SetAppAwareProcessing sets AppAwareProcessing field to given value.


### GetGuestFSIndexing

`func (o *BackupJobGuestProcessingModel) GetGuestFSIndexing() GuestFileSystemIndexingModel`

GetGuestFSIndexing returns the GuestFSIndexing field if non-nil, zero value otherwise.

### GetGuestFSIndexingOk

`func (o *BackupJobGuestProcessingModel) GetGuestFSIndexingOk() (*GuestFileSystemIndexingModel, bool)`

GetGuestFSIndexingOk returns a tuple with the GuestFSIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestFSIndexing

`func (o *BackupJobGuestProcessingModel) SetGuestFSIndexing(v GuestFileSystemIndexingModel)`

SetGuestFSIndexing sets GuestFSIndexing field to given value.


### GetGuestInteractionProxies

`func (o *BackupJobGuestProcessingModel) GetGuestInteractionProxies() GuestInteractionProxiesSettingsModel`

GetGuestInteractionProxies returns the GuestInteractionProxies field if non-nil, zero value otherwise.

### GetGuestInteractionProxiesOk

`func (o *BackupJobGuestProcessingModel) GetGuestInteractionProxiesOk() (*GuestInteractionProxiesSettingsModel, bool)`

GetGuestInteractionProxiesOk returns a tuple with the GuestInteractionProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestInteractionProxies

`func (o *BackupJobGuestProcessingModel) SetGuestInteractionProxies(v GuestInteractionProxiesSettingsModel)`

SetGuestInteractionProxies sets GuestInteractionProxies field to given value.

### HasGuestInteractionProxies

`func (o *BackupJobGuestProcessingModel) HasGuestInteractionProxies() bool`

HasGuestInteractionProxies returns a boolean if a field has been set.

### GetGuestCredentials

`func (o *BackupJobGuestProcessingModel) GetGuestCredentials() GuestOsCredentialsModel`

GetGuestCredentials returns the GuestCredentials field if non-nil, zero value otherwise.

### GetGuestCredentialsOk

`func (o *BackupJobGuestProcessingModel) GetGuestCredentialsOk() (*GuestOsCredentialsModel, bool)`

GetGuestCredentialsOk returns a tuple with the GuestCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCredentials

`func (o *BackupJobGuestProcessingModel) SetGuestCredentials(v GuestOsCredentialsModel)`

SetGuestCredentials sets GuestCredentials field to given value.

### HasGuestCredentials

`func (o *BackupJobGuestProcessingModel) HasGuestCredentials() bool`

HasGuestCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


