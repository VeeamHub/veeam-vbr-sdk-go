package client_test

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	"github.com/veeamhub/veeam-vbr-sdk-go/v2/pkg/client"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
)

const vbrhost = "https://127.0.0.1:9419"
const uservbr = "user"
const passvbr = "pass"
const vbrapiver = "1.1-rev0"

func ExampleNewClientWithResponses() {
	cl, err := client.NewClientWithResponses(vbrhost)
	if err != nil {
		log.Fatal(err)
	}
	rsi, err := cl.GetServerInfoWithResponse(context.Background(), &client.GetServerInfoParams{
		XApiVersion: vbrapiver,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(rsi)
}

func ExampleNewClientWithResponses_nossl() {

	tlsClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	cl, err := client.NewClientWithResponses(vbrhost, client.WithHTTPClient(tlsClient))
	if err != nil {
		log.Fatal(err)
	}
	_ = cl
}

func ExampleNewClientWithResponses_auth() {
	cl, err := client.NewClientWithResponses(vbrhost)
	if err != nil {
		log.Fatal(err)
	}
	user := uservbr
	pass := passvbr
	rl, err := cl.CreateTokenWithFormdataBodyWithResponse(context.Background(), &client.CreateTokenParams{
		XApiVersion: vbrapiver,
	}, client.CreateTokenFormdataRequestBody{
		GrantType: "password",
		Username:  &user,
		Password:  &pass,
	})
	if err != nil {
		log.Fatal(err)
	}

	bearerTokenProvider, bearerTokenProviderErr := securityprovider.NewSecurityProviderBearerToken(rl.JSON200.AccessToken)
	if bearerTokenProviderErr != nil {
		panic(bearerTokenProviderErr)
	}

	authcl, err := client.NewClientWithResponses(vbrhost, client.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		log.Fatal(err)
	}
	_ = authcl
}

func ExampleClientWithResponses_GetServerInfoWithResponse() {
	cl, err := client.NewClientWithResponses(vbrhost)
	if err != nil {
		log.Fatal(err)
	}
	rsi, err := cl.GetServerInfoWithResponse(context.Background(), &client.GetServerInfoParams{
		XApiVersion: vbrapiver,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(rsi)
}

func ExampleClientWithResponses_GetAllRepositoriesWithResponse() {
	cl, err := client.NewClientWithResponses(vbrhost)
	if err != nil {
		log.Fatal(err)
	}
	user := uservbr
	pass := passvbr
	rl, err := cl.CreateTokenWithFormdataBodyWithResponse(context.Background(), &client.CreateTokenParams{
		XApiVersion: vbrapiver,
	}, client.CreateTokenFormdataRequestBody{
		GrantType: "password",
		Username:  &user,
		Password:  &pass,
	})
	if err != nil {
		log.Fatal(err)
	}

	bearerTokenProvider, bearerTokenProviderErr := securityprovider.NewSecurityProviderBearerToken(rl.JSON200.AccessToken)
	if bearerTokenProviderErr != nil {
		panic(bearerTokenProviderErr)
	}

	authcl, err := client.NewClientWithResponses(vbrhost, client.WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		log.Fatal(err)
	}

	nameFilter := "repname"
	rgr, err := authcl.GetAllRepositoriesWithResponse(context.Background(), &client.GetAllRepositoriesParams{
		XApiVersion: vbrapiver,
		NameFilter:  &nameFilter,
	})
	if err != nil {
		log.Fatal(err)
	}
	for rmi := range rgr.JSON200.Data {
		rm := rgr.JSON200.Data[rmi]
		log.Println(rm)
	}
}
