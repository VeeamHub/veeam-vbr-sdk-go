# ComputerRecoveryTokenResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ComputerRecoveryTokenModel**](ComputerRecoveryTokenModel.md) | Array of recovery tokens. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewComputerRecoveryTokenResult

`func NewComputerRecoveryTokenResult(data []ComputerRecoveryTokenModel, pagination PaginationResult, ) *ComputerRecoveryTokenResult`

NewComputerRecoveryTokenResult instantiates a new ComputerRecoveryTokenResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerRecoveryTokenResultWithDefaults

`func NewComputerRecoveryTokenResultWithDefaults() *ComputerRecoveryTokenResult`

NewComputerRecoveryTokenResultWithDefaults instantiates a new ComputerRecoveryTokenResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ComputerRecoveryTokenResult) GetData() []ComputerRecoveryTokenModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ComputerRecoveryTokenResult) GetDataOk() (*[]ComputerRecoveryTokenModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ComputerRecoveryTokenResult) SetData(v []ComputerRecoveryTokenModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *ComputerRecoveryTokenResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ComputerRecoveryTokenResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ComputerRecoveryTokenResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


