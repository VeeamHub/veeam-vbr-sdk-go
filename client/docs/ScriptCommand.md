# ScriptCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, script execution is enabled. | 
**Command** | Pointer to **string** | Path to the script. | [optional] 

## Methods

### NewScriptCommand

`func NewScriptCommand(isEnabled bool, ) *ScriptCommand`

NewScriptCommand instantiates a new ScriptCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptCommandWithDefaults

`func NewScriptCommandWithDefaults() *ScriptCommand`

NewScriptCommandWithDefaults instantiates a new ScriptCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ScriptCommand) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ScriptCommand) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ScriptCommand) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetCommand

`func (o *ScriptCommand) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ScriptCommand) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ScriptCommand) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *ScriptCommand) HasCommand() bool`

HasCommand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


