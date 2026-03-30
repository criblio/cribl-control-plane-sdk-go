# HealthServerStatusStatus

Health state: <code>healthy</code>, <code>standby</code>, or <code>shutting down</code>.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.HealthServerStatusStatusHealthy

// Open enum: custom values can be created with a direct type cast
custom := components.HealthServerStatusStatus("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `HealthServerStatusStatusHealthy`      | healthy                                |
| `HealthServerStatusStatusShuttingDown` | shutting down                          |
| `HealthServerStatusStatusStandby`      | standby                                |