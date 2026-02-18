# PipelineFunctionSerializeConf


## Supported Types

### SerializeTypeKvp

```go
pipelineFunctionSerializeConf := components.CreatePipelineFunctionSerializeConfKvp(components.SerializeTypeKvp{/* values here */})
```

### SerializeTypeDelim

```go
pipelineFunctionSerializeConf := components.CreatePipelineFunctionSerializeConfDelim(components.SerializeTypeDelim{/* values here */})
```

### SerializeTypeCsv

```go
pipelineFunctionSerializeConf := components.CreatePipelineFunctionSerializeConfCsv(components.SerializeTypeCsv{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch pipelineFunctionSerializeConf.Type {
	case components.PipelineFunctionSerializeConfTypeKvp:
		// pipelineFunctionSerializeConf.SerializeTypeKvp is populated
	case components.PipelineFunctionSerializeConfTypeDelim:
		// pipelineFunctionSerializeConf.SerializeTypeDelim is populated
	case components.PipelineFunctionSerializeConfTypeCsv:
		// pipelineFunctionSerializeConf.SerializeTypeCsv is populated
	default:
		// Unknown type - use pipelineFunctionSerializeConf.GetUnknownRaw() for raw JSON
}
```
