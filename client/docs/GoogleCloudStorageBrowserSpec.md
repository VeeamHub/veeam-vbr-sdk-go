# GoogleCloudStorageBrowserSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayServerId** | Pointer to **string** | ID of a gateway server you want to use to connect to the object storage. Specify this parameter to check internet connection of the server. As a gateway server you can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 
**Filters** | Pointer to [**GoogleCloudStorageBrowserFilters**](GoogleCloudStorageBrowserFilters.md) |  | [optional] 

## Methods

### NewGoogleCloudStorageBrowserSpec

`func NewGoogleCloudStorageBrowserSpec() *GoogleCloudStorageBrowserSpec`

NewGoogleCloudStorageBrowserSpec instantiates a new GoogleCloudStorageBrowserSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudStorageBrowserSpecWithDefaults

`func NewGoogleCloudStorageBrowserSpecWithDefaults() *GoogleCloudStorageBrowserSpec`

NewGoogleCloudStorageBrowserSpecWithDefaults instantiates a new GoogleCloudStorageBrowserSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayServerId

`func (o *GoogleCloudStorageBrowserSpec) GetGatewayServerId() string`

GetGatewayServerId returns the GatewayServerId field if non-nil, zero value otherwise.

### GetGatewayServerIdOk

`func (o *GoogleCloudStorageBrowserSpec) GetGatewayServerIdOk() (*string, bool)`

GetGatewayServerIdOk returns a tuple with the GatewayServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServerId

`func (o *GoogleCloudStorageBrowserSpec) SetGatewayServerId(v string)`

SetGatewayServerId sets GatewayServerId field to given value.

### HasGatewayServerId

`func (o *GoogleCloudStorageBrowserSpec) HasGatewayServerId() bool`

HasGatewayServerId returns a boolean if a field has been set.

### GetFilters

`func (o *GoogleCloudStorageBrowserSpec) GetFilters() GoogleCloudStorageBrowserFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GoogleCloudStorageBrowserSpec) GetFiltersOk() (*GoogleCloudStorageBrowserFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GoogleCloudStorageBrowserSpec) SetFilters(v GoogleCloudStorageBrowserFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GoogleCloudStorageBrowserSpec) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


