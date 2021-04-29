# ManageServerImportSpecCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WindowsHosts** | Pointer to [**[]WindowsHostImportSpec**](WindowsHostImportSpec.md) | Array of managed Microsoft Windows servers. | [optional] 
**LinuxHosts** | Pointer to [**[]LinuxHostImportSpec**](LinuxHostImportSpec.md) | Array of managed Linux servers. | [optional] 
**ViHosts** | Pointer to [**[]ViHostImportSpec**](ViHostImportSpec.md) | Array of VMware vSphere servers. | [optional] 

## Methods

### NewManageServerImportSpecCollection

`func NewManageServerImportSpecCollection() *ManageServerImportSpecCollection`

NewManageServerImportSpecCollection instantiates a new ManageServerImportSpecCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageServerImportSpecCollectionWithDefaults

`func NewManageServerImportSpecCollectionWithDefaults() *ManageServerImportSpecCollection`

NewManageServerImportSpecCollectionWithDefaults instantiates a new ManageServerImportSpecCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindowsHosts

`func (o *ManageServerImportSpecCollection) GetWindowsHosts() []WindowsHostImportSpec`

GetWindowsHosts returns the WindowsHosts field if non-nil, zero value otherwise.

### GetWindowsHostsOk

`func (o *ManageServerImportSpecCollection) GetWindowsHostsOk() (*[]WindowsHostImportSpec, bool)`

GetWindowsHostsOk returns a tuple with the WindowsHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsHosts

`func (o *ManageServerImportSpecCollection) SetWindowsHosts(v []WindowsHostImportSpec)`

SetWindowsHosts sets WindowsHosts field to given value.

### HasWindowsHosts

`func (o *ManageServerImportSpecCollection) HasWindowsHosts() bool`

HasWindowsHosts returns a boolean if a field has been set.

### GetLinuxHosts

`func (o *ManageServerImportSpecCollection) GetLinuxHosts() []LinuxHostImportSpec`

GetLinuxHosts returns the LinuxHosts field if non-nil, zero value otherwise.

### GetLinuxHostsOk

`func (o *ManageServerImportSpecCollection) GetLinuxHostsOk() (*[]LinuxHostImportSpec, bool)`

GetLinuxHostsOk returns a tuple with the LinuxHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxHosts

`func (o *ManageServerImportSpecCollection) SetLinuxHosts(v []LinuxHostImportSpec)`

SetLinuxHosts sets LinuxHosts field to given value.

### HasLinuxHosts

`func (o *ManageServerImportSpecCollection) HasLinuxHosts() bool`

HasLinuxHosts returns a boolean if a field has been set.

### GetViHosts

`func (o *ManageServerImportSpecCollection) GetViHosts() []ViHostImportSpec`

GetViHosts returns the ViHosts field if non-nil, zero value otherwise.

### GetViHostsOk

`func (o *ManageServerImportSpecCollection) GetViHostsOk() (*[]ViHostImportSpec, bool)`

GetViHostsOk returns a tuple with the ViHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViHosts

`func (o *ManageServerImportSpecCollection) SetViHosts(v []ViHostImportSpec)`

SetViHosts sets ViHosts field to given value.

### HasViHosts

`func (o *ManageServerImportSpecCollection) HasViHosts() bool`

HasViHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


