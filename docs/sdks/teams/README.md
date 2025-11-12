# Teams
(*Groups.ACL.Teams*)

## Overview

### Available Operations

* [Get](#get) - Get the Access Control List for teams with permissions on a Worker Group or Edge Fleet for the specified Cribl product

## Get

Get the Access Control List (ACL) for teams that have permissions on a Worker Group or Edge Fleet for the specified Cribl product.

### Example Usage

<!-- UsageSnippet language="go" operationID="getConfigGroupAclTeamsByProductAndId" method="get" path="/products/{product}/groups/{id}/acl/teams" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Groups.ACL.Teams.Get(ctx, components.ProductsCoreEdge, "<id>", components.RbacResourceMacros.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedTeamAccessControlList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |
| `product`                                                                             | [components.ProductsCore](../../models/components/productscore.md)                    | :heavy_check_mark:                                                                    | Name of the Cribl product that contains the Worker Group or Edge Fleet.               |
| `id`                                                                                  | *string*                                                                              | :heavy_check_mark:                                                                    | The <code>id</code> of the Worker Group or Edge Fleet to get the team ACL for.        |
| `type_`                                                                               | [*components.RbacResource](../../models/components/rbacresource.md)                   | :heavy_minus_sign:                                                                    | Filter for limiting the response to ACL entries for the specified RBAC resource type. |
| `opts`                                                                                | [][operations.Option](../../models/operations/option.md)                              | :heavy_minus_sign:                                                                    | The options for this request.                                                         |

### Response

**[*operations.GetConfigGroupACLTeamsByProductAndIDResponse](../../models/operations/getconfiggroupaclteamsbyproductandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |