# SessionsFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**OrderColumn** | Pointer to [**ESessionsFiltersOrderColumn**](ESessionsFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** |  | [optional] 
**NameFilter** | Pointer to **string** |  | [optional] 
**CreatedAfterFilter** | Pointer to **time.Time** |  | [optional] 
**CreatedBeforeFilter** | Pointer to **time.Time** |  | [optional] 
**EndedAfterFilter** | Pointer to **time.Time** |  | [optional] 
**EndedBeforeFilter** | Pointer to **time.Time** |  | [optional] 
**TypeFilter** | Pointer to [**ESessionType**](ESessionType.md) |  | [optional] 
**StateFilter** | Pointer to [**ESessionState**](ESessionState.md) |  | [optional] 
**ResultFilter** | Pointer to [**ESessionResult**](ESessionResult.md) |  | [optional] 
**JobIdFilter** | Pointer to **string** |  | [optional] 

## Methods

### NewSessionsFilters

`func NewSessionsFilters() *SessionsFilters`

NewSessionsFilters instantiates a new SessionsFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionsFiltersWithDefaults

`func NewSessionsFiltersWithDefaults() *SessionsFilters`

NewSessionsFiltersWithDefaults instantiates a new SessionsFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *SessionsFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *SessionsFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *SessionsFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *SessionsFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *SessionsFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SessionsFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SessionsFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SessionsFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *SessionsFilters) GetOrderColumn() ESessionsFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *SessionsFilters) GetOrderColumnOk() (*ESessionsFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *SessionsFilters) SetOrderColumn(v ESessionsFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *SessionsFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *SessionsFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *SessionsFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *SessionsFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *SessionsFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.

### GetNameFilter

`func (o *SessionsFilters) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *SessionsFilters) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *SessionsFilters) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *SessionsFilters) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetCreatedAfterFilter

`func (o *SessionsFilters) GetCreatedAfterFilter() time.Time`

GetCreatedAfterFilter returns the CreatedAfterFilter field if non-nil, zero value otherwise.

### GetCreatedAfterFilterOk

`func (o *SessionsFilters) GetCreatedAfterFilterOk() (*time.Time, bool)`

GetCreatedAfterFilterOk returns a tuple with the CreatedAfterFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAfterFilter

`func (o *SessionsFilters) SetCreatedAfterFilter(v time.Time)`

SetCreatedAfterFilter sets CreatedAfterFilter field to given value.

### HasCreatedAfterFilter

`func (o *SessionsFilters) HasCreatedAfterFilter() bool`

HasCreatedAfterFilter returns a boolean if a field has been set.

### GetCreatedBeforeFilter

`func (o *SessionsFilters) GetCreatedBeforeFilter() time.Time`

GetCreatedBeforeFilter returns the CreatedBeforeFilter field if non-nil, zero value otherwise.

### GetCreatedBeforeFilterOk

`func (o *SessionsFilters) GetCreatedBeforeFilterOk() (*time.Time, bool)`

GetCreatedBeforeFilterOk returns a tuple with the CreatedBeforeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBeforeFilter

`func (o *SessionsFilters) SetCreatedBeforeFilter(v time.Time)`

SetCreatedBeforeFilter sets CreatedBeforeFilter field to given value.

### HasCreatedBeforeFilter

`func (o *SessionsFilters) HasCreatedBeforeFilter() bool`

HasCreatedBeforeFilter returns a boolean if a field has been set.

### GetEndedAfterFilter

`func (o *SessionsFilters) GetEndedAfterFilter() time.Time`

GetEndedAfterFilter returns the EndedAfterFilter field if non-nil, zero value otherwise.

### GetEndedAfterFilterOk

`func (o *SessionsFilters) GetEndedAfterFilterOk() (*time.Time, bool)`

GetEndedAfterFilterOk returns a tuple with the EndedAfterFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAfterFilter

`func (o *SessionsFilters) SetEndedAfterFilter(v time.Time)`

SetEndedAfterFilter sets EndedAfterFilter field to given value.

### HasEndedAfterFilter

`func (o *SessionsFilters) HasEndedAfterFilter() bool`

HasEndedAfterFilter returns a boolean if a field has been set.

### GetEndedBeforeFilter

`func (o *SessionsFilters) GetEndedBeforeFilter() time.Time`

GetEndedBeforeFilter returns the EndedBeforeFilter field if non-nil, zero value otherwise.

### GetEndedBeforeFilterOk

`func (o *SessionsFilters) GetEndedBeforeFilterOk() (*time.Time, bool)`

GetEndedBeforeFilterOk returns a tuple with the EndedBeforeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedBeforeFilter

`func (o *SessionsFilters) SetEndedBeforeFilter(v time.Time)`

SetEndedBeforeFilter sets EndedBeforeFilter field to given value.

### HasEndedBeforeFilter

`func (o *SessionsFilters) HasEndedBeforeFilter() bool`

HasEndedBeforeFilter returns a boolean if a field has been set.

### GetTypeFilter

`func (o *SessionsFilters) GetTypeFilter() ESessionType`

GetTypeFilter returns the TypeFilter field if non-nil, zero value otherwise.

### GetTypeFilterOk

`func (o *SessionsFilters) GetTypeFilterOk() (*ESessionType, bool)`

GetTypeFilterOk returns a tuple with the TypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFilter

`func (o *SessionsFilters) SetTypeFilter(v ESessionType)`

SetTypeFilter sets TypeFilter field to given value.

### HasTypeFilter

`func (o *SessionsFilters) HasTypeFilter() bool`

HasTypeFilter returns a boolean if a field has been set.

### GetStateFilter

`func (o *SessionsFilters) GetStateFilter() ESessionState`

GetStateFilter returns the StateFilter field if non-nil, zero value otherwise.

### GetStateFilterOk

`func (o *SessionsFilters) GetStateFilterOk() (*ESessionState, bool)`

GetStateFilterOk returns a tuple with the StateFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateFilter

`func (o *SessionsFilters) SetStateFilter(v ESessionState)`

SetStateFilter sets StateFilter field to given value.

### HasStateFilter

`func (o *SessionsFilters) HasStateFilter() bool`

HasStateFilter returns a boolean if a field has been set.

### GetResultFilter

`func (o *SessionsFilters) GetResultFilter() ESessionResult`

GetResultFilter returns the ResultFilter field if non-nil, zero value otherwise.

### GetResultFilterOk

`func (o *SessionsFilters) GetResultFilterOk() (*ESessionResult, bool)`

GetResultFilterOk returns a tuple with the ResultFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultFilter

`func (o *SessionsFilters) SetResultFilter(v ESessionResult)`

SetResultFilter sets ResultFilter field to given value.

### HasResultFilter

`func (o *SessionsFilters) HasResultFilter() bool`

HasResultFilter returns a boolean if a field has been set.

### GetJobIdFilter

`func (o *SessionsFilters) GetJobIdFilter() string`

GetJobIdFilter returns the JobIdFilter field if non-nil, zero value otherwise.

### GetJobIdFilterOk

`func (o *SessionsFilters) GetJobIdFilterOk() (*string, bool)`

GetJobIdFilterOk returns a tuple with the JobIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIdFilter

`func (o *SessionsFilters) SetJobIdFilter(v string)`

SetJobIdFilter sets JobIdFilter field to given value.

### HasJobIdFilter

`func (o *SessionsFilters) HasJobIdFilter() bool`

HasJobIdFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


