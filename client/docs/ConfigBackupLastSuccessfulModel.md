# ConfigBackupLastSuccessfulModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSuccessfulTime** | Pointer to **time.Time** | Date and time when the last successful backup was created. | [optional] 
**SessionId** | Pointer to **string** | ID of the job session. | [optional] 

## Methods

### NewConfigBackupLastSuccessfulModel

`func NewConfigBackupLastSuccessfulModel() *ConfigBackupLastSuccessfulModel`

NewConfigBackupLastSuccessfulModel instantiates a new ConfigBackupLastSuccessfulModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigBackupLastSuccessfulModelWithDefaults

`func NewConfigBackupLastSuccessfulModelWithDefaults() *ConfigBackupLastSuccessfulModel`

NewConfigBackupLastSuccessfulModelWithDefaults instantiates a new ConfigBackupLastSuccessfulModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastSuccessfulTime

`func (o *ConfigBackupLastSuccessfulModel) GetLastSuccessfulTime() time.Time`

GetLastSuccessfulTime returns the LastSuccessfulTime field if non-nil, zero value otherwise.

### GetLastSuccessfulTimeOk

`func (o *ConfigBackupLastSuccessfulModel) GetLastSuccessfulTimeOk() (*time.Time, bool)`

GetLastSuccessfulTimeOk returns a tuple with the LastSuccessfulTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulTime

`func (o *ConfigBackupLastSuccessfulModel) SetLastSuccessfulTime(v time.Time)`

SetLastSuccessfulTime sets LastSuccessfulTime field to given value.

### HasLastSuccessfulTime

`func (o *ConfigBackupLastSuccessfulModel) HasLastSuccessfulTime() bool`

HasLastSuccessfulTime returns a boolean if a field has been set.

### GetSessionId

`func (o *ConfigBackupLastSuccessfulModel) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ConfigBackupLastSuccessfulModel) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ConfigBackupLastSuccessfulModel) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ConfigBackupLastSuccessfulModel) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


