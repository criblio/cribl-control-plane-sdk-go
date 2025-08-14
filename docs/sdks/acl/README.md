# ACL
(*Groups.ACL*)

## Overview

### Available Operations

* [Get](#get) - Retrieve the Access Control List (ACL) for a Worker Group or Edge Fleet

## Get

ACL of members with permissions for resources in this Group

### Example Usage

<!-- UsageSnippet language="go" operationID="getGroupsAclById" method="get" path="/master/groups/{id}/acl" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Groups.ACL.Get(ctx, "<id>", operations.GetGroupsACLByIDTypeInsights.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |
| `id`                                                                                | *string*                                                                            | :heavy_check_mark:                                                                  | Group id                                                                            |
| `type_`                                                                             | [*operations.GetGroupsACLByIDType](../../models/operations/getgroupsaclbyidtype.md) | :heavy_minus_sign:                                                                  | resource type by which to filter access levels                                      |
| `opts`                                                                              | [][operations.Option](../../models/operations/option.md)                            | :heavy_minus_sign:                                                                  | The options for this request.                                                       |

### Response

**[*operations.GetGroupsACLByIDResponse](../../models/operations/getgroupsaclbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |