# GoogleCloudStorageBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | ID of a server used to connect to the object storage. | [optional] 
**Regions** | Pointer to [**[]GoogleCloudStorageRegionBrowserModel**](GoogleCloudStorageRegionBrowserModel.md) | Array of regions. | [optional] 

## Methods

### NewGoogleCloudStorageBrowserModel

`func NewGoogleCloudStorageBrowserModel() *GoogleCloudStorageBrowserModel`

NewGoogleCloudStorageBrowserModel instantiates a new GoogleCloudStorageBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudStorageBrowserModelWithDefaults

`func NewGoogleCloudStorageBrowserModelWithDefaults() *GoogleCloudStorageBrowserModel`

NewGoogleCloudStorageBrowserModelWithDefaults instantiates a new GoogleCloudStorageBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *GoogleCloudStorageBrowserModel) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *GoogleCloudStorageBrowserModel) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *GoogleCloudStorageBrowserModel) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *GoogleCloudStorageBrowserModel) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetRegions

`func (o *GoogleCloudStorageBrowserModel) GetRegions() []GoogleCloudStorageRegionBrowserModel`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *GoogleCloudStorageBrowserModel) GetRegionsOk() (*[]GoogleCloudStorageRegionBrowserModel, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *GoogleCloudStorageBrowserModel) SetRegions(v []GoogleCloudStorageRegionBrowserModel)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *GoogleCloudStorageBrowserModel) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


