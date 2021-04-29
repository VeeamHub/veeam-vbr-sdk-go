# ProxyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the backup proxy. | 
**Name** | **string** | Name of the backup proxy. | 
**Description** | **string** | Description of the backup proxy. | 
**Type** | [**EProxyType**](EProxyType.md) |  | 

## Methods

### NewProxyModel

`func NewProxyModel(id string, name string, description string, type_ EProxyType, ) *ProxyModel`

NewProxyModel instantiates a new ProxyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyModelWithDefaults

`func NewProxyModelWithDefaults() *ProxyModel`

NewProxyModelWithDefaults instantiates a new ProxyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProxyModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProxyModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProxyModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ProxyModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProxyModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProxyModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProxyModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProxyModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProxyModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ProxyModel) GetType() EProxyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProxyModel) GetTypeOk() (*EProxyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProxyModel) SetType(v EProxyType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


