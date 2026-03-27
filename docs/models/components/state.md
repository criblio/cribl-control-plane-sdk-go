# State

State of the Job

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.StateInitializing

// Open enum: custom values can be created with a direct type cast
custom := components.State(999)
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `StateInitializing` | 0                   |
| `StatePending`      | 1                   |
| `StateRunning`      | 2                   |
| `StatePaused`       | 3                   |
| `StateCancelled`    | 4                   |
| `StateFinished`     | 5                   |
| `StateFailed`       | 6                   |
| `StateOrphaned`     | 7                   |
| `StateUnknown`      | 8                   |
| `StateLength`       | 9                   |