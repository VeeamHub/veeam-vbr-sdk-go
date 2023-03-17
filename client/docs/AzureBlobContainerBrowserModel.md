# AzureBlobContainerBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Container name. | [optional] 
**Folders** | Pointer to **[]string** | Array of folders loated in the container. | [optional] 

## Methods

### NewAzureBlobContainerBrowserModel

`func NewAzureBlobContainerBrowserModel() *AzureBlobContainerBrowserModel`

NewAzureBlobContainerBrowserModel instantiates a new AzureBlobContainerBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobContainerBrowserModelWithDefaults

`func NewAzureBlobContainerBrowserModelWithDefaults() *AzureBlobContainerBrowserModel`

NewAzureBlobContainerBrowserModelWithDefaults instantiates a new AzureBlobContainerBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AzureBlobContainerBrowserModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureBlobContainerBrowserModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureBlobContainerBrowserModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureBlobContainerBrowserModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFolders

`func (o *AzureBlobContainerBrowserModel) GetFolders() []string`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *AzureBlobContainerBrowserModel) GetFoldersOk() (*[]string, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *AzureBlobContainerBrowserModel) SetFolders(v []string)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *AzureBlobContainerBrowserModel) HasFolders() bool`

HasFolders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


