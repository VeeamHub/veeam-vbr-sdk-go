# TokenLoginSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | [**ELoginGrantType**](ELoginGrantType.md) |  | [default to ELOGINGRANTTYPE_PASSWORD]
**Username** | Pointer to **string** | User name. Required if the &#x60;grant_type&#x60; value is &#x60;password&#x60;. | [optional] 
**Password** | Pointer to **string** | Password. Required if the &#x60;grant_type&#x60; value is &#x60;password&#x60;. | [optional] 
**RefreshToken** | Pointer to **string** | Refresh token. Required if the &#x60;grant_type&#x60; value is &#x60;refresh_token&#x60;. | [optional] 
**Code** | Pointer to **string** | Authorization code. Required if the &#x60;grant_type&#x60; value is &#x60;authorization_code&#x60;. | [optional] 
**UseShortTermRefresh** | Pointer to **bool** | If *true*, a short-term refresh token is used. Lifetime of the short-term refresh token is the access token lifetime plus 15 minutes. | [optional] 
**VbrToken** | Pointer to **string** | Veeam Backup &amp; Replication platform service token. | [optional] 

## Methods

### NewTokenLoginSpec

`func NewTokenLoginSpec(grantType ELoginGrantType, ) *TokenLoginSpec`

NewTokenLoginSpec instantiates a new TokenLoginSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenLoginSpecWithDefaults

`func NewTokenLoginSpecWithDefaults() *TokenLoginSpec`

NewTokenLoginSpecWithDefaults instantiates a new TokenLoginSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *TokenLoginSpec) GetGrantType() ELoginGrantType`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *TokenLoginSpec) GetGrantTypeOk() (*ELoginGrantType, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *TokenLoginSpec) SetGrantType(v ELoginGrantType)`

SetGrantType sets GrantType field to given value.


### GetUsername

`func (o *TokenLoginSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TokenLoginSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TokenLoginSpec) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TokenLoginSpec) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *TokenLoginSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TokenLoginSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TokenLoginSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TokenLoginSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRefreshToken

`func (o *TokenLoginSpec) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *TokenLoginSpec) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *TokenLoginSpec) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *TokenLoginSpec) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetCode

`func (o *TokenLoginSpec) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TokenLoginSpec) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TokenLoginSpec) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TokenLoginSpec) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetUseShortTermRefresh

`func (o *TokenLoginSpec) GetUseShortTermRefresh() bool`

GetUseShortTermRefresh returns the UseShortTermRefresh field if non-nil, zero value otherwise.

### GetUseShortTermRefreshOk

`func (o *TokenLoginSpec) GetUseShortTermRefreshOk() (*bool, bool)`

GetUseShortTermRefreshOk returns a tuple with the UseShortTermRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseShortTermRefresh

`func (o *TokenLoginSpec) SetUseShortTermRefresh(v bool)`

SetUseShortTermRefresh sets UseShortTermRefresh field to given value.

### HasUseShortTermRefresh

`func (o *TokenLoginSpec) HasUseShortTermRefresh() bool`

HasUseShortTermRefresh returns a boolean if a field has been set.

### GetVbrToken

`func (o *TokenLoginSpec) GetVbrToken() string`

GetVbrToken returns the VbrToken field if non-nil, zero value otherwise.

### GetVbrTokenOk

`func (o *TokenLoginSpec) GetVbrTokenOk() (*string, bool)`

GetVbrTokenOk returns a tuple with the VbrToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVbrToken

`func (o *TokenLoginSpec) SetVbrToken(v string)`

SetVbrToken sets VbrToken field to given value.

### HasVbrToken

`func (o *TokenLoginSpec) HasVbrToken() bool`

HasVbrToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


