# WindowsHostPortsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]WindowsHostComponentPortModel**](WindowsHostComponentPortModel.md) | Array of Veeam Backup &amp; Replication components. | [optional] 
**PortRangeStart** | Pointer to **int32** | Start port used for data transfer. | [optional] 
**PortRangeEnd** | Pointer to **int32** | End port used for data transfer. | [optional] 
**ServerThisSide** | Pointer to **bool** | If *true*, the server is run on this side. | [optional] 

## Methods

### NewWindowsHostPortsModel

`func NewWindowsHostPortsModel() *WindowsHostPortsModel`

NewWindowsHostPortsModel instantiates a new WindowsHostPortsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsHostPortsModelWithDefaults

`func NewWindowsHostPortsModelWithDefaults() *WindowsHostPortsModel`

NewWindowsHostPortsModelWithDefaults instantiates a new WindowsHostPortsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *WindowsHostPortsModel) GetComponents() []WindowsHostComponentPortModel`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *WindowsHostPortsModel) GetComponentsOk() (*[]WindowsHostComponentPortModel, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *WindowsHostPortsModel) SetComponents(v []WindowsHostComponentPortModel)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *WindowsHostPortsModel) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetPortRangeStart

`func (o *WindowsHostPortsModel) GetPortRangeStart() int32`

GetPortRangeStart returns the PortRangeStart field if non-nil, zero value otherwise.

### GetPortRangeStartOk

`func (o *WindowsHostPortsModel) GetPortRangeStartOk() (*int32, bool)`

GetPortRangeStartOk returns a tuple with the PortRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeStart

`func (o *WindowsHostPortsModel) SetPortRangeStart(v int32)`

SetPortRangeStart sets PortRangeStart field to given value.

### HasPortRangeStart

`func (o *WindowsHostPortsModel) HasPortRangeStart() bool`

HasPortRangeStart returns a boolean if a field has been set.

### GetPortRangeEnd

`func (o *WindowsHostPortsModel) GetPortRangeEnd() int32`

GetPortRangeEnd returns the PortRangeEnd field if non-nil, zero value otherwise.

### GetPortRangeEndOk

`func (o *WindowsHostPortsModel) GetPortRangeEndOk() (*int32, bool)`

GetPortRangeEndOk returns a tuple with the PortRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeEnd

`func (o *WindowsHostPortsModel) SetPortRangeEnd(v int32)`

SetPortRangeEnd sets PortRangeEnd field to given value.

### HasPortRangeEnd

`func (o *WindowsHostPortsModel) HasPortRangeEnd() bool`

HasPortRangeEnd returns a boolean if a field has been set.

### GetServerThisSide

`func (o *WindowsHostPortsModel) GetServerThisSide() bool`

GetServerThisSide returns the ServerThisSide field if non-nil, zero value otherwise.

### GetServerThisSideOk

`func (o *WindowsHostPortsModel) GetServerThisSideOk() (*bool, bool)`

GetServerThisSideOk returns a tuple with the ServerThisSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerThisSide

`func (o *WindowsHostPortsModel) SetServerThisSide(v bool)`

SetServerThisSide sets ServerThisSide field to given value.

### HasServerThisSide

`func (o *WindowsHostPortsModel) HasServerThisSide() bool`

HasServerThisSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


