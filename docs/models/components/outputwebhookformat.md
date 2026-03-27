# OutputWebhookFormat

How to format events before sending out

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputWebhookFormatNdjson

// Open enum: custom values can be created with a direct type cast
custom := components.OutputWebhookFormat("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `OutputWebhookFormatNdjson`    | ndjson                         |
| `OutputWebhookFormatJSONArray` | json_array                     |
| `OutputWebhookFormatCustom`    | custom                         |
| `OutputWebhookFormatAdvanced`  | advanced                       |