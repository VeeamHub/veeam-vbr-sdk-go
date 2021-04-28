# EncryptionPasswordsFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**OrderColumn** | Pointer to [**EEncryptionPasswordsFiltersOrderColumn**](EEncryptionPasswordsFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** |  | [optional] 

## Methods

### NewEncryptionPasswordsFilters

`func NewEncryptionPasswordsFilters() *EncryptionPasswordsFilters`

NewEncryptionPasswordsFilters instantiates a new EncryptionPasswordsFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionPasswordsFiltersWithDefaults

`func NewEncryptionPasswordsFiltersWithDefaults() *EncryptionPasswordsFilters`

NewEncryptionPasswordsFiltersWithDefaults instantiates a new EncryptionPasswordsFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *EncryptionPasswordsFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *EncryptionPasswordsFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *EncryptionPasswordsFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *EncryptionPasswordsFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *EncryptionPasswordsFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *EncryptionPasswordsFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *EncryptionPasswordsFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *EncryptionPasswordsFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *EncryptionPasswordsFilters) GetOrderColumn() EEncryptionPasswordsFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *EncryptionPasswordsFilters) GetOrderColumnOk() (*EEncryptionPasswordsFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *EncryptionPasswordsFilters) SetOrderColumn(v EEncryptionPasswordsFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *EncryptionPasswordsFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *EncryptionPasswordsFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *EncryptionPasswordsFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *EncryptionPasswordsFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *EncryptionPasswordsFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


