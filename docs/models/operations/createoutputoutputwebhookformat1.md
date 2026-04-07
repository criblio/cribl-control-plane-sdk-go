# CreateOutputOutputWebhookFormat1

How to format events before sending out

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputOutputWebhookFormat1Ndjson

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputOutputWebhookFormat1("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `CreateOutputOutputWebhookFormat1Ndjson`    | ndjson                                      |
| `CreateOutputOutputWebhookFormat1JSONArray` | json_array                                  |
| `CreateOutputOutputWebhookFormat1Custom`    | custom                                      |
| `CreateOutputOutputWebhookFormat1Advanced`  | advanced                                    |