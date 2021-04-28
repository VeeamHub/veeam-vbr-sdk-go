# SessionLogResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalRecords** | Pointer to **int32** | Total number of records. | [optional] 
**Records** | Pointer to [**[]SessionLogRecordModel**](SessionLogRecordModel.md) | Array of log records. | [optional] 

## Methods

### NewSessionLogResult

`func NewSessionLogResult() *SessionLogResult`

NewSessionLogResult instantiates a new SessionLogResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionLogResultWithDefaults

`func NewSessionLogResultWithDefaults() *SessionLogResult`

NewSessionLogResultWithDefaults instantiates a new SessionLogResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalRecords

`func (o *SessionLogResult) GetTotalRecords() int32`

GetTotalRecords returns the TotalRecords field if non-nil, zero value otherwise.

### GetTotalRecordsOk

`func (o *SessionLogResult) GetTotalRecordsOk() (*int32, bool)`

GetTotalRecordsOk returns a tuple with the TotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecords

`func (o *SessionLogResult) SetTotalRecords(v int32)`

SetTotalRecords sets TotalRecords field to given value.

### HasTotalRecords

`func (o *SessionLogResult) HasTotalRecords() bool`

HasTotalRecords returns a boolean if a field has been set.

### GetRecords

`func (o *SessionLogResult) GetRecords() []SessionLogRecordModel`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *SessionLogResult) GetRecordsOk() (*[]SessionLogRecordModel, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *SessionLogResult) SetRecords(v []SessionLogRecordModel)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *SessionLogResult) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


