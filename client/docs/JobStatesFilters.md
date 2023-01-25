# JobStatesFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** | Skips the specified number of jobs. | [optional] 
**Limit** | Pointer to **int32** | Returns the specified number of jobs. | [optional] 
**OrderColumn** | Pointer to [**EJobStatesFiltersOrderColumn**](EJobStatesFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** | Sorts jobs in the ascending order by the &#x60;orderColumn&#x60; parameter. | [optional] 
**IdFilter** | Pointer to **string** |  | [optional] 
**NameFilter** | Pointer to **string** | Filters jobs by the &#x60;nameFilter&#x60; pattern. The pattern can match any job state parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both. | [optional] 
**TypeFilter** | Pointer to [**EJobType**](EJobType.md) |  | [optional] 
**LastResultFilter** | Pointer to [**ESessionResult**](ESessionResult.md) |  | [optional] 
**StatusFilter** | Pointer to [**EJobStatus**](EJobStatus.md) |  | [optional] 
**WorkloadFilter** | Pointer to [**EJobWorkload**](EJobWorkload.md) |  | [optional] 
**LastRunAfterFilter** | Pointer to **time.Time** |  | [optional] 
**LastRunBeforeFilter** | Pointer to **time.Time** |  | [optional] 
**IsHighPriorityJobFilter** | Pointer to **bool** |  | [optional] 
**RepositoryIdFilter** | Pointer to **string** |  | [optional] 
**ObjectsCountFilter** | Pointer to **int32** |  | [optional] 

## Methods

### NewJobStatesFilters

`func NewJobStatesFilters() *JobStatesFilters`

NewJobStatesFilters instantiates a new JobStatesFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobStatesFiltersWithDefaults

`func NewJobStatesFiltersWithDefaults() *JobStatesFilters`

NewJobStatesFiltersWithDefaults instantiates a new JobStatesFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *JobStatesFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *JobStatesFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *JobStatesFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *JobStatesFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *JobStatesFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *JobStatesFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *JobStatesFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *JobStatesFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *JobStatesFilters) GetOrderColumn() EJobStatesFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *JobStatesFilters) GetOrderColumnOk() (*EJobStatesFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *JobStatesFilters) SetOrderColumn(v EJobStatesFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *JobStatesFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *JobStatesFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *JobStatesFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *JobStatesFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *JobStatesFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.

### GetIdFilter

`func (o *JobStatesFilters) GetIdFilter() string`

GetIdFilter returns the IdFilter field if non-nil, zero value otherwise.

### GetIdFilterOk

`func (o *JobStatesFilters) GetIdFilterOk() (*string, bool)`

GetIdFilterOk returns a tuple with the IdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdFilter

`func (o *JobStatesFilters) SetIdFilter(v string)`

SetIdFilter sets IdFilter field to given value.

### HasIdFilter

`func (o *JobStatesFilters) HasIdFilter() bool`

HasIdFilter returns a boolean if a field has been set.

### GetNameFilter

`func (o *JobStatesFilters) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *JobStatesFilters) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *JobStatesFilters) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *JobStatesFilters) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetTypeFilter

`func (o *JobStatesFilters) GetTypeFilter() EJobType`

GetTypeFilter returns the TypeFilter field if non-nil, zero value otherwise.

### GetTypeFilterOk

`func (o *JobStatesFilters) GetTypeFilterOk() (*EJobType, bool)`

GetTypeFilterOk returns a tuple with the TypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFilter

`func (o *JobStatesFilters) SetTypeFilter(v EJobType)`

SetTypeFilter sets TypeFilter field to given value.

### HasTypeFilter

`func (o *JobStatesFilters) HasTypeFilter() bool`

HasTypeFilter returns a boolean if a field has been set.

### GetLastResultFilter

`func (o *JobStatesFilters) GetLastResultFilter() ESessionResult`

GetLastResultFilter returns the LastResultFilter field if non-nil, zero value otherwise.

### GetLastResultFilterOk

`func (o *JobStatesFilters) GetLastResultFilterOk() (*ESessionResult, bool)`

GetLastResultFilterOk returns a tuple with the LastResultFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResultFilter

`func (o *JobStatesFilters) SetLastResultFilter(v ESessionResult)`

SetLastResultFilter sets LastResultFilter field to given value.

### HasLastResultFilter

`func (o *JobStatesFilters) HasLastResultFilter() bool`

HasLastResultFilter returns a boolean if a field has been set.

### GetStatusFilter

`func (o *JobStatesFilters) GetStatusFilter() EJobStatus`

GetStatusFilter returns the StatusFilter field if non-nil, zero value otherwise.

### GetStatusFilterOk

`func (o *JobStatesFilters) GetStatusFilterOk() (*EJobStatus, bool)`

GetStatusFilterOk returns a tuple with the StatusFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusFilter

`func (o *JobStatesFilters) SetStatusFilter(v EJobStatus)`

SetStatusFilter sets StatusFilter field to given value.

### HasStatusFilter

`func (o *JobStatesFilters) HasStatusFilter() bool`

HasStatusFilter returns a boolean if a field has been set.

### GetWorkloadFilter

`func (o *JobStatesFilters) GetWorkloadFilter() EJobWorkload`

GetWorkloadFilter returns the WorkloadFilter field if non-nil, zero value otherwise.

### GetWorkloadFilterOk

`func (o *JobStatesFilters) GetWorkloadFilterOk() (*EJobWorkload, bool)`

GetWorkloadFilterOk returns a tuple with the WorkloadFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadFilter

`func (o *JobStatesFilters) SetWorkloadFilter(v EJobWorkload)`

SetWorkloadFilter sets WorkloadFilter field to given value.

### HasWorkloadFilter

`func (o *JobStatesFilters) HasWorkloadFilter() bool`

HasWorkloadFilter returns a boolean if a field has been set.

### GetLastRunAfterFilter

`func (o *JobStatesFilters) GetLastRunAfterFilter() time.Time`

GetLastRunAfterFilter returns the LastRunAfterFilter field if non-nil, zero value otherwise.

### GetLastRunAfterFilterOk

`func (o *JobStatesFilters) GetLastRunAfterFilterOk() (*time.Time, bool)`

GetLastRunAfterFilterOk returns a tuple with the LastRunAfterFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunAfterFilter

`func (o *JobStatesFilters) SetLastRunAfterFilter(v time.Time)`

SetLastRunAfterFilter sets LastRunAfterFilter field to given value.

### HasLastRunAfterFilter

`func (o *JobStatesFilters) HasLastRunAfterFilter() bool`

HasLastRunAfterFilter returns a boolean if a field has been set.

### GetLastRunBeforeFilter

`func (o *JobStatesFilters) GetLastRunBeforeFilter() time.Time`

GetLastRunBeforeFilter returns the LastRunBeforeFilter field if non-nil, zero value otherwise.

### GetLastRunBeforeFilterOk

`func (o *JobStatesFilters) GetLastRunBeforeFilterOk() (*time.Time, bool)`

GetLastRunBeforeFilterOk returns a tuple with the LastRunBeforeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunBeforeFilter

`func (o *JobStatesFilters) SetLastRunBeforeFilter(v time.Time)`

SetLastRunBeforeFilter sets LastRunBeforeFilter field to given value.

### HasLastRunBeforeFilter

`func (o *JobStatesFilters) HasLastRunBeforeFilter() bool`

HasLastRunBeforeFilter returns a boolean if a field has been set.

### GetIsHighPriorityJobFilter

`func (o *JobStatesFilters) GetIsHighPriorityJobFilter() bool`

GetIsHighPriorityJobFilter returns the IsHighPriorityJobFilter field if non-nil, zero value otherwise.

### GetIsHighPriorityJobFilterOk

`func (o *JobStatesFilters) GetIsHighPriorityJobFilterOk() (*bool, bool)`

GetIsHighPriorityJobFilterOk returns a tuple with the IsHighPriorityJobFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHighPriorityJobFilter

`func (o *JobStatesFilters) SetIsHighPriorityJobFilter(v bool)`

SetIsHighPriorityJobFilter sets IsHighPriorityJobFilter field to given value.

### HasIsHighPriorityJobFilter

`func (o *JobStatesFilters) HasIsHighPriorityJobFilter() bool`

HasIsHighPriorityJobFilter returns a boolean if a field has been set.

### GetRepositoryIdFilter

`func (o *JobStatesFilters) GetRepositoryIdFilter() string`

GetRepositoryIdFilter returns the RepositoryIdFilter field if non-nil, zero value otherwise.

### GetRepositoryIdFilterOk

`func (o *JobStatesFilters) GetRepositoryIdFilterOk() (*string, bool)`

GetRepositoryIdFilterOk returns a tuple with the RepositoryIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryIdFilter

`func (o *JobStatesFilters) SetRepositoryIdFilter(v string)`

SetRepositoryIdFilter sets RepositoryIdFilter field to given value.

### HasRepositoryIdFilter

`func (o *JobStatesFilters) HasRepositoryIdFilter() bool`

HasRepositoryIdFilter returns a boolean if a field has been set.

### GetObjectsCountFilter

`func (o *JobStatesFilters) GetObjectsCountFilter() int32`

GetObjectsCountFilter returns the ObjectsCountFilter field if non-nil, zero value otherwise.

### GetObjectsCountFilterOk

`func (o *JobStatesFilters) GetObjectsCountFilterOk() (*int32, bool)`

GetObjectsCountFilterOk returns a tuple with the ObjectsCountFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectsCountFilter

`func (o *JobStatesFilters) SetObjectsCountFilter(v int32)`

SetObjectsCountFilter sets ObjectsCountFilter field to given value.

### HasObjectsCountFilter

`func (o *JobStatesFilters) HasObjectsCountFilter() bool`

HasObjectsCountFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


