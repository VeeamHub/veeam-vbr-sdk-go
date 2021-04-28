# WindowsHostComponentPortModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentName** | [**EWindowsHostComponentType**](EWindowsHostComponentType.md) |  | 
**Port** | **int32** | Port used by the component. | 

## Methods

### NewWindowsHostComponentPortModel

`func NewWindowsHostComponentPortModel(componentName EWindowsHostComponentType, port int32, ) *WindowsHostComponentPortModel`

NewWindowsHostComponentPortModel instantiates a new WindowsHostComponentPortModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsHostComponentPortModelWithDefaults

`func NewWindowsHostComponentPortModelWithDefaults() *WindowsHostComponentPortModel`

NewWindowsHostComponentPortModelWithDefaults instantiates a new WindowsHostComponentPortModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentName

`func (o *WindowsHostComponentPortModel) GetComponentName() EWindowsHostComponentType`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *WindowsHostComponentPortModel) GetComponentNameOk() (*EWindowsHostComponentType, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *WindowsHostComponentPortModel) SetComponentName(v EWindowsHostComponentType)`

SetComponentName sets ComponentName field to given value.


### GetPort

`func (o *WindowsHostComponentPortModel) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *WindowsHostComponentPortModel) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *WindowsHostComponentPortModel) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


