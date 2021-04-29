# RepositoryImportSpecCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WindowsLocalRepositories** | Pointer to [**[]WindowsLocalStorageImportSpec**](WindowsLocalStorageImportSpec.md) | Array of windows local storages. | [optional] 
**LinuxLocalRepositories** | Pointer to [**[]LinuxLocalStorageImportSpec**](LinuxLocalStorageImportSpec.md) | Array of linux local storages. | [optional] 
**SmbRepositories** | Pointer to [**[]SmbStorageImportSpec**](SmbStorageImportSpec.md) | Array of network attached storages. | [optional] 
**NfsRepositories** | Pointer to [**[]NfsStorageImportSpec**](NfsStorageImportSpec.md) | Array of network attached storages. | [optional] 

## Methods

### NewRepositoryImportSpecCollection

`func NewRepositoryImportSpecCollection() *RepositoryImportSpecCollection`

NewRepositoryImportSpecCollection instantiates a new RepositoryImportSpecCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryImportSpecCollectionWithDefaults

`func NewRepositoryImportSpecCollectionWithDefaults() *RepositoryImportSpecCollection`

NewRepositoryImportSpecCollectionWithDefaults instantiates a new RepositoryImportSpecCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindowsLocalRepositories

`func (o *RepositoryImportSpecCollection) GetWindowsLocalRepositories() []WindowsLocalStorageImportSpec`

GetWindowsLocalRepositories returns the WindowsLocalRepositories field if non-nil, zero value otherwise.

### GetWindowsLocalRepositoriesOk

`func (o *RepositoryImportSpecCollection) GetWindowsLocalRepositoriesOk() (*[]WindowsLocalStorageImportSpec, bool)`

GetWindowsLocalRepositoriesOk returns a tuple with the WindowsLocalRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsLocalRepositories

`func (o *RepositoryImportSpecCollection) SetWindowsLocalRepositories(v []WindowsLocalStorageImportSpec)`

SetWindowsLocalRepositories sets WindowsLocalRepositories field to given value.

### HasWindowsLocalRepositories

`func (o *RepositoryImportSpecCollection) HasWindowsLocalRepositories() bool`

HasWindowsLocalRepositories returns a boolean if a field has been set.

### GetLinuxLocalRepositories

`func (o *RepositoryImportSpecCollection) GetLinuxLocalRepositories() []LinuxLocalStorageImportSpec`

GetLinuxLocalRepositories returns the LinuxLocalRepositories field if non-nil, zero value otherwise.

### GetLinuxLocalRepositoriesOk

`func (o *RepositoryImportSpecCollection) GetLinuxLocalRepositoriesOk() (*[]LinuxLocalStorageImportSpec, bool)`

GetLinuxLocalRepositoriesOk returns a tuple with the LinuxLocalRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxLocalRepositories

`func (o *RepositoryImportSpecCollection) SetLinuxLocalRepositories(v []LinuxLocalStorageImportSpec)`

SetLinuxLocalRepositories sets LinuxLocalRepositories field to given value.

### HasLinuxLocalRepositories

`func (o *RepositoryImportSpecCollection) HasLinuxLocalRepositories() bool`

HasLinuxLocalRepositories returns a boolean if a field has been set.

### GetSmbRepositories

`func (o *RepositoryImportSpecCollection) GetSmbRepositories() []SmbStorageImportSpec`

GetSmbRepositories returns the SmbRepositories field if non-nil, zero value otherwise.

### GetSmbRepositoriesOk

`func (o *RepositoryImportSpecCollection) GetSmbRepositoriesOk() (*[]SmbStorageImportSpec, bool)`

GetSmbRepositoriesOk returns a tuple with the SmbRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbRepositories

`func (o *RepositoryImportSpecCollection) SetSmbRepositories(v []SmbStorageImportSpec)`

SetSmbRepositories sets SmbRepositories field to given value.

### HasSmbRepositories

`func (o *RepositoryImportSpecCollection) HasSmbRepositories() bool`

HasSmbRepositories returns a boolean if a field has been set.

### GetNfsRepositories

`func (o *RepositoryImportSpecCollection) GetNfsRepositories() []NfsStorageImportSpec`

GetNfsRepositories returns the NfsRepositories field if non-nil, zero value otherwise.

### GetNfsRepositoriesOk

`func (o *RepositoryImportSpecCollection) GetNfsRepositoriesOk() (*[]NfsStorageImportSpec, bool)`

GetNfsRepositoriesOk returns a tuple with the NfsRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsRepositories

`func (o *RepositoryImportSpecCollection) SetNfsRepositories(v []NfsStorageImportSpec)`

SetNfsRepositories sets NfsRepositories field to given value.

### HasNfsRepositories

`func (o *RepositoryImportSpecCollection) HasNfsRepositories() bool`

HasNfsRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


