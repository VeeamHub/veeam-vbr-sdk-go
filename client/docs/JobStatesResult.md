# JobStatesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]JobStateModel**](JobStateModel.md) | Array of job states. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewJobStatesResult

`func NewJobStatesResult(data []JobStateModel, pagination PaginationResult, ) *JobStatesResult`

NewJobStatesResult instantiates a new JobStatesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobStatesResultWithDefaults

`func NewJobStatesResultWithDefaults() *JobStatesResult`

NewJobStatesResultWithDefaults instantiates a new JobStatesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *JobStatesResult) GetData() []JobStateModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *JobStatesResult) GetDataOk() (*[]JobStateModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *JobStatesResult) SetData(v []JobStateModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *JobStatesResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *JobStatesResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *JobStatesResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


