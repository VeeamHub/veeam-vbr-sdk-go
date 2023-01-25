# CloudBrowserNewFolderSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialsId** | **string** | ID of a cloud credentials record requiered to connect to the object storage. | 
**ServiceType** | [**ECloudServiceType**](ECloudServiceType.md) |  | 
**NewFolderName** | **string** | Name of the new folder. | 
**HostId** | Pointer to **string** | ID of a server you want to use to connect to the object storage. You can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 
**RegionType** | [**EAmazonRegionType**](EAmazonRegionType.md) |  | 
**ContainerName** | **string** | Name of the container where you want to store your backup data. | 
**FolderType** | Pointer to [**ECloudBrowserFolderType**](ECloudBrowserFolderType.md) |  | [optional] 
**RegionId** | **string** | Region where the bucket is located. | 
**BucketName** | **string** | Name of the bucket where you want to store your backup data. | 
**ConnectionPoint** | **string** | Endpoint address and port number of the IBM Cloud object storage. | 
**ServicePoint** | **string** | Service endpoint address of the Azure Data Box device. | 

## Methods

### NewCloudBrowserNewFolderSpec

`func NewCloudBrowserNewFolderSpec(credentialsId string, serviceType ECloudServiceType, newFolderName string, regionType EAmazonRegionType, containerName string, regionId string, bucketName string, connectionPoint string, servicePoint string, ) *CloudBrowserNewFolderSpec`

NewCloudBrowserNewFolderSpec instantiates a new CloudBrowserNewFolderSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBrowserNewFolderSpecWithDefaults

`func NewCloudBrowserNewFolderSpecWithDefaults() *CloudBrowserNewFolderSpec`

NewCloudBrowserNewFolderSpecWithDefaults instantiates a new CloudBrowserNewFolderSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialsId

`func (o *CloudBrowserNewFolderSpec) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *CloudBrowserNewFolderSpec) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *CloudBrowserNewFolderSpec) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetServiceType

`func (o *CloudBrowserNewFolderSpec) GetServiceType() ECloudServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CloudBrowserNewFolderSpec) GetServiceTypeOk() (*ECloudServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CloudBrowserNewFolderSpec) SetServiceType(v ECloudServiceType)`

SetServiceType sets ServiceType field to given value.


### GetNewFolderName

`func (o *CloudBrowserNewFolderSpec) GetNewFolderName() string`

GetNewFolderName returns the NewFolderName field if non-nil, zero value otherwise.

### GetNewFolderNameOk

`func (o *CloudBrowserNewFolderSpec) GetNewFolderNameOk() (*string, bool)`

GetNewFolderNameOk returns a tuple with the NewFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFolderName

`func (o *CloudBrowserNewFolderSpec) SetNewFolderName(v string)`

SetNewFolderName sets NewFolderName field to given value.


### GetHostId

`func (o *CloudBrowserNewFolderSpec) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *CloudBrowserNewFolderSpec) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *CloudBrowserNewFolderSpec) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *CloudBrowserNewFolderSpec) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetRegionType

`func (o *CloudBrowserNewFolderSpec) GetRegionType() EAmazonRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *CloudBrowserNewFolderSpec) GetRegionTypeOk() (*EAmazonRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *CloudBrowserNewFolderSpec) SetRegionType(v EAmazonRegionType)`

SetRegionType sets RegionType field to given value.


### GetContainerName

`func (o *CloudBrowserNewFolderSpec) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *CloudBrowserNewFolderSpec) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *CloudBrowserNewFolderSpec) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetFolderType

`func (o *CloudBrowserNewFolderSpec) GetFolderType() ECloudBrowserFolderType`

GetFolderType returns the FolderType field if non-nil, zero value otherwise.

### GetFolderTypeOk

`func (o *CloudBrowserNewFolderSpec) GetFolderTypeOk() (*ECloudBrowserFolderType, bool)`

GetFolderTypeOk returns a tuple with the FolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderType

`func (o *CloudBrowserNewFolderSpec) SetFolderType(v ECloudBrowserFolderType)`

SetFolderType sets FolderType field to given value.

### HasFolderType

`func (o *CloudBrowserNewFolderSpec) HasFolderType() bool`

HasFolderType returns a boolean if a field has been set.

### GetRegionId

`func (o *CloudBrowserNewFolderSpec) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *CloudBrowserNewFolderSpec) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *CloudBrowserNewFolderSpec) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetBucketName

`func (o *CloudBrowserNewFolderSpec) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *CloudBrowserNewFolderSpec) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *CloudBrowserNewFolderSpec) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetConnectionPoint

`func (o *CloudBrowserNewFolderSpec) GetConnectionPoint() string`

GetConnectionPoint returns the ConnectionPoint field if non-nil, zero value otherwise.

### GetConnectionPointOk

`func (o *CloudBrowserNewFolderSpec) GetConnectionPointOk() (*string, bool)`

GetConnectionPointOk returns a tuple with the ConnectionPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPoint

`func (o *CloudBrowserNewFolderSpec) SetConnectionPoint(v string)`

SetConnectionPoint sets ConnectionPoint field to given value.


### GetServicePoint

`func (o *CloudBrowserNewFolderSpec) GetServicePoint() string`

GetServicePoint returns the ServicePoint field if non-nil, zero value otherwise.

### GetServicePointOk

`func (o *CloudBrowserNewFolderSpec) GetServicePointOk() (*string, bool)`

GetServicePointOk returns a tuple with the ServicePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePoint

`func (o *CloudBrowserNewFolderSpec) SetServicePoint(v string)`

SetServicePoint sets ServicePoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


