# GuestInteractionProxiesSettingsImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticSelection** | **bool** | If *true*, Veeam Backup &amp; Replication automatically selects the guest interaction proxy. | [default to true]
**Proxies** | Pointer to **[]string** | Array of proxies specified explicitly. The array must contain Microsoft Windows servers added to the backup infrastructure only. | [optional] 

## Methods

### NewGuestInteractionProxiesSettingsImportModel

`func NewGuestInteractionProxiesSettingsImportModel(automaticSelection bool, ) *GuestInteractionProxiesSettingsImportModel`

NewGuestInteractionProxiesSettingsImportModel instantiates a new GuestInteractionProxiesSettingsImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestInteractionProxiesSettingsImportModelWithDefaults

`func NewGuestInteractionProxiesSettingsImportModelWithDefaults() *GuestInteractionProxiesSettingsImportModel`

NewGuestInteractionProxiesSettingsImportModelWithDefaults instantiates a new GuestInteractionProxiesSettingsImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticSelection

`func (o *GuestInteractionProxiesSettingsImportModel) GetAutomaticSelection() bool`

GetAutomaticSelection returns the AutomaticSelection field if non-nil, zero value otherwise.

### GetAutomaticSelectionOk

`func (o *GuestInteractionProxiesSettingsImportModel) GetAutomaticSelectionOk() (*bool, bool)`

GetAutomaticSelectionOk returns a tuple with the AutomaticSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticSelection

`func (o *GuestInteractionProxiesSettingsImportModel) SetAutomaticSelection(v bool)`

SetAutomaticSelection sets AutomaticSelection field to given value.


### GetProxies

`func (o *GuestInteractionProxiesSettingsImportModel) GetProxies() []string`

GetProxies returns the Proxies field if non-nil, zero value otherwise.

### GetProxiesOk

`func (o *GuestInteractionProxiesSettingsImportModel) GetProxiesOk() (*[]string, bool)`

GetProxiesOk returns a tuple with the Proxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxies

`func (o *GuestInteractionProxiesSettingsImportModel) SetProxies(v []string)`

SetProxies sets Proxies field to given value.

### HasProxies

`func (o *GuestInteractionProxiesSettingsImportModel) HasProxies() bool`

HasProxies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


