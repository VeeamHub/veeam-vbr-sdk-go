# InstantViVMRecoveryMountsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]InstantViVMRecoveryMount**](InstantViVMRecoveryMount.md) | Array of VM mounts. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewInstantViVMRecoveryMountsResult

`func NewInstantViVMRecoveryMountsResult(data []InstantViVMRecoveryMount, pagination PaginationResult, ) *InstantViVMRecoveryMountsResult`

NewInstantViVMRecoveryMountsResult instantiates a new InstantViVMRecoveryMountsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstantViVMRecoveryMountsResultWithDefaults

`func NewInstantViVMRecoveryMountsResultWithDefaults() *InstantViVMRecoveryMountsResult`

NewInstantViVMRecoveryMountsResultWithDefaults instantiates a new InstantViVMRecoveryMountsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InstantViVMRecoveryMountsResult) GetData() []InstantViVMRecoveryMount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InstantViVMRecoveryMountsResult) GetDataOk() (*[]InstantViVMRecoveryMount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InstantViVMRecoveryMountsResult) SetData(v []InstantViVMRecoveryMount)`

SetData sets Data field to given value.


### GetPagination

`func (o *InstantViVMRecoveryMountsResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *InstantViVMRecoveryMountsResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *InstantViVMRecoveryMountsResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


