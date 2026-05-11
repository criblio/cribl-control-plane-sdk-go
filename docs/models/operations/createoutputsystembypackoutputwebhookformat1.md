# CreateOutputSystemByPackOutputWebhookFormat1

How to format events before sending out

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackOutputWebhookFormat1Ndjson

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackOutputWebhookFormat1("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `CreateOutputSystemByPackOutputWebhookFormat1Ndjson`    | ndjson                                                  |
| `CreateOutputSystemByPackOutputWebhookFormat1JSONArray` | json_array                                              |
| `CreateOutputSystemByPackOutputWebhookFormat1Custom`    | custom                                                  |
| `CreateOutputSystemByPackOutputWebhookFormat1Advanced`  | advanced                                                |