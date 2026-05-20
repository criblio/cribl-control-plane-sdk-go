# OutputDatadogSendLogsAs

The content type to use when sending logs

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputDatadogSendLogsAsText

// Open enum: custom values can be created with a direct type cast
custom := components.OutputDatadogSendLogsAs("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `OutputDatadogSendLogsAsText` | text                          |
| `OutputDatadogSendLogsAsJSON` | json                          |