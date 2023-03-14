# NetworkRepositorySettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableTaskLimit** | Pointer to **bool** | If *true*, the maximum number of concurrent tasks is limited. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**EnableReadWriteLimit** | Pointer to **bool** | If *true*, reading and writing speed is limited. | [optional] 
**ReadWriteRate** | Pointer to **int32** | Maximum rate that restricts the total speed of reading and writing data to the backup repository disk. | [optional] 
**AdvancedSettings** | Pointer to [**RepositoryAdvancedSettingsModel**](RepositoryAdvancedSettingsModel.md) |  | [optional] 

## Methods

### NewNetworkRepositorySettingsModel

`func NewNetworkRepositorySettingsModel() *NetworkRepositorySettingsModel`

NewNetworkRepositorySettingsModel instantiates a new NetworkRepositorySettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRepositorySettingsModelWithDefaults

`func NewNetworkRepositorySettingsModelWithDefaults() *NetworkRepositorySettingsModel`

NewNetworkRepositorySettingsModelWithDefaults instantiates a new NetworkRepositorySettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableTaskLimit

`func (o *NetworkRepositorySettingsModel) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *NetworkRepositorySettingsModel) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *NetworkRepositorySettingsModel) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *NetworkRepositorySettingsModel) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *NetworkRepositorySettingsModel) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *NetworkRepositorySettingsModel) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *NetworkRepositorySettingsModel) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *NetworkRepositorySettingsModel) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetEnableReadWriteLimit

`func (o *NetworkRepositorySettingsModel) GetEnableReadWriteLimit() bool`

GetEnableReadWriteLimit returns the EnableReadWriteLimit field if non-nil, zero value otherwise.

### GetEnableReadWriteLimitOk

`func (o *NetworkRepositorySettingsModel) GetEnableReadWriteLimitOk() (*bool, bool)`

GetEnableReadWriteLimitOk returns a tuple with the EnableReadWriteLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableReadWriteLimit

`func (o *NetworkRepositorySettingsModel) SetEnableReadWriteLimit(v bool)`

SetEnableReadWriteLimit sets EnableReadWriteLimit field to given value.

### HasEnableReadWriteLimit

`func (o *NetworkRepositorySettingsModel) HasEnableReadWriteLimit() bool`

HasEnableReadWriteLimit returns a boolean if a field has been set.

### GetReadWriteRate

`func (o *NetworkRepositorySettingsModel) GetReadWriteRate() int32`

GetReadWriteRate returns the ReadWriteRate field if non-nil, zero value otherwise.

### GetReadWriteRateOk

`func (o *NetworkRepositorySettingsModel) GetReadWriteRateOk() (*int32, bool)`

GetReadWriteRateOk returns a tuple with the ReadWriteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadWriteRate

`func (o *NetworkRepositorySettingsModel) SetReadWriteRate(v int32)`

SetReadWriteRate sets ReadWriteRate field to given value.

### HasReadWriteRate

`func (o *NetworkRepositorySettingsModel) HasReadWriteRate() bool`

HasReadWriteRate returns a boolean if a field has been set.

### GetAdvancedSettings

`func (o *NetworkRepositorySettingsModel) GetAdvancedSettings() RepositoryAdvancedSettingsModel`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *NetworkRepositorySettingsModel) GetAdvancedSettingsOk() (*RepositoryAdvancedSettingsModel, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *NetworkRepositorySettingsModel) SetAdvancedSettings(v RepositoryAdvancedSettingsModel)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *NetworkRepositorySettingsModel) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


