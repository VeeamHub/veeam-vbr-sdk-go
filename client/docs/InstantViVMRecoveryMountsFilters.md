# InstantViVMRecoveryMountsFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** | Number of mounts to skip. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of mounts to return. | [optional] 
**OrderColumn** | Pointer to [**EInstantViVMRecoveryMountsFiltersOrderColumn**](EInstantViVMRecoveryMountsFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** | Sorts mounts in the ascending order by the &#x60;orderColumn&#x60; parameter. | [optional] 
**NameFilter** | Pointer to **string** | Filters mounts by the &#x60;nameFilter&#x60; pattern. The pattern can match any mount parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both. | [optional] 
**StateFilter** | Pointer to [**EInstantRecoveryMountState**](EInstantRecoveryMountState.md) |  | [optional] 

## Methods

### NewInstantViVMRecoveryMountsFilters

`func NewInstantViVMRecoveryMountsFilters() *InstantViVMRecoveryMountsFilters`

NewInstantViVMRecoveryMountsFilters instantiates a new InstantViVMRecoveryMountsFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstantViVMRecoveryMountsFiltersWithDefaults

`func NewInstantViVMRecoveryMountsFiltersWithDefaults() *InstantViVMRecoveryMountsFilters`

NewInstantViVMRecoveryMountsFiltersWithDefaults instantiates a new InstantViVMRecoveryMountsFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *InstantViVMRecoveryMountsFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *InstantViVMRecoveryMountsFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *InstantViVMRecoveryMountsFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *InstantViVMRecoveryMountsFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *InstantViVMRecoveryMountsFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InstantViVMRecoveryMountsFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InstantViVMRecoveryMountsFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *InstantViVMRecoveryMountsFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *InstantViVMRecoveryMountsFilters) GetOrderColumn() EInstantViVMRecoveryMountsFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *InstantViVMRecoveryMountsFilters) GetOrderColumnOk() (*EInstantViVMRecoveryMountsFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *InstantViVMRecoveryMountsFilters) SetOrderColumn(v EInstantViVMRecoveryMountsFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *InstantViVMRecoveryMountsFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *InstantViVMRecoveryMountsFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *InstantViVMRecoveryMountsFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *InstantViVMRecoveryMountsFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *InstantViVMRecoveryMountsFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.

### GetNameFilter

`func (o *InstantViVMRecoveryMountsFilters) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *InstantViVMRecoveryMountsFilters) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *InstantViVMRecoveryMountsFilters) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *InstantViVMRecoveryMountsFilters) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetStateFilter

`func (o *InstantViVMRecoveryMountsFilters) GetStateFilter() EInstantRecoveryMountState`

GetStateFilter returns the StateFilter field if non-nil, zero value otherwise.

### GetStateFilterOk

`func (o *InstantViVMRecoveryMountsFilters) GetStateFilterOk() (*EInstantRecoveryMountState, bool)`

GetStateFilterOk returns a tuple with the StateFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateFilter

`func (o *InstantViVMRecoveryMountsFilters) SetStateFilter(v EInstantRecoveryMountState)`

SetStateFilter sets StateFilter field to given value.

### HasStateFilter

`func (o *InstantViVMRecoveryMountsFilters) HasStateFilter() bool`

HasStateFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


