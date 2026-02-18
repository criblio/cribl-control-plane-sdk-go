# PipelineFunctionAggregateMetricsConf


## Supported Types

### AggregateMetricsCumulativeTrue

```go
pipelineFunctionAggregateMetricsConf := components.CreatePipelineFunctionAggregateMetricsConfAggregateMetricsCumulativeTrue(components.AggregateMetricsCumulativeTrue{/* values here */})
```

### AggregateMetricsCumulativeFalse

```go
pipelineFunctionAggregateMetricsConf := components.CreatePipelineFunctionAggregateMetricsConfAggregateMetricsCumulativeFalse(components.AggregateMetricsCumulativeFalse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch pipelineFunctionAggregateMetricsConf.Type {
	case components.PipelineFunctionAggregateMetricsConfTypeAggregateMetricsCumulativeTrue:
		// pipelineFunctionAggregateMetricsConf.AggregateMetricsCumulativeTrue is populated
	case components.PipelineFunctionAggregateMetricsConfTypeAggregateMetricsCumulativeFalse:
		// pipelineFunctionAggregateMetricsConf.AggregateMetricsCumulativeFalse is populated
}
```
