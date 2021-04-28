# BackupPlacementSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtentName** | **string** | Name of a performance extent. | 
**AllowedBackups** | [**EAllowedBackupsType**](EAllowedBackupsType.md) |  | 

## Methods

### NewBackupPlacementSettingsModel

`func NewBackupPlacementSettingsModel(extentName string, allowedBackups EAllowedBackupsType, ) *BackupPlacementSettingsModel`

NewBackupPlacementSettingsModel instantiates a new BackupPlacementSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPlacementSettingsModelWithDefaults

`func NewBackupPlacementSettingsModelWithDefaults() *BackupPlacementSettingsModel`

NewBackupPlacementSettingsModelWithDefaults instantiates a new BackupPlacementSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtentName

`func (o *BackupPlacementSettingsModel) GetExtentName() string`

GetExtentName returns the ExtentName field if non-nil, zero value otherwise.

### GetExtentNameOk

`func (o *BackupPlacementSettingsModel) GetExtentNameOk() (*string, bool)`

GetExtentNameOk returns a tuple with the ExtentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtentName

`func (o *BackupPlacementSettingsModel) SetExtentName(v string)`

SetExtentName sets ExtentName field to given value.


### GetAllowedBackups

`func (o *BackupPlacementSettingsModel) GetAllowedBackups() EAllowedBackupsType`

GetAllowedBackups returns the AllowedBackups field if non-nil, zero value otherwise.

### GetAllowedBackupsOk

`func (o *BackupPlacementSettingsModel) GetAllowedBackupsOk() (*EAllowedBackupsType, bool)`

GetAllowedBackupsOk returns a tuple with the AllowedBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedBackups

`func (o *BackupPlacementSettingsModel) SetAllowedBackups(v EAllowedBackupsType)`

SetAllowedBackups sets AllowedBackups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


