# PaginationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of results. | 
**Count** | **int32** | Number of returned results. | 
**Skip** | Pointer to **int32** | Number of skipped results. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of results to return. | [optional] 

## Methods

### NewPaginationResult

`func NewPaginationResult(total int32, count int32, ) *PaginationResult`

NewPaginationResult instantiates a new PaginationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationResultWithDefaults

`func NewPaginationResultWithDefaults() *PaginationResult`

NewPaginationResultWithDefaults instantiates a new PaginationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *PaginationResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaginationResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaginationResult) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetCount

`func (o *PaginationResult) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginationResult) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginationResult) SetCount(v int32)`

SetCount sets Count field to given value.


### GetSkip

`func (o *PaginationResult) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *PaginationResult) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *PaginationResult) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *PaginationResult) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *PaginationResult) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaginationResult) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaginationResult) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PaginationResult) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


