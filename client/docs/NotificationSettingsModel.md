# NotificationSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendSNMPNotifications** | Pointer to **bool** | If *true*, SNMP notifications are enabled for this job. | [optional] 
**EmailNotifications** | Pointer to [**EmailNotificationSettingsModel**](EmailNotificationSettingsModel.md) |  | [optional] 
**VmAttribute** | Pointer to [**NotificationVmAttributeSettingsModel**](NotificationVmAttributeSettingsModel.md) |  | [optional] 

## Methods

### NewNotificationSettingsModel

`func NewNotificationSettingsModel() *NotificationSettingsModel`

NewNotificationSettingsModel instantiates a new NotificationSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSettingsModelWithDefaults

`func NewNotificationSettingsModelWithDefaults() *NotificationSettingsModel`

NewNotificationSettingsModelWithDefaults instantiates a new NotificationSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendSNMPNotifications

`func (o *NotificationSettingsModel) GetSendSNMPNotifications() bool`

GetSendSNMPNotifications returns the SendSNMPNotifications field if non-nil, zero value otherwise.

### GetSendSNMPNotificationsOk

`func (o *NotificationSettingsModel) GetSendSNMPNotificationsOk() (*bool, bool)`

GetSendSNMPNotificationsOk returns a tuple with the SendSNMPNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendSNMPNotifications

`func (o *NotificationSettingsModel) SetSendSNMPNotifications(v bool)`

SetSendSNMPNotifications sets SendSNMPNotifications field to given value.

### HasSendSNMPNotifications

`func (o *NotificationSettingsModel) HasSendSNMPNotifications() bool`

HasSendSNMPNotifications returns a boolean if a field has been set.

### GetEmailNotifications

`func (o *NotificationSettingsModel) GetEmailNotifications() EmailNotificationSettingsModel`

GetEmailNotifications returns the EmailNotifications field if non-nil, zero value otherwise.

### GetEmailNotificationsOk

`func (o *NotificationSettingsModel) GetEmailNotificationsOk() (*EmailNotificationSettingsModel, bool)`

GetEmailNotificationsOk returns a tuple with the EmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotifications

`func (o *NotificationSettingsModel) SetEmailNotifications(v EmailNotificationSettingsModel)`

SetEmailNotifications sets EmailNotifications field to given value.

### HasEmailNotifications

`func (o *NotificationSettingsModel) HasEmailNotifications() bool`

HasEmailNotifications returns a boolean if a field has been set.

### GetVmAttribute

`func (o *NotificationSettingsModel) GetVmAttribute() NotificationVmAttributeSettingsModel`

GetVmAttribute returns the VmAttribute field if non-nil, zero value otherwise.

### GetVmAttributeOk

`func (o *NotificationSettingsModel) GetVmAttributeOk() (*NotificationVmAttributeSettingsModel, bool)`

GetVmAttributeOk returns a tuple with the VmAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmAttribute

`func (o *NotificationSettingsModel) SetVmAttribute(v NotificationVmAttributeSettingsModel)`

SetVmAttribute sets VmAttribute field to given value.

### HasVmAttribute

`func (o *NotificationSettingsModel) HasVmAttribute() bool`

HasVmAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


