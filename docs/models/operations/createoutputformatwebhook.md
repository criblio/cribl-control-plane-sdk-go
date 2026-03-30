# CreateOutputFormatWebhook

How to format events before sending out

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputFormatWebhookNdjson

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputFormatWebhook("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `CreateOutputFormatWebhookNdjson`    | ndjson                               |
| `CreateOutputFormatWebhookJSONArray` | json_array                           |
| `CreateOutputFormatWebhookCustom`    | custom                               |
| `CreateOutputFormatWebhookAdvanced`  | advanced                             |