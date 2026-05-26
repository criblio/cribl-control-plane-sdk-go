# OutputResponseBlobAccessTier

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseBlobAccessTierInferred

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseBlobAccessTier("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `OutputResponseBlobAccessTierInferred` | Inferred                               |
| `OutputResponseBlobAccessTierHot`      | Hot                                    |
| `OutputResponseBlobAccessTierCool`     | Cool                                   |
| `OutputResponseBlobAccessTierCold`     | Cold                                   |
| `OutputResponseBlobAccessTierArchive`  | Archive                                |