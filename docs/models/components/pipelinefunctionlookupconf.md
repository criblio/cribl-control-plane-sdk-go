# PipelineFunctionLookupConf


## Supported Types

### LookupDbLookupTrue

```go
pipelineFunctionLookupConf := components.CreatePipelineFunctionLookupConfLookupDbLookupTrue(components.LookupDbLookupTrue{/* values here */})
```

### LookupDbLookupFalse

```go
pipelineFunctionLookupConf := components.CreatePipelineFunctionLookupConfLookupDbLookupFalse(components.LookupDbLookupFalse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch pipelineFunctionLookupConf.Type {
	case components.PipelineFunctionLookupConfTypeLookupDbLookupTrue:
		// pipelineFunctionLookupConf.LookupDbLookupTrue is populated
	case components.PipelineFunctionLookupConfTypeLookupDbLookupFalse:
		// pipelineFunctionLookupConf.LookupDbLookupFalse is populated
}
```
