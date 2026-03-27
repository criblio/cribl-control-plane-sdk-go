# BagExpansionMode

decides if bag-values are expanded to bags or arrays

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.BagExpansionModeBag

// Open enum: custom values can be created with a direct type cast
custom := components.BagExpansionMode("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `BagExpansionModeBag`   | bag                     |
| `BagExpansionModeArray` | array                   |