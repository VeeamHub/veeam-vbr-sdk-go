# GoogleCloudCredentialsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** | Access ID of a Google HMAC key. | 
**SecretKey** | **string** | Secret linked to the access ID. | 
**Tag** | Pointer to **string** | Tag used to identify the cloud credentials record. | [optional] 

## Methods

### NewGoogleCloudCredentialsSpec

`func NewGoogleCloudCredentialsSpec(accessKey string, secretKey string, ) *GoogleCloudCredentialsSpec`

NewGoogleCloudCredentialsSpec instantiates a new GoogleCloudCredentialsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudCredentialsSpecWithDefaults

`func NewGoogleCloudCredentialsSpecWithDefaults() *GoogleCloudCredentialsSpec`

NewGoogleCloudCredentialsSpecWithDefaults instantiates a new GoogleCloudCredentialsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *GoogleCloudCredentialsSpec) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *GoogleCloudCredentialsSpec) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *GoogleCloudCredentialsSpec) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *GoogleCloudCredentialsSpec) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *GoogleCloudCredentialsSpec) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *GoogleCloudCredentialsSpec) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetTag

`func (o *GoogleCloudCredentialsSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GoogleCloudCredentialsSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GoogleCloudCredentialsSpec) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GoogleCloudCredentialsSpec) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


