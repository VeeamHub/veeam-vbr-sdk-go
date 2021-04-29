# BackupsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BackupModel**](BackupModel.md) | Array of backups. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewBackupsResult

`func NewBackupsResult(data []BackupModel, pagination PaginationResult, ) *BackupsResult`

NewBackupsResult instantiates a new BackupsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupsResultWithDefaults

`func NewBackupsResultWithDefaults() *BackupsResult`

NewBackupsResultWithDefaults instantiates a new BackupsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BackupsResult) GetData() []BackupModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BackupsResult) GetDataOk() (*[]BackupModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BackupsResult) SetData(v []BackupModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *BackupsResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *BackupsResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *BackupsResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


