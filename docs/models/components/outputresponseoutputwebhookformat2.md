# OutputResponseOutputWebhookFormat2

How to format events before sending out

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseOutputWebhookFormat2Ndjson

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseOutputWebhookFormat2("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `OutputResponseOutputWebhookFormat2Ndjson`    | ndjson                                        |
| `OutputResponseOutputWebhookFormat2JSONArray` | json_array                                    |
| `OutputResponseOutputWebhookFormat2Custom`    | custom                                        |
| `OutputResponseOutputWebhookFormat2Advanced`  | advanced                                      |