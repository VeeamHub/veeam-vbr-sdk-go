# CloudCredentialsFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** | Number of cloud credentials records to skip. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of cloud credentials records to return. | [optional] 
**OrderColumn** | Pointer to [**ECloudCredentialsFiltersOrderColumn**](ECloudCredentialsFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** | Sorts cloud credentials in the ascending order by the &#x60;orderColumn&#x60; parameter. | [optional] 
**NameFilter** | Pointer to **string** | Filters cloud credentials by the &#x60;nameFilter&#x60; pattern. The pattern can match any cloud credentials parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | [optional] 
**TypeFilter** | Pointer to [**ECloudCredentialsType**](ECloudCredentialsType.md) |  | [optional] 

## Methods

### NewCloudCredentialsFilters

`func NewCloudCredentialsFilters() *CloudCredentialsFilters`

NewCloudCredentialsFilters instantiates a new CloudCredentialsFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsFiltersWithDefaults

`func NewCloudCredentialsFiltersWithDefaults() *CloudCredentialsFilters`

NewCloudCredentialsFiltersWithDefaults instantiates a new CloudCredentialsFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *CloudCredentialsFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *CloudCredentialsFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *CloudCredentialsFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *CloudCredentialsFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *CloudCredentialsFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CloudCredentialsFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CloudCredentialsFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CloudCredentialsFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *CloudCredentialsFilters) GetOrderColumn() ECloudCredentialsFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *CloudCredentialsFilters) GetOrderColumnOk() (*ECloudCredentialsFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *CloudCredentialsFilters) SetOrderColumn(v ECloudCredentialsFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *CloudCredentialsFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *CloudCredentialsFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *CloudCredentialsFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *CloudCredentialsFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *CloudCredentialsFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.

### GetNameFilter

`func (o *CloudCredentialsFilters) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *CloudCredentialsFilters) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *CloudCredentialsFilters) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *CloudCredentialsFilters) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetTypeFilter

`func (o *CloudCredentialsFilters) GetTypeFilter() ECloudCredentialsType`

GetTypeFilter returns the TypeFilter field if non-nil, zero value otherwise.

### GetTypeFilterOk

`func (o *CloudCredentialsFilters) GetTypeFilterOk() (*ECloudCredentialsType, bool)`

GetTypeFilterOk returns a tuple with the TypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFilter

`func (o *CloudCredentialsFilters) SetTypeFilter(v ECloudCredentialsType)`

SetTypeFilter sets TypeFilter field to given value.

### HasTypeFilter

`func (o *CloudCredentialsFilters) HasTypeFilter() bool`

HasTypeFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


