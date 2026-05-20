# OutputWebhookFormat2

How to format events before sending out

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputWebhookFormat2Ndjson

// Open enum: custom values can be created with a direct type cast
custom := components.OutputWebhookFormat2("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `OutputWebhookFormat2Ndjson`    | ndjson                          |
| `OutputWebhookFormat2JSONArray` | json_array                      |
| `OutputWebhookFormat2Custom`    | custom                          |
| `OutputWebhookFormat2Advanced`  | advanced                        |