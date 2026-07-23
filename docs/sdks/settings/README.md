# System.Settings

## Overview

### Available Operations

* [Restart](#restart) - Restart the Cribl server

## Restart

Restart the Cribl server.<br/><br/>This operation requires <code>system.restart</code> to be set to <code>api</code> in <code>cribl.yml</code>. If this setting is not configured, the request returns a <code>403</code> error.<br/><br/>Restarting the server causes a brief period of downtime while the process stops and restarts. All in-flight events are drained before the process exits. Use <code>POST /system/settings/reload</code> to apply configuration changes without a full restart.

### Example Usage

<!-- UsageSnippet language="go" operationID="createSystemSettingsRestart" method="post" path="/system/settings/restart" example="RestartSystemExamplesDefault" -->
```go
package main

import(
	"context"
	"os"
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/Cribl-Community/cribl-control-plane-sdk-go"
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

    res, err := s.System.Settings.Restart(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSystemRestartResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateSystemSettingsRestartResponse](../../models/operations/createsystemsettingsrestartresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401                | application/json   |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |