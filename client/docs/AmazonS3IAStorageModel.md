# AmazonS3IAStorageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | If *true*, Standard Infrequent Access is enabled. | [optional] 
**SingleZoneEnabled** | Pointer to **bool** | If *true*, Amazon S3 One Zone-Infrequent Access is enabled. | [optional] 

## Methods

### NewAmazonS3IAStorageModel

`func NewAmazonS3IAStorageModel() *AmazonS3IAStorageModel`

NewAmazonS3IAStorageModel instantiates a new AmazonS3IAStorageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3IAStorageModelWithDefaults

`func NewAmazonS3IAStorageModelWithDefaults() *AmazonS3IAStorageModel`

NewAmazonS3IAStorageModelWithDefaults instantiates a new AmazonS3IAStorageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *AmazonS3IAStorageModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AmazonS3IAStorageModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AmazonS3IAStorageModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *AmazonS3IAStorageModel) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetSingleZoneEnabled

`func (o *AmazonS3IAStorageModel) GetSingleZoneEnabled() bool`

GetSingleZoneEnabled returns the SingleZoneEnabled field if non-nil, zero value otherwise.

### GetSingleZoneEnabledOk

`func (o *AmazonS3IAStorageModel) GetSingleZoneEnabledOk() (*bool, bool)`

GetSingleZoneEnabledOk returns a tuple with the SingleZoneEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleZoneEnabled

`func (o *AmazonS3IAStorageModel) SetSingleZoneEnabled(v bool)`

SetSingleZoneEnabled sets SingleZoneEnabled field to given value.

### HasSingleZoneEnabled

`func (o *AmazonS3IAStorageModel) HasSingleZoneEnabled() bool`

HasSingleZoneEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


