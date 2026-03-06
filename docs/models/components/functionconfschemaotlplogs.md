# FunctionConfSchemaOtlpLogs


## Supported Types

### OTLPLogsBatchOTLPLogsFalse

```go
functionConfSchemaOtlpLogs := components.CreateFunctionConfSchemaOtlpLogsOTLPLogsBatchOTLPLogsFalse(components.OTLPLogsBatchOTLPLogsFalse{/* values here */})
```

### OTLPLogsBatchOTLPLogsTrue

```go
functionConfSchemaOtlpLogs := components.CreateFunctionConfSchemaOtlpLogsOTLPLogsBatchOTLPLogsTrue(components.OTLPLogsBatchOTLPLogsTrue{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch functionConfSchemaOtlpLogs.Type {
	case components.FunctionConfSchemaOtlpLogsTypeOTLPLogsBatchOTLPLogsFalse:
		// functionConfSchemaOtlpLogs.OTLPLogsBatchOTLPLogsFalse is populated
	case components.FunctionConfSchemaOtlpLogsTypeOTLPLogsBatchOTLPLogsTrue:
		// functionConfSchemaOtlpLogs.OTLPLogsBatchOTLPLogsTrue is populated
}
```
