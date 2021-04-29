# SessionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the session. | 
**Name** | **string** | Name of the session. | 
**ActivityId** | **string** | ID of the activity. | 
**SessionType** | [**ESessionType**](ESessionType.md) |  | 
**CreationTime** | **time.Time** | Date and time the session was created. | 
**EndTime** | Pointer to **time.Time** | Date and time the session was ended. | [optional] 
**State** | [**ESessionState**](ESessionState.md) |  | 
**ProgressPercent** | Pointer to **int32** | Progress percentage of the session. | [optional] 
**Result** | Pointer to [**SessionResultModel**](SessionResultModel.md) |  | [optional] 
**ResourceId** | Pointer to **string** | ID of the resource. | [optional] 
**ResourceReference** | Pointer to **string** | URI of the resource. | [optional] 
**ParentSessionId** | Pointer to **string** | ID of the parent session. | [optional] 
**Usn** | **int64** | Update sequence number. | 

## Methods

### NewSessionModel

`func NewSessionModel(id string, name string, activityId string, sessionType ESessionType, creationTime time.Time, state ESessionState, usn int64, ) *SessionModel`

NewSessionModel instantiates a new SessionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionModelWithDefaults

`func NewSessionModelWithDefaults() *SessionModel`

NewSessionModelWithDefaults instantiates a new SessionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SessionModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SessionModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SessionModel) SetName(v string)`

SetName sets Name field to given value.


### GetActivityId

`func (o *SessionModel) GetActivityId() string`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *SessionModel) GetActivityIdOk() (*string, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *SessionModel) SetActivityId(v string)`

SetActivityId sets ActivityId field to given value.


### GetSessionType

`func (o *SessionModel) GetSessionType() ESessionType`

GetSessionType returns the SessionType field if non-nil, zero value otherwise.

### GetSessionTypeOk

`func (o *SessionModel) GetSessionTypeOk() (*ESessionType, bool)`

GetSessionTypeOk returns a tuple with the SessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionType

`func (o *SessionModel) SetSessionType(v ESessionType)`

SetSessionType sets SessionType field to given value.


### GetCreationTime

`func (o *SessionModel) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *SessionModel) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *SessionModel) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetEndTime

`func (o *SessionModel) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SessionModel) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SessionModel) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *SessionModel) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetState

`func (o *SessionModel) GetState() ESessionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SessionModel) GetStateOk() (*ESessionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SessionModel) SetState(v ESessionState)`

SetState sets State field to given value.


### GetProgressPercent

`func (o *SessionModel) GetProgressPercent() int32`

GetProgressPercent returns the ProgressPercent field if non-nil, zero value otherwise.

### GetProgressPercentOk

`func (o *SessionModel) GetProgressPercentOk() (*int32, bool)`

GetProgressPercentOk returns a tuple with the ProgressPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressPercent

`func (o *SessionModel) SetProgressPercent(v int32)`

SetProgressPercent sets ProgressPercent field to given value.

### HasProgressPercent

`func (o *SessionModel) HasProgressPercent() bool`

HasProgressPercent returns a boolean if a field has been set.

### GetResult

`func (o *SessionModel) GetResult() SessionResultModel`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SessionModel) GetResultOk() (*SessionResultModel, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SessionModel) SetResult(v SessionResultModel)`

SetResult sets Result field to given value.

### HasResult

`func (o *SessionModel) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetResourceId

`func (o *SessionModel) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *SessionModel) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *SessionModel) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *SessionModel) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceReference

`func (o *SessionModel) GetResourceReference() string`

GetResourceReference returns the ResourceReference field if non-nil, zero value otherwise.

### GetResourceReferenceOk

`func (o *SessionModel) GetResourceReferenceOk() (*string, bool)`

GetResourceReferenceOk returns a tuple with the ResourceReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceReference

`func (o *SessionModel) SetResourceReference(v string)`

SetResourceReference sets ResourceReference field to given value.

### HasResourceReference

`func (o *SessionModel) HasResourceReference() bool`

HasResourceReference returns a boolean if a field has been set.

### GetParentSessionId

`func (o *SessionModel) GetParentSessionId() string`

GetParentSessionId returns the ParentSessionId field if non-nil, zero value otherwise.

### GetParentSessionIdOk

`func (o *SessionModel) GetParentSessionIdOk() (*string, bool)`

GetParentSessionIdOk returns a tuple with the ParentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSessionId

`func (o *SessionModel) SetParentSessionId(v string)`

SetParentSessionId sets ParentSessionId field to given value.

### HasParentSessionId

`func (o *SessionModel) HasParentSessionId() bool`

HasParentSessionId returns a boolean if a field has been set.

### GetUsn

`func (o *SessionModel) GetUsn() int64`

GetUsn returns the Usn field if non-nil, zero value otherwise.

### GetUsnOk

`func (o *SessionModel) GetUsnOk() (*int64, bool)`

GetUsnOk returns a tuple with the Usn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsn

`func (o *SessionModel) SetUsn(v int64)`

SetUsn sets Usn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


