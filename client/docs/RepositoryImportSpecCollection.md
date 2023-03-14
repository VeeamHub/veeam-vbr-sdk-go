# RepositoryImportSpecCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WindowsLocalRepositories** | Pointer to [**[]WindowsLocalStorageImportSpec**](WindowsLocalStorageImportSpec.md) | Array of Microsoft Windows-based repositories. | [optional] 
**LinuxLocalRepositories** | Pointer to [**[]LinuxLocalStorageImportSpec**](LinuxLocalStorageImportSpec.md) | Array of Linux-based repositories. | [optional] 
**SmbRepositories** | Pointer to [**[]SmbStorageImportSpec**](SmbStorageImportSpec.md) | Array of SMB backup repositories. | [optional] 
**NfsRepositories** | Pointer to [**[]NfsStorageImportSpec**](NfsStorageImportSpec.md) | Array of NFS backup repositories. | [optional] 
**AzureBlobStorages** | Pointer to [**[]AzureBlobStorageImportSpec**](AzureBlobStorageImportSpec.md) | Array of Microsoft Azure Blob storages. | [optional] 
**AzureDataBoxStorages** | Pointer to [**[]AzureDataBoxStorageImportSpec**](AzureDataBoxStorageImportSpec.md) | Array of Microsoft Azure Data Box storages. | [optional] 
**AmazonS3Storages** | Pointer to [**[]AmazonS3StorageImportSpec**](AmazonS3StorageImportSpec.md) | Array of Amazon S3 storages. | [optional] 
**AmazonSnowballEdgeStorages** | Pointer to [**[]AmazonSnowballEdgeStorageImportSpec**](AmazonSnowballEdgeStorageImportSpec.md) | Array of AWS Snowball Edge storages. | [optional] 
**S3CompatibleStorages** | Pointer to [**[]S3CompatibleStorageImportSpec**](S3CompatibleStorageImportSpec.md) | Array of S3 compatible storages. | [optional] 
**GoogleCloudStorages** | Pointer to [**[]GoogleCloudStorageImportSpec**](GoogleCloudStorageImportSpec.md) | Array of Google Cloud storages. | [optional] 
**IBMCloudStorages** | Pointer to [**[]IBMCloudStorageImportSpec**](IBMCloudStorageImportSpec.md) | Array of IBM Cloud storages. | [optional] 
**AmazonS3GlacierStorages** | Pointer to [**[]AmazonS3GlacierStorageImportSpec**](AmazonS3GlacierStorageImportSpec.md) | Array of Amazon S3 Glacier storages. | [optional] 
**AzureArchiveStorages** | Pointer to [**[]AzureArchiveStorageImportSpec**](AzureArchiveStorageImportSpec.md) | Array of Microsoft Azure Archive storages. | [optional] 
**WasabiCloudStorages** | Pointer to [**[]WasabiCloudStorageImportSpec**](WasabiCloudStorageImportSpec.md) | Array of Wasabi Cloud storages. | [optional] 
**LinuxHardenedRepositories** | Pointer to [**[]LinuxHardenedStorageImportSpec**](LinuxHardenedStorageImportSpec.md) | Array of Linux hardened repositories. | [optional] 

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

### GetAzureBlobStorages

`func (o *RepositoryImportSpecCollection) GetAzureBlobStorages() []AzureBlobStorageImportSpec`

GetAzureBlobStorages returns the AzureBlobStorages field if non-nil, zero value otherwise.

### GetAzureBlobStoragesOk

`func (o *RepositoryImportSpecCollection) GetAzureBlobStoragesOk() (*[]AzureBlobStorageImportSpec, bool)`

GetAzureBlobStoragesOk returns a tuple with the AzureBlobStorages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureBlobStorages

`func (o *RepositoryImportSpecCollection) SetAzureBlobStorages(v []AzureBlobStorageImportSpec)`

SetAzureBlobStorages sets AzureBlobStorages field to given value.

### HasAzureBlobStorages

`func (o *RepositoryImportSpecCollection) HasAzureBlobStorages() bool`

HasAzureBlobStorages returns a boolean if a field has been set.

### GetAzureDataBoxStorages

`func (o *RepositoryImportSpecCollection) GetAzureDataBoxStorages() []AzureDataBoxStorageImportSpec`

GetAzureDataBoxStorages returns the AzureDataBoxStorages field if non-nil, zero value otherwise.

### GetAzureDataBoxStoragesOk

`func (o *RepositoryImportSpecCollection) GetAzureDataBoxStoragesOk() (*[]AzureDataBoxStorageImportSpec, bool)`

GetAzureDataBoxStoragesOk returns a tuple with the AzureDataBoxStorages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureDataBoxStorages

