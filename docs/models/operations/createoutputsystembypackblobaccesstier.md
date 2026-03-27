# CreateOutputSystemByPackBlobAccessTier

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackBlobAccessTierInferred

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackBlobAccessTier("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `CreateOutputSystemByPackBlobAccessTierInferred` | Inferred                                         |
| `CreateOutputSystemByPackBlobAccessTierHot`      | Hot                                              |
| `CreateOutputSystemByPackBlobAccessTierCool`     | Cool                                             |
| `CreateOutputSystemByPackBlobAccessTierCold`     | Cold                                             |
| `CreateOutputSystemByPackBlobAccessTierArchive`  | Archive                                          |