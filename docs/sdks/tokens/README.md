# Auth.Tokens

## Overview

### Available Operations

* [Get](#get) - Log in and fetch an authentication token

## Get

This endpoint is unavailable on Cribl.Cloud.Instead, follow the instructions at https://docs.cribl.io/stream/api-tutorials/#criblcloud to get an Auth token for Cribl.Cloud.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAuthLogin" method="post" path="/auth/login" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
    )

    res, err := s.Auth.Tokens.Get(ctx, components.LoginInfo{
        Password: "6j50J9421x29IhO",
        Username: "Lilly_Weissnat",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [components.LoginInfo](../../models/components/logininfo.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |
| `opts`                                                       | [][operations.Option](../../models/operations/option.md)     | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.CreateAuthLoginResponse](../../models/operations/createauthloginresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |