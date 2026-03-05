# CreateOutputSystemByPackFormatWebhook

How to format events before sending out

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackFormatWebhookNdjson

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackFormatWebhook("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `CreateOutputSystemByPackFormatWebhookNdjson`    | ndjson                                           |
| `CreateOutputSystemByPackFormatWebhookJSONArray` | json_array                                       |
| `CreateOutputSystemByPackFormatWebhookCustom`    | custom                                           |
| `CreateOutputSystemByPackFormatWebhookAdvanced`  | advanced                                         |