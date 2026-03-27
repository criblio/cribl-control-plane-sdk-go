# NodeSkippedUpgradeStatus

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.NodeSkippedUpgradeStatusDownloadError

// Open enum: custom values can be created with a direct type cast
custom := components.NodeSkippedUpgradeStatus(999)
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `NodeSkippedUpgradeStatusDownloadError`  | 0                                        |
| `NodeSkippedUpgradeStatusInstallType`    | 1                                        |
| `NodeSkippedUpgradeStatusMissingPackage` | 2                                        |
| `NodeSkippedUpgradeStatusTooOld`         | 3                                        |