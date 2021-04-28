# GeneralOptionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailSettings** | Pointer to [**GeneralOptionsEmailNotificationsModel**](GeneralOptionsEmailNotificationsModel.md) |  | [optional] 
**Notifications** | Pointer to [**GeneralOptionsNotificationsModel**](GeneralOptionsNotificationsModel.md) |  | [optional] 

## Methods

### NewGeneralOptionsModel

`func NewGeneralOptionsModel() *GeneralOptionsModel`

NewGeneralOptionsModel instantiates a new GeneralOptionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralOptionsModelWithDefaults

`func NewGeneralOptionsModelWithDefaults() *GeneralOptionsModel`

NewGeneralOptionsModelWithDefaults instantiates a new GeneralOptionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailSettings

`func (o *GeneralOptionsModel) GetEmailSettings() GeneralOptionsEmailNotificationsModel`

GetEmailSettings returns the EmailSettings field if non-nil, zero value otherwise.

### GetEmailSettingsOk

`func (o *GeneralOptionsModel) GetEmailSettingsOk() (*GeneralOptionsEmailNotificationsModel, bool)`

GetEmailSettingsOk returns a tuple with the EmailSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSettings

`func (o *GeneralOptionsModel) SetEmailSettings(v GeneralOptionsEmailNotificationsModel)`

SetEmailSettings sets EmailSettings field to given value.

### HasEmailSettings

`func (o *GeneralOptionsModel) HasEmailSettings() bool`

HasEmailSettings returns a boolean if a field has been set.

### GetNotifications

`func (o *GeneralOptionsModel) GetNotifications() GeneralOptionsNotificationsModel`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *GeneralOptionsModel) GetNotificationsOk() (*GeneralOptionsNotificationsModel, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *GeneralOptionsModel) SetNotifications(v GeneralOptionsNotificationsModel)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *GeneralOptionsModel) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


