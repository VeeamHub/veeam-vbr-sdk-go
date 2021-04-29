# RepositoryShareGatewayModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSelect** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication automatically selects a gateway server. | [optional] 
**GatewayServerId** | Pointer to **string** | ID of the gateway server. | [optional] 

## Methods

### NewRepositoryShareGatewayModel

`func NewRepositoryShareGatewayModel() *RepositoryShareGatewayModel`

NewRepositoryShareGatewayModel instantiates a new RepositoryShareGatewayModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryShareGatewayModelWithDefaults

`func NewRepositoryShareGatewayModelWithDefaults() *RepositoryShareGatewayModel`

NewRepositoryShareGatewayModelWithDefaults instantiates a new RepositoryShareGatewayModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSelect

`func (o *RepositoryShareGatewayModel) GetAutoSelect() bool`

GetAutoSelect returns the AutoSelect field if non-nil, zero value otherwise.

### GetAutoSelectOk

`func (o *RepositoryShareGatewayModel) GetAutoSelectOk() (*bool, bool)`

GetAutoSelectOk returns a tuple with the AutoSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelect

`func (o *RepositoryShareGatewayModel) SetAutoSelect(v bool)`

SetAutoSelect sets AutoSelect field to given value.

### HasAutoSelect

`func (o *RepositoryShareGatewayModel) HasAutoSelect() bool`

HasAutoSelect returns a boolean if a field has been set.

### GetGatewayServerId

`func (o *RepositoryShareGatewayModel) GetGatewayServerId() string`

GetGatewayServerId returns the GatewayServerId field if non-nil, zero value otherwise.

### GetGatewayServerIdOk

`func (o *RepositoryShareGatewayModel) GetGatewayServerIdOk() (*string, bool)`

GetGatewayServerIdOk returns a tuple with the GatewayServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServerId

`func (o *RepositoryShareGatewayModel) SetGatewayServerId(v string)`

SetGatewayServerId sets GatewayServerId field to given value.

### HasGatewayServerId

`func (o *RepositoryShareGatewayModel) HasGatewayServerId() bool`

HasGatewayServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


