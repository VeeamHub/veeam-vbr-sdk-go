# \AgentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComputerRecoveryToken**](AgentsApi.md#CreateComputerRecoveryToken) | **Post** /api/v1/agents/recoveryTokens | Create Recovery Token
[**DeleteComputerRecoveryToken**](AgentsApi.md#DeleteComputerRecoveryToken) | **Delete** /api/v1/agents/recoveryTokens/{id} | Delete Recovery Token
[**GetAllComputerRecoveryTokens**](AgentsApi.md#GetAllComputerRecoveryTokens) | **Get** /api/v1/agents/recoveryTokens | Get All Recovery Tokens
[**GetComputerRecoveryToken**](AgentsApi.md#GetComputerRecoveryToken) | **Get** /api/v1/agents/recoveryTokens/{id} | Get Recovery Token
[**UpdateComputerRecoveryToken**](AgentsApi.md#UpdateComputerRecoveryToken) | **Put** /api/v1/agents/recoveryTokens/{id} | Edit Recovery Token



## CreateComputerRecoveryToken

> ComputerRecoveryTokenModel CreateComputerRecoveryToken(ctx).XApiVersion(xApiVersion).ComputerRecoveryTokenSpec(computerRecoveryTokenSpec).Execute()

Create Recovery Token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    computerRecoveryTokenSpec := *openapiclient.NewComputerRecoveryTokenSpec([]string{"BackupIds_example"}, time.Now()) // ComputerRecoveryTokenSpec | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentsApi.CreateComputerRecoveryToken(context.Background()).XApiVersion(xApiVersion).ComputerRecoveryTokenSpec(computerRecoveryTokenSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.CreateComputerRecoveryToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateComputerRecoveryToken`: ComputerRecoveryTokenModel
    fmt.Fprintf(os.Stdout, "Response from `AgentsApi.CreateComputerRecoveryToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateComputerRecoveryTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **computerRecoveryTokenSpec** | [**ComputerRecoveryTokenSpec**](ComputerRecoveryTokenSpec.md) |  | 

### Return type

[**ComputerRecoveryTokenModel**](ComputerRecoveryTokenModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteComputerRecoveryToken

> map[string]interface{} DeleteComputerRecoveryToken(ctx, id).XApiVersion(xApiVersion).Execute()

Delete Recovery Token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the recovery token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentsApi.DeleteComputerRecoveryToken(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.DeleteComputerRecoveryToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteComputerRecoveryToken`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AgentsApi.DeleteComputerRecoveryToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the recovery token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteComputerRecoveryTokenRequest struct via the builder pattern


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


## GetAllComputerRecoveryTokens

> ComputerRecoveryTokenResult GetAllComputerRecoveryTokens(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).Execute()

Get All Recovery Tokens



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    skip := int32(56) // int32 | Number of recovery tokens to skip. (optional)
    limit := int32(56) // int32 | Maximum number of recovery tokens to return. (optional)
    orderColumn := openapiclient.EComputerRecoveryTokenFiltersOrderColumn("Name") // EComputerRecoveryTokenFiltersOrderColumn | Sorts recovery tokens by one of the parameters. (optional)
    orderAsc := true // bool | Sorts recovery tokens in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters recovery tokens by the `nameFilter` pattern. The pattern can match any repository parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentsApi.GetAllComputerRecoveryTokens(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.GetAllComputerRecoveryTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllComputerRecoveryTokens`: ComputerRecoveryTokenResult
    fmt.Fprintf(os.Stdout, "Response from `AgentsApi.GetAllComputerRecoveryTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllComputerRecoveryTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **skip** | **int32** | Number of recovery tokens to skip. | 
 **limit** | **int32** | Maximum number of recovery tokens to return. | 
 **orderColumn** | [**EComputerRecoveryTokenFiltersOrderColumn**](EComputerRecoveryTokenFiltersOrderColumn.md) | Sorts recovery tokens by one of the parameters. | 
 **orderAsc** | **bool** | Sorts recovery tokens in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **nameFilter** | **string** | Filters recovery tokens by the &#x60;nameFilter&#x60; pattern. The pattern can match any repository parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both. | 

### Return type

[**ComputerRecoveryTokenResult**](ComputerRecoveryTokenResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComputerRecoveryToken

> ComputerRecoveryTokenModel GetComputerRecoveryToken(ctx, id).XApiVersion(xApiVersion).Execute()

Get Recovery Token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the recovery token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentsApi.GetComputerRecoveryToken(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.GetComputerRecoveryToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComputerRecoveryToken`: ComputerRecoveryTokenModel
    fmt.Fprintf(os.Stdout, "Response from `AgentsApi.GetComputerRecoveryToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the recovery token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComputerRecoveryTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


### Return type

[**ComputerRecoveryTokenModel**](ComputerRecoveryTokenModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComputerRecoveryToken

> ComputerRecoveryTokenModel UpdateComputerRecoveryToken(ctx, id).XApiVersion(xApiVersion).ComputerRecoveryTokenModel(computerRecoveryTokenModel).Execute()

Edit Recovery Token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format&#58; `<version>-<revision>`. (default to "1.1-rev0")
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the recovery token.
    computerRecoveryTokenModel := *openapiclient.NewComputerRecoveryTokenModel("Id_example", "Name_example", time.Now()) // ComputerRecoveryTokenModel | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgentsApi.UpdateComputerRecoveryToken(context.Background(), id).XApiVersion(xApiVersion).ComputerRecoveryTokenModel(computerRecoveryTokenModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgentsApi.UpdateComputerRecoveryToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateComputerRecoveryToken`: ComputerRecoveryTokenModel
    fmt.Fprintf(os.Stdout, "Response from `AgentsApi.UpdateComputerRecoveryToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the recovery token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateComputerRecoveryTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **computerRecoveryTokenModel** | [**ComputerRecoveryTokenModel**](ComputerRecoveryTokenModel.md) |  | 

### Return type

[**ComputerRecoveryTokenModel**](ComputerRecoveryTokenModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

