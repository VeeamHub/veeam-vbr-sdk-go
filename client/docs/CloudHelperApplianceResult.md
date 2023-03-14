# CloudHelperApplianceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CloudHelperApplianceModel**](CloudHelperApplianceModel.md) | Array of Lunux-based helper appliances. | 
**Pagination** | [**PaginationResult**](PaginationResult.md) |  | 

## Methods

### NewCloudHelperApplianceResult

`func NewCloudHelperApplianceResult(data []CloudHelperApplianceModel, pagination PaginationResult, ) *CloudHelperApplianceResult`

NewCloudHelperApplianceResult instantiates a new CloudHelperApplianceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudHelperApplianceResultWithDefaults

`func NewCloudHelperApplianceResultWithDefaults() *CloudHelperApplianceResult`

NewCloudHelperApplianceResultWithDefaults instantiates a new CloudHelperApplianceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudHelperApplianceResult) GetData() []CloudHelperApplianceModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudHelperApplianceResult) GetDataOk() (*[]CloudHelperApplianceModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudHelperApplianceResult) SetData(v []CloudHelperApplianceModel)`

SetData sets Data field to given value.


### GetPagination

`func (o *CloudHelperApplianceResult) GetPagination() PaginationResult`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CloudHelperApplianceResult) GetPaginationOk() (*PaginationResult, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CloudHelperApplianceResult) SetPagination(v PaginationResult)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


