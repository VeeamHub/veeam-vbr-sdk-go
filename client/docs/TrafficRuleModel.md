# TrafficRuleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the rule. | [optional] 
**Name** | **string** | Name of the rule. | 
**SourceIPStart** | **string** | Start IP address of the range for the backup infrastructure components on the source side. | 
**SourceIPEnd** | **string** | End IP address of the range for the backup infrastructure components on the source side. | 
**TargetIPStart** | **string** | Start IP address of the range for the backup infrastructure components on the target side. | 
**TargetIPEnd** | **string** | End IP address of the range for the backup infrastructure components on the target side. | 
**EncryptionEnabled** | Pointer to **bool** | If *true*, traffic encryption is enabled. | [optional] 
**ThrottlingEnabled** | Pointer to **bool** | If *true*, traffic throttling is enabled. | [optional] 
**ThrottlingUnit** | Pointer to [**ESpeedUnit**](ESpeedUnit.md) |  | [optional] 
**ThrottlingValue** | Pointer to **int32** | Maximum speed that must be used to transfer data from source to target. | [optional] 
**ThrottlingWindowEnabled** | Pointer to **bool** | If *true*, throttling window during which the speed must be limited is enabled. | [optional] 
**ThrottlingWindowOptions** | Pointer to [**BackupWindowSettingModel**](BackupWindowSettingModel.md) |  | [optional] 

## Methods

### NewTrafficRuleModel

`func NewTrafficRuleModel(name string, sourceIPStart string, sourceIPEnd string, targetIPStart string, targetIPEnd string, ) *TrafficRuleModel`

NewTrafficRuleModel instantiates a new TrafficRuleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficRuleModelWithDefaults

`func NewTrafficRuleModelWithDefaults() *TrafficRuleModel`

NewTrafficRuleModelWithDefaults instantiates a new TrafficRuleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrafficRuleModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrafficRuleModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrafficRuleModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrafficRuleModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TrafficRuleModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrafficRuleModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrafficRuleModel) SetName(v string)`

SetName sets Name field to given value.


### GetSourceIPStart

`func (o *TrafficRuleModel) GetSourceIPStart() string`

GetSourceIPStart returns the SourceIPStart field if non-nil, zero value otherwise.

### GetSourceIPStartOk

`func (o *TrafficRuleModel) GetSourceIPStartOk() (*string, bool)`

GetSourceIPStartOk returns a tuple with the SourceIPStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIPStart

`func (o *TrafficRuleModel) SetSourceIPStart(v string)`

SetSourceIPStart sets SourceIPStart field to given value.


### GetSourceIPEnd

`func (o *TrafficRuleModel) GetSourceIPEnd() string`

GetSourceIPEnd returns the SourceIPEnd field if non-nil, zero value otherwise.

### GetSourceIPEndOk

`func (o *TrafficRuleModel) GetSourceIPEndOk() (*string, bool)`

GetSourceIPEndOk returns a tuple with the SourceIPEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIPEnd

`func (o *TrafficRuleModel) SetSourceIPEnd(v string)`

SetSourceIPEnd sets SourceIPEnd field to given value.


### GetTargetIPStart

`func (o *TrafficRuleModel) GetTargetIPStart() string`

GetTargetIPStart returns the TargetIPStart field if non-nil, zero value otherwise.

### GetTargetIPStartOk

`func (o *TrafficRuleModel) GetTargetIPStartOk() (*string, bool)`

GetTargetIPStartOk returns a tuple with the TargetIPStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIPStart

`func (o *TrafficRuleModel) SetTargetIPStart(v string)`

SetTargetIPStart sets TargetIPStart field to given value.


### GetTargetIPEnd

`func (o *TrafficRuleModel) GetTargetIPEnd() string`

GetTargetIPEnd returns the TargetIPEnd field if non-nil, zero value otherwise.

### GetTargetIPEndOk

`func (o *TrafficRuleModel) GetTargetIPEndOk() (*string, bool)`

GetTargetIPEndOk returns a tuple with the TargetIPEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIPEnd

`func (o *TrafficRuleModel) SetTargetIPEnd(v string)`

SetTargetIPEnd sets TargetIPEnd field to given value.


### GetEncryptionEnabled

`func (o *TrafficRuleModel) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *TrafficRuleModel) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *TrafficRuleModel) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *TrafficRuleModel) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### GetThrottlingEnabled

