# PrivateKeyChangeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | **string** | Private key. | 
**Passphrase** | Pointer to **string** | Passphrase that protects the private key. | [optional] 

## Methods

### NewPrivateKeyChangeSpec

`func NewPrivateKeyChangeSpec(privateKey string, ) *PrivateKeyChangeSpec`

NewPrivateKeyChangeSpec instantiates a new PrivateKeyChangeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateKeyChangeSpecWithDefaults

`func NewPrivateKeyChangeSpecWithDefaults() *PrivateKeyChangeSpec`

NewPrivateKeyChangeSpecWithDefaults instantiates a new PrivateKeyChangeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *PrivateKeyChangeSpec) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PrivateKeyChangeSpec) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PrivateKeyChangeSpec) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetPassphrase

`func (o *PrivateKeyChangeSpec) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *PrivateKeyChangeSpec) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *PrivateKeyChangeSpec) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *PrivateKeyChangeSpec) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


