# CreateOutputBlobAccessTier

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputBlobAccessTierInferred

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputBlobAccessTier("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `CreateOutputBlobAccessTierInferred` | Inferred                             |
| `CreateOutputBlobAccessTierHot`      | Hot                                  |
| `CreateOutputBlobAccessTierCool`     | Cool                                 |
| `CreateOutputBlobAccessTierCold`     | Cold                                 |
| `CreateOutputBlobAccessTierArchive`  | Archive                              |