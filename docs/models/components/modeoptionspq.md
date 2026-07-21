# ModeOptionsPq

With Smart mode (deprecated), PQ will write events to the filesystem only when it detects backpressure from the processing engine. Smart mode will have no new development starting July 2026, followed by End of Support and feature removal (auto-migrating to Always On) in January 2027. We recommend using Always On mode instead. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ModeOptionsPqSmart

// Open enum: custom values can be created with a direct type cast
custom := components.ModeOptionsPq("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `ModeOptionsPqSmart`  | smart                 |
| `ModeOptionsPqAlways` | always                |