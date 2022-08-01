# VmwareFcdInstantRecoveryMountsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VmwareFcdInstantRecoveryMount**](VmwareFcdInstantRecoveryMount.md) | Array of FCD mounts. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewVmwareFcdInstantRecoveryMountsResult

`func NewVmwareFcdInstantRecoveryMountsResult(data []VmwareFcdInstantRecoveryMount, pagination PaginationResult, ) *VmwareFcdInstantRecoveryMountsResult`

NewVmwareFcdInstantRecoveryMountsResult instantiates a new VmwareFcdInstantRecoveryMountsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareFcdInstantRecoveryMountsResultWithDefaults

`func NewVmwareFcdInstantRecoveryMountsResultWithDefaults() *VmwareFcdInstantRecoveryMountsResult`

NewVmwareFcdInstantRecoveryMountsResultWithDefaults instantiates a new VmwareFcdInstantRecoveryMountsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VmwareFcdInstantRecoveryMountsResult) GetData() []VmwareFcdInstantRecoveryMount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VmwareFcdInstantRecoveryMountsResult) GetDataOk() (*[]VmwareFcdInstantRecoveryMount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VmwareFcdInstantRecoveryMountsResult) SetData(v []VmwareFcdInstantRecoveryMount)`

SetData sets Data field to given value.


### GetPagination

`func (o *VmwareFcdInstantRecoveryMountsResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *VmwareFcdInstantRecoveryMountsResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *VmwareFcdInstantRecoveryMountsResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


