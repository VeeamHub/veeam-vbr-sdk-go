# RepositoriesFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**OrderColumn** | Pointer to [**ERepositoryFiltersOrderColumn**](ERepositoryFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** |  | [optional] 
**NameFilter** | Pointer to **string** |  | [optional] 
**TypeFilter** | Pointer to [**ERepositoryType**](ERepositoryType.md) |  | [optional] 
**HostIdFilter** | Pointer to **string** |  | [optional] 
**PathFilter** | Pointer to **string** |  | [optional] 
**VmbApiFilter** | Pointer to **string** | VmbApiFilterModel as json serialized in base64 format (see VmbApiFilterModel) | [optional] 

## Methods

### NewRepositoriesFilters

`func NewRepositoriesFilters() *RepositoriesFilters`

NewRepositoriesFilters instantiates a new RepositoriesFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoriesFiltersWithDefaults

`func NewRepositoriesFiltersWithDefaults() *RepositoriesFilters`

NewRepositoriesFiltersWithDefaults instantiates a new RepositoriesFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *RepositoriesFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *RepositoriesFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *RepositoriesFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *RepositoriesFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *RepositoriesFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RepositoriesFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RepositoriesFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RepositoriesFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *RepositoriesFilters) GetOrderColumn() ERepositoryFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *RepositoriesFilters) GetOrderColumnOk() (*ERepositoryFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *RepositoriesFilters) SetOrderColumn(v ERepositoryFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *RepositoriesFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *RepositoriesFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *RepositoriesFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *RepositoriesFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *RepositoriesFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.

### GetNameFilter

`func (o *RepositoriesFilters) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *RepositoriesFilters) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *RepositoriesFilters) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *RepositoriesFilters) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetTypeFilter

`func (o *RepositoriesFilters) GetTypeFilter() ERepositoryType`

GetTypeFilter returns the TypeFilter field if non-nil, zero value otherwise.

### GetTypeFilterOk

`func (o *RepositoriesFilters) GetTypeFilterOk() (*ERepositoryType, bool)`

GetTypeFilterOk returns a tuple with the TypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFilter

`func (o *RepositoriesFilters) SetTypeFilter(v ERepositoryType)`

SetTypeFilter sets TypeFilter field to given value.

### HasTypeFilter

`func (o *RepositoriesFilters) HasTypeFilter() bool`

HasTypeFilter returns a boolean if a field has been set.

### GetHostIdFilter

`func (o *RepositoriesFilters) GetHostIdFilter() string`

GetHostIdFilter returns the HostIdFilter field if non-nil, zero value otherwise.

### GetHostIdFilterOk

`func (o *RepositoriesFilters) GetHostIdFilterOk() (*string, bool)`

GetHostIdFilterOk returns a tuple with the HostIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIdFilter

`func (o *RepositoriesFilters) SetHostIdFilter(v string)`

SetHostIdFilter sets HostIdFilter field to given value.

### HasHostIdFilter

`func (o *RepositoriesFilters) HasHostIdFilter() bool`

HasHostIdFilter returns a boolean if a field has been set.

### GetPathFilter

`func (o *RepositoriesFilters) GetPathFilter() string`

GetPathFilter returns the PathFilter field if non-nil, zero value otherwise.

### GetPathFilterOk

`func (o *RepositoriesFilters) GetPathFilterOk() (*string, bool)`

GetPathFilterOk returns a tuple with the PathFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathFilter

`func (o *RepositoriesFilters) SetPathFilter(v string)`

SetPathFilter sets PathFilter field to given value.

### HasPathFilter

`func (o *RepositoriesFilters) HasPathFilter() bool`

HasPathFilter returns a boolean if a field has been set.

### GetVmbApiFilter

`func (o *RepositoriesFilters) GetVmbApiFilter() string`

GetVmbApiFilter returns the VmbApiFilter field if non-nil, zero value otherwise.

### GetVmbApiFilterOk

`func (o *RepositoriesFilters) GetVmbApiFilterOk() (*string, bool)`

GetVmbApiFilterOk returns a tuple with the VmbApiFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmbApiFilter

`func (o *RepositoriesFilters) SetVmbApiFilter(v string)`

SetVmbApiFilter sets VmbApiFilter field to given value.

### HasVmbApiFilter

`func (o *RepositoriesFilters) HasVmbApiFilter() bool`

HasVmbApiFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


