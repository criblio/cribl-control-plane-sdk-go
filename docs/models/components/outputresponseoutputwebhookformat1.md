# OutputResponseOutputWebhookFormat1

How to format events before sending out

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseOutputWebhookFormat1Ndjson

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseOutputWebhookFormat1("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `OutputResponseOutputWebhookFormat1Ndjson`    | ndjson                                        |
| `OutputResponseOutputWebhookFormat1JSONArray` | json_array                                    |
| `OutputResponseOutputWebhookFormat1Custom`    | custom                                        |
| `OutputResponseOutputWebhookFormat1Advanced`  | advanced                                      |