# AzureDataBoxBrowserDestinationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | ID of a server you want to use to connect to the object storage. You can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 
**ContainerName** | **string** | Name of the container where you want to store your backup data. | 
**ServicePoint** | **string** | Service endpoint address of the Azure Data Box device. | 

## Methods

### NewAzureDataBoxBrowserDestinationSpec

`func NewAzureDataBoxBrowserDestinationSpec(containerName string, servicePoint string, ) *AzureDataBoxBrowserDestinationSpec`

NewAzureDataBoxBrowserDestinationSpec instantiates a new AzureDataBoxBrowserDestinationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDataBoxBrowserDestinationSpecWithDefaults

`func NewAzureDataBoxBrowserDestinationSpecWithDefaults() *AzureDataBoxBrowserDestinationSpec`

NewAzureDataBoxBrowserDestinationSpecWithDefaults instantiates a new AzureDataBoxBrowserDestinationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *AzureDataBoxBrowserDestinationSpec) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *AzureDataBoxBrowserDestinationSpec) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *AzureDataBoxBrowserDestinationSpec) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *AzureDataBoxBrowserDestinationSpec) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetContainerName

`func (o *AzureDataBoxBrowserDestinationSpec) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *AzureDataBoxBrowserDestinationSpec) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *AzureDataBoxBrowserDestinationSpec) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetServicePoint

`func (o *AzureDataBoxBrowserDestinationSpec) GetServicePoint() string`

GetServicePoint returns the ServicePoint field if non-nil, zero value otherwise.

### GetServicePointOk

`func (o *AzureDataBoxBrowserDestinationSpec) GetServicePointOk() (*string, bool)`

GetServicePointOk returns a tuple with the ServicePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePoint

`func (o *AzureDataBoxBrowserDestinationSpec) SetServicePoint(v string)`

SetServicePoint sets ServicePoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


