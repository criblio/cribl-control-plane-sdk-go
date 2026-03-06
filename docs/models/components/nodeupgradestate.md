# NodeUpgradeState

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.NodeUpgradeStateActive

// Open enum: custom values can be created with a direct type cast
custom := components.NodeUpgradeState(999)
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `NodeUpgradeStateActive`  | 0                         |
| `NodeUpgradeStateCurrent` | 1                         |
| `NodeUpgradeStateFailed`  | 2                         |
| `NodeUpgradeStateSkipped` | 3                         |