# ScaleOutRepositoriesFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** | Number of repositories to skip. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of repositories to return. | [optional] 
**OrderColumn** | Pointer to [**EScaleOutRepositoryFiltersOrderColumn**](EScaleOutRepositoryFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** | Sorts repositories in the ascending order by the &#x60;orderColumn&#x60; parameter. | [optional] 
**NameFilter** | Pointer to **string** | Filters repositories by the &#x60;nameFilter&#x60; pattern. The pattern can match any repository parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both. | [optional] 

## Methods

### NewScaleOutRepositoriesFilters

`func NewScaleOutRepositoriesFilters() *ScaleOutRepositoriesFilters`

NewScaleOutRepositoriesFilters instantiates a new ScaleOutRepositoriesFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScaleOutRepositoriesFiltersWithDefaults

`func NewScaleOutRepositoriesFiltersWithDefaults() *ScaleOutRepositoriesFilters`

NewScaleOutRepositoriesFiltersWithDefaults instantiates a new ScaleOutRepositoriesFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *ScaleOutRepositoriesFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *ScaleOutRepositoriesFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *ScaleOutRepositoriesFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *ScaleOutRepositoriesFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *ScaleOutRepositoriesFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ScaleOutRepositoriesFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ScaleOutRepositoriesFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ScaleOutRepositoriesFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *ScaleOutRepositoriesFilters) GetOrderColumn() EScaleOutRepositoryFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *ScaleOutRepositoriesFilters) GetOrderColumnOk() (*EScaleOutRepositoryFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *ScaleOutRepositoriesFilters) SetOrderColumn(v EScaleOutRepositoryFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *ScaleOutRepositoriesFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *ScaleOutRepositoriesFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *ScaleOutRepositoriesFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *ScaleOutRepositoriesFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *ScaleOutRepositoriesFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.

### GetNameFilter

`func (o *ScaleOutRepositoriesFilters) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *ScaleOutRepositoriesFilters) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *ScaleOutRepositoriesFilters) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *ScaleOutRepositoriesFilters) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