`func (o *RepositoryImportSpecCollection) SetAzureDataBoxStorages(v []AzureDataBoxStorageImportSpec)`

SetAzureDataBoxStorages sets AzureDataBoxStorages field to given value.

### HasAzureDataBoxStorages

`func (o *RepositoryImportSpecCollection) HasAzureDataBoxStorages() bool`

HasAzureDataBoxStorages returns a boolean if a field has been set.

### GetAmazonS3Storages

`func (o *RepositoryImportSpecCollection) GetAmazonS3Storages() []AmazonS3StorageImportSpec`

GetAmazonS3Storages returns the AmazonS3Storages field if non-nil, zero value otherwise.

### GetAmazonS3StoragesOk

`func (o *RepositoryImportSpecCollection) GetAmazonS3StoragesOk() (*[]AmazonS3StorageImportSpec, bool)`

GetAmazonS3StoragesOk returns a tuple with the AmazonS3Storages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonS3Storages

`func (o *RepositoryImportSpecCollection) SetAmazonS3Storages(v []AmazonS3StorageImportSpec)`

SetAmazonS3Storages sets AmazonS3Storages field to given value.

### HasAmazonS3Storages

`func (o *RepositoryImportSpecCollection) HasAmazonS3Storages() bool`

HasAmazonS3Storages returns a boolean if a field has been set.

### GetAmazonSnowballEdgeStorages

`func (o *RepositoryImportSpecCollection) GetAmazonSnowballEdgeStorages() []AmazonSnowballEdgeStorageImportSpec`

GetAmazonSnowballEdgeStorages returns the AmazonSnowballEdgeStorages field if non-nil, zero value otherwise.

### GetAmazonSnowballEdgeStoragesOk

`func (o *RepositoryImportSpecCollection) GetAmazonSnowballEdgeStoragesOk() (*[]AmazonSnowballEdgeStorageImportSpec, bool)`

GetAmazonSnowballEdgeStoragesOk returns a tuple with the AmazonSnowballEdgeStorages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonSnowballEdgeStorages

`func (o *RepositoryImportSpecCollection) SetAmazonSnowballEdgeStorages(v []AmazonSnowballEdgeStorageImportSpec)`

SetAmazonSnowballEdgeStorages sets AmazonSnowballEdgeStorages field to given value.

### HasAmazonSnowballEdgeStorages

`func (o *RepositoryImportSpecCollection) HasAmazonSnowballEdgeStorages() bool`

HasAmazonSnowballEdgeStorages returns a boolean if a field has been set.

### GetS3CompatibleStorages

`func (o *RepositoryImportSpecCollection) GetS3CompatibleStorages() []S3CompatibleStorageImportSpec`

GetS3CompatibleStorages returns the S3CompatibleStorages field if non-nil, zero value otherwise.

### GetS3CompatibleStoragesOk

`func (o *RepositoryImportSpecCollection) GetS3CompatibleStoragesOk() (*[]S3CompatibleStorageImportSpec, bool)`

GetS3CompatibleStoragesOk returns a tuple with the S3CompatibleStorages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3CompatibleStorages

`func (o *RepositoryImportSpecCollection) SetS3CompatibleStorages(v []S3CompatibleStorageImportSpec)`

SetS3CompatibleStorages sets S3CompatibleStorages field to given value.

### HasS3CompatibleStorages

`func (o *RepositoryImportSpecCollection) HasS3CompatibleStorages() bool`

HasS3CompatibleStorages returns a boolean if a field has been set.

### GetGoogleCloudStorages

`func (o *RepositoryImportSpecCollection) GetGoogleCloudStorages() []GoogleCloudStorageImportSpec`

GetGoogleCloudStorages returns the GoogleCloudStorages field if non-nil, zero value otherwise.

### GetGoogleCloudStoragesOk

`func (o *RepositoryImportSpecCollection) GetGoogleCloudStoragesOk() (*[]GoogleCloudStorageImportSpec, bool)`

GetGoogleCloudStoragesOk returns a tuple with the GoogleCloudStorages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleCloudStorages

`func (o *RepositoryImportSpecCollection) SetGoogleCloudStorages(v []GoogleCloudStorageImportSpec)`

SetGoogleCloudStorages sets GoogleCloudStorages field to given value.

### HasGoogleCloudStorages

`func (o *RepositoryImportSpecCollection) HasGoogleCloudStorages() bool`

HasGoogleCloudStorages returns a boolean if a field has been set.

### GetIBMCloudStorages

`func (o *RepositoryImportSpecCollection) GetIBMCloudStorages() []IBMCloudStorageImportSpec`

GetIBMCloudStorages returns the IBMCloudStorages field if non-nil, zero value otherwise.

