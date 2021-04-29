# BackupProxyImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the backup proxy. | 
**Type** | [**EBackupProxyImportType**](EBackupProxyImportType.md) |  | 
**Tag** | Pointer to **string** | Tag assigned to the backup proxy. | [optional] 

## Methods

### NewBackupProxyImportModel

`func NewBackupProxyImportModel(name string, type_ EBackupProxyImportType, ) *BackupProxyImportModel`

NewBackupProxyImportModel instantiates a new BackupProxyImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupProxyImportModelWithDefaults

`func NewBackupProxyImportModelWithDefaults() *BackupProxyImportModel`

NewBackupProxyImportModelWithDefaults instantiates a new BackupProxyImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BackupProxyImportModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupProxyImportModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupProxyImportModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BackupProxyImportModel) GetType() EBackupProxyImportType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupProxyImportModel) GetTypeOk() (*EBackupProxyImportType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupProxyImportModel) SetType(v EBackupProxyImportType)`

SetType sets Type field to given value.


### GetTag

`func (o *BackupProxyImportModel) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *BackupProxyImportModel) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *BackupProxyImportModel) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *BackupProxyImportModel) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


