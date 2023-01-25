# VCenterInventoryFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skip** | Pointer to **int32** | Number of objects to skip. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of objects to return. | [optional] 
**OrderColumn** | Pointer to [**EvCentersInventoryFiltersOrderColumn**](EvCentersInventoryFiltersOrderColumn.md) |  | [optional] 
**OrderAsc** | Pointer to **bool** | Sorts objects in the ascending order by the &#x60;orderColumn&#x60; parameter. | [optional] 
**ObjectIdFilter** | Pointer to **string** | Filters objects by object ID. | [optional] 
**HierarchyTypeFilter** | Pointer to [**EHierarchyType**](EHierarchyType.md) |  | [optional] 
**NameFilter** | Pointer to **string** | Filters objects by the &#x60;nameFilter&#x60; pattern. The pattern can match any object parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both. | [optional] 
**TypeFilter** | Pointer to [**EVmwareInventoryType**](EVmwareInventoryType.md) |  | [optional] 
**ParentContainerNameFilter** | Pointer to **string** | Filters objects by name of the parent container. | [optional] 

## Methods

### NewVCenterInventoryFilters

`func NewVCenterInventoryFilters() *VCenterInventoryFilters`

NewVCenterInventoryFilters instantiates a new VCenterInventoryFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVCenterInventoryFiltersWithDefaults

`func NewVCenterInventoryFiltersWithDefaults() *VCenterInventoryFilters`

NewVCenterInventoryFiltersWithDefaults instantiates a new VCenterInventoryFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkip

`func (o *VCenterInventoryFilters) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *VCenterInventoryFilters) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *VCenterInventoryFilters) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *VCenterInventoryFilters) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetLimit

`func (o *VCenterInventoryFilters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *VCenterInventoryFilters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *VCenterInventoryFilters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *VCenterInventoryFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderColumn

`func (o *VCenterInventoryFilters) GetOrderColumn() EvCentersInventoryFiltersOrderColumn`

GetOrderColumn returns the OrderColumn field if non-nil, zero value otherwise.

### GetOrderColumnOk

`func (o *VCenterInventoryFilters) GetOrderColumnOk() (*EvCentersInventoryFiltersOrderColumn, bool)`

GetOrderColumnOk returns a tuple with the OrderColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderColumn

`func (o *VCenterInventoryFilters) SetOrderColumn(v EvCentersInventoryFiltersOrderColumn)`

SetOrderColumn sets OrderColumn field to given value.

### HasOrderColumn

`func (o *VCenterInventoryFilters) HasOrderColumn() bool`

HasOrderColumn returns a boolean if a field has been set.

### GetOrderAsc

`func (o *VCenterInventoryFilters) GetOrderAsc() bool`

GetOrderAsc returns the OrderAsc field if non-nil, zero value otherwise.

### GetOrderAscOk

`func (o *VCenterInventoryFilters) GetOrderAscOk() (*bool, bool)`

GetOrderAscOk returns a tuple with the OrderAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAsc

`func (o *VCenterInventoryFilters) SetOrderAsc(v bool)`

SetOrderAsc sets OrderAsc field to given value.

### HasOrderAsc

`func (o *VCenterInventoryFilters) HasOrderAsc() bool`

HasOrderAsc returns a boolean if a field has been set.

### GetObjectIdFilter

`func (o *VCenterInventoryFilters) GetObjectIdFilter() string`

GetObjectIdFilter returns the ObjectIdFilter field if non-nil, zero value otherwise.

### GetObjectIdFilterOk

`func (o *VCenterInventoryFilters) GetObjectIdFilterOk() (*string, bool)`

GetObjectIdFilterOk returns a tuple with the ObjectIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIdFilter

`func (o *VCenterInventoryFilters) SetObjectIdFilter(v string)`

SetObjectIdFilter sets ObjectIdFilter field to given value.

### HasObjectIdFilter

`func (o *VCenterInventoryFilters) HasObjectIdFilter() bool`

HasObjectIdFilter returns a boolean if a field has been set.

### GetHierarchyTypeFilter

`func (o *VCenterInventoryFilters) GetHierarchyTypeFilter() EHierarchyType`

GetHierarchyTypeFilter returns the HierarchyTypeFilter field if non-nil, zero value otherwise.

### GetHierarchyTypeFilterOk

`func (o *VCenterInventoryFilters) GetHierarchyTypeFilterOk() (*EHierarchyType, bool)`

GetHierarchyTypeFilterOk returns a tuple with the HierarchyTypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyTypeFilter

`func (o *VCenterInventoryFilters) SetHierarchyTypeFilter(v EHierarchyType)`

SetHierarchyTypeFilter sets HierarchyTypeFilter field to given value.

### HasHierarchyTypeFilter

`func (o *VCenterInventoryFilters) HasHierarchyTypeFilter() bool`

HasHierarchyTypeFilter returns a boolean if a field has been set.

### GetNameFilter

`func (o *VCenterInventoryFilters) GetNameFilter() string`

GetNameFilter returns the NameFilter field if non-nil, zero value otherwise.

### GetNameFilterOk

`func (o *VCenterInventoryFilters) GetNameFilterOk() (*string, bool)`

GetNameFilterOk returns a tuple with the NameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFilter

`func (o *VCenterInventoryFilters) SetNameFilter(v string)`

SetNameFilter sets NameFilter field to given value.

### HasNameFilter

`func (o *VCenterInventoryFilters) HasNameFilter() bool`

HasNameFilter returns a boolean if a field has been set.

### GetTypeFilter

`func (o *VCenterInventoryFilters) GetTypeFilter() EVmwareInventoryType`

GetTypeFilter returns the TypeFilter field if non-nil, zero value otherwise.

### GetTypeFilterOk

`func (o *VCenterInventoryFilters) GetTypeFilterOk() (*EVmwareInventoryType, bool)`

GetTypeFilterOk returns a tuple with the TypeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeFilter

`func (o *VCenterInventoryFilters) SetTypeFilter(v EVmwareInventoryType)`

SetTypeFilter sets TypeFilter field to given value.

### HasTypeFilter

`func (o *VCenterInventoryFilters) HasTypeFilter() bool`

HasTypeFilter returns a boolean if a field has been set.

### GetParentContainerNameFilter

`func (o *VCenterInventoryFilters) GetParentContainerNameFilter() string`

GetParentContainerNameFilter returns the ParentContainerNameFilter field if non-nil, zero value otherwise.

### GetParentContainerNameFilterOk

`func (o *VCenterInventoryFilters) GetParentContainerNameFilterOk() (*string, bool)`

GetParentContainerNameFilterOk returns a tuple with the ParentContainerNameFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentContainerNameFilter

`func (o *VCenterInventoryFilters) SetParentContainerNameFilter(v string)`

SetParentContainerNameFilter sets ParentContainerNameFilter field to given value.

### HasParentContainerNameFilter

`func (o *VCenterInventoryFilters) HasParentContainerNameFilter() bool`

HasParentContainerNameFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


