# ObjectRestorePointDisksResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ObjectRestorePointDiskModel**](ObjectRestorePointDiskModel.md) | Array of disks. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewObjectRestorePointDisksResult

`func NewObjectRestorePointDisksResult(data []ObjectRestorePointDiskModel, pagination PaginationResult, ) *ObjectRestorePointDisksResult`

NewObjectRestorePointDisksResult instantiates a new ObjectRestorePointDisksResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectRestorePointDisksResultWithDefaults

`func NewObjectRestorePointDisksResultWithDefaults() *ObjectRestorePointDisksResult`

NewObjectRestorePointDisksResultWithDefaults instantiates a new ObjectRestorePointDisksResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ObjectRestorePointDisksResult) GetData() []ObjectRestorePointDiskModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ObjectRestorePointDisksResult) GetDataOk() (*[]ObjectRestorePointDiskModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ObjectRestorePointDisksResult) SetData(v []ObjectRestorePointDiskModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *ObjectRestorePointDisksResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ObjectRestorePointDisksResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ObjectRestorePointDisksResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


