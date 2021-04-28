# ProxiesFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**OrderColumn** | Pointer to [**EProxiesFiltersOrderColumn**](EProxiesFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** |  | [optional] 
**NameFilter** | Pointer to **string** |  | [optional] 
**TypeFilter** | Pointer to [**EProxyType**](EProxyType.md) |  | [optional] 
**HostIdFilter** | Pointer to **string** |  | [optional] 

## Methods

### NewProxiesFilters

`func NewProxiesFilters() *ProxiesFilters`

NewProxiesFilters instantiates a new ProxiesFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxiesFiltersWithDefaults

`func NewProxiesFiltersWithDefaults() *ProxiesFilters`

NewProxiesFiltersWithDefaults instantiates a new ProxiesFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *ProxiesFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *ProxiesFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *ProxiesFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *ProxiesFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *ProxiesFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ProxiesFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ProxiesFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ProxiesFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *ProxiesFilters) GetOrderColumn() EProxiesFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *ProxiesFilters) GetOrderColumnOk() (*EProxiesFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *ProxiesFilters) SetOrderColumn(v EProxiesFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *ProxiesFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *ProxiesFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *ProxiesFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *ProxiesFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *ProxiesFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.

### GetNameFilter

`func (o *ProxiesFilters) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *ProxiesFilters) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *ProxiesFilters) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *ProxiesFilters) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetTypeFilter

`func (o *ProxiesFilters) GetTypeFilter() EProxyType`

GetTypeFilter returns the TypeFilter field if non-nil, zero value otherwise.

### GetTypeFilterOk

`func (o *ProxiesFilters) GetTypeFilterOk() (*EProxyType, bool)`

GetTypeFilterOk returns a tuple with the TypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFilter

`func (o *ProxiesFilters) SetTypeFilter(v EProxyType)`

SetTypeFilter sets TypeFilter field to given value.

### HasTypeFilter

`func (o *ProxiesFilters) HasTypeFilter() bool`

HasTypeFilter returns a boolean if a field has been set.

### GetHostIdFilter

`func (o *ProxiesFilters) GetHostIdFilter() string`

GetHostIdFilter returns the HostIdFilter field if non-nil, zero value otherwise.

### GetHostIdFilterOk

`func (o *ProxiesFilters) GetHostIdFilterOk() (*string, bool)`

GetHostIdFilterOk returns a tuple with the HostIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIdFilter

`func (o *ProxiesFilters) SetHostIdFilter(v string)`

SetHostIdFilter sets HostIdFilter field to given value.

### HasHostIdFilter

`func (o *ProxiesFilters) HasHostIdFilter() bool`

HasHostIdFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


