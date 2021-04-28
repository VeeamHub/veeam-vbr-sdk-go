# EmailNotificationSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, email notifications are enabled for this job. | 
**Recipients** | Pointer to **[]string** | Array of recipient’s email addresses. | [optional] 
**NotificationType** | Pointer to [**EEmailNotificationType**](EEmailNotificationType.md) |  | [optional] 
**CustomNotificationSettings** | Pointer to [**EmailCustomNotificationType**](EmailCustomNotificationType.md) |  | [optional] 

## Methods

### NewEmailNotificationSettingsModel

`func NewEmailNotificationSettingsModel(isEnabled bool, ) *EmailNotificationSettingsModel`

NewEmailNotificationSettingsModel instantiates a new EmailNotificationSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailNotificationSettingsModelWithDefaults

`func NewEmailNotificationSettingsModelWithDefaults() *EmailNotificationSettingsModel`

NewEmailNotificationSettingsModelWithDefaults instantiates a new EmailNotificationSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *EmailNotificationSettingsModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *EmailNotificationSettingsModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *EmailNotificationSettingsModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetRecipients

`func (o *EmailNotificationSettingsModel) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *EmailNotificationSettingsModel) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *EmailNotificationSettingsModel) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *EmailNotificationSettingsModel) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetNotificationType

`func (o *EmailNotificationSettingsModel) GetNotificationType() EEmailNotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *EmailNotificationSettingsModel) GetNotificationTypeOk() (*EEmailNotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *EmailNotificationSettingsModel) SetNotificationType(v EEmailNotificationType)`

SetNotificationType sets NotificationType field to given value.

### HasNotificationType

`func (o *EmailNotificationSettingsModel) HasNotificationType() bool`

HasNotificationType returns a boolean if a field has been set.

### GetCustomNotificationSettings

`func (o *EmailNotificationSettingsModel) GetCustomNotificationSettings() EmailCustomNotificationType`

GetCustomNotificationSettings returns the CustomNotificationSettings field if non-nil, zero value otherwise.

### GetCustomNotificationSettingsOk

`func (o *EmailNotificationSettingsModel) GetCustomNotificationSettingsOk() (*EmailCustomNotificationType, bool)`

GetCustomNotificationSettingsOk returns a tuple with the CustomNotificationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNotificationSettings

`func (o *EmailNotificationSettingsModel) SetCustomNotificationSettings(v EmailCustomNotificationType)`

SetCustomNotificationSettings sets CustomNotificationSettings field to given value.

### HasCustomNotificationSettings

`func (o *EmailNotificationSettingsModel) HasCustomNotificationSettings() bool`

HasCustomNotificationSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


