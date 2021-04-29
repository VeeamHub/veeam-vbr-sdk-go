# RepositoryStatesFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** | Skips the specified number of repositories. | [optional] 
**Limit** | Pointer to **int32** | Returns the specified number of repositories. | [optional] 
**OrderColumn** | Pointer to [**ERepositoryStatesFiltersOrderColumn**](ERepositoryStatesFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** | Sorts repositories in the ascending order by the &#x60;orderColumn&#x60; parameter. | [optional] 
**IdFilter** | Pointer to **string** |  | [optional] 
**NameFilter** | Pointer to **string** | Filters repositories by the &#x60;nameFilter&#x60; pattern. The pattern can match any repository parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | [optional] 
**TypeFilter** | Pointer to [**ERepositoryType**](ERepositoryType.md) |  | [optional] 
**CapacityFilter** | Pointer to **float64** | Filters repositories by repository capacity. | [optional] 
**FreeSpaceFilter** | Pointer to **float64** | Filters repositories by repository free space. | [optional] 
**UsedSpaceFilter** | Pointer to **float64** | Filters repositories by repository used space. | [optional] 

## Methods

### NewRepositoryStatesFilters

`func NewRepositoryStatesFilters() *RepositoryStatesFilters`

NewRepositoryStatesFilters instantiates a new RepositoryStatesFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryStatesFiltersWithDefaults

`func NewRepositoryStatesFiltersWithDefaults() *RepositoryStatesFilters`

NewRepositoryStatesFiltersWithDefaults instantiates a new RepositoryStatesFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *RepositoryStatesFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *RepositoryStatesFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *RepositoryStatesFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *RepositoryStatesFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *RepositoryStatesFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RepositoryStatesFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RepositoryStatesFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RepositoryStatesFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *RepositoryStatesFilters) GetOrderColumn() ERepositoryStatesFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *RepositoryStatesFilters) GetOrderColumnOk() (*ERepositoryStatesFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *RepositoryStatesFilters) SetOrderColumn(v ERepositoryStatesFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *RepositoryStatesFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *RepositoryStatesFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *RepositoryStatesFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *RepositoryStatesFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *RepositoryStatesFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.

### GetIdFilter

`func (o *RepositoryStatesFilters) GetIdFilter() string`

GetIdFilter returns the IdFilter field if non-nil, zero value otherwise.

### GetIdFilterOk

`func (o *RepositoryStatesFilters) GetIdFilterOk() (*string, bool)`

GetIdFilterOk returns a tuple with the IdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdFilter

`func (o *RepositoryStatesFilters) SetIdFilter(v string)`

SetIdFilter sets IdFilter field to given value.

### HasIdFilter

`func (o *RepositoryStatesFilters) HasIdFilter() bool`

HasIdFilter returns a boolean if a field has been set.

### GetNameFilter

`func (o *RepositoryStatesFilters) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *RepositoryStatesFilters) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *RepositoryStatesFilters) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *RepositoryStatesFilters) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetTypeFilter

`func (o *RepositoryStatesFilters) GetTypeFilter() ERepositoryType`

GetTypeFilter returns the TypeFilter field if non-nil, zero value otherwise.

### GetTypeFilterOk

`func (o *RepositoryStatesFilters) GetTypeFilterOk() (*ERepositoryType, bool)`

GetTypeFilterOk returns a tuple with the TypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFilter

`func (o *RepositoryStatesFilters) SetTypeFilter(v ERepositoryType)`

SetTypeFilter sets TypeFilter field to given value.

### HasTypeFilter

`func (o *RepositoryStatesFilters) HasTypeFilter() bool`

HasTypeFilter returns a boolean if a field has been set.

### GetCapacityFilter

`func (o *RepositoryStatesFilters) GetCapacityFilter() float64`

GetCapacityFilter returns the CapacityFilter field if non-nil, zero value otherwise.

### GetCapacityFilterOk

`func (o *RepositoryStatesFilters) GetCapacityFilterOk() (*float64, bool)`

GetCapacityFilterOk returns a tuple with the CapacityFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityFilter

`func (o *RepositoryStatesFilters) SetCapacityFilter(v float64)`

SetCapacityFilter sets CapacityFilter field to given value.

### HasCapacityFilter

`func (o *RepositoryStatesFilters) HasCapacityFilter() bool`

HasCapacityFilter returns a boolean if a field has been set.

### GetFreeSpaceFilter

`func (o *RepositoryStatesFilters) GetFreeSpaceFilter() float64`

GetFreeSpaceFilter returns the FreeSpaceFilter field if non-nil, zero value otherwise.

### GetFreeSpaceFilterOk

`func (o *RepositoryStatesFilters) GetFreeSpaceFilterOk() (*float64, bool)`

GetFreeSpaceFilterOk returns a tuple with the FreeSpaceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpaceFilter

`func (o *RepositoryStatesFilters) SetFreeSpaceFilter(v float64)`

SetFreeSpaceFilter sets FreeSpaceFilter field to given value.

### HasFreeSpaceFilter

`func (o *RepositoryStatesFilters) HasFreeSpaceFilter() bool`

HasFreeSpaceFilter returns a boolean if a field has been set.

### GetUsedSpaceFilter

`func (o *RepositoryStatesFilters) GetUsedSpaceFilter() float64`

GetUsedSpaceFilter returns the UsedSpaceFilter field if non-nil, zero value otherwise.

### GetUsedSpaceFilterOk

`func (o *RepositoryStatesFilters) GetUsedSpaceFilterOk() (*float64, bool)`

GetUsedSpaceFilterOk returns a tuple with the UsedSpaceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpaceFilter

`func (o *RepositoryStatesFilters) SetUsedSpaceFilter(v float64)`

SetUsedSpaceFilter sets UsedSpaceFilter field to given value.

### HasUsedSpaceFilter

`func (o *RepositoryStatesFilters) HasUsedSpaceFilter() bool`

HasUsedSpaceFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


