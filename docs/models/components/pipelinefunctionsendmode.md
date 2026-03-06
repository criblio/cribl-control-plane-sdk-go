# PipelineFunctionSendMode

In Sender mode, forwards search results directly to the destination. In Metrics mode, accumulates metrics from federated send operators, and forwards the aggregate metrics.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.PipelineFunctionSendModeSender

// Open enum: custom values can be created with a direct type cast
custom := components.PipelineFunctionSendMode("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PipelineFunctionSendModeSender`  | sender                            |
| `PipelineFunctionSendModeMetrics` | metrics                           |