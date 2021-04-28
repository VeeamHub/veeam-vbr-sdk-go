# ViRootsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VmwareObjectSizeModel**](VmwareObjectSizeModel.md) | Array of VMware vSphere servers. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewViRootsResult

`func NewViRootsResult(data []VmwareObjectSizeModel, pagination PaginationResult, ) *ViRootsResult`

NewViRootsResult instantiates a new ViRootsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViRootsResultWithDefaults

`func NewViRootsResultWithDefaults() *ViRootsResult`

NewViRootsResultWithDefaults instantiates a new ViRootsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ViRootsResult) GetData() []VmwareObjectSizeModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ViRootsResult) GetDataOk() (*[]VmwareObjectSizeModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ViRootsResult) SetData(v []VmwareObjectSizeModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *ViRootsResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ViRootsResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ViRootsResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


