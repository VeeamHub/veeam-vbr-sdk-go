# JobScriptsSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreCommand** | Pointer to [**ScriptCommand**](ScriptCommand.md) |  | [optional] 
**PostCommand** | Pointer to [**ScriptCommand**](ScriptCommand.md) |  | [optional] 
**PeriodicityType** | Pointer to [**EScriptPeriodicityType**](EScriptPeriodicityType.md) |  | [optional] 
**RunScriptEvery** | Pointer to **int32** | Number of the backup job session after which the scripts must be executed. | [optional] 
**DayOfWeek** | Pointer to [**[]EDayOfWeek**](EDayOfWeek.md) | Days of the week when the scripts must be executed. | [optional] 

## Methods

### NewJobScriptsSettingsModel

`func NewJobScriptsSettingsModel() *JobScriptsSettingsModel`

NewJobScriptsSettingsModel instantiates a new JobScriptsSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobScriptsSettingsModelWithDefaults

`func NewJobScriptsSettingsModelWithDefaults() *JobScriptsSettingsModel`

NewJobScriptsSettingsModelWithDefaults instantiates a new JobScriptsSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreCommand

`func (o *JobScriptsSettingsModel) GetPreCommand() ScriptCommand`

GetPreCommand returns the PreCommand field if non-nil, zero value otherwise.

### GetPreCommandOk

`func (o *JobScriptsSettingsModel) GetPreCommandOk() (*ScriptCommand, bool)`

GetPreCommandOk returns a tuple with the PreCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreCommand

`func (o *JobScriptsSettingsModel) SetPreCommand(v ScriptCommand)`

SetPreCommand sets PreCommand field to given value.

### HasPreCommand

`func (o *JobScriptsSettingsModel) HasPreCommand() bool`

HasPreCommand returns a boolean if a field has been set.

### GetPostCommand

`func (o *JobScriptsSettingsModel) GetPostCommand() ScriptCommand`

GetPostCommand returns the PostCommand field if non-nil, zero value otherwise.

### GetPostCommandOk

`func (o *JobScriptsSettingsModel) GetPostCommandOk() (*ScriptCommand, bool)`

GetPostCommandOk returns a tuple with the PostCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCommand

`func (o *JobScriptsSettingsModel) SetPostCommand(v ScriptCommand)`

SetPostCommand sets PostCommand field to given value.

### HasPostCommand

`func (o *JobScriptsSettingsModel) HasPostCommand() bool`

HasPostCommand returns a boolean if a field has been set.

### GetPeriodicityType

`func (o *JobScriptsSettingsModel) GetPeriodicityType() EScriptPeriodicityType`

GetPeriodicityType returns the PeriodicityType field if non-nil, zero value otherwise.

### GetPeriodicityTypeOk

`func (o *JobScriptsSettingsModel) GetPeriodicityTypeOk() (*EScriptPeriodicityType, bool)`

GetPeriodicityTypeOk returns a tuple with the PeriodicityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicityType

`func (o *JobScriptsSettingsModel) SetPeriodicityType(v EScriptPeriodicityType)`

SetPeriodicityType sets PeriodicityType field to given value.

### HasPeriodicityType

`func (o *JobScriptsSettingsModel) HasPeriodicityType() bool`

HasPeriodicityType returns a boolean if a field has been set.

### GetRunScriptEvery

`func (o *JobScriptsSettingsModel) GetRunScriptEvery() int32`

GetRunScriptEvery returns the RunScriptEvery field if non-nil, zero value otherwise.

### GetRunScriptEveryOk

`func (o *JobScriptsSettingsModel) GetRunScriptEveryOk() (*int32, bool)`

GetRunScriptEveryOk returns a tuple with the RunScriptEvery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunScriptEvery

`func (o *JobScriptsSettingsModel) SetRunScriptEvery(v int32)`

SetRunScriptEvery sets RunScriptEvery field to given value.

### HasRunScriptEvery

`func (o *JobScriptsSettingsModel) HasRunScriptEvery() bool`

HasRunScriptEvery returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *JobScriptsSettingsModel) GetDayOfWeek() []EDayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *JobScriptsSettingsModel) GetDayOfWeekOk() (*[]EDayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *JobScriptsSettingsModel) SetDayOfWeek(v []EDayOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *JobScriptsSettingsModel) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


