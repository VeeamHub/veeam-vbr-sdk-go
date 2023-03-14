# EncryptionPasswordExportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModificationTimeFrom** | Pointer to **time.Time** | Date and time when the password was last modified. | [optional] 
**Ids** | Pointer to **[]string** | Array of password IDs. | [optional] 
**Hints** | Pointer to **[]string** | Array of password hints. | [optional] 
**Tags** | Pointer to **[]string** | Array of password tags. | [optional] 

## Methods

### NewEncryptionPasswordExportSpec

`func NewEncryptionPasswordExportSpec() *EncryptionPasswordExportSpec`

NewEncryptionPasswordExportSpec instantiates a new EncryptionPasswordExportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionPasswordExportSpecWithDefaults

`func NewEncryptionPasswordExportSpecWithDefaults() *EncryptionPasswordExportSpec`

NewEncryptionPasswordExportSpecWithDefaults instantiates a new EncryptionPasswordExportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModificationTimeFrom

`func (o *EncryptionPasswordExportSpec) GetModificationTimeFrom() time.Time`

GetModificationTimeFrom returns the ModificationTimeFrom field if non-nil, zero value otherwise.

### GetModificationTimeFromOk

`func (o *EncryptionPasswordExportSpec) GetModificationTimeFromOk() (*time.Time, bool)`

GetModificationTimeFromOk returns a tuple with the ModificationTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationTimeFrom

`func (o *EncryptionPasswordExportSpec) SetModificationTimeFrom(v time.Time)`

SetModificationTimeFrom sets ModificationTimeFrom field to given value.

### HasModificationTimeFrom

`func (o *EncryptionPasswordExportSpec) HasModificationTimeFrom() bool`

HasModificationTimeFrom returns a boolean if a field has been set.

### GetIds

`func (o *EncryptionPasswordExportSpec) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *EncryptionPasswordExportSpec) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *EncryptionPasswordExportSpec) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *EncryptionPasswordExportSpec) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetHints

`func (o *EncryptionPasswordExportSpec) GetHints() []string`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *EncryptionPasswordExportSpec) GetHintsOk() (*[]string, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *EncryptionPasswordExportSpec) SetHints(v []string)`

SetHints sets Hints field to given value.

### HasHints

`func (o *EncryptionPasswordExportSpec) HasHints() bool`

HasHints returns a boolean if a field has been set.

### GetTags

`func (o *EncryptionPasswordExportSpec) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EncryptionPasswordExportSpec) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EncryptionPasswordExportSpec) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EncryptionPasswordExportSpec) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


