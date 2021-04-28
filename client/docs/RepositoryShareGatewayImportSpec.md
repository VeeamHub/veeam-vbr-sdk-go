# RepositoryShareGatewayImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSelect** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication automatically selects a gateway server. | [optional] 
**GatewayServerName** | Pointer to **string** | Name of the gateway server. | [optional] 

## Methods

### NewRepositoryShareGatewayImportSpec

`func NewRepositoryShareGatewayImportSpec() *RepositoryShareGatewayImportSpec`

NewRepositoryShareGatewayImportSpec instantiates a new RepositoryShareGatewayImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryShareGatewayImportSpecWithDefaults

`func NewRepositoryShareGatewayImportSpecWithDefaults() *RepositoryShareGatewayImportSpec`

NewRepositoryShareGatewayImportSpecWithDefaults instantiates a new RepositoryShareGatewayImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSelect

`func (o *RepositoryShareGatewayImportSpec) GetAutoSelect() bool`

GetAutoSelect returns the AutoSelect field if non-nil, zero value otherwise.

### GetAutoSelectOk

`func (o *RepositoryShareGatewayImportSpec) GetAutoSelectOk() (*bool, bool)`

GetAutoSelectOk returns a tuple with the AutoSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelect

`func (o *RepositoryShareGatewayImportSpec) SetAutoSelect(v bool)`

SetAutoSelect sets AutoSelect field to given value.

### HasAutoSelect

`func (o *RepositoryShareGatewayImportSpec) HasAutoSelect() bool`

HasAutoSelect returns a boolean if a field has been set.

### GetGatewayServerName

`func (o *RepositoryShareGatewayImportSpec) GetGatewayServerName() string`

GetGatewayServerName returns the GatewayServerName field if non-nil, zero value otherwise.

### GetGatewayServerNameOk

`func (o *RepositoryShareGatewayImportSpec) GetGatewayServerNameOk() (*string, bool)`

GetGatewayServerNameOk returns a tuple with the GatewayServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServerName

`func (o *RepositoryShareGatewayImportSpec) SetGatewayServerName(v string)`

SetGatewayServerName sets GatewayServerName field to given value.

### HasGatewayServerName

`func (o *RepositoryShareGatewayImportSpec) HasGatewayServerName() bool`

HasGatewayServerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


