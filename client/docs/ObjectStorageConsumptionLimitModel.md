# ObjectStorageConsumptionLimitModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | If *true*, the consumption limit is enabled. | [optional] 
**ConsumptionLimitCount** | Pointer to **int32** | Limit value. | [optional] 
**ConsumptionLimitKind** | Pointer to [**EConsumptionLimitKind**](EConsumptionLimitKind.md) |  | [optional] 

## Methods

### NewObjectStorageConsumptionLimitModel

`func NewObjectStorageConsumptionLimitModel() *ObjectStorageConsumptionLimitModel`

NewObjectStorageConsumptionLimitModel instantiates a new ObjectStorageConsumptionLimitModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageConsumptionLimitModelWithDefaults

`func NewObjectStorageConsumptionLimitModelWithDefaults() *ObjectStorageConsumptionLimitModel`

NewObjectStorageConsumptionLimitModelWithDefaults instantiates a new ObjectStorageConsumptionLimitModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ObjectStorageConsumptionLimitModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ObjectStorageConsumptionLimitModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ObjectStorageConsumptionLimitModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ObjectStorageConsumptionLimitModel) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetConsumptionLimitCount

`func (o *ObjectStorageConsumptionLimitModel) GetConsumptionLimitCount() int32`

GetConsumptionLimitCount returns the ConsumptionLimitCount field if non-nil, zero value otherwise.

### GetConsumptionLimitCountOk

`func (o *ObjectStorageConsumptionLimitModel) GetConsumptionLimitCountOk() (*int32, bool)`

GetConsumptionLimitCountOk returns a tuple with the ConsumptionLimitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumptionLimitCount

`func (o *ObjectStorageConsumptionLimitModel) SetConsumptionLimitCount(v int32)`

SetConsumptionLimitCount sets ConsumptionLimitCount field to given value.

### HasConsumptionLimitCount

`func (o *ObjectStorageConsumptionLimitModel) HasConsumptionLimitCount() bool`

HasConsumptionLimitCount returns a boolean if a field has been set.

### GetConsumptionLimitKind

`func (o *ObjectStorageConsumptionLimitModel) GetConsumptionLimitKind() EConsumptionLimitKind`

GetConsumptionLimitKind returns the ConsumptionLimitKind field if non-nil, zero value otherwise.

### GetConsumptionLimitKindOk

`func (o *ObjectStorageConsumptionLimitModel) GetConsumptionLimitKindOk() (*EConsumptionLimitKind, bool)`

GetConsumptionLimitKindOk returns a tuple with the ConsumptionLimitKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumptionLimitKind

`func (o *ObjectStorageConsumptionLimitModel) SetConsumptionLimitKind(v EConsumptionLimitKind)`

SetConsumptionLimitKind sets ConsumptionLimitKind field to given value.

### HasConsumptionLimitKind

`func (o *ObjectStorageConsumptionLimitModel) HasConsumptionLimitKind() bool`

HasConsumptionLimitKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


