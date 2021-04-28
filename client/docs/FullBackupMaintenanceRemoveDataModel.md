# FullBackupMaintenanceRemoveDataModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, Veeam Backup &amp; Replication keeps the backup data of deleted VMs for the &#x60;afterDays&#x60; number of days. | 
**AfterDays** | Pointer to **int32** | Number of days. | [optional] 

## Methods

### NewFullBackupMaintenanceRemoveDataModel

`func NewFullBackupMaintenanceRemoveDataModel(isEnabled bool, ) *FullBackupMaintenanceRemoveDataModel`

NewFullBackupMaintenanceRemoveDataModel instantiates a new FullBackupMaintenanceRemoveDataModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullBackupMaintenanceRemoveDataModelWithDefaults

`func NewFullBackupMaintenanceRemoveDataModelWithDefaults() *FullBackupMaintenanceRemoveDataModel`

NewFullBackupMaintenanceRemoveDataModelWithDefaults instantiates a new FullBackupMaintenanceRemoveDataModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *FullBackupMaintenanceRemoveDataModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FullBackupMaintenanceRemoveDataModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FullBackupMaintenanceRemoveDataModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetAfterDays

`func (o *FullBackupMaintenanceRemoveDataModel) GetAfterDays() int32`

GetAfterDays returns the AfterDays field if non-nil, zero value otherwise.

### GetAfterDaysOk

`func (o *FullBackupMaintenanceRemoveDataModel) GetAfterDaysOk() (*int32, bool)`

GetAfterDaysOk returns a tuple with the AfterDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterDays

`func (o *FullBackupMaintenanceRemoveDataModel) SetAfterDays(v int32)`

SetAfterDays sets AfterDays field to given value.

### HasAfterDays

`func (o *FullBackupMaintenanceRemoveDataModel) HasAfterDays() bool`

HasAfterDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


