# SampleMode

Defines how sample rate will be derived: log(previousPeriodCount) or sqrt(previousPeriodCount)

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SampleModeLog

// Open enum: custom values can be created with a direct type cast
custom := components.SampleMode("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `SampleModeLog`  | log              |
| `SampleModeSqrt` | sqrt             |