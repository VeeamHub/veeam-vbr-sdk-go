# AmazonCloudCredentialsSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** | ID of the access key. | 
**SecretKey** | **string** | Secret key. | 
**Tag** | Pointer to **string** | Tag used to identify the cloud credentials record. | [optional] 

## Methods

### NewAmazonCloudCredentialsSpecAllOf

`func NewAmazonCloudCredentialsSpecAllOf(accessKey string, secretKey string, ) *AmazonCloudCredentialsSpecAllOf`

NewAmazonCloudCredentialsSpecAllOf instantiates a new AmazonCloudCredentialsSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonCloudCredentialsSpecAllOfWithDefaults

`func NewAmazonCloudCredentialsSpecAllOfWithDefaults() *AmazonCloudCredentialsSpecAllOf`

NewAmazonCloudCredentialsSpecAllOfWithDefaults instantiates a new AmazonCloudCredentialsSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *AmazonCloudCredentialsSpecAllOf) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AmazonCloudCredentialsSpecAllOf) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AmazonCloudCredentialsSpecAllOf) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *AmazonCloudCredentialsSpecAllOf) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AmazonCloudCredentialsSpecAllOf) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AmazonCloudCredentialsSpecAllOf) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetTag

`func (o *AmazonCloudCredentialsSpecAllOf) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AmazonCloudCredentialsSpecAllOf) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AmazonCloudCredentialsSpecAllOf) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AmazonCloudCredentialsSpecAllOf) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


