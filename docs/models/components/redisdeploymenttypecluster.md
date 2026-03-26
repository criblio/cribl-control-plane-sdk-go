# RedisDeploymentTypeCluster


## Supported Types

### RedisDeploymentTypeClusterTLSTrue

```go
redisDeploymentTypeCluster := components.CreateRedisDeploymentTypeClusterRedisDeploymentTypeClusterTLSTrue(components.RedisDeploymentTypeClusterTLSTrue{/* values here */})
```

### RedisDeploymentTypeClusterTLSFalse

```go
redisDeploymentTypeCluster := components.CreateRedisDeploymentTypeClusterRedisDeploymentTypeClusterTLSFalse(components.RedisDeploymentTypeClusterTLSFalse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch redisDeploymentTypeCluster.Type {
	case components.RedisDeploymentTypeClusterTypeRedisDeploymentTypeClusterTLSTrue:
		// redisDeploymentTypeCluster.RedisDeploymentTypeClusterTLSTrue is populated
	case components.RedisDeploymentTypeClusterTypeRedisDeploymentTypeClusterTLSFalse:
		// redisDeploymentTypeCluster.RedisDeploymentTypeClusterTLSFalse is populated
}
```
