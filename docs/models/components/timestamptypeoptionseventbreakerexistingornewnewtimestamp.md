# TimestampTypeOptionsEventBreakerExistingOrNewNewTimestamp

Method to use for timestamp extraction. Use <code>auto</code> for automatic detection, <code>format</code> to specify a strptime format, or <code>current</code> to use the current system time.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.TimestampTypeOptionsEventBreakerExistingOrNewNewTimestampAuto

// Open enum: custom values can be created with a direct type cast
custom := components.TimestampTypeOptionsEventBreakerExistingOrNewNewTimestamp("custom_value")
```


## Values

| Name                                                               | Value                                                              |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `TimestampTypeOptionsEventBreakerExistingOrNewNewTimestampAuto`    | auto                                                               |
| `TimestampTypeOptionsEventBreakerExistingOrNewNewTimestampFormat`  | format                                                             |
| `TimestampTypeOptionsEventBreakerExistingOrNewNewTimestampCurrent` | current                                                            |