# RequestFormatOptions

When set to JSON, the event is automatically formatted with required fields before sending. When set to Raw, only the event's `_raw` value is sent.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RequestFormatOptionsJSON

// Open enum: custom values can be created with a direct type cast
custom := components.RequestFormatOptions("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `RequestFormatOptionsJSON` | JSON                       |
| `RequestFormatOptionsRaw`  | raw                        |