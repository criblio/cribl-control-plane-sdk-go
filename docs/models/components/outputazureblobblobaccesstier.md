# OutputAzureBlobBlobAccessTier

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputAzureBlobBlobAccessTierInferred

// Open enum: custom values can be created with a direct type cast
custom := components.OutputAzureBlobBlobAccessTier("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `OutputAzureBlobBlobAccessTierInferred` | Inferred                                |
| `OutputAzureBlobBlobAccessTierHot`      | Hot                                     |
| `OutputAzureBlobBlobAccessTierCool`     | Cool                                    |
| `OutputAzureBlobBlobAccessTierCold`     | Cold                                    |
| `OutputAzureBlobBlobAccessTierArchive`  | Archive                                 |