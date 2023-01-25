# AdvancedSmtpOptionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int32** | Port number for the SMTP server. | 
**TimeoutMs** | **int32** | Connection timeout for the SMTP server. | 
**SSLEnabled** | **bool** | If *true*, secure connection for email operations is used. | 
**AuthRequired** | **bool** | If *true*, the &#x60;credentialsId&#x60; credentials are used to connect to the SMTP server. | 
**CredentialsId** | Pointer to **string** | ID of the credentials used to connect to the server. | [optional] 

## Methods

### NewAdvancedSmtpOptionsModel

`func NewAdvancedSmtpOptionsModel(port int32, timeoutMs int32, sSLEnabled bool, authRequired bool, ) *AdvancedSmtpOptionsModel`

NewAdvancedSmtpOptionsModel instantiates a new AdvancedSmtpOptionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedSmtpOptionsModelWithDefaults

`func NewAdvancedSmtpOptionsModelWithDefaults() *AdvancedSmtpOptionsModel`

NewAdvancedSmtpOptionsModelWithDefaults instantiates a new AdvancedSmtpOptionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *AdvancedSmtpOptionsModel) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AdvancedSmtpOptionsModel) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AdvancedSmtpOptionsModel) SetPort(v int32)`

SetPort sets Port field to given value.


### GetTimeoutMs

`func (o *AdvancedSmtpOptionsModel) GetTimeoutMs() int32`

GetTimeoutMs returns the TimeoutMs field if non-nil, zero value otherwise.

### GetTimeoutMsOk

`func (o *AdvancedSmtpOptionsModel) GetTimeoutMsOk() (*int32, bool)`

GetTimeoutMsOk returns a tuple with the TimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMs

`func (o *AdvancedSmtpOptionsModel) SetTimeoutMs(v int32)`

SetTimeoutMs sets TimeoutMs field to given value.


### GetSSLEnabled

`func (o *AdvancedSmtpOptionsModel) GetSSLEnabled() bool`

GetSSLEnabled returns the SSLEnabled field if non-nil, zero value otherwise.

### GetSSLEnabledOk

`func (o *AdvancedSmtpOptionsModel) GetSSLEnabledOk() (*bool, bool)`

GetSSLEnabledOk returns a tuple with the SSLEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSLEnabled

`func (o *AdvancedSmtpOptionsModel) SetSSLEnabled(v bool)`

SetSSLEnabled sets SSLEnabled field to given value.


### GetAuthRequired

`func (o *AdvancedSmtpOptionsModel) GetAuthRequired() bool`

GetAuthRequired returns the AuthRequired field if non-nil, zero value otherwise.

### GetAuthRequiredOk

`func (o *AdvancedSmtpOptionsModel) GetAuthRequiredOk() (*bool, bool)`

GetAuthRequiredOk returns a tuple with the AuthRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthRequired

`func (o *AdvancedSmtpOptionsModel) SetAuthRequired(v bool)`

SetAuthRequired sets AuthRequired field to given value.


### GetCredentialsId

`func (o *AdvancedSmtpOptionsModel) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *AdvancedSmtpOptionsModel) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *AdvancedSmtpOptionsModel) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.

### HasCredentialsId

`func (o *AdvancedSmtpOptionsModel) HasCredentialsId() bool`

HasCredentialsId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


