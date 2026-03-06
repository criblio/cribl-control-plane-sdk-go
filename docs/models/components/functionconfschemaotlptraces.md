# FunctionConfSchemaOtlpTraces


## Supported Types

### OTLPTracesBatchOTLPTracesFalse

```go
functionConfSchemaOtlpTraces := components.CreateFunctionConfSchemaOtlpTracesOTLPTracesBatchOTLPTracesFalse(components.OTLPTracesBatchOTLPTracesFalse{/* values here */})
```

### OTLPTracesBatchOTLPTracesTrue

```go
functionConfSchemaOtlpTraces := components.CreateFunctionConfSchemaOtlpTracesOTLPTracesBatchOTLPTracesTrue(components.OTLPTracesBatchOTLPTracesTrue{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch functionConfSchemaOtlpTraces.Type {
	case components.FunctionConfSchemaOtlpTracesTypeOTLPTracesBatchOTLPTracesFalse:
		// functionConfSchemaOtlpTraces.OTLPTracesBatchOTLPTracesFalse is populated
	case components.FunctionConfSchemaOtlpTracesTypeOTLPTracesBatchOTLPTracesTrue:
		// functionConfSchemaOtlpTraces.OTLPTracesBatchOTLPTracesTrue is populated
}
```
