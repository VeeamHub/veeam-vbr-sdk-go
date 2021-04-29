# RepositorySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the backup repository. | 
**Description** | **string** | Description of the backup repository. | 
**Tag** | Pointer to **string** | VMware vSphere tag assigned to the backup repository. | [optional] 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 

## Methods

### NewRepositorySpec

`func NewRepositorySpec(name string, description string, type_ ERepositoryType, ) *RepositorySpec`

NewRepositorySpec instantiates a new RepositorySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositorySpecWithDefaults

`func NewRepositorySpecWithDefaults() *RepositorySpec`

NewRepositorySpecWithDefaults instantiates a new RepositorySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RepositorySpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositorySpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositorySpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RepositorySpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RepositorySpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RepositorySpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *RepositorySpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *RepositorySpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *RepositorySpec) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *RepositorySpec) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetType

`func (o *RepositorySpec) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RepositorySpec) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RepositorySpec) SetType(v ERepositoryType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


