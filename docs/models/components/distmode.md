# DistMode

Distributed deployment mode for this instance.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.DistModeEdge

// Open enum: custom values can be created with a direct type cast
custom := components.DistMode("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `DistModeEdge`             | edge                       |
| `DistModeManagedEdge`      | managed-edge               |
| `DistModeMaster`           | master                     |
| `DistModeOutpost`          | outpost                    |
| `DistModeSearchSupervisor` | search-supervisor          |
| `DistModeSingle`           | single                     |
| `DistModeWorker`           | worker                     |