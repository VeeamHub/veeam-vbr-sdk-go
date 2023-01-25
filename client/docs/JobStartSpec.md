# JobStartSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerformActiveFull** | **bool** | If *true*, Veeam Backup &amp; Replication will perform an active full backup. | [default to false]
**StartChainedJobs** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication will start chained jobs as well. | [optional] [default to false]

## Methods

### NewJobStartSpec

`func NewJobStartSpec(performActiveFull bool, ) *JobStartSpec`

NewJobStartSpec instantiates a new JobStartSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobStartSpecWithDefaults

`func NewJobStartSpecWithDefaults() *JobStartSpec`

NewJobStartSpecWithDefaults instantiates a new JobStartSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerformActiveFull

`func (o *JobStartSpec) GetPerformActiveFull() bool`

GetPerformActiveFull returns the PerformActiveFull field if non-nil, zero value otherwise.

### GetPerformActiveFullOk

`func (o *JobStartSpec) GetPerformActiveFullOk() (*bool, bool)`

GetPerformActiveFullOk returns a tuple with the PerformActiveFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformActiveFull

`func (o *JobStartSpec) SetPerformActiveFull(v bool)`

SetPerformActiveFull sets PerformActiveFull field to given value.


### GetStartChainedJobs

`func (o *JobStartSpec) GetStartChainedJobs() bool`

GetStartChainedJobs returns the StartChainedJobs field if non-nil, zero value otherwise.

### GetStartChainedJobsOk

`func (o *JobStartSpec) GetStartChainedJobsOk() (*bool, bool)`

GetStartChainedJobsOk returns a tuple with the StartChainedJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartChainedJobs

`func (o *JobStartSpec) SetStartChainedJobs(v bool)`

SetStartChainedJobs sets StartChainedJobs field to given value.

### HasStartChainedJobs

`func (o *JobStartSpec) HasStartChainedJobs() bool`

HasStartChainedJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


