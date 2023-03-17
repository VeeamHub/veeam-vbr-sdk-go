# \GeneralOptionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGeneralOptions**](GeneralOptionsApi.md#GetGeneralOptions) | **Get** /api/v1/generalOptions | Get Notification Settings
[**UpdateGeneralOptions**](GeneralOptionsApi.md#UpdateGeneralOptions) | **Put** /api/v1/generalOptions | Edit Notification Settings



## GetGeneralOptions

> GeneralOptionsModel GetGeneralOptions(ctx).XApiVersion(xApiVersion).Execute()

Get Notification Settings



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
    resp, r, err := apiClient.GeneralOptionsApi.GetGeneralOptions(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneralOptionsApi.GetGeneralOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGeneralOptions`: GeneralOptionsModel
    fmt.Fprintf(os.Stdout, "Response from `GeneralOptionsApi.GetGeneralOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGeneralOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

### Return type

[**GeneralOptionsModel**](GeneralOptionsModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGeneralOptions

> GeneralOptionsModel UpdateGeneralOptions(ctx).XApiVersion(xApiVersion).GeneralOptionsModel(generalOptionsModel).Execute()

Edit Notification Settings



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
    generalOptionsModel := *client.NewGeneralOptionsModel() // GeneralOptionsModel |

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)
    resp, r, err := apiClient.GeneralOptionsApi.UpdateGeneralOptions(context.Background()).XApiVersion(xApiVersion).GeneralOptionsModel(generalOptionsModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneralOptionsApi.UpdateGeneralOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGeneralOptions`: GeneralOptionsModel
    fmt.Fprintf(os.Stdout, "Response from `GeneralOptionsApi.UpdateGeneralOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGeneralOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **generalOptionsModel** | [**GeneralOptionsModel**](GeneralOptionsModel.md) |  | 

### Return type

[**GeneralOptionsModel**](GeneralOptionsModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

