# GuestInteractionProxiesSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSelection** | **bool** | If *true*, Veeam Backup &amp; Replication automatically selects the guest interaction proxy. | [default to true]
**ProxyIds** | Pointer to **[]string** | Array of proxies specified explicitly. The array must contain Microsoft Windows servers added to the backup infrastructure only. | [optional] 

## Methods

### NewGuestInteractionProxiesSettingsModel

`func NewGuestInteractionProxiesSettingsModel(autoSelection bool, ) *GuestInteractionProxiesSettingsModel`

NewGuestInteractionProxiesSettingsModel instantiates a new GuestInteractionProxiesSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestInteractionProxiesSettingsModelWithDefaults

`func NewGuestInteractionProxiesSettingsModelWithDefaults() *GuestInteractionProxiesSettingsModel`

NewGuestInteractionProxiesSettingsModelWithDefaults instantiates a new GuestInteractionProxiesSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSelection

`func (o *GuestInteractionProxiesSettingsModel) GetAutoSelection() bool`

GetAutoSelection returns the AutoSelection field if non-nil, zero value otherwise.

### GetAutoSelectionOk

`func (o *GuestInteractionProxiesSettingsModel) GetAutoSelectionOk() (*bool, bool)`

GetAutoSelectionOk returns a tuple with the AutoSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelection

`func (o *GuestInteractionProxiesSettingsModel) SetAutoSelection(v bool)`

SetAutoSelection sets AutoSelection field to given value.


### GetProxyIds

`func (o *GuestInteractionProxiesSettingsModel) GetProxyIds() []string`

GetProxyIds returns the ProxyIds field if non-nil, zero value otherwise.

### GetProxyIdsOk

`func (o *GuestInteractionProxiesSettingsModel) GetProxyIdsOk() (*[]string, bool)`

GetProxyIdsOk returns a tuple with the ProxyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyIds

`func (o *GuestInteractionProxiesSettingsModel) SetProxyIds(v []string)`

SetProxyIds sets ProxyIds field to given value.

### HasProxyIds

`func (o *GuestInteractionProxiesSettingsModel) HasProxyIds() bool`

HasProxyIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


