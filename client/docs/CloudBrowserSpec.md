# CloudBrowserSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialsId** | **string** | ID of the object storage account (for browsing either storage or compute infrastructure). | 
**ServiceType** | [**ECloudServiceType**](ECloudServiceType.md) |  | 
**RegionType** | [**EAmazonRegionType**](EAmazonRegionType.md) |  | 
**FolderType** | Pointer to [**ECloudBrowserFolderType**](ECloudBrowserFolderType.md) |  | [optional] 
**GatewayServerId** | Pointer to **string** | ID of a gateway server you want to use to connect to the object storage. Specify this parameter to check internet connection of the server. As a gateway server you can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 
**Filters** | Pointer to [**AmazonEC2BrowserFilters**](AmazonEC2BrowserFilters.md) |  | [optional] 
**RegionId** | **string** | Region ID. | 
**ServicePoint** | **string** | Endpoint address and port number of the IBM Cloud object storage. | 

## Methods

### NewCloudBrowserSpec

`func NewCloudBrowserSpec(credentialsId string, serviceType ECloudServiceType, regionType EAmazonRegionType, regionId string, servicePoint string, ) *CloudBrowserSpec`

NewCloudBrowserSpec instantiates a new CloudBrowserSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBrowserSpecWithDefaults

`func NewCloudBrowserSpecWithDefaults() *CloudBrowserSpec`

NewCloudBrowserSpecWithDefaults instantiates a new CloudBrowserSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialsId

`func (o *CloudBrowserSpec) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *CloudBrowserSpec) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *CloudBrowserSpec) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetServiceType

`func (o *CloudBrowserSpec) GetServiceType() ECloudServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CloudBrowserSpec) GetServiceTypeOk() (*ECloudServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CloudBrowserSpec) SetServiceType(v ECloudServiceType)`

SetServiceType sets ServiceType field to given value.


### GetRegionType

`func (o *CloudBrowserSpec) GetRegionType() EAmazonRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *CloudBrowserSpec) GetRegionTypeOk() (*EAmazonRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *CloudBrowserSpec) SetRegionType(v EAmazonRegionType)`

SetRegionType sets RegionType field to given value.


### GetFolderType

`func (o *CloudBrowserSpec) GetFolderType() ECloudBrowserFolderType`

GetFolderType returns the FolderType field if non-nil, zero value otherwise.

### GetFolderTypeOk

`func (o *CloudBrowserSpec) GetFolderTypeOk() (*ECloudBrowserFolderType, bool)`

GetFolderTypeOk returns a tuple with the FolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderType

`func (o *CloudBrowserSpec) SetFolderType(v ECloudBrowserFolderType)`

SetFolderType sets FolderType field to given value.

### HasFolderType

`func (o *CloudBrowserSpec) HasFolderType() bool`

HasFolderType returns a boolean if a field has been set.

### GetGatewayServerId

`func (o *CloudBrowserSpec) GetGatewayServerId() string`

GetGatewayServerId returns the GatewayServerId field if non-nil, zero value otherwise.

### GetGatewayServerIdOk

`func (o *CloudBrowserSpec) GetGatewayServerIdOk() (*string, bool)`

GetGatewayServerIdOk returns a tuple with the GatewayServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServerId

`func (o *CloudBrowserSpec) SetGatewayServerId(v string)`

SetGatewayServerId sets GatewayServerId field to given value.

### HasGatewayServerId

`func (o *CloudBrowserSpec) HasGatewayServerId() bool`

HasGatewayServerId returns a boolean if a field has been set.

### GetFilters

`func (o *CloudBrowserSpec) GetFilters() AmazonEC2BrowserFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CloudBrowserSpec) GetFiltersOk() (*AmazonEC2BrowserFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CloudBrowserSpec) SetFilters(v AmazonEC2BrowserFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *CloudBrowserSpec) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetRegionId

`func (o *CloudBrowserSpec) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *CloudBrowserSpec) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *CloudBrowserSpec) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetServicePoint

`func (o *CloudBrowserSpec) GetServicePoint() string`

GetServicePoint returns the ServicePoint field if non-nil, zero value otherwise.

### GetServicePointOk

`func (o *CloudBrowserSpec) GetServicePointOk() (*string, bool)`

GetServicePointOk returns a tuple with the ServicePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePoint

`func (o *CloudBrowserSpec) SetServicePoint(v string)`

SetServicePoint sets ServicePoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


