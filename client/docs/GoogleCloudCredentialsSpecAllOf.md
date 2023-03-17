# GoogleCloudCredentialsSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** | Access ID of a Google HMAC key. | 
**SecretKey** | **string** | Secret linked to the access ID. | 
**Tag** | Pointer to **string** | Tag used to identify the cloud credentials record. | [optional] 

## Methods

### NewGoogleCloudCredentialsSpecAllOf

`func NewGoogleCloudCredentialsSpecAllOf(accessKey string, secretKey string, ) *GoogleCloudCredentialsSpecAllOf`

NewGoogleCloudCredentialsSpecAllOf instantiates a new GoogleCloudCredentialsSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudCredentialsSpecAllOfWithDefaults

`func NewGoogleCloudCredentialsSpecAllOfWithDefaults() *GoogleCloudCredentialsSpecAllOf`

NewGoogleCloudCredentialsSpecAllOfWithDefaults instantiates a new GoogleCloudCredentialsSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *GoogleCloudCredentialsSpecAllOf) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *GoogleCloudCredentialsSpecAllOf) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *GoogleCloudCredentialsSpecAllOf) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *GoogleCloudCredentialsSpecAllOf) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *GoogleCloudCredentialsSpecAllOf) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *GoogleCloudCredentialsSpecAllOf) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetTag

`func (o *GoogleCloudCredentialsSpecAllOf) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GoogleCloudCredentialsSpecAllOf) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GoogleCloudCredentialsSpecAllOf) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GoogleCloudCredentialsSpecAllOf) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


