# RedisDeploymentTypeSentinel


## Supported Types

### RedisDeploymentTypeSentinelTLSTrue

```go
redisDeploymentTypeSentinel := components.CreateRedisDeploymentTypeSentinelRedisDeploymentTypeSentinelTLSTrue(components.RedisDeploymentTypeSentinelTLSTrue{/* values here */})
```

### RedisDeploymentTypeSentinelTLSFalse

```go
redisDeploymentTypeSentinel := components.CreateRedisDeploymentTypeSentinelRedisDeploymentTypeSentinelTLSFalse(components.RedisDeploymentTypeSentinelTLSFalse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch redisDeploymentTypeSentinel.Type {
	case components.RedisDeploymentTypeSentinelTypeRedisDeploymentTypeSentinelTLSTrue:
		// redisDeploymentTypeSentinel.RedisDeploymentTypeSentinelTLSTrue is populated
	case components.RedisDeploymentTypeSentinelTypeRedisDeploymentTypeSentinelTLSFalse:
		// redisDeploymentTypeSentinel.RedisDeploymentTypeSentinelTLSFalse is populated
}
```
