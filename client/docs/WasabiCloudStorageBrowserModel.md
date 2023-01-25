# WasabiCloudStorageBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | ID of a server used to connect to the object storage. | [optional] 
**Regions** | Pointer to [**[]WasabiCloudStorageRegionBrowserModel**](WasabiCloudStorageRegionBrowserModel.md) | Array of regions. | [optional] 

## Methods

### NewWasabiCloudStorageBrowserModel

`func NewWasabiCloudStorageBrowserModel() *WasabiCloudStorageBrowserModel`

NewWasabiCloudStorageBrowserModel instantiates a new WasabiCloudStorageBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWasabiCloudStorageBrowserModelWithDefaults

`func NewWasabiCloudStorageBrowserModelWithDefaults() *WasabiCloudStorageBrowserModel`

NewWasabiCloudStorageBrowserModelWithDefaults instantiates a new WasabiCloudStorageBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *WasabiCloudStorageBrowserModel) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *WasabiCloudStorageBrowserModel) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *WasabiCloudStorageBrowserModel) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *WasabiCloudStorageBrowserModel) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetRegions

`func (o *WasabiCloudStorageBrowserModel) GetRegions() []WasabiCloudStorageRegionBrowserModel`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *WasabiCloudStorageBrowserModel) GetRegionsOk() (*[]WasabiCloudStorageRegionBrowserModel, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *WasabiCloudStorageBrowserModel) SetRegions(v []WasabiCloudStorageRegionBrowserModel)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *WasabiCloudStorageBrowserModel) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


