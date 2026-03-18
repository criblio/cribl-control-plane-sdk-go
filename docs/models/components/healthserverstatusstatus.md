# HealthServerStatusStatus

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.HealthServerStatusStatusShuttingDown

// Open enum: custom values can be created with a direct type cast
custom := components.HealthServerStatusStatus("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `HealthServerStatusStatusShuttingDown` | shutting down                          |
| `HealthServerStatusStatusHealthy`      | healthy                                |
| `HealthServerStatusStatusStandby`      | standby                                |