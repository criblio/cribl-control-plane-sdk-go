# CreateInputDisplayValue

ServiceNow reference field display mode. Allows raw values, display values, or both (sysparm_display_value).

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputDisplayValueFalse

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputDisplayValue("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `CreateInputDisplayValueFalse` | false                          |
| `CreateInputDisplayValueTrue`  | true                           |
| `CreateInputDisplayValueAll`   | all                            |