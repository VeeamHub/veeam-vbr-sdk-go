# ObjectStorageImmutabilityModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | If *true*, storage immutability is enabled. | [optional] 
**DaysCount** | Pointer to **int32** | Immutability period. | [optional] 

## Methods

### NewObjectStorageImmutabilityModel

`func NewObjectStorageImmutabilityModel() *ObjectStorageImmutabilityModel`

NewObjectStorageImmutabilityModel instantiates a new ObjectStorageImmutabilityModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageImmutabilityModelWithDefaults

`func NewObjectStorageImmutabilityModelWithDefaults() *ObjectStorageImmutabilityModel`

NewObjectStorageImmutabilityModelWithDefaults instantiates a new ObjectStorageImmutabilityModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ObjectStorageImmutabilityModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ObjectStorageImmutabilityModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ObjectStorageImmutabilityModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ObjectStorageImmutabilityModel) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetDaysCount

`func (o *ObjectStorageImmutabilityModel) GetDaysCount() int32`

GetDaysCount returns the DaysCount field if non-nil, zero value otherwise.

### GetDaysCountOk

`func (o *ObjectStorageImmutabilityModel) GetDaysCountOk() (*int32, bool)`

GetDaysCountOk returns a tuple with the DaysCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysCount

`func (o *ObjectStorageImmutabilityModel) SetDaysCount(v int32)`

SetDaysCount sets DaysCount field to given value.

### HasDaysCount

`func (o *ObjectStorageImmutabilityModel) HasDaysCount() bool`

HasDaysCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


