# ProxyDatastoreSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSelect** | **bool** | If *true*, all datastores that the backup proxy can access are detected automatically. | 
**Datastores** | Pointer to [**[]ProxyDatastoreModel**](ProxyDatastoreModel.md) | Array of datastores to which the backup proxy has a direct SAN or NFS connection. | [optional] 

## Methods

### NewProxyDatastoreSettingsModel

`func NewProxyDatastoreSettingsModel(autoSelect bool, ) *ProxyDatastoreSettingsModel`

NewProxyDatastoreSettingsModel instantiates a new ProxyDatastoreSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyDatastoreSettingsModelWithDefaults

`func NewProxyDatastoreSettingsModelWithDefaults() *ProxyDatastoreSettingsModel`

NewProxyDatastoreSettingsModelWithDefaults instantiates a new ProxyDatastoreSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSelect

`func (o *ProxyDatastoreSettingsModel) GetAutoSelect() bool`

GetAutoSelect returns the AutoSelect field if non-nil, zero value otherwise.

### GetAutoSelectOk

`func (o *ProxyDatastoreSettingsModel) GetAutoSelectOk() (*bool, bool)`

GetAutoSelectOk returns a tuple with the AutoSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelect

`func (o *ProxyDatastoreSettingsModel) SetAutoSelect(v bool)`

SetAutoSelect sets AutoSelect field to given value.


### GetDatastores

`func (o *ProxyDatastoreSettingsModel) GetDatastores() []ProxyDatastoreModel`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *ProxyDatastoreSettingsModel) GetDatastoresOk() (*[]ProxyDatastoreModel, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *ProxyDatastoreSettingsModel) SetDatastores(v []ProxyDatastoreModel)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *ProxyDatastoreSettingsModel) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


