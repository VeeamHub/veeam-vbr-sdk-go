# GFSPolicySettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, the long-term (GFS) retention policy is enabled. | 
**Weekly** | Pointer to [**GFSPolicySettingsWeeklyModel**](GFSPolicySettingsWeeklyModel.md) |  | [optional] 
**Monthly** | Pointer to [**GFSPolicySettingsMonthlyModel**](GFSPolicySettingsMonthlyModel.md) |  | [optional] 
**Yearly** | Pointer to [**GFSPolicySettingsYearlyModel**](GFSPolicySettingsYearlyModel.md) |  | [optional] 

## Methods

### NewGFSPolicySettingsModel

`func NewGFSPolicySettingsModel(isEnabled bool, ) *GFSPolicySettingsModel`

NewGFSPolicySettingsModel instantiates a new GFSPolicySettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGFSPolicySettingsModelWithDefaults

`func NewGFSPolicySettingsModelWithDefaults() *GFSPolicySettingsModel`

NewGFSPolicySettingsModelWithDefaults instantiates a new GFSPolicySettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *GFSPolicySettingsModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *GFSPolicySettingsModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *GFSPolicySettingsModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetWeekly

`func (o *GFSPolicySettingsModel) GetWeekly() GFSPolicySettingsWeeklyModel`

GetWeekly returns the Weekly field if non-nil, zero value otherwise.

### GetWeeklyOk

`func (o *GFSPolicySettingsModel) GetWeeklyOk() (*GFSPolicySettingsWeeklyModel, bool)`

GetWeeklyOk returns a tuple with the Weekly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekly

`func (o *GFSPolicySettingsModel) SetWeekly(v GFSPolicySettingsWeeklyModel)`

SetWeekly sets Weekly field to given value.

### HasWeekly

`func (o *GFSPolicySettingsModel) HasWeekly() bool`

HasWeekly returns a boolean if a field has been set.

### GetMonthly

`func (o *GFSPolicySettingsModel) GetMonthly() GFSPolicySettingsMonthlyModel`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *GFSPolicySettingsModel) GetMonthlyOk() (*GFSPolicySettingsMonthlyModel, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *GFSPolicySettingsModel) SetMonthly(v GFSPolicySettingsMonthlyModel)`

SetMonthly sets Monthly field to given value.

### HasMonthly

`func (o *GFSPolicySettingsModel) HasMonthly() bool`

HasMonthly returns a boolean if a field has been set.

### GetYearly

`func (o *GFSPolicySettingsModel) GetYearly() GFSPolicySettingsYearlyModel`

GetYearly returns the Yearly field if non-nil, zero value otherwise.

### GetYearlyOk

`func (o *GFSPolicySettingsModel) GetYearlyOk() (*GFSPolicySettingsYearlyModel, bool)`

GetYearlyOk returns a tuple with the Yearly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearly

`func (o *GFSPolicySettingsModel) SetYearly(v GFSPolicySettingsYearlyModel)`

SetYearly sets Yearly field to given value.

### HasYearly

`func (o *GFSPolicySettingsModel) HasYearly() bool`

HasYearly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


