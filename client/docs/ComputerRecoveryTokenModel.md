# ComputerRecoveryTokenModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the recovery token. | 
**Name** | **string** | Friendly name of the recovery token. | 
**RecoveryToken** | Pointer to **string** | Recovery token. | [optional] 
**ExpirationDate** | **time.Time** | Date and time when the recovery token expires. | 

## Methods

### NewComputerRecoveryTokenModel

`func NewComputerRecoveryTokenModel(id string, name string, expirationDate time.Time, ) *ComputerRecoveryTokenModel`

NewComputerRecoveryTokenModel instantiates a new ComputerRecoveryTokenModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerRecoveryTokenModelWithDefaults

`func NewComputerRecoveryTokenModelWithDefaults() *ComputerRecoveryTokenModel`

NewComputerRecoveryTokenModelWithDefaults instantiates a new ComputerRecoveryTokenModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComputerRecoveryTokenModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputerRecoveryTokenModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputerRecoveryTokenModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ComputerRecoveryTokenModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputerRecoveryTokenModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputerRecoveryTokenModel) SetName(v string)`

SetName sets Name field to given value.


### GetRecoveryToken

`func (o *ComputerRecoveryTokenModel) GetRecoveryToken() string`

GetRecoveryToken returns the RecoveryToken field if non-nil, zero value otherwise.

### GetRecoveryTokenOk

`func (o *ComputerRecoveryTokenModel) GetRecoveryTokenOk() (*string, bool)`

GetRecoveryTokenOk returns a tuple with the RecoveryToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryToken

`func (o *ComputerRecoveryTokenModel) SetRecoveryToken(v string)`

SetRecoveryToken sets RecoveryToken field to given value.

### HasRecoveryToken

`func (o *ComputerRecoveryTokenModel) HasRecoveryToken() bool`

HasRecoveryToken returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ComputerRecoveryTokenModel) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ComputerRecoveryTokenModel) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ComputerRecoveryTokenModel) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


