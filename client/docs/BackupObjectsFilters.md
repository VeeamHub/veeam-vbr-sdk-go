# BackupObjectsFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**OrderColumn** | Pointer to [**EBackupObjectsFiltersOrderColumn**](EBackupObjectsFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** |  | [optional] 
**NameFilter** | Pointer to **string** |  | [optional] 
**PlatformNameFilter** | Pointer to [**EPlatformType**](EPlatformType.md) |  | [optional] 
**PlatformIdFilter** | Pointer to **string** |  | [optional] 
**TypeFilter** | Pointer to **string** |  | [optional] 
**ViTypeFilter** | Pointer to **string** |  | [optional] 

## Methods

### NewBackupObjectsFilters

`func NewBackupObjectsFilters() *BackupObjectsFilters`

NewBackupObjectsFilters instantiates a new BackupObjectsFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupObjectsFiltersWithDefaults

`func NewBackupObjectsFiltersWithDefaults() *BackupObjectsFilters`

NewBackupObjectsFiltersWithDefaults instantiates a new BackupObjectsFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *BackupObjectsFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *BackupObjectsFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *BackupObjectsFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *BackupObjectsFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *BackupObjectsFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BackupObjectsFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BackupObjectsFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *BackupObjectsFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *BackupObjectsFilters) GetOrderColumn() EBackupObjectsFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *BackupObjectsFilters) GetOrderColumnOk() (*EBackupObjectsFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *BackupObjectsFilters) SetOrderColumn(v EBackupObjectsFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *BackupObjectsFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *BackupObjectsFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *BackupObjectsFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *BackupObjectsFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *BackupObjectsFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.

### GetNameFilter

`func (o *BackupObjectsFilters) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *BackupObjectsFilters) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *BackupObjectsFilters) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *BackupObjectsFilters) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetPlatformNameFilter

`func (o *BackupObjectsFilters) GetPlatformNameFilter() EPlatformType`

GetPlatformNameFilter returns the PlatformNameFilter field if non-nil, zero value otherwise.

### GetPlatformNameFilterOk

`func (o *BackupObjectsFilters) GetPlatformNameFilterOk() (*EPlatformType, bool)`

GetPlatformNameFilterOk returns a tuple with the PlatformNameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformNameFilter

`func (o *BackupObjectsFilters) SetPlatformNameFilter(v EPlatformType)`

SetPlatformNameFilter sets PlatformNameFilter field to given value.

### HasPlatformNameFilter

`func (o *BackupObjectsFilters) HasPlatformNameFilter() bool`

HasPlatformNameFilter returns a boolean if a field has been set.

### GetPlatformIdFilter

`func (o *BackupObjectsFilters) GetPlatformIdFilter() string`

GetPlatformIdFilter returns the PlatformIdFilter field if non-nil, zero value otherwise.

### GetPlatformIdFilterOk

`func (o *BackupObjectsFilters) GetPlatformIdFilterOk() (*string, bool)`

GetPlatformIdFilterOk returns a tuple with the PlatformIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIdFilter

`func (o *BackupObjectsFilters) SetPlatformIdFilter(v string)`

SetPlatformIdFilter sets PlatformIdFilter field to given value.

### HasPlatformIdFilter

`func (o *BackupObjectsFilters) HasPlatformIdFilter() bool`

HasPlatformIdFilter returns a boolean if a field has been set.

### GetTypeFilter

`func (o *BackupObjectsFilters) GetTypeFilter() string`

GetTypeFilter returns the TypeFilter field if non-nil, zero value otherwise.

### GetTypeFilterOk

`func (o *BackupObjectsFilters) GetTypeFilterOk() (*string, bool)`

GetTypeFilterOk returns a tuple with the TypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFilter

`func (o *BackupObjectsFilters) SetTypeFilter(v string)`

SetTypeFilter sets TypeFilter field to given value.

### HasTypeFilter

`func (o *BackupObjectsFilters) HasTypeFilter() bool`

HasTypeFilter returns a boolean if a field has been set.

### GetViTypeFilter

`func (o *BackupObjectsFilters) GetViTypeFilter() string`

GetViTypeFilter returns the ViTypeFilter field if non-nil, zero value otherwise.

### GetViTypeFilterOk

`func (o *BackupObjectsFilters) GetViTypeFilterOk() (*string, bool)`

GetViTypeFilterOk returns a tuple with the ViTypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViTypeFilter

`func (o *BackupObjectsFilters) SetViTypeFilter(v string)`

SetViTypeFilter sets ViTypeFilter field to given value.

### HasViTypeFilter

`func (o *BackupObjectsFilters) HasViTypeFilter() bool`

HasViTypeFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


