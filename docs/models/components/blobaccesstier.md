# BlobAccessTier

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.BlobAccessTierInferred

// Open enum: custom values can be created with a direct type cast
custom := components.BlobAccessTier("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `BlobAccessTierInferred` | Inferred                 |
| `BlobAccessTierHot`      | Hot                      |
| `BlobAccessTierCool`     | Cool                     |
| `BlobAccessTierCold`     | Cold                     |
| `BlobAccessTierArchive`  | Archive                  |