# TokenModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | String that represents authorization issued to the client. It must be specified in all requests. An access token can be used multiple times, but its lifetime is 15 minutes. | 
**TokenType** | **string** | Type of the access token. | 
**RefreshToken** | **string** | String that is used to obtain a new access token if the current access token expires or becomes lost. A refresh token can be used only once, and its default lifetime is 14 days. | 
**ExpiresIn** | **int32** | Lifetime of the access token, in seconds. | 
**Issued** | **time.Time** | Date and time the access token is issued. | 
**Expires** | **time.Time** | Date and time the access token expires. | 

## Methods

### NewTokenModel

`func NewTokenModel(accessToken string, tokenType string, refreshToken string, expiresIn int32, issued time.Time, expires time.Time, ) *TokenModel`

NewTokenModel instantiates a new TokenModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenModelWithDefaults

`func NewTokenModelWithDefaults() *TokenModel`

NewTokenModelWithDefaults instantiates a new TokenModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *TokenModel) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenModel) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenModel) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetTokenType

`func (o *TokenModel) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TokenModel) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TokenModel) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetRefreshToken

`func (o *TokenModel) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *TokenModel) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *TokenModel) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.


### GetExpiresIn

`func (o *TokenModel) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *TokenModel) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *TokenModel) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.


### GetIssued

`func (o *TokenModel) GetIssued() time.Time`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *TokenModel) GetIssuedOk() (*time.Time, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *TokenModel) SetIssued(v time.Time)`

SetIssued sets Issued field to given value.


### GetExpires

`func (o *TokenModel) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TokenModel) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TokenModel) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


