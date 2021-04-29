# JobsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]JobModel**](JobModel.md) | Array of jobs. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewJobsResult

`func NewJobsResult(data []JobModel, pagination PaginationResult, ) *JobsResult`

NewJobsResult instantiates a new JobsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobsResultWithDefaults

`func NewJobsResultWithDefaults() *JobsResult`

NewJobsResultWithDefaults instantiates a new JobsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *JobsResult) GetData() []JobModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JobsResult) GetDataOk() (*[]JobModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JobsResult) SetData(v []JobModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *JobsResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *JobsResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *JobsResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


