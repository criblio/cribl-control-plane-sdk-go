# AggregateMetricsCumulativeTrueMetricType

The output metric type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AggregateMetricsCumulativeTrueMetricTypeAutomatic

// Open enum: custom values can be created with a direct type cast
custom := components.AggregateMetricsCumulativeTrueMetricType("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `AggregateMetricsCumulativeTrueMetricTypeAutomatic`    | automatic                                              |
| `AggregateMetricsCumulativeTrueMetricTypeCounter`      | counter                                                |
| `AggregateMetricsCumulativeTrueMetricTypeDistribution` | distribution                                           |
| `AggregateMetricsCumulativeTrueMetricTypeGauge`        | gauge                                                  |
| `AggregateMetricsCumulativeTrueMetricTypeHistogram`    | histogram                                              |
| `AggregateMetricsCumulativeTrueMetricTypeSummary`      | summary                                                |
| `AggregateMetricsCumulativeTrueMetricTypeTimer`        | timer                                                  |