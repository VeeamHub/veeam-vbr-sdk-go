# RestoreTargetFolderSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmName** | Pointer to **string** | VM name. | [optional] 
**Folder** | [**VmwareObjectModel**](VmwareObjectModel.md) |  | 
**RestoreVmTags** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication restores tags that were assigned to the original VMs, and assigns them to the restored VMs. | [optional] 

## Methods

### NewRestoreTargetFolderSpec

`func NewRestoreTargetFolderSpec(folder VmwareObjectModel, ) *RestoreTargetFolderSpec`

NewRestoreTargetFolderSpec instantiates a new RestoreTargetFolderSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreTargetFolderSpecWithDefaults

`func NewRestoreTargetFolderSpecWithDefaults() *RestoreTargetFolderSpec`

NewRestoreTargetFolderSpecWithDefaults instantiates a new RestoreTargetFolderSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmName

`func (o *RestoreTargetFolderSpec) GetVmName() string`

GetVmName returns the VmName field if non-nil, zero value otherwise.

### GetVmNameOk

`func (o *RestoreTargetFolderSpec) GetVmNameOk() (*string, bool)`

GetVmNameOk returns a tuple with the VmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmName

`func (o *RestoreTargetFolderSpec) SetVmName(v string)`

SetVmName sets VmName field to given value.

### HasVmName

`func (o *RestoreTargetFolderSpec) HasVmName() bool`

HasVmName returns a boolean if a field has been set.

### GetFolder

`func (o *RestoreTargetFolderSpec) GetFolder() VmwareObjectModel`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *RestoreTargetFolderSpec) GetFolderOk() (*VmwareObjectModel, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *RestoreTargetFolderSpec) SetFolder(v VmwareObjectModel)`

SetFolder sets Folder field to given value.


### GetRestoreVmTags

`func (o *RestoreTargetFolderSpec) GetRestoreVmTags() bool`

GetRestoreVmTags returns the RestoreVmTags field if non-nil, zero value otherwise.

### GetRestoreVmTagsOk

`func (o *RestoreTargetFolderSpec) GetRestoreVmTagsOk() (*bool, bool)`

GetRestoreVmTagsOk returns a tuple with the RestoreVmTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreVmTags

`func (o *RestoreTargetFolderSpec) SetRestoreVmTags(v bool)`

SetRestoreVmTags sets RestoreVmTags field to given value.

### HasRestoreVmTags

`func (o *RestoreTargetFolderSpec) HasRestoreVmTags() bool`

HasRestoreVmTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


