# BackupJobGuestProcessingImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationAwareProcessing** | [**BackupApplicationAwareProcessingImportModel**](BackupApplicationAwareProcessingImportModel.md) |  | 
**GuestFileSystemIndexing** | [**GuestFileSystemIndexingModel**](GuestFileSystemIndexingModel.md) |  | 
**GuestInteractionProxies** | Pointer to [**GuestInteractionProxiesSettingsImportModel**](GuestInteractionProxiesSettingsImportModel.md) |  | [optional] 
**GuestCredentials** | Pointer to [**GuestOsCredentialsImportModel**](GuestOsCredentialsImportModel.md) |  | [optional] 

## Methods

### NewBackupJobGuestProcessingImportModel

`func NewBackupJobGuestProcessingImportModel(applicationAwareProcessing BackupApplicationAwareProcessingImportModel, guestFileSystemIndexing GuestFileSystemIndexingModel, ) *BackupJobGuestProcessingImportModel`

NewBackupJobGuestProcessingImportModel instantiates a new BackupJobGuestProcessingImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobGuestProcessingImportModelWithDefaults

`func NewBackupJobGuestProcessingImportModelWithDefaults() *BackupJobGuestProcessingImportModel`

NewBackupJobGuestProcessingImportModelWithDefaults instantiates a new BackupJobGuestProcessingImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationAwareProcessing

`func (o *BackupJobGuestProcessingImportModel) GetApplicationAwareProcessing() BackupApplicationAwareProcessingImportModel`

GetApplicationAwareProcessing returns the ApplicationAwareProcessing field if non-nil, zero value otherwise.

### GetApplicationAwareProcessingOk

`func (o *BackupJobGuestProcessingImportModel) GetApplicationAwareProcessingOk() (*BackupApplicationAwareProcessingImportModel, bool)`

GetApplicationAwareProcessingOk returns a tuple with the ApplicationAwareProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAwareProcessing

`func (o *BackupJobGuestProcessingImportModel) SetApplicationAwareProcessing(v BackupApplicationAwareProcessingImportModel)`

SetApplicationAwareProcessing sets ApplicationAwareProcessing field to given value.


### GetGuestFileSystemIndexing

`func (o *BackupJobGuestProcessingImportModel) GetGuestFileSystemIndexing() GuestFileSystemIndexingModel`

GetGuestFileSystemIndexing returns the GuestFileSystemIndexing field if non-nil, zero value otherwise.

### GetGuestFileSystemIndexingOk

`func (o *BackupJobGuestProcessingImportModel) GetGuestFileSystemIndexingOk() (*GuestFileSystemIndexingModel, bool)`

GetGuestFileSystemIndexingOk returns a tuple with the GuestFileSystemIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestFileSystemIndexing

`func (o *BackupJobGuestProcessingImportModel) SetGuestFileSystemIndexing(v GuestFileSystemIndexingModel)`

SetGuestFileSystemIndexing sets GuestFileSystemIndexing field to given value.


### GetGuestInteractionProxies

`func (o *BackupJobGuestProcessingImportModel) GetGuestInteractionProxies() GuestInteractionProxiesSettingsImportModel`

GetGuestInteractionProxies returns the GuestInteractionProxies field if non-nil, zero value otherwise.

### GetGuestInteractionProxiesOk

`func (o *BackupJobGuestProcessingImportModel) GetGuestInteractionProxiesOk() (*GuestInteractionProxiesSettingsImportModel, bool)`

GetGuestInteractionProxiesOk returns a tuple with the GuestInteractionProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestInteractionProxies

`func (o *BackupJobGuestProcessingImportModel) SetGuestInteractionProxies(v GuestInteractionProxiesSettingsImportModel)`

SetGuestInteractionProxies sets GuestInteractionProxies field to given value.

### HasGuestInteractionProxies

`func (o *BackupJobGuestProcessingImportModel) HasGuestInteractionProxies() bool`

HasGuestInteractionProxies returns a boolean if a field has been set.

### GetGuestCredentials

`func (o *BackupJobGuestProcessingImportModel) GetGuestCredentials() GuestOsCredentialsImportModel`

GetGuestCredentials returns the GuestCredentials field if non-nil, zero value otherwise.

### GetGuestCredentialsOk

`func (o *BackupJobGuestProcessingImportModel) GetGuestCredentialsOk() (*GuestOsCredentialsImportModel, bool)`

GetGuestCredentialsOk returns a tuple with the GuestCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCredentials

`func (o *BackupJobGuestProcessingImportModel) SetGuestCredentials(v GuestOsCredentialsImportModel)`

SetGuestCredentials sets GuestCredentials field to given value.

### HasGuestCredentials

`func (o *BackupJobGuestProcessingImportModel) HasGuestCredentials() bool`

HasGuestCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


