# CreateOutputSendLogsAs

The content type to use when sending logs

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSendLogsAsText

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSendLogsAs("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `CreateOutputSendLogsAsText` | text                         |
| `CreateOutputSendLogsAsJSON` | json                         |