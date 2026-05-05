# OutputResponseSendLogsAs

The content type to use when sending logs

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseSendLogsAsText

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseSendLogsAs("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `OutputResponseSendLogsAsText` | text                           |
| `OutputResponseSendLogsAsJSON` | json                           |