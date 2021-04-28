# \RepositoriesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRepository**](RepositoriesApi.md#CreateRepository) | **Post** /api/v1/backupInfrastructure/repositories | Add Repository
[**DeleteRepository**](RepositoriesApi.md#DeleteRepository) | **Delete** /api/v1/backupInfrastructure/repositories/{id} | Remove Repository
[**DisableScaleOutExtentMaintenanceMode**](RepositoriesApi.md#DisableScaleOutExtentMaintenanceMode) | **Post** /api/v1/backupInfrastructure/scaleOutRepositories/{id}/disableMaintenanceMode | Disable Maintenance Mode
[**EnableScaleOutExtentMaintenanceMode**](RepositoriesApi.md#EnableScaleOutExtentMaintenanceMode) | **Post** /api/v1/backupInfrastructure/scaleOutRepositories/{id}/enableMaintenanceMode | Enable Maintenance Mode
[**GetAllRepositories**](RepositoriesApi.md#GetAllRepositories) | **Get** /api/v1/backupInfrastructure/repositories | Get All Repositories
[**GetAllRepositoriesStates**](RepositoriesApi.md#GetAllRepositoriesStates) | **Get** /api/v1/backupInfrastructure/repositories/states | Get All Repository States
[**GetAllScaleOutRepositories**](RepositoriesApi.md#GetAllScaleOutRepositories) | **Get** /api/v1/backupInfrastructure/scaleOutRepositories | Get All Scale-Out Backup Repositories
[**GetRepository**](RepositoriesApi.md#GetRepository) | **Get** /api/v1/backupInfrastructure/repositories/{id} | Get Repository
[**GetScaleOutRepository**](RepositoriesApi.md#GetScaleOutRepository) | **Get** /api/v1/backupInfrastructure/scaleOutRepositories/{id} | Get Scale-Out Backup Repository
[**UpdateRepository**](RepositoriesApi.md#UpdateRepository) | **Put** /api/v1/backupInfrastructure/repositories/{id} | Edit Repository



## CreateRepository

> SessionModel CreateRepository(ctx).XApiVersion(xApiVersion).RepositorySpec(repositorySpec).Execute()

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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    repositorySpec := *openapiclient.NewRepositorySpec("Name_example", "Description_example", openapiclient.ERepositoryType("WinLocal")) // RepositorySpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.CreateRepository(context.Background()).XApiVersion(xApiVersion).RepositorySpec(repositorySpec).Execute()
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
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **repositorySpec** | [**RepositorySpec**](RepositorySpec.md) |  | 

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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    id := TODO // string | ID of the backup repository.
    deleteBackups := true // bool | If *true*, Veeam Backup & Replication will remove backup files. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.DeleteRepository(context.Background(), id).XApiVersion(xApiVersion).DeleteBackups(deleteBackups).Execute()
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
**id** | [**string**](.md) | ID of the backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]

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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    id := TODO // string | ID of the scale-out backup repository extent.
    scaleOutExtentMaintenanceSpec := *openapiclient.NewScaleOutExtentMaintenanceSpec([]string{"RepositoryIds_example"}) // ScaleOutExtentMaintenanceSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.DisableScaleOutExtentMaintenanceMode(context.Background(), id).XApiVersion(xApiVersion).ScaleOutExtentMaintenanceSpec(scaleOutExtentMaintenanceSpec).Execute()
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
**id** | [**string**](.md) | ID of the scale-out backup repository extent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableScaleOutExtentMaintenanceModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]

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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    id := TODO // string | ID of the scale-out backup repository extent.
    scaleOutExtentMaintenanceSpec := *openapiclient.NewScaleOutExtentMaintenanceSpec([]string{"RepositoryIds_example"}) // ScaleOutExtentMaintenanceSpec | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.EnableScaleOutExtentMaintenanceMode(context.Background(), id).XApiVersion(xApiVersion).ScaleOutExtentMaintenanceSpec(scaleOutExtentMaintenanceSpec).Execute()
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
**id** | [**string**](.md) | ID of the scale-out backup repository extent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableScaleOutExtentMaintenanceModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]

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

> RepositoriesResult GetAllRepositories(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).TypeFilter(typeFilter).HostIdFilter(hostIdFilter).PathFilter(pathFilter).VmbApiFilter(vmbApiFilter).Execute()

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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    skip := int32(56) // int32 | Number of repositories to skip. (optional)
    limit := int32(56) // int32 | Maximum number of repositories to return. (optional)
    orderColumn := openapiclient.ERepositoryFiltersOrderColumn("Name") // ERepositoryFiltersOrderColumn | Sorts repositories by one of the repository parameters. (optional)
    orderAsc := true // bool | Sorts repositories in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters repositories by the `nameFilter` pattern. The pattern can match any repository parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. (optional)
    typeFilter := openapiclient.ERepositoryType("WinLocal") // ERepositoryType | Filters repositories by repository type. (optional)
    hostIdFilter := TODO // string | Filters repositories by ID of the backup server. (optional)
    pathFilter := "pathFilter_example" // string | Filters repositories by path to the folder where backup files are stored. (optional)
    vmbApiFilter := "vmbApiFilter_example" // string | Filters repositories by VM Backup API parameters converted to the base64 string.<br> To compose the base64 string:<br> <ol>   <li>Prepare VM Backup API parameters in the JSON format.</li>   <code>{<br>   \"protocolVersion\":\"string\",<br>   \"assemblyVersion\":\"string\",<br>   \"productId\":\"string\",<br>   \"versionFlags\":\"string\"<br>   }<br></code>   <li>Convert the JSON object into a string.</li>   <li>Encode the string with base64 encoding.</li> </ol>  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.GetAllRepositories(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).TypeFilter(typeFilter).HostIdFilter(hostIdFilter).PathFilter(pathFilter).VmbApiFilter(vmbApiFilter).Execute()
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
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **skip** | **int32** | Number of repositories to skip. | 
 **limit** | **int32** | Maximum number of repositories to return. | 
 **orderColumn** | [**ERepositoryFiltersOrderColumn**](ERepositoryFiltersOrderColumn.md) | Sorts repositories by one of the repository parameters. | 
 **orderAsc** | **bool** | Sorts repositories in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **nameFilter** | **string** | Filters repositories by the &#x60;nameFilter&#x60; pattern. The pattern can match any repository parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | 
 **typeFilter** | [**ERepositoryType**](ERepositoryType.md) | Filters repositories by repository type. | 
 **hostIdFilter** | [**string**](string.md) | Filters repositories by ID of the backup server. | 
 **pathFilter** | **string** | Filters repositories by path to the folder where backup files are stored. | 
 **vmbApiFilter** | **string** | Filters repositories by VM Backup API parameters converted to the base64 string.&lt;br&gt; To compose the base64 string:&lt;br&gt; &lt;ol&gt;   &lt;li&gt;Prepare VM Backup API parameters in the JSON format.&lt;/li&gt;   &lt;code&gt;{&lt;br&gt;   \&quot;protocolVersion\&quot;:\&quot;string\&quot;,&lt;br&gt;   \&quot;assemblyVersion\&quot;:\&quot;string\&quot;,&lt;br&gt;   \&quot;productId\&quot;:\&quot;string\&quot;,&lt;br&gt;   \&quot;versionFlags\&quot;:\&quot;string\&quot;&lt;br&gt;   }&lt;br&gt;&lt;/code&gt;   &lt;li&gt;Convert the JSON object into a string.&lt;/li&gt;   &lt;li&gt;Encode the string with base64 encoding.&lt;/li&gt; &lt;/ol&gt;  | 

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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    skip := int32(56) // int32 | Number of repository states to skip. (optional)
    limit := int32(56) // int32 | Maximum number of repository states to return. (optional)
    orderColumn := openapiclient.ERepositoryStatesFiltersOrderColumn("Name") // ERepositoryStatesFiltersOrderColumn | Sorts repository states by one of the state parameters. (optional)
    orderAsc := true // bool | Sorts repository states in the ascending order by the `orderColumn` parameter. (optional)
    idFilter := TODO // string | Filters repository states by repository ID. (optional)
    nameFilter := "nameFilter_example" // string | Filters repository states by the `nameFilter` pattern. The pattern can match any repository state parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. (optional)
    typeFilter := openapiclient.ERepositoryType("WinLocal") // ERepositoryType | Filters repository states by repository type. (optional)
    capacityFilter := float64(1.2) // float64 | Filters repository states by repository capacity. (optional)
    freeSpaceFilter := float64(1.2) // float64 | Filters repository states by repository free space. (optional)
    usedSpaceFilter := float64(1.2) // float64 | Filters repository states by repository used space. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.GetAllRepositoriesStates(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).IdFilter(idFilter).NameFilter(nameFilter).TypeFilter(typeFilter).CapacityFilter(capacityFilter).FreeSpaceFilter(freeSpaceFilter).UsedSpaceFilter(usedSpaceFilter).Execute()
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
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **skip** | **int32** | Number of repository states to skip. | 
 **limit** | **int32** | Maximum number of repository states to return. | 
 **orderColumn** | [**ERepositoryStatesFiltersOrderColumn**](ERepositoryStatesFiltersOrderColumn.md) | Sorts repository states by one of the state parameters. | 
 **orderAsc** | **bool** | Sorts repository states in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **idFilter** | [**string**](string.md) | Filters repository states by repository ID. | 
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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    skip := int32(56) // int32 | Number of repositories to skip. (optional)
    limit := int32(56) // int32 | Maximum number of repositories to return. (optional)
    orderColumn := openapiclient.EScaleOutRepositoryFiltersOrderColumn("Name") // EScaleOutRepositoryFiltersOrderColumn | Sorts repositories by one of the repository parameters. (optional)
    orderAsc := true // bool | Sorts repositories in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters repositories by the `nameFilter` substring. The substring can be part of any repository parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.GetAllScaleOutRepositories(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).Execute()
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
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **skip** | **int32** | Number of repositories to skip. | 
 **limit** | **int32** | Maximum number of repositories to return. | 
 **orderColumn** | [**EScaleOutRepositoryFiltersOrderColumn**](EScaleOutRepositoryFiltersOrderColumn.md) | Sorts repositories by one of the repository parameters. | 
 **orderAsc** | **bool** | Sorts repositories in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **nameFilter** | **string** | Filters repositories by the &#x60;nameFilter&#x60; substring. The substring can be part of any repository parameter. | 

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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    id := TODO // string | ID of the backup repository.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.GetRepository(context.Background(), id).XApiVersion(xApiVersion).Execute()
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
**id** | [**string**](.md) | ID of the backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]


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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    id := TODO // string | ID of the scale-out backup repository.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.GetScaleOutRepository(context.Background(), id).XApiVersion(xApiVersion).Execute()
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
**id** | [**string**](.md) | ID of the scale-out backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScaleOutRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]


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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    id := TODO // string | ID of the backup repository.
    repositoryModel := *openapiclient.NewRepositoryModel("Id_example", "Name_example", "Description_example", openapiclient.ERepositoryType("WinLocal")) // RepositoryModel | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RepositoriesApi.UpdateRepository(context.Background(), id).XApiVersion(xApiVersion).RepositoryModel(repositoryModel).Execute()
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
**id** | [**string**](.md) | ID of the backup repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]

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

