# EncryptionPasswordImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | Password. | 
**Hint** | **string** | Hint for the encryption password. | 
**Tag** | Pointer to **string** | Tag for the encryption password. | [optional] 

## Methods

### NewEncryptionPasswordImportSpec

`func NewEncryptionPasswordImportSpec(password string, hint string, ) *EncryptionPasswordImportSpec`

NewEncryptionPasswordImportSpec instantiates a new EncryptionPasswordImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionPasswordImportSpecWithDefaults

`func NewEncryptionPasswordImportSpecWithDefaults() *EncryptionPasswordImportSpec`

NewEncryptionPasswordImportSpecWithDefaults instantiates a new EncryptionPasswordImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *EncryptionPasswordImportSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EncryptionPasswordImportSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EncryptionPasswordImportSpec) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetHint

`func (o *EncryptionPasswordImportSpec) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *EncryptionPasswordImportSpec) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *EncryptionPasswordImportSpec) SetHint(v string)`

SetHint sets Hint field to given value.


### GetTag

`func (o *EncryptionPasswordImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *EncryptionPasswordImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *EncryptionPasswordImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *EncryptionPasswordImportSpec) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


