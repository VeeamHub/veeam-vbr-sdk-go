# \BackupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllBackups**](BackupsApi.md#GetAllBackups) | **Get** /api/v1/backups | Get All Backups
[**GetBackup**](BackupsApi.md#GetBackup) | **Get** /api/v1/backups/{id} | Get Backup
[**GetBackupObjects**](BackupsApi.md#GetBackupObjects) | **Get** /api/v1/backups/{id}/objects | Get Backup Objects



## GetAllBackups

> BackupsResult GetAllBackups(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).CreatedAfterFilter(createdAfterFilter).CreatedBeforeFilter(createdBeforeFilter).PlatformIdFilter(platformIdFilter).JobIdFilter(jobIdFilter).PolicyTagFilter(policyTagFilter).Execute()

Get All Backups



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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev2")
    skip := int32(56) // int32 | Number of backups to skip. (optional)
    limit := int32(56) // int32 | Maximum number of backups to return. (optional)
    orderColumn := openapiclient.EBackupsFiltersOrderColumn("Name") // EBackupsFiltersOrderColumn | Sorts backups by one of the backup parameters. (optional)
    orderAsc := true // bool | Sorts backups in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters backups by the `nameFilter` pattern. The pattern can match any backup parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. (optional)
    createdAfterFilter := time.Now() // time.Time | Returns backups that are created after the specified date and time. (optional)
    createdBeforeFilter := time.Now() // time.Time | Returns backups that are created before the specified date and time. (optional)
    platformIdFilter := TODO // string | Filters backups by ID of the backup object platform. (optional)
    jobIdFilter := TODO // string | Filters backups by ID of the parent job. (optional)
    policyTagFilter := "policyTagFilter_example" // string | Filters backups by retention policy tag. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.GetAllBackups(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).CreatedAfterFilter(createdAfterFilter).CreatedBeforeFilter(createdBeforeFilter).PlatformIdFilter(platformIdFilter).JobIdFilter(jobIdFilter).PolicyTagFilter(policyTagFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.GetAllBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllBackups`: BackupsResult
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.GetAllBackups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]
 **skip** | **int32** | Number of backups to skip. | 
 **limit** | **int32** | Maximum number of backups to return. | 
 **orderColumn** | [**EBackupsFiltersOrderColumn**](EBackupsFiltersOrderColumn.md) | Sorts backups by one of the backup parameters. | 
 **orderAsc** | **bool** | Sorts backups in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **nameFilter** | **string** | Filters backups by the &#x60;nameFilter&#x60; pattern. The pattern can match any backup parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | 
 **createdAfterFilter** | **time.Time** | Returns backups that are created after the specified date and time. | 
 **createdBeforeFilter** | **time.Time** | Returns backups that are created before the specified date and time. | 
 **platformIdFilter** | [**string**](string.md) | Filters backups by ID of the backup object platform. | 
 **jobIdFilter** | [**string**](string.md) | Filters backups by ID of the parent job. | 
 **policyTagFilter** | **string** | Filters backups by retention policy tag. | 

### Return type

[**BackupsResult**](BackupsResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackup

> BackupModel GetBackup(ctx, id).XApiVersion(xApiVersion).Execute()

Get Backup



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
    id := TODO // string | ID of the backup.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.GetBackup(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.GetBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBackup`: BackupModel
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.GetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the backup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]


### Return type

[**BackupModel**](BackupModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupObjects

> BackupObjectsResult GetBackupObjects(ctx, id).XApiVersion(xApiVersion).Execute()

Get Backup Objects



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
    id := TODO // string | ID of the backup.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupsApi.GetBackupObjects(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.GetBackupObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBackupObjects`: BackupObjectsResult
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.GetBackupObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the backup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]


### Return type

[**BackupObjectsResult**](BackupObjectsResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

