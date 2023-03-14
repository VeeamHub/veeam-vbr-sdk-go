# AzureComputeCloudDeviceCodeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Verification URI. | 
**VerificationCode** | **string** | Verification code. | 
**ExpirationTime** | **time.Time** | Expiration date and time of the verification code. By default, the code is valid for 15 minutes. | 

## Methods

### NewAzureComputeCloudDeviceCodeModel

`func NewAzureComputeCloudDeviceCodeModel(url string, verificationCode string, expirationTime time.Time, ) *AzureComputeCloudDeviceCodeModel`

NewAzureComputeCloudDeviceCodeModel instantiates a new AzureComputeCloudDeviceCodeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureComputeCloudDeviceCodeModelWithDefaults

`func NewAzureComputeCloudDeviceCodeModelWithDefaults() *AzureComputeCloudDeviceCodeModel`

NewAzureComputeCloudDeviceCodeModelWithDefaults instantiates a new AzureComputeCloudDeviceCodeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AzureComputeCloudDeviceCodeModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AzureComputeCloudDeviceCodeModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AzureComputeCloudDeviceCodeModel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVerificationCode

`func (o *AzureComputeCloudDeviceCodeModel) GetVerificationCode() string`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *AzureComputeCloudDeviceCodeModel) GetVerificationCodeOk() (*string, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *AzureComputeCloudDeviceCodeModel) SetVerificationCode(v string)`

SetVerificationCode sets VerificationCode field to given value.


### GetExpirationTime

`func (o *AzureComputeCloudDeviceCodeModel) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *AzureComputeCloudDeviceCodeModel) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *AzureComputeCloudDeviceCodeModel) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


