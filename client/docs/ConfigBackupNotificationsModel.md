# ConfigBackupNotificationsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNMPEnabled** | **bool** | If *true*, SNMP traps are enabled for this job. | 
**SMTPSettings** | Pointer to [**ConfigBackupSMTPSettigsModel**](ConfigBackupSMTPSettigsModel.md) |  | [optional] 

## Methods

### NewConfigBackupNotificationsModel

`func NewConfigBackupNotificationsModel(sNMPEnabled bool, ) *ConfigBackupNotificationsModel`

NewConfigBackupNotificationsModel instantiates a new ConfigBackupNotificationsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigBackupNotificationsModelWithDefaults

`func NewConfigBackupNotificationsModelWithDefaults() *ConfigBackupNotificationsModel`

NewConfigBackupNotificationsModelWithDefaults instantiates a new ConfigBackupNotificationsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNMPEnabled

`func (o *ConfigBackupNotificationsModel) GetSNMPEnabled() bool`

GetSNMPEnabled returns the SNMPEnabled field if non-nil, zero value otherwise.

### GetSNMPEnabledOk

`func (o *ConfigBackupNotificationsModel) GetSNMPEnabledOk() (*bool, bool)`

GetSNMPEnabledOk returns a tuple with the SNMPEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNMPEnabled

`func (o *ConfigBackupNotificationsModel) SetSNMPEnabled(v bool)`

SetSNMPEnabled sets SNMPEnabled field to given value.


### GetSMTPSettings

`func (o *ConfigBackupNotificationsModel) GetSMTPSettings() ConfigBackupSMTPSettigsModel`

GetSMTPSettings returns the SMTPSettings field if non-nil, zero value otherwise.

### GetSMTPSettingsOk

`func (o *ConfigBackupNotificationsModel) GetSMTPSettingsOk() (*ConfigBackupSMTPSettigsModel, bool)`

GetSMTPSettingsOk returns a tuple with the SMTPSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMTPSettings

`func (o *ConfigBackupNotificationsModel) SetSMTPSettings(v ConfigBackupSMTPSettigsModel)`

SetSMTPSettings sets SMTPSettings field to given value.

### HasSMTPSettings

`func (o *ConfigBackupNotificationsModel) HasSMTPSettings() bool`

HasSMTPSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


