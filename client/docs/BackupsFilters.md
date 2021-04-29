# BackupsFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**OrderColumn** | Pointer to [**EBackupsFiltersOrderColumn**](EBackupsFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** |  | [optional] 
**NameFilter** | Pointer to **string** |  | [optional] 
**CreatedAfterFilter** | Pointer to **time.Time** |  | [optional] 
**CreatedBeforeFilter** | Pointer to **time.Time** |  | [optional] 
**PlatformIdFilter** | Pointer to **string** |  | [optional] 
**JobIdFilter** | Pointer to **string** |  | [optional] 
**PolicyTagFilter** | Pointer to **string** |  | [optional] 

## Methods

### NewBackupsFilters

`func NewBackupsFilters() *BackupsFilters`

NewBackupsFilters instantiates a new BackupsFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupsFiltersWithDefaults

`func NewBackupsFiltersWithDefaults() *BackupsFilters`

NewBackupsFiltersWithDefaults instantiates a new BackupsFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *BackupsFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *BackupsFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *BackupsFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *BackupsFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *BackupsFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BackupsFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BackupsFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *BackupsFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *BackupsFilters) GetOrderColumn() EBackupsFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *BackupsFilters) GetOrderColumnOk() (*EBackupsFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *BackupsFilters) SetOrderColumn(v EBackupsFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *BackupsFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *BackupsFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *BackupsFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *BackupsFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *BackupsFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.

### GetNameFilter

`func (o *BackupsFilters) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *BackupsFilters) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *BackupsFilters) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *BackupsFilters) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetCreatedAfterFilter

`func (o *BackupsFilters) GetCreatedAfterFilter() time.Time`

GetCreatedAfterFilter returns the CreatedAfterFilter field if non-nil, zero value otherwise.

### GetCreatedAfterFilterOk

`func (o *BackupsFilters) GetCreatedAfterFilterOk() (*time.Time, bool)`

GetCreatedAfterFilterOk returns a tuple with the CreatedAfterFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAfterFilter

`func (o *BackupsFilters) SetCreatedAfterFilter(v time.Time)`

SetCreatedAfterFilter sets CreatedAfterFilter field to given value.

### HasCreatedAfterFilter

`func (o *BackupsFilters) HasCreatedAfterFilter() bool`

HasCreatedAfterFilter returns a boolean if a field has been set.

### GetCreatedBeforeFilter

`func (o *BackupsFilters) GetCreatedBeforeFilter() time.Time`

GetCreatedBeforeFilter returns the CreatedBeforeFilter field if non-nil, zero value otherwise.

### GetCreatedBeforeFilterOk

`func (o *BackupsFilters) GetCreatedBeforeFilterOk() (*time.Time, bool)`

GetCreatedBeforeFilterOk returns a tuple with the CreatedBeforeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBeforeFilter

`func (o *BackupsFilters) SetCreatedBeforeFilter(v time.Time)`

SetCreatedBeforeFilter sets CreatedBeforeFilter field to given value.

### HasCreatedBeforeFilter

`func (o *BackupsFilters) HasCreatedBeforeFilter() bool`

HasCreatedBeforeFilter returns a boolean if a field has been set.

### GetPlatformIdFilter

`func (o *BackupsFilters) GetPlatformIdFilter() string`

GetPlatformIdFilter returns the PlatformIdFilter field if non-nil, zero value otherwise.

### GetPlatformIdFilterOk

`func (o *BackupsFilters) GetPlatformIdFilterOk() (*string, bool)`

GetPlatformIdFilterOk returns a tuple with the PlatformIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIdFilter

`func (o *BackupsFilters) SetPlatformIdFilter(v string)`

SetPlatformIdFilter sets PlatformIdFilter field to given value.

### HasPlatformIdFilter

`func (o *BackupsFilters) HasPlatformIdFilter() bool`

HasPlatformIdFilter returns a boolean if a field has been set.

### GetJobIdFilter

`func (o *BackupsFilters) GetJobIdFilter() string`

GetJobIdFilter returns the JobIdFilter field if non-nil, zero value otherwise.

### GetJobIdFilterOk

`func (o *BackupsFilters) GetJobIdFilterOk() (*string, bool)`

GetJobIdFilterOk returns a tuple with the JobIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIdFilter

`func (o *BackupsFilters) SetJobIdFilter(v string)`

SetJobIdFilter sets JobIdFilter field to given value.

### HasJobIdFilter

`func (o *BackupsFilters) HasJobIdFilter() bool`

HasJobIdFilter returns a boolean if a field has been set.

### GetPolicyTagFilter

`func (o *BackupsFilters) GetPolicyTagFilter() string`

GetPolicyTagFilter returns the PolicyTagFilter field if non-nil, zero value otherwise.

### GetPolicyTagFilterOk

`func (o *BackupsFilters) GetPolicyTagFilterOk() (*string, bool)`

GetPolicyTagFilterOk returns a tuple with the PolicyTagFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTagFilter

`func (o *BackupsFilters) SetPolicyTagFilter(v string)`

SetPolicyTagFilter sets PolicyTagFilter field to given value.

### HasPolicyTagFilter

`func (o *BackupsFilters) HasPolicyTagFilter() bool`

HasPolicyTagFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


