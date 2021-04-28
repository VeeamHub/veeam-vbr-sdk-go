# ObjectRestorePointsFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**OrderColumn** | Pointer to [**EObjectRestorePointsFiltersOrderColumn**](EObjectRestorePointsFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** |  | [optional] 
**CreatedAfterFilter** | Pointer to **time.Time** |  | [optional] 
**CreatedBeforeFilter** | Pointer to **time.Time** |  | [optional] 
**NameFilter** | Pointer to **string** |  | [optional] 
**PlatformNameFilter** | Pointer to [**EPlatformType**](EPlatformType.md) |  | [optional] 
**PlatformIdFilter** | Pointer to **string** |  | [optional] 
**BackupIdFilter** | Pointer to **string** |  | [optional] 
**BackupObjectIdFilter** | Pointer to **string** |  | [optional] 

## Methods

### NewObjectRestorePointsFilters

`func NewObjectRestorePointsFilters() *ObjectRestorePointsFilters`

NewObjectRestorePointsFilters instantiates a new ObjectRestorePointsFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectRestorePointsFiltersWithDefaults

`func NewObjectRestorePointsFiltersWithDefaults() *ObjectRestorePointsFilters`

NewObjectRestorePointsFiltersWithDefaults instantiates a new ObjectRestorePointsFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *ObjectRestorePointsFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *ObjectRestorePointsFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *ObjectRestorePointsFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *ObjectRestorePointsFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *ObjectRestorePointsFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ObjectRestorePointsFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ObjectRestorePointsFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ObjectRestorePointsFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *ObjectRestorePointsFilters) GetOrderColumn() EObjectRestorePointsFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *ObjectRestorePointsFilters) GetOrderColumnOk() (*EObjectRestorePointsFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *ObjectRestorePointsFilters) SetOrderColumn(v EObjectRestorePointsFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *ObjectRestorePointsFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *ObjectRestorePointsFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *ObjectRestorePointsFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *ObjectRestorePointsFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *ObjectRestorePointsFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.

### GetCreatedAfterFilter

`func (o *ObjectRestorePointsFilters) GetCreatedAfterFilter() time.Time`

GetCreatedAfterFilter returns the CreatedAfterFilter field if non-nil, zero value otherwise.

### GetCreatedAfterFilterOk

`func (o *ObjectRestorePointsFilters) GetCreatedAfterFilterOk() (*time.Time, bool)`

GetCreatedAfterFilterOk returns a tuple with the CreatedAfterFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAfterFilter

`func (o *ObjectRestorePointsFilters) SetCreatedAfterFilter(v time.Time)`

SetCreatedAfterFilter sets CreatedAfterFilter field to given value.

### HasCreatedAfterFilter

`func (o *ObjectRestorePointsFilters) HasCreatedAfterFilter() bool`

HasCreatedAfterFilter returns a boolean if a field has been set.

### GetCreatedBeforeFilter

`func (o *ObjectRestorePointsFilters) GetCreatedBeforeFilter() time.Time`

GetCreatedBeforeFilter returns the CreatedBeforeFilter field if non-nil, zero value otherwise.

### GetCreatedBeforeFilterOk

`func (o *ObjectRestorePointsFilters) GetCreatedBeforeFilterOk() (*time.Time, bool)`

GetCreatedBeforeFilterOk returns a tuple with the CreatedBeforeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBeforeFilter

`func (o *ObjectRestorePointsFilters) SetCreatedBeforeFilter(v time.Time)`

SetCreatedBeforeFilter sets CreatedBeforeFilter field to given value.

### HasCreatedBeforeFilter

`func (o *ObjectRestorePointsFilters) HasCreatedBeforeFilter() bool`

HasCreatedBeforeFilter returns a boolean if a field has been set.

### GetNameFilter

`func (o *ObjectRestorePointsFilters) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *ObjectRestorePointsFilters) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *ObjectRestorePointsFilters) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *ObjectRestorePointsFilters) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetPlatformNameFilter

`func (o *ObjectRestorePointsFilters) GetPlatformNameFilter() EPlatformType`

GetPlatformNameFilter returns the PlatformNameFilter field if non-nil, zero value otherwise.

### GetPlatformNameFilterOk

`func (o *ObjectRestorePointsFilters) GetPlatformNameFilterOk() (*EPlatformType, bool)`

GetPlatformNameFilterOk returns a tuple with the PlatformNameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformNameFilter

`func (o *ObjectRestorePointsFilters) SetPlatformNameFilter(v EPlatformType)`

SetPlatformNameFilter sets PlatformNameFilter field to given value.

### HasPlatformNameFilter

`func (o *ObjectRestorePointsFilters) HasPlatformNameFilter() bool`

HasPlatformNameFilter returns a boolean if a field has been set.

### GetPlatformIdFilter

`func (o *ObjectRestorePointsFilters) GetPlatformIdFilter() string`

GetPlatformIdFilter returns the PlatformIdFilter field if non-nil, zero value otherwise.

### GetPlatformIdFilterOk

`func (o *ObjectRestorePointsFilters) GetPlatformIdFilterOk() (*string, bool)`

GetPlatformIdFilterOk returns a tuple with the PlatformIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIdFilter

`func (o *ObjectRestorePointsFilters) SetPlatformIdFilter(v string)`

SetPlatformIdFilter sets PlatformIdFilter field to given value.

### HasPlatformIdFilter

`func (o *ObjectRestorePointsFilters) HasPlatformIdFilter() bool`

HasPlatformIdFilter returns a boolean if a field has been set.

### GetBackupIdFilter

`func (o *ObjectRestorePointsFilters) GetBackupIdFilter() string`

GetBackupIdFilter returns the BackupIdFilter field if non-nil, zero value otherwise.

### GetBackupIdFilterOk

`func (o *ObjectRestorePointsFilters) GetBackupIdFilterOk() (*string, bool)`

GetBackupIdFilterOk returns a tuple with the BackupIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupIdFilter

`func (o *ObjectRestorePointsFilters) SetBackupIdFilter(v string)`

SetBackupIdFilter sets BackupIdFilter field to given value.

### HasBackupIdFilter

`func (o *ObjectRestorePointsFilters) HasBackupIdFilter() bool`

HasBackupIdFilter returns a boolean if a field has been set.

### GetBackupObjectIdFilter

`func (o *ObjectRestorePointsFilters) GetBackupObjectIdFilter() string`

GetBackupObjectIdFilter returns the BackupObjectIdFilter field if non-nil, zero value otherwise.

### GetBackupObjectIdFilterOk

`func (o *ObjectRestorePointsFilters) GetBackupObjectIdFilterOk() (*string, bool)`

GetBackupObjectIdFilterOk returns a tuple with the BackupObjectIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupObjectIdFilter

`func (o *ObjectRestorePointsFilters) SetBackupObjectIdFilter(v string)`

SetBackupObjectIdFilter sets BackupObjectIdFilter field to given value.

### HasBackupObjectIdFilter

`func (o *ObjectRestorePointsFilters) HasBackupObjectIdFilter() bool`

HasBackupObjectIdFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


