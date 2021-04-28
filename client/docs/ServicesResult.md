# ServicesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ServicesModel**](ServicesModel.md) | Array of services. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewServicesResult

`func NewServicesResult(data []ServicesModel, pagination PaginationResult, ) *ServicesResult`

NewServicesResult instantiates a new ServicesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicesResultWithDefaults

`func NewServicesResultWithDefaults() *ServicesResult`

NewServicesResultWithDefaults instantiates a new ServicesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServicesResult) GetData() []ServicesModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServicesResult) GetDataOk() (*[]ServicesModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServicesResult) SetData(v []ServicesModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *ServicesResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ServicesResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ServicesResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


