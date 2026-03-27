# CountComparator

Operation to be applied over the results count

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.CountComparatorGreaterThan

// Open enum: custom values can be created with a direct type cast
custom := components.CountComparator("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `CountComparatorGreaterThan`      | >                                 |
| `CountComparatorLessThan`         | <                                 |
| `CountComparatorEqualEqualEqual`  | ===                               |
| `CountComparatorNotEqualEqual`    | !==                               |
| `CountComparatorGreaterThanEqual` | >=                                |
| `CountComparatorLessThanEqual`    | <=                                |