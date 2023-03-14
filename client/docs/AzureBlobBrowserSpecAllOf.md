# AzureBlobBrowserSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionType** | Pointer to [**EAzureRegionType**](EAzureRegionType.md) |  | [optional] 
**FolderType** | Pointer to [**ECloudBrowserFolderType**](ECloudBrowserFolderType.md) |  | [optional] 
**GatewayServerId** | Pointer to **string** | ID of a gateway server you want to use to connect to the object storage. Specify this parameter to check internet connection of the server. As a gateway server you can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 

## Methods

### NewAzureBlobBrowserSpecAllOf

`func NewAzureBlobBrowserSpecAllOf() *AzureBlobBrowserSpecAllOf`

NewAzureBlobBrowserSpecAllOf instantiates a new AzureBlobBrowserSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobBrowserSpecAllOfWithDefaults

`func NewAzureBlobBrowserSpecAllOfWithDefaults() *AzureBlobBrowserSpecAllOf`

NewAzureBlobBrowserSpecAllOfWithDefaults instantiates a new AzureBlobBrowserSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionType

`func (o *AzureBlobBrowserSpecAllOf) GetRegionType() EAzureRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AzureBlobBrowserSpecAllOf) GetRegionTypeOk() (*EAzureRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AzureBlobBrowserSpecAllOf) SetRegionType(v EAzureRegionType)`

SetRegionType sets RegionType field to given value.

### HasRegionType

`func (o *AzureBlobBrowserSpecAllOf) HasRegionType() bool`

HasRegionType returns a boolean if a field has been set.

### GetFolderType

`func (o *AzureBlobBrowserSpecAllOf) GetFolderType() ECloudBrowserFolderType`

GetFolderType returns the FolderType field if non-nil, zero value otherwise.

### GetFolderTypeOk

`func (o *AzureBlobBrowserSpecAllOf) GetFolderTypeOk() (*ECloudBrowserFolderType, bool)`

GetFolderTypeOk returns a tuple with the FolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderType

`func (o *AzureBlobBrowserSpecAllOf) SetFolderType(v ECloudBrowserFolderType)`

SetFolderType sets FolderType field to given value.

### HasFolderType

`func (o *AzureBlobBrowserSpecAllOf) HasFolderType() bool`

HasFolderType returns a boolean if a field has been set.

### GetGatewayServerId

`func (o *AzureBlobBrowserSpecAllOf) GetGatewayServerId() string`

GetGatewayServerId returns the GatewayServerId field if non-nil, zero value otherwise.

### GetGatewayServerIdOk

`func (o *AzureBlobBrowserSpecAllOf) GetGatewayServerIdOk() (*string, bool)`

GetGatewayServerIdOk returns a tuple with the GatewayServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServerId

`func (o *AzureBlobBrowserSpecAllOf) SetGatewayServerId(v string)`

SetGatewayServerId sets GatewayServerId field to given value.

### HasGatewayServerId

`func (o *AzureBlobBrowserSpecAllOf) HasGatewayServerId() bool`

HasGatewayServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


