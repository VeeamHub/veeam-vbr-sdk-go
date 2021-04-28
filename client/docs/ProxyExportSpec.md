# ProxyExportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | Array of backup proxy IDs. | [optional] 
**Types** | Pointer to [**[]EProxyType**](EProxyType.md) | Array of backup proxy types. | [optional] 
**Names** | Pointer to **[]string** | Array of backup proxy names. Wildcard characters are supported. | [optional] 

## Methods

### NewProxyExportSpec

`func NewProxyExportSpec() *ProxyExportSpec`

NewProxyExportSpec instantiates a new ProxyExportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyExportSpecWithDefaults

`func NewProxyExportSpecWithDefaults() *ProxyExportSpec`

NewProxyExportSpecWithDefaults instantiates a new ProxyExportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *ProxyExportSpec) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ProxyExportSpec) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ProxyExportSpec) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ProxyExportSpec) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetTypes

`func (o *ProxyExportSpec) GetTypes() []EProxyType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *ProxyExportSpec) GetTypesOk() (*[]EProxyType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *ProxyExportSpec) SetTypes(v []EProxyType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *ProxyExportSpec) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetNames

`func (o *ProxyExportSpec) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *ProxyExportSpec) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *ProxyExportSpec) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *ProxyExportSpec) HasNames() bool`

HasNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


