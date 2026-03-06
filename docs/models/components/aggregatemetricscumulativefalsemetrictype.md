# AggregateMetricsCumulativeFalseMetricType

The output metric type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AggregateMetricsCumulativeFalseMetricTypeAutomatic

// Open enum: custom values can be created with a direct type cast
custom := components.AggregateMetricsCumulativeFalseMetricType("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `AggregateMetricsCumulativeFalseMetricTypeAutomatic`    | automatic                                               |
| `AggregateMetricsCumulativeFalseMetricTypeCounter`      | counter                                                 |
| `AggregateMetricsCumulativeFalseMetricTypeDistribution` | distribution                                            |
| `AggregateMetricsCumulativeFalseMetricTypeGauge`        | gauge                                                   |
| `AggregateMetricsCumulativeFalseMetricTypeHistogram`    | histogram                                               |
| `AggregateMetricsCumulativeFalseMetricTypeSummary`      | summary                                                 |
| `AggregateMetricsCumulativeFalseMetricTypeTimer`        | timer                                                   |