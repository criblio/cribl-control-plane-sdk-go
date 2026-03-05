# ModeOptionsPq

With Smart mode, PQ will write events to the filesystem only when it detects backpressure from the processing engine. With Always On mode, PQ will always write events directly to the queue before forwarding them to the processing engine.

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