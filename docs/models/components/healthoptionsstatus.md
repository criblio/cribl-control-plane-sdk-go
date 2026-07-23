# HealthOptionsStatus

Overall health status of the Source or Destination.

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.HealthOptionsStatusGreen

// Open enum: custom values can be created with a direct type cast
custom := components.HealthOptionsStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `HealthOptionsStatusGreen`   | Green                        |
| `HealthOptionsStatusRed`     | Red                          |
| `HealthOptionsStatusUnknown` | Unknown                      |
| `HealthOptionsStatusYellow`  | Yellow                       |