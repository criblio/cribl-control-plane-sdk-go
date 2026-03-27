# OsUnion


## Supported Types

### HostOsTypeHeartbeatMetadata

```go
osUnion := components.CreateOsUnionHostOsTypeHeartbeatMetadata(components.HostOsTypeHeartbeatMetadata{/* values here */})
```

### Os

```go
osUnion := components.CreateOsUnionOs(components.Os{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch osUnion.Type {
	case components.OsUnionTypeHostOsTypeHeartbeatMetadata:
		// osUnion.HostOsTypeHeartbeatMetadata is populated
	case components.OsUnionTypeOs:
		// osUnion.Os is populated
}
```
