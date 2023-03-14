# \TrafficRulesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllTrafficRules**](TrafficRulesApi.md#GetAllTrafficRules) | **Get** /api/v1/trafficRules | Get Traffic Rules
[**UpdateTrafficRules**](TrafficRulesApi.md#UpdateTrafficRules) | **Put** /api/v1/trafficRules | Edit Traffic Rules



## GetAllTrafficRules

> GlobalNetworkTrafficRulesModel GetAllTrafficRules(ctx).XApiVersion(xApiVersion).Execute()

Get Traffic Rules



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrafficRulesApi.GetAllTrafficRules(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrafficRulesApi.GetAllTrafficRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTrafficRules`: GlobalNetworkTrafficRulesModel
    fmt.Fprintf(os.Stdout, "Response from `TrafficRulesApi.GetAllTrafficRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTrafficRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

### Return type

[**GlobalNetworkTrafficRulesModel**](GlobalNetworkTrafficRulesModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrafficRules

> GlobalNetworkTrafficRulesModel UpdateTrafficRules(ctx).XApiVersion(xApiVersion).GlobalNetworkTrafficRulesModel(globalNetworkTrafficRulesModel).Execute()

Edit Traffic Rules



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
    globalNetworkTrafficRulesModel := *openapiclient.NewGlobalNetworkTrafficRulesModel(false) // GlobalNetworkTrafficRulesModel | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrafficRulesApi.UpdateTrafficRules(context.Background()).XApiVersion(xApiVersion).GlobalNetworkTrafficRulesModel(globalNetworkTrafficRulesModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrafficRulesApi.UpdateTrafficRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTrafficRules`: GlobalNetworkTrafficRulesModel
    fmt.Fprintf(os.Stdout, "Response from `TrafficRulesApi.UpdateTrafficRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrafficRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **globalNetworkTrafficRulesModel** | [**GlobalNetworkTrafficRulesModel**](GlobalNetworkTrafficRulesModel.md) |  | 

### Return type

[**GlobalNetworkTrafficRulesModel**](GlobalNetworkTrafficRulesModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

