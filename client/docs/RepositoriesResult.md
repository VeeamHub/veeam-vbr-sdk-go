# RepositoriesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]RepositoryModel**](RepositoryModel.md) | Array of backup repositories. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewRepositoriesResult

`func NewRepositoriesResult(data []RepositoryModel, pagination PaginationResult, ) *RepositoriesResult`

NewRepositoriesResult instantiates a new RepositoriesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoriesResultWithDefaults

`func NewRepositoriesResultWithDefaults() *RepositoriesResult`

NewRepositoriesResultWithDefaults instantiates a new RepositoriesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RepositoriesResult) GetData() []RepositoryModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RepositoriesResult) GetDataOk() (*[]RepositoryModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RepositoriesResult) SetData(v []RepositoryModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *RepositoriesResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RepositoriesResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RepositoriesResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


