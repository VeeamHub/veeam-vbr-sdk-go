# \ObjectRestorePointsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllObjectRestorePoints**](ObjectRestorePointsApi.md#GetAllObjectRestorePoints) | **Get** /api/v1/objectRestorePoints | Get All Restore Points
[**GetObjectRestorePoint**](ObjectRestorePointsApi.md#GetObjectRestorePoint) | **Get** /api/v1/objectRestorePoints/{id} | Get Restore Point
[**GetObjectRestorePointDisks**](ObjectRestorePointsApi.md#GetObjectRestorePointDisks) | **Get** /api/v1/objectRestorePoints/{id}/disks | Get Restore Point Disks



## GetAllObjectRestorePoints

> ObjectRestorePointsResult GetAllObjectRestorePoints(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).CreatedAfterFilter(createdAfterFilter).CreatedBeforeFilter(createdBeforeFilter).NameFilter(nameFilter).PlatformNameFilter(platformNameFilter).PlatformIdFilter(platformIdFilter).BackupIdFilter(backupIdFilter).BackupObjectIdFilter(backupObjectIdFilter).Execute()

Get All Restore Points



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
    skip := int32(56) // int32 | Number of restore points to skip. (optional)
    limit := int32(56) // int32 | Maximum number of restore points to return. (optional)
    orderColumn := openapiclient.EObjectRestorePointsFiltersOrderColumn("CreationTime") // EObjectRestorePointsFiltersOrderColumn | Sorts restore points by one of the restore point parameters. (optional)
    orderAsc := true // bool | Sorts restore points in the ascending order by the `orderColumn` parameter. (optional)
    createdAfterFilter := time.Now() // time.Time | Returns restore points that are created after the specified date and time. (optional)
    createdBeforeFilter := time.Now() // time.Time | Returns restore points that are created before the specified date and time. (optional)
    nameFilter := "nameFilter_example" // string | Filters restore points by the `nameFilter` pattern. The pattern can match any restore point parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. (optional)
    platformNameFilter := openapiclient.EPlatformType("VMware") // EPlatformType | Filters restore points by name of the backup object platform. (optional)
    platformIdFilter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filters restore points by ID of the backup object platform. (optional)
    backupIdFilter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filters restore points by backup ID. (optional)
    backupObjectIdFilter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filters restore points by backup object ID. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectRestorePointsApi.GetAllObjectRestorePoints(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).CreatedAfterFilter(createdAfterFilter).CreatedBeforeFilter(createdBeforeFilter).NameFilter(nameFilter).PlatformNameFilter(platformNameFilter).PlatformIdFilter(platformIdFilter).BackupIdFilter(backupIdFilter).BackupObjectIdFilter(backupObjectIdFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectRestorePointsApi.GetAllObjectRestorePoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllObjectRestorePoints`: ObjectRestorePointsResult
    fmt.Fprintf(os.Stdout, "Response from `ObjectRestorePointsApi.GetAllObjectRestorePoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllObjectRestorePointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **skip** | **int32** | Number of restore points to skip. | 
 **limit** | **int32** | Maximum number of restore points to return. | 
 **orderColumn** | [**EObjectRestorePointsFiltersOrderColumn**](EObjectRestorePointsFiltersOrderColumn.md) | Sorts restore points by one of the restore point parameters. | 
 **orderAsc** | **bool** | Sorts restore points in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **createdAfterFilter** | **time.Time** | Returns restore points that are created after the specified date and time. | 
 **createdBeforeFilter** | **time.Time** | Returns restore points that are created before the specified date and time. | 
 **nameFilter** | **string** | Filters restore points by the &#x60;nameFilter&#x60; pattern. The pattern can match any restore point parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | 
 **platformNameFilter** | [**EPlatformType**](EPlatformType.md) | Filters restore points by name of the backup object platform. | 
 **platformIdFilter** | **string** | Filters restore points by ID of the backup object platform. | 
 **backupIdFilter** | **string** | Filters restore points by backup ID. | 
 **backupObjectIdFilter** | **string** | Filters restore points by backup object ID. | 

### Return type

[**ObjectRestorePointsResult**](ObjectRestorePointsResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectRestorePoint

> ObjectRestorePointModel GetObjectRestorePoint(ctx, id).XApiVersion(xApiVersion).Execute()

Get Restore Point



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the restore point.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectRestorePointsApi.GetObjectRestorePoint(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectRestorePointsApi.GetObjectRestorePoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectRestorePoint`: ObjectRestorePointModel
    fmt.Fprintf(os.Stdout, "Response from `ObjectRestorePointsApi.GetObjectRestorePoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the restore point. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectRestorePointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


### Return type

[**ObjectRestorePointModel**](ObjectRestorePointModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectRestorePointDisks

> ObjectRestorePointDisksResult GetObjectRestorePointDisks(ctx, id).XApiVersion(xApiVersion).Execute()

Get Restore Point Disks



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the restore point.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectRestorePointsApi.GetObjectRestorePointDisks(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectRestorePointsApi.GetObjectRestorePointDisks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectRestorePointDisks`: ObjectRestorePointDisksResult
    fmt.Fprintf(os.Stdout, "Response from `ObjectRestorePointsApi.GetObjectRestorePointDisks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the restore point. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectRestorePointDisksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


### Return type

[**ObjectRestorePointDisksResult**](ObjectRestorePointDisksResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

