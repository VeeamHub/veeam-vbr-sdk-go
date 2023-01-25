# \RepositoriesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRepository**](RepositoriesApi.md#CreateRepository) | **Post** /api/v1/backupInfrastructure/repositories | Add Repository
[**CreateScaleOutRepository**](RepositoriesApi.md#CreateScaleOutRepository) | **Post** /api/v1/backupInfrastructure/scaleOutRepositories | Add Scale-Out Backup Repository
[**DeleteRepository**](RepositoriesApi.md#DeleteRepository) | **Delete** /api/v1/backupInfrastructure/repositories/{id} | Remove Repository
[**DeleteScaleOutRepository**](RepositoriesApi.md#DeleteScaleOutRepository) | **Delete** /api/v1/backupInfrastructure/scaleOutRepositories/{id} | Remove Scale-Out Backup Repository
[**DisableScaleOutExtentMaintenanceMode**](RepositoriesApi.md#DisableScaleOutExtentMaintenanceMode) | **Post** /api/v1/backupInfrastructure/scaleOutRepositories/{id}/disableMaintenanceMode | Disable Maintenance Mode
[**DisableScaleOutExtentSealedMode**](RepositoriesApi.md#DisableScaleOutExtentSealedMode) | **Post** /api/v1/backupInfrastructure/scaleOutRepositories/{id}/disableSealedMode | Disable Sealed Mode
[**EnableScaleOutExtentMaintenanceMode**](RepositoriesApi.md#EnableScaleOutExtentMaintenanceMode) | **Post** /api/v1/backupInfrastructure/scaleOutRepositories/{id}/enableMaintenanceMode | Enable Maintenance Mode
[**EnableScaleOutExtentSealedMode**](RepositoriesApi.md#EnableScaleOutExtentSealedMode) | **Post** /api/v1/backupInfrastructure/scaleOutRepositories/{id}/enableSealedMode | Enable Sealed Mode
[**EvacuateBackupsFromScaleOutExtent**](RepositoriesApi.md#EvacuateBackupsFromScaleOutExtent) | **Post** /api/v1/backupInfrastructure/scaleOutRepositories/{id}/evacuateBackups | Evacuate Backups from Performance Extent
[**GetAllRepositories**](RepositoriesApi.md#GetAllRepositories) | **Get** /api/v1/backupInfrastructure/repositories | Get All Repositories
[**GetAllRepositoriesStates**](RepositoriesApi.md#GetAllRepositoriesStates) | **Get** /api/v1/backupInfrastructure/repositories/states | Get All Repository States
[**GetAllScaleOutRepositories**](RepositoriesApi.md#GetAllScaleOutRepositories) | **Get** /api/v1/backupInfrastructure/scaleOutRepositories | Get All Scale-Out Backup Repositories
[**GetRepository**](RepositoriesApi.md#GetRepository) | **Get** /api/v1/backupInfrastructure/repositories/{id} | Get Repository
[**GetRepositoryAccessPermissions**](RepositoriesApi.md#GetRepositoryAccessPermissions) | **Get** /api/v1/backupInfrastructure/repositories/{id}/accessPermissions | Get Repository Access Permissions
[**GetScaleOutRepository**](RepositoriesApi.md#GetScaleOutRepository) | **Get** /api/v1/backupInfrastructure/scaleOutRepositories/{id} | Get Scale-Out Backup Repository
[**GetScaleOutRepositoryAccessPermissions**](RepositoriesApi.md#GetScaleOutRepositoryAccessPermissions) | **Get** /api/v1/backupInfrastructure/scaleOutRepositories/{id}/accessPermissions | Get Scale-Out Backup Repository Access Permissions
[**UpdateRepository**](RepositoriesApi.md#UpdateRepository) | **Put** /api/v1/backupInfrastructure/repositories/{id} | Edit Repository
[**UpdateRepositoryAccessPermissions**](RepositoriesApi.md#UpdateRepositoryAccessPermissions) | **Put** /api/v1/backupInfrastructure/repositories/{id}/accessPermissions | Edit Repository Access Permissions
[**UpdateScaleOutRepository**](RepositoriesApi.md#UpdateScaleOutRepository) | **Put** /api/v1/backupInfrastructure/scaleOutRepositories/{id} | Edit Scale-Out Backup Repository
[**UpdateScaleOutRepositoryAccessPermissions**](RepositoriesApi.md#UpdateScaleOutRepositoryAccessPermissions) | **Put** /api/v1/backupInfrastructure/scaleOutRepositories/{id}/accessPermissions | Edit Scale-Out Backup Repository Access Permissions



## CreateRepository

> SessionModel CreateRepository(ctx).XApiVersion(xApiVersion).RepositorySpec(repositorySpec).OverwriteOwner(overwriteOwner).Execute()

Add Repository



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
    repositorySpec := openapiclient.RepositorySpec{AmazonS3GlacierStorageSpec: openapiclient.NewAmazonS3GlacierStorageSpec(*openapiclient.NewAmazonS3StorageAccountModel("CredentialsId_example", openapiclient.EAmazonRegionType("China")), *openapiclient.NewAmazonS3GlacierStorageBucketModel("RegionId_example", "BucketName_example", "FolderName_example"), "HostId_example", *openapiclient.NewLinuxHardenedRepositorySettingsModel(), *openapiclient.NewMountServerSettingsModel("MountServerId_example", "WriteCacheFolder_example", false), *openapiclient.NewSmbRepositoryShareSettingsModel("SharePath_example", "CredentialsId_example"), *openapiclient.NewAzureArchiveStorageContainerModel("ContainerName_example", "FolderName_example"))} // RepositorySpec | 
    overwriteOwner := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.CreateRepository(context.Background()).XApiVersion(xApiVersion).RepositorySpec(repositorySpec).OverwriteOwner(overwriteOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.CreateRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRepository`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.CreateRepository`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **repositorySpec** | [**RepositorySpec**](RepositorySpec.md) |  | 
 **overwriteOwner** | **bool** |  | 

### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateScaleOutRepository

> SessionModel CreateScaleOutRepository(ctx).XApiVersion(xApiVersion).ScaleOutRepositorySpec(scaleOutRepositorySpec).Execute()

Add Scale-Out Backup Repository



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
    scaleOutRepositorySpec := *openapiclient.NewScaleOutRepositorySpec("Name_example", "Description_example", *openapiclient.NewPerformanceTierSpec([]openapiclient.PerformanceExtentSpec{*openapiclient.NewPerformanceExtentSpec("Id_example")})) // ScaleOutRepositorySpec | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.CreateScaleOutRepository(context.Background()).XApiVersion(xApiVersion).ScaleOutRepositorySpec(scaleOutRepositorySpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.CreateScaleOutRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateScaleOutRepository`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.CreateScaleOutRepository`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateScaleOutRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **scaleOutRepositorySpec** | [**ScaleOutRepositorySpec**](ScaleOutRepositorySpec.md) |  | 

### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRepository

> map[string]interface{} DeleteRepository(ctx, id).XApiVersion(xApiVersion).DeleteBackups(deleteBackups).Execute()

Remove Repository



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the backup repository.
    deleteBackups := true // bool | If *true*, Veeam Backup & Replication will remove backup files. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.DeleteRepository(context.Background(), id).XApiVersion(xApiVersion).DeleteBackups(deleteBackups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.DeleteRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRepository`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.DeleteRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **deleteBackups** | **bool** | If *true*, Veeam Backup &amp; Replication will remove backup files. | 

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


## DeleteScaleOutRepository

> map[string]interface{} DeleteScaleOutRepository(ctx, id).XApiVersion(xApiVersion).DeleteBackups(deleteBackups).Execute()

Remove Scale-Out Backup Repository



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the scale-out backup repository.
    deleteBackups := true // bool | If *true*, Veeam Backup & Replication will remove backup files. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.DeleteScaleOutRepository(context.Background(), id).XApiVersion(xApiVersion).DeleteBackups(deleteBackups).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.DeleteScaleOutRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteScaleOutRepository`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.DeleteScaleOutRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the scale-out backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScaleOutRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **deleteBackups** | **bool** | If *true*, Veeam Backup &amp; Replication will remove backup files. | 

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


## DisableScaleOutExtentMaintenanceMode

> SessionModel DisableScaleOutExtentMaintenanceMode(ctx, id).XApiVersion(xApiVersion).ScaleOutExtentMaintenanceSpec(scaleOutExtentMaintenanceSpec).Execute()

Disable Maintenance Mode



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the scale-out backup repository.
    scaleOutExtentMaintenanceSpec := *openapiclient.NewScaleOutExtentMaintenanceSpec([]string{"RepositoryIds_example"}) // ScaleOutExtentMaintenanceSpec | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.DisableScaleOutExtentMaintenanceMode(context.Background(), id).XApiVersion(xApiVersion).ScaleOutExtentMaintenanceSpec(scaleOutExtentMaintenanceSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.DisableScaleOutExtentMaintenanceMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableScaleOutExtentMaintenanceMode`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.DisableScaleOutExtentMaintenanceMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the scale-out backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableScaleOutExtentMaintenanceModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **scaleOutExtentMaintenanceSpec** | [**ScaleOutExtentMaintenanceSpec**](ScaleOutExtentMaintenanceSpec.md) |  | 

### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableScaleOutExtentSealedMode

> SessionModel DisableScaleOutExtentSealedMode(ctx, id).XApiVersion(xApiVersion).ScaleOutExtentMaintenanceSpec(scaleOutExtentMaintenanceSpec).Execute()

Disable Sealed Mode



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the scale-out backup repository.
    scaleOutExtentMaintenanceSpec := *openapiclient.NewScaleOutExtentMaintenanceSpec([]string{"RepositoryIds_example"}) // ScaleOutExtentMaintenanceSpec | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.DisableScaleOutExtentSealedMode(context.Background(), id).XApiVersion(xApiVersion).ScaleOutExtentMaintenanceSpec(scaleOutExtentMaintenanceSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.DisableScaleOutExtentSealedMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableScaleOutExtentSealedMode`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.DisableScaleOutExtentSealedMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the scale-out backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableScaleOutExtentSealedModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **scaleOutExtentMaintenanceSpec** | [**ScaleOutExtentMaintenanceSpec**](ScaleOutExtentMaintenanceSpec.md) |  | 

### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableScaleOutExtentMaintenanceMode

> SessionModel EnableScaleOutExtentMaintenanceMode(ctx, id).XApiVersion(xApiVersion).ScaleOutExtentMaintenanceSpec(scaleOutExtentMaintenanceSpec).Execute()

Enable Maintenance Mode



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the scale-out backup repository.
    scaleOutExtentMaintenanceSpec := *openapiclient.NewScaleOutExtentMaintenanceSpec([]string{"RepositoryIds_example"}) // ScaleOutExtentMaintenanceSpec | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.EnableScaleOutExtentMaintenanceMode(context.Background(), id).XApiVersion(xApiVersion).ScaleOutExtentMaintenanceSpec(scaleOutExtentMaintenanceSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.EnableScaleOutExtentMaintenanceMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableScaleOutExtentMaintenanceMode`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.EnableScaleOutExtentMaintenanceMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the scale-out backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableScaleOutExtentMaintenanceModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **scaleOutExtentMaintenanceSpec** | [**ScaleOutExtentMaintenanceSpec**](ScaleOutExtentMaintenanceSpec.md) |  | 

### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableScaleOutExtentSealedMode

> SessionModel EnableScaleOutExtentSealedMode(ctx, id).XApiVersion(xApiVersion).ScaleOutExtentMaintenanceSpec(scaleOutExtentMaintenanceSpec).Execute()

Enable Sealed Mode



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the scale-out backup repository.
    scaleOutExtentMaintenanceSpec := *openapiclient.NewScaleOutExtentMaintenanceSpec([]string{"RepositoryIds_example"}) // ScaleOutExtentMaintenanceSpec | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.EnableScaleOutExtentSealedMode(context.Background(), id).XApiVersion(xApiVersion).ScaleOutExtentMaintenanceSpec(scaleOutExtentMaintenanceSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.EnableScaleOutExtentSealedMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableScaleOutExtentSealedMode`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.EnableScaleOutExtentSealedMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the scale-out backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableScaleOutExtentSealedModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **scaleOutExtentMaintenanceSpec** | [**ScaleOutExtentMaintenanceSpec**](ScaleOutExtentMaintenanceSpec.md) |  | 

### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EvacuateBackupsFromScaleOutExtent

> SessionModel EvacuateBackupsFromScaleOutExtent(ctx, id).XApiVersion(xApiVersion).ScaleOutExtentMaintenanceSpec(scaleOutExtentMaintenanceSpec).Execute()

Evacuate Backups from Performance Extent



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the scale-out backup repository.
    scaleOutExtentMaintenanceSpec := *openapiclient.NewScaleOutExtentMaintenanceSpec([]string{"RepositoryIds_example"}) // ScaleOutExtentMaintenanceSpec | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.EvacuateBackupsFromScaleOutExtent(context.Background(), id).XApiVersion(xApiVersion).ScaleOutExtentMaintenanceSpec(scaleOutExtentMaintenanceSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.EvacuateBackupsFromScaleOutExtent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EvacuateBackupsFromScaleOutExtent`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.EvacuateBackupsFromScaleOutExtent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the scale-out backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvacuateBackupsFromScaleOutExtentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **scaleOutExtentMaintenanceSpec** | [**ScaleOutExtentMaintenanceSpec**](ScaleOutExtentMaintenanceSpec.md) |  | 

### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRepositories

> RepositoriesResult GetAllRepositories(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).TypeFilter(typeFilter).HostIdFilter(hostIdFilter).PathFilter(pathFilter).VmbApiFilter(vmbApiFilter).VmbApiPlatform(vmbApiPlatform).Execute()

Get All Repositories



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
    skip := int32(56) // int32 | Number of repositories to skip. (optional)
    limit := int32(56) // int32 | Maximum number of repositories to return. (optional)
    orderColumn := openapiclient.ERepositoryFiltersOrderColumn("Name") // ERepositoryFiltersOrderColumn | Sorts repositories by one of the repository parameters. (optional)
    orderAsc := true // bool | Sorts repositories in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters repositories by the `nameFilter` pattern. The pattern can match any repository parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. (optional)
    typeFilter := openapiclient.ERepositoryType("WinLocal") // ERepositoryType | Filters repositories by repository type. (optional)
    hostIdFilter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filters repositories by ID of the backup server. (optional)
    pathFilter := "pathFilter_example" // string | Filters repositories by path to the folder where backup files are stored. (optional)
    vmbApiFilter := "vmbApiFilter_example" // string | Filters repositories by VM Backup API parameters converted to the base64 string. To obtain the string, call the `GetApiProductInfoString` method of VM Backup API. (optional)
    vmbApiPlatform := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filters repositories by ID of a platform that you use to communicate with VM Backup API. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.GetAllRepositories(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).TypeFilter(typeFilter).HostIdFilter(hostIdFilter).PathFilter(pathFilter).VmbApiFilter(vmbApiFilter).VmbApiPlatform(vmbApiPlatform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.GetAllRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllRepositories`: RepositoriesResult
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.GetAllRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **skip** | **int32** | Number of repositories to skip. | 
 **limit** | **int32** | Maximum number of repositories to return. | 
 **orderColumn** | [**ERepositoryFiltersOrderColumn**](ERepositoryFiltersOrderColumn.md) | Sorts repositories by one of the repository parameters. | 
 **orderAsc** | **bool** | Sorts repositories in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **nameFilter** | **string** | Filters repositories by the &#x60;nameFilter&#x60; pattern. The pattern can match any repository parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | 
 **typeFilter** | [**ERepositoryType**](ERepositoryType.md) | Filters repositories by repository type. | 
 **hostIdFilter** | **string** | Filters repositories by ID of the backup server. | 
 **pathFilter** | **string** | Filters repositories by path to the folder where backup files are stored. | 
 **vmbApiFilter** | **string** | Filters repositories by VM Backup API parameters converted to the base64 string. To obtain the string, call the &#x60;GetApiProductInfoString&#x60; method of VM Backup API. | 
 **vmbApiPlatform** | **string** | Filters repositories by ID of a platform that you use to communicate with VM Backup API. | 

### Return type

[**RepositoriesResult**](RepositoriesResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRepositoriesStates

> RepositoryStatesResult GetAllRepositoriesStates(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).IdFilter(idFilter).NameFilter(nameFilter).TypeFilter(typeFilter).CapacityFilter(capacityFilter).FreeSpaceFilter(freeSpaceFilter).UsedSpaceFilter(usedSpaceFilter).Execute()

Get All Repository States



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
    skip := int32(56) // int32 | Number of repository states to skip. (optional)
    limit := int32(56) // int32 | Maximum number of repository states to return. (optional)
    orderColumn := openapiclient.ERepositoryStatesFiltersOrderColumn("Name") // ERepositoryStatesFiltersOrderColumn | Sorts repository states by one of the state parameters. (optional)
    orderAsc := true // bool | Sorts repository states in the ascending order by the `orderColumn` parameter. (optional)
    idFilter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filters repository states by repository ID. (optional)
    nameFilter := "nameFilter_example" // string | Filters repository states by the `nameFilter` pattern. The pattern can match any repository state parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. (optional)
    typeFilter := openapiclient.ERepositoryType("WinLocal") // ERepositoryType | Filters repository states by repository type. (optional)
    capacityFilter := float64(1.2) // float64 | Filters repository states by repository capacity. (optional)
    freeSpaceFilter := float64(1.2) // float64 | Filters repository states by repository free space. (optional)
    usedSpaceFilter := float64(1.2) // float64 | Filters repository states by repository used space. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.GetAllRepositoriesStates(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).IdFilter(idFilter).NameFilter(nameFilter).TypeFilter(typeFilter).CapacityFilter(capacityFilter).FreeSpaceFilter(freeSpaceFilter).UsedSpaceFilter(usedSpaceFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.GetAllRepositoriesStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllRepositoriesStates`: RepositoryStatesResult
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.GetAllRepositoriesStates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRepositoriesStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **skip** | **int32** | Number of repository states to skip. | 
 **limit** | **int32** | Maximum number of repository states to return. | 
 **orderColumn** | [**ERepositoryStatesFiltersOrderColumn**](ERepositoryStatesFiltersOrderColumn.md) | Sorts repository states by one of the state parameters. | 
 **orderAsc** | **bool** | Sorts repository states in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **idFilter** | **string** | Filters repository states by repository ID. | 
 **nameFilter** | **string** | Filters repository states by the &#x60;nameFilter&#x60; pattern. The pattern can match any repository state parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | 
 **typeFilter** | [**ERepositoryType**](ERepositoryType.md) | Filters repository states by repository type. | 
 **capacityFilter** | **float64** | Filters repository states by repository capacity. | 
 **freeSpaceFilter** | **float64** | Filters repository states by repository free space. | 
 **usedSpaceFilter** | **float64** | Filters repository states by repository used space. | 

### Return type

[**RepositoryStatesResult**](RepositoryStatesResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllScaleOutRepositories

> ScaleOutRepositoriesResult GetAllScaleOutRepositories(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).Execute()

Get All Scale-Out Backup Repositories



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
    skip := int32(56) // int32 | Number of repositories to skip. (optional)
    limit := int32(56) // int32 | Maximum number of repositories to return. (optional)
    orderColumn := openapiclient.EScaleOutRepositoryFiltersOrderColumn("Name") // EScaleOutRepositoryFiltersOrderColumn | Sorts repositories by one of the repository parameters. (optional)
    orderAsc := true // bool | Sorts repositories in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters repositories by the `nameFilter` pattern. The pattern can match any repository parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.GetAllScaleOutRepositories(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.GetAllScaleOutRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllScaleOutRepositories`: ScaleOutRepositoriesResult
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.GetAllScaleOutRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllScaleOutRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **skip** | **int32** | Number of repositories to skip. | 
 **limit** | **int32** | Maximum number of repositories to return. | 
 **orderColumn** | [**EScaleOutRepositoryFiltersOrderColumn**](EScaleOutRepositoryFiltersOrderColumn.md) | Sorts repositories by one of the repository parameters. | 
 **orderAsc** | **bool** | Sorts repositories in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **nameFilter** | **string** | Filters repositories by the &#x60;nameFilter&#x60; pattern. The pattern can match any repository parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both. | 

### Return type

[**ScaleOutRepositoriesResult**](ScaleOutRepositoriesResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository

> RepositoryModel GetRepository(ctx, id).XApiVersion(xApiVersion).Execute()

Get Repository



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the backup repository.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.GetRepository(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.GetRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepository`: RepositoryModel
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.GetRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


### Return type

[**RepositoryModel**](RepositoryModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryAccessPermissions

> RepositoryAccessPermissionsModel GetRepositoryAccessPermissions(ctx, id).XApiVersion(xApiVersion).Execute()

Get Repository Access Permissions



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the backup repository.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.GetRepositoryAccessPermissions(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.GetRepositoryAccessPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositoryAccessPermissions`: RepositoryAccessPermissionsModel
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.GetRepositoryAccessPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryAccessPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


### Return type

[**RepositoryAccessPermissionsModel**](RepositoryAccessPermissionsModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScaleOutRepository

> ScaleOutRepositoryModel GetScaleOutRepository(ctx, id).XApiVersion(xApiVersion).Execute()

Get Scale-Out Backup Repository



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the scale-out backup repository.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.GetScaleOutRepository(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.GetScaleOutRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScaleOutRepository`: ScaleOutRepositoryModel
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.GetScaleOutRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the scale-out backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScaleOutRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


### Return type

[**ScaleOutRepositoryModel**](ScaleOutRepositoryModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScaleOutRepositoryAccessPermissions

> RepositoryAccessPermissionsModel GetScaleOutRepositoryAccessPermissions(ctx, id).XApiVersion(xApiVersion).Execute()

Get Scale-Out Backup Repository Access Permissions



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the scale-out backup repository.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.GetScaleOutRepositoryAccessPermissions(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.GetScaleOutRepositoryAccessPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScaleOutRepositoryAccessPermissions`: RepositoryAccessPermissionsModel
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.GetScaleOutRepositoryAccessPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the scale-out backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScaleOutRepositoryAccessPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


### Return type

[**RepositoryAccessPermissionsModel**](RepositoryAccessPermissionsModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRepository

> SessionModel UpdateRepository(ctx, id).XApiVersion(xApiVersion).RepositoryModel(repositoryModel).Execute()

Edit Repository



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the backup repository.
    repositoryModel := openapiclient.RepositoryModel{AmazonS3GlacierStorageModel: openapiclient.NewAmazonS3GlacierStorageModel(*openapiclient.NewAmazonS3StorageAccountModel("CredentialsId_example", openapiclient.EAmazonRegionType("China")), *openapiclient.NewAmazonS3GlacierStorageBucketModel("RegionId_example", "BucketName_example", "FolderName_example"), "HostId_example", *openapiclient.NewLinuxHardenedRepositorySettingsModel(), *openapiclient.NewMountServerSettingsModel("MountServerId_example", "WriteCacheFolder_example", false), *openapiclient.NewSmbRepositoryShareSettingsModel("SharePath_example", "CredentialsId_example"), *openapiclient.NewAzureArchiveStorageContainerModel("ContainerName_example", "FolderName_example"))} // RepositoryModel | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.UpdateRepository(context.Background(), id).XApiVersion(xApiVersion).RepositoryModel(repositoryModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.UpdateRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRepository`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.UpdateRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **repositoryModel** | [**RepositoryModel**](RepositoryModel.md) |  | 

### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRepositoryAccessPermissions

> RepositoryAccessPermissionsModel UpdateRepositoryAccessPermissions(ctx, id).XApiVersion(xApiVersion).RepositoryAccessPermissionsModel(repositoryAccessPermissionsModel).Execute()

Edit Repository Access Permissions



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the backup repository.
    repositoryAccessPermissionsModel := *openapiclient.NewRepositoryAccessPermissionsModel(openapiclient.ERepositoryAccessType("DenyAll"), false) // RepositoryAccessPermissionsModel | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.UpdateRepositoryAccessPermissions(context.Background(), id).XApiVersion(xApiVersion).RepositoryAccessPermissionsModel(repositoryAccessPermissionsModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.UpdateRepositoryAccessPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRepositoryAccessPermissions`: RepositoryAccessPermissionsModel
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.UpdateRepositoryAccessPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryAccessPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **repositoryAccessPermissionsModel** | [**RepositoryAccessPermissionsModel**](RepositoryAccessPermissionsModel.md) |  | 

### Return type

[**RepositoryAccessPermissionsModel**](RepositoryAccessPermissionsModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScaleOutRepository

> SessionModel UpdateScaleOutRepository(ctx, id).XApiVersion(xApiVersion).ScaleOutRepositoryModel(scaleOutRepositoryModel).Execute()

Edit Scale-Out Backup Repository



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the scale-out backup repository.
    scaleOutRepositoryModel := *openapiclient.NewScaleOutRepositoryModel("Id_example", "Name_example", "Description_example", *openapiclient.NewPerformanceTierModel([]openapiclient.PerformanceExtentModel{*openapiclient.NewPerformanceExtentModel("Id_example", "Name_example")})) // ScaleOutRepositoryModel | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.UpdateScaleOutRepository(context.Background(), id).XApiVersion(xApiVersion).ScaleOutRepositoryModel(scaleOutRepositoryModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.UpdateScaleOutRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScaleOutRepository`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.UpdateScaleOutRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the scale-out backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScaleOutRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **scaleOutRepositoryModel** | [**ScaleOutRepositoryModel**](ScaleOutRepositoryModel.md) |  | 

### Return type

[**SessionModel**](SessionModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScaleOutRepositoryAccessPermissions

> RepositoryAccessPermissionsModel UpdateScaleOutRepositoryAccessPermissions(ctx, id).XApiVersion(xApiVersion).RepositoryAccessPermissionsModel(repositoryAccessPermissionsModel).Execute()

Edit Scale-Out Backup Repository Access Permissions



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the scale-out backup repository.
    repositoryAccessPermissionsModel := *openapiclient.NewRepositoryAccessPermissionsModel(openapiclient.ERepositoryAccessType("DenyAll"), false) // RepositoryAccessPermissionsModel | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesApi.UpdateScaleOutRepositoryAccessPermissions(context.Background(), id).XApiVersion(xApiVersion).RepositoryAccessPermissionsModel(repositoryAccessPermissionsModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesApi.UpdateScaleOutRepositoryAccessPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScaleOutRepositoryAccessPermissions`: RepositoryAccessPermissionsModel
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesApi.UpdateScaleOutRepositoryAccessPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the scale-out backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScaleOutRepositoryAccessPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **repositoryAccessPermissionsModel** | [**RepositoryAccessPermissionsModel**](RepositoryAccessPermissionsModel.md) |  | 

### Return type

[**RepositoryAccessPermissionsModel**](RepositoryAccessPermissionsModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

