# AzureResourceGroupBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceGroup** | **string** | Resource group name. | 
**VirtualNetworks** | [**[]AzureVirtualNetworkBrowserModel**](AzureVirtualNetworkBrowserModel.md) | Array of virtual networks available in the resource group. | 

## Methods

### NewAzureResourceGroupBrowserModel

`func NewAzureResourceGroupBrowserModel(resourceGroup string, virtualNetworks []AzureVirtualNetworkBrowserModel, ) *AzureResourceGroupBrowserModel`

NewAzureResourceGroupBrowserModel instantiates a new AzureResourceGroupBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureResourceGroupBrowserModelWithDefaults

`func NewAzureResourceGroupBrowserModelWithDefaults() *AzureResourceGroupBrowserModel`

NewAzureResourceGroupBrowserModelWithDefaults instantiates a new AzureResourceGroupBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceGroup

`func (o *AzureResourceGroupBrowserModel) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AzureResourceGroupBrowserModel) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AzureResourceGroupBrowserModel) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.


### GetVirtualNetworks

`func (o *AzureResourceGroupBrowserModel) GetVirtualNetworks() []AzureVirtualNetworkBrowserModel`

GetVirtualNetworks returns the VirtualNetworks field if non-nil, zero value otherwise.

### GetVirtualNetworksOk

`func (o *AzureResourceGroupBrowserModel) GetVirtualNetworksOk() (*[]AzureVirtualNetworkBrowserModel, bool)`

GetVirtualNetworksOk returns a tuple with the VirtualNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworks

`func (o *AzureResourceGroupBrowserModel) SetVirtualNetworks(v []AzureVirtualNetworkBrowserModel)`

SetVirtualNetworks sets VirtualNetworks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


