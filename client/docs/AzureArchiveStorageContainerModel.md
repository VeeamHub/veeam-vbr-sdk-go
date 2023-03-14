# AzureArchiveStorageContainerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerName** | **string** | Name of an Azure Archive container. | 
**FolderName** | **string** | Name of the folder that the object storage repository is mapped to. | 

## Methods

### NewAzureArchiveStorageContainerModel

`func NewAzureArchiveStorageContainerModel(containerName string, folderName string, ) *AzureArchiveStorageContainerModel`

NewAzureArchiveStorageContainerModel instantiates a new AzureArchiveStorageContainerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureArchiveStorageContainerModelWithDefaults

`func NewAzureArchiveStorageContainerModelWithDefaults() *AzureArchiveStorageContainerModel`

NewAzureArchiveStorageContainerModelWithDefaults instantiates a new AzureArchiveStorageContainerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerName

`func (o *AzureArchiveStorageContainerModel) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *AzureArchiveStorageContainerModel) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *AzureArchiveStorageContainerModel) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetFolderName

`func (o *AzureArchiveStorageContainerModel) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *AzureArchiveStorageContainerModel) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *AzureArchiveStorageContainerModel) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


