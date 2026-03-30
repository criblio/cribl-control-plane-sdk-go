# DisplayValue

ServiceNow reference field display mode. Allows raw values, display values, or both (sysparm_display_value).

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.DisplayValueFalse

// Open enum: custom values can be created with a direct type cast
custom := components.DisplayValue("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `DisplayValueFalse` | false               |
| `DisplayValueTrue`  | true                |
| `DisplayValueAll`   | all                 |