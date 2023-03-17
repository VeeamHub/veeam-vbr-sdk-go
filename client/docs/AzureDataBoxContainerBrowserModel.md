# AzureDataBoxContainerBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Container name. | [optional] 
**Folders** | Pointer to **[]string** | Array of folders located in the container. | [optional] 

## Methods

### NewAzureDataBoxContainerBrowserModel

`func NewAzureDataBoxContainerBrowserModel() *AzureDataBoxContainerBrowserModel`

NewAzureDataBoxContainerBrowserModel instantiates a new AzureDataBoxContainerBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDataBoxContainerBrowserModelWithDefaults

`func NewAzureDataBoxContainerBrowserModelWithDefaults() *AzureDataBoxContainerBrowserModel`

NewAzureDataBoxContainerBrowserModelWithDefaults instantiates a new AzureDataBoxContainerBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AzureDataBoxContainerBrowserModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureDataBoxContainerBrowserModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureDataBoxContainerBrowserModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureDataBoxContainerBrowserModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFolders

`func (o *AzureDataBoxContainerBrowserModel) GetFolders() []string`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *AzureDataBoxContainerBrowserModel) GetFoldersOk() (*[]string, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *AzureDataBoxContainerBrowserModel) SetFolders(v []string)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *AzureDataBoxContainerBrowserModel) HasFolders() bool`

HasFolders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