`func (o *TrafficRuleModel) GetThrottlingEnabled() bool`

GetThrottlingEnabled returns the ThrottlingEnabled field if non-nil, zero value otherwise.

### GetThrottlingEnabledOk

`func (o *TrafficRuleModel) GetThrottlingEnabledOk() (*bool, bool)`

GetThrottlingEnabledOk returns a tuple with the ThrottlingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingEnabled

`func (o *TrafficRuleModel) SetThrottlingEnabled(v bool)`

SetThrottlingEnabled sets ThrottlingEnabled field to given value.

### HasThrottlingEnabled

`func (o *TrafficRuleModel) HasThrottlingEnabled() bool`

HasThrottlingEnabled returns a boolean if a field has been set.

### GetThrottlingUnit

`func (o *TrafficRuleModel) GetThrottlingUnit() ESpeedUnit`

GetThrottlingUnit returns the ThrottlingUnit field if non-nil, zero value otherwise.

### GetThrottlingUnitOk

`func (o *TrafficRuleModel) GetThrottlingUnitOk() (*ESpeedUnit, bool)`

GetThrottlingUnitOk returns a tuple with the ThrottlingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingUnit

`func (o *TrafficRuleModel) SetThrottlingUnit(v ESpeedUnit)`

SetThrottlingUnit sets ThrottlingUnit field to given value.

### HasThrottlingUnit

`func (o *TrafficRuleModel) HasThrottlingUnit() bool`

HasThrottlingUnit returns a boolean if a field has been set.

### GetThrottlingValue

`func (o *TrafficRuleModel) GetThrottlingValue() int32`

GetThrottlingValue returns the ThrottlingValue field if non-nil, zero value otherwise.

### GetThrottlingValueOk

`func (o *TrafficRuleModel) GetThrottlingValueOk() (*int32, bool)`

GetThrottlingValueOk returns a tuple with the ThrottlingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingValue

`func (o *TrafficRuleModel) SetThrottlingValue(v int32)`

SetThrottlingValue sets ThrottlingValue field to given value.

### HasThrottlingValue

`func (o *TrafficRuleModel) HasThrottlingValue() bool`

HasThrottlingValue returns a boolean if a field has been set.

### GetThrottlingWindowEnabled

`func (o *TrafficRuleModel) GetThrottlingWindowEnabled() bool`

GetThrottlingWindowEnabled returns the ThrottlingWindowEnabled field if non-nil, zero value otherwise.

### GetThrottlingWindowEnabledOk

`func (o *TrafficRuleModel) GetThrottlingWindowEnabledOk() (*bool, bool)`

GetThrottlingWindowEnabledOk returns a tuple with the ThrottlingWindowEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingWindowEnabled

`func (o *TrafficRuleModel) SetThrottlingWindowEnabled(v bool)`

SetThrottlingWindowEnabled sets ThrottlingWindowEnabled field to given value.

### HasThrottlingWindowEnabled

`func (o *TrafficRuleModel) HasThrottlingWindowEnabled() bool`

HasThrottlingWindowEnabled returns a boolean if a field has been set.

### GetThrottlingWindowOptions

`func (o *TrafficRuleModel) GetThrottlingWindowOptions() BackupWindowSettingModel`

GetThrottlingWindowOptions returns the ThrottlingWindowOptions field if non-nil, zero value otherwise.

### GetThrottlingWindowOptionsOk

`func (o *TrafficRuleModel) GetThrottlingWindowOptionsOk() (*BackupWindowSettingModel, bool)`

GetThrottlingWindowOptionsOk returns a tuple with the ThrottlingWindowOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingWindowOptions

`func (o *TrafficRuleModel) SetThrottlingWindowOptions(v BackupWindowSettingModel)`

SetThrottlingWindowOptions sets ThrottlingWindowOptions field to given value.

### HasThrottlingWindowOptions

`func (o *TrafficRuleModel) HasThrottlingWindowOptions() bool`

HasThrottlingWindowOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


