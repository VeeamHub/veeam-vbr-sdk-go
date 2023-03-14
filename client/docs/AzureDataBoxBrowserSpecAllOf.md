# AzureDataBoxBrowserSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePoint** | Pointer to **string** | Service endpoint used to connect to the Azure Data Box object storage. | [optional] 
**GatewayServerId** | Pointer to **string** | ID of a gateway server you want to use to connect to the object storage. Specify this parameter to check internet connection of the server. As a gateway server you can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 

## Methods

### NewAzureDataBoxBrowserSpecAllOf

`func NewAzureDataBoxBrowserSpecAllOf() *AzureDataBoxBrowserSpecAllOf`

NewAzureDataBoxBrowserSpecAllOf instantiates a new AzureDataBoxBrowserSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDataBoxBrowserSpecAllOfWithDefaults

`func NewAzureDataBoxBrowserSpecAllOfWithDefaults() *AzureDataBoxBrowserSpecAllOf`

NewAzureDataBoxBrowserSpecAllOfWithDefaults instantiates a new AzureDataBoxBrowserSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePoint

`func (o *AzureDataBoxBrowserSpecAllOf) GetServicePoint() string`

GetServicePoint returns the ServicePoint field if non-nil, zero value otherwise.

### GetServicePointOk

`func (o *AzureDataBoxBrowserSpecAllOf) GetServicePointOk() (*string, bool)`

GetServicePointOk returns a tuple with the ServicePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePoint

`func (o *AzureDataBoxBrowserSpecAllOf) SetServicePoint(v string)`

SetServicePoint sets ServicePoint field to given value.

### HasServicePoint

`func (o *AzureDataBoxBrowserSpecAllOf) HasServicePoint() bool`

HasServicePoint returns a boolean if a field has been set.

### GetGatewayServerId

`func (o *AzureDataBoxBrowserSpecAllOf) GetGatewayServerId() string`

GetGatewayServerId returns the GatewayServerId field if non-nil, zero value otherwise.

### GetGatewayServerIdOk

`func (o *AzureDataBoxBrowserSpecAllOf) GetGatewayServerIdOk() (*string, bool)`

GetGatewayServerIdOk returns a tuple with the GatewayServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServerId

`func (o *AzureDataBoxBrowserSpecAllOf) SetGatewayServerId(v string)`

SetGatewayServerId sets GatewayServerId field to given value.

### HasGatewayServerId

`func (o *AzureDataBoxBrowserSpecAllOf) HasGatewayServerId() bool`

HasGatewayServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


