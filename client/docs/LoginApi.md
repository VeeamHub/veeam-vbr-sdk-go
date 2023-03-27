# \LoginApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthorizationCode**](LoginApi.md#CreateAuthorizationCode) | **Post** /api/oauth2/authorization_code | Get Authorization Code
[**CreateToken**](LoginApi.md#CreateToken) | **Post** /api/oauth2/token | Get Access Token
[**Logout**](LoginApi.md#Logout) | **Post** /api/oauth2/logout | Log Out



## CreateAuthorizationCode

> AuthorizationCodeModel CreateAuthorizationCode(ctx).XApiVersion(xApiVersion).Execute()

Get Authorization Code



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginApi.CreateAuthorizationCode(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.CreateAuthorizationCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthorizationCode`: AuthorizationCodeModel
    fmt.Fprintf(os.Stdout, "Response from `LoginApi.CreateAuthorizationCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthorizationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

### Return type

[**AuthorizationCodeModel**](AuthorizationCodeModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateToken

> TokenModel CreateToken(ctx).XApiVersion(xApiVersion).GrantType(grantType).Username(username).Password(password).RefreshToken(refreshToken).Code(code).UseShortTermRefresh(useShortTermRefresh).VbrToken(vbrToken).Execute()

Get Access Token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    grantType := client.ELoginGrantType("password") // ELoginGrantType |  (default to "password")
    username := "username_example" // string | User name. Required if the `grant_type` value is `password`. (optional)
    password := "password_example" // string | Password. Required if the `grant_type` value is `password`. (optional)
    refreshToken := "refreshToken_example" // string | Refresh token. Required if the `grant_type` value is `refresh_token`. (optional)
    code := "code_example" // string | Authorization code. Required if the `grant_type` value is `authorization_code`. (optional)
    useShortTermRefresh := true // bool | If *true*, a short-term refresh token is used. Lifetime of the short-term refresh token is the access token lifetime plus 15 minutes. (optional)
    vbrToken := "vbrToken_example" // string | Veeam Backup & Replication platform service token. (optional)

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginApi.CreateToken(context.Background()).XApiVersion(xApiVersion).GrantType(grantType).Username(username).Password(password).RefreshToken(refreshToken).Code(code).UseShortTermRefresh(useShortTermRefresh).VbrToken(vbrToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.CreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToken`: TokenModel
    fmt.Fprintf(os.Stdout, "Response from `LoginApi.CreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **grantType** | [**ELoginGrantType**](ELoginGrantType.md) |  | [default to &quot;password&quot;]
 **username** | **string** | User name. Required if the &#x60;grant_type&#x60; value is &#x60;password&#x60;. | 
 **password** | **string** | Password. Required if the &#x60;grant_type&#x60; value is &#x60;password&#x60;. | 
 **refreshToken** | **string** | Refresh token. Required if the &#x60;grant_type&#x60; value is &#x60;refresh_token&#x60;. | 
 **code** | **string** | Authorization code. Required if the &#x60;grant_type&#x60; value is &#x60;authorization_code&#x60;. | 
 **useShortTermRefresh** | **bool** | If *true*, a short-term refresh token is used. Lifetime of the short-term refresh token is the access token lifetime plus 15 minutes. | 
 **vbrToken** | **string** | Veeam Backup &amp; Replication platform service token. | 

### Return type

[**TokenModel**](TokenModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> map[string]interface{} Logout(ctx).XApiVersion(xApiVersion).Execute()

Log Out



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/veeamhub/veeam-vbr-sdk-go/client"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginApi.Logout(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginApi.Logout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Logout`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LoginApi.Logout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

