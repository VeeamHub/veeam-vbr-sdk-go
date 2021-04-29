# JobStateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the job. | 
**Name** | **string** | Name of the job. | 
**Type** | [**EJobType**](EJobType.md) |  | 
**Description** | **string** | Description of the job. | 
**Status** | [**EJobStatus**](EJobStatus.md) |  | 
**LastRun** | Pointer to **time.Time** | Last run of the job. | [optional] 
**LastResult** | [**ESessionResult**](ESessionResult.md) |  | 
**NextRun** | Pointer to **time.Time** | Next run of the job. | [optional] 
**Workload** | [**EJobWorkload**](EJobWorkload.md) |  | 
**RepositoryId** | Pointer to **string** | ID of the backup repository. | [optional] 
**RepositoryName** | Pointer to **string** | Name of the backup repository. | [optional] 
**ObjectsCount** | **int32** | Number of objects processed by the job. | 
**SessionId** | Pointer to **string** | ID of the last job session. | [optional] 

## Methods

### NewJobStateModel

`func NewJobStateModel(id string, name string, type_ EJobType, description string, status EJobStatus, lastResult ESessionResult, workload EJobWorkload, objectsCount int32, ) *JobStateModel`

NewJobStateModel instantiates a new JobStateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobStateModelWithDefaults

`func NewJobStateModelWithDefaults() *JobStateModel`

NewJobStateModelWithDefaults instantiates a new JobStateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobStateModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobStateModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobStateModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *JobStateModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobStateModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobStateModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *JobStateModel) GetType() EJobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobStateModel) GetTypeOk() (*EJobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobStateModel) SetType(v EJobType)`

SetType sets Type field to given value.


### GetDescription

`func (o *JobStateModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JobStateModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JobStateModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *JobStateModel) GetStatus() EJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobStateModel) GetStatusOk() (*EJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobStateModel) SetStatus(v EJobStatus)`

SetStatus sets Status field to given value.


### GetLastRun

`func (o *JobStateModel) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *JobStateModel) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *JobStateModel) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *JobStateModel) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetLastResult

`func (o *JobStateModel) GetLastResult() ESessionResult`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *JobStateModel) GetLastResultOk() (*ESessionResult, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *JobStateModel) SetLastResult(v ESessionResult)`

SetLastResult sets LastResult field to given value.


### GetNextRun

`func (o *JobStateModel) GetNextRun() time.Time`

GetNextRun returns the NextRun field if non-nil, zero value otherwise.

### GetNextRunOk

`func (o *JobStateModel) GetNextRunOk() (*time.Time, bool)`

GetNextRunOk returns a tuple with the NextRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRun

`func (o *JobStateModel) SetNextRun(v time.Time)`

SetNextRun sets NextRun field to given value.

### HasNextRun

`func (o *JobStateModel) HasNextRun() bool`

HasNextRun returns a boolean if a field has been set.

### GetWorkload

`func (o *JobStateModel) GetWorkload() EJobWorkload`

GetWorkload returns the Workload field if non-nil, zero value otherwise.

### GetWorkloadOk

`func (o *JobStateModel) GetWorkloadOk() (*EJobWorkload, bool)`

GetWorkloadOk returns a tuple with the Workload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkload

`func (o *JobStateModel) SetWorkload(v EJobWorkload)`

SetWorkload sets Workload field to given value.


### GetRepositoryId

`func (o *JobStateModel) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *JobStateModel) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *JobStateModel) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *JobStateModel) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetRepositoryName

`func (o *JobStateModel) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *JobStateModel) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *JobStateModel) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.

### HasRepositoryName

`func (o *JobStateModel) HasRepositoryName() bool`

HasRepositoryName returns a boolean if a field has been set.

### GetObjectsCount

`func (o *JobStateModel) GetObjectsCount() int32`

GetObjectsCount returns the ObjectsCount field if non-nil, zero value otherwise.

### GetObjectsCountOk

`func (o *JobStateModel) GetObjectsCountOk() (*int32, bool)`

GetObjectsCountOk returns a tuple with the ObjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectsCount

`func (o *JobStateModel) SetObjectsCount(v int32)`

SetObjectsCount sets ObjectsCount field to given value.


### GetSessionId

`func (o *JobStateModel) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *JobStateModel) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *JobStateModel) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *JobStateModel) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


