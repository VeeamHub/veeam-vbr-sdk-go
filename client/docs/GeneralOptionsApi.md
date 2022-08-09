# \GeneralOptionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGeneralOptions**](GeneralOptionsApi.md#GetGeneralOptions) | **Get** /api/v1/generalOptions | Get General Options
[**UpdateGeneralOptions**](GeneralOptionsApi.md#UpdateGeneralOptions) | **Put** /api/v1/generalOptions | Edit General Options



## GetGeneralOptions

> GeneralOptionsModel GetGeneralOptions(ctx).XApiVersion(xApiVersion).Execute()

Get General Options



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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev2")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneralOptionsApi.GetGeneralOptions(context.Background()).XApiVersion(xApiVersion).Execute()
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
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]

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

Edit General Options



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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev2")
    generalOptionsModel := *openapiclient.NewGeneralOptionsModel() // GeneralOptionsModel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneralOptionsApi.UpdateGeneralOptions(context.Background()).XApiVersion(xApiVersion).GeneralOptionsModel(generalOptionsModel).Execute()
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
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]
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

