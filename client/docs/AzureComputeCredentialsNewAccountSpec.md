# AzureComputeCredentialsNewAccountSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | [**EAzureRegionType**](EAzureRegionType.md) |  | 
**VerificationCode** | **string** | Single-use verification code. Use this code to sign in on the https://microsoft.com/devicelogin page. | 

## Methods

### NewAzureComputeCredentialsNewAccountSpec

`func NewAzureComputeCredentialsNewAccountSpec(region EAzureRegionType, verificationCode string, ) *AzureComputeCredentialsNewAccountSpec`

NewAzureComputeCredentialsNewAccountSpec instantiates a new AzureComputeCredentialsNewAccountSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureComputeCredentialsNewAccountSpecWithDefaults

`func NewAzureComputeCredentialsNewAccountSpecWithDefaults() *AzureComputeCredentialsNewAccountSpec`

NewAzureComputeCredentialsNewAccountSpecWithDefaults instantiates a new AzureComputeCredentialsNewAccountSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *AzureComputeCredentialsNewAccountSpec) GetRegion() EAzureRegionType`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AzureComputeCredentialsNewAccountSpec) GetRegionOk() (*EAzureRegionType, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AzureComputeCredentialsNewAccountSpec) SetRegion(v EAzureRegionType)`

SetRegion sets Region field to given value.


### GetVerificationCode

`func (o *AzureComputeCredentialsNewAccountSpec) GetVerificationCode() string`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *AzureComputeCredentialsNewAccountSpec) GetVerificationCodeOk() (*string, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *AzureComputeCredentialsNewAccountSpec) SetVerificationCode(v string)`

SetVerificationCode sets VerificationCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


