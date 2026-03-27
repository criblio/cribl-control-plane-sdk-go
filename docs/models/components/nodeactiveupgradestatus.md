# NodeActiveUpgradeStatus

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.NodeActiveUpgradeStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.NodeActiveUpgradeStatus(999)
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `NodeActiveUpgradeStatusPending`   | 0                                  |
| `NodeActiveUpgradeStatusQueued`    | 1                                  |
| `NodeActiveUpgradeStatusUpgrading` | 2                                  |