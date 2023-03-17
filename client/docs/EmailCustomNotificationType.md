# EmailCustomNotificationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** | Notification subject. Use the following variables in the subject:&lt;ul&gt; &lt;li&gt;*%Time%* — completion time&lt;/li&gt; &lt;li&gt;*%JobName%* — job name&lt;/li&gt; &lt;li&gt;*%JobResult%* — job result&lt;/li&gt; &lt;li&gt;*%ObjectCount%* — number of VMs in the job&lt;/li&gt; &lt;li&gt;*%Issues%* — number of VMs in the job that have finished with the Warning or Failed status&lt;/li&gt;&lt;/ul&gt; | [optional] 
**NotifyOnSuccess** | Pointer to **bool** | If *true*, email notifications are sent when the job completes successfully. | [optional] 
**NotifyOnWarning** | Pointer to **bool** | If *true*, email notifications are sent when the job completes with a warning. | [optional] 
**NotifyOnError** | Pointer to **bool** | If *true*, email notifications are sent when the job fails. | [optional] 
**SuppressNotificationUntilLastRetry** | Pointer to **bool** | If *true*, email notifications are sent about the final job status only (not per every job retry). | [optional] 

## Methods

### NewEmailCustomNotificationType

`func NewEmailCustomNotificationType() *EmailCustomNotificationType`

NewEmailCustomNotificationType instantiates a new EmailCustomNotificationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailCustomNotificationTypeWithDefaults

`func NewEmailCustomNotificationTypeWithDefaults() *EmailCustomNotificationType`

NewEmailCustomNotificationTypeWithDefaults instantiates a new EmailCustomNotificationType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *EmailCustomNotificationType) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailCustomNotificationType) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailCustomNotificationType) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailCustomNotificationType) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetNotifyOnSuccess

`func (o *EmailCustomNotificationType) GetNotifyOnSuccess() bool`

GetNotifyOnSuccess returns the NotifyOnSuccess field if non-nil, zero value otherwise.

### GetNotifyOnSuccessOk

`func (o *EmailCustomNotificationType) GetNotifyOnSuccessOk() (*bool, bool)`

GetNotifyOnSuccessOk returns a tuple with the NotifyOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnSuccess

`func (o *EmailCustomNotificationType) SetNotifyOnSuccess(v bool)`

SetNotifyOnSuccess sets NotifyOnSuccess field to given value.

### HasNotifyOnSuccess

`func (o *EmailCustomNotificationType) HasNotifyOnSuccess() bool`

HasNotifyOnSuccess returns a boolean if a field has been set.

### GetNotifyOnWarning

`func (o *EmailCustomNotificationType) GetNotifyOnWarning() bool`

GetNotifyOnWarning returns the NotifyOnWarning field if non-nil, zero value otherwise.

### GetNotifyOnWarningOk

`func (o *EmailCustomNotificationType) GetNotifyOnWarningOk() (*bool, bool)`

GetNotifyOnWarningOk returns a tuple with the NotifyOnWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnWarning

`func (o *EmailCustomNotificationType) SetNotifyOnWarning(v bool)`

SetNotifyOnWarning sets NotifyOnWarning field to given value.

### HasNotifyOnWarning

`func (o *EmailCustomNotificationType) HasNotifyOnWarning() bool`

HasNotifyOnWarning returns a boolean if a field has been set.

### GetNotifyOnError

`func (o *EmailCustomNotificationType) GetNotifyOnError() bool`

GetNotifyOnError returns the NotifyOnError field if non-nil, zero value otherwise.

### GetNotifyOnErrorOk

`func (o *EmailCustomNotificationType) GetNotifyOnErrorOk() (*bool, bool)`

GetNotifyOnErrorOk returns a tuple with the NotifyOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnError

`func (o *EmailCustomNotificationType) SetNotifyOnError(v bool)`

SetNotifyOnError sets NotifyOnError field to given value.

### HasNotifyOnError

`func (o *EmailCustomNotificationType) HasNotifyOnError() bool`

HasNotifyOnError returns a boolean if a field has been set.

### GetSuppressNotificationUntilLastRetry

`func (o *EmailCustomNotificationType) GetSuppressNotificationUntilLastRetry() bool`

GetSuppressNotificationUntilLastRetry returns the SuppressNotificationUntilLastRetry field if non-nil, zero value otherwise.

### GetSuppressNotificationUntilLastRetryOk

`func (o *EmailCustomNotificationType) GetSuppressNotificationUntilLastRetryOk() (*bool, bool)`

GetSuppressNotificationUntilLastRetryOk returns a tuple with the SuppressNotificationUntilLastRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressNotificationUntilLastRetry

`func (o *EmailCustomNotificationType) SetSuppressNotificationUntilLastRetry(v bool)`

SetSuppressNotificationUntilLastRetry sets SuppressNotificationUntilLastRetry field to given value.

### HasSuppressNotificationUntilLastRetry

`func (o *EmailCustomNotificationType) HasSuppressNotificationUntilLastRetry() bool`

HasSuppressNotificationUntilLastRetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


