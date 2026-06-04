# Collector

Collector configuration


## Supported Types

### CollectorAzureBlob

```go
collector := components.CreateCollectorAzureBlob(components.CollectorAzureBlob{/* values here */})
```

### CollectorCriblLake

```go
collector := components.CreateCollectorCriblLake(components.CollectorCriblLake{/* values here */})
```

### CollectorDatabase

```go
collector := components.CreateCollectorDatabase(components.CollectorDatabase{/* values here */})
```

### CollectorFilesystem

```go
collector := components.CreateCollectorFilesystem(components.CollectorFilesystem{/* values here */})
```

### CollectorGoogleCloudStorage

```go
collector := components.CreateCollectorGoogleCloudStorage(components.CollectorGoogleCloudStorage{/* values here */})
```

### CollectorHealthCheck

```go
collector := components.CreateCollectorHealthCheck(components.CollectorHealthCheck{/* values here */})
```

### CollectorRest

```go
collector := components.CreateCollectorRest(components.CollectorRest{/* values here */})
```

### CollectorS3

```go
collector := components.CreateCollectorS3(components.CollectorS3{/* values here */})
```

### CollectorScript

```go
collector := components.CreateCollectorScript(components.CollectorScript{/* values here */})
```

### CollectorSplunk

```go
collector := components.CreateCollectorSplunk(components.CollectorSplunk{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch collector.Type {
	case components.CollectorUnionTypeAzureBlob:
		// collector.CollectorAzureBlob is populated
	case components.CollectorUnionTypeCriblLake:
		// collector.CollectorCriblLake is populated
	case components.CollectorUnionTypeDatabase:
		// collector.CollectorDatabase is populated
	case components.CollectorUnionTypeFilesystem:
		// collector.CollectorFilesystem is populated
	case components.CollectorUnionTypeGoogleCloudStorage:
		// collector.CollectorGoogleCloudStorage is populated
	case components.CollectorUnionTypeHealthCheck:
		// collector.CollectorHealthCheck is populated
	case components.CollectorUnionTypeRest:
		// collector.CollectorRest is populated
	case components.CollectorUnionTypeS3:
		// collector.CollectorS3 is populated
	case components.CollectorUnionTypeScript:
		// collector.CollectorScript is populated
	case components.CollectorUnionTypeSplunk:
		// collector.CollectorSplunk is populated
	default:
		// Unknown type - use collector.GetUnknownRaw() for raw JSON
}
```
