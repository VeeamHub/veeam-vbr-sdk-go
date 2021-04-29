# RepositoryStatesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]RepositoryStateModel**](RepositoryStateModel.md) | Array of repository states. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewRepositoryStatesResult

`func NewRepositoryStatesResult(data []RepositoryStateModel, pagination PaginationResult, ) *RepositoryStatesResult`

NewRepositoryStatesResult instantiates a new RepositoryStatesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryStatesResultWithDefaults

`func NewRepositoryStatesResultWithDefaults() *RepositoryStatesResult`

NewRepositoryStatesResultWithDefaults instantiates a new RepositoryStatesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RepositoryStatesResult) GetData() []RepositoryStateModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RepositoryStatesResult) GetDataOk() (*[]RepositoryStateModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RepositoryStatesResult) SetData(v []RepositoryStateModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *RepositoryStatesResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RepositoryStatesResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RepositoryStatesResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


