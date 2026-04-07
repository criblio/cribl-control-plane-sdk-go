# CreateOutputOutputWebhookFormat2

How to format events before sending out

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputOutputWebhookFormat2Ndjson

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputOutputWebhookFormat2("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `CreateOutputOutputWebhookFormat2Ndjson`    | ndjson                                      |
| `CreateOutputOutputWebhookFormat2JSONArray` | json_array                                  |
| `CreateOutputOutputWebhookFormat2Custom`    | custom                                      |
| `CreateOutputOutputWebhookFormat2Advanced`  | advanced                                    |