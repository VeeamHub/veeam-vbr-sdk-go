# GeneralOptionsEmailNotificationsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, global email notification settings are enabled. | 
**SmtpServerName** | **string** | Full DNS name or IP address of the SMTP server. | 
**AdvancedSmtpOptions** | [**AdvancedSmtpOptionsModel**](AdvancedSmtpOptionsModel.md) |  | 
**From** | **string** | Email address from which email notifications must be sent. | 
**To** | **string** | Recipient email addresses. Use a semicolon to separate multiple addresses. | 
**Subject** | **string** | Notification subject. Use the following variables in the subject:&lt;ul&gt; &lt;li&gt;%Time% — completion time&lt;/li&gt; &lt;li&gt;%JobName% — job name&lt;/li&gt; &lt;li&gt;%JobResult% — job result&lt;/li&gt; &lt;li&gt;%ObjectCount% — number of VMs in the job&lt;/li&gt; &lt;li&gt;%Issues% — number of VMs in the job that have been processed with the Warning or Failed status&lt;/li&gt;&lt;/ul&gt; | 
**SendDailyReportsAt** | **time.Time** | Time when Veeam Backup &amp; Replication sends daily email reports. | 
**NotifyOnSuccess** | **bool** | If *true*, email notifications are sent when the job completes successfully. | 
**NotifyOnWarning** | **bool** | If *true*, email notifications are sent when the job completes with a warning. | 
**NotifyOnFailure** | **bool** | If *true*, email notifications are sent when the job fails. | 
**NotifyOnLastRetry** | **bool** | If *true*, email notifications are sent about the final job status only (not per every job retry). | 

## Methods

### NewGeneralOptionsEmailNotificationsModel

`func NewGeneralOptionsEmailNotificationsModel(isEnabled bool, smtpServerName string, advancedSmtpOptions AdvancedSmtpOptionsModel, from string, to string, subject string, sendDailyReportsAt time.Time, notifyOnSuccess bool, notifyOnWarning bool, notifyOnFailure bool, notifyOnLastRetry bool, ) *GeneralOptionsEmailNotificationsModel`

NewGeneralOptionsEmailNotificationsModel instantiates a new GeneralOptionsEmailNotificationsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralOptionsEmailNotificationsModelWithDefaults

`func NewGeneralOptionsEmailNotificationsModelWithDefaults() *GeneralOptionsEmailNotificationsModel`

NewGeneralOptionsEmailNotificationsModelWithDefaults instantiates a new GeneralOptionsEmailNotificationsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *GeneralOptionsEmailNotificationsModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *GeneralOptionsEmailNotificationsModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *GeneralOptionsEmailNotificationsModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetSmtpServerName

`func (o *GeneralOptionsEmailNotificationsModel) GetSmtpServerName() string`

GetSmtpServerName returns the SmtpServerName field if non-nil, zero value otherwise.

### GetSmtpServerNameOk

`func (o *GeneralOptionsEmailNotificationsModel) GetSmtpServerNameOk() (*string, bool)`

GetSmtpServerNameOk returns a tuple with the SmtpServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServerName

`func (o *GeneralOptionsEmailNotificationsModel) SetSmtpServerName(v string)`

SetSmtpServerName sets SmtpServerName field to given value.


### GetAdvancedSmtpOptions

`func (o *GeneralOptionsEmailNotificationsModel) GetAdvancedSmtpOptions() AdvancedSmtpOptionsModel`

GetAdvancedSmtpOptions returns the AdvancedSmtpOptions field if non-nil, zero value otherwise.

### GetAdvancedSmtpOptionsOk

`func (o *GeneralOptionsEmailNotificationsModel) GetAdvancedSmtpOptionsOk() (*AdvancedSmtpOptionsModel, bool)`

GetAdvancedSmtpOptionsOk returns a tuple with the AdvancedSmtpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSmtpOptions

`func (o *GeneralOptionsEmailNotificationsModel) SetAdvancedSmtpOptions(v AdvancedSmtpOptionsModel)`

