# Health

Health status of the persistent queue.

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.HealthGreen

// Open enum: custom values can be created with a direct type cast
custom := components.Health("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `HealthGreen`   | Green           |
| `HealthRed`     | Red             |
| `HealthUnknown` | Unknown         |
| `HealthYellow`  | Yellow          |