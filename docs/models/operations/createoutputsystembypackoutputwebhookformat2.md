# CreateOutputSystemByPackOutputWebhookFormat2

How to format events before sending out

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackOutputWebhookFormat2Ndjson

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackOutputWebhookFormat2("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `CreateOutputSystemByPackOutputWebhookFormat2Ndjson`    | ndjson                                                  |
| `CreateOutputSystemByPackOutputWebhookFormat2JSONArray` | json_array                                              |
| `CreateOutputSystemByPackOutputWebhookFormat2Custom`    | custom                                                  |
| `CreateOutputSystemByPackOutputWebhookFormat2Advanced`  | advanced                                                |