SetAdvancedSmtpOptions sets AdvancedSmtpOptions field to given value.


### GetFrom

`func (o *GeneralOptionsEmailNotificationsModel) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GeneralOptionsEmailNotificationsModel) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GeneralOptionsEmailNotificationsModel) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *GeneralOptionsEmailNotificationsModel) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GeneralOptionsEmailNotificationsModel) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GeneralOptionsEmailNotificationsModel) SetTo(v string)`

SetTo sets To field to given value.


### GetSubject

`func (o *GeneralOptionsEmailNotificationsModel) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GeneralOptionsEmailNotificationsModel) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GeneralOptionsEmailNotificationsModel) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetSendDailyReportsAt

`func (o *GeneralOptionsEmailNotificationsModel) GetSendDailyReportsAt() time.Time`

GetSendDailyReportsAt returns the SendDailyReportsAt field if non-nil, zero value otherwise.

### GetSendDailyReportsAtOk

`func (o *GeneralOptionsEmailNotificationsModel) GetSendDailyReportsAtOk() (*time.Time, bool)`

GetSendDailyReportsAtOk returns a tuple with the SendDailyReportsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDailyReportsAt

`func (o *GeneralOptionsEmailNotificationsModel) SetSendDailyReportsAt(v time.Time)`

SetSendDailyReportsAt sets SendDailyReportsAt field to given value.


### GetNotifyOnSuccess

`func (o *GeneralOptionsEmailNotificationsModel) GetNotifyOnSuccess() bool`

GetNotifyOnSuccess returns the NotifyOnSuccess field if non-nil, zero value otherwise.

### GetNotifyOnSuccessOk

`func (o *GeneralOptionsEmailNotificationsModel) GetNotifyOnSuccessOk() (*bool, bool)`

GetNotifyOnSuccessOk returns a tuple with the NotifyOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnSuccess

`func (o *GeneralOptionsEmailNotificationsModel) SetNotifyOnSuccess(v bool)`

SetNotifyOnSuccess sets NotifyOnSuccess field to given value.


### GetNotifyOnWarning

`func (o *GeneralOptionsEmailNotificationsModel) GetNotifyOnWarning() bool`

GetNotifyOnWarning returns the NotifyOnWarning field if non-nil, zero value otherwise.

### GetNotifyOnWarningOk

`func (o *GeneralOptionsEmailNotificationsModel) GetNotifyOnWarningOk() (*bool, bool)`

GetNotifyOnWarningOk returns a tuple with the NotifyOnWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnWarning

`func (o *GeneralOptionsEmailNotificationsModel) SetNotifyOnWarning(v bool)`

SetNotifyOnWarning sets NotifyOnWarning field to given value.


### GetNotifyOnFailure

`func (o *GeneralOptionsEmailNotificationsModel) GetNotifyOnFailure() bool`

GetNotifyOnFailure returns the NotifyOnFailure field if non-nil, zero value otherwise.

### GetNotifyOnFailureOk

`func (o *GeneralOptionsEmailNotificationsModel) GetNotifyOnFailureOk() (*bool, bool)`

GetNotifyOnFailureOk returns a tuple with the NotifyOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnFailure

`func (o *GeneralOptionsEmailNotificationsModel) SetNotifyOnFailure(v bool)`

SetNotifyOnFailure sets NotifyOnFailure field to given value.


### GetNotifyOnLastRetry

`func (o *GeneralOptionsEmailNotificationsModel) GetNotifyOnLastRetry() bool`

GetNotifyOnLastRetry returns the NotifyOnLastRetry field if non-nil, zero value otherwise.

### GetNotifyOnLastRetryOk

`func (o *GeneralOptionsEmailNotificationsModel) GetNotifyOnLastRetryOk() (*bool, bool)`

GetNotifyOnLastRetryOk returns a tuple with the NotifyOnLastRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnLastRetry

`func (o *GeneralOptionsEmailNotificationsModel) SetNotifyOnLastRetry(v bool)`

SetNotifyOnLastRetry sets NotifyOnLastRetry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


