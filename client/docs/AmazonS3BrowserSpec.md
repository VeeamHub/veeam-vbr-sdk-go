# AmazonS3BrowserSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionType** | [**EAmazonRegionType**](EAmazonRegionType.md) |  | 
**GatewayServerId** | Pointer to **string** | ID of a gateway server you want to use to connect to the object storage. Specify this parameter to check internet connection of the server. As a gateway server you can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 
**FolderType** | Pointer to [**ECloudBrowserFolderType**](ECloudBrowserFolderType.md) |  | [optional] 
**Filters** | Pointer to [**AmazonS3CloudBrowserFilters**](AmazonS3CloudBrowserFilters.md) |  | [optional] 

## Methods

### NewAmazonS3BrowserSpec

`func NewAmazonS3BrowserSpec(regionType EAmazonRegionType, ) *AmazonS3BrowserSpec`

NewAmazonS3BrowserSpec instantiates a new AmazonS3BrowserSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3BrowserSpecWithDefaults

`func NewAmazonS3BrowserSpecWithDefaults() *AmazonS3BrowserSpec`

NewAmazonS3BrowserSpecWithDefaults instantiates a new AmazonS3BrowserSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionType

`func (o *AmazonS3BrowserSpec) GetRegionType() EAmazonRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AmazonS3BrowserSpec) GetRegionTypeOk() (*EAmazonRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AmazonS3BrowserSpec) SetRegionType(v EAmazonRegionType)`

SetRegionType sets RegionType field to given value.


### GetGatewayServerId

`func (o *AmazonS3BrowserSpec) GetGatewayServerId() string`

GetGatewayServerId returns the GatewayServerId field if non-nil, zero value otherwise.

### GetGatewayServerIdOk

`func (o *AmazonS3BrowserSpec) GetGatewayServerIdOk() (*string, bool)`

GetGatewayServerIdOk returns a tuple with the GatewayServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServerId

`func (o *AmazonS3BrowserSpec) SetGatewayServerId(v string)`

SetGatewayServerId sets GatewayServerId field to given value.

### HasGatewayServerId

`func (o *AmazonS3BrowserSpec) HasGatewayServerId() bool`

HasGatewayServerId returns a boolean if a field has been set.

### GetFolderType

`func (o *AmazonS3BrowserSpec) GetFolderType() ECloudBrowserFolderType`

GetFolderType returns the FolderType field if non-nil, zero value otherwise.

### GetFolderTypeOk

`func (o *AmazonS3BrowserSpec) GetFolderTypeOk() (*ECloudBrowserFolderType, bool)`

GetFolderTypeOk returns a tuple with the FolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderType

`func (o *AmazonS3BrowserSpec) SetFolderType(v ECloudBrowserFolderType)`

SetFolderType sets FolderType field to given value.

### HasFolderType

`func (o *AmazonS3BrowserSpec) HasFolderType() bool`

HasFolderType returns a boolean if a field has been set.

### GetFilters

`func (o *AmazonS3BrowserSpec) GetFilters() AmazonS3CloudBrowserFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AmazonS3BrowserSpec) GetFiltersOk() (*AmazonS3CloudBrowserFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AmazonS3BrowserSpec) SetFilters(v AmazonS3CloudBrowserFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AmazonS3BrowserSpec) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


