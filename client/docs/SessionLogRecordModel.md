# SessionLogRecordModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the log record. | [optional] 
**Status** | Pointer to [**ETaskLogRecordStatus**](ETaskLogRecordStatus.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** | Date and time the operation was started. | [optional] 
**UpdateTime** | Pointer to **time.Time** | Date and time the log record was updated. | [optional] 
**Title** | Pointer to **string** | Title of the log record. | [optional] 
**Description** | Pointer to **string** | Description of the log record. | [optional] 

## Methods

### NewSessionLogRecordModel

`func NewSessionLogRecordModel() *SessionLogRecordModel`

NewSessionLogRecordModel instantiates a new SessionLogRecordModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionLogRecordModelWithDefaults

`func NewSessionLogRecordModelWithDefaults() *SessionLogRecordModel`

NewSessionLogRecordModelWithDefaults instantiates a new SessionLogRecordModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionLogRecordModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionLogRecordModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionLogRecordModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SessionLogRecordModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *SessionLogRecordModel) GetStatus() ETaskLogRecordStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SessionLogRecordModel) GetStatusOk() (*ETaskLogRecordStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SessionLogRecordModel) SetStatus(v ETaskLogRecordStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SessionLogRecordModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartTime

`func (o *SessionLogRecordModel) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SessionLogRecordModel) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SessionLogRecordModel) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SessionLogRecordModel) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *SessionLogRecordModel) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *SessionLogRecordModel) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *SessionLogRecordModel) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *SessionLogRecordModel) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetTitle

`func (o *SessionLogRecordModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SessionLogRecordModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SessionLogRecordModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SessionLogRecordModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *SessionLogRecordModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SessionLogRecordModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SessionLogRecordModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SessionLogRecordModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


