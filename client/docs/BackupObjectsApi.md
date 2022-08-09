# \BackupObjectsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllBackupObjects**](BackupObjectsApi.md#GetAllBackupObjects) | **Get** /api/v1/backupObjects | Get All Backup Objects
[**GetBackupObject**](BackupObjectsApi.md#GetBackupObject) | **Get** /api/v1/backupObjects/{id} | Get Backup Object
[**GetBackupObjectRestorePoints**](BackupObjectsApi.md#GetBackupObjectRestorePoints) | **Get** /api/v1/backupObjects/{id}/restorePoints | Get Restore Points



## GetAllBackupObjects

> BackupObjectsResult GetAllBackupObjects(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).PlatformNameFilter(platformNameFilter).PlatformIdFilter(platformIdFilter).TypeFilter(typeFilter).ViTypeFilter(viTypeFilter).Execute()

Get All Backup Objects



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
    skip := int32(56) // int32 | Number of backup objects to skip. (optional)
    limit := int32(56) // int32 | Maximum number of backup objects to return. (optional)
    orderColumn := openapiclient.EBackupObjectsFiltersOrderColumn("Name") // EBackupObjectsFiltersOrderColumn | Sorts backup objects by one of the backup object parameters. (optional)
    orderAsc := true // bool | Sorts backup objects in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters backup objects by the `nameFilter` pattern. The pattern can match any backup object parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. (optional)
    platformNameFilter := openapiclient.EPlatformType("VMware") // EPlatformType | Filters backup objects by platform ID. (optional)
    platformIdFilter := TODO // string | Filters backup objects by platform ID. (optional)
    typeFilter := "typeFilter_example" // string | Filters backup objects by object type. (optional)
    viTypeFilter := "viTypeFilter_example" // string | Filters backup objects by the type of VMware vSphere server. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupObjectsApi.GetAllBackupObjects(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).PlatformNameFilter(platformNameFilter).PlatformIdFilter(platformIdFilter).TypeFilter(typeFilter).ViTypeFilter(viTypeFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupObjectsApi.GetAllBackupObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllBackupObjects`: BackupObjectsResult
    fmt.Fprintf(os.Stdout, "Response from `BackupObjectsApi.GetAllBackupObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllBackupObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]
 **skip** | **int32** | Number of backup objects to skip. | 
 **limit** | **int32** | Maximum number of backup objects to return. | 
 **orderColumn** | [**EBackupObjectsFiltersOrderColumn**](EBackupObjectsFiltersOrderColumn.md) | Sorts backup objects by one of the backup object parameters. | 
 **orderAsc** | **bool** | Sorts backup objects in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **nameFilter** | **string** | Filters backup objects by the &#x60;nameFilter&#x60; pattern. The pattern can match any backup object parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | 
 **platformNameFilter** | [**EPlatformType**](EPlatformType.md) | Filters backup objects by platform ID. | 
 **platformIdFilter** | [**string**](string.md) | Filters backup objects by platform ID. | 
 **typeFilter** | **string** | Filters backup objects by object type. | 
 **viTypeFilter** | **string** | Filters backup objects by the type of VMware vSphere server. | 

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


## GetBackupObject

> BackupObjectModel GetBackupObject(ctx, id).XApiVersion(xApiVersion).Execute()

Get Backup Object



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
    id := TODO // string | ID of the backup object.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupObjectsApi.GetBackupObject(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupObjectsApi.GetBackupObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBackupObject`: BackupObjectModel
    fmt.Fprintf(os.Stdout, "Response from `BackupObjectsApi.GetBackupObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the backup object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]


### Return type

[**BackupObjectModel**](BackupObjectModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackupObjectRestorePoints

> ObjectRestorePointsResult GetBackupObjectRestorePoints(ctx, id).XApiVersion(xApiVersion).Execute()

Get Restore Points



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
    id := TODO // string | ID of the backup object.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BackupObjectsApi.GetBackupObjectRestorePoints(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupObjectsApi.GetBackupObjectRestorePoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBackupObjectRestorePoints`: ObjectRestorePointsResult
    fmt.Fprintf(os.Stdout, "Response from `BackupObjectsApi.GetBackupObjectRestorePoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the backup object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupObjectRestorePointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev2&quot;]


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

