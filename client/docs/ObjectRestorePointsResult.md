# ObjectRestorePointsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ObjectRestorePointModel**](ObjectRestorePointModel.md) | Array of restore points. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewObjectRestorePointsResult

`func NewObjectRestorePointsResult(data []ObjectRestorePointModel, pagination PaginationResult, ) *ObjectRestorePointsResult`

NewObjectRestorePointsResult instantiates a new ObjectRestorePointsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectRestorePointsResultWithDefaults

`func NewObjectRestorePointsResultWithDefaults() *ObjectRestorePointsResult`

NewObjectRestorePointsResultWithDefaults instantiates a new ObjectRestorePointsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ObjectRestorePointsResult) GetData() []ObjectRestorePointModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ObjectRestorePointsResult) GetDataOk() (*[]ObjectRestorePointModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ObjectRestorePointsResult) SetData(v []ObjectRestorePointModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *ObjectRestorePointsResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ObjectRestorePointsResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ObjectRestorePointsResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


