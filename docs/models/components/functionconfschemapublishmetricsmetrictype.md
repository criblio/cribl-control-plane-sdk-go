# FunctionConfSchemaPublishMetricsMetricType

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.FunctionConfSchemaPublishMetricsMetricTypeCounter

// Open enum: custom values can be created with a direct type cast
custom := components.FunctionConfSchemaPublishMetricsMetricType("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `FunctionConfSchemaPublishMetricsMetricTypeCounter`      | counter                                                  |
| `FunctionConfSchemaPublishMetricsMetricTypeTimer`        | timer                                                    |
| `FunctionConfSchemaPublishMetricsMetricTypeGauge`        | gauge                                                    |
| `FunctionConfSchemaPublishMetricsMetricTypeDistribution` | distribution                                             |
| `FunctionConfSchemaPublishMetricsMetricTypeSummary`      | summary                                                  |
| `FunctionConfSchemaPublishMetricsMetricTypeHistogram`    | histogram                                                |