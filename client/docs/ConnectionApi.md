# \ConnectionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConnectionCertificate**](ConnectionApi.md#GetConnectionCertificate) | **Post** /api/v1/connectionCertificate | Request TLS Certificate or SSH Fingerprint



## GetConnectionCertificate

> ConnectionCertificateModel GetConnectionCertificate(ctx).XApiVersion(xApiVersion).HostConnectionSpec(hostConnectionSpec).Execute()

Request TLS Certificate or SSH Fingerprint



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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    hostConnectionSpec := *openapiclient.NewHostConnectionSpec("ServerName_example", "CredentialsId_example", openapiclient.EManagedServerType("WindowsHost")) // HostConnectionSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionApi.GetConnectionCertificate(context.Background()).XApiVersion(xApiVersion).HostConnectionSpec(hostConnectionSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionApi.GetConnectionCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectionCertificate`: ConnectionCertificateModel
    fmt.Fprintf(os.Stdout, "Response from `ConnectionApi.GetConnectionCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **hostConnectionSpec** | [**HostConnectionSpec**](HostConnectionSpec.md) |  | 

### Return type

[**ConnectionCertificateModel**](ConnectionCertificateModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

