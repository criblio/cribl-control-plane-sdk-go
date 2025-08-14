# Teams
(*Groups.ACL.Teams*)

## Overview

### Available Operations

* [Get](#get) - Retrieve the Access Control List (ACL) for teams with permissions on a Worker Group or Edge Fleet for the specified Cribl product

## Get

ACL of team with permissions for resources in this Group

### Example Usage

<!-- UsageSnippet language="go" operationID="getProductsGroupsAclTeamsByProductAndId" method="get" path="/products/{product}/groups/{id}/acl/teams" -->
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

    res, err := s.Groups.ACL.Teams.Get(ctx, operations.GetProductsGroupsACLTeamsByProductAndIDProductStream, "<id>", operations.GetProductsGroupsACLTeamsByProductAndIDTypeDatasets.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `product`                                                                                                                              | [operations.GetProductsGroupsACLTeamsByProductAndIDProduct](../../models/operations/getproductsgroupsaclteamsbyproductandidproduct.md) | :heavy_check_mark:                                                                                                                     | Cribl Product                                                                                                                          |
| `id`                                                                                                                                   | *string*                                                                                                                               | :heavy_check_mark:                                                                                                                     | Group ID                                                                                                                               |
| `type_`                                                                                                                                | [*operations.GetProductsGroupsACLTeamsByProductAndIDType](../../models/operations/getproductsgroupsaclteamsbyproductandidtype.md)      | :heavy_minus_sign:                                                                                                                     | resource type by which to filter access levels                                                                                         |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.GetProductsGroupsACLTeamsByProductAndIDResponse](../../models/operations/getproductsgroupsaclteamsbyproductandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |