# Os

Operating system metadata collected from the node.


## Supported Types

### NodeOsInfo

```go
os := components.CreateOsNodeOsInfo(components.NodeOsInfo{/* values here */})
```

### OsTypeHeartbeatMetadata

```go
os := components.CreateOsOsTypeHeartbeatMetadata(components.OsTypeHeartbeatMetadata{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch os.Type {
	case components.OsTypeNodeOsInfo:
		// os.NodeOsInfo is populated
	case components.OsTypeOsTypeHeartbeatMetadata:
		// os.OsTypeHeartbeatMetadata is populated
}
```
