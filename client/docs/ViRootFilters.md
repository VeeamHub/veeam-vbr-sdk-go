# ViRootFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** | Number of VMware vSphere servers to skip. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of VMware vSphere servers to return. | [optional] 
**OrderColumn** | Pointer to [**EViRootFiltersOrderColumn**](EViRootFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** | Sorts VMware vSphere servers in the ascending order by the &#x60;orderColumn&#x60; parameter. | [optional] 
**NameFilter** | Pointer to **string** | Filters VMware vSphere servers by the &#x60;nameFilter&#x60; pattern. The pattern can match any VMware vSphere server parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | [optional] 

## Methods

### NewViRootFilters

`func NewViRootFilters() *ViRootFilters`

NewViRootFilters instantiates a new ViRootFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViRootFiltersWithDefaults

`func NewViRootFiltersWithDefaults() *ViRootFilters`

NewViRootFiltersWithDefaults instantiates a new ViRootFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *ViRootFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *ViRootFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *ViRootFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *ViRootFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *ViRootFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ViRootFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ViRootFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ViRootFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *ViRootFilters) GetOrderColumn() EViRootFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *ViRootFilters) GetOrderColumnOk() (*EViRootFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *ViRootFilters) SetOrderColumn(v EViRootFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *ViRootFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *ViRootFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *ViRootFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *ViRootFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *ViRootFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.

### GetNameFilter

`func (o *ViRootFilters) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *ViRootFilters) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *ViRootFilters) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *ViRootFilters) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


