# ConfigBackupSMTPSettigsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, email notifications are enabled for this job. | 
**Recipients** | **[]string** | Array of recipients&#39; email addresses. | 
**SettingsType** | [**EConfigBackupSMTPSettingsType**](EConfigBackupSMTPSettingsType.md) |  | 
**Subject** | **string** | Notification subject. Use the following variables in the subject:&lt;ul&gt; &lt;li&gt;*%Time%* — completion time&lt;/li&gt; &lt;li&gt;*%JobName%* — job name&lt;/li&gt; &lt;li&gt;*%JobResult%* — job result&lt;/li&gt;&lt;/ul&gt; | 
**NotifyOnSuccess** | **bool** | If *true*, email notifications are sent when the job completes successfully. | 
**NotifyOnWarning** | **bool** | If *true*, email notifications are sent when the job completes with a warning. | 
**NotifyOnError** | **bool** | If *true*, email notifications are sent when the job fails. | 

## Methods

### NewConfigBackupSMTPSettigsModel

`func NewConfigBackupSMTPSettigsModel(isEnabled bool, recipients []string, settingsType EConfigBackupSMTPSettingsType, subject string, notifyOnSuccess bool, notifyOnWarning bool, notifyOnError bool, ) *ConfigBackupSMTPSettigsModel`

NewConfigBackupSMTPSettigsModel instantiates a new ConfigBackupSMTPSettigsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigBackupSMTPSettigsModelWithDefaults

`func NewConfigBackupSMTPSettigsModelWithDefaults() *ConfigBackupSMTPSettigsModel`

NewConfigBackupSMTPSettigsModelWithDefaults instantiates a new ConfigBackupSMTPSettigsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ConfigBackupSMTPSettigsModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ConfigBackupSMTPSettigsModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ConfigBackupSMTPSettigsModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetRecipients

`func (o *ConfigBackupSMTPSettigsModel) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ConfigBackupSMTPSettigsModel) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ConfigBackupSMTPSettigsModel) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.


### GetSettingsType

`func (o *ConfigBackupSMTPSettigsModel) GetSettingsType() EConfigBackupSMTPSettingsType`

GetSettingsType returns the SettingsType field if non-nil, zero value otherwise.

### GetSettingsTypeOk

`func (o *ConfigBackupSMTPSettigsModel) GetSettingsTypeOk() (*EConfigBackupSMTPSettingsType, bool)`

GetSettingsTypeOk returns a tuple with the SettingsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsType

`func (o *ConfigBackupSMTPSettigsModel) SetSettingsType(v EConfigBackupSMTPSettingsType)`

SetSettingsType sets SettingsType field to given value.


### GetSubject

`func (o *ConfigBackupSMTPSettigsModel) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ConfigBackupSMTPSettigsModel) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ConfigBackupSMTPSettigsModel) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetNotifyOnSuccess

`func (o *ConfigBackupSMTPSettigsModel) GetNotifyOnSuccess() bool`

GetNotifyOnSuccess returns the NotifyOnSuccess field if non-nil, zero value otherwise.

### GetNotifyOnSuccessOk

`func (o *ConfigBackupSMTPSettigsModel) GetNotifyOnSuccessOk() (*bool, bool)`

GetNotifyOnSuccessOk returns a tuple with the NotifyOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnSuccess

`func (o *ConfigBackupSMTPSettigsModel) SetNotifyOnSuccess(v bool)`

SetNotifyOnSuccess sets NotifyOnSuccess field to given value.


### GetNotifyOnWarning

`func (o *ConfigBackupSMTPSettigsModel) GetNotifyOnWarning() bool`

GetNotifyOnWarning returns the NotifyOnWarning field if non-nil, zero value otherwise.

### GetNotifyOnWarningOk

`func (o *ConfigBackupSMTPSettigsModel) GetNotifyOnWarningOk() (*bool, bool)`

GetNotifyOnWarningOk returns a tuple with the NotifyOnWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnWarning

`func (o *ConfigBackupSMTPSettigsModel) SetNotifyOnWarning(v bool)`

SetNotifyOnWarning sets NotifyOnWarning field to given value.


### GetNotifyOnError

`func (o *ConfigBackupSMTPSettigsModel) GetNotifyOnError() bool`

GetNotifyOnError returns the NotifyOnError field if non-nil, zero value otherwise.

### GetNotifyOnErrorOk

`func (o *ConfigBackupSMTPSettigsModel) GetNotifyOnErrorOk() (*bool, bool)`

GetNotifyOnErrorOk returns a tuple with the NotifyOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnError

`func (o *ConfigBackupSMTPSettigsModel) SetNotifyOnError(v bool)`

SetNotifyOnError sets NotifyOnError field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


