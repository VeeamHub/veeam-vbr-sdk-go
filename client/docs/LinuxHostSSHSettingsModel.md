# LinuxHostSSHSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshTimeOutMs** | Pointer to **int32** | SSH timeout, in ms. If a task targeted at the server is inactive after the timeout, the task is terminated. | [optional] 
**PortRangeStart** | Pointer to **int32** | Start port used for data transfer. | [optional] 
**PortRangeEnd** | Pointer to **int32** | End port used for data transfer. | [optional] 
**ServerThisSide** | Pointer to **bool** | If *true*, the server is run on this side. | [optional] 
**ManagementPort** | Pointer to **int32** | Port used as a control channel from the Veeam Backup &amp; Replication console to the Linux server. | [optional] 

## Methods

### NewLinuxHostSSHSettingsModel

`func NewLinuxHostSSHSettingsModel() *LinuxHostSSHSettingsModel`

NewLinuxHostSSHSettingsModel instantiates a new LinuxHostSSHSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinuxHostSSHSettingsModelWithDefaults

`func NewLinuxHostSSHSettingsModelWithDefaults() *LinuxHostSSHSettingsModel`

NewLinuxHostSSHSettingsModelWithDefaults instantiates a new LinuxHostSSHSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshTimeOutMs

`func (o *LinuxHostSSHSettingsModel) GetSshTimeOutMs() int32`

GetSshTimeOutMs returns the SshTimeOutMs field if non-nil, zero value otherwise.

### GetSshTimeOutMsOk

`func (o *LinuxHostSSHSettingsModel) GetSshTimeOutMsOk() (*int32, bool)`

GetSshTimeOutMsOk returns a tuple with the SshTimeOutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshTimeOutMs

`func (o *LinuxHostSSHSettingsModel) SetSshTimeOutMs(v int32)`

SetSshTimeOutMs sets SshTimeOutMs field to given value.

### HasSshTimeOutMs

`func (o *LinuxHostSSHSettingsModel) HasSshTimeOutMs() bool`

HasSshTimeOutMs returns a boolean if a field has been set.

### GetPortRangeStart

`func (o *LinuxHostSSHSettingsModel) GetPortRangeStart() int32`

GetPortRangeStart returns the PortRangeStart field if non-nil, zero value otherwise.

### GetPortRangeStartOk

`func (o *LinuxHostSSHSettingsModel) GetPortRangeStartOk() (*int32, bool)`

GetPortRangeStartOk returns a tuple with the PortRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeStart

`func (o *LinuxHostSSHSettingsModel) SetPortRangeStart(v int32)`

SetPortRangeStart sets PortRangeStart field to given value.

### HasPortRangeStart

`func (o *LinuxHostSSHSettingsModel) HasPortRangeStart() bool`

HasPortRangeStart returns a boolean if a field has been set.

### GetPortRangeEnd

`func (o *LinuxHostSSHSettingsModel) GetPortRangeEnd() int32`

GetPortRangeEnd returns the PortRangeEnd field if non-nil, zero value otherwise.

### GetPortRangeEndOk

`func (o *LinuxHostSSHSettingsModel) GetPortRangeEndOk() (*int32, bool)`

GetPortRangeEndOk returns a tuple with the PortRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRangeEnd

`func (o *LinuxHostSSHSettingsModel) SetPortRangeEnd(v int32)`

SetPortRangeEnd sets PortRangeEnd field to given value.

### HasPortRangeEnd

`func (o *LinuxHostSSHSettingsModel) HasPortRangeEnd() bool`

HasPortRangeEnd returns a boolean if a field has been set.

### GetServerThisSide

`func (o *LinuxHostSSHSettingsModel) GetServerThisSide() bool`

GetServerThisSide returns the ServerThisSide field if non-nil, zero value otherwise.

### GetServerThisSideOk

`func (o *LinuxHostSSHSettingsModel) GetServerThisSideOk() (*bool, bool)`

GetServerThisSideOk returns a tuple with the ServerThisSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerThisSide

`func (o *LinuxHostSSHSettingsModel) SetServerThisSide(v bool)`

SetServerThisSide sets ServerThisSide field to given value.

### HasServerThisSide

`func (o *LinuxHostSSHSettingsModel) HasServerThisSide() bool`

HasServerThisSide returns a boolean if a field has been set.

### GetManagementPort

`func (o *LinuxHostSSHSettingsModel) GetManagementPort() int32`

GetManagementPort returns the ManagementPort field if non-nil, zero value otherwise.

### GetManagementPortOk

`func (o *LinuxHostSSHSettingsModel) GetManagementPortOk() (*int32, bool)`

GetManagementPortOk returns a tuple with the ManagementPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementPort

`func (o *LinuxHostSSHSettingsModel) SetManagementPort(v int32)`

SetManagementPort sets ManagementPort field to given value.

### HasManagementPort

`func (o *LinuxHostSSHSettingsModel) HasManagementPort() bool`

HasManagementPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


