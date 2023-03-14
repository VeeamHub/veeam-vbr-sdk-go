# AmazonVPCBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcName** | Pointer to **string** | VPC name. | [optional] 
**Subnets** | Pointer to **[]string** | Array of VPC subnets. | [optional] 
**SecurityGroups** | Pointer to **[]string** | Array of security groups. | [optional] 

## Methods

### NewAmazonVPCBrowserModel

`func NewAmazonVPCBrowserModel() *AmazonVPCBrowserModel`

NewAmazonVPCBrowserModel instantiates a new AmazonVPCBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonVPCBrowserModelWithDefaults

`func NewAmazonVPCBrowserModelWithDefaults() *AmazonVPCBrowserModel`

NewAmazonVPCBrowserModelWithDefaults instantiates a new AmazonVPCBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcName

`func (o *AmazonVPCBrowserModel) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *AmazonVPCBrowserModel) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *AmazonVPCBrowserModel) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *AmazonVPCBrowserModel) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### GetSubnets

`func (o *AmazonVPCBrowserModel) GetSubnets() []string`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *AmazonVPCBrowserModel) GetSubnetsOk() (*[]string, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *AmazonVPCBrowserModel) SetSubnets(v []string)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *AmazonVPCBrowserModel) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *AmazonVPCBrowserModel) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *AmazonVPCBrowserModel) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *AmazonVPCBrowserModel) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *AmazonVPCBrowserModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


