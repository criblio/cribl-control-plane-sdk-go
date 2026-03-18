# PipelineFunctionAggregateMetricsMetricType

The output metric type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.PipelineFunctionAggregateMetricsMetricTypeAutomatic

// Open enum: custom values can be created with a direct type cast
custom := components.PipelineFunctionAggregateMetricsMetricType("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `PipelineFunctionAggregateMetricsMetricTypeAutomatic`    | automatic                                                |
| `PipelineFunctionAggregateMetricsMetricTypeCounter`      | counter                                                  |
| `PipelineFunctionAggregateMetricsMetricTypeDistribution` | distribution                                             |
| `PipelineFunctionAggregateMetricsMetricTypeGauge`        | gauge                                                    |
| `PipelineFunctionAggregateMetricsMetricTypeHistogram`    | histogram                                                |
| `PipelineFunctionAggregateMetricsMetricTypeSummary`      | summary                                                  |
| `PipelineFunctionAggregateMetricsMetricTypeTimer`        | timer                                                    |