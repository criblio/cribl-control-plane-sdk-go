# Remote


## Supported Types

### 

```go
remote := components.CreateRemoteStr(string{/* values here */})
```

### RemoteEnum

```go
remote := components.CreateRemoteRemoteEnum(components.RemoteEnum{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch remote.Type {
	case components.RemoteTypeStr:
		// remote.Str is populated
	case components.RemoteTypeRemoteEnum:
		// remote.RemoteEnum is populated
}
```
