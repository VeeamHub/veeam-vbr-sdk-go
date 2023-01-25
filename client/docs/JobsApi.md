# \JobsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJob**](JobsApi.md#CreateJob) | **Post** /api/v1/jobs | Create Job
[**DeleteJob**](JobsApi.md#DeleteJob) | **Delete** /api/v1/jobs/{id} | Delete Job
[**DisableJob**](JobsApi.md#DisableJob) | **Post** /api/v1/jobs/{id}/disable | Disable Job
[**EnableJob**](JobsApi.md#EnableJob) | **Post** /api/v1/jobs/{id}/enable | Enable Job
[**GetAllJobs**](JobsApi.md#GetAllJobs) | **Get** /api/v1/jobs | Get All Jobs
[**GetAllJobsStates**](JobsApi.md#GetAllJobsStates) | **Get** /api/v1/jobs/states | Get All Job States
[**GetJob**](JobsApi.md#GetJob) | **Get** /api/v1/jobs/{id} | Get Job
[**StartJob**](JobsApi.md#StartJob) | **Post** /api/v1/jobs/{id}/start | Start Job
[**StopJob**](JobsApi.md#StopJob) | **Post** /api/v1/jobs/{id}/stop | Stop Job
[**UpdateJob**](JobsApi.md#UpdateJob) | **Put** /api/v1/jobs/{id} | Edit Job



## CreateJob

> JobModel CreateJob(ctx).XApiVersion(xApiVersion).JobSpec(jobSpec).Execute()

Create Job



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
    jobSpec := openapiclient.JobSpec{BackupJobSpec: openapiclient.NewBackupJobSpec(false, *openapiclient.NewBackupJobVirtualMachinesSpec([]openapiclient.VmwareObjectModel{*openapiclient.NewVmwareObjectModel("HostName_example", "Name_example", openapiclient.EVmwareInventoryType("Unknown"))}), *openapiclient.NewBackupJobStorageModel("BackupRepositoryId_example", *openapiclient.NewBackupProxiesSettingsModel(false), *openapiclient.NewBackupJobRetentionPolicySettingsModel(openapiclient.ERetentionPolicyType("RestorePoints"), int32(123))), *openapiclient.NewBackupJobGuestProcessingModel(*openapiclient.NewBackupApplicationAwareProcessingModel(false), *openapiclient.NewGuestFileSystemIndexingModel(false)), *openapiclient.NewBackupScheduleModel(false))} // JobSpec | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.CreateJob(context.Background()).XApiVersion(xApiVersion).JobSpec(jobSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.CreateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJob`: JobModel
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.CreateJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **jobSpec** | [**JobSpec**](JobSpec.md) |  | 

### Return type

[**JobModel**](JobModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJob

> map[string]interface{} DeleteJob(ctx, id).XApiVersion(xApiVersion).Execute()

Delete Job



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.DeleteJob(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.DeleteJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteJob`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.DeleteJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


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


## DisableJob

> map[string]interface{} DisableJob(ctx, id).XApiVersion(xApiVersion).Execute()

Disable Job



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.DisableJob(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.DisableJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableJob`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.DisableJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


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


## EnableJob

> map[string]interface{} EnableJob(ctx, id).XApiVersion(xApiVersion).Execute()

Enable Job



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.EnableJob(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.EnableJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableJob`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.EnableJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


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


## GetAllJobs

> JobsResult GetAllJobs(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).TypeFilter(typeFilter).Execute()

Get All Jobs



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
    skip := int32(56) // int32 | Number of jobs to skip. (optional)
    limit := int32(56) // int32 | Maximum number of jobs to return. (optional)
    orderColumn := openapiclient.EJobFiltersOrderColumn("Name") // EJobFiltersOrderColumn | Sorts jobs by one of the job parameters. (optional)
    orderAsc := true // bool | Sorts jobs in the ascending order by the `orderColumn` parameter. (optional)
    nameFilter := "nameFilter_example" // string | Filters jobs by the `nameFilter` pattern. The pattern can match any job parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both. (optional)
    typeFilter := openapiclient.EJobType("Backup") // EJobType | Filters jobs by job type. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.GetAllJobs(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).NameFilter(nameFilter).TypeFilter(typeFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetAllJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllJobs`: JobsResult
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetAllJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **skip** | **int32** | Number of jobs to skip. | 
 **limit** | **int32** | Maximum number of jobs to return. | 
 **orderColumn** | [**EJobFiltersOrderColumn**](EJobFiltersOrderColumn.md) | Sorts jobs by one of the job parameters. | 
 **orderAsc** | **bool** | Sorts jobs in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **nameFilter** | **string** | Filters jobs by the &#x60;nameFilter&#x60; pattern. The pattern can match any job parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both. | 
 **typeFilter** | [**EJobType**](EJobType.md) | Filters jobs by job type. | 

### Return type

[**JobsResult**](JobsResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllJobsStates

> JobStatesResult GetAllJobsStates(ctx).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).IdFilter(idFilter).NameFilter(nameFilter).TypeFilter(typeFilter).LastResultFilter(lastResultFilter).StatusFilter(statusFilter).WorkloadFilter(workloadFilter).LastRunAfterFilter(lastRunAfterFilter).LastRunBeforeFilter(lastRunBeforeFilter).IsHighPriorityJobFilter(isHighPriorityJobFilter).RepositoryIdFilter(repositoryIdFilter).ObjectsCountFilter(objectsCountFilter).Execute()

Get All Job States



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
    skip := int32(56) // int32 | Number of job states to skip. (optional)
    limit := int32(56) // int32 | Maximum number of job states to return. (optional)
    orderColumn := openapiclient.EJobStatesFiltersOrderColumn("Name") // EJobStatesFiltersOrderColumn | Sorts job states by one of the state parameters. (optional)
    orderAsc := true // bool | Sorts job states in the ascending order by the `orderColumn` parameter. (optional)
    idFilter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filters job states by job ID. (optional)
    nameFilter := "nameFilter_example" // string | Filters job states by the `nameFilter` pattern. The pattern can match any state parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both. (optional)
    typeFilter := openapiclient.EJobType("Backup") // EJobType | Filters job states by job type. (optional)
    lastResultFilter := openapiclient.ESessionResult("None") // ESessionResult | Filters job states by status with which jobs must finish. (optional)
    statusFilter := openapiclient.EJobStatus("running") // EJobStatus | Filters job states by current status of the job. (optional)
    workloadFilter := openapiclient.EJobWorkload("application") // EJobWorkload | Filters job states by workloads that jobs must process. (optional)
    lastRunAfterFilter := time.Now() // time.Time | Returns job states for jobs that have run after the specified date and time. (optional)
    lastRunBeforeFilter := time.Now() // time.Time | Returns job states for jobs that have not run after the specified date and time. (optional)
    isHighPriorityJobFilter := true // bool | If *true*, Returns job states for jobs with high priority. (optional)
    repositoryIdFilter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filters job states by repository ID. (optional)
    objectsCountFilter := int32(56) // int32 | Filters job states by number of objects processed by the job. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.GetAllJobsStates(context.Background()).XApiVersion(xApiVersion).Skip(skip).Limit(limit).OrderColumn(orderColumn).OrderAsc(orderAsc).IdFilter(idFilter).NameFilter(nameFilter).TypeFilter(typeFilter).LastResultFilter(lastResultFilter).StatusFilter(statusFilter).WorkloadFilter(workloadFilter).LastRunAfterFilter(lastRunAfterFilter).LastRunBeforeFilter(lastRunBeforeFilter).IsHighPriorityJobFilter(isHighPriorityJobFilter).RepositoryIdFilter(repositoryIdFilter).ObjectsCountFilter(objectsCountFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetAllJobsStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllJobsStates`: JobStatesResult
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetAllJobsStates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllJobsStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]
 **skip** | **int32** | Number of job states to skip. | 
 **limit** | **int32** | Maximum number of job states to return. | 
 **orderColumn** | [**EJobStatesFiltersOrderColumn**](EJobStatesFiltersOrderColumn.md) | Sorts job states by one of the state parameters. | 
 **orderAsc** | **bool** | Sorts job states in the ascending order by the &#x60;orderColumn&#x60; parameter. | 
 **idFilter** | **string** | Filters job states by job ID. | 
 **nameFilter** | **string** | Filters job states by the &#x60;nameFilter&#x60; pattern. The pattern can match any state parameter. To substitute one or more characters, use the asterisk (*) character at the beginning, at the end or both. | 
 **typeFilter** | [**EJobType**](EJobType.md) | Filters job states by job type. | 
 **lastResultFilter** | [**ESessionResult**](ESessionResult.md) | Filters job states by status with which jobs must finish. | 
 **statusFilter** | [**EJobStatus**](EJobStatus.md) | Filters job states by current status of the job. | 
 **workloadFilter** | [**EJobWorkload**](EJobWorkload.md) | Filters job states by workloads that jobs must process. | 
 **lastRunAfterFilter** | **time.Time** | Returns job states for jobs that have run after the specified date and time. | 
 **lastRunBeforeFilter** | **time.Time** | Returns job states for jobs that have not run after the specified date and time. | 
 **isHighPriorityJobFilter** | **bool** | If *true*, Returns job states for jobs with high priority. | 
 **repositoryIdFilter** | **string** | Filters job states by repository ID. | 
 **objectsCountFilter** | **int32** | Filters job states by number of objects processed by the job. | 

### Return type

[**JobStatesResult**](JobStatesResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJob

> JobModel GetJob(ctx, id).XApiVersion(xApiVersion).Execute()

Get Job



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.GetJob(context.Background(), id).XApiVersion(xApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJob`: JobModel
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]


### Return type

[**JobModel**](JobModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartJob

> SessionModel StartJob(ctx, id).XApiVersion(xApiVersion).JobStartSpec(jobStartSpec).Execute()

Start Job



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the job.
    jobStartSpec := *openapiclient.NewJobStartSpec(false) // JobStartSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.StartJob(context.Background(), id).XApiVersion(xApiVersion).JobStartSpec(jobStartSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.StartJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartJob`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.StartJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **jobStartSpec** | [**JobStartSpec**](JobStartSpec.md) |  | 

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


## StopJob

> SessionModel StopJob(ctx, id).XApiVersion(xApiVersion).JobStopSpec(jobStopSpec).Execute()

Stop Job



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the job.
    jobStopSpec := *openapiclient.NewJobStopSpec(false) // JobStopSpec |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.StopJob(context.Background(), id).XApiVersion(xApiVersion).JobStopSpec(jobStopSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.StopJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopJob`: SessionModel
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.StopJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **jobStopSpec** | [**JobStopSpec**](JobStopSpec.md) |  | 

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


## UpdateJob

> JobModel UpdateJob(ctx, id).XApiVersion(xApiVersion).JobModel(jobModel).Execute()

Edit Job



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the job.
    jobModel := openapiclient.JobModel{BackupJobModel: openapiclient.NewBackupJobModel(false, *openapiclient.NewBackupJobVirtualMachinesModel([]openapiclient.VmwareObjectSizeModel{*openapiclient.NewVmwareObjectSizeModel(*openapiclient.NewVmwareObjectModel("HostName_example", "Name_example", openapiclient.EVmwareInventoryType("Unknown")))}), *openapiclient.NewBackupJobStorageModel("BackupRepositoryId_example", *openapiclient.NewBackupProxiesSettingsModel(false), *openapiclient.NewBackupJobRetentionPolicySettingsModel(openapiclient.ERetentionPolicyType("RestorePoints"), int32(123))), *openapiclient.NewBackupJobGuestProcessingModel(*openapiclient.NewBackupApplicationAwareProcessingModel(false), *openapiclient.NewGuestFileSystemIndexingModel(false)), *openapiclient.NewBackupScheduleModel(false))} // JobModel | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.UpdateJob(context.Background(), id).XApiVersion(xApiVersion).JobModel(jobModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.UpdateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateJob`: JobModel
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.UpdateJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiVersion** | **string** | Version and revision of the client REST API. Must be in the following format&amp;#58; &#x60;&lt;version&gt;-&lt;revision&gt;&#x60;. | [default to &quot;1.1-rev0&quot;]

 **jobModel** | [**JobModel**](JobModel.md) |  | 

### Return type

[**JobModel**](JobModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

