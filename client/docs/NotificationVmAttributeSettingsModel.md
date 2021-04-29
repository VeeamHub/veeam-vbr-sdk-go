# NotificationVmAttributeSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, information about successfully performed backup is written to a VM attribute. | 
**Notes** | Pointer to **string** | Name of the VM attribute. | [optional] 
**AppendToExisitingValue** | Pointer to **bool** | If *true*, information about successfully performed backup is appended to the existing value of the attribute added by the user. | [optional] 

## Methods

### NewNotificationVmAttributeSettingsModel

`func NewNotificationVmAttributeSettingsModel(isEnabled bool, ) *NotificationVmAttributeSettingsModel`

NewNotificationVmAttributeSettingsModel instantiates a new NotificationVmAttributeSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationVmAttributeSettingsModelWithDefaults

`func NewNotificationVmAttributeSettingsModelWithDefaults() *NotificationVmAttributeSettingsModel`

NewNotificationVmAttributeSettingsModelWithDefaults instantiates a new NotificationVmAttributeSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *NotificationVmAttributeSettingsModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *NotificationVmAttributeSettingsModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *NotificationVmAttributeSettingsModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetNotes

`func (o *NotificationVmAttributeSettingsModel) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *NotificationVmAttributeSettingsModel) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *NotificationVmAttributeSettingsModel) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *NotificationVmAttributeSettingsModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAppendToExisitingValue

`func (o *NotificationVmAttributeSettingsModel) GetAppendToExisitingValue() bool`

GetAppendToExisitingValue returns the AppendToExisitingValue field if non-nil, zero value otherwise.

### GetAppendToExisitingValueOk

`func (o *NotificationVmAttributeSettingsModel) GetAppendToExisitingValueOk() (*bool, bool)`

GetAppendToExisitingValueOk returns a tuple with the AppendToExisitingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppendToExisitingValue

`func (o *NotificationVmAttributeSettingsModel) SetAppendToExisitingValue(v bool)`

SetAppendToExisitingValue sets AppendToExisitingValue field to given value.

### HasAppendToExisitingValue

`func (o *NotificationVmAttributeSettingsModel) HasAppendToExisitingValue() bool`

HasAppendToExisitingValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


