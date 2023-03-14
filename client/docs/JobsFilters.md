# JobsFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** | Skips the specified number of jobs. | [optional] 
**Limit** | Pointer to **int32** | Returns the specified number of jobs. | [optional] 
**OrderColumn** | Pointer to [**EJobFiltersOrderColumn**](EJobFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** | Sorts jobs in the ascending order by the &#x60;orderColumn&#x60; parameter. | [optional] 
**NameFilter** | Pointer to **string** | Filters jobs by the &#x60;nameFilter&#x60; pattern. The pattern can match any job parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both. | [optional] 
**TypeFilter** | Pointer to [**EJobType**](EJobType.md) |  | [optional] 

## Methods

### NewJobsFilters

`func NewJobsFilters() *JobsFilters`

NewJobsFilters instantiates a new JobsFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobsFiltersWithDefaults

`func NewJobsFiltersWithDefaults() *JobsFilters`

NewJobsFiltersWithDefaults instantiates a new JobsFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *JobsFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *JobsFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *JobsFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *JobsFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *JobsFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *JobsFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *JobsFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *JobsFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *JobsFilters) GetOrderColumn() EJobFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *JobsFilters) GetOrderColumnOk() (*EJobFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *JobsFilters) SetOrderColumn(v EJobFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *JobsFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *JobsFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *JobsFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *JobsFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *JobsFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.

### GetNameFilter

`func (o *JobsFilters) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *JobsFilters) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *JobsFilters) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *JobsFilters) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetTypeFilter

`func (o *JobsFilters) GetTypeFilter() EJobType`

GetTypeFilter returns the TypeFilter field if non-nil, zero value otherwise.

### GetTypeFilterOk

`func (o *JobsFilters) GetTypeFilterOk() (*EJobType, bool)`

GetTypeFilterOk returns a tuple with the TypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFilter

`func (o *JobsFilters) SetTypeFilter(v EJobType)`

SetTypeFilter sets TypeFilter field to given value.

### HasTypeFilter

`func (o *JobsFilters) HasTypeFilter() bool`

HasTypeFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


