# SessionResultModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | [**ESessionResult**](ESessionResult.md) |  | 
**Message** | Pointer to **string** | Message that explains the session result. | [optional] 
**IsCanceled** | Pointer to **bool** | If *true*, the session has been canceled. | [optional] 

## Methods

### NewSessionResultModel

`func NewSessionResultModel(result ESessionResult, ) *SessionResultModel`

NewSessionResultModel instantiates a new SessionResultModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionResultModelWithDefaults

`func NewSessionResultModelWithDefaults() *SessionResultModel`

NewSessionResultModelWithDefaults instantiates a new SessionResultModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *SessionResultModel) GetResult() ESessionResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SessionResultModel) GetResultOk() (*ESessionResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SessionResultModel) SetResult(v ESessionResult)`

SetResult sets Result field to given value.


### GetMessage

`func (o *SessionResultModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SessionResultModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SessionResultModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SessionResultModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetIsCanceled

`func (o *SessionResultModel) GetIsCanceled() bool`

GetIsCanceled returns the IsCanceled field if non-nil, zero value otherwise.

### GetIsCanceledOk

`func (o *SessionResultModel) GetIsCanceledOk() (*bool, bool)`

GetIsCanceledOk returns a tuple with the IsCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCanceled

`func (o *SessionResultModel) SetIsCanceled(v bool)`

SetIsCanceled sets IsCanceled field to given value.

### HasIsCanceled

`func (o *SessionResultModel) HasIsCanceled() bool`

HasIsCanceled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


