# \ConfigurationBackupApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigBackupOptions**](ConfigurationBackupApi.md#GetConfigBackupOptions) | **Get** /api/v1/configBackup | Get Configuration Backup
[**StartConfigBackup**](ConfigurationBackupApi.md#StartConfigBackup) | **Post** /api/v1/configBackup/backup | Start Configuration Backup
[**UpdateConfigBackupOptions**](ConfigurationBackupApi.md#UpdateConfigBackupOptions) | **Put** /api/v1/configBackup | Edit Configuration Backup



## GetConfigBackupOptions

> ConfigBackupModel GetConfigBackupOptions(ctx).XApiVersion(xApiVersion).Execute()

Get Configuration Backup



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
    resp, r, err := apiClient.ConfigurationBackupApi.GetConfigBackupOptions(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationBackupApi.GetConfigBackupOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigBackupOptions`: ConfigBackupModel
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationBackupApi.GetConfigBackupOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigBackupOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

### Return type

[**ConfigBackupModel**](ConfigBackupModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartConfigBackup

> SessionModel StartConfigBackup(ctx).XApiVersion(xApiVersion).Execute()

Start Configuration Backup



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
    resp, r, err := apiClient.ConfigurationBackupApi.StartConfigBackup(context.Background()).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationBackupApi.StartConfigBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartConfigBackup`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationBackupApi.StartConfigBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartConfigBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigBackupOptions

> ConfigBackupModel UpdateConfigBackupOptions(ctx).XApiVersion(xApiVersion).ConfigBackupModel(configBackupModel).Execute()

Edit Configuration Backup



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
    configBackupModel := *openapiclient.NewConfigBackupModel(false, "BackupRepositoryId_example", int32(123), *openapiclient.NewConfigBackupNotificationsModel(false), *openapiclient.NewConfigBackupScheduleModel(false), *openapiclient.NewConfigBackupLastSuccessfulModel(), *openapiclient.NewConfigBackupEncryptionModel(false, "PasswordId_example")) // ConfigBackupModel | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationBackupApi.UpdateConfigBackupOptions(context.Background()).XApiVersion(xApiVersion).ConfigBackupModel(configBackupModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationBackupApi.UpdateConfigBackupOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfigBackupOptions`: ConfigBackupModel
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationBackupApi.UpdateConfigBackupOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigBackupOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **configBackupModel** | [**ConfigBackupModel**](ConfigBackupModel.md) |  | 

### Return type

[**ConfigBackupModel**](ConfigBackupModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

