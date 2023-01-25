# ComputerRecoveryTokenFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**OrderColumn** | Pointer to [**EComputerRecoveryTokenFiltersOrderColumn**](EComputerRecoveryTokenFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** |  | [optional] 
**NameFilter** | Pointer to **string** |  | [optional] 
**ExpirationDateBefore** | Pointer to **time.Time** |  | [optional] 
**ExpirationDateAfter** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewComputerRecoveryTokenFilters

`func NewComputerRecoveryTokenFilters() *ComputerRecoveryTokenFilters`

NewComputerRecoveryTokenFilters instantiates a new ComputerRecoveryTokenFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerRecoveryTokenFiltersWithDefaults

`func NewComputerRecoveryTokenFiltersWithDefaults() *ComputerRecoveryTokenFilters`

NewComputerRecoveryTokenFiltersWithDefaults instantiates a new ComputerRecoveryTokenFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *ComputerRecoveryTokenFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *ComputerRecoveryTokenFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *ComputerRecoveryTokenFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *ComputerRecoveryTokenFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *ComputerRecoveryTokenFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ComputerRecoveryTokenFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ComputerRecoveryTokenFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ComputerRecoveryTokenFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *ComputerRecoveryTokenFilters) GetOrderColumn() EComputerRecoveryTokenFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *ComputerRecoveryTokenFilters) GetOrderColumnOk() (*EComputerRecoveryTokenFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *ComputerRecoveryTokenFilters) SetOrderColumn(v EComputerRecoveryTokenFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *ComputerRecoveryTokenFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *ComputerRecoveryTokenFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *ComputerRecoveryTokenFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *ComputerRecoveryTokenFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *ComputerRecoveryTokenFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.

### GetNameFilter

`func (o *ComputerRecoveryTokenFilters) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *ComputerRecoveryTokenFilters) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *ComputerRecoveryTokenFilters) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *ComputerRecoveryTokenFilters) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetExpirationDateBefore

`func (o *ComputerRecoveryTokenFilters) GetExpirationDateBefore() time.Time`

GetExpirationDateBefore returns the ExpirationDateBefore field if non-nil, zero value otherwise.

### GetExpirationDateBeforeOk

`func (o *ComputerRecoveryTokenFilters) GetExpirationDateBeforeOk() (*time.Time, bool)`

GetExpirationDateBeforeOk returns a tuple with the ExpirationDateBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateBefore

`func (o *ComputerRecoveryTokenFilters) SetExpirationDateBefore(v time.Time)`

SetExpirationDateBefore sets ExpirationDateBefore field to given value.

### HasExpirationDateBefore

`func (o *ComputerRecoveryTokenFilters) HasExpirationDateBefore() bool`

HasExpirationDateBefore returns a boolean if a field has been set.

### GetExpirationDateAfter

`func (o *ComputerRecoveryTokenFilters) GetExpirationDateAfter() time.Time`

GetExpirationDateAfter returns the ExpirationDateAfter field if non-nil, zero value otherwise.

### GetExpirationDateAfterOk

`func (o *ComputerRecoveryTokenFilters) GetExpirationDateAfterOk() (*time.Time, bool)`

GetExpirationDateAfterOk returns a tuple with the ExpirationDateAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateAfter

`func (o *ComputerRecoveryTokenFilters) SetExpirationDateAfter(v time.Time)`

SetExpirationDateAfter sets ExpirationDateAfter field to given value.

### HasExpirationDateAfter

`func (o *ComputerRecoveryTokenFilters) HasExpirationDateAfter() bool`

HasExpirationDateAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


