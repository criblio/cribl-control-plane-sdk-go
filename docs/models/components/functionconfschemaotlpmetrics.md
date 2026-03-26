# FunctionConfSchemaOtlpMetrics


## Supported Types

### OTLPMetricsBatchOTLPMetricsFalse

```go
functionConfSchemaOtlpMetrics := components.CreateFunctionConfSchemaOtlpMetricsOTLPMetricsBatchOTLPMetricsFalse(components.OTLPMetricsBatchOTLPMetricsFalse{/* values here */})
```

### OTLPMetricsBatchOTLPMetricsTrue

```go
functionConfSchemaOtlpMetrics := components.CreateFunctionConfSchemaOtlpMetricsOTLPMetricsBatchOTLPMetricsTrue(components.OTLPMetricsBatchOTLPMetricsTrue{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch functionConfSchemaOtlpMetrics.Type {
	case components.FunctionConfSchemaOtlpMetricsTypeOTLPMetricsBatchOTLPMetricsFalse:
		// functionConfSchemaOtlpMetrics.OTLPMetricsBatchOTLPMetricsFalse is populated
	case components.FunctionConfSchemaOtlpMetricsTypeOTLPMetricsBatchOTLPMetricsTrue:
		// functionConfSchemaOtlpMetrics.OTLPMetricsBatchOTLPMetricsTrue is populated
}
```
