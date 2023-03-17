# AmazonEC2RegionBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionId** | Pointer to **string** | Region ID. | [optional] 
**Vpcs** | Pointer to [**[]AmazonVPCBrowserModel**](AmazonVPCBrowserModel.md) | Array of Amazon Virtual Private Cloud (Amazon VPC) networks. | [optional] 
**InstanceTypes** | Pointer to **[]string** | Array of Amazon instance types. | [optional] 

## Methods

### NewAmazonEC2RegionBrowserModel

`func NewAmazonEC2RegionBrowserModel() *AmazonEC2RegionBrowserModel`

NewAmazonEC2RegionBrowserModel instantiates a new AmazonEC2RegionBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonEC2RegionBrowserModelWithDefaults

`func NewAmazonEC2RegionBrowserModelWithDefaults() *AmazonEC2RegionBrowserModel`

NewAmazonEC2RegionBrowserModelWithDefaults instantiates a new AmazonEC2RegionBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionId

`func (o *AmazonEC2RegionBrowserModel) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AmazonEC2RegionBrowserModel) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AmazonEC2RegionBrowserModel) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *AmazonEC2RegionBrowserModel) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetVpcs

`func (o *AmazonEC2RegionBrowserModel) GetVpcs() []AmazonVPCBrowserModel`

GetVpcs returns the Vpcs field if non-nil, zero value otherwise.

### GetVpcsOk

`func (o *AmazonEC2RegionBrowserModel) GetVpcsOk() (*[]AmazonVPCBrowserModel, bool)`

GetVpcsOk returns a tuple with the Vpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcs

`func (o *AmazonEC2RegionBrowserModel) SetVpcs(v []AmazonVPCBrowserModel)`

SetVpcs sets Vpcs field to given value.

### HasVpcs

`func (o *AmazonEC2RegionBrowserModel) HasVpcs() bool`

HasVpcs returns a boolean if a field has been set.

### GetInstanceTypes

`func (o *AmazonEC2RegionBrowserModel) GetInstanceTypes() []string`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *AmazonEC2RegionBrowserModel) GetInstanceTypesOk() (*[]string, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *AmazonEC2RegionBrowserModel) SetInstanceTypes(v []string)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *AmazonEC2RegionBrowserModel) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


