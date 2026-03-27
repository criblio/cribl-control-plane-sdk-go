# GaugeUpdate

The operation to use when rolling up gauge metrics. Defaults to last.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.GaugeUpdateLast

// Open enum: custom values can be created with a direct type cast
custom := components.GaugeUpdate("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `GaugeUpdateLast` | last              |
| `GaugeUpdateMax`  | max               |
| `GaugeUpdateMin`  | min               |
| `GaugeUpdateAvg`  | avg               |