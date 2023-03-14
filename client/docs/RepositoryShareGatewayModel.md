# RepositoryShareGatewayModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSelect** | **bool** | If *true*, Veeam Backup &amp; Replication automatically selects a gateway server. | 
**GatewayServerIds** | Pointer to **[]string** | Array of gateway server IDs. | [optional] 

## Methods

### NewRepositoryShareGatewayModel

`func NewRepositoryShareGatewayModel(autoSelect bool, ) *RepositoryShareGatewayModel`

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


### GetGatewayServerIds

`func (o *RepositoryShareGatewayModel) GetGatewayServerIds() []string`

GetGatewayServerIds returns the GatewayServerIds field if non-nil, zero value otherwise.

### GetGatewayServerIdsOk

`func (o *RepositoryShareGatewayModel) GetGatewayServerIdsOk() (*[]string, bool)`

GetGatewayServerIdsOk returns a tuple with the GatewayServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServerIds

`func (o *RepositoryShareGatewayModel) SetGatewayServerIds(v []string)`

SetGatewayServerIds sets GatewayServerIds field to given value.

### HasGatewayServerIds

`func (o *RepositoryShareGatewayModel) HasGatewayServerIds() bool`

HasGatewayServerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


