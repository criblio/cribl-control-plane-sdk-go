# PipelineFunctionSerdeConf


## Supported Types

### SerdeTypeKvp

```go
pipelineFunctionSerdeConf := components.CreatePipelineFunctionSerdeConfKvp(components.SerdeTypeKvp{/* values here */})
```

### SerdeTypeDelim

```go
pipelineFunctionSerdeConf := components.CreatePipelineFunctionSerdeConfDelim(components.SerdeTypeDelim{/* values here */})
```

### SerdeTypeCsv

```go
pipelineFunctionSerdeConf := components.CreatePipelineFunctionSerdeConfCsv(components.SerdeTypeCsv{/* values here */})
```

### SerdeTypeJSON

```go
pipelineFunctionSerdeConf := components.CreatePipelineFunctionSerdeConfJSON(components.SerdeTypeJSON{/* values here */})
```

### SerdeTypeRegex

```go
pipelineFunctionSerdeConf := components.CreatePipelineFunctionSerdeConfRegex(components.SerdeTypeRegex{/* values here */})
```

### SerdeTypeGrok

```go
pipelineFunctionSerdeConf := components.CreatePipelineFunctionSerdeConfGrok(components.SerdeTypeGrok{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch pipelineFunctionSerdeConf.Type {
	case components.PipelineFunctionSerdeConfTypeKvp:
		// pipelineFunctionSerdeConf.SerdeTypeKvp is populated
	case components.PipelineFunctionSerdeConfTypeDelim:
		// pipelineFunctionSerdeConf.SerdeTypeDelim is populated
	case components.PipelineFunctionSerdeConfTypeCsv:
		// pipelineFunctionSerdeConf.SerdeTypeCsv is populated
	case components.PipelineFunctionSerdeConfTypeJSON:
		// pipelineFunctionSerdeConf.SerdeTypeJSON is populated
	case components.PipelineFunctionSerdeConfTypeRegex:
		// pipelineFunctionSerdeConf.SerdeTypeRegex is populated
	case components.PipelineFunctionSerdeConfTypeGrok:
		// pipelineFunctionSerdeConf.SerdeTypeGrok is populated
	default:
		// Unknown type - use pipelineFunctionSerdeConf.GetUnknownRaw() for raw JSON
}
```
