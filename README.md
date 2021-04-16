# Veeam Backup & Replication SDK for Go

veeam-vbr-sdk-go is the unofficial Veeam Backup & Replication SDK for the Go programming language.

This SDK was generated using [go-swagger](https://github.com/go-swagger/go-swagger). As such, any API call in the [Veeam Backup & Replication v11 REST API](https://helpcenter.veeam.com/docs/backup/vbr_rest/overview.html?ver=110) is present in this SDK.

## 📗 Documentation

### Installation

Use `go get` to retrieve the SDK to add it to your `GOPATH` workspace, or
project's Go module dependencies.

```bash
go get github.com/veeamhub/veeam-vbr-sdk-go
```

To update the SDK use `go get -u` to retrieve the latest version of the SDK.

```bash
go get -u github.com/veeamhub/veeam-vbr-sdk-go
```

#### Dependencies

The SDK includes a `vendor` folder containing the runtime dependencies of the SDK. The metadata of the SDK's dependencies can be found in the Go module file `go.mod`.

### Example

This example shows a complete working Go file which will list the names of all Backup Jobs and demonstrate how to configure the timeout logic that will cancel the request if it takes too long. This example highlights how to login to a VBR server with trusted or self-signed certs, make a request, handle the error, and process the response.

```go
package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/veeamhub/veeam-vbr-sdk-go/client"
	"github.com/veeamhub/veeam-vbr-sdk-go/client/jobs"
	"github.com/veeamhub/veeam-vbr-sdk-go/client/login"
)

// Retrieves and prints out all Backup Job names of the specified VBR server.
//
// The Variables below need to be set according to your environment.
func main() {
	// Setting variables
	host := "vbr.contoso.local:9419"  // default API port 9419
	u := "contoso\\jsmith"  // VBR username
	p := strfmt.Password("password")  // VBR password
	timeout := 15 * time.Second  // 15 seconds

	// Using untrusted (self-signed) certificates
	skipTlsClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: timeout,
	}
	transport := httptransport.NewWithClient(host, "/", []string{"https"}, skipTlsClient)
	veeam := client.New(transport, nil)

	// Using trusted certificates
	// veeam := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
	//     Host:     host,
	//     BasePath: "/",
	//     Schemes:  []string{"https"},
	// })

	// Settings parameters
	params := login.NewCreateTokenParams().WithDefaults()
	params.Username = &u
	params.Password = &p

	// Authenticating to VBR API
	login, err := veeam.Login.CreateToken(params)
	if err != nil {
		panic(err)
	}
	token := *login.Payload.AccessToken

	// Setting authorization
	// From this point, all calls will require the 'auth' variable
	auth := httptransport.BearerToken(token)

	// Retrieving all backup jobs
	jobs, err := veeam.Jobs.GetAllJobs(
		jobs.NewGetAllJobsParams().WithDefaults(),
		auth)
	if err != nil {
		panic(err)
	}

	// Printing out job names
	// Working with the response payload
	fmt.Printf("Job Names:\n\n")
	for _, job := range jobs.Payload.Data {
		fmt.Println(*job.Name)
	}
}
```

## ✍ Contributions

We welcome contributions from the community! We encourage you to create [issues](https://github.com/VeeamHub/veeam-vbr-sdk-go/issues/new/choose) for Bugs & Feature Requests and submit Pull Requests. For more detailed information, refer to our [Contributing Guide](CONTRIBUTING.md).

## 🤝🏾 License

* [MIT License](LICENSE)

## 🤔 Questions

If you have any questions or something is unclear, please don't hesitate to [create an issue](https://github.com/VeeamHub/veeam-vbr-sdk-go/issues/new/choose) and let us know!
