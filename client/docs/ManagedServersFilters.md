# ManagedServersFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**OrderColumn** | Pointer to [**EManagedServersFiltersOrderColumn**](EManagedServersFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** |  | [optional] 
**NameFilter** | Pointer to **string** |  | [optional] 
**TypeFilter** | Pointer to [**EManagedServerType**](EManagedServerType.md) |  | [optional] 
**ViTypeFilter** | Pointer to [**EViHostType**](EViHostType.md) |  | [optional] 

## Methods

### NewManagedServersFilters

`func NewManagedServersFilters() *ManagedServersFilters`

NewManagedServersFilters instantiates a new ManagedServersFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedServersFiltersWithDefaults

`func NewManagedServersFiltersWithDefaults() *ManagedServersFilters`

NewManagedServersFiltersWithDefaults instantiates a new ManagedServersFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *ManagedServersFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *ManagedServersFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *ManagedServersFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *ManagedServersFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *ManagedServersFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ManagedServersFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ManagedServersFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ManagedServersFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *ManagedServersFilters) GetOrderColumn() EManagedServersFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *ManagedServersFilters) GetOrderColumnOk() (*EManagedServersFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *ManagedServersFilters) SetOrderColumn(v EManagedServersFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *ManagedServersFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *ManagedServersFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *ManagedServersFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *ManagedServersFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *ManagedServersFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.

### GetNameFilter

`func (o *ManagedServersFilters) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *ManagedServersFilters) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *ManagedServersFilters) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *ManagedServersFilters) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetTypeFilter

`func (o *ManagedServersFilters) GetTypeFilter() EManagedServerType`

GetTypeFilter returns the TypeFilter field if non-nil, zero value otherwise.

### GetTypeFilterOk

`func (o *ManagedServersFilters) GetTypeFilterOk() (*EManagedServerType, bool)`

GetTypeFilterOk returns a tuple with the TypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFilter

`func (o *ManagedServersFilters) SetTypeFilter(v EManagedServerType)`

SetTypeFilter sets TypeFilter field to given value.

### HasTypeFilter

`func (o *ManagedServersFilters) HasTypeFilter() bool`

HasTypeFilter returns a boolean if a field has been set.

### GetViTypeFilter

`func (o *ManagedServersFilters) GetViTypeFilter() EViHostType`

GetViTypeFilter returns the ViTypeFilter field if non-nil, zero value otherwise.

### GetViTypeFilterOk

`func (o *ManagedServersFilters) GetViTypeFilterOk() (*EViHostType, bool)`

GetViTypeFilterOk returns a tuple with the ViTypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViTypeFilter

`func (o *ManagedServersFilters) SetViTypeFilter(v EViHostType)`

SetViTypeFilter sets ViTypeFilter field to given value.

### HasViTypeFilter

`func (o *ManagedServersFilters) HasViTypeFilter() bool`

HasViTypeFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


