# OutputWebhookFormat1

How to format events before sending out

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputWebhookFormat1Ndjson

// Open enum: custom values can be created with a direct type cast
custom := components.OutputWebhookFormat1("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `OutputWebhookFormat1Ndjson`    | ndjson                          |
| `OutputWebhookFormat1JSONArray` | json_array                      |
| `OutputWebhookFormat1Custom`    | custom                          |
| `OutputWebhookFormat1Advanced`  | advanced                        |