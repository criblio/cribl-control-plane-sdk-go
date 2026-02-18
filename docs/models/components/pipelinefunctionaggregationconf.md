# PipelineFunctionAggregationConf


## Supported Types

### AggregationCumulativeTrue

```go
pipelineFunctionAggregationConf := components.CreatePipelineFunctionAggregationConfAggregationCumulativeTrue(components.AggregationCumulativeTrue{/* values here */})
```

### AggregationCumulativeFalse

```go
pipelineFunctionAggregationConf := components.CreatePipelineFunctionAggregationConfAggregationCumulativeFalse(components.AggregationCumulativeFalse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch pipelineFunctionAggregationConf.Type {
	case components.PipelineFunctionAggregationConfTypeAggregationCumulativeTrue:
		// pipelineFunctionAggregationConf.AggregationCumulativeTrue is populated
	case components.PipelineFunctionAggregationConfTypeAggregationCumulativeFalse:
		// pipelineFunctionAggregationConf.AggregationCumulativeFalse is populated
}
```
