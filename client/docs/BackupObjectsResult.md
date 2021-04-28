# BackupObjectsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BackupObjectModel**](BackupObjectModel.md) | Array of backup objects. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewBackupObjectsResult

`func NewBackupObjectsResult(data []BackupObjectModel, pagination PaginationResult, ) *BackupObjectsResult`

NewBackupObjectsResult instantiates a new BackupObjectsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupObjectsResultWithDefaults

`func NewBackupObjectsResultWithDefaults() *BackupObjectsResult`

NewBackupObjectsResultWithDefaults instantiates a new BackupObjectsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BackupObjectsResult) GetData() []BackupObjectModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BackupObjectsResult) GetDataOk() (*[]BackupObjectModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BackupObjectsResult) SetData(v []BackupObjectModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *BackupObjectsResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *BackupObjectsResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *BackupObjectsResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


