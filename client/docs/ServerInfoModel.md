# ServerInfoModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VbrId** | Pointer to **string** | Veeam Backup &amp; Replication installation ID. | [optional] 
**Name** | **string** | Full DNS name or IP address of the backup server. | 
**BuildVersion** | **string** | Veeam Backup &amp; Replication build number. | 
**Patches** | **[]string** | Array of Veeam Backup &amp; Replication cumulative patches installed on the backup server. | 

## Methods

### NewServerInfoModel

`func NewServerInfoModel(name string, buildVersion string, patches []string, ) *ServerInfoModel`

NewServerInfoModel instantiates a new ServerInfoModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInfoModelWithDefaults

`func NewServerInfoModelWithDefaults() *ServerInfoModel`

NewServerInfoModelWithDefaults instantiates a new ServerInfoModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVbrId

`func (o *ServerInfoModel) GetVbrId() string`

GetVbrId returns the VbrId field if non-nil, zero value otherwise.

### GetVbrIdOk

`func (o *ServerInfoModel) GetVbrIdOk() (*string, bool)`

GetVbrIdOk returns a tuple with the VbrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVbrId

`func (o *ServerInfoModel) SetVbrId(v string)`

SetVbrId sets VbrId field to given value.

### HasVbrId

`func (o *ServerInfoModel) HasVbrId() bool`

HasVbrId returns a boolean if a field has been set.

### GetName

`func (o *ServerInfoModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerInfoModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerInfoModel) SetName(v string)`

SetName sets Name field to given value.


### GetBuildVersion

`func (o *ServerInfoModel) GetBuildVersion() string`

GetBuildVersion returns the BuildVersion field if non-nil, zero value otherwise.

### GetBuildVersionOk

`func (o *ServerInfoModel) GetBuildVersionOk() (*string, bool)`

GetBuildVersionOk returns a tuple with the BuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersion

`func (o *ServerInfoModel) SetBuildVersion(v string)`

SetBuildVersion sets BuildVersion field to given value.


### GetPatches

`func (o *ServerInfoModel) GetPatches() []string`

GetPatches returns the Patches field if non-nil, zero value otherwise.

### GetPatchesOk

`func (o *ServerInfoModel) GetPatchesOk() (*[]string, bool)`

GetPatchesOk returns a tuple with the Patches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatches

`func (o *ServerInfoModel) SetPatches(v []string)`

SetPatches sets Patches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


