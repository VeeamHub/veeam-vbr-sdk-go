# WindowsHostSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkSettings** | Pointer to [**WindowsHostPortsModel**](WindowsHostPortsModel.md) |  | [optional] 

## Methods

### NewWindowsHostSpec

`func NewWindowsHostSpec() *WindowsHostSpec`

NewWindowsHostSpec instantiates a new WindowsHostSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsHostSpecWithDefaults

`func NewWindowsHostSpecWithDefaults() *WindowsHostSpec`

NewWindowsHostSpecWithDefaults instantiates a new WindowsHostSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkSettings

`func (o *WindowsHostSpec) GetNetworkSettings() WindowsHostPortsModel`

GetNetworkSettings returns the NetworkSettings field if non-nil, zero value otherwise.

### GetNetworkSettingsOk

`func (o *WindowsHostSpec) GetNetworkSettingsOk() (*WindowsHostPortsModel, bool)`

GetNetworkSettingsOk returns a tuple with the NetworkSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSettings

`func (o *WindowsHostSpec) SetNetworkSettings(v WindowsHostPortsModel)`

SetNetworkSettings sets NetworkSettings field to given value.

### HasNetworkSettings

`func (o *WindowsHostSpec) HasNetworkSettings() bool`

HasNetworkSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


