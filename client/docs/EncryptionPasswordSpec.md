# EncryptionPasswordSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | Password for data encryption. If you lose the password, you will not be able to restore it. | 
**Hint** | **string** | Hint for the encryption password. Provide a meaningful hint that will help you recall the password. | 
**Tag** | Pointer to **string** | Tag for the encryption password. | [optional] 

## Methods

### NewEncryptionPasswordSpec

`func NewEncryptionPasswordSpec(password string, hint string, ) *EncryptionPasswordSpec`

NewEncryptionPasswordSpec instantiates a new EncryptionPasswordSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionPasswordSpecWithDefaults

`func NewEncryptionPasswordSpecWithDefaults() *EncryptionPasswordSpec`

NewEncryptionPasswordSpecWithDefaults instantiates a new EncryptionPasswordSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *EncryptionPasswordSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EncryptionPasswordSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EncryptionPasswordSpec) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetHint

`func (o *EncryptionPasswordSpec) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *EncryptionPasswordSpec) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *EncryptionPasswordSpec) SetHint(v string)`

SetHint sets Hint field to given value.


### GetTag

`func (o *EncryptionPasswordSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *EncryptionPasswordSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *EncryptionPasswordSpec) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *EncryptionPasswordSpec) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


