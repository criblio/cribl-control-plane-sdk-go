# OutputSentinelOneAISIEMAISIEMEndpointPath

Endpoint to send events to. Use /services/collector/event for structured JSON payloads with standard HEC top-level fields. Use /services/collector/raw for unstructured log lines (plain text).

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputSentinelOneAISIEMAISIEMEndpointPathRootServicesCollectorEvent

// Open enum: custom values can be created with a direct type cast
custom := components.OutputSentinelOneAISIEMAISIEMEndpointPath("custom_value")
```


## Values

| Name                                                                  | Value                                                                 |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `OutputSentinelOneAISIEMAISIEMEndpointPathRootServicesCollectorEvent` | /services/collector/event                                             |
| `OutputSentinelOneAISIEMAISIEMEndpointPathRootServicesCollectorRaw`   | /services/collector/raw                                               |