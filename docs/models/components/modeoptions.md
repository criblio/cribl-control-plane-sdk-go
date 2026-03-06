# ModeOptions

In Error mode, PQ writes events to the filesystem if the Destination is unavailable. In Backpressure mode, PQ writes events to the filesystem when it detects backpressure from the Destination. In Always On mode, PQ always writes events to the filesystem.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ModeOptionsError

// Open enum: custom values can be created with a direct type cast
custom := components.ModeOptions("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ModeOptionsError`        | error                     |
| `ModeOptionsAlways`       | always                    |
| `ModeOptionsBackpressure` | backpressure              |