# SendLogsAs

The content type to use when sending logs

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SendLogsAsText

// Open enum: custom values can be created with a direct type cast
custom := components.SendLogsAs("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `SendLogsAsText` | text             |
| `SendLogsAsJSON` | json             |