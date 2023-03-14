# AzureVirtualNetworkBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualNetworkName** | Pointer to **string** | Virtual network name. | [optional] 
**Subnets** | Pointer to **[]string** | Array of subnets. | [optional] 

## Methods

### NewAzureVirtualNetworkBrowserModel

`func NewAzureVirtualNetworkBrowserModel() *AzureVirtualNetworkBrowserModel`

NewAzureVirtualNetworkBrowserModel instantiates a new AzureVirtualNetworkBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVirtualNetworkBrowserModelWithDefaults

`func NewAzureVirtualNetworkBrowserModelWithDefaults() *AzureVirtualNetworkBrowserModel`

NewAzureVirtualNetworkBrowserModelWithDefaults instantiates a new AzureVirtualNetworkBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualNetworkName

`func (o *AzureVirtualNetworkBrowserModel) GetVirtualNetworkName() string`

GetVirtualNetworkName returns the VirtualNetworkName field if non-nil, zero value otherwise.

### GetVirtualNetworkNameOk

`func (o *AzureVirtualNetworkBrowserModel) GetVirtualNetworkNameOk() (*string, bool)`

GetVirtualNetworkNameOk returns a tuple with the VirtualNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkName

`func (o *AzureVirtualNetworkBrowserModel) SetVirtualNetworkName(v string)`

SetVirtualNetworkName sets VirtualNetworkName field to given value.

### HasVirtualNetworkName

`func (o *AzureVirtualNetworkBrowserModel) HasVirtualNetworkName() bool`

HasVirtualNetworkName returns a boolean if a field has been set.

### GetSubnets

`func (o *AzureVirtualNetworkBrowserModel) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *AzureVirtualNetworkBrowserModel) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *AzureVirtualNetworkBrowserModel) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *AzureVirtualNetworkBrowserModel) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


