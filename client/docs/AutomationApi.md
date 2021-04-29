# \AutomationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportCredentials**](AutomationApi.md#ExportCredentials) | **Post** /api/v1/automation/credentials/export | Export Credentials
[**ExportEncryptionPasswords**](AutomationApi.md#ExportEncryptionPasswords) | **Post** /api/v1/automation/encryptionPasswords/export | Export Encryption Passwords
[**ExportJobs**](AutomationApi.md#ExportJobs) | **Post** /api/v1/automation/jobs/export | Export Jobs
[**ExportManagedServers**](AutomationApi.md#ExportManagedServers) | **Post** /api/v1/automation/managedServers/export | Export Servers
[**ExportProxies**](AutomationApi.md#ExportProxies) | **Post** /api/v1/automation/proxies/export | Export Proxies
[**ExportRepositories**](AutomationApi.md#ExportRepositories) | **Post** /api/v1/automation/repositories/export | Export Repositories
[**GetAllAutomationSessions**](AutomationApi.md#GetAllAutomationSessions) | **Get** /api/v1/automation/sessions | Get All Automation Sessions
[**GetAutomationSession**](AutomationApi.md#GetAutomationSession) | **Get** /api/v1/automation/sessions/{id} | Get Automation Session
[**GetAutomationSessionLogs**](AutomationApi.md#GetAutomationSessionLogs) | **Get** /api/v1/automation/sessions/{id}/logs | Get Automation Session Logs
[**ImportCredentials**](AutomationApi.md#ImportCredentials) | **Post** /api/v1/automation/credentials/import | Import Credentials
[**ImportEncryptionPasswords**](AutomationApi.md#ImportEncryptionPasswords) | **Post** /api/v1/automation/encryptionPasswords/import | Import Encryption Passwords
[**ImportJobs**](AutomationApi.md#ImportJobs) | **Post** /api/v1/automation/jobs/import | Import Jobs
[**ImportManagedServers**](AutomationApi.md#ImportManagedServers) | **Post** /api/v1/automation/managedServers/import | Import Servers
[**ImportProxies**](AutomationApi.md#ImportProxies) | **Post** /api/v1/automation/proxies/import | Import Proxies
[**ImportRepositories**](AutomationApi.md#ImportRepositories) | **Post** /api/v1/automation/repositories/import | Import Repositories
[**StopAutomationSession**](AutomationApi.md#StopAutomationSession) | **Post** /api/v1/automation/sessions/{id}/stop | Stop Automation Session



## ExportCredentials

> CredentialsImportSpecCollection ExportCredentials(ctx).XApiVersion(xApiVersion).CredentialsExportSpec(credentialsExportSpec).Execute()

Export Credentials



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
    credentialsExportSpec := *openapiclient.NewCredentialsExportSpec() // CredentialsExportSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ExportCredentials(context.Background()).XApiVersion(xApiVersion).CredentialsExportSpec(credentialsExportSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ExportCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportCredentials`: CredentialsImportSpecCollection
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ExportCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **credentialsExportSpec** | [**CredentialsExportSpec**](CredentialsExportSpec.md) |  | 

### Return type

[**CredentialsImportSpecCollection**](CredentialsImportSpecCollection.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportEncryptionPasswords

> EncryptionPasswordImportSpecCollection ExportEncryptionPasswords(ctx).XApiVersion(xApiVersion).EncryptionPasswordExportSpec(encryptionPasswordExportSpec).Execute()

Export Encryption Passwords



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
    encryptionPasswordExportSpec := *openapiclient.NewEncryptionPasswordExportSpec() // EncryptionPasswordExportSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ExportEncryptionPasswords(context.Background()).XApiVersion(xApiVersion).EncryptionPasswordExportSpec(encryptionPasswordExportSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ExportEncryptionPasswords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportEncryptionPasswords`: EncryptionPasswordImportSpecCollection
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ExportEncryptionPasswords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportEncryptionPasswordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **encryptionPasswordExportSpec** | [**EncryptionPasswordExportSpec**](EncryptionPasswordExportSpec.md) |  | 

### Return type

[**EncryptionPasswordImportSpecCollection**](EncryptionPasswordImportSpecCollection.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportJobs

> JobImportSpecCollection ExportJobs(ctx).XApiVersion(xApiVersion).JobExportSpec(jobExportSpec).Execute()

Export Jobs



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
    jobExportSpec := *openapiclient.NewJobExportSpec() // JobExportSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ExportJobs(context.Background()).XApiVersion(xApiVersion).JobExportSpec(jobExportSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ExportJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportJobs`: JobImportSpecCollection
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ExportJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **jobExportSpec** | [**JobExportSpec**](JobExportSpec.md) |  | 

### Return type

[**JobImportSpecCollection**](JobImportSpecCollection.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportManagedServers

> ManageServerImportSpecCollection ExportManagedServers(ctx).XApiVersion(xApiVersion).ManageServerExportSpec(manageServerExportSpec).Execute()

Export Servers



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
    manageServerExportSpec := *openapiclient.NewManageServerExportSpec() // ManageServerExportSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ExportManagedServers(context.Background()).XApiVersion(xApiVersion).ManageServerExportSpec(manageServerExportSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ExportManagedServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportManagedServers`: ManageServerImportSpecCollection
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ExportManagedServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportManagedServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **manageServerExportSpec** | [**ManageServerExportSpec**](ManageServerExportSpec.md) |  | 

### Return type

[**ManageServerImportSpecCollection**](ManageServerImportSpecCollection.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportProxies

> ProxyImportSpecCollection ExportProxies(ctx).XApiVersion(xApiVersion).ProxyExportSpec(proxyExportSpec).Execute()

Export Proxies



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
    proxyExportSpec := *openapiclient.NewProxyExportSpec() // ProxyExportSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ExportProxies(context.Background()).XApiVersion(xApiVersion).ProxyExportSpec(proxyExportSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ExportProxies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportProxies`: ProxyImportSpecCollection
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ExportProxies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportProxiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **proxyExportSpec** | [**ProxyExportSpec**](ProxyExportSpec.md) |  | 

### Return type

[**ProxyImportSpecCollection**](ProxyImportSpecCollection.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportRepositories

> RepositoryImportSpecCollection ExportRepositories(ctx).XApiVersion(xApiVersion).RepositoryExportSpec(repositoryExportSpec).Execute()

Export Repositories



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
    repositoryExportSpec := *openapiclient.NewRepositoryExportSpec() // RepositoryExportSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ExportRepositories(context.Background()).XApiVersion(xApiVersion).RepositoryExportSpec(repositoryExportSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ExportRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportRepositories`: RepositoryImportSpecCollection
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ExportRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **repositoryExportSpec** | [**RepositoryExportSpec**](RepositoryExportSpec.md) |  | 

### Return type

[**RepositoryImportSpecCollection**](RepositoryImportSpecCollection.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAutomationSessions

> SessionsResult GetAllAutomationSessions(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).CreatedAfterFilter(createdAfterFilter).CreatedBeforeFilter(createdBeforeFilter).EndedAfterFilter(endedAfterFilter).EndedBeforeFilter(endedBeforeFilter).TypeFilter(typeFilter).StateFilter(stateFilter).ResultFilter(resultFilter).JobIdFilter(jobIdFilter).Execute()

Get All Automation Sessions



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
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")
    skip := int32(56) // int32 | Number of sessions to skip. (optional)
    limit := int32(56) // int32 | Maximum number of sessions to return. (optional)
    orderColumn := openapiclient.ESessionsFiltersOrderColumn("Name") // ESessionsFiltersOrderColumn | Sorts sessions by one of the session parameters. (optional)
    orderAsc := true // bool | Sorts sessions in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters sessions by the `nameFilter` pattern. The pattern can match any session parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. (optional)
    createdAfterFilter := time.Now() // time.Time | Returns sessions that are created after the specified date and time. (optional)
    createdBeforeFilter := time.Now() // time.Time | Returns sessions that are created before the specified date and time. (optional)
    endedAfterFilter := time.Now() // time.Time | Returns sessions that are finished after the specified date and time. (optional)
    endedBeforeFilter := time.Now() // time.Time | Returns sessions that are finished before the specified date and time. (optional)
    typeFilter := openapiclient.ESessionType("Infrastructure") // ESessionType | Filters sessions by session type. (optional)
    stateFilter := openapiclient.ESessionState("Stopped") // ESessionState | Filters sessions by session state. (optional)
    resultFilter := openapiclient.ESessionResult("None") // ESessionResult | Filters sessions by session result. (optional)
    jobIdFilter := TODO // string | Filters sessions by job ID. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.GetAllAutomationSessions(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).CreatedAfterFilter(createdAfterFilter).CreatedBeforeFilter(createdBeforeFilter).EndedAfterFilter(endedAfterFilter).EndedBeforeFilter(endedBeforeFilter).TypeFilter(typeFilter).StateFilter(stateFilter).ResultFilter(resultFilter).JobIdFilter(jobIdFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.GetAllAutomationSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllAutomationSessions`: SessionsResult
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.GetAllAutomationSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAutomationSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **skip** | **int32** | Number of sessions to skip. | 
 **limit** | **int32** | Maximum number of sessions to return. | 
 **orderColumn** | [**ESessionsFiltersOrderColumn**](ESessionsFiltersOrderColumn.md) | Sorts sessions by one of the session parameters. | 
 **orderAsc** | **bool** | Sorts sessions in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **nameFilter** | **string** | Filters sessions by the &#x60;nameFilter&#x60; pattern. The pattern can match any session parameter. To substitute one or more characters, use the asterisk (*) character at the beginning and/or at the end. | 
 **createdAfterFilter** | **time.Time** | Returns sessions that are created after the specified date and time. | 
 **createdBeforeFilter** | **time.Time** | Returns sessions that are created before the specified date and time. | 
 **endedAfterFilter** | **time.Time** | Returns sessions that are finished after the specified date and time. | 
 **endedBeforeFilter** | **time.Time** | Returns sessions that are finished before the specified date and time. | 
 **typeFilter** | [**ESessionType**](ESessionType.md) | Filters sessions by session type. | 
 **stateFilter** | [**ESessionState**](ESessionState.md) | Filters sessions by session state. | 
 **resultFilter** | [**ESessionResult**](ESessionResult.md) | Filters sessions by session result. | 
 **jobIdFilter** | [**string**](string.md) | Filters sessions by job ID. | 

### Return type

[**SessionsResult**](SessionsResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationSession

> SessionModel GetAutomationSession(ctx, id).XApiVersion(xApiVersion).Execute()

Get Automation Session



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
    id := TODO // string | ID of the session.
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.GetAutomationSession(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.GetAutomationSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomationSession`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.GetAutomationSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]

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


## GetAutomationSessionLogs

> SessionLogResult GetAutomationSessionLogs(ctx, id).XApiVersion(xApiVersion).Execute()

Get Automation Session Logs



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
    id := TODO // string | ID of the session.
    xApiVersion := "xApiVersion_example" // string | Version and revision of the client REST API. Must be in the following format: *\\<version\\>-\\<revision\\>*.  (default to "1.0-rev1")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.GetAutomationSessionLogs(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.GetAutomationSessionLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomationSessionLogs`: SessionLogResult
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.GetAutomationSessionLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationSessionLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]

### Return type

[**SessionLogResult**](SessionLogResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCredentials

> SessionModel ImportCredentials(ctx).XApiVersion(xApiVersion).CredentialsImportSpecCollection(credentialsImportSpecCollection).Execute()

Import Credentials



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
    credentialsImportSpecCollection := *openapiclient.NewCredentialsImportSpecCollection([]openapiclient.CredentialsImportSpec{*openapiclient.NewCredentialsImportSpec("Username_example", "Tag_example", openapiclient.ECredentialsType("Standard"))}) // CredentialsImportSpecCollection | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ImportCredentials(context.Background()).XApiVersion(xApiVersion).CredentialsImportSpecCollection(credentialsImportSpecCollection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ImportCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportCredentials`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ImportCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **credentialsImportSpecCollection** | [**CredentialsImportSpecCollection**](CredentialsImportSpecCollection.md) |  | 

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


## ImportEncryptionPasswords

> SessionModel ImportEncryptionPasswords(ctx).XApiVersion(xApiVersion).EncryptionPasswordImportSpecCollection(encryptionPasswordImportSpecCollection).Execute()

Import Encryption Passwords



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
    encryptionPasswordImportSpecCollection := *openapiclient.NewEncryptionPasswordImportSpecCollection() // EncryptionPasswordImportSpecCollection | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ImportEncryptionPasswords(context.Background()).XApiVersion(xApiVersion).EncryptionPasswordImportSpecCollection(encryptionPasswordImportSpecCollection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ImportEncryptionPasswords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportEncryptionPasswords`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ImportEncryptionPasswords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportEncryptionPasswordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **encryptionPasswordImportSpecCollection** | [**EncryptionPasswordImportSpecCollection**](EncryptionPasswordImportSpecCollection.md) |  | 

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


## ImportJobs

> SessionModel ImportJobs(ctx).XApiVersion(xApiVersion).JobImportSpecCollection(jobImportSpecCollection).Execute()

Import Jobs



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
    jobImportSpecCollection := *openapiclient.NewJobImportSpecCollection([]openapiclient.JobImportSpec{*openapiclient.NewJobImportSpec("Name_example", "Description_example", false, openapiclient.EJobType("Backup"), *openapiclient.NewBackupJobVirtualMachinesSpec([]openapiclient.VmwareObjectModel{*openapiclient.NewVmwareObjectModel("HostName_example", "Name_example", openapiclient.EVmwareInventoryType("Unknown"))}), *openapiclient.NewBackupJobStorageImportModel(*openapiclient.NewBackupRepositoryImportModel("Name_example"), *openapiclient.NewBackupJobImportProxiesModel(false), *openapiclient.NewBackupJobRetentionPolicySettingsModel(openapiclient.ERetentionPolicyType("RestorePoints"), int32(123))), *openapiclient.NewBackupJobGuestProcessingImportModel(*openapiclient.NewBackupApplicationAwareProcessingImportModel(false), *openapiclient.NewGuestFileSystemIndexingModel(false)), *openapiclient.NewBackupScheduleModel(false))}) // JobImportSpecCollection | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ImportJobs(context.Background()).XApiVersion(xApiVersion).JobImportSpecCollection(jobImportSpecCollection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ImportJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportJobs`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ImportJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **jobImportSpecCollection** | [**JobImportSpecCollection**](JobImportSpecCollection.md) |  | 

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


## ImportManagedServers

> SessionModel ImportManagedServers(ctx).XApiVersion(xApiVersion).ManageServerImportSpecCollection(manageServerImportSpecCollection).Execute()

Import Servers



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
    manageServerImportSpecCollection := *openapiclient.NewManageServerImportSpecCollection() // ManageServerImportSpecCollection | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ImportManagedServers(context.Background()).XApiVersion(xApiVersion).ManageServerImportSpecCollection(manageServerImportSpecCollection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ImportManagedServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportManagedServers`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ImportManagedServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportManagedServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **manageServerImportSpecCollection** | [**ManageServerImportSpecCollection**](ManageServerImportSpecCollection.md) |  | 

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


## ImportProxies

> SessionModel ImportProxies(ctx).XApiVersion(xApiVersion).ProxyImportSpecCollection(proxyImportSpecCollection).Execute()

Import Proxies



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
    proxyImportSpecCollection := *openapiclient.NewProxyImportSpecCollection([]openapiclient.ProxyImportSpec{*openapiclient.NewProxyImportSpec("Description_example", openapiclient.EProxyType("ViProxy"), *openapiclient.NewProxyServerSettingsImportSpec("HostName_example"))}) // ProxyImportSpecCollection | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ImportProxies(context.Background()).XApiVersion(xApiVersion).ProxyImportSpecCollection(proxyImportSpecCollection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ImportProxies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportProxies`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ImportProxies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportProxiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **proxyImportSpecCollection** | [**ProxyImportSpecCollection**](ProxyImportSpecCollection.md) |  | 

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


## ImportRepositories

> SessionModel ImportRepositories(ctx).XApiVersion(xApiVersion).RepositoryImportSpecCollection(repositoryImportSpecCollection).Execute()

Import Repositories



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
    repositoryImportSpecCollection := *openapiclient.NewRepositoryImportSpecCollection() // RepositoryImportSpecCollection | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.ImportRepositories(context.Background()).XApiVersion(xApiVersion).RepositoryImportSpecCollection(repositoryImportSpecCollection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.ImportRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportRepositories`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.ImportRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]
 **repositoryImportSpecCollection** | [**RepositoryImportSpecCollection**](RepositoryImportSpecCollection.md) |  | 

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


## StopAutomationSession

> map[string]interface{} StopAutomationSession(ctx, id).XApiVersion(xApiVersion).Execute()

Stop Automation Session



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
    id := TODO // string | ID of the session.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutomationApi.StopAutomationSession(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomationApi.StopAutomationSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopAutomationSession`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AutomationApi.StopAutomationSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) | ID of the session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopAutomationSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format: *\\&lt;version\\&gt;-\\&lt;revision\\&gt;*.  | [default to &quot;1.0-rev1&quot;]


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

