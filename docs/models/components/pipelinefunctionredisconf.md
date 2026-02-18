# PipelineFunctionRedisConf


## Supported Types

### RedisAuthTypeNone

```go
pipelineFunctionRedisConf := components.CreatePipelineFunctionRedisConfNone(components.RedisAuthTypeNone{/* values here */})
```

### RedisAuthTypeManual

```go
pipelineFunctionRedisConf := components.CreatePipelineFunctionRedisConfManual(components.RedisAuthTypeManual{/* values here */})
```

### RedisAuthTypeCredentialsSecret

```go
pipelineFunctionRedisConf := components.CreatePipelineFunctionRedisConfCredentialsSecret(components.RedisAuthTypeCredentialsSecret{/* values here */})
```

### RedisAuthTypeTextSecret

```go
pipelineFunctionRedisConf := components.CreatePipelineFunctionRedisConfTextSecret(components.RedisAuthTypeTextSecret{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch pipelineFunctionRedisConf.Type {
	case components.PipelineFunctionRedisConfTypeNone:
		// pipelineFunctionRedisConf.RedisAuthTypeNone is populated
	case components.PipelineFunctionRedisConfTypeManual:
		// pipelineFunctionRedisConf.RedisAuthTypeManual is populated
	case components.PipelineFunctionRedisConfTypeCredentialsSecret:
		// pipelineFunctionRedisConf.RedisAuthTypeCredentialsSecret is populated
	case components.PipelineFunctionRedisConfTypeTextSecret:
		// pipelineFunctionRedisConf.RedisAuthTypeTextSecret is populated
	default:
		// Unknown type - use pipelineFunctionRedisConf.GetUnknownRaw() for raw JSON
}
```