### GetIBMCloudStoragesOk

`func (o *RepositoryImportSpecCollection) GetIBMCloudStoragesOk() (*[]IBMCloudStorageImportSpec, bool)`

GetIBMCloudStoragesOk returns a tuple with the IBMCloudStorages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIBMCloudStorages

`func (o *RepositoryImportSpecCollection) SetIBMCloudStorages(v []IBMCloudStorageImportSpec)`

SetIBMCloudStorages sets IBMCloudStorages field to given value.

### HasIBMCloudStorages

`func (o *RepositoryImportSpecCollection) HasIBMCloudStorages() bool`

HasIBMCloudStorages returns a boolean if a field has been set.

### GetAmazonS3GlacierStorages

`func (o *RepositoryImportSpecCollection) GetAmazonS3GlacierStorages() []AmazonS3GlacierStorageImportSpec`

GetAmazonS3GlacierStorages returns the AmazonS3GlacierStorages field if non-nil, zero value otherwise.

### GetAmazonS3GlacierStoragesOk

`func (o *RepositoryImportSpecCollection) GetAmazonS3GlacierStoragesOk() (*[]AmazonS3GlacierStorageImportSpec, bool)`

GetAmazonS3GlacierStoragesOk returns a tuple with the AmazonS3GlacierStorages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonS3GlacierStorages

`func (o *RepositoryImportSpecCollection) SetAmazonS3GlacierStorages(v []AmazonS3GlacierStorageImportSpec)`

SetAmazonS3GlacierStorages sets AmazonS3GlacierStorages field to given value.

### HasAmazonS3GlacierStorages

`func (o *RepositoryImportSpecCollection) HasAmazonS3GlacierStorages() bool`

HasAmazonS3GlacierStorages returns a boolean if a field has been set.

### GetAzureArchiveStorages

`func (o *RepositoryImportSpecCollection) GetAzureArchiveStorages() []AzureArchiveStorageImportSpec`

GetAzureArchiveStorages returns the AzureArchiveStorages field if non-nil, zero value otherwise.

### GetAzureArchiveStoragesOk

`func (o *RepositoryImportSpecCollection) GetAzureArchiveStoragesOk() (*[]AzureArchiveStorageImportSpec, bool)`

GetAzureArchiveStoragesOk returns a tuple with the AzureArchiveStorages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureArchiveStorages

`func (o *RepositoryImportSpecCollection) SetAzureArchiveStorages(v []AzureArchiveStorageImportSpec)`

SetAzureArchiveStorages sets AzureArchiveStorages field to given value.

### HasAzureArchiveStorages

`func (o *RepositoryImportSpecCollection) HasAzureArchiveStorages() bool`

HasAzureArchiveStorages returns a boolean if a field has been set.

### GetWasabiCloudStorages

`func (o *RepositoryImportSpecCollection) GetWasabiCloudStorages() []WasabiCloudStorageImportSpec`

GetWasabiCloudStorages returns the WasabiCloudStorages field if non-nil, zero value otherwise.

### GetWasabiCloudStoragesOk

`func (o *RepositoryImportSpecCollection) GetWasabiCloudStoragesOk() (*[]WasabiCloudStorageImportSpec, bool)`

GetWasabiCloudStoragesOk returns a tuple with the WasabiCloudStorages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasabiCloudStorages

`func (o *RepositoryImportSpecCollection) SetWasabiCloudStorages(v []WasabiCloudStorageImportSpec)`

SetWasabiCloudStorages sets WasabiCloudStorages field to given value.

### HasWasabiCloudStorages

`func (o *RepositoryImportSpecCollection) HasWasabiCloudStorages() bool`

HasWasabiCloudStorages returns a boolean if a field has been set.

### GetLinuxHardenedRepositories

`func (o *RepositoryImportSpecCollection) GetLinuxHardenedRepositories() []LinuxHardenedStorageImportSpec`

GetLinuxHardenedRepositories returns the LinuxHardenedRepositories field if non-nil, zero value otherwise.

### GetLinuxHardenedRepositoriesOk

`func (o *RepositoryImportSpecCollection) GetLinuxHardenedRepositoriesOk() (*[]LinuxHardenedStorageImportSpec, bool)`

GetLinuxHardenedRepositoriesOk returns a tuple with the LinuxHardenedRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxHardenedRepositories

`func (o *RepositoryImportSpecCollection) SetLinuxHardenedRepositories(v []LinuxHardenedStorageImportSpec)`

SetLinuxHardenedRepositories sets LinuxHardenedRepositories field to given value.

### HasLinuxHardenedRepositories

`func (o *RepositoryImportSpecCollection) HasLinuxHardenedRepositories() bool`

HasLinuxHardenedRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


