# BackupPlacementSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtentId** | **string** | ID of a performance extent. | 
**AllowedBackups** | [**EAllowedBackupsType**](EAllowedBackupsType.md) |  | 

## Methods

### NewBackupPlacementSettingsModel

`func NewBackupPlacementSettingsModel(extentId string, allowedBackups EAllowedBackupsType, ) *BackupPlacementSettingsModel`

NewBackupPlacementSettingsModel instantiates a new BackupPlacementSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupPlacementSettingsModelWithDefaults

`func NewBackupPlacementSettingsModelWithDefaults() *BackupPlacementSettingsModel`

NewBackupPlacementSettingsModelWithDefaults instantiates a new BackupPlacementSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtentId

`func (o *BackupPlacementSettingsModel) GetExtentId() string`

GetExtentId returns the ExtentId field if non-nil, zero value otherwise.

### GetExtentIdOk

`func (o *BackupPlacementSettingsModel) GetExtentIdOk() (*string, bool)`

GetExtentIdOk returns a tuple with the ExtentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtentId

`func (o *BackupPlacementSettingsModel) SetExtentId(v string)`

SetExtentId sets ExtentId field to given value.


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


