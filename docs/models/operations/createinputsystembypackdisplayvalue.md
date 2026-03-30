# CreateInputSystemByPackDisplayValue

ServiceNow reference field display mode. Allows raw values, display values, or both (sysparm_display_value).

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackDisplayValueFalse

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackDisplayValue("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `CreateInputSystemByPackDisplayValueFalse` | false                                      |
| `CreateInputSystemByPackDisplayValueTrue`  | true                                       |
| `CreateInputSystemByPackDisplayValueAll`   | all                                        |