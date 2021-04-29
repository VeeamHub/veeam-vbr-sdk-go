# EncryptionPasswordModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the encryption password. | 
**Hint** | **string** | Hint for the encryption password. | 
**ModificationTime** | Pointer to **time.Time** | Date and time the password was last modified. | [optional] 
**Tag** | Pointer to **string** | Tag for the encryption password. | [optional] 

## Methods

### NewEncryptionPasswordModel

`func NewEncryptionPasswordModel(id string, hint string, ) *EncryptionPasswordModel`

NewEncryptionPasswordModel instantiates a new EncryptionPasswordModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionPasswordModelWithDefaults

`func NewEncryptionPasswordModelWithDefaults() *EncryptionPasswordModel`

NewEncryptionPasswordModelWithDefaults instantiates a new EncryptionPasswordModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EncryptionPasswordModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EncryptionPasswordModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EncryptionPasswordModel) SetId(v string)`

SetId sets Id field to given value.


### GetHint

`func (o *EncryptionPasswordModel) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *EncryptionPasswordModel) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *EncryptionPasswordModel) SetHint(v string)`

SetHint sets Hint field to given value.


### GetModificationTime

`func (o *EncryptionPasswordModel) GetModificationTime() time.Time`

GetModificationTime returns the ModificationTime field if non-nil, zero value otherwise.

### GetModificationTimeOk

`func (o *EncryptionPasswordModel) GetModificationTimeOk() (*time.Time, bool)`

GetModificationTimeOk returns a tuple with the ModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationTime

`func (o *EncryptionPasswordModel) SetModificationTime(v time.Time)`

SetModificationTime sets ModificationTime field to given value.

### HasModificationTime

`func (o *EncryptionPasswordModel) HasModificationTime() bool`

HasModificationTime returns a boolean if a field has been set.

### GetTag

`func (o *EncryptionPasswordModel) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *EncryptionPasswordModel) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *EncryptionPasswordModel) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *EncryptionPasswordModel) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


