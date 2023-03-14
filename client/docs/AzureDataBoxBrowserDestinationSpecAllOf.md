# AzureDataBoxBrowserDestinationSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | ID of a server you want to use to connect to the object storage. You can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 
**ContainerName** | **string** | Name of the container where you want to store your backup data. | 
**ServicePoint** | **string** | Service endpoint address of the Azure Data Box device. | 

## Methods

### NewAzureDataBoxBrowserDestinationSpecAllOf

`func NewAzureDataBoxBrowserDestinationSpecAllOf(containerName string, servicePoint string, ) *AzureDataBoxBrowserDestinationSpecAllOf`

NewAzureDataBoxBrowserDestinationSpecAllOf instantiates a new AzureDataBoxBrowserDestinationSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDataBoxBrowserDestinationSpecAllOfWithDefaults

`func NewAzureDataBoxBrowserDestinationSpecAllOfWithDefaults() *AzureDataBoxBrowserDestinationSpecAllOf`

NewAzureDataBoxBrowserDestinationSpecAllOfWithDefaults instantiates a new AzureDataBoxBrowserDestinationSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *AzureDataBoxBrowserDestinationSpecAllOf) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *AzureDataBoxBrowserDestinationSpecAllOf) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *AzureDataBoxBrowserDestinationSpecAllOf) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *AzureDataBoxBrowserDestinationSpecAllOf) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetContainerName

`func (o *AzureDataBoxBrowserDestinationSpecAllOf) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *AzureDataBoxBrowserDestinationSpecAllOf) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *AzureDataBoxBrowserDestinationSpecAllOf) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetServicePoint

`func (o *AzureDataBoxBrowserDestinationSpecAllOf) GetServicePoint() string`

GetServicePoint returns the ServicePoint field if non-nil, zero value otherwise.

### GetServicePointOk

`func (o *AzureDataBoxBrowserDestinationSpecAllOf) GetServicePointOk() (*string, bool)`

GetServicePointOk returns a tuple with the ServicePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePoint

`func (o *AzureDataBoxBrowserDestinationSpecAllOf) SetServicePoint(v string)`

SetServicePoint sets ServicePoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


