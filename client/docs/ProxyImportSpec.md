# ProxyImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the backup proxy. | 
**Type** | [**EProxyType**](EProxyType.md) |  | 
**Server** | [**ProxyServerSettingsImportSpec**](ProxyServerSettingsImportSpec.md) |  | 

## Methods

### NewProxyImportSpec

`func NewProxyImportSpec(description string, type_ EProxyType, server ProxyServerSettingsImportSpec, ) *ProxyImportSpec`

NewProxyImportSpec instantiates a new ProxyImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyImportSpecWithDefaults

`func NewProxyImportSpecWithDefaults() *ProxyImportSpec`

NewProxyImportSpecWithDefaults instantiates a new ProxyImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ProxyImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProxyImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProxyImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ProxyImportSpec) GetType() EProxyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProxyImportSpec) GetTypeOk() (*EProxyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProxyImportSpec) SetType(v EProxyType)`

SetType sets Type field to given value.


### GetServer

`func (o *ProxyImportSpec) GetServer() ProxyServerSettingsImportSpec`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ProxyImportSpec) GetServerOk() (*ProxyServerSettingsImportSpec, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ProxyImportSpec) SetServer(v ProxyServerSettingsImportSpec)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


