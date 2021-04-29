# ScaleOutRepositoriesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ScaleOutRepositoryModel**](ScaleOutRepositoryModel.md) | Array of scale-out backup repositories. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewScaleOutRepositoriesResult

`func NewScaleOutRepositoriesResult(data []ScaleOutRepositoryModel, pagination PaginationResult, ) *ScaleOutRepositoriesResult`

NewScaleOutRepositoriesResult instantiates a new ScaleOutRepositoriesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScaleOutRepositoriesResultWithDefaults

`func NewScaleOutRepositoriesResultWithDefaults() *ScaleOutRepositoriesResult`

NewScaleOutRepositoriesResultWithDefaults instantiates a new ScaleOutRepositoriesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ScaleOutRepositoriesResult) GetData() []ScaleOutRepositoryModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ScaleOutRepositoriesResult) GetDataOk() (*[]ScaleOutRepositoryModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ScaleOutRepositoriesResult) SetData(v []ScaleOutRepositoryModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *ScaleOutRepositoriesResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ScaleOutRepositoriesResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ScaleOutRepositoriesResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


