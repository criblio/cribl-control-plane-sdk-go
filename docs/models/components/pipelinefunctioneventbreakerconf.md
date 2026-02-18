# PipelineFunctionEventBreakerConf


## Supported Types

### EventBreakerExistingOrNewNew

```go
pipelineFunctionEventBreakerConf := components.CreatePipelineFunctionEventBreakerConfNew(components.EventBreakerExistingOrNewNew{/* values here */})
```

### EventBreakerExistingOrNewExisting

```go
pipelineFunctionEventBreakerConf := components.CreatePipelineFunctionEventBreakerConfExisting(components.EventBreakerExistingOrNewExisting{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch pipelineFunctionEventBreakerConf.Type {
	case components.PipelineFunctionEventBreakerConfTypeNew:
		// pipelineFunctionEventBreakerConf.EventBreakerExistingOrNewNew is populated
	case components.PipelineFunctionEventBreakerConfTypeExisting:
		// pipelineFunctionEventBreakerConf.EventBreakerExistingOrNewExisting is populated
	default:
		// Unknown type - use pipelineFunctionEventBreakerConf.GetUnknownRaw() for raw JSON
}
```
