# ProxiesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ProxyModel**](ProxyModel.md) | Array of backup proxies. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewProxiesResult

`func NewProxiesResult(data []ProxyModel, pagination PaginationResult, ) *ProxiesResult`

NewProxiesResult instantiates a new ProxiesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxiesResultWithDefaults

`func NewProxiesResultWithDefaults() *ProxiesResult`

NewProxiesResultWithDefaults instantiates a new ProxiesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProxiesResult) GetData() []ProxyModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProxiesResult) GetDataOk() (*[]ProxyModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProxiesResult) SetData(v []ProxyModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *ProxiesResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ProxiesResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ProxiesResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


