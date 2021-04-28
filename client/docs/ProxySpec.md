# ProxySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the backup proxy. | 
**Type** | [**EProxyType**](EProxyType.md) |  | 

## Methods

### NewProxySpec

`func NewProxySpec(description string, type_ EProxyType, ) *ProxySpec`

NewProxySpec instantiates a new ProxySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxySpecWithDefaults

`func NewProxySpecWithDefaults() *ProxySpec`

NewProxySpecWithDefaults instantiates a new ProxySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ProxySpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProxySpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProxySpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ProxySpec) GetType() EProxyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProxySpec) GetTypeOk() (*EProxyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProxySpec) SetType(v EProxyType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